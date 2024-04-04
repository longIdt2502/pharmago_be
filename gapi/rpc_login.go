package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *ServerGRPC) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	account, err := server.store.GetAccountByUseName(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to find user: %e", err))
	}
	if !account.IsVerify {
		return nil, status.Errorf(codes.Unauthenticated, "account not verify")
	}

	password := req.GetPassword()
	err = utils.CheckPassword(password, account.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("password incorrect: %e", err))
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(username, account.ID, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	refreshTokenDuration, err := time.ParseDuration("24h")
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create refresh token duration: %e", err))
	}
	refreshToken, refreshTokenPayload, _ := server.tokenMaker.CreateToken(username, account.ID, refreshTokenDuration)

	metadata := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshTokenPayload.ID,
		Username:     username,
		RefreshToken: refreshToken,
		UserAgent:    metadata.UserAgent,
		ClientIp:     metadata.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshTokenPayload.ExpireAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %e", err)
	}

	accountResponse := mapper.AccountMapper(account)

	rsp := &pb.LoginResponse{
		Code:    200,
		Message: "success",
		Details: &pb.LoginResponseData{
			Account:               accountResponse,
			SessionId:             session.ID.String(),
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiresAt:  timestamppb.New(accessTokenPayload.ExpireAt),
			RefreshTokenExpiresAt: timestamppb.New(refreshTokenPayload.ExpireAt),
		},
	}
	return rsp, nil
}
