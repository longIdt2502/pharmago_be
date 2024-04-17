// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/conversation/rpc_list_conversation.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ListConversationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OaId    int64  `protobuf:"varint,1,opt,name=oa_id,json=oaId,proto3" json:"oa_id,omitempty"`
	Serach  string `protobuf:"bytes,2,opt,name=serach,proto3" json:"serach,omitempty"`
	Page    int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32  `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *ListConversationRequest) Reset() {
	*x = ListConversationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_conversation_rpc_list_conversation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConversationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConversationRequest) ProtoMessage() {}

func (x *ListConversationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_conversation_rpc_list_conversation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConversationRequest.ProtoReflect.Descriptor instead.
func (*ListConversationRequest) Descriptor() ([]byte, []int) {
	return file_rpc_conversation_rpc_list_conversation_proto_rawDescGZIP(), []int{0}
}

func (x *ListConversationRequest) GetOaId() int64 {
	if x != nil {
		return x.OaId
	}
	return 0
}

func (x *ListConversationRequest) GetSerach() string {
	if x != nil {
		return x.Serach
	}
	return ""
}

func (x *ListConversationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListConversationRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListConversationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*Conversation `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ListConversationResponse) Reset() {
	*x = ListConversationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_conversation_rpc_list_conversation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConversationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConversationResponse) ProtoMessage() {}

func (x *ListConversationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_conversation_rpc_list_conversation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConversationResponse.ProtoReflect.Descriptor instead.
func (*ListConversationResponse) Descriptor() ([]byte, []int) {
	return file_rpc_conversation_rpc_list_conversation_proto_rawDescGZIP(), []int{1}
}

func (x *ListConversationResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListConversationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListConversationResponse) GetDetails() []*Conversation {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_rpc_conversation_rpc_list_conversation_proto protoreflect.FileDescriptor

var file_rpc_conversation_rpc_list_conversation_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x95, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x6f,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32,
	0x11, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0xd2, 0x01, 0x05, 0x6f, 0x61, 0x5f, 0x69, 0x64, 0x52, 0x04, 0x6f, 0x61, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x61, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x61, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e,
	0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61,
	0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_conversation_rpc_list_conversation_proto_rawDescOnce sync.Once
	file_rpc_conversation_rpc_list_conversation_proto_rawDescData = file_rpc_conversation_rpc_list_conversation_proto_rawDesc
)

func file_rpc_conversation_rpc_list_conversation_proto_rawDescGZIP() []byte {
	file_rpc_conversation_rpc_list_conversation_proto_rawDescOnce.Do(func() {
		file_rpc_conversation_rpc_list_conversation_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_conversation_rpc_list_conversation_proto_rawDescData)
	})
	return file_rpc_conversation_rpc_list_conversation_proto_rawDescData
}

var file_rpc_conversation_rpc_list_conversation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_conversation_rpc_list_conversation_proto_goTypes = []interface{}{
	(*ListConversationRequest)(nil),  // 0: pb.ListConversationRequest
	(*ListConversationResponse)(nil), // 1: pb.ListConversationResponse
	(*Conversation)(nil),             // 2: pb.Conversation
}
var file_rpc_conversation_rpc_list_conversation_proto_depIdxs = []int32{
	2, // 0: pb.ListConversationResponse.details:type_name -> pb.Conversation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_conversation_rpc_list_conversation_proto_init() }
func file_rpc_conversation_rpc_list_conversation_proto_init() {
	if File_rpc_conversation_rpc_list_conversation_proto != nil {
		return
	}
	file_entities_conversation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_conversation_rpc_list_conversation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConversationRequest); i {
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
		file_rpc_conversation_rpc_list_conversation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConversationResponse); i {
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
			RawDescriptor: file_rpc_conversation_rpc_list_conversation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_conversation_rpc_list_conversation_proto_goTypes,
		DependencyIndexes: file_rpc_conversation_rpc_list_conversation_proto_depIdxs,
		MessageInfos:      file_rpc_conversation_rpc_list_conversation_proto_msgTypes,
	}.Build()
	File_rpc_conversation_rpc_list_conversation_proto = out.File
	file_rpc_conversation_rpc_list_conversation_proto_rawDesc = nil
	file_rpc_conversation_rpc_list_conversation_proto_goTypes = nil
	file_rpc_conversation_rpc_list_conversation_proto_depIdxs = nil
}
