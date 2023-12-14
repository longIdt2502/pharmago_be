package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

func (server *ServerGRPC) VerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.VerifyAccountResponse, error) {
	violations := validateVerifyAccount(req)
	if violations != nil {
		return nil, config.InvalidArgumentError(violations)
	}

	verify, err := server.store.GetVerify(ctx, int64(req.IdVerify))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "verify id invalid:", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get verify:", err)
	}

	if verify.ExpiredAt.Before(time.Now()) {
		return nil, status.Errorf(codes.Unavailable, "verify have been expired")
	}

	if verify.SecretCode != req.SecretCode {
		return nil, status.Errorf(codes.InvalidArgument, "code incorrect")
	}

	_, err = server.store.UpdateVerify(ctx, int64(req.IdVerify))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify account:", err)
	}

	_, err = server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		IsVerify: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		ID: sql.NullInt64{},
		Username: sql.NullString{
			String: verify.Username,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update verify account:", err)
	}

	rsp := &pb.VerifyAccountResponse{
		Code:    200,
		Message: "success",
		Details: true,
	}

	return rsp, nil
}

func validateVerifyAccount(req *pb.VerifyAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if len(req.SecretCode) != 6 {
		violations = append(violations, config.FieldViolation("username", fmt.Errorf("code must be 6 charecters")))
	}

	if _, err := strconv.Atoi(req.SecretCode); err != nil || len(req.SecretCode) != 6 {
		violations = append(violations, config.FieldViolation("username", err))
	}

	return violations
}
