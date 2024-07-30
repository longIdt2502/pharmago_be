package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
)

func (server *ServerGRPC) MedicalBillCreate(ctx context.Context, req *pb.MedicalBill) (*pb.MedicalBillResponse, error) {
	account, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("SCHEDULE-%s", utils.RandomString(6))
	medicalBill, err := server.store.CreateMedicalBill(ctx, db.CreateMedicalBillParams{
		Uuid:        uuid.New(),
		Code:        code,
		Customer:    sql.NullInt32{Int32: req.GetCustomerId(), Valid: true},
		Company:     sql.NullInt32{Int32: req.GetCompany(), Valid: true},
		Doctor:      sql.NullInt32{Int32: req.GetDoctorId(), Valid: true},
		Symptoms:    sql.NullString{String: req.GetSymptoms(), Valid: req.Symptoms != nil},
		Diagnostic:  sql.NullString{String: req.GetDiagnostic(), Valid: req.Diagnostic != nil},
		IsDone:      req.GetIsDone(),
		MeetingAt:   time.Unix(req.GetMeetingAt().GetSeconds(), 0),
		UserCreated: account.UserID,
		UserUpdated: sql.NullInt32{
			Int32: account.UserID,
			Valid: true,
		},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi tạo lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	for _, item := range req.GetServices() {
		_, err = server.store.CreateScheduleService(ctx, db.CreateScheduleServiceParams{
			MbUuid: uuid.NullUUID{
				UUID:  medicalBill.Uuid,
				Valid: true,
			},
			Service: sql.NullInt32{
				Int32: item.GetServiceId(),
				Valid: true,
			},
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.MedicalBillResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: "Lỗi tạo dịch vụ lịch hẹn",
				Log:          errApp.Log,
			}, nil
		}
	}

	// for _, item := range req.GetDrugs() {
	// 	_, err = server.store.CreateScheduleDrug(ctx, db.CreateScheduleDrugParams{
	// 		MbUuid: uuid.NullUUID{
	// 			UUID:  medicalBill.Uuid,
	// 			Valid: true,
	// 		},
	// 		Variant:  sql.NullInt32{Int32: item.GetVariantId(), Valid: true},
	// 		LieuDung: sql.NullString{String: item.GetLieuDung(), Valid: item.LieuDung != nil},
	// 		Quantity: item.GetQuantity(),
	// 	})
	// 	if err != nil {
	// 		errApp := common.ErrDB(err)
	// 		return &pb.MedicalBillResponse{
	// 			Code:         int32(errApp.StatusCode),
	// 			Message:      errApp.Message,
	// 			MessageTrans: "Lỗi tạo đơn thuốc",
	// 			Log:          errApp.Log,
	// 		}, nil
	// 	}
	// }

	return &pb.MedicalBillResponse{
		Code:    200,
		Message: "success",
		Details: &pb.MedicalBill{
			Id:   medicalBill.ID,
			Uuid: medicalBill.Uuid.String(),
		},
	}, nil
}

func (server *ServerGRPC) MedicalBillList(ctx context.Context, req *pb.MedicalBillListRequest) (*pb.MedicalBillListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	medicalBills, err := server.store.GetListMedicalBill(ctx, db.GetListMedicalBillParams{
		Company: sql.NullInt32{
			Int32: req.Company,
			Valid: true,
		},
		Search:       sql.NullString{String: req.GetSearch(), Valid: req.Search != nil},
		Doctor:       sql.NullInt32{Int32: req.GetDoctor(), Valid: req.Doctor != nil},
		CreatedStart: sql.NullTime{Time: time.Unix(req.GetTimeStart().GetSeconds(), 0), Valid: req.TimeStart != nil},
		CreatedEnd:   sql.NullTime{Time: time.Unix(req.GetTimeEnd().GetSeconds(), 0), Valid: req.TimeEnd != nil},
		Page:         sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:        sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillListResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	var medicalBillPb []*pb.MedicalBill
	for _, item := range medicalBills {
		itemPb := mapper.MedicalBillMapper(ctx, server.store, item)
		medicalBillPb = append(medicalBillPb, itemPb)
	}

	return &pb.MedicalBillListResponse{
		Code:    200,
		Message: "success",
		Details: medicalBillPb,
	}, nil
}

func (server *ServerGRPC) MedicalBillDetail(ctx context.Context, req *pb.MedicalBill) (*pb.MedicalBillResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	uuidParse, _ := uuid.Parse(req.GetUuid())

	medicalBill, err := server.store.GetListMedicalBill(ctx, db.GetListMedicalBillParams{
		Company: sql.NullInt32{
			Int32: req.Company,
			Valid: true,
		},
		Uuid: uuid.NullUUID{UUID: uuidParse, Valid: true},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	if len(medicalBill) == 0 {
		return &pb.MedicalBillResponse{
			Code:         404,
			MessageTrans: "Dữ liệu lịch hẹn không tồn tại",
		}, nil
	}

	itemPb := mapper.MedicalBillMapper(ctx, server.store, medicalBill[0])

	var paymentsPB []*pb.Payment
	paymentSell, err := server.store.PaymentOrderByMedicalBill(ctx, medicalBill[0].Uuid)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu thông tin thanh toán đơn bán",
			Log:          errApp.Log,
		}, nil
	}
	paymentsPB = append(paymentsPB, &pb.Payment{
		Code:     "SELL",
		MustPaid: float32(paymentSell.TotalMustPaid),
		HadPaid:  float32(paymentSell.TotalHadPaid),
		NeedPay:  float32(paymentSell.TotalNeedPay),
	})

	paymentService, err := server.store.PaymentOrderServiceByMedicalBill(ctx, medicalBill[0].Uuid)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu thông tin thanh toán đơn dịch vụ",
			Log:          errApp.Log,
		}, nil
	}
	paymentsPB = append(paymentsPB, &pb.Payment{
		Code:     "SERVICE",
		MustPaid: float32(paymentService.TotalMustPaid),
		HadPaid:  float32(paymentService.TotalHadPaid),
		NeedPay:  float32(paymentService.TotalNeedPay),
	})

	itemPb.Payment = paymentsPB

	return &pb.MedicalBillResponse{
		Code:    200,
		Message: "success",
		Details: itemPb,
	}, nil
}

func (server *ServerGRPC) MedicalBillUpdate(ctx context.Context, req *pb.MedicalBillUpdateRequest) (*pb.MedicalBillUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	uuidParse, _ := uuid.Parse(req.GetUuid())

	medicalBill, err := server.store.DetailMedicalBill(ctx, uuidParse)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	if len(req.Files) != 0 {
		_, err = server.CreateMannyMediaRecord(ctx, req.GetFiles(), req.GetType().String(), tokenPayload.UserID, medicalBill.Customer.Int32, nil, &uuidParse)
		if err != nil {
			errApp := common.ErrInternal(err)
			return &pb.MedicalBillUpdateResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: "Lỗi tạo dữ liệu",
				Log:          errApp.Log,
			}, nil
		}
	}

	_, err = server.store.UpdateMedicalBill(ctx, db.UpdateMedicalBillParams{
		Diagnostic: sql.NullString{String: req.GetDiagnostic(), Valid: req.Diagnostic != nil},
		Symptoms:   sql.NullString{String: req.GetSymptoms(), Valid: req.Symptoms != nil},
		Uuid:       uuidParse,
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.MedicalBillUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi cập nhật dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	return &pb.MedicalBillUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}
