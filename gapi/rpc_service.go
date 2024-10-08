package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ServiceList(ctx context.Context, req *pb.ServiceListRequest) (*pb.ServiceListResponse, error) {

	serviceDb, err := server.store.GetListService(ctx, db.GetListServiceParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list service: %e", err)
	}

	var servicePb []*pb.Service
	for _, item := range serviceDb {
		servicePb = append(servicePb, mapper.ServiceGetListServiceRowMapper(item))
	}

	return &pb.ServiceListResponse{
		Code:    200,
		Message: "success",
		Details: servicePb,
	}, nil
}

func (server *ServerGRPC) ServicesByCustomer(ctx context.Context, req *pb.ServicesByCustomerRequest) (*pb.ServicesByCustomerResponse, error) {

	serviceDb, err := server.store.GetServicesByCustomer(ctx, db.GetServicesByCustomerParams{
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
		Customer: req.GetCustomer(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list service: %e", err)
	}

	var servicePb []*pb.Service
	for _, item := range serviceDb {
		servicePb = append(servicePb, mapper.ServiceByCustomerMapper(item))
	}

	return &pb.ServicesByCustomerResponse{
		Code:    200,
		Message: "success",
		Details: servicePb,
	}, nil
}

func (server *ServerGRPC) ServiceCreate(ctx context.Context, req *pb.ServiceCreateRequest) (*pb.ServiceCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	result, err := server.store.CreateServiceTx(ctx, db.CreateServiceTxParams{
		ServiceCreateRequest: req,
		TokenPayload:         tokenPayload,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create service: %e", err))
	}

	return &pb.ServiceCreateResponse{
		Code:    200,
		Message: "success",
		Details: result.Id,
	}, nil
}

func (server *ServerGRPC) ServiceDetail(ctx context.Context, req *pb.ServiceDetailRequest) (*pb.ServiceDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	service, err := server.store.DetailService(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "service not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get service: %e", err)
	}

	serviceVariants, err := server.store.ListServiceVariant(ctx, service.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list variants: %e", err)
	}

	var variantsPb []*pb.Variant
	for _, item := range serviceVariants {
		data := mapper.VariantMapper(ctx, server.store, db.GetVariantsRow{
			ID:   item.ID_2.Int32,
			Name: item.Name.String,
			Code: item.Code.String,
			Barcode: sql.NullString{
				String: item.Barcode.String,
				Valid:  item.Barcode.Valid,
			},
		})
		variantsPb = append(variantsPb, data)
	}

	var accountPb *pb.Account
	if service.Staff.Valid {
		accountDb, err := server.store.GetAccount(ctx, service.Staff.Int32)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "faield to get account: %e", err)
		}
		accountPb = mapper.AccountRowMapper(accountDb)
	}

	servicePb := mapper.ServiceDetailRowMapper(service)
	servicePb.Variants = variantsPb
	servicePb.Staff = accountPb

	return &pb.ServiceDetailResponse{
		Code:    200,
		Message: "success",
		Details: servicePb,
	}, nil
}

func (server *ServerGRPC) ServiceUpdate(ctx context.Context, req *pb.ServiceUpdateRequest) (*pb.ServiceUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	serviceDb, err := server.store.UpdateService(ctx, db.UpdateServiceParams{
		Image:          sql.NullInt32{},
		Title:          sql.NullString{String: req.GetTitle(), Valid: req.Title != nil},
		Entity:         sql.NullString{String: req.GetEntity(), Valid: req.Entity != nil},
		Staff:          sql.NullInt32{Int32: req.GetStaff(), Valid: req.Staff != nil},
		Frequency:      sql.NullString{String: req.GetFrequency(), Valid: req.Frequency != nil},
		Unit:           sql.NullString{String: req.GetUnit(), Valid: req.Unit != nil},
		Price:          sql.NullFloat64{Float64: req.GetPrice(), Valid: req.Price != nil},
		Description:    sql.NullString{String: req.GetDescription(), Valid: req.Description != nil},
		Brand:          sql.NullString{},
		ActionTime:     sql.NullString{String: req.GetActionTime(), Valid: req.ActionTime != nil},
		ChiDinh:        sql.NullString{String: req.GetChiDinh(), Valid: req.ChiDinh != nil},
		ChongChiDinh:   sql.NullString{String: req.GetChongChiDinh(), Valid: req.ChongChiDinh != nil},
		CongDung:       sql.NullString{String: req.GetCongDung(), Valid: req.CongDung != nil},
		Caution:        sql.NullString{String: req.GetCaution(), Valid: req.Caution != nil},
		HinhThuc:       sql.NullString{String: req.GetHinhThuc(), Valid: req.HinhThuc != nil},
		TacDungPhu:     sql.NullString{String: req.GetTacDungPhu(), Valid: req.TacDungPhu != nil},
		NumberRegister: sql.NullString{String: req.GetNumberRegister(), Valid: req.NumberRegister != nil},
		NumberDecision: sql.NullString{String: req.GetNumberDecision(), Valid: req.NumberDecision != nil},
		CongTyDk:       sql.NullString{String: req.GetCongTyDk(), Valid: req.CongTyDk != nil},
		Message:        sql.NullString{String: req.GetMessage(), Valid: req.Message != nil},
		UserUpdated:    sql.NullInt32{Int32: tokenPayload.UserID, Valid: true},
		ID:             req.GetId(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "service not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to update service: %e", err)
	}

	serviceVariantDb, err := server.store.ListServiceVariant(ctx, serviceDb.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get service variant: %e", err)
	}

	for _, item := range serviceVariantDb {
		if !(slices.Contains(req.GetVariants(), item.Variant.Int32)) {
			server.store.DeleteServiceVariant(ctx, item.ID)
		}
	}

	for _, item := range req.GetVariants() {
		server.store.CreateServiceVariant(ctx, db.CreateServiceVariantParams{
			Service: sql.NullInt32{
				Int32: serviceDb.ID,
				Valid: true,
			},
			Variant: sql.NullInt32{
				Int32: item,
				Valid: true,
			},
		})
	}

	return &pb.ServiceUpdateResponse{
		Code:    200,
		Message: "success",
		Details: serviceDb.ID,
	}, nil
}

func (server *ServerGRPC) ServiceDelete(ctx context.Context, req *pb.ServiceDeleteRequest) (*pb.ServiceDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteService(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "service not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete service: %e", err)
	}

	return &pb.ServiceDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
