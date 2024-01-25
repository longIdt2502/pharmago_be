// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/order.proto

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

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Variant     *Variant     `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
	Value       int32        `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	TotalPrice  float32      `protobuf:"fixed32,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Consignment *Consignment `protobuf:"bytes,5,opt,name=consignment,proto3" json:"consignment,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_entities_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_entities_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItem) GetVariant() *Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *OrderItem) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *OrderItem) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderItem) GetConsignment() *Consignment {
	if x != nil {
		return x.Consignment
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code         string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	TotalPrice   float32                `protobuf:"fixed32,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Description  string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Vat          float32                `protobuf:"fixed32,5,opt,name=vat,proto3" json:"vat,omitempty"`
	Discount     string                 `protobuf:"bytes,6,opt,name=discount,proto3" json:"discount,omitempty"`
	ServicePrice float32                `protobuf:"fixed32,7,opt,name=service_price,json=servicePrice,proto3" json:"service_price,omitempty"`
	MustPaid     float32                `protobuf:"fixed32,8,opt,name=must_paid,json=mustPaid,proto3" json:"must_paid,omitempty"`
	Customer     *CustomerDetail        `protobuf:"bytes,9,opt,name=customer,proto3" json:"customer,omitempty"`
	Address      *Address               `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Type         *SimpleData            `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	Status       *SimpleData            `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	Qr           string                 `protobuf:"bytes,13,opt,name=qr,proto3" json:"qr,omitempty"`
	Company      int32                  `protobuf:"varint,14,opt,name=company,proto3" json:"company,omitempty"`
	UserCreated  string                 `protobuf:"bytes,15,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	UserUpdated  string                 `protobuf:"bytes,16,opt,name=user_updated,json=userUpdated,proto3" json:"user_updated,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Payment      *Payment               `protobuf:"bytes,19,opt,name=payment,proto3" json:"payment,omitempty"`
	Items        []*OrderItem           `protobuf:"bytes,20,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_entities_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_entities_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Order) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetVat() float32 {
	if x != nil {
		return x.Vat
	}
	return 0
}

func (x *Order) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *Order) GetServicePrice() float32 {
	if x != nil {
		return x.ServicePrice
	}
	return 0
}

func (x *Order) GetMustPaid() float32 {
	if x != nil {
		return x.MustPaid
	}
	return 0
}

func (x *Order) GetCustomer() *CustomerDetail {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Order) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetType() *SimpleData {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Order) GetStatus() *SimpleData {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Order) GetQr() string {
	if x != nil {
		return x.Qr
	}
	return ""
}

func (x *Order) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *Order) GetUserCreated() string {
	if x != nil {
		return x.UserCreated
	}
	return ""
}

func (x *Order) GetUserUpdated() string {
	if x != nil {
		return x.UserUpdated
	}
	return ""
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Order) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type OrderPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code         string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	TotalPrice   float32                `protobuf:"fixed32,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status       *SimpleData            `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Description  string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CustomerName string                 `protobuf:"bytes,6,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	UserCreated  string                 `protobuf:"bytes,7,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *OrderPreview) Reset() {
	*x = OrderPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPreview) ProtoMessage() {}

func (x *OrderPreview) ProtoReflect() protoreflect.Message {
	mi := &file_entities_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPreview.ProtoReflect.Descriptor instead.
func (*OrderPreview) Descriptor() ([]byte, []int) {
	return file_entities_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderPreview) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderPreview) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OrderPreview) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *OrderPreview) GetStatus() *SimpleData {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *OrderPreview) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrderPreview) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *OrderPreview) GetUserCreated() string {
	if x != nil {
		return x.UserCreated
	}
	return ""
}

func (x *OrderPreview) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_entities_order_proto protoreflect.FileDescriptor

var file_entities_order_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xb3, 0x05, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x76, 0x61, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6d, 0x75, 0x73, 0x74, 0x50, 0x61, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x71, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x71, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0xa0, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f,
	0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_order_proto_rawDescOnce sync.Once
	file_entities_order_proto_rawDescData = file_entities_order_proto_rawDesc
)

func file_entities_order_proto_rawDescGZIP() []byte {
	file_entities_order_proto_rawDescOnce.Do(func() {
		file_entities_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_order_proto_rawDescData)
	})
	return file_entities_order_proto_rawDescData
}

var file_entities_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entities_order_proto_goTypes = []interface{}{
	(*OrderItem)(nil),             // 0: pb.OrderItem
	(*Order)(nil),                 // 1: pb.Order
	(*OrderPreview)(nil),          // 2: pb.OrderPreview
	(*Variant)(nil),               // 3: pb.Variant
	(*Consignment)(nil),           // 4: pb.Consignment
	(*CustomerDetail)(nil),        // 5: pb.CustomerDetail
	(*Address)(nil),               // 6: pb.Address
	(*SimpleData)(nil),            // 7: pb.SimpleData
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Payment)(nil),               // 9: pb.Payment
}
var file_entities_order_proto_depIdxs = []int32{
	3,  // 0: pb.OrderItem.variant:type_name -> pb.Variant
	4,  // 1: pb.OrderItem.consignment:type_name -> pb.Consignment
	5,  // 2: pb.Order.customer:type_name -> pb.CustomerDetail
	6,  // 3: pb.Order.address:type_name -> pb.Address
	7,  // 4: pb.Order.type:type_name -> pb.SimpleData
	7,  // 5: pb.Order.status:type_name -> pb.SimpleData
	8,  // 6: pb.Order.created_at:type_name -> google.protobuf.Timestamp
	8,  // 7: pb.Order.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 8: pb.Order.payment:type_name -> pb.Payment
	0,  // 9: pb.Order.items:type_name -> pb.OrderItem
	7,  // 10: pb.OrderPreview.status:type_name -> pb.SimpleData
	8,  // 11: pb.OrderPreview.created_at:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_entities_order_proto_init() }
func file_entities_order_proto_init() {
	if File_entities_order_proto != nil {
		return
	}
	file_entities_company_proto_init()
	file_entities_customer_proto_init()
	file_entities_address_proto_init()
	file_entities_simple_data_proto_init()
	file_entities_payment_proto_init()
	file_entities_variant_proto_init()
	file_entities_consigment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_entities_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_entities_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPreview); i {
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
			RawDescriptor: file_entities_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_order_proto_goTypes,
		DependencyIndexes: file_entities_order_proto_depIdxs,
		MessageInfos:      file_entities_order_proto_msgTypes,
	}.Build()
	File_entities_order_proto = out.File
	file_entities_order_proto_rawDesc = nil
	file_entities_order_proto_goTypes = nil
	file_entities_order_proto_depIdxs = nil
}
