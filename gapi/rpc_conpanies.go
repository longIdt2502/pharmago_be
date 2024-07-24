package gapi

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/longIdt2502/pharmago_be/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListCompanies(ctx context.Context, req *pb.GetCompaniesRequest) (*pb.GetCompaniesResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticated")
	}

	accountRequest, err := server.store.GetAccountByUseName(ctx, tokenPayload.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user doesn't exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user")
	}

	companies, err := server.store.GetCompanies(ctx, db.GetCompaniesParams{
		Owner:  sql.NullInt32{Int32: int32(accountRequest.ID), Valid: true},
		Search: sql.NullString{String: req.GetSearch(), Valid: req.Search != nil},
		Parent: sql.NullInt32{Int32: req.GetParent(), Valid: req.Parent != nil},
		Page:   sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:  sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list company")
	}

	var companyPb []*pb.Company
	for _, value := range companies {
		data := mapper.CompanyMapper(ctx, server.store, value)
		companyPb = append(companyPb, data)
	}

	rsp := &pb.GetCompaniesResponse{
		Code:    200,
		Message: "success",
		Details: companyPb,
	}

	return rsp, nil
}

func (server *ServerGRPC) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyDataRequest) (*pb.UpdateCompanyResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticated")
	}

	company, err := server.store.UpdateCompany(ctx, db.UpdateCompanyParams{
		Name:        sql.NullString{String: req.GetName(), Valid: req.Name != nil},
		Type:        sql.NullString{String: req.GetType(), Valid: req.Type != nil},
		Manager:     sql.NullInt32{Int32: req.GetManager(), Valid: req.Manager != nil},
		IsActive:    sql.NullBool{Bool: req.GetIsActive(), Valid: req.IsActive != nil},
		TimeOpen:    sql.NullTime{Time: time.Unix(req.GetTimeOpen().AsTime().Unix(), 0), Valid: req.TimeOpen != nil},
		TimeClose:   sql.NullTime{Time: time.Unix(req.GetTimeClose().AsTime().Unix(), 0), Valid: req.TimeClose != nil},
		UserUpdated: sql.NullInt32{Int32: tokenPayload.UserID, Valid: true},
		ID:          req.GetId(),
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.UpdateCompanyResponse{
			Code:    int32(errApp.StatusCode),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	return &pb.UpdateCompanyResponse{
		Code:    200,
		Message: "success",
		Details: company.ID,
	}, nil
}

func (server *ServerGRPC) DetailCompany(ctx context.Context, req *pb.DetailCompanyDataRequest) (*pb.DetailCompanyResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticated")
	}

	company, err := server.store.DetailCompany(ctx, req.Id)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.DetailCompanyResponse{
			Code:    int32(errApp.StatusCode),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	companyPb := mapper.CompanyDetailMapper(ctx, server.store, company)
	return &pb.DetailCompanyResponse{
		Code:    200,
		Message: "success",
		Details: companyPb,
	}, nil
}

func (server *ServerGRPC) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	violations := validateCreateCompany(req)
	if violations != nil {
		return nil, config.InvalidArgumentError(violations)
	}

	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticated")
	}

	accountRequest, err := server.store.GetAccountByUseName(ctx, tokenPayload.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user doesn't exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user")
	}

	code := utils.RandomString(9)

	address, err := server.store.CreateAddress(ctx, db.CreateAddressParams{
		Lat: float64(req.Address.Lat),
		Lng: float64(req.Address.Lng),
		Province: sql.NullString{
			String: req.Address.GetProvince(),
			Valid:  true,
		},
		District: sql.NullString{
			String: req.Address.GetDistrict(),
			Valid:  true,
		},
		Ward: sql.NullString{
			String: req.Address.GetWard(),
			Valid:  req.Address.Ward != nil,
		},
		Title:       req.Address.GetTitle(),
		UserCreated: accountRequest.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create address: %e", err)
	}

	company, err := server.store.CreateCompany(ctx, db.CreateCompanyParams{
		Name:        req.Company.GetName(),
		Code:        code,
		Type:        req.Company.GetType(),
		TaxCode:     sql.NullString{String: req.Company.GetTaxCode(), Valid: req.Company.TaxCode != nil},
		Phone:       sql.NullString{String: req.Company.GetPhone(), Valid: req.Company.Phone != nil},
		Description: sql.NullString{String: req.Company.GetDescription(), Valid: req.Company.Description != nil},
		Address:     sql.NullInt32{Int32: address.ID, Valid: true},
		Owner:       accountRequest.ID,
		TimeOpen:    sql.NullTime{Time: time.Unix(req.Company.GetTimeOpen().AsTime().Unix(), 0), Valid: req.Company.TimeOpen.IsValid()},
		TimeClose:   sql.NullTime{Time: time.Unix(req.Company.GetTimeClose().AsTime().Unix(), 0), Valid: req.Company.TimeClose.IsValid()},
		UserCreated: sql.NullInt32{Int32: tokenPayload.UserID, Valid: true},
		UserUpdated: sql.NullInt32{Valid: false},
		Parent:      sql.NullInt32{Int32: req.Company.GetCompanyParent(), Valid: req.Company.CompanyParent != nil},
		Manager:     sql.NullInt32{Int32: req.Company.GetManager(), Valid: req.Company.Manager != nil},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create company: %e", err)
	}

	companyData := mapper.CompanyMapper(ctx, server.store, company)

	rsp := &pb.CreateCompanyResponse{
		Code:    200,
		Message: "success",
		Details: companyData,
	}

	return rsp, nil
}

func (server *ServerGRPC) AssignEmployee(ctx context.Context, req *pb.AssignCompanyReq) (*pb.AssignCompanyRes, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to authenticated")
	}

	for _, item := range req.GetAssign() {
		_, err := server.store.AssignEmployee(ctx, db.AssignEmployeeParams{
			Company: sql.NullInt32{Int32: req.GetCompany(), Valid: true},
			Account: item,
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.AssignCompanyRes{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}, nil
		}
	}

	for _, item := range req.GetRemove() {
		_, err := server.store.AssignEmployee(ctx, db.AssignEmployeeParams{
			Company: sql.NullInt32{},
			Account: item,
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.AssignCompanyRes{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}, nil
		}
	}

	return &pb.AssignCompanyRes{
		Code:    200,
		Message: "success",
	}, nil
}

func validateCreateCompany(req *pb.CreateCompanyRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateFullName(req.Company.Name); err != nil {
		violations = append(violations, config.FieldViolation("name", err))
	}
	return violations
}
