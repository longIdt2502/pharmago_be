// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/promotions/rpc_promotion_check.proto

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

type PromotionCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company    int32   `protobuf:"varint,1,opt,name=company,proto3" json:"company,omitempty"`
	TotalPrice float32 `protobuf:"fixed32,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Products   []int32 `protobuf:"varint,3,rep,packed,name=products,proto3" json:"products,omitempty"`
	Services   []int32 `protobuf:"varint,4,rep,packed,name=services,proto3" json:"services,omitempty"`
}

func (x *PromotionCheckRequest) Reset() {
	*x = PromotionCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_promotions_rpc_promotion_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionCheckRequest) ProtoMessage() {}

func (x *PromotionCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_promotions_rpc_promotion_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionCheckRequest.ProtoReflect.Descriptor instead.
func (*PromotionCheckRequest) Descriptor() ([]byte, []int) {
	return file_rpc_promotions_rpc_promotion_check_proto_rawDescGZIP(), []int{0}
}

func (x *PromotionCheckRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *PromotionCheckRequest) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *PromotionCheckRequest) GetProducts() []int32 {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *PromotionCheckRequest) GetServices() []int32 {
	if x != nil {
		return x.Services
	}
	return nil
}

type PromotionCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	RootErr string       `protobuf:"bytes,2,opt,name=root_err,json=rootErr,proto3" json:"root_err,omitempty"`
	Message string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Log     string       `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty"`
	Details []*Promotion `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *PromotionCheckResponse) Reset() {
	*x = PromotionCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_promotions_rpc_promotion_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionCheckResponse) ProtoMessage() {}

func (x *PromotionCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_promotions_rpc_promotion_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionCheckResponse.ProtoReflect.Descriptor instead.
func (*PromotionCheckResponse) Descriptor() ([]byte, []int) {
	return file_rpc_promotions_rpc_promotion_check_proto_rawDescGZIP(), []int{1}
}

func (x *PromotionCheckResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PromotionCheckResponse) GetRootErr() string {
	if x != nil {
		return x.RootErr
	}
	return ""
}

func (x *PromotionCheckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PromotionCheckResponse) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *PromotionCheckResponse) GetDetails() []*Promotion {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_rpc_promotions_rpc_promotion_check_proto protoreflect.FileDescriptor

var file_rpc_promotions_rpc_promotion_check_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x65,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x45, 0x72,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x27, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35,
	0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_promotions_rpc_promotion_check_proto_rawDescOnce sync.Once
	file_rpc_promotions_rpc_promotion_check_proto_rawDescData = file_rpc_promotions_rpc_promotion_check_proto_rawDesc
)

func file_rpc_promotions_rpc_promotion_check_proto_rawDescGZIP() []byte {
	file_rpc_promotions_rpc_promotion_check_proto_rawDescOnce.Do(func() {
		file_rpc_promotions_rpc_promotion_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_promotions_rpc_promotion_check_proto_rawDescData)
	})
	return file_rpc_promotions_rpc_promotion_check_proto_rawDescData
}

var file_rpc_promotions_rpc_promotion_check_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_promotions_rpc_promotion_check_proto_goTypes = []interface{}{
	(*PromotionCheckRequest)(nil),  // 0: pb.PromotionCheckRequest
	(*PromotionCheckResponse)(nil), // 1: pb.PromotionCheckResponse
	(*Promotion)(nil),              // 2: pb.Promotion
}
var file_rpc_promotions_rpc_promotion_check_proto_depIdxs = []int32{
	2, // 0: pb.PromotionCheckResponse.details:type_name -> pb.Promotion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_promotions_rpc_promotion_check_proto_init() }
func file_rpc_promotions_rpc_promotion_check_proto_init() {
	if File_rpc_promotions_rpc_promotion_check_proto != nil {
		return
	}
	file_entities_promotions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_promotions_rpc_promotion_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionCheckRequest); i {
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
		file_rpc_promotions_rpc_promotion_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionCheckResponse); i {
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
			RawDescriptor: file_rpc_promotions_rpc_promotion_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_promotions_rpc_promotion_check_proto_goTypes,
		DependencyIndexes: file_rpc_promotions_rpc_promotion_check_proto_depIdxs,
		MessageInfos:      file_rpc_promotions_rpc_promotion_check_proto_msgTypes,
	}.Build()
	File_rpc_promotions_rpc_promotion_check_proto = out.File
	file_rpc_promotions_rpc_promotion_check_proto_rawDesc = nil
	file_rpc_promotions_rpc_promotion_check_proto_goTypes = nil
	file_rpc_promotions_rpc_promotion_check_proto_depIdxs = nil
}
