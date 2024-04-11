package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/longIdt2502/pharmago_be/woker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *ServerGRPC) SendCode(ctx context.Context, req *pb.SendCodeRequest) (*pb.SendCodeResponse, error) {

	account, err := server.store.GetAccountByMail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account")
	}

	code := utils.RandomInt(100000, 999999)

	_ = server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, &woker.PayloadSendVerifyEmail{
		Username: account.Username,
		Code:     fmt.Sprintf("%d", code),
	})

	verify, err := server.store.CreateVerify(ctx, db.CreateVerifyParams{
		Username:   account.Username,
		Email:      account.Email,
		SecretCode: strconv.FormatInt(code, 10),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send code to your mail")
	}

	return &pb.SendCodeResponse{
		Code:    200,
		Message: "success",
		Details: verify.ID,
	}, nil
}

func (server *ServerGRPC) VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.VerifyCodeResponse, error) {
	verify, err := server.store.GetVerify(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account")
	}

	if verify.SecretCode != req.GetCode() {
		return nil, status.Errorf(codes.Unimplemented, "failed to verify")
	}

	_, err = server.store.UpdateVerify(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify")
	}

	return &pb.VerifyCodeResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	verify, err := server.store.GetVerify(ctx, req.GetIdVerify())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "verify not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get verify")
	}
	if verify.SecretCode != req.GetCodeVerify() {
		return nil, status.Errorf(codes.Internal, "failed to verify")
	}

	hashPass, err := utils.HashedPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	_, err = server.store.ResetPassword(ctx, db.ResetPasswordParams{
		HashedPassword: sql.NullString{
			String: hashPass,
			Valid:  true,
		},
		Email: sql.NullString{
			String: verify.Email,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to reset password")
	}

	return &pb.ResetPasswordResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) CheckEmail(ctx context.Context, req *pb.CheckEmailRequest) (*pb.CheckEmailResponse, error) {
	_, err := server.store.GetAccountByMail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account")
	}

	return &pb.CheckEmailResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) CheckPhone(ctx context.Context, req *pb.CheckPhoneRequest) (*pb.CheckPhoneResponse, error) {
	_, err := server.store.GetAccountByPhone(ctx, req.GetPhone())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account")
	}

	return &pb.CheckPhoneResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	auth, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	timePb := timestamppb.New(auth.ExpireAt)

	return &pb.CheckTokenResponse{
		Code:    200,
		Message: "success",
		Details: timePb,
	}, nil
}
