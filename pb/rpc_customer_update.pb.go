// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/customer/rpc_customer_update.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type CustomerUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code           *string                `protobuf:"bytes,2,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone          string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Address        *AddressPayload        `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Birthday       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=birthday,proto3,oneof" json:"birthday,omitempty"`
	Gender         *int32                 `protobuf:"varint,7,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	Group          *int32                 `protobuf:"varint,8,opt,name=group,proto3,oneof" json:"group,omitempty"`
	Company        int32                  `protobuf:"varint,9,opt,name=company,proto3" json:"company,omitempty"`
	Title          *string                `protobuf:"bytes,10,opt,name=title,proto3,oneof" json:"title,omitempty"`
	LicenseDate    *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=license_date,json=licenseDate,proto3,oneof" json:"license_date,omitempty"`
	ContactName    *string                `protobuf:"bytes,12,opt,name=contact_name,json=contactName,proto3,oneof" json:"contact_name,omitempty"`
	ContactTitle   *string                `protobuf:"bytes,13,opt,name=contact_title,json=contactTitle,proto3,oneof" json:"contact_title,omitempty"`
	ContactPhone   *string                `protobuf:"bytes,14,opt,name=contact_phone,json=contactPhone,proto3,oneof" json:"contact_phone,omitempty"`
	ContactEmail   *string                `protobuf:"bytes,15,opt,name=contact_email,json=contactEmail,proto3,oneof" json:"contact_email,omitempty"`
	ContactAddress *AddressPayload        `protobuf:"bytes,16,opt,name=contact_address,json=contactAddress,proto3,oneof" json:"contact_address,omitempty"`
	AccountNumber  *string                `protobuf:"bytes,17,opt,name=account_number,json=accountNumber,proto3,oneof" json:"account_number,omitempty"`
	BankName       *string                `protobuf:"bytes,18,opt,name=bank_name,json=bankName,proto3,oneof" json:"bank_name,omitempty"`
	BankBranch     *string                `protobuf:"bytes,19,opt,name=bank_branch,json=bankBranch,proto3,oneof" json:"bank_branch,omitempty"`
}

func (x *CustomerUpdateRequest) Reset() {
	*x = CustomerUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_rpc_customer_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUpdateRequest) ProtoMessage() {}

func (x *CustomerUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_rpc_customer_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUpdateRequest.ProtoReflect.Descriptor instead.
func (*CustomerUpdateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_customer_rpc_customer_update_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerUpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerUpdateRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *CustomerUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerUpdateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CustomerUpdateRequest) GetAddress() *AddressPayload {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CustomerUpdateRequest) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *CustomerUpdateRequest) GetGender() int32 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *CustomerUpdateRequest) GetGroup() int32 {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return 0
}

func (x *CustomerUpdateRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *CustomerUpdateRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *CustomerUpdateRequest) GetLicenseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LicenseDate
	}
	return nil
}

func (x *CustomerUpdateRequest) GetContactName() string {
	if x != nil && x.ContactName != nil {
		return *x.ContactName
	}
	return ""
}

func (x *CustomerUpdateRequest) GetContactTitle() string {
	if x != nil && x.ContactTitle != nil {
		return *x.ContactTitle
	}
	return ""
}

func (x *CustomerUpdateRequest) GetContactPhone() string {
	if x != nil && x.ContactPhone != nil {
		return *x.ContactPhone
	}
	return ""
}

func (x *CustomerUpdateRequest) GetContactEmail() string {
	if x != nil && x.ContactEmail != nil {
		return *x.ContactEmail
	}
	return ""
}

func (x *CustomerUpdateRequest) GetContactAddress() *AddressPayload {
	if x != nil {
		return x.ContactAddress
	}
	return nil
}

func (x *CustomerUpdateRequest) GetAccountNumber() string {
	if x != nil && x.AccountNumber != nil {
		return *x.AccountNumber
	}
	return ""
}

func (x *CustomerUpdateRequest) GetBankName() string {
	if x != nil && x.BankName != nil {
		return *x.BankName
	}
	return ""
}

func (x *CustomerUpdateRequest) GetBankBranch() string {
	if x != nil && x.BankBranch != nil {
		return *x.BankBranch
	}
	return ""
}

type CustomerUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CustomerUpdateResponse) Reset() {
	*x = CustomerUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_customer_rpc_customer_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUpdateResponse) ProtoMessage() {}

func (x *CustomerUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_customer_rpc_customer_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUpdateResponse.ProtoReflect.Descriptor instead.
func (*CustomerUpdateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_customer_rpc_customer_update_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerUpdateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CustomerUpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_customer_rpc_customer_update_proto protoreflect.FileDescriptor

var file_rpc_customer_rpc_customer_update_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4,
	0x07, 0x0a, 0x15, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x03, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05, 0x52, 0x0b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x0a, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0c,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0d, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x62, 0x61, 0x6e, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x46, 0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e,
	0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61,
	0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_customer_rpc_customer_update_proto_rawDescOnce sync.Once
	file_rpc_customer_rpc_customer_update_proto_rawDescData = file_rpc_customer_rpc_customer_update_proto_rawDesc
)

func file_rpc_customer_rpc_customer_update_proto_rawDescGZIP() []byte {
	file_rpc_customer_rpc_customer_update_proto_rawDescOnce.Do(func() {
		file_rpc_customer_rpc_customer_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_customer_rpc_customer_update_proto_rawDescData)
	})
	return file_rpc_customer_rpc_customer_update_proto_rawDescData
}

var file_rpc_customer_rpc_customer_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_customer_rpc_customer_update_proto_goTypes = []interface{}{
	(*CustomerUpdateRequest)(nil),  // 0: pb.CustomerUpdateRequest
	(*CustomerUpdateResponse)(nil), // 1: pb.CustomerUpdateResponse
	(*AddressPayload)(nil),         // 2: pb.AddressPayload
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_rpc_customer_rpc_customer_update_proto_depIdxs = []int32{
	2, // 0: pb.CustomerUpdateRequest.address:type_name -> pb.AddressPayload
	3, // 1: pb.CustomerUpdateRequest.birthday:type_name -> google.protobuf.Timestamp
	3, // 2: pb.CustomerUpdateRequest.license_date:type_name -> google.protobuf.Timestamp
	2, // 3: pb.CustomerUpdateRequest.contact_address:type_name -> pb.AddressPayload
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_customer_rpc_customer_update_proto_init() }
func file_rpc_customer_rpc_customer_update_proto_init() {
	if File_rpc_customer_rpc_customer_update_proto != nil {
		return
	}
	file_entities_customer_proto_init()
	file_entities_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_customer_rpc_customer_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUpdateRequest); i {
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
		file_rpc_customer_rpc_customer_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUpdateResponse); i {
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
	file_rpc_customer_rpc_customer_update_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_customer_rpc_customer_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_customer_rpc_customer_update_proto_goTypes,
		DependencyIndexes: file_rpc_customer_rpc_customer_update_proto_depIdxs,
		MessageInfos:      file_rpc_customer_rpc_customer_update_proto_msgTypes,
	}.Build()
	File_rpc_customer_rpc_customer_update_proto = out.File
	file_rpc_customer_rpc_customer_update_proto_rawDesc = nil
	file_rpc_customer_rpc_customer_update_proto_goTypes = nil
	file_rpc_customer_rpc_customer_update_proto_depIdxs = nil
}
