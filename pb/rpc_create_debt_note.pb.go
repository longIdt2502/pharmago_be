// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/debt_note/rpc_create_debt_note.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDebtNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company   int32                  `protobuf:"varint,1,opt,name=company,proto3" json:"company,omitempty"`
	Code      *string                `protobuf:"bytes,2,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Title     *string                `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Entity    string                 `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Money     float32                `protobuf:"fixed32,5,opt,name=money,proto3" json:"money,omitempty"`
	Paymented float32                `protobuf:"fixed32,6,opt,name=paymented,proto3" json:"paymented,omitempty"`
	Note      *string                `protobuf:"bytes,7,opt,name=note,proto3,oneof" json:"note,omitempty"`
	Type      string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Status    *string                `protobuf:"bytes,9,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Exprise   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=exprise,proto3" json:"exprise,omitempty"`
}

func (x *CreateDebtNoteRequest) Reset() {
	*x = CreateDebtNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDebtNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDebtNoteRequest) ProtoMessage() {}

func (x *CreateDebtNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDebtNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateDebtNoteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_debt_note_rpc_create_debt_note_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDebtNoteRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *CreateDebtNoteRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *CreateDebtNoteRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *CreateDebtNoteRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *CreateDebtNoteRequest) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *CreateDebtNoteRequest) GetPaymented() float32 {
	if x != nil {
		return x.Paymented
	}
	return 0
}

func (x *CreateDebtNoteRequest) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

func (x *CreateDebtNoteRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateDebtNoteRequest) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *CreateDebtNoteRequest) GetExprise() *timestamppb.Timestamp {
	if x != nil {
		return x.Exprise
	}
	return nil
}

type CreateDebtNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details int32  `protobuf:"varint,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *CreateDebtNoteResponse) Reset() {
	*x = CreateDebtNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDebtNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDebtNoteResponse) ProtoMessage() {}

func (x *CreateDebtNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDebtNoteResponse.ProtoReflect.Descriptor instead.
func (*CreateDebtNoteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_debt_note_rpc_create_debt_note_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDebtNoteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateDebtNoteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateDebtNoteResponse) GetDetails() int32 {
	if x != nil {
		return x.Details
	}
	return 0
}

var File_rpc_debt_note_rpc_create_debt_note_proto protoreflect.FileDescriptor

var file_rpc_debt_note_rpc_create_debt_note_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x62, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x62, 0x74, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x18,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x62, 0x74, 0x5f, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x62, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x17, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x72, 0x69, 0x73, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x60, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x62, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35,
	0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_debt_note_rpc_create_debt_note_proto_rawDescOnce sync.Once
	file_rpc_debt_note_rpc_create_debt_note_proto_rawDescData = file_rpc_debt_note_rpc_create_debt_note_proto_rawDesc
)

func file_rpc_debt_note_rpc_create_debt_note_proto_rawDescGZIP() []byte {
	file_rpc_debt_note_rpc_create_debt_note_proto_rawDescOnce.Do(func() {
		file_rpc_debt_note_rpc_create_debt_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_debt_note_rpc_create_debt_note_proto_rawDescData)
	})
	return file_rpc_debt_note_rpc_create_debt_note_proto_rawDescData
}

var file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_debt_note_rpc_create_debt_note_proto_goTypes = []interface{}{
	(*CreateDebtNoteRequest)(nil),  // 0: pb.CreateDebtNoteRequest
	(*CreateDebtNoteResponse)(nil), // 1: pb.CreateDebtNoteResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_rpc_debt_note_rpc_create_debt_note_proto_depIdxs = []int32{
	2, // 0: pb.CreateDebtNoteRequest.exprise:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_debt_note_rpc_create_debt_note_proto_init() }
func file_rpc_debt_note_rpc_create_debt_note_proto_init() {
	if File_rpc_debt_note_rpc_create_debt_note_proto != nil {
		return
	}
	file_entities_debt_note_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDebtNoteRequest); i {
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
		file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDebtNoteResponse); i {
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
	file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_debt_note_rpc_create_debt_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_debt_note_rpc_create_debt_note_proto_goTypes,
		DependencyIndexes: file_rpc_debt_note_rpc_create_debt_note_proto_depIdxs,
		MessageInfos:      file_rpc_debt_note_rpc_create_debt_note_proto_msgTypes,
	}.Build()
	File_rpc_debt_note_rpc_create_debt_note_proto = out.File
	file_rpc_debt_note_rpc_create_debt_note_proto_rawDesc = nil
	file_rpc_debt_note_rpc_create_debt_note_proto_goTypes = nil
	file_rpc_debt_note_rpc_create_debt_note_proto_depIdxs = nil
}
