package gapi

import (
	"database/sql"
	"errors"
	"fmt"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) AccountDetail(ctx context.Context, _ *pb.AccountDetailRequest) (*pb.AccountDetailResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	account, err := server.store.GetAccount(ctx, tokenPayload.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account: %e", err)
	}

	accountPb := mapper.AccountMapper(account)

	companies, _ := server.store.GetCompanies(ctx, db.GetCompaniesParams{
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
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	var accountId int32

	if req.Id != nil {
		account, err := server.store.GetAccount(ctx, req.GetId())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "account not exists")
			}
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get account: %v", err))
		}
		accountId = account.ID

		if account.Type == 3 && req.GetId() != tokenPayload.UserID {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
	} else {
		accountId = tokenPayload.UserID
	}

	_, err = server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID: sql.NullInt32{
			Int32: accountId,
			Valid: true,
		},
		IsVerify: sql.NullBool{
			Bool:  req.GetStatus(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to inactive user: %v", err))
	}

	return &pb.AccountInactiveResponse{
		Code:    200,
		Message: "success",
	}, nil
}
