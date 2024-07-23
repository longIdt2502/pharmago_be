// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/prescriptions.proto

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

// / ============== ENTITIES ============
type Prescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid          string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Code          string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	CustomerId    int32                  `protobuf:"varint,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Customer      *Account               `protobuf:"bytes,5,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
	Company       int32                  `protobuf:"varint,6,opt,name=company,proto3" json:"company,omitempty"`
	DoctorId      int32                  `protobuf:"varint,7,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	Doctor        *Account               `protobuf:"bytes,8,opt,name=doctor,proto3" json:"doctor,omitempty"`
	Symptoms      *string                `protobuf:"bytes,9,opt,name=symptoms,proto3,oneof" json:"symptoms,omitempty"`
	Diagnostic    *string                `protobuf:"bytes,10,opt,name=diagnostic,proto3,oneof" json:"diagnostic,omitempty"`
	Items         []*PrescriptionItem    `protobuf:"bytes,11,rep,name=items,proto3" json:"items,omitempty"`
	Payment       []*Payment             `protobuf:"bytes,12,rep,name=payment,proto3" json:"payment,omitempty"`
	UserCreatedId int32                  `protobuf:"varint,13,opt,name=user_created_id,json=userCreatedId,proto3" json:"user_created_id,omitempty"`
	UserCreated   *Account               `protobuf:"bytes,14,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	UserUpdatedId *int32                 `protobuf:"varint,15,opt,name=user_updated_id,json=userUpdatedId,proto3,oneof" json:"user_updated_id,omitempty"`
	UserUpdated   *Account               `protobuf:"bytes,16,opt,name=user_updated,json=userUpdated,proto3,oneof" json:"user_updated,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	MbUuid        *string                `protobuf:"bytes,20,opt,name=mb_uuid,json=mbUuid,proto3,oneof" json:"mb_uuid,omitempty"`
}

func (x *Prescription) Reset() {
	*x = Prescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_prescriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prescription) ProtoMessage() {}

func (x *Prescription) ProtoReflect() protoreflect.Message {
	mi := &file_entities_prescriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prescription.ProtoReflect.Descriptor instead.
func (*Prescription) Descriptor() ([]byte, []int) {
	return file_entities_prescriptions_proto_rawDescGZIP(), []int{0}
}

func (x *Prescription) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Prescription) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Prescription) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Prescription) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Prescription) GetCustomer() *Account {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Prescription) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *Prescription) GetDoctorId() int32 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

func (x *Prescription) GetDoctor() *Account {
	if x != nil {
		return x.Doctor
	}
	return nil
}

func (x *Prescription) GetSymptoms() string {
	if x != nil && x.Symptoms != nil {
		return *x.Symptoms
	}
	return ""
}

func (x *Prescription) GetDiagnostic() string {
	if x != nil && x.Diagnostic != nil {
		return *x.Diagnostic
	}
	return ""
}

func (x *Prescription) GetItems() []*PrescriptionItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Prescription) GetPayment() []*Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

func (x *Prescription) GetUserCreatedId() int32 {
	if x != nil {
		return x.UserCreatedId
	}
	return 0
}

func (x *Prescription) GetUserCreated() *Account {
	if x != nil {
		return x.UserCreated
	}
	return nil
}

func (x *Prescription) GetUserUpdatedId() int32 {
	if x != nil && x.UserUpdatedId != nil {
		return *x.UserUpdatedId
	}
	return 0
}

func (x *Prescription) GetUserUpdated() *Account {
	if x != nil {
		return x.UserUpdated
	}
	return nil
}

func (x *Prescription) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Prescription) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Prescription) GetMbUuid() string {
	if x != nil && x.MbUuid != nil {
		return *x.MbUuid
	}
	return ""
}

type PrescriptionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AsUuid    string   `protobuf:"bytes,2,opt,name=as_uuid,json=asUuid,proto3" json:"as_uuid,omitempty"`
	VariantId int32    `protobuf:"varint,3,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	Variant   *Variant `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant,omitempty"`
	LieuDung  *string  `protobuf:"bytes,5,opt,name=lieu_dung,json=lieuDung,proto3,oneof" json:"lieu_dung,omitempty"`
	Quantity  int32    `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *PrescriptionItem) Reset() {
	*x = PrescriptionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_prescriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrescriptionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrescriptionItem) ProtoMessage() {}

func (x *PrescriptionItem) ProtoReflect() protoreflect.Message {
	mi := &file_entities_prescriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrescriptionItem.ProtoReflect.Descriptor instead.
func (*PrescriptionItem) Descriptor() ([]byte, []int) {
	return file_entities_prescriptions_proto_rawDescGZIP(), []int{1}
}

func (x *PrescriptionItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PrescriptionItem) GetAsUuid() string {
	if x != nil {
		return x.AsUuid
	}
	return ""
}

func (x *PrescriptionItem) GetVariantId() int32 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

func (x *PrescriptionItem) GetVariant() *Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *PrescriptionItem) GetLieuDung() string {
	if x != nil && x.LieuDung != nil {
		return *x.LieuDung
	}
	return ""
}

func (x *PrescriptionItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// ===================== REQ/RES ========================
type PrescriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageTrans string        `protobuf:"bytes,3,opt,name=message_trans,json=messageTrans,proto3" json:"message_trans,omitempty"`
	Log          string        `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty"`
	Details      *Prescription `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *PrescriptionResponse) Reset() {
	*x = PrescriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_prescriptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrescriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrescriptionResponse) ProtoMessage() {}

func (x *PrescriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entities_prescriptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrescriptionResponse.ProtoReflect.Descriptor instead.
func (*PrescriptionResponse) Descriptor() ([]byte, []int) {
	return file_entities_prescriptions_proto_rawDescGZIP(), []int{2}
}

func (x *PrescriptionResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PrescriptionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PrescriptionResponse) GetMessageTrans() string {
	if x != nil {
		return x.MessageTrans
	}
	return ""
}

func (x *PrescriptionResponse) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *PrescriptionResponse) GetDetails() *Prescription {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_entities_prescriptions_proto protoreflect.FileDescriptor

var file_entities_prescriptions_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x06, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x79, 0x6d,
	0x70, 0x74, 0x6f, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x73,
	0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x64, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0a, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x0f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x04, 0x52, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6d, 0x62, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06, 0x6d, 0x62, 0x55, 0x75,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x62, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x22, 0xcd,
	0x01, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x73, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x73, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x69, 0x65, 0x75, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x69, 0x65, 0x75, 0x44, 0x75, 0x6e,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x69, 0x65, 0x75, 0x5f, 0x64, 0x75, 0x6e, 0x67, 0x22, 0xa7,
	0x01, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x2a, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67,
	0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_prescriptions_proto_rawDescOnce sync.Once
	file_entities_prescriptions_proto_rawDescData = file_entities_prescriptions_proto_rawDesc
)

func file_entities_prescriptions_proto_rawDescGZIP() []byte {
	file_entities_prescriptions_proto_rawDescOnce.Do(func() {
		file_entities_prescriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_prescriptions_proto_rawDescData)
	})
	return file_entities_prescriptions_proto_rawDescData
}

var file_entities_prescriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entities_prescriptions_proto_goTypes = []interface{}{
	(*Prescription)(nil),          // 0: pb.Prescription
	(*PrescriptionItem)(nil),      // 1: pb.PrescriptionItem
	(*PrescriptionResponse)(nil),  // 2: pb.PrescriptionResponse
	(*Account)(nil),               // 3: pb.Account
	(*Payment)(nil),               // 4: pb.Payment
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Variant)(nil),               // 6: pb.Variant
}
var file_entities_prescriptions_proto_depIdxs = []int32{
	3,  // 0: pb.Prescription.customer:type_name -> pb.Account
	3,  // 1: pb.Prescription.doctor:type_name -> pb.Account
	1,  // 2: pb.Prescription.items:type_name -> pb.PrescriptionItem
	4,  // 3: pb.Prescription.payment:type_name -> pb.Payment
	3,  // 4: pb.Prescription.user_created:type_name -> pb.Account
	3,  // 5: pb.Prescription.user_updated:type_name -> pb.Account
	5,  // 6: pb.Prescription.created_at:type_name -> google.protobuf.Timestamp
	5,  // 7: pb.Prescription.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 8: pb.PrescriptionItem.variant:type_name -> pb.Variant
	0,  // 9: pb.PrescriptionResponse.details:type_name -> pb.Prescription
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_entities_prescriptions_proto_init() }
func file_entities_prescriptions_proto_init() {
	if File_entities_prescriptions_proto != nil {
		return
	}
	file_entities_account_proto_init()
	file_entities_service_proto_init()
	file_entities_order_proto_init()
	file_entities_variant_proto_init()
	file_entities_customer_proto_init()
	file_entities_appointment_schedule_proto_init()
	file_entities_payment_proto_init()
	file_rpc_customer_rpc_medical_record_create_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_prescriptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prescription); i {
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
		file_entities_prescriptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrescriptionItem); i {
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
		file_entities_prescriptions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrescriptionResponse); i {
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
	file_entities_prescriptions_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_entities_prescriptions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entities_prescriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_prescriptions_proto_goTypes,
		DependencyIndexes: file_entities_prescriptions_proto_depIdxs,
		MessageInfos:      file_entities_prescriptions_proto_msgTypes,
	}.Build()
	File_entities_prescriptions_proto = out.File
	file_entities_prescriptions_proto_rawDesc = nil
	file_entities_prescriptions_proto_goTypes = nil
	file_entities_prescriptions_proto_depIdxs = nil
}