package gapi

import (
	"context"
	"database/sql"
	"errors"

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
		Name: req.Company.GetName(),
		Code: code,
		Type: req.Company.GetType(),
		TaxCode: sql.NullString{
			String: req.Company.GetTaxCode(),
			Valid:  req.Company.TaxCode != nil,
		},
		Phone: sql.NullString{
			String: req.Company.GetPhone(),
			Valid:  req.Company.Phone != nil,
		},
		Description: sql.NullString{
			String: req.Company.GetDescription(),
			Valid:  req.Company.Description != nil,
		},
		Address: sql.NullInt32{
			Int32: address.ID,
			Valid: true,
		},
		Owner: accountRequest.ID,
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

func validateCreateCompany(req *pb.CreateCompanyRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateFullName(req.Company.Name); err != nil {
		violations = append(violations, config.FieldViolation("name", err))
	}
	return violations
}
