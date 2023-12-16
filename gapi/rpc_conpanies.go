package gapi

import (
	"context"
	"database/sql"
	"errors"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
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
		Owner: sql.NullInt32{
			Int32: int32(accountRequest.ID),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list company")
	}
	print(accountRequest.ID)
	print(req.GetPage())
	print(req.GetLimit())
	print(req.GetSearch())
	print(companies)

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
