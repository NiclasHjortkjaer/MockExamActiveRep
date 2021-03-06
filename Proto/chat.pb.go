// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/chat.proto

package Proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{1}
}

type PortsAndClocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfPorts  []int32  `protobuf:"varint,1,rep,packed,name=listOfPorts,proto3" json:"listOfPorts,omitempty"`
	ListOfClocks []uint32 `protobuf:"varint,2,rep,packed,name=listOfClocks,proto3" json:"listOfClocks,omitempty"`
}

func (x *PortsAndClocks) Reset() {
	*x = PortsAndClocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortsAndClocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortsAndClocks) ProtoMessage() {}

func (x *PortsAndClocks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortsAndClocks.ProtoReflect.Descriptor instead.
func (*PortsAndClocks) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{2}
}

func (x *PortsAndClocks) GetListOfPorts() []int32 {
	if x != nil {
		return x.ListOfPorts
	}
	return nil
}

func (x *PortsAndClocks) GetListOfClocks() []uint32 {
	if x != nil {
		return x.ListOfClocks
	}
	return nil
}

type ElectionPorts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfPorts []int32 `protobuf:"varint,1,rep,packed,name=listOfPorts,proto3" json:"listOfPorts,omitempty"`
}

func (x *ElectionPorts) Reset() {
	*x = ElectionPorts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionPorts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionPorts) ProtoMessage() {}

func (x *ElectionPorts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionPorts.ProtoReflect.Descriptor instead.
func (*ElectionPorts) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ElectionPorts) GetListOfPorts() []int32 {
	if x != nil {
		return x.ListOfPorts
	}
	return nil
}

var File_proto_chat_proto protoreflect.FileDescriptor

var file_proto_chat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x56, 0x0a, 0x0e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x69,
	0x73, 0x74, 0x4f, 0x66, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x31, 0x0a, 0x0d, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x32, 0xea, 0x01,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x0e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x0c, 0x52, 0x69, 0x6e, 0x67, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x41,
	0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x4e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x69, 0x63, 0x6c, 0x61, 0x73, 0x48, 0x6a, 0x6f, 0x72, 0x74, 0x6b, 0x6a, 0x61, 0x65,
	0x72, 0x2f, 0x4d, 0x6f, 0x63, 0x6b, 0x45, 0x78, 0x61, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chat_proto_rawDescOnce sync.Once
	file_proto_chat_proto_rawDescData = file_proto_chat_proto_rawDesc
)

func file_proto_chat_proto_rawDescGZIP() []byte {
	file_proto_chat_proto_rawDescOnce.Do(func() {
		file_proto_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chat_proto_rawDescData)
	})
	return file_proto_chat_proto_rawDescData
}

var file_proto_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_chat_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: proto.Message
	(*Void)(nil),           // 1: proto.Void
	(*PortsAndClocks)(nil), // 2: proto.PortsAndClocks
	(*ElectionPorts)(nil),  // 3: proto.ElectionPorts
}
var file_proto_chat_proto_depIdxs = []int32{
	0, // 0: proto.ChatService.IncrementValue:input_type -> proto.Message
	2, // 1: proto.ChatService.RingElection:input_type -> proto.PortsAndClocks
	1, // 2: proto.ChatService.SelectNewLeader:input_type -> proto.Void
	3, // 3: proto.ChatService.BroadcastNewLeader:input_type -> proto.ElectionPorts
	0, // 4: proto.ChatService.IncrementValue:output_type -> proto.Message
	1, // 5: proto.ChatService.RingElection:output_type -> proto.Void
	3, // 6: proto.ChatService.SelectNewLeader:output_type -> proto.ElectionPorts
	1, // 7: proto.ChatService.BroadcastNewLeader:output_type -> proto.Void
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chat_proto_init() }
func file_proto_chat_proto_init() {
	if File_proto_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortsAndClocks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionPorts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chat_proto_goTypes,
		DependencyIndexes: file_proto_chat_proto_depIdxs,
		MessageInfos:      file_proto_chat_proto_msgTypes,
	}.Build()
	File_proto_chat_proto = out.File
	file_proto_chat_proto_rawDesc = nil
	file_proto_chat_proto_goTypes = nil
	file_proto_chat_proto_depIdxs = nil
}
