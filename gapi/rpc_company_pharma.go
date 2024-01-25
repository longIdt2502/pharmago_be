package gapi

import (
	"context"
	"database/sql"
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) CompanyPharmaList(ctx context.Context, req *pb.CompanyPharmaListRequest) (*pb.CompanyPharmaListResponse, error) {
	companyPharma, err := server.store.GetListCompanyPharma(ctx, db.GetListCompanyPharmaParams{
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
		Type: req.GetType(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list preparation type: %w", err)
	}

	var companyPharmaPb []*pb.CompanyPharma
	for _, value := range companyPharma {
		companyPharmaPb = append(companyPharmaPb, mapper.CompanyPharmaMapper(value))
	}

	return &pb.CompanyPharmaListResponse{
		Code:    200,
		Message: "success",
		Details: companyPharmaPb,
	}, nil
}

func (server *ServerGRPC) CompanyPharmaCreate(ctx context.Context, req *pb.CompanyPharmaCreateRequest) (*pb.CompanyPharmaCreateResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	companyPharma, err := server.store.CreateCompanyPharma(ctx, db.CreateCompanyPharmaParams{
		Name: req.GetName(),
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  true,
		},
		Country: sql.NullString{
			String: req.GetCountry(),
			Valid:  true,
		},
		Address: sql.NullString{
			String: req.GetAddress(),
			Valid:  true,
		},
		CompanyPharmaType: req.GetType().String(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record company pharma: %v", err))
	}

	return &pb.CompanyPharmaCreateResponse{
		Code:    200,
		Message: "success",
		Details: companyPharma.ID,
	}, nil
}

func (server *ServerGRPC) CompanyPharmaDetail(ctx context.Context, req *pb.CompanyPharmaDetailRequest) (*pb.CompanyPharmaDetailResponse, error) {
	companyPharma, err := server.store.GetCompanyPharmaDetail(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list preparation type: %w", err)
	}

	companyPharmaPb := mapper.CompanyPharmaMapper(companyPharma)

	return &pb.CompanyPharmaDetailResponse{
		Code:    200,
		Message: "success",
		Details: companyPharmaPb,
	}, nil
}

func (server *ServerGRPC) CompanyPharmaUpdate(ctx context.Context, req *pb.CompanyPharmaUpdateRequest) (*pb.CompanyPharmaUpdateResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	companyPharma, err := server.store.UpdateCompanyPharma(ctx, db.UpdateCompanyPharmaParams{
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  true,
		},
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  true,
		},
		Country: sql.NullString{
			String: req.GetCountry(),
			Valid:  true,
		},
		Address: sql.NullString{
			String: req.GetAddress(),
			Valid:  true,
		},
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update company pharma: %v", err))
	}

	return &pb.CompanyPharmaUpdateResponse{
		Code:    200,
		Message: "success",
		Details: companyPharma.ID,
	}, nil
}

func (server *ServerGRPC) CompanyPharmaDelete(ctx context.Context, req *pb.CompanyPharmaDeleteRequest) (*pb.CompanyPharmaDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteCompanyPharma(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete company pharma: %v", err))
	}

	return &pb.CompanyPharmaDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
