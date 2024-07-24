package gapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
)

func (server *ServerGRPC) PrescriptionCreate(ctx context.Context, req *pb.Prescription) (*pb.PrescriptionResponse, error) {
	account, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("DT-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))

	prescription, err := server.store.CreatePrescription(ctx, db.CreatePrescriptionParams{
		Uuid:        uuid.New(),
		Code:        code,
		Symptoms:    sql.NullString{String: req.GetSymptoms(), Valid: req.Symptoms != nil},
		Diagnostic:  sql.NullString{String: req.GetDiagnostic(), Valid: req.Diagnostic != nil},
		Customer:    sql.NullInt32{Int32: req.GetCustomerId(), Valid: req.CustomerId != nil},
		Doctor:      sql.NullInt32{Int32: req.GetDoctorId(), Valid: true},
		Company:     req.GetCompany(),
		UserCreated: account.UserID,
		UserUpdated: sql.NullInt32{},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.PrescriptionResponse{
			Code:    int32(errApp.StatusCode),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	if req.MbUuid != nil {
		mbUuidParse, err := uuid.Parse(req.GetMbUuid())
		if err != nil {
			errApp := common.ErrInternal(err)
			return &pb.PrescriptionResponse{
				Code:    int32(errApp.StatusCode),
				Message: errApp.Message,
				Log:     errApp.Log,
			}, nil
		}
		_, err = server.store.UpdateMedicalBill(ctx, db.UpdateMedicalBillParams{
			Prescription: uuid.NullUUID{UUID: prescription.Uuid, Valid: true},
			Uuid:         mbUuidParse,
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.PrescriptionResponse{
				Code:    int32(errApp.StatusCode),
				Message: errApp.Message,
				Log:     errApp.Log,
			}, nil
		}
	}

	for _, item := range req.GetItems() {
		_, err = server.store.CreatePrescriptionItem(ctx, db.CreatePrescriptionItemParams{
			PrescriptionUuid: uuid.NullUUID{
				UUID:  prescription.Uuid,
				Valid: true,
			},
			Variant:  sql.NullInt32{Int32: item.GetVariantId(), Valid: true},
			LieuDung: sql.NullString{String: item.GetLieuDung(), Valid: item.LieuDung != nil},
			Quantity: item.Quantity,
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.PrescriptionResponse{
				Code:    int32(errApp.StatusCode),
				Message: errApp.Message,
				Log:     errApp.Log,
			}, nil
		}
	}

	return &pb.PrescriptionResponse{
		Code:    200,
		Message: "success",
		Details: &pb.Prescription{
			Id:   prescription.ID,
			Uuid: prescription.Uuid.String(),
		},
	}, nil
}

func (server *ServerGRPC) PrescriptionDetail(ctx context.Context, req *pb.Prescription) (*pb.PrescriptionResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	uuidPare, _ := uuid.Parse(req.GetUuid())

	prescription, err := server.store.DetailPrescription(ctx, uuidPare)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.PrescriptionResponse{
			Code:    int32(errApp.StatusCode),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	return &pb.PrescriptionResponse{
		Code:    200,
		Message: "success",
		Details: mapper.PrescriptionMapper(ctx, server.store, prescription),
	}, nil
}

func (server *ServerGRPC) PrescriptionList(ctx context.Context, req *pb.PrescriptionListRequest) (*pb.PrescriptionListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	prescriptions, err := server.store.ListPrescription(ctx, db.ListPrescriptionParams{
		Company: sql.NullInt32{Int32: req.GetCompany(), Valid: true},
		Search:  sql.NullString{String: req.GetSearch(), Valid: true},
		Page:    sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:   sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.PrescriptionListResponse{
			Code:    int32(errApp.StatusCode),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	var prescriptionsPb []*pb.Prescription
	for _, item := range prescriptions {
		prescriptionsPb = append(prescriptionsPb, mapper.PrescriptionListItemMapper(ctx, server.store, item))
	}

	return &pb.PrescriptionListResponse{
		Code:    200,
		Message: "success",
		Details: prescriptionsPb,
	}, nil
}
