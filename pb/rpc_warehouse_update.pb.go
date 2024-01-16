// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: rpc/warehouse/rpc_warehouse_update.proto

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

type WarehouseUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    *string         `protobuf:"bytes,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Name    *string         `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Address *AddressPayload `protobuf:"bytes,3,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Id      int32           `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WarehouseUpdateRequest) Reset() {
	*x = WarehouseUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseUpdateRequest) ProtoMessage() {}

func (x *WarehouseUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseUpdateRequest.ProtoReflect.Descriptor instead.
func (*WarehouseUpdateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_warehouse_rpc_warehouse_update_proto_rawDescGZIP(), []int{0}
}

func (x *WarehouseUpdateRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *WarehouseUpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *WarehouseUpdateRequest) GetAddress() *AddressPayload {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *WarehouseUpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type WarehouseUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WarehouseUpdateResponse) Reset() {
	*x = WarehouseUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseUpdateResponse) ProtoMessage() {}

func (x *WarehouseUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseUpdateResponse.ProtoReflect.Descriptor instead.
func (*WarehouseUpdateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_warehouse_rpc_warehouse_update_proto_rawDescGZIP(), []int{1}
}

func (x *WarehouseUpdateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *WarehouseUpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_warehouse_rpc_warehouse_update_proto protoreflect.FileDescriptor

var file_rpc_warehouse_rpc_warehouse_update_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x16,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x02, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x47, 0x0a, 0x17, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e,
	0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61,
	0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_warehouse_rpc_warehouse_update_proto_rawDescOnce sync.Once
	file_rpc_warehouse_rpc_warehouse_update_proto_rawDescData = file_rpc_warehouse_rpc_warehouse_update_proto_rawDesc
)

func file_rpc_warehouse_rpc_warehouse_update_proto_rawDescGZIP() []byte {
	file_rpc_warehouse_rpc_warehouse_update_proto_rawDescOnce.Do(func() {
		file_rpc_warehouse_rpc_warehouse_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_warehouse_rpc_warehouse_update_proto_rawDescData)
	})
	return file_rpc_warehouse_rpc_warehouse_update_proto_rawDescData
}

var file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_warehouse_rpc_warehouse_update_proto_goTypes = []interface{}{
	(*WarehouseUpdateRequest)(nil),  // 0: pb.WarehouseUpdateRequest
	(*WarehouseUpdateResponse)(nil), // 1: pb.WarehouseUpdateResponse
	(*AddressPayload)(nil),          // 2: pb.AddressPayload
}
var file_rpc_warehouse_rpc_warehouse_update_proto_depIdxs = []int32{
	2, // 0: pb.WarehouseUpdateRequest.address:type_name -> pb.AddressPayload
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_warehouse_rpc_warehouse_update_proto_init() }
func file_rpc_warehouse_rpc_warehouse_update_proto_init() {
	if File_rpc_warehouse_rpc_warehouse_update_proto != nil {
		return
	}
	file_entities_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseUpdateRequest); i {
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
		file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseUpdateResponse); i {
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
	file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_warehouse_rpc_warehouse_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_warehouse_rpc_warehouse_update_proto_goTypes,
		DependencyIndexes: file_rpc_warehouse_rpc_warehouse_update_proto_depIdxs,
		MessageInfos:      file_rpc_warehouse_rpc_warehouse_update_proto_msgTypes,
	}.Build()
	File_rpc_warehouse_rpc_warehouse_update_proto = out.File
	file_rpc_warehouse_rpc_warehouse_update_proto_rawDesc = nil
	file_rpc_warehouse_rpc_warehouse_update_proto_goTypes = nil
	file_rpc_warehouse_rpc_warehouse_update_proto_depIdxs = nil
}