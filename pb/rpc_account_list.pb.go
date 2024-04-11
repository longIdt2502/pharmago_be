// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/account/rpc_account_list.proto

package pb

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

type AccountListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company int32  `protobuf:"varint,1,opt,name=company,proto3" json:"company,omitempty"`
	Type    *int32 `protobuf:"varint,2,opt,name=type,proto3,oneof" json:"type,omitempty"`
	Search  string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Page    *int32 `protobuf:"varint,4,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit   *int32 `protobuf:"varint,5,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *AccountListRequest) Reset() {
	*x = AccountListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_account_rpc_account_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListRequest) ProtoMessage() {}

func (x *AccountListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_account_rpc_account_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListRequest.ProtoReflect.Descriptor instead.
func (*AccountListRequest) Descriptor() ([]byte, []int) {
	return file_rpc_account_rpc_account_list_proto_rawDescGZIP(), []int{0}
}

func (x *AccountListRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *AccountListRequest) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *AccountListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *AccountListRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *AccountListRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type AccountListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*Account `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *AccountListResponse) Reset() {
	*x = AccountListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_account_rpc_account_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountListResponse) ProtoMessage() {}

func (x *AccountListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_account_rpc_account_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountListResponse.ProtoReflect.Descriptor instead.
func (*AccountListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_account_rpc_account_list_proto_rawDescGZIP(), []int{1}
}

func (x *AccountListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AccountListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AccountListResponse) GetDetails() []*Account {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_rpc_account_rpc_account_list_proto protoreflect.FileDescriptor

var file_rpc_account_rpc_account_list_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaf, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x6a, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61,
	0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d,
	0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_account_rpc_account_list_proto_rawDescOnce sync.Once
	file_rpc_account_rpc_account_list_proto_rawDescData = file_rpc_account_rpc_account_list_proto_rawDesc
)

func file_rpc_account_rpc_account_list_proto_rawDescGZIP() []byte {
	file_rpc_account_rpc_account_list_proto_rawDescOnce.Do(func() {
		file_rpc_account_rpc_account_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_account_rpc_account_list_proto_rawDescData)
	})
	return file_rpc_account_rpc_account_list_proto_rawDescData
}

var file_rpc_account_rpc_account_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_account_rpc_account_list_proto_goTypes = []interface{}{
	(*AccountListRequest)(nil),  // 0: pb.AccountListRequest
	(*AccountListResponse)(nil), // 1: pb.AccountListResponse
	(*Account)(nil),             // 2: pb.Account
}
var file_rpc_account_rpc_account_list_proto_depIdxs = []int32{
	2, // 0: pb.AccountListResponse.details:type_name -> pb.Account
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_account_rpc_account_list_proto_init() }
func file_rpc_account_rpc_account_list_proto_init() {
	if File_rpc_account_rpc_account_list_proto != nil {
		return
	}
	file_entities_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_account_rpc_account_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListRequest); i {
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
		file_rpc_account_rpc_account_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountListResponse); i {
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
	file_rpc_account_rpc_account_list_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_account_rpc_account_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_account_rpc_account_list_proto_goTypes,
		DependencyIndexes: file_rpc_account_rpc_account_list_proto_depIdxs,
		MessageInfos:      file_rpc_account_rpc_account_list_proto_msgTypes,
	}.Build()
	File_rpc_account_rpc_account_list_proto = out.File
	file_rpc_account_rpc_account_list_proto_rawDesc = nil
	file_rpc_account_rpc_account_list_proto_goTypes = nil
	file_rpc_account_rpc_account_list_proto_depIdxs = nil
}
