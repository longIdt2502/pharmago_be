// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: service_pharmago.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_pharmago_proto protoreflect.FileDescriptor

var file_service_pharmago_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72,
	0x70, 0x63, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x72, 0x70, 0x63, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe7, 0x09, 0x0a, 0x08,
	0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x12, 0x64, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x92, 0x41, 0x17, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x55, 0x73, 0x65, 0x72,
	0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x87,
	0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x92, 0x41, 0x1f, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3f, 0x92, 0x41, 0x1f, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a,
	0x12, 0x80, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x92, 0x41, 0x1a, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x65,
	0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x92, 0x41, 0x1b, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x10, 0x47, 0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x7c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x92,
	0x41, 0x1c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x11, 0x47, 0x65, 0x74,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7c, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x12, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x92, 0x41, 0x1c,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x11, 0x47, 0x65, 0x74, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x61, 0x72, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x57, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x92,
	0x41, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0d, 0x47, 0x65, 0x74,
	0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x77, 0x61, 0x72, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x10, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61,
	0x72, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x83, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x92, 0x41,
	0x1d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x92, 0x41, 0x1d,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x4e, 0x65, 0x77, 0x20, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x85, 0x02, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x61, 0x6e, 0x67, 0x4c, 0x6f, 0x6e, 0x67, 0x32, 0x35,
	0x30, 0x32, 0x2f, 0x70, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x92, 0x41, 0xd8, 0x01, 0x12, 0x6d, 0x0a, 0x13, 0x41, 0x20, 0x42, 0x69, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x51, 0x0a, 0x0b,
	0x50, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x67, 0x6f, 0x5f, 0x42, 0x65, 0x12, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x12, 0x6c, 0x6f, 0x6e,
	0x67, 0x2e, 0x6c, 0x62, 0x67, 0x40, 0x69, 0x64, 0x74, 0x69, 0x6e, 0x63, 0x2e, 0x63, 0x6f, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62,
	0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_pharmago_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),          // 0: pb.LoginRequest
	(*CreateAccountRequest)(nil),  // 1: pb.CreateAccountRequest
	(*VerifyAccountRequest)(nil),  // 2: pb.VerifyAccountRequest
	(*CreateCompanyRequest)(nil),  // 3: pb.CreateCompanyRequest
	(*GetCompaniesRequest)(nil),   // 4: pb.GetCompaniesRequest
	(*ProvincesRequest)(nil),      // 5: pb.ProvincesRequest
	(*DistrictsRequest)(nil),      // 6: pb.DistrictsRequest
	(*WardsRequest)(nil),          // 7: pb.WardsRequest
	(*CreateProductRequest)(nil),  // 8: pb.CreateProductRequest
	(*ListProductRequest)(nil),    // 9: pb.ListProductRequest
	(*LoginResponse)(nil),         // 10: pb.LoginResponse
	(*CreateAccountResponse)(nil), // 11: pb.CreateAccountResponse
	(*VerifyAccountResponse)(nil), // 12: pb.VerifyAccountResponse
	(*CreateCompanyResponse)(nil), // 13: pb.CreateCompanyResponse
	(*GetCompaniesResponse)(nil),  // 14: pb.GetCompaniesResponse
	(*ProvincesResponse)(nil),     // 15: pb.ProvincesResponse
	(*DistrictsResponse)(nil),     // 16: pb.DistrictsResponse
	(*WardsResponse)(nil),         // 17: pb.WardsResponse
	(*CreateProductResponse)(nil), // 18: pb.CreateProductResponse
	(*ListProductResponse)(nil),   // 19: pb.ListProductResponse
}
var file_service_pharmago_proto_depIdxs = []int32{
	0,  // 0: pb.Pharmago.Login:input_type -> pb.LoginRequest
	1,  // 1: pb.Pharmago.CreateAccount:input_type -> pb.CreateAccountRequest
	2,  // 2: pb.Pharmago.VerifyAccount:input_type -> pb.VerifyAccountRequest
	3,  // 3: pb.Pharmago.CreateCompany:input_type -> pb.CreateCompanyRequest
	4,  // 4: pb.Pharmago.ListCompanies:input_type -> pb.GetCompaniesRequest
	5,  // 5: pb.Pharmago.ListProvinces:input_type -> pb.ProvincesRequest
	6,  // 6: pb.Pharmago.ListDistricts:input_type -> pb.DistrictsRequest
	7,  // 7: pb.Pharmago.ListWards:input_type -> pb.WardsRequest
	8,  // 8: pb.Pharmago.CreateProduct:input_type -> pb.CreateProductRequest
	9,  // 9: pb.Pharmago.ListProduct:input_type -> pb.ListProductRequest
	10, // 10: pb.Pharmago.Login:output_type -> pb.LoginResponse
	11, // 11: pb.Pharmago.CreateAccount:output_type -> pb.CreateAccountResponse
	12, // 12: pb.Pharmago.VerifyAccount:output_type -> pb.VerifyAccountResponse
	13, // 13: pb.Pharmago.CreateCompany:output_type -> pb.CreateCompanyResponse
	14, // 14: pb.Pharmago.ListCompanies:output_type -> pb.GetCompaniesResponse
	15, // 15: pb.Pharmago.ListProvinces:output_type -> pb.ProvincesResponse
	16, // 16: pb.Pharmago.ListDistricts:output_type -> pb.DistrictsResponse
	17, // 17: pb.Pharmago.ListWards:output_type -> pb.WardsResponse
	18, // 18: pb.Pharmago.CreateProduct:output_type -> pb.CreateProductResponse
	19, // 19: pb.Pharmago.ListProduct:output_type -> pb.ListProductResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_pharmago_proto_init() }
func file_service_pharmago_proto_init() {
	if File_service_pharmago_proto != nil {
		return
	}
	file_rpc_rpc_login_proto_init()
	file_rpc_rpc_create_account_proto_init()
	file_rpc_rpc_verify_account_proto_init()
	file_rpc_address_rpc_provinces_proto_init()
	file_rpc_address_rpc_districts_proto_init()
	file_rpc_address_rpc_wards_proto_init()
	file_rpc_company_rpc_create_company_proto_init()
	file_rpc_company_rpc_companies_proto_init()
	file_rpc_product_rpc_create_product_proto_init()
	file_rpc_product_rpc_list_product_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_pharmago_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_pharmago_proto_goTypes,
		DependencyIndexes: file_service_pharmago_proto_depIdxs,
	}.Build()
	File_service_pharmago_proto = out.File
	file_service_pharmago_proto_rawDesc = nil
	file_service_pharmago_proto_goTypes = nil
	file_service_pharmago_proto_depIdxs = nil
}
