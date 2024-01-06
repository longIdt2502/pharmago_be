// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: rpc/order/rpc_order_create.proto

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

type PaymentItemCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value     float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	IsPaid    bool    `protobuf:"varint,3,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	ExtraNote *string `protobuf:"bytes,4,opt,name=extra_note,json=extraNote,proto3,oneof" json:"extra_note,omitempty"`
}

func (x *PaymentItemCreate) Reset() {
	*x = PaymentItemCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentItemCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentItemCreate) ProtoMessage() {}

func (x *PaymentItemCreate) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentItemCreate.ProtoReflect.Descriptor instead.
func (*PaymentItemCreate) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_create_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentItemCreate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PaymentItemCreate) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PaymentItemCreate) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *PaymentItemCreate) GetExtraNote() string {
	if x != nil && x.ExtraNote != nil {
		return *x.ExtraNote
	}
	return ""
}

type PaymentCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MustPaid float32 `protobuf:"fixed32,1,opt,name=must_paid,json=mustPaid,proto3" json:"must_paid,omitempty"`
	HadPaid  float32 `protobuf:"fixed32,2,opt,name=had_paid,json=hadPaid,proto3" json:"had_paid,omitempty"`
	NeedPay  float32 `protobuf:"fixed32,3,opt,name=need_pay,json=needPay,proto3" json:"need_pay,omitempty"`
}

func (x *PaymentCreate) Reset() {
	*x = PaymentCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreate) ProtoMessage() {}

func (x *PaymentCreate) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreate.ProtoReflect.Descriptor instead.
func (*PaymentCreate) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_create_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentCreate) GetMustPaid() float32 {
	if x != nil {
		return x.MustPaid
	}
	return 0
}

func (x *PaymentCreate) GetHadPaid() float32 {
	if x != nil {
		return x.HadPaid
	}
	return 0
}

func (x *PaymentCreate) GetNeedPay() float32 {
	if x != nil {
		return x.NeedPay
	}
	return 0
}

type OrderItemCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variant     int32   `protobuf:"varint,1,opt,name=variant,proto3" json:"variant,omitempty"`
	Value       int32   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Consignment *int32  `protobuf:"varint,3,opt,name=consignment,proto3,oneof" json:"consignment,omitempty"`
	TotalPrice  float32 `protobuf:"fixed32,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *OrderItemCreate) Reset() {
	*x = OrderItemCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemCreate) ProtoMessage() {}

func (x *OrderItemCreate) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemCreate.ProtoReflect.Descriptor instead.
func (*OrderItemCreate) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_create_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItemCreate) GetVariant() int32 {
	if x != nil {
		return x.Variant
	}
	return 0
}

func (x *OrderItemCreate) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *OrderItemCreate) GetConsignment() int32 {
	if x != nil && x.Consignment != nil {
		return *x.Consignment
	}
	return 0
}

func (x *OrderItemCreate) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type OrderCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          *string `protobuf:"bytes,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	TotalPrice    float32 `protobuf:"fixed32,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Description   *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Vat           float32 `protobuf:"fixed32,4,opt,name=vat,proto3" json:"vat,omitempty"`
	Discount      string  `protobuf:"bytes,5,opt,name=discount,proto3" json:"discount,omitempty"`
	ServicePrice  float32 `protobuf:"fixed32,6,opt,name=service_price,json=servicePrice,proto3" json:"service_price,omitempty"`
	MustPaid      float32 `protobuf:"fixed32,7,opt,name=must_paid,json=mustPaid,proto3" json:"must_paid,omitempty"`
	Customer      *int32  `protobuf:"varint,8,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
	Status        string  `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Type          string  `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Company       int32   `protobuf:"varint,11,opt,name=company,proto3" json:"company,omitempty"`
	CustomerPhone *string `protobuf:"bytes,12,opt,name=customer_phone,json=customerPhone,proto3,oneof" json:"customer_phone,omitempty"`
}

func (x *OrderCreate) Reset() {
	*x = OrderCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreate) ProtoMessage() {}

func (x *OrderCreate) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreate.ProtoReflect.Descriptor instead.
func (*OrderCreate) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_create_proto_rawDescGZIP(), []int{3}
}

func (x *OrderCreate) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *OrderCreate) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderCreate) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *OrderCreate) GetVat() float32 {
	if x != nil {
		return x.Vat
	}
	return 0
}

func (x *OrderCreate) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *OrderCreate) GetServicePrice() float32 {
	if x != nil {
		return x.ServicePrice
	}
	return 0
}

func (x *OrderCreate) GetMustPaid() float32 {
	if x != nil {
		return x.MustPaid
	}
	return 0
}

func (x *OrderCreate) GetCustomer() int32 {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return 0
}

func (x *OrderCreate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderCreate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderCreate) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *OrderCreate) GetCustomerPhone() string {
	if x != nil && x.CustomerPhone != nil {
		return *x.CustomerPhone
	}
	return ""
}

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order        *OrderCreate         `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	OrderItems   []*OrderItemCreate   `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	Payment      *PaymentCreate       `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
	PaymentItems []*PaymentItemCreate `protobuf:"bytes,4,rep,name=payment_items,json=paymentItems,proto3" json:"payment_items,omitempty"`
	Warehouse    int32                `protobuf:"varint,5,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_create_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_create_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_create_proto_rawDescGZIP(), []int{4}
}

func (x *OrderCreateRequest) GetOrder() *OrderCreate {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *OrderCreateRequest) GetOrderItems() []*OrderItemCreate {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

func (x *OrderCreateRequest) GetPayment() *PaymentCreate {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *OrderCreateRequest) GetPaymentItems() []*PaymentItemCreate {
	if x != nil {
		return x.PaymentItems
	}
	return nil
}

func (x *OrderCreateRequest) GetWarehouse() int32 {
	if x != nil {
		return x.Warehouse
	}
	return 0
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details int32  `protobuf:"varint,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_create_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_create_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_create_proto_rawDescGZIP(), []int{5}
}

func (x *OrderCreateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OrderCreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderCreateResponse) GetDetails() int32 {
	if x != nil {
		return x.Details
	}
	return 0
}

var File_rpc_order_rpc_order_create_proto protoreflect.FileDescriptor

var file_rpc_order_rpc_order_create_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x89, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4e, 0x6f, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6e, 0x6f,
	0x74, 0x65, 0x22, 0x62, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x75, 0x73, 0x74, 0x50, 0x61, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x68, 0x61, 0x64, 0x50, 0x61, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x65, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e,
	0x65, 0x65, 0x64, 0x50, 0x61, 0x79, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0xaa, 0x03, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x76, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x70,
	0x61, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x75, 0x73, 0x74, 0x50,
	0x61, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0xf8, 0x01, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x0c,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e,
	0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_order_rpc_order_create_proto_rawDescOnce sync.Once
	file_rpc_order_rpc_order_create_proto_rawDescData = file_rpc_order_rpc_order_create_proto_rawDesc
)

func file_rpc_order_rpc_order_create_proto_rawDescGZIP() []byte {
	file_rpc_order_rpc_order_create_proto_rawDescOnce.Do(func() {
		file_rpc_order_rpc_order_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_order_rpc_order_create_proto_rawDescData)
	})
	return file_rpc_order_rpc_order_create_proto_rawDescData
}

var file_rpc_order_rpc_order_create_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_order_rpc_order_create_proto_goTypes = []interface{}{
	(*PaymentItemCreate)(nil),   // 0: pb.PaymentItemCreate
	(*PaymentCreate)(nil),       // 1: pb.PaymentCreate
	(*OrderItemCreate)(nil),     // 2: pb.OrderItemCreate
	(*OrderCreate)(nil),         // 3: pb.OrderCreate
	(*OrderCreateRequest)(nil),  // 4: pb.OrderCreateRequest
	(*OrderCreateResponse)(nil), // 5: pb.OrderCreateResponse
}
var file_rpc_order_rpc_order_create_proto_depIdxs = []int32{
	3, // 0: pb.OrderCreateRequest.order:type_name -> pb.OrderCreate
	2, // 1: pb.OrderCreateRequest.order_items:type_name -> pb.OrderItemCreate
	1, // 2: pb.OrderCreateRequest.payment:type_name -> pb.PaymentCreate
	0, // 3: pb.OrderCreateRequest.payment_items:type_name -> pb.PaymentItemCreate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_order_rpc_order_create_proto_init() }
func file_rpc_order_rpc_order_create_proto_init() {
	if File_rpc_order_rpc_order_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_order_rpc_order_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentItemCreate); i {
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
		file_rpc_order_rpc_order_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreate); i {
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
		file_rpc_order_rpc_order_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemCreate); i {
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
		file_rpc_order_rpc_order_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreate); i {
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
		file_rpc_order_rpc_order_create_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_rpc_order_rpc_order_create_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
	file_rpc_order_rpc_order_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_rpc_order_rpc_order_create_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_rpc_order_rpc_order_create_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_order_rpc_order_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_order_rpc_order_create_proto_goTypes,
		DependencyIndexes: file_rpc_order_rpc_order_create_proto_depIdxs,
		MessageInfos:      file_rpc_order_rpc_order_create_proto_msgTypes,
	}.Build()
	File_rpc_order_rpc_order_create_proto = out.File
	file_rpc_order_rpc_order_create_proto_rawDesc = nil
	file_rpc_order_rpc_order_create_proto_goTypes = nil
	file_rpc_order_rpc_order_create_proto_depIdxs = nil
}
