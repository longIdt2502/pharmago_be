// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: service_pharmago.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PharmagoClient is the client API for Pharmago service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PharmagoClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	VerifyAccount(ctx context.Context, in *VerifyAccountRequest, opts ...grpc.CallOption) (*VerifyAccountResponse, error)
}

type pharmagoClient struct {
	cc grpc.ClientConnInterface
}

func NewPharmagoClient(cc grpc.ClientConnInterface) PharmagoClient {
	return &pharmagoClient{cc}
}

func (c *pharmagoClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) VerifyAccount(ctx context.Context, in *VerifyAccountRequest, opts ...grpc.CallOption) (*VerifyAccountResponse, error) {
	out := new(VerifyAccountResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/VerifyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PharmagoServer is the server API for Pharmago service.
// All implementations must embed UnimplementedPharmagoServer
// for forward compatibility
type PharmagoServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	VerifyAccount(context.Context, *VerifyAccountRequest) (*VerifyAccountResponse, error)
	mustEmbedUnimplementedPharmagoServer()
}

// UnimplementedPharmagoServer must be embedded to have forward compatible implementations.
type UnimplementedPharmagoServer struct {
}

func (UnimplementedPharmagoServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPharmagoServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedPharmagoServer) VerifyAccount(context.Context, *VerifyAccountRequest) (*VerifyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
func (UnimplementedPharmagoServer) mustEmbedUnimplementedPharmagoServer() {}

// UnsafePharmagoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PharmagoServer will
// result in compilation errors.
type UnsafePharmagoServer interface {
	mustEmbedUnimplementedPharmagoServer()
}

func RegisterPharmagoServer(s grpc.ServiceRegistrar, srv PharmagoServer) {
	s.RegisterService(&Pharmago_ServiceDesc, srv)
}

func _Pharmago_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_VerifyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).VerifyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/VerifyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).VerifyAccount(ctx, req.(*VerifyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pharmago_ServiceDesc is the grpc.ServiceDesc for Pharmago service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pharmago_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Pharmago",
	HandlerType: (*PharmagoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Pharmago_Login_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Pharmago_CreateAccount_Handler,
		},
		{
			MethodName: "VerifyAccount",
			Handler:    _Pharmago_VerifyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_pharmago.proto",
}
