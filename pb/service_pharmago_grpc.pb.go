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
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
	ListCompanies(ctx context.Context, in *GetCompaniesRequest, opts ...grpc.CallOption) (*GetCompaniesResponse, error)
	ListProvinces(ctx context.Context, in *ProvincesRequest, opts ...grpc.CallOption) (*ProvincesResponse, error)
	ListDistricts(ctx context.Context, in *DistrictsRequest, opts ...grpc.CallOption) (*DistrictsResponse, error)
	ListWards(ctx context.Context, in *WardsRequest, opts ...grpc.CallOption) (*WardsResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	GetPriceList(ctx context.Context, in *PriceListRequest, opts ...grpc.CallOption) (*PriceListResponse, error)
	DetailPriceList(ctx context.Context, in *DetailPriceListRequest, opts ...grpc.CallOption) (*DetailPriceListResponse, error)
	UpdatePriceList(ctx context.Context, in *UpdatePriceListRequest, opts ...grpc.CallOption) (*UpdatePriceListResponse, error)
	// ================== PRODUCT MASTER DATA ===================
	ClassifyList(ctx context.Context, in *ClassifyListRequest, opts ...grpc.CallOption) (*ClassifyListResponse, error)
	ProductionStandardList(ctx context.Context, in *ProductionStandardListRequest, opts ...grpc.CallOption) (*ProductionStandardListResponse, error)
	PreparationTypeList(ctx context.Context, in *PreparationTypeListRequest, opts ...grpc.CallOption) (*PreparationTypeListResponse, error)
	CompanyPharmaList(ctx context.Context, in *CompanyPharmaListRequest, opts ...grpc.CallOption) (*CompanyPharmaListResponse, error)
	// ================== BRAND ===================
	BrandList(ctx context.Context, in *BrandListRequest, opts ...grpc.CallOption) (*BrandListResponse, error)
	// ================== CATEGORY ===================
	CategoryList(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error)
	// ================== PRODUCT_TYPE ===================
	ProductTypeList(ctx context.Context, in *ProductTypeListRequest, opts ...grpc.CallOption) (*ProductTypeListResponse, error)
	// ================== IMPORT ===================
	ImportCompany(ctx context.Context, in *ImportCompanyRequest, opts ...grpc.CallOption) (*ImportCompanyResponse, error)
	ImportProduct(ctx context.Context, in *ImportProductRequest, opts ...grpc.CallOption) (*ImportProductResponse, error)
	ImportProductMasterData(ctx context.Context, in *ImportProductMasterDataRequest, opts ...grpc.CallOption) (*ImportProductMasterDataResponse, error)
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

func (c *pharmagoClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ListCompanies(ctx context.Context, in *GetCompaniesRequest, opts ...grpc.CallOption) (*GetCompaniesResponse, error) {
	out := new(GetCompaniesResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ListCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ListProvinces(ctx context.Context, in *ProvincesRequest, opts ...grpc.CallOption) (*ProvincesResponse, error) {
	out := new(ProvincesResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ListProvinces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ListDistricts(ctx context.Context, in *DistrictsRequest, opts ...grpc.CallOption) (*DistrictsResponse, error) {
	out := new(DistrictsResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ListDistricts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ListWards(ctx context.Context, in *WardsRequest, opts ...grpc.CallOption) (*WardsResponse, error) {
	out := new(WardsResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ListWards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ListProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) GetPriceList(ctx context.Context, in *PriceListRequest, opts ...grpc.CallOption) (*PriceListResponse, error) {
	out := new(PriceListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/GetPriceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) DetailPriceList(ctx context.Context, in *DetailPriceListRequest, opts ...grpc.CallOption) (*DetailPriceListResponse, error) {
	out := new(DetailPriceListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/DetailPriceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) UpdatePriceList(ctx context.Context, in *UpdatePriceListRequest, opts ...grpc.CallOption) (*UpdatePriceListResponse, error) {
	out := new(UpdatePriceListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/UpdatePriceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ClassifyList(ctx context.Context, in *ClassifyListRequest, opts ...grpc.CallOption) (*ClassifyListResponse, error) {
	out := new(ClassifyListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ClassifyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ProductionStandardList(ctx context.Context, in *ProductionStandardListRequest, opts ...grpc.CallOption) (*ProductionStandardListResponse, error) {
	out := new(ProductionStandardListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ProductionStandardList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) PreparationTypeList(ctx context.Context, in *PreparationTypeListRequest, opts ...grpc.CallOption) (*PreparationTypeListResponse, error) {
	out := new(PreparationTypeListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/PreparationTypeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) CompanyPharmaList(ctx context.Context, in *CompanyPharmaListRequest, opts ...grpc.CallOption) (*CompanyPharmaListResponse, error) {
	out := new(CompanyPharmaListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/CompanyPharmaList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) BrandList(ctx context.Context, in *BrandListRequest, opts ...grpc.CallOption) (*BrandListResponse, error) {
	out := new(BrandListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/BrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) CategoryList(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/CategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ProductTypeList(ctx context.Context, in *ProductTypeListRequest, opts ...grpc.CallOption) (*ProductTypeListResponse, error) {
	out := new(ProductTypeListResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ProductTypeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ImportCompany(ctx context.Context, in *ImportCompanyRequest, opts ...grpc.CallOption) (*ImportCompanyResponse, error) {
	out := new(ImportCompanyResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ImportCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ImportProduct(ctx context.Context, in *ImportProductRequest, opts ...grpc.CallOption) (*ImportProductResponse, error) {
	out := new(ImportProductResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ImportProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pharmagoClient) ImportProductMasterData(ctx context.Context, in *ImportProductMasterDataRequest, opts ...grpc.CallOption) (*ImportProductMasterDataResponse, error) {
	out := new(ImportProductMasterDataResponse)
	err := c.cc.Invoke(ctx, "/pb.Pharmago/ImportProductMasterData", in, out, opts...)
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
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
	ListCompanies(context.Context, *GetCompaniesRequest) (*GetCompaniesResponse, error)
	ListProvinces(context.Context, *ProvincesRequest) (*ProvincesResponse, error)
	ListDistricts(context.Context, *DistrictsRequest) (*DistrictsResponse, error)
	ListWards(context.Context, *WardsRequest) (*WardsResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
	GetPriceList(context.Context, *PriceListRequest) (*PriceListResponse, error)
	DetailPriceList(context.Context, *DetailPriceListRequest) (*DetailPriceListResponse, error)
	UpdatePriceList(context.Context, *UpdatePriceListRequest) (*UpdatePriceListResponse, error)
	// ================== PRODUCT MASTER DATA ===================
	ClassifyList(context.Context, *ClassifyListRequest) (*ClassifyListResponse, error)
	ProductionStandardList(context.Context, *ProductionStandardListRequest) (*ProductionStandardListResponse, error)
	PreparationTypeList(context.Context, *PreparationTypeListRequest) (*PreparationTypeListResponse, error)
	CompanyPharmaList(context.Context, *CompanyPharmaListRequest) (*CompanyPharmaListResponse, error)
	// ================== BRAND ===================
	BrandList(context.Context, *BrandListRequest) (*BrandListResponse, error)
	// ================== CATEGORY ===================
	CategoryList(context.Context, *CategoryListRequest) (*CategoryListResponse, error)
	// ================== PRODUCT_TYPE ===================
	ProductTypeList(context.Context, *ProductTypeListRequest) (*ProductTypeListResponse, error)
	// ================== IMPORT ===================
	ImportCompany(context.Context, *ImportCompanyRequest) (*ImportCompanyResponse, error)
	ImportProduct(context.Context, *ImportProductRequest) (*ImportProductResponse, error)
	ImportProductMasterData(context.Context, *ImportProductMasterDataRequest) (*ImportProductMasterDataResponse, error)
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
func (UnimplementedPharmagoServer) CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedPharmagoServer) ListCompanies(context.Context, *GetCompaniesRequest) (*GetCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanies not implemented")
}
func (UnimplementedPharmagoServer) ListProvinces(context.Context, *ProvincesRequest) (*ProvincesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProvinces not implemented")
}
func (UnimplementedPharmagoServer) ListDistricts(context.Context, *DistrictsRequest) (*DistrictsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDistricts not implemented")
}
func (UnimplementedPharmagoServer) ListWards(context.Context, *WardsRequest) (*WardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWards not implemented")
}
func (UnimplementedPharmagoServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedPharmagoServer) ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedPharmagoServer) GetPriceList(context.Context, *PriceListRequest) (*PriceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPriceList not implemented")
}
func (UnimplementedPharmagoServer) DetailPriceList(context.Context, *DetailPriceListRequest) (*DetailPriceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailPriceList not implemented")
}
func (UnimplementedPharmagoServer) UpdatePriceList(context.Context, *UpdatePriceListRequest) (*UpdatePriceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePriceList not implemented")
}
func (UnimplementedPharmagoServer) ClassifyList(context.Context, *ClassifyListRequest) (*ClassifyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassifyList not implemented")
}
func (UnimplementedPharmagoServer) ProductionStandardList(context.Context, *ProductionStandardListRequest) (*ProductionStandardListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductionStandardList not implemented")
}
func (UnimplementedPharmagoServer) PreparationTypeList(context.Context, *PreparationTypeListRequest) (*PreparationTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreparationTypeList not implemented")
}
func (UnimplementedPharmagoServer) CompanyPharmaList(context.Context, *CompanyPharmaListRequest) (*CompanyPharmaListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompanyPharmaList not implemented")
}
func (UnimplementedPharmagoServer) BrandList(context.Context, *BrandListRequest) (*BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrandList not implemented")
}
func (UnimplementedPharmagoServer) CategoryList(context.Context, *CategoryListRequest) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryList not implemented")
}
func (UnimplementedPharmagoServer) ProductTypeList(context.Context, *ProductTypeListRequest) (*ProductTypeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductTypeList not implemented")
}
func (UnimplementedPharmagoServer) ImportCompany(context.Context, *ImportCompanyRequest) (*ImportCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCompany not implemented")
}
func (UnimplementedPharmagoServer) ImportProduct(context.Context, *ImportProductRequest) (*ImportProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProduct not implemented")
}
func (UnimplementedPharmagoServer) ImportProductMasterData(context.Context, *ImportProductMasterDataRequest) (*ImportProductMasterDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProductMasterData not implemented")
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

func _Pharmago_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ListCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ListCompanies(ctx, req.(*GetCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ListProvinces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvincesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ListProvinces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ListProvinces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ListProvinces(ctx, req.(*ProvincesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ListDistricts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistrictsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ListDistricts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ListDistricts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ListDistricts(ctx, req.(*DistrictsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ListWards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ListWards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ListWards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ListWards(ctx, req.(*WardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ListProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_GetPriceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).GetPriceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/GetPriceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).GetPriceList(ctx, req.(*PriceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_DetailPriceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailPriceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).DetailPriceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/DetailPriceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).DetailPriceList(ctx, req.(*DetailPriceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_UpdatePriceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePriceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).UpdatePriceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/UpdatePriceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).UpdatePriceList(ctx, req.(*UpdatePriceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ClassifyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassifyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ClassifyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ClassifyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ClassifyList(ctx, req.(*ClassifyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ProductionStandardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductionStandardListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ProductionStandardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ProductionStandardList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ProductionStandardList(ctx, req.(*ProductionStandardListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_PreparationTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreparationTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).PreparationTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/PreparationTypeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).PreparationTypeList(ctx, req.(*PreparationTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_CompanyPharmaList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyPharmaListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).CompanyPharmaList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/CompanyPharmaList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).CompanyPharmaList(ctx, req.(*CompanyPharmaListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_BrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).BrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/BrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).BrandList(ctx, req.(*BrandListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_CategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).CategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/CategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).CategoryList(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ProductTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductTypeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ProductTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ProductTypeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ProductTypeList(ctx, req.(*ProductTypeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ImportCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ImportCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ImportCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ImportCompany(ctx, req.(*ImportCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ImportProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ImportProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ImportProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ImportProduct(ctx, req.(*ImportProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pharmago_ImportProductMasterData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportProductMasterDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PharmagoServer).ImportProductMasterData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Pharmago/ImportProductMasterData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PharmagoServer).ImportProductMasterData(ctx, req.(*ImportProductMasterDataRequest))
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
		{
			MethodName: "CreateCompany",
			Handler:    _Pharmago_CreateCompany_Handler,
		},
		{
			MethodName: "ListCompanies",
			Handler:    _Pharmago_ListCompanies_Handler,
		},
		{
			MethodName: "ListProvinces",
			Handler:    _Pharmago_ListProvinces_Handler,
		},
		{
			MethodName: "ListDistricts",
			Handler:    _Pharmago_ListDistricts_Handler,
		},
		{
			MethodName: "ListWards",
			Handler:    _Pharmago_ListWards_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _Pharmago_CreateProduct_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _Pharmago_ListProduct_Handler,
		},
		{
			MethodName: "GetPriceList",
			Handler:    _Pharmago_GetPriceList_Handler,
		},
		{
			MethodName: "DetailPriceList",
			Handler:    _Pharmago_DetailPriceList_Handler,
		},
		{
			MethodName: "UpdatePriceList",
			Handler:    _Pharmago_UpdatePriceList_Handler,
		},
		{
			MethodName: "ClassifyList",
			Handler:    _Pharmago_ClassifyList_Handler,
		},
		{
			MethodName: "ProductionStandardList",
			Handler:    _Pharmago_ProductionStandardList_Handler,
		},
		{
			MethodName: "PreparationTypeList",
			Handler:    _Pharmago_PreparationTypeList_Handler,
		},
		{
			MethodName: "CompanyPharmaList",
			Handler:    _Pharmago_CompanyPharmaList_Handler,
		},
		{
			MethodName: "BrandList",
			Handler:    _Pharmago_BrandList_Handler,
		},
		{
			MethodName: "CategoryList",
			Handler:    _Pharmago_CategoryList_Handler,
		},
		{
			MethodName: "ProductTypeList",
			Handler:    _Pharmago_ProductTypeList_Handler,
		},
		{
			MethodName: "ImportCompany",
			Handler:    _Pharmago_ImportCompany_Handler,
		},
		{
			MethodName: "ImportProduct",
			Handler:    _Pharmago_ImportProduct_Handler,
		},
		{
			MethodName: "ImportProductMasterData",
			Handler:    _Pharmago_ImportProductMasterData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_pharmago.proto",
}
