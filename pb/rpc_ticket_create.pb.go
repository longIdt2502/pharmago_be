// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: rpc/warehouse/rpc_ticket_create.proto

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

type TicketCreateRequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       *string `protobuf:"bytes,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Type       string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Status     string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Note       string  `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	TotalPrice float32 `protobuf:"fixed32,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	ExportTo   int32   `protobuf:"varint,6,opt,name=export_to,json=exportTo,proto3" json:"export_to,omitempty"`
	ImportFrom int32   `protobuf:"varint,7,opt,name=import_from,json=importFrom,proto3" json:"import_from,omitempty"`
	Warehouse  int32   `protobuf:"varint,8,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
}

func (x *TicketCreateRequestInfo) Reset() {
	*x = TicketCreateRequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCreateRequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCreateRequestInfo) ProtoMessage() {}

func (x *TicketCreateRequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCreateRequestInfo.ProtoReflect.Descriptor instead.
func (*TicketCreateRequestInfo) Descriptor() ([]byte, []int) {
	return file_rpc_warehouse_rpc_ticket_create_proto_rawDescGZIP(), []int{0}
}

func (x *TicketCreateRequestInfo) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *TicketCreateRequestInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TicketCreateRequestInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TicketCreateRequestInfo) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *TicketCreateRequestInfo) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *TicketCreateRequestInfo) GetExportTo() int32 {
	if x != nil {
		return x.ExportTo
	}
	return 0
}

func (x *TicketCreateRequestInfo) GetImportFrom() int32 {
	if x != nil {
		return x.ImportFrom
	}
	return 0
}

func (x *TicketCreateRequestInfo) GetWarehouse() int32 {
	if x != nil {
		return x.Warehouse
	}
	return 0
}

type TicketCreateRequestConsignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       *string                `protobuf:"bytes,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Quantity   int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Inventory  int32                  `protobuf:"varint,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Variant    int32                  `protobuf:"varint,4,opt,name=variant,proto3" json:"variant,omitempty"`
	ExpiredAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	ProducedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=produced_at,json=producedAt,proto3" json:"produced_at,omitempty"`
}

func (x *TicketCreateRequestConsignment) Reset() {
	*x = TicketCreateRequestConsignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCreateRequestConsignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCreateRequestConsignment) ProtoMessage() {}

func (x *TicketCreateRequestConsignment) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCreateRequestConsignment.ProtoReflect.Descriptor instead.
func (*TicketCreateRequestConsignment) Descriptor() ([]byte, []int) {
	return file_rpc_warehouse_rpc_ticket_create_proto_rawDescGZIP(), []int{1}
}

func (x *TicketCreateRequestConsignment) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *TicketCreateRequestConsignment) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *TicketCreateRequestConsignment) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *TicketCreateRequestConsignment) GetVariant() int32 {
	if x != nil {
		return x.Variant
	}
	return 0
}

func (x *TicketCreateRequestConsignment) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *TicketCreateRequestConsignment) GetProducedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ProducedAt
	}
	return nil
}

type TicketCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket      *TicketCreateRequestInfo          `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Consignment []*TicketCreateRequestConsignment `protobuf:"bytes,2,rep,name=consignment,proto3" json:"consignment,omitempty"`
}

func (x *TicketCreateRequest) Reset() {
	*x = TicketCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCreateRequest) ProtoMessage() {}

func (x *TicketCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCreateRequest.ProtoReflect.Descriptor instead.
func (*TicketCreateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_warehouse_rpc_ticket_create_proto_rawDescGZIP(), []int{2}
}

func (x *TicketCreateRequest) GetTicket() *TicketCreateRequestInfo {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *TicketCreateRequest) GetConsignment() []*TicketCreateRequestConsignment {
	if x != nil {
		return x.Consignment
	}
	return nil
}

type TicketCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details int32  `protobuf:"varint,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *TicketCreateResponse) Reset() {
	*x = TicketCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketCreateResponse) ProtoMessage() {}

func (x *TicketCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketCreateResponse.ProtoReflect.Descriptor instead.
func (*TicketCreateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_warehouse_rpc_ticket_create_proto_rawDescGZIP(), []int{3}
}

func (x *TicketCreateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TicketCreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TicketCreateResponse) GetDetails() int32 {
	if x != nil {
		return x.Details
	}
	return 0
}

var File_rpc_warehouse_rpc_ticket_create_proto protoreflect.FileDescriptor

var file_rpc_warehouse_rpc_ticket_create_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a,
	0x17, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x8e, 0x02, 0x0a, 0x1e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x33, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x44, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c,
	0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f,
	0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_warehouse_rpc_ticket_create_proto_rawDescOnce sync.Once
	file_rpc_warehouse_rpc_ticket_create_proto_rawDescData = file_rpc_warehouse_rpc_ticket_create_proto_rawDesc
)

func file_rpc_warehouse_rpc_ticket_create_proto_rawDescGZIP() []byte {
	file_rpc_warehouse_rpc_ticket_create_proto_rawDescOnce.Do(func() {
		file_rpc_warehouse_rpc_ticket_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_warehouse_rpc_ticket_create_proto_rawDescData)
	})
	return file_rpc_warehouse_rpc_ticket_create_proto_rawDescData
}

var file_rpc_warehouse_rpc_ticket_create_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_warehouse_rpc_ticket_create_proto_goTypes = []interface{}{
	(*TicketCreateRequestInfo)(nil),        // 0: pb.TicketCreateRequestInfo
	(*TicketCreateRequestConsignment)(nil), // 1: pb.TicketCreateRequestConsignment
	(*TicketCreateRequest)(nil),            // 2: pb.TicketCreateRequest
	(*TicketCreateResponse)(nil),           // 3: pb.TicketCreateResponse
	(*timestamppb.Timestamp)(nil),          // 4: google.protobuf.Timestamp
}
var file_rpc_warehouse_rpc_ticket_create_proto_depIdxs = []int32{
	4, // 0: pb.TicketCreateRequestConsignment.expired_at:type_name -> google.protobuf.Timestamp
	4, // 1: pb.TicketCreateRequestConsignment.produced_at:type_name -> google.protobuf.Timestamp
	0, // 2: pb.TicketCreateRequest.ticket:type_name -> pb.TicketCreateRequestInfo
	1, // 3: pb.TicketCreateRequest.consignment:type_name -> pb.TicketCreateRequestConsignment
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_warehouse_rpc_ticket_create_proto_init() }
func file_rpc_warehouse_rpc_ticket_create_proto_init() {
	if File_rpc_warehouse_rpc_ticket_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketCreateRequestInfo); i {
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
		file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketCreateRequestConsignment); i {
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
		file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketCreateRequest); i {
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
		file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketCreateResponse); i {
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
	file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_rpc_warehouse_rpc_ticket_create_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_warehouse_rpc_ticket_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_warehouse_rpc_ticket_create_proto_goTypes,
		DependencyIndexes: file_rpc_warehouse_rpc_ticket_create_proto_depIdxs,
		MessageInfos:      file_rpc_warehouse_rpc_ticket_create_proto_msgTypes,
	}.Build()
	File_rpc_warehouse_rpc_ticket_create_proto = out.File
	file_rpc_warehouse_rpc_ticket_create_proto_rawDesc = nil
	file_rpc_warehouse_rpc_ticket_create_proto_goTypes = nil
	file_rpc_warehouse_rpc_ticket_create_proto_depIdxs = nil
}
