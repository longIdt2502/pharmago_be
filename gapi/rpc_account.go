package gapi

import (
	"database/sql"
	"errors"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) AccountDetail(ctx context.Context, req *pb.AccountDetailRequest) (*pb.AccountDetailResponse, error) {
	account, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account: ", err)
	}

	accountPb := mapper.AccountMapper(account)

	companies, err := server.store.GetCompanies(ctx, db.GetCompaniesParams{
		Owner: sql.NullInt32{
			Int32: account.ID,
			Valid: true,
		},
	})

	var companiesPb []*pb.Company
	for _, value := range companies {
		dataPb := mapper.CompanyMapper(ctx, server.store, value)
		companiesPb = append(companiesPb, dataPb)
	}

	return &pb.AccountDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.AccountDetailResponseDetail{
			Account: accountPb,
			Company: companiesPb,
		},
	}, nil
}

func (server *ServerGRPC) AccountInactive(ctx context.Context, req *pb.AccountInactiveRequest) (*pb.AccountInactiveResponse, error) {
	return nil, nil
}
