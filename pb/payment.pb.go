// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/payment.proto

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

type PaymentItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      *SimpleData `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value     float32     `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	IsPaid    bool        `protobuf:"varint,4,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	ExtraNote string      `protobuf:"bytes,5,opt,name=extra_note,json=extraNote,proto3" json:"extra_note,omitempty"`
}

func (x *PaymentItem) Reset() {
	*x = PaymentItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentItem) ProtoMessage() {}

func (x *PaymentItem) ProtoReflect() protoreflect.Message {
	mi := &file_entities_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentItem.ProtoReflect.Descriptor instead.
func (*PaymentItem) Descriptor() ([]byte, []int) {
	return file_entities_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentItem) GetType() *SimpleData {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *PaymentItem) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PaymentItem) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *PaymentItem) GetExtraNote() string {
	if x != nil {
		return x.ExtraNote
	}
	return ""
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code     string         `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	MustPaid float32        `protobuf:"fixed32,3,opt,name=must_paid,json=mustPaid,proto3" json:"must_paid,omitempty"`
	HadPaid  float32        `protobuf:"fixed32,4,opt,name=had_paid,json=hadPaid,proto3" json:"had_paid,omitempty"`
	NeedPay  float32        `protobuf:"fixed32,5,opt,name=need_pay,json=needPay,proto3" json:"need_pay,omitempty"`
	Items    []*PaymentItem `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_entities_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_entities_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Payment) GetMustPaid() float32 {
	if x != nil {
		return x.MustPaid
	}
	return 0
}

func (x *Payment) GetHadPaid() float32 {
	if x != nil {
		return x.HadPaid
	}
	return 0
}

func (x *Payment) GetNeedPay() float32 {
	if x != nil {
		return x.NeedPay
	}
	return 0
}

func (x *Payment) GetItems() []*PaymentItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_entities_payment_proto protoreflect.FileDescriptor

var file_entities_payment_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1a, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x75, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6d, 0x75, 0x73, 0x74, 0x50, 0x61, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x64, 0x5f,
	0x70, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x68, 0x61, 0x64, 0x50,
	0x61, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e, 0x65, 0x65, 0x64, 0x50, 0x61, 0x79, 0x12, 0x25,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30,
	0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_payment_proto_rawDescOnce sync.Once
	file_entities_payment_proto_rawDescData = file_entities_payment_proto_rawDesc
)

func file_entities_payment_proto_rawDescGZIP() []byte {
	file_entities_payment_proto_rawDescOnce.Do(func() {
		file_entities_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_payment_proto_rawDescData)
	})
	return file_entities_payment_proto_rawDescData
}

var file_entities_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entities_payment_proto_goTypes = []interface{}{
	(*PaymentItem)(nil), // 0: pb.PaymentItem
	(*Payment)(nil),     // 1: pb.Payment
	(*SimpleData)(nil),  // 2: pb.SimpleData
}
var file_entities_payment_proto_depIdxs = []int32{
	2, // 0: pb.PaymentItem.type:type_name -> pb.SimpleData
	0, // 1: pb.Payment.items:type_name -> pb.PaymentItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entities_payment_proto_init() }
func file_entities_payment_proto_init() {
	if File_entities_payment_proto != nil {
		return
	}
	file_entities_simple_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentItem); i {
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
		file_entities_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
			RawDescriptor: file_entities_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_payment_proto_goTypes,
		DependencyIndexes: file_entities_payment_proto_depIdxs,
		MessageInfos:      file_entities_payment_proto_msgTypes,
	}.Build()
	File_entities_payment_proto = out.File
	file_entities_payment_proto_rawDesc = nil
	file_entities_payment_proto_goTypes = nil
	file_entities_payment_proto_depIdxs = nil
}
