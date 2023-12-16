package gapi

import (
	"context"
	"database/sql"
	"errors"
	"github.com/hibiken/asynq"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/longIdt2502/pharmago_be/validate"
	"github.com/longIdt2502/pharmago_be/woker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

func (server *ServerGRPC) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	violation := validateCreateAccountRequest(req)
	if violation != nil {
		return nil, config.InvalidArgumentError(violation)
	}

	password := req.GetPassword()
	hashedPassword, err := utils.HashedPassword(password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't hashed password:", err)
	}

	accountType, err := server.store.GetAccountType(ctx, db.GetAccountTypeParams{
		ID: sql.NullInt64{},
		Code: sql.NullString{
			String: req.AccountType,
			Valid:  true,
		},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account code doesn't exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account type")
	}

	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Type:           accountType.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create record account:", err)
	}

	randomCode := utils.RandomInt(100000, 999999)
	secretCode := strconv.Itoa(int(randomCode))
	verify, err := server.store.CreateVerify(ctx, db.CreateVerifyParams{
		Username:   req.Username,
		Email:      req.Email,
		SecretCode: secretCode,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create verify:", err)
	}

	// ===== Redis task distributor send email
	taskPayload := &woker.PayloadSendVerifyEmail{
		Username: req.Username,
		Code:     secretCode,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(woker.QueueCritical),
	}

	err = server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to run task email: ", err)
	}

	accountMapper := mapper.AccountMapper(account)
	rsp := &pb.CreateAccountResponse{
		Code:     int32(200),
		Message:  "success",
		Details:  accountMapper,
		VerifyId: int32(verify.ID),
	}
	return rsp, nil
}

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, config.FieldViolation("username", err))
	}

	if err := validate.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, config.FieldViolation("full_name", err))
	}

	if err := validate.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, config.FieldViolation("password", err))
	}

	if err := validate.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, config.FieldViolation("email", err))
	}
	return violations
}
