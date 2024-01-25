// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/ticket.proto

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

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code   string      `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Type   *SimpleData `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status *SimpleData `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Note   string      `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	Qr     string      `protobuf:"bytes,6,opt,name=qr,proto3" json:"qr,omitempty"`
	// Address export_to = 7;
	// Address import_from = 8;
	TotalPrice   float32        `protobuf:"fixed32,9,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Supplier     *Supplier      `protobuf:"bytes,10,opt,name=supplier,proto3,oneof" json:"supplier,omitempty"`
	Customer     *Customer      `protobuf:"bytes,11,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
	Consignments []*Consignment `protobuf:"bytes,12,rep,name=consignments,proto3" json:"consignments,omitempty"`
	Warehouse    *Warehouse     `protobuf:"bytes,13,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	UserCreated  string         `protobuf:"bytes,14,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	UserUpdated  string         `protobuf:"bytes,15,opt,name=user_updated,json=userUpdated,proto3" json:"user_updated,omitempty"`
	// google.protobuf.Timestamp updated_at = 16;
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_entities_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_entities_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ticket) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Ticket) GetType() *SimpleData {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Ticket) GetStatus() *SimpleData {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Ticket) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Ticket) GetQr() string {
	if x != nil {
		return x.Qr
	}
	return ""
}

func (x *Ticket) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Ticket) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

func (x *Ticket) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Ticket) GetConsignments() []*Consignment {
	if x != nil {
		return x.Consignments
	}
	return nil
}

func (x *Ticket) GetWarehouse() *Warehouse {
	if x != nil {
		return x.Warehouse
	}
	return nil
}

func (x *Ticket) GetUserCreated() string {
	if x != nil {
		return x.UserCreated
	}
	return ""
}

func (x *Ticket) GetUserUpdated() string {
	if x != nil {
		return x.UserUpdated
	}
	return ""
}

func (x *Ticket) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TicketPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code           string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Type           *SimpleData            `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status         *SimpleData            `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Note           string                 `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	Qr             string                 `protobuf:"bytes,6,opt,name=qr,proto3" json:"qr,omitempty"`
	TotalItems     int32                  `protobuf:"varint,7,opt,name=totalItems,proto3" json:"totalItems,omitempty"`
	TotalItemsType int32                  `protobuf:"varint,8,opt,name=totalItemsType,proto3" json:"totalItemsType,omitempty"`
	TotalPrice     float32                `protobuf:"fixed32,9,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	WarehouseName  string                 `protobuf:"bytes,10,opt,name=warehouse_name,json=warehouseName,proto3" json:"warehouse_name,omitempty"`
	UserCreated    string                 `protobuf:"bytes,11,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *TicketPreview) Reset() {
	*x = TicketPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketPreview) ProtoMessage() {}

func (x *TicketPreview) ProtoReflect() protoreflect.Message {
	mi := &file_entities_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketPreview.ProtoReflect.Descriptor instead.
func (*TicketPreview) Descriptor() ([]byte, []int) {
	return file_entities_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *TicketPreview) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TicketPreview) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TicketPreview) GetType() *SimpleData {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *TicketPreview) GetStatus() *SimpleData {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *TicketPreview) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *TicketPreview) GetQr() string {
	if x != nil {
		return x.Qr
	}
	return ""
}

func (x *TicketPreview) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *TicketPreview) GetTotalItemsType() int32 {
	if x != nil {
		return x.TotalItemsType
	}
	return 0
}

func (x *TicketPreview) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *TicketPreview) GetWarehouseName() string {
	if x != nil {
		return x.WarehouseName
	}
	return ""
}

func (x *TicketPreview) GetUserCreated() string {
	if x != nil {
		return x.UserCreated
	}
	return ""
}

func (x *TicketPreview) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_entities_ticket_proto protoreflect.FileDescriptor

var file_entities_ticket_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1a, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x04, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x71, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x01, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x22, 0x91, 0x03, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x71, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x71, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35,
	0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_ticket_proto_rawDescOnce sync.Once
	file_entities_ticket_proto_rawDescData = file_entities_ticket_proto_rawDesc
)

func file_entities_ticket_proto_rawDescGZIP() []byte {
	file_entities_ticket_proto_rawDescOnce.Do(func() {
		file_entities_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_ticket_proto_rawDescData)
	})
	return file_entities_ticket_proto_rawDescData
}

var file_entities_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entities_ticket_proto_goTypes = []interface{}{
	(*Ticket)(nil),                // 0: pb.Ticket
	(*TicketPreview)(nil),         // 1: pb.TicketPreview
	(*SimpleData)(nil),            // 2: pb.SimpleData
	(*Supplier)(nil),              // 3: pb.Supplier
	(*Customer)(nil),              // 4: pb.Customer
	(*Consignment)(nil),           // 5: pb.Consignment
	(*Warehouse)(nil),             // 6: pb.Warehouse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_entities_ticket_proto_depIdxs = []int32{
	2,  // 0: pb.Ticket.type:type_name -> pb.SimpleData
	2,  // 1: pb.Ticket.status:type_name -> pb.SimpleData
	3,  // 2: pb.Ticket.supplier:type_name -> pb.Supplier
	4,  // 3: pb.Ticket.customer:type_name -> pb.Customer
	5,  // 4: pb.Ticket.consignments:type_name -> pb.Consignment
	6,  // 5: pb.Ticket.warehouse:type_name -> pb.Warehouse
	7,  // 6: pb.Ticket.created_at:type_name -> google.protobuf.Timestamp
	2,  // 7: pb.TicketPreview.type:type_name -> pb.SimpleData
	2,  // 8: pb.TicketPreview.status:type_name -> pb.SimpleData
	7,  // 9: pb.TicketPreview.created_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_entities_ticket_proto_init() }
func file_entities_ticket_proto_init() {
	if File_entities_ticket_proto != nil {
		return
	}
	file_entities_simple_data_proto_init()
	file_entities_address_proto_init()
	file_entities_supplier_proto_init()
	file_entities_warehouse_proto_init()
	file_entities_consigment_proto_init()
	file_entities_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_entities_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketPreview); i {
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
	file_entities_ticket_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entities_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_ticket_proto_goTypes,
		DependencyIndexes: file_entities_ticket_proto_depIdxs,
		MessageInfos:      file_entities_ticket_proto_msgTypes,
	}.Build()
	File_entities_ticket_proto = out.File
	file_entities_ticket_proto_rawDesc = nil
	file_entities_ticket_proto_goTypes = nil
	file_entities_ticket_proto_depIdxs = nil
}
