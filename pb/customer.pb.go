// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: entities/customer.proto

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

type MedicalRecordType int32

const (
	MedicalRecordType_test       MedicalRecordType = 0
	MedicalRecordType_patient    MedicalRecordType = 1
	MedicalRecordType_diagnostic MedicalRecordType = 2
)

// Enum value maps for MedicalRecordType.
var (
	MedicalRecordType_name = map[int32]string{
		0: "test",
		1: "patient",
		2: "diagnostic",
	}
	MedicalRecordType_value = map[string]int32{
		"test":       0,
		"patient":    1,
		"diagnostic": 2,
	}
)

func (x MedicalRecordType) Enum() *MedicalRecordType {
	p := new(MedicalRecordType)
	*p = x
	return p
}

func (x MedicalRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MedicalRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_entities_customer_proto_enumTypes[0].Descriptor()
}

func (MedicalRecordType) Type() protoreflect.EnumType {
	return &file_entities_customer_proto_enumTypes[0]
}

func (x MedicalRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MedicalRecordType.Descriptor instead.
func (MedicalRecordType) EnumDescriptor() ([]byte, []int) {
	return file_entities_customer_proto_rawDescGZIP(), []int{0}
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code         string        `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	FullName     string        `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Company      int32         `protobuf:"varint,4,opt,name=company,proto3" json:"company,omitempty"`
	Address      *Address      `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phone        string        `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Email        *string       `protobuf:"bytes,7,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Orders       int32         `protobuf:"varint,8,opt,name=orders,proto3" json:"orders,omitempty"`
	Revenue      float32       `protobuf:"fixed32,9,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Conversation *Conversation `protobuf:"bytes,10,opt,name=conversation,proto3,oneof" json:"conversation,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_entities_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_entities_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Customer) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Customer) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *Customer) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Customer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Customer) GetOrders() int32 {
	if x != nil {
		return x.Orders
	}
	return 0
}

func (x *Customer) GetRevenue() float32 {
	if x != nil {
		return x.Revenue
	}
	return 0
}

func (x *Customer) GetConversation() *Conversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

type CustomerDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code           string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	FullName       string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Company        int32                  `protobuf:"varint,4,opt,name=company,proto3" json:"company,omitempty"`
	Address        *Address               `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phone          string                 `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Email          *string                `protobuf:"bytes,7,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Birthday       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=birthday,proto3,oneof" json:"birthday,omitempty"`
	Gender         *string                `protobuf:"bytes,9,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	Title          *string                `protobuf:"bytes,10,opt,name=title,proto3,oneof" json:"title,omitempty"`
	LicenseDate    *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=license_date,json=licenseDate,proto3,oneof" json:"license_date,omitempty"`
	IssuedBy       *string                `protobuf:"bytes,12,opt,name=issued_by,json=issuedBy,proto3,oneof" json:"issued_by,omitempty"`
	ContactName    *string                `protobuf:"bytes,13,opt,name=contact_name,json=contactName,proto3,oneof" json:"contact_name,omitempty"`
	ContactTitle   *string                `protobuf:"bytes,14,opt,name=contact_title,json=contactTitle,proto3,oneof" json:"contact_title,omitempty"`
	ContactPhone   *string                `protobuf:"bytes,15,opt,name=contact_phone,json=contactPhone,proto3,oneof" json:"contact_phone,omitempty"`
	ContactEmail   *string                `protobuf:"bytes,16,opt,name=contact_email,json=contactEmail,proto3,oneof" json:"contact_email,omitempty"`
	ContactAddress *Address               `protobuf:"bytes,17,opt,name=contact_address,json=contactAddress,proto3,oneof" json:"contact_address,omitempty"`
	AccountNumber  *string                `protobuf:"bytes,18,opt,name=account_number,json=accountNumber,proto3,oneof" json:"account_number,omitempty"`
	BankName       *string                `protobuf:"bytes,19,opt,name=bank_name,json=bankName,proto3,oneof" json:"bank_name,omitempty"`
	BankBranch     *string                `protobuf:"bytes,20,opt,name=bank_branch,json=bankBranch,proto3,oneof" json:"bank_branch,omitempty"`
	License        *string                `protobuf:"bytes,21,opt,name=license,proto3,oneof" json:"license,omitempty"`
	UserCreated    string                 `protobuf:"bytes,22,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserUpdated    *string                `protobuf:"bytes,24,opt,name=user_updated,json=userUpdated,proto3,oneof" json:"user_updated,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *CustomerDetail) Reset() {
	*x = CustomerDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDetail) ProtoMessage() {}

func (x *CustomerDetail) ProtoReflect() protoreflect.Message {
	mi := &file_entities_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDetail.ProtoReflect.Descriptor instead.
func (*CustomerDetail) Descriptor() ([]byte, []int) {
	return file_entities_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerDetail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CustomerDetail) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CustomerDetail) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *CustomerDetail) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CustomerDetail) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CustomerDetail) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *CustomerDetail) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *CustomerDetail) GetGender() string {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return ""
}

func (x *CustomerDetail) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *CustomerDetail) GetLicenseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LicenseDate
	}
	return nil
}

func (x *CustomerDetail) GetIssuedBy() string {
	if x != nil && x.IssuedBy != nil {
		return *x.IssuedBy
	}
	return ""
}

func (x *CustomerDetail) GetContactName() string {
	if x != nil && x.ContactName != nil {
		return *x.ContactName
	}
	return ""
}

func (x *CustomerDetail) GetContactTitle() string {
	if x != nil && x.ContactTitle != nil {
		return *x.ContactTitle
	}
	return ""
}

func (x *CustomerDetail) GetContactPhone() string {
	if x != nil && x.ContactPhone != nil {
		return *x.ContactPhone
	}
	return ""
}

func (x *CustomerDetail) GetContactEmail() string {
	if x != nil && x.ContactEmail != nil {
		return *x.ContactEmail
	}
	return ""
}

func (x *CustomerDetail) GetContactAddress() *Address {
	if x != nil {
		return x.ContactAddress
	}
	return nil
}

func (x *CustomerDetail) GetAccountNumber() string {
	if x != nil && x.AccountNumber != nil {
		return *x.AccountNumber
	}
	return ""
}

func (x *CustomerDetail) GetBankName() string {
	if x != nil && x.BankName != nil {
		return *x.BankName
	}
	return ""
}

func (x *CustomerDetail) GetBankBranch() string {
	if x != nil && x.BankBranch != nil {
		return *x.BankBranch
	}
	return ""
}

func (x *CustomerDetail) GetLicense() string {
	if x != nil && x.License != nil {
		return *x.License
	}
	return ""
}

func (x *CustomerDetail) GetUserCreated() string {
	if x != nil {
		return x.UserCreated
	}
	return ""
}

func (x *CustomerDetail) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CustomerDetail) GetUserUpdated() string {
	if x != nil && x.UserUpdated != nil {
		return *x.UserUpdated
	}
	return ""
}

func (x *CustomerDetail) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type MedicalRecordLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // serial
	Uuid                string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
	Type                MedicalRecordType      `protobuf:"varint,3,opt,name=type,proto3,enum=pb.MedicalRecordType" json:"type,omitempty"`
	TypeName            string                 `protobuf:"bytes,4,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"` // medical_record_type
	Title               string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`                       // varchar
	Url                 string                 `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`                           // varchar
	Size                int32                  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Customer            int32                  `protobuf:"varint,8,opt,name=customer,proto3" json:"customer,omitempty"`                                                       // serial
	AppointmentSchedule *string                `protobuf:"bytes,9,opt,name=appointment_schedule,json=appointmentSchedule,proto3,oneof" json:"appointment_schedule,omitempty"` // uuid
	UserCreated         int32                  `protobuf:"varint,10,opt,name=user_created,json=userCreated,proto3" json:"user_created,omitempty"`                             // serial
	CreatedAt           *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                    // timestamp
}

func (x *MedicalRecordLink) Reset() {
	*x = MedicalRecordLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicalRecordLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicalRecordLink) ProtoMessage() {}

func (x *MedicalRecordLink) ProtoReflect() protoreflect.Message {
	mi := &file_entities_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicalRecordLink.ProtoReflect.Descriptor instead.
func (*MedicalRecordLink) Descriptor() ([]byte, []int) {
	return file_entities_customer_proto_rawDescGZIP(), []int{2}
}

func (x *MedicalRecordLink) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MedicalRecordLink) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MedicalRecordLink) GetType() MedicalRecordType {
	if x != nil {
		return x.Type
	}
	return MedicalRecordType_test
}

func (x *MedicalRecordLink) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *MedicalRecordLink) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MedicalRecordLink) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MedicalRecordLink) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MedicalRecordLink) GetCustomer() int32 {
	if x != nil {
		return x.Customer
	}
	return 0
}

func (x *MedicalRecordLink) GetAppointmentSchedule() string {
	if x != nil && x.AppointmentSchedule != nil {
		return *x.AppointmentSchedule
	}
	return ""
}

func (x *MedicalRecordLink) GetUserCreated() int32 {
	if x != nil {
		return x.UserCreated
	}
	return 0
}

func (x *MedicalRecordLink) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_entities_customer_proto protoreflect.FileDescriptor

var file_entities_customer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x72, 0x65,
	0x76, 0x65, 0x6e, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdb, 0x09, 0x0a, 0x0e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x3b, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x01, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x04, 0x52, 0x0b, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52,
	0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x06, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x39, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x0a, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0c, 0x52, 0x08,
	0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x0d, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0e, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x0f, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x10, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x86, 0x03, 0x0a, 0x11, 0x4d, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2a, 0x3a, 0x0a, 0x11, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x10, 0x02, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61,
	0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d,
	0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_entities_customer_proto_rawDescOnce sync.Once
	file_entities_customer_proto_rawDescData = file_entities_customer_proto_rawDesc
)

func file_entities_customer_proto_rawDescGZIP() []byte {
	file_entities_customer_proto_rawDescOnce.Do(func() {
		file_entities_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_customer_proto_rawDescData)
	})
	return file_entities_customer_proto_rawDescData
}

var file_entities_customer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entities_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entities_customer_proto_goTypes = []interface{}{
	(MedicalRecordType)(0),        // 0: pb.MedicalRecordType
	(*Customer)(nil),              // 1: pb.Customer
	(*CustomerDetail)(nil),        // 2: pb.CustomerDetail
	(*MedicalRecordLink)(nil),     // 3: pb.MedicalRecordLink
	(*Address)(nil),               // 4: pb.Address
	(*Conversation)(nil),          // 5: pb.Conversation
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_entities_customer_proto_depIdxs = []int32{
	4,  // 0: pb.Customer.address:type_name -> pb.Address
	5,  // 1: pb.Customer.conversation:type_name -> pb.Conversation
	4,  // 2: pb.CustomerDetail.address:type_name -> pb.Address
	6,  // 3: pb.CustomerDetail.birthday:type_name -> google.protobuf.Timestamp
	6,  // 4: pb.CustomerDetail.license_date:type_name -> google.protobuf.Timestamp
	4,  // 5: pb.CustomerDetail.contact_address:type_name -> pb.Address
	6,  // 6: pb.CustomerDetail.created_at:type_name -> google.protobuf.Timestamp
	6,  // 7: pb.CustomerDetail.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 8: pb.MedicalRecordLink.type:type_name -> pb.MedicalRecordType
	6,  // 9: pb.MedicalRecordLink.created_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_entities_customer_proto_init() }
func file_entities_customer_proto_init() {
	if File_entities_customer_proto != nil {
		return
	}
	file_entities_address_proto_init()
	file_entities_conversation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_entities_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDetail); i {
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
		file_entities_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicalRecordLink); i {
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
	file_entities_customer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_entities_customer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_entities_customer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entities_customer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_customer_proto_goTypes,
		DependencyIndexes: file_entities_customer_proto_depIdxs,
		EnumInfos:         file_entities_customer_proto_enumTypes,
		MessageInfos:      file_entities_customer_proto_msgTypes,
	}.Build()
	File_entities_customer_proto = out.File
	file_entities_customer_proto_rawDesc = nil
	file_entities_customer_proto_goTypes = nil
	file_entities_customer_proto_depIdxs = nil
}
