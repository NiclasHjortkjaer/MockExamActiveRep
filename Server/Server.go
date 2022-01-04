package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/NiclasHjortkjaer/MockExam/Proto"
	"google.golang.org/grpc"
)

const (
	MainPort       = 5001
	MainPortString = ":5001"
)

type Server struct {
	ID int
	Proto.UnimplementedChatServiceServer
	PrimePulse        time.Time
	lock              sync.Mutex
	Ports             []int32
	Port              int32
	Value             int32
	receivedNewLeader bool
	ServerTimestamp   Proto.LamportClock
}

func (s *Server) IncrementValue(ctx context.Context, msg *Proto.Message) (*Proto.Message, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	log.Print("__________________________________")
	log.Printf("Server recieved increment wish with a value of %d",
		msg.Value)

	returnMsg := Proto.Message{
		Value: s.Value,
	}

	s.Value += msg.Value
	log.Printf("Current value is now: %d", s.Value)

	return &returnMsg, nil
}

func main() {
	id := flag.Int("I", -1, "id")
	flag.Parse()

	portFile, err := os.Open("../ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(portFile)

	var ports []int32

	for scanner.Scan() {
		nextPort, _ := strconv.Atoi(scanner.Text())
		ports = append(ports, int32(nextPort))
	}

	s := &Server{
		ID:                             *id,
		UnimplementedChatServiceServer: Proto.UnimplementedChatServiceServer{},
		PrimePulse:                     time.Now(),
		lock:                           sync.Mutex{},
		Ports:                          ports,
		Port:                           ports[*id],
		Value:                          0,
		ServerTimestamp:                Proto.LamportClock{},
	}

	filename := "Server" + strconv.Itoa(s.ID) + "logs.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	grpcServer := grpc.NewServer()
	Proto.RegisterChatServiceServer(grpcServer, s)
	defer func() {
		lis.Close()
		log.Printf("Server stopped listening")
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d: %v", MainPort, err)
	}
}

/*
	LEADER ELECTION
*/
func (s *Server) LastLeader() {
	updatedPort := s.FindIndex(s.Port)
	s.Ports = removePort(s.Ports, int32(updatedPort))
	s.Port = MainPort

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	log.Printf("New server open at: %v", lis.Addr())

	Proto.RegisterChatServiceServer(grpcServer, s)
	go func(listener net.Listener) {
		defer func() {
			lis.Close()
			log.Printf("Server stopped listening")
		}()

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve server: %v", err)
		}
	}(lis)
}

/*
	HELPER METHODS
*/

func (s *Server) GetPort() int32 {
	return s.Port
}

func (s *Server) FindIndex(port int32) int {
	index := -1

	for i := 0; i < len(s.Ports); i++ {
		if s.Ports[i] == port {
			index = i
			break
		}
	}

	return index
}

func (s *Server) FindNextPort(index int) string {

	nextPort := s.Ports[(index+1)%len(s.Ports)]
	if nextPort == MainPort {
		nextPort = s.Ports[(index+2)%len(s.Ports)]
	}

	return ":" + strconv.Itoa(int(nextPort))
}

func removePort(s []int32, i int32) []int32 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
