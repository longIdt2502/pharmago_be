// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/account/rpc_create_employee.proto

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

type CreateEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FullName    string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email       string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	AccountType string                 `protobuf:"bytes,5,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	Role        *int32                 `protobuf:"varint,6,opt,name=role,proto3,oneof" json:"role,omitempty"`
	Gender      *string                `protobuf:"bytes,7,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	Licence     *string                `protobuf:"bytes,8,opt,name=licence,proto3,oneof" json:"licence,omitempty"`
	Dob         *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=dob,proto3,oneof" json:"dob,omitempty"`
	Address     *Address               `protobuf:"bytes,10,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Company     int32                  `protobuf:"varint,11,opt,name=company,proto3" json:"company,omitempty"`
	Active      bool                   `protobuf:"varint,12,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *CreateEmployeeRequest) Reset() {
	*x = CreateEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_account_rpc_create_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeRequest) ProtoMessage() {}

func (x *CreateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_account_rpc_create_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*CreateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_account_rpc_create_employee_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEmployeeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateEmployeeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateEmployeeRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateEmployeeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *CreateEmployeeRequest) GetRole() int32 {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return 0
}

func (x *CreateEmployeeRequest) GetGender() string {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return ""
}

func (x *CreateEmployeeRequest) GetLicence() string {
	if x != nil && x.Licence != nil {
		return *x.Licence
	}
	return ""
}

func (x *CreateEmployeeRequest) GetDob() *timestamppb.Timestamp {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *CreateEmployeeRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CreateEmployeeRequest) GetCompany() int32 {
	if x != nil {
		return x.Company
	}
	return 0
}

func (x *CreateEmployeeRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type CreateEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details int32  `protobuf:"varint,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *CreateEmployeeResponse) Reset() {
	*x = CreateEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_account_rpc_create_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeResponse) ProtoMessage() {}

func (x *CreateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_account_rpc_create_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*CreateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_account_rpc_create_employee_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmployeeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateEmployeeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateEmployeeResponse) GetDetails() int32 {
	if x != nil {
		return x.Details
	}
	return 0
}

var File_rpc_account_rpc_create_employee_proto protoreflect.FileDescriptor

var file_rpc_account_rpc_create_employee_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6c, 0x69,
	0x63, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x03, 0x64, 0x6f, 0x62,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x03, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x04, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64,
	0x6f, 0x62, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x60,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48,
	0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35, 0x30, 0x32, 0x2f, 0x70, 0x68, 0x61,
	0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_account_rpc_create_employee_proto_rawDescOnce sync.Once
	file_rpc_account_rpc_create_employee_proto_rawDescData = file_rpc_account_rpc_create_employee_proto_rawDesc
)

func file_rpc_account_rpc_create_employee_proto_rawDescGZIP() []byte {
	file_rpc_account_rpc_create_employee_proto_rawDescOnce.Do(func() {
		file_rpc_account_rpc_create_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_account_rpc_create_employee_proto_rawDescData)
	})
	return file_rpc_account_rpc_create_employee_proto_rawDescData
}

var file_rpc_account_rpc_create_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_account_rpc_create_employee_proto_goTypes = []interface{}{
	(*CreateEmployeeRequest)(nil),  // 0: pb.CreateEmployeeRequest
	(*CreateEmployeeResponse)(nil), // 1: pb.CreateEmployeeResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*Address)(nil),                // 3: pb.Address
}
var file_rpc_account_rpc_create_employee_proto_depIdxs = []int32{
	2, // 0: pb.CreateEmployeeRequest.dob:type_name -> google.protobuf.Timestamp
	3, // 1: pb.CreateEmployeeRequest.address:type_name -> pb.Address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_account_rpc_create_employee_proto_init() }
func file_rpc_account_rpc_create_employee_proto_init() {
	if File_rpc_account_rpc_create_employee_proto != nil {
		return
	}
	file_entities_account_proto_init()
	file_entities_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_account_rpc_create_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmployeeRequest); i {
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
		file_rpc_account_rpc_create_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmployeeResponse); i {
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
	file_rpc_account_rpc_create_employee_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_account_rpc_create_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_account_rpc_create_employee_proto_goTypes,
		DependencyIndexes: file_rpc_account_rpc_create_employee_proto_depIdxs,
		MessageInfos:      file_rpc_account_rpc_create_employee_proto_msgTypes,
	}.Build()
	File_rpc_account_rpc_create_employee_proto = out.File
	file_rpc_account_rpc_create_employee_proto_rawDesc = nil
	file_rpc_account_rpc_create_employee_proto_goTypes = nil
	file_rpc_account_rpc_create_employee_proto_depIdxs = nil
}
