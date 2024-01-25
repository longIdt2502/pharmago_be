// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/product/rpc_company_pharma_list.proto

package pb

import (
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

type CompanyPharmaListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *int32  `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit  *int32  `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Search *string `protobuf:"bytes,3,opt,name=search,proto3,oneof" json:"search,omitempty"`
	Type   *string `protobuf:"bytes,4,opt,name=type,proto3,oneof" json:"type,omitempty"`
}

func (x *CompanyPharmaListRequest) Reset() {
	*x = CompanyPharmaListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_product_rpc_company_pharma_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyPharmaListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyPharmaListRequest) ProtoMessage() {}

func (x *CompanyPharmaListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_rpc_company_pharma_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyPharmaListRequest.ProtoReflect.Descriptor instead.
func (*CompanyPharmaListRequest) Descriptor() ([]byte, []int) {
	return file_rpc_product_rpc_company_pharma_list_proto_rawDescGZIP(), []int{0}
}

func (x *CompanyPharmaListRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *CompanyPharmaListRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *CompanyPharmaListRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

func (x *CompanyPharmaListRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type CompanyPharmaListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*CompanyPharma `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *CompanyPharmaListResponse) Reset() {
	*x = CompanyPharmaListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_product_rpc_company_pharma_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyPharmaListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyPharmaListResponse) ProtoMessage() {}

func (x *CompanyPharmaListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_product_rpc_company_pharma_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyPharmaListResponse.ProtoReflect.Descriptor instead.
func (*CompanyPharmaListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_product_rpc_company_pharma_list_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyPharmaListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CompanyPharmaListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CompanyPharmaListResponse) GetDetails() []*CompanyPharma {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_rpc_product_rpc_company_pharma_list_proto protoreflect.FileDescriptor

var file_rpc_product_rpc_company_pharma_list_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xab, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x68, 0x61, 0x72, 0x6d,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x76, 0x0a,
	0x19, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30,
	0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_product_rpc_company_pharma_list_proto_rawDescOnce sync.Once
	file_rpc_product_rpc_company_pharma_list_proto_rawDescData = file_rpc_product_rpc_company_pharma_list_proto_rawDesc
)

func file_rpc_product_rpc_company_pharma_list_proto_rawDescGZIP() []byte {
	file_rpc_product_rpc_company_pharma_list_proto_rawDescOnce.Do(func() {
		file_rpc_product_rpc_company_pharma_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_product_rpc_company_pharma_list_proto_rawDescData)
	})
	return file_rpc_product_rpc_company_pharma_list_proto_rawDescData
}

var file_rpc_product_rpc_company_pharma_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_product_rpc_company_pharma_list_proto_goTypes = []interface{}{
	(*CompanyPharmaListRequest)(nil),  // 0: pb.CompanyPharmaListRequest
	(*CompanyPharmaListResponse)(nil), // 1: pb.CompanyPharmaListResponse
	(*CompanyPharma)(nil),             // 2: pb.CompanyPharma
}
var file_rpc_product_rpc_company_pharma_list_proto_depIdxs = []int32{
	2, // 0: pb.CompanyPharmaListResponse.details:type_name -> pb.CompanyPharma
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_product_rpc_company_pharma_list_proto_init() }
func file_rpc_product_rpc_company_pharma_list_proto_init() {
	if File_rpc_product_rpc_company_pharma_list_proto != nil {
		return
	}
	file_entities_company_pharma_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_product_rpc_company_pharma_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyPharmaListRequest); i {
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
		file_rpc_product_rpc_company_pharma_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyPharmaListResponse); i {
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
	file_rpc_product_rpc_company_pharma_list_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_product_rpc_company_pharma_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_product_rpc_company_pharma_list_proto_goTypes,
		DependencyIndexes: file_rpc_product_rpc_company_pharma_list_proto_depIdxs,
		MessageInfos:      file_rpc_product_rpc_company_pharma_list_proto_msgTypes,
	}.Build()
	File_rpc_product_rpc_company_pharma_list_proto = out.File
	file_rpc_product_rpc_company_pharma_list_proto_rawDesc = nil
	file_rpc_product_rpc_company_pharma_list_proto_goTypes = nil
	file_rpc_product_rpc_company_pharma_list_proto_depIdxs = nil
}
