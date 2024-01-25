// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/product/rpc_preparation_type_create.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PreparationTypeCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company     int32   `protobuf:"varint,1,opt,name=company,proto3" json:"company,omitempty"`
	Code        *string `protobuf:"bytes,2,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PreparationTypeCreateRequest) Reset() {
	*x = PreparationTypeCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_product_rpc_preparation_type_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparationTypeCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparationTypeCreateRequest) ProtoMessage() {}

func (x *PreparationTypeCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_rpc_preparation_type_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparationTypeCreateRequest.ProtoReflect.Descriptor instead.
func (*PreparationTypeCreateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_rpc_preparation_type_create_proto_rawDescGZIP(), []int{0}
}

func (x *PreparationTypeCreateRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *PreparationTypeCreateRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *PreparationTypeCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PreparationTypeCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PreparationTypeCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details int32  `protobuf:"varint,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *PreparationTypeCreateResponse) Reset() {
	*x = PreparationTypeCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_product_rpc_preparation_type_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparationTypeCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparationTypeCreateResponse) ProtoMessage() {}

func (x *PreparationTypeCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_rpc_preparation_type_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparationTypeCreateResponse.ProtoReflect.Descriptor instead.
func (*PreparationTypeCreateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_rpc_preparation_type_create_proto_rawDescGZIP(), []int{1}
}

func (x *PreparationTypeCreateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PreparationTypeCreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PreparationTypeCreateResponse) GetDetails() int32 {
	if x != nil {
		return x.Details
	}
	return 0
}

var File_rpc_product_rpc_preparation_type_create_proto protoreflect.FileDescriptor

var file_rpc_product_rpc_preparation_type_create_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd2, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x20, 0x92, 0x41, 0x1d, 0x32, 0x11, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x20, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0xd2, 0x01, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x17, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0xd2, 0x01, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x67, 0x0a, 0x1d, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61,
	0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d,
	0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_product_rpc_preparation_type_create_proto_rawDescOnce sync.Once
	file_rpc_product_rpc_preparation_type_create_proto_rawDescData = file_rpc_product_rpc_preparation_type_create_proto_rawDesc
)

func file_rpc_product_rpc_preparation_type_create_proto_rawDescGZIP() []byte {
	file_rpc_product_rpc_preparation_type_create_proto_rawDescOnce.Do(func() {
		file_rpc_product_rpc_preparation_type_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_product_rpc_preparation_type_create_proto_rawDescData)
	})
	return file_rpc_product_rpc_preparation_type_create_proto_rawDescData
}

var file_rpc_product_rpc_preparation_type_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_product_rpc_preparation_type_create_proto_goTypes = []interface{}{
	(*PreparationTypeCreateRequest)(nil),  // 0: pb.PreparationTypeCreateRequest
	(*PreparationTypeCreateResponse)(nil), // 1: pb.PreparationTypeCreateResponse
}
var file_rpc_product_rpc_preparation_type_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_product_rpc_preparation_type_create_proto_init() }
func file_rpc_product_rpc_preparation_type_create_proto_init() {
	if File_rpc_product_rpc_preparation_type_create_proto != nil {
		return
	}
	file_entities_simple_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_product_rpc_preparation_type_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparationTypeCreateRequest); i {
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
		file_rpc_product_rpc_preparation_type_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparationTypeCreateResponse); i {
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
	file_rpc_product_rpc_preparation_type_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_product_rpc_preparation_type_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_product_rpc_preparation_type_create_proto_goTypes,
		DependencyIndexes: file_rpc_product_rpc_preparation_type_create_proto_depIdxs,
		MessageInfos:      file_rpc_product_rpc_preparation_type_create_proto_msgTypes,
	}.Build()
	File_rpc_product_rpc_preparation_type_create_proto = out.File
	file_rpc_product_rpc_preparation_type_create_proto_rawDesc = nil
	file_rpc_product_rpc_preparation_type_create_proto_goTypes = nil
	file_rpc_product_rpc_preparation_type_create_proto_depIdxs = nil
}