// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: grpc/proto.proto

package proto

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

type ClientReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientAddress string `protobuf:"bytes,1,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
	ClientPort    int32  `protobuf:"varint,2,opt,name=client_port,json=clientPort,proto3" json:"client_port,omitempty"`
	ClientName    string `protobuf:"bytes,3,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
}

func (x *ClientReference) Reset() {
	*x = ClientReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReference) ProtoMessage() {}

func (x *ClientReference) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReference.ProtoReflect.Descriptor instead.
func (*ClientReference) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

func (x *ClientReference) GetClientAddress() string {
	if x != nil {
		return x.ClientAddress
	}
	return ""
}

func (x *ClientReference) GetClientPort() int32 {
	if x != nil {
		return x.ClientPort
	}
	return 0
}

func (x *ClientReference) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time will be represented by Lamport clocks incremented when a message is received or sended
	ClientReference *ClientReference `protobuf:"bytes,1,opt,name=client_reference,json=clientReference,proto3" json:"client_reference,omitempty"`
	Time            int32            `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetClientReference() *ClientReference {
	if x != nil {
		return x.ClientReference
	}
	return nil
}

func (x *Question) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply bool  `protobuf:"varint,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Time  int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

func (x *Answer) GetReply() bool {
	if x != nil {
		return x.Reply
	}
	return false
}

func (x *Answer) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x41, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x48, 0x0a, 0x15,
	0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x78, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0d, 0x41, 0x73, 0x6b, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_proto_rawDescOnce sync.Once
	file_grpc_proto_proto_rawDescData = file_grpc_proto_proto_rawDesc
)

func file_grpc_proto_proto_rawDescGZIP() []byte {
	file_grpc_proto_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proto_rawDescData)
	})
	return file_grpc_proto_proto_rawDescData
}

var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_proto_proto_goTypes = []interface{}{
	(*ClientReference)(nil), // 0: proto.ClientReference
	(*Question)(nil),        // 1: proto.Question
	(*Answer)(nil),          // 2: proto.Answer
}
var file_grpc_proto_proto_depIdxs = []int32{
	0, // 0: proto.Question.client_reference:type_name -> proto.ClientReference
	1, // 1: proto.MutualExlusionService.AskPermission:input_type -> proto.Question
	2, // 2: proto.MutualExlusionService.AskPermission:output_type -> proto.Answer
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_proto_init() }
func file_grpc_proto_proto_init() {
	if File_grpc_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReference); i {
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
		file_grpc_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_grpc_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
			RawDescriptor: file_grpc_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proto_depIdxs,
		MessageInfos:      file_grpc_proto_proto_msgTypes,
	}.Build()
	File_grpc_proto_proto = out.File
	file_grpc_proto_proto_rawDesc = nil
	file_grpc_proto_proto_goTypes = nil
	file_grpc_proto_proto_depIdxs = nil
}
