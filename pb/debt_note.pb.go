// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/debt_note.proto

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

type DebtNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code        string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Entity      string                 `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Money       float32                `protobuf:"fixed32,5,opt,name=money,proto3" json:"money,omitempty"`
	Paymented   float32                `protobuf:"fixed32,6,opt,name=paymented,proto3" json:"paymented,omitempty"`
	Note        string                 `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	Type        string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Status      string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Company     int32                  `protobuf:"varint,10,opt,name=company,proto3" json:"company,omitempty"`
	UserCreated int32                  `protobuf:"varint,11,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	Exprise     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=exprise,proto3" json:"exprise,omitempty"`
	DabtNoteAt  *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=dabt_note_at,json=dabtNoteAt,proto3" json:"dabt_note_at,omitempty"`
	Repayments  []*DebtRepayment       `protobuf:"bytes,14,rep,name=repayments,proto3" json:"repayments,omitempty"`
}

func (x *DebtNote) Reset() {
	*x = DebtNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_debt_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebtNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebtNote) ProtoMessage() {}

func (x *DebtNote) ProtoReflect() protoreflect.Message {
	mi := &file_entities_debt_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebtNote.ProtoReflect.Descriptor instead.
func (*DebtNote) Descriptor() ([]byte, []int) {
	return file_entities_debt_note_proto_rawDescGZIP(), []int{0}
}

func (x *DebtNote) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DebtNote) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DebtNote) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DebtNote) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *DebtNote) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *DebtNote) GetPaymented() float32 {
	if x != nil {
		return x.Paymented
	}
	return 0
}

func (x *DebtNote) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *DebtNote) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DebtNote) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DebtNote) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *DebtNote) GetUserCreated() int32 {
	if x != nil {
		return x.UserCreated
	}
	return 0
}

func (x *DebtNote) GetExprise() *timestamppb.Timestamp {
	if x != nil {
		return x.Exprise
	}
	return nil
}

func (x *DebtNote) GetDabtNoteAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DabtNoteAt
	}
	return nil
}

func (x *DebtNote) GetRepayments() []*DebtRepayment {
	if x != nil {
		return x.Repayments
	}
	return nil
}

type DebtRepayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code        string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Money       float32                `protobuf:"fixed32,3,opt,name=money,proto3" json:"money,omitempty"`
	Debt        int32                  `protobuf:"varint,4,opt,name=debt,proto3" json:"debt,omitempty"`
	UserCreated int32                  `protobuf:"varint,5,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *DebtRepayment) Reset() {
	*x = DebtRepayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_debt_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebtRepayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebtRepayment) ProtoMessage() {}

func (x *DebtRepayment) ProtoReflect() protoreflect.Message {
	mi := &file_entities_debt_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebtRepayment.ProtoReflect.Descriptor instead.
func (*DebtRepayment) Descriptor() ([]byte, []int) {
	return file_entities_debt_note_proto_rawDescGZIP(), []int{1}
}

func (x *DebtRepayment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DebtRepayment) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DebtRepayment) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *DebtRepayment) GetDebt() int32 {
	if x != nil {
		return x.Debt
	}
	return 0
}

func (x *DebtRepayment) GetUserCreated() int32 {
	if x != nil {
		return x.UserCreated
	}
	return 0
}

func (x *DebtRepayment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type DebtReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chart   []*DebtReportChart   `protobuf:"bytes,1,rep,name=chart,proto3" json:"chart,omitempty"`
	Revenue []*DebtReportRevenue `protobuf:"bytes,2,rep,name=revenue,proto3" json:"revenue,omitempty"`
}

func (x *DebtReport) Reset() {
	*x = DebtReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_debt_note_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebtReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebtReport) ProtoMessage() {}

func (x *DebtReport) ProtoReflect() protoreflect.Message {
	mi := &file_entities_debt_note_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebtReport.ProtoReflect.Descriptor instead.
func (*DebtReport) Descriptor() ([]byte, []int) {
	return file_entities_debt_note_proto_rawDescGZIP(), []int{2}
}

func (x *DebtReport) GetChart() []*DebtReportChart {
	if x != nil {
		return x.Chart
	}
	return nil
}

func (x *DebtReport) GetRevenue() []*DebtReportRevenue {
	if x != nil {
		return x.Revenue
	}
	return nil
}

type DebtReportChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Ticket int32                  `protobuf:"varint,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Money  float32                `protobuf:"fixed32,3,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *DebtReportChart) Reset() {
	*x = DebtReportChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_debt_note_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebtReportChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebtReportChart) ProtoMessage() {}

func (x *DebtReportChart) ProtoReflect() protoreflect.Message {
	mi := &file_entities_debt_note_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebtReportChart.ProtoReflect.Descriptor instead.
func (*DebtReportChart) Descriptor() ([]byte, []int) {
	return file_entities_debt_note_proto_rawDescGZIP(), []int{3}
}

func (x *DebtReportChart) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *DebtReportChart) GetTicket() int32 {
	if x != nil {
		return x.Ticket
	}
	return 0
}

func (x *DebtReportChart) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

type DebtReportRevenue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Quantity int32   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Money    float32 `protobuf:"fixed32,3,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *DebtReportRevenue) Reset() {
	*x = DebtReportRevenue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_debt_note_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebtReportRevenue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebtReportRevenue) ProtoMessage() {}

func (x *DebtReportRevenue) ProtoReflect() protoreflect.Message {
	mi := &file_entities_debt_note_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebtReportRevenue.ProtoReflect.Descriptor instead.
func (*DebtReportRevenue) Descriptor() ([]byte, []int) {
	return file_entities_debt_note_proto_rawDescGZIP(), []int{4}
}

func (x *DebtReportRevenue) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DebtReportRevenue) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *DebtReportRevenue) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

var File_entities_debt_note_proto protoreflect.FileDescriptor

var file_entities_debt_note_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x62, 0x74, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb4, 0x03, 0x0a, 0x08, 0x44, 0x65, 0x62, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x72, 0x69, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78,
	0x70, 0x72, 0x69, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x61, 0x62, 0x74, 0x5f, 0x6e, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x61, 0x62, 0x74, 0x4e, 0x6f, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x62,
	0x74, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x62, 0x74, 0x52,
	0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x62, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x64, 0x65, 0x62, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x68, 0x0a, 0x0a, 0x44, 0x65, 0x62, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x62, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x12, 0x2f, 0x0a,
	0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x62, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x22, 0x6f,
	0x0a, 0x0f, 0x44, 0x65, 0x62, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22,
	0x59, 0x0a, 0x11, 0x44, 0x65, 0x62, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x76,
	0x65, 0x6e, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f,
	0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f,
	0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_debt_note_proto_rawDescOnce sync.Once
	file_entities_debt_note_proto_rawDescData = file_entities_debt_note_proto_rawDesc
)

func file_entities_debt_note_proto_rawDescGZIP() []byte {
	file_entities_debt_note_proto_rawDescOnce.Do(func() {
		file_entities_debt_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_debt_note_proto_rawDescData)
	})
	return file_entities_debt_note_proto_rawDescData
}

var file_entities_debt_note_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_entities_debt_note_proto_goTypes = []interface{}{
	(*DebtNote)(nil),              // 0: pb.DebtNote
	(*DebtRepayment)(nil),         // 1: pb.DebtRepayment
	(*DebtReport)(nil),            // 2: pb.DebtReport
	(*DebtReportChart)(nil),       // 3: pb.DebtReportChart
	(*DebtReportRevenue)(nil),     // 4: pb.DebtReportRevenue
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_entities_debt_note_proto_depIdxs = []int32{
	5, // 0: pb.DebtNote.exprise:type_name -> google.protobuf.Timestamp
	5, // 1: pb.DebtNote.dabt_note_at:type_name -> google.protobuf.Timestamp
	1, // 2: pb.DebtNote.repayments:type_name -> pb.DebtRepayment
	5, // 3: pb.DebtRepayment.created_at:type_name -> google.protobuf.Timestamp
	3, // 4: pb.DebtReport.chart:type_name -> pb.DebtReportChart
	4, // 5: pb.DebtReport.revenue:type_name -> pb.DebtReportRevenue
	5, // 6: pb.DebtReportChart.date:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_entities_debt_note_proto_init() }
func file_entities_debt_note_proto_init() {
	if File_entities_debt_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entities_debt_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebtNote); i {
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
		file_entities_debt_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebtRepayment); i {
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
		file_entities_debt_note_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebtReport); i {
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
		file_entities_debt_note_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebtReportChart); i {
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
		file_entities_debt_note_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebtReportRevenue); i {
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
			RawDescriptor: file_entities_debt_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_debt_note_proto_goTypes,
		DependencyIndexes: file_entities_debt_note_proto_depIdxs,
		MessageInfos:      file_entities_debt_note_proto_msgTypes,
	}.Build()
	File_entities_debt_note_proto = out.File
	file_entities_debt_note_proto_rawDesc = nil
	file_entities_debt_note_proto_goTypes = nil
	file_entities_debt_note_proto_depIdxs = nil
}
