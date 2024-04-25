// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/customer/rpc_customer_group_detail.proto

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

type CustomerGroupDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerGroupDetailRequest) Reset() {
	*x = CustomerGroupDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_rpc_customer_group_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerGroupDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerGroupDetailRequest) ProtoMessage() {}

func (x *CustomerGroupDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_rpc_customer_group_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerGroupDetailRequest.ProtoReflect.Descriptor instead.
func (*CustomerGroupDetailRequest) Descriptor() ([]byte, []int) {
	return file_rpc_customer_rpc_customer_group_detail_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerGroupDetailRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CustomerGroupDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details *SimpleData `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *CustomerGroupDetailResponse) Reset() {
	*x = CustomerGroupDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_rpc_customer_group_detail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerGroupDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerGroupDetailResponse) ProtoMessage() {}

func (x *CustomerGroupDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_rpc_customer_group_detail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerGroupDetailResponse.ProtoReflect.Descriptor instead.
func (*CustomerGroupDetailResponse) Descriptor() ([]byte, []int) {
	return file_rpc_customer_rpc_customer_group_detail_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerGroupDetailResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CustomerGroupDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CustomerGroupDetailResponse) GetDetails() *SimpleData {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_rpc_customer_rpc_customer_group_detail_proto protoreflect.FileDescriptor

var file_rpc_customer_rpc_customer_group_detail_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x32, 0x11, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0xd2, 0x01, 0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x1b, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48,
	0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61,
	0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_customer_rpc_customer_group_detail_proto_rawDescOnce sync.Once
	file_rpc_customer_rpc_customer_group_detail_proto_rawDescData = file_rpc_customer_rpc_customer_group_detail_proto_rawDesc
)

func file_rpc_customer_rpc_customer_group_detail_proto_rawDescGZIP() []byte {
	file_rpc_customer_rpc_customer_group_detail_proto_rawDescOnce.Do(func() {
		file_rpc_customer_rpc_customer_group_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_customer_rpc_customer_group_detail_proto_rawDescData)
	})
	return file_rpc_customer_rpc_customer_group_detail_proto_rawDescData
}

var file_rpc_customer_rpc_customer_group_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_customer_rpc_customer_group_detail_proto_goTypes = []interface{}{
	(*CustomerGroupDetailRequest)(nil),  // 0: pb.CustomerGroupDetailRequest
	(*CustomerGroupDetailResponse)(nil), // 1: pb.CustomerGroupDetailResponse
	(*SimpleData)(nil),                  // 2: pb.SimpleData
}
var file_rpc_customer_rpc_customer_group_detail_proto_depIdxs = []int32{
	2, // 0: pb.CustomerGroupDetailResponse.details:type_name -> pb.SimpleData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_customer_rpc_customer_group_detail_proto_init() }
func file_rpc_customer_rpc_customer_group_detail_proto_init() {
	if File_rpc_customer_rpc_customer_group_detail_proto != nil {
		return
	}
	file_entities_simple_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_customer_rpc_customer_group_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerGroupDetailRequest); i {
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
		file_rpc_customer_rpc_customer_group_detail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerGroupDetailResponse); i {
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
			RawDescriptor: file_rpc_customer_rpc_customer_group_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_customer_rpc_customer_group_detail_proto_goTypes,
		DependencyIndexes: file_rpc_customer_rpc_customer_group_detail_proto_depIdxs,
		MessageInfos:      file_rpc_customer_rpc_customer_group_detail_proto_msgTypes,
	}.Build()
	File_rpc_customer_rpc_customer_group_detail_proto = out.File
	file_rpc_customer_rpc_customer_group_detail_proto_rawDesc = nil
	file_rpc_customer_rpc_customer_group_detail_proto_goTypes = nil
	file_rpc_customer_rpc_customer_group_detail_proto_depIdxs = nil
}
