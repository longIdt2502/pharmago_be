package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListMedicalRecord(ctx context.Context, req *pb.ListMedicalRecordRequest) (*pb.ListMedicalRecordResponse, error) {

	medicalRecords, err := server.store.ListMedicalRecord(ctx, db.ListMedicalRecordParams{
		Customer: req.GetCustomer(),
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  true,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: true,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get record")
	}

	var medicalRecordsPb []*pb.MedicalRecord
	for _, item := range medicalRecords {
		medicalRecordsPb = append(medicalRecordsPb, mapper.MedicalRecordMapper(item))
	}

	return &pb.ListMedicalRecordResponse{
		Code:    200,
		Message: "success",
		Details: medicalRecordsPb,
	}, nil
}

func (server *ServerGRPC) CreateMedicalRecord(ctx context.Context, req *pb.CreateMedicalRecordRequest) (*pb.CreateMedicalRecordResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("MR-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))

	medicalRecord, err := server.store.CreateMedicalRecord(ctx, db.CreateMedicalRecordParams{
		Code:     code,
		Customer: req.GetCustomer(),
		Weight: sql.NullFloat64{
			Float64: float64(req.GetWeight()),
			Valid:   req.Weight != nil,
		},
		Long: sql.NullFloat64{
			Float64: float64(req.GetLong()),
			Valid:   req.Long != nil,
		},
		Symptom:    req.GetSymptom(),
		Diagnostic: req.GetDiagnostic(),
		Result:     req.GetResult(),
		Doctor: sql.NullInt32{
			Int32: req.GetDoctor(),
			Valid: true,
		},
		ReExamination: req.GetReExamination(),
		Note: sql.NullString{
			String: req.GetNote(),
			Valid:  true,
		},
		UserCreated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create medical record: %e", err))
	}

	return &pb.CreateMedicalRecordResponse{
		Code:    200,
		Message: "success",
		Details: medicalRecord.ID,
	}, nil
}

func (server *ServerGRPC) DetailMedicalRecord(ctx context.Context, req *pb.DetailMedicalRecordRequest) (*pb.DetailMedicalRecordResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	medicalRecord, err := server.store.DetailMedicalRecord(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "medical record not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get medical record")
	}

	dataPb := mapper.MedicalRecordMapper(medicalRecord)

	return &pb.DetailMedicalRecordResponse{
		Code:    200,
		Message: "success",
		Details: dataPb,
	}, nil
}
