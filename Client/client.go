package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/NiclasHjortkjaer/MockExam/Proto"
	"google.golang.org/grpc"
)

const (
	MainPort = ":5001"
)

var client Proto.ChatServiceClient
var wait *sync.WaitGroup

func init() {
	wait = &sync.WaitGroup{}
}

func main() {
	//init channel
	done := make(chan int)

	// Set up a connection to the server.
	conn, err := grpc.Dial(MainPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	client = Proto.NewChatServiceClient(conn)

	//Send messages
	wait.Add(1)

	go func() {
		defer wait.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			if strings.TrimSpace(scanner.Text()) != "" {
				input, err := strconv.Atoi(scanner.Text())
				if err != nil {
					log.Println("Please only input numbers")
				} else if input < 1 {
					log.Println("Only input positive integers")
				} else {
					msg := &Proto.Message{
						Value: int32(input),
					}

					response, err := client.IncrementValue(context.Background(), msg)
					if err != nil {
						log.Println("Something went wrong. Try again in a few seconds")
						log.Println("Try again in a few seconds")
					} else {
						log.Printf("Value before was: %d", response.Value)
					}
				}
			}
		}
	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done

}
