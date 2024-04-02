package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListDebtNote(ctx context.Context, req *pb.ListDebtNoteRequest) (*pb.ListDebtNoteResponse, error) {
	debtNoteDb, err := server.store.GetListDebtNote(ctx, db.GetListDebtNoteParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
		Status: sql.NullString{
			String: req.GetStatus(),
			Valid:  req.Status != nil,
		},
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
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get debt note: %e", err))
	}

	var debtNotePb []*pb.DebtNote
	for _, item := range debtNoteDb {
		pb := mapper.DebtNoteMapper(item, nil)
		debtNotePb = append(debtNotePb, pb)
	}

	return &pb.ListDebtNoteResponse{
		Code:    200,
		Message: "succes",
		Details: debtNotePb,
	}, nil
}

func (server *ServerGRPC) CreateDebtNote(ctx context.Context, req *pb.CreateDebtNoteRequest) (*pb.CreateDebtNoteResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	debtCode := fmt.Sprintf("CNT-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		debtCode = req.GetCode()
	}

	statusDebt := "OPEN"
	if req.GetPaymented() != 0 && req.GetPaymented() < req.GetMoney() {
		statusDebt = "REPAYING"
	} else if req.GetPaymented() == req.GetMoney() {
		statusDebt = "SETTLED"
	}

	debtNoteDb, err := server.store.CreateDebtNote(ctx, db.CreateDebtNoteParams{
		Code: debtCode,
		Title: sql.NullString{
			String: req.GetTitle(),
			Valid:  req.Title != nil,
		},
		Entity:    req.GetEntity(),
		Money:     float64(req.GetMoney()),
		Paymented: float64(req.GetPaymented()),
		Note: sql.NullString{
			String: req.GetNote(),
			Valid:  req.Note != nil,
		},
		Type:        req.GetType(),
		Status:      statusDebt,
		Company:     req.GetCompany(),
		UserCreated: tokenPayload.UserID,
		Exprise:     time.Unix(req.GetExprise().GetSeconds(), 0),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record debt note: %e", err))
	}

	if req.Paymented != 0 {
		repaymentCode := fmt.Sprintf("CNT-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
		_, _ = server.store.CreateRepayment(ctx, db.CreateRepaymentParams{
			Code:        repaymentCode,
			Money:       float64(req.GetPaymented()),
			UserCreated: tokenPayload.UserID,
			Debt:        debtNoteDb.ID,
		})
	}

	return &pb.CreateDebtNoteResponse{
		Code:    200,
		Message: "success",
		Details: debtNoteDb.ID,
	}, nil
}

func (server *ServerGRPC) DetailDebtNote(ctx context.Context, req *pb.DetailDebtNoteRequest) (*pb.DetailDebtNoteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	debtNote, err := server.store.DetailDebtNote(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "debt note not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get detail debt note")
	}

	debtPayment, err := server.store.ListRepayment(ctx, debtNote.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list debt repayment")
	}

	debtNotePb := mapper.DebtNoteMapper(debtNote, &debtPayment)
	return &pb.DetailDebtNoteResponse{
		Code:    200,
		Message: "success",
		Details: debtNotePb,
	}, nil
}

func (server *ServerGRPC) CreateDebtRepayment(ctx context.Context, req *pb.CreateDebtRepaymentRequest) (*pb.CreateDebtRepaymentResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	repaymentCode := fmt.Sprintf("CNT-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	repayment, err := server.store.CreateRepayment(ctx, db.CreateRepaymentParams{
		Code:        repaymentCode,
		Money:       float64(req.GetMoney()),
		UserCreated: tokenPayload.UserID,
		Debt:        req.GetDebt(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record repayment: %e", err))
	}

	return &pb.CreateDebtRepaymentResponse{
		Code:    200,
		Message: "success",
		Details: repayment.ID,
	}, nil
}
