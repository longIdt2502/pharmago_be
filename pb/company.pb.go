// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/company.proto

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

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code          string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Type          string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	TaxCode       *string  `protobuf:"bytes,5,opt,name=tax_code,json=taxCode,proto3,oneof" json:"tax_code,omitempty"`
	Phone         *string  `protobuf:"bytes,6,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Description   *string  `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Address       *Address `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Owner         int32    `protobuf:"varint,9,opt,name=owner,proto3" json:"owner,omitempty"`
	OaId          *string  `protobuf:"bytes,10,opt,name=oa_id,json=oaId,proto3,oneof" json:"oa_id,omitempty"`
	TimeOpen      *string  `protobuf:"bytes,11,opt,name=time_open,json=timeOpen,proto3,oneof" json:"time_open,omitempty"`
	TimeClose     *string  `protobuf:"bytes,12,opt,name=time_close,json=timeClose,proto3,oneof" json:"time_close,omitempty"`
	TotalEmployee int32    `protobuf:"varint,13,opt,name=total_employee,json=totalEmployee,proto3" json:"total_employee,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_entities_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_entities_company_proto_rawDescGZIP(), []int{0}
}

func (x *Company) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Company) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Company) GetTaxCode() string {
	if x != nil && x.TaxCode != nil {
		return *x.TaxCode
	}
	return ""
}

func (x *Company) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *Company) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Company) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Company) GetOwner() int32 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *Company) GetOaId() string {
	if x != nil && x.OaId != nil {
		return *x.OaId
	}
	return ""
}

func (x *Company) GetTimeOpen() string {
	if x != nil && x.TimeOpen != nil {
		return *x.TimeOpen
	}
	return ""
}

func (x *Company) GetTimeClose() string {
	if x != nil && x.TimeClose != nil {
		return *x.TimeClose
	}
	return ""
}

func (x *Company) GetTotalEmployee() int32 {
	if x != nil {
		return x.TotalEmployee
	}
	return 0
}

var File_entities_company_proto protoreflect.FileDescriptor

var file_entities_company_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x08, 0x74, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x74, 0x61, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x05,
	0x6f, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x6f,
	0x61, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x4f, 0x70, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f,
	0x61, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68,
	0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_company_proto_rawDescOnce sync.Once
	file_entities_company_proto_rawDescData = file_entities_company_proto_rawDesc
)

func file_entities_company_proto_rawDescGZIP() []byte {
	file_entities_company_proto_rawDescOnce.Do(func() {
		file_entities_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_company_proto_rawDescData)
	})
	return file_entities_company_proto_rawDescData
}

var file_entities_company_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entities_company_proto_goTypes = []interface{}{
	(*Company)(nil), // 0: pb.Company
	(*Address)(nil), // 1: pb.Address
}
var file_entities_company_proto_depIdxs = []int32{
	1, // 0: pb.Company.address:type_name -> pb.Address
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entities_company_proto_init() }
func file_entities_company_proto_init() {
	if File_entities_company_proto != nil {
		return
	}
	file_entities_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
	file_entities_company_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entities_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_company_proto_goTypes,
		DependencyIndexes: file_entities_company_proto_depIdxs,
		MessageInfos:      file_entities_company_proto_msgTypes,
	}.Build()
	File_entities_company_proto = out.File
	file_entities_company_proto_rawDesc = nil
	file_entities_company_proto_goTypes = nil
	file_entities_company_proto_depIdxs = nil
}
