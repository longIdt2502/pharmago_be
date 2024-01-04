// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: rpc/order/rpc_order_list.proto

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

type OrderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company      int32                  `protobuf:"varint,1,opt,name=company,proto3" json:"company,omitempty"`
	Warehouse    *int32                 `protobuf:"varint,2,opt,name=warehouse,proto3,oneof" json:"warehouse,omitempty"`
	Search       *string                `protobuf:"bytes,3,opt,name=search,proto3,oneof" json:"search,omitempty"`
	Page         *int32                 `protobuf:"varint,4,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit        *int32                 `protobuf:"varint,5,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Status       *string                `protobuf:"bytes,6,opt,name=status,proto3,oneof" json:"status,omitempty"`
	OrderBy      *string                `protobuf:"bytes,7,opt,name=order_by,json=orderBy,proto3,oneof" json:"order_by,omitempty"`
	CreatedStart *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_start,json=createdStart,proto3,oneof" json:"created_start,omitempty"`
	CreatedEnd   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_end,json=createdEnd,proto3,oneof" json:"created_end,omitempty"`
	UpdatedStart *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_start,json=updatedStart,proto3,oneof" json:"updated_start,omitempty"`
	UpdatedEnd   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_end,json=updatedEnd,proto3,oneof" json:"updated_end,omitempty"`
}

func (x *OrderListRequest) Reset() {
	*x = OrderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListRequest) ProtoMessage() {}

func (x *OrderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListRequest.ProtoReflect.Descriptor instead.
func (*OrderListRequest) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_list_proto_rawDescGZIP(), []int{0}
}

func (x *OrderListRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *OrderListRequest) GetWarehouse() int32 {
	if x != nil && x.Warehouse != nil {
		return *x.Warehouse
	}
	return 0
}

func (x *OrderListRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

func (x *OrderListRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *OrderListRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *OrderListRequest) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *OrderListRequest) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *OrderListRequest) GetCreatedStart() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedStart
	}
	return nil
}

func (x *OrderListRequest) GetCreatedEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedEnd
	}
	return nil
}

func (x *OrderListRequest) GetUpdatedStart() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedStart
	}
	return nil
}

func (x *OrderListRequest) GetUpdatedEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedEnd
	}
	return nil
}

type OrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*OrderPreview         `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	Count   *OrderListResponseCount `protobuf:"bytes,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *OrderListResponse) Reset() {
	*x = OrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResponse) ProtoMessage() {}

func (x *OrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResponse.ProtoReflect.Descriptor instead.
func (*OrderListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_list_proto_rawDescGZIP(), []int{1}
}

func (x *OrderListResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OrderListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderListResponse) GetDetails() []*OrderPreview {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *OrderListResponse) GetCount() *OrderListResponseCount {
	if x != nil {
		return x.Count
	}
	return nil
}

type OrderListResponseCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Draft     int32 `protobuf:"varint,1,opt,name=draft,proto3" json:"draft,omitempty"`
	InProcess int32 `protobuf:"varint,2,opt,name=in_process,json=inProcess,proto3" json:"in_process,omitempty"`
	Complete  int32 `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	Cancel    int32 `protobuf:"varint,4,opt,name=cancel,proto3" json:"cancel,omitempty"`
}

func (x *OrderListResponseCount) Reset() {
	*x = OrderListResponseCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_order_rpc_order_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResponseCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResponseCount) ProtoMessage() {}

func (x *OrderListResponseCount) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_order_rpc_order_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResponseCount.ProtoReflect.Descriptor instead.
func (*OrderListResponseCount) Descriptor() ([]byte, []int) {
	return file_rpc_order_rpc_order_list_proto_rawDescGZIP(), []int{2}
}

func (x *OrderListResponseCount) GetDraft() int32 {
	if x != nil {
		return x.Draft
	}
	return 0
}

func (x *OrderListResponseCount) GetInProcess() int32 {
	if x != nil {
		return x.InProcess
	}
	return 0
}

func (x *OrderListResponseCount) GetComplete() int32 {
	if x != nil {
		return x.Complete
	}
	return 0
}

func (x *OrderListResponseCount) GetCancel() int32 {
	if x != nil {
		return x.Cancel
	}
	return 0
}

var File_rpc_order_rpc_order_list_proto protoreflect.FileDescriptor

var file_rpc_order_rpc_order_list_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x04, 0x0a, 0x10,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x21, 0x0a, 0x09, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x06,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x07, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x08, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x09, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x65, 0x6e, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e,
	0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_order_rpc_order_list_proto_rawDescOnce sync.Once
	file_rpc_order_rpc_order_list_proto_rawDescData = file_rpc_order_rpc_order_list_proto_rawDesc
)

func file_rpc_order_rpc_order_list_proto_rawDescGZIP() []byte {
	file_rpc_order_rpc_order_list_proto_rawDescOnce.Do(func() {
		file_rpc_order_rpc_order_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_order_rpc_order_list_proto_rawDescData)
	})
	return file_rpc_order_rpc_order_list_proto_rawDescData
}

var file_rpc_order_rpc_order_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_order_rpc_order_list_proto_goTypes = []interface{}{
	(*OrderListRequest)(nil),       // 0: pb.OrderListRequest
	(*OrderListResponse)(nil),      // 1: pb.OrderListResponse
	(*OrderListResponseCount)(nil), // 2: pb.OrderListResponseCount
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*OrderPreview)(nil),           // 4: pb.OrderPreview
}
var file_rpc_order_rpc_order_list_proto_depIdxs = []int32{
	3, // 0: pb.OrderListRequest.created_start:type_name -> google.protobuf.Timestamp
	3, // 1: pb.OrderListRequest.created_end:type_name -> google.protobuf.Timestamp
	3, // 2: pb.OrderListRequest.updated_start:type_name -> google.protobuf.Timestamp
	3, // 3: pb.OrderListRequest.updated_end:type_name -> google.protobuf.Timestamp
	4, // 4: pb.OrderListResponse.details:type_name -> pb.OrderPreview
	2, // 5: pb.OrderListResponse.count:type_name -> pb.OrderListResponseCount
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rpc_order_rpc_order_list_proto_init() }
func file_rpc_order_rpc_order_list_proto_init() {
	if File_rpc_order_rpc_order_list_proto != nil {
		return
	}
	file_entities_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_order_rpc_order_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListRequest); i {
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
		file_rpc_order_rpc_order_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResponse); i {
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
		file_rpc_order_rpc_order_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResponseCount); i {
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
	file_rpc_order_rpc_order_list_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_order_rpc_order_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_order_rpc_order_list_proto_goTypes,
		DependencyIndexes: file_rpc_order_rpc_order_list_proto_depIdxs,
		MessageInfos:      file_rpc_order_rpc_order_list_proto_msgTypes,
	}.Build()
	File_rpc_order_rpc_order_list_proto = out.File
	file_rpc_order_rpc_order_list_proto_rawDesc = nil
	file_rpc_order_rpc_order_list_proto_goTypes = nil
	file_rpc_order_rpc_order_list_proto_depIdxs = nil
}
