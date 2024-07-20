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

func (server *ServerGRPC) ScheduleCreate(ctx context.Context, req *pb.AppointmentSchedule) (*pb.AppointmentScheduleResponse, error) {
	account, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("SCHEDULE-%s", utils.RandomString(6))
	schedule, err := server.store.CreateSchedule(ctx, db.CreateScheduleParams{
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
		},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.AppointmentScheduleResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi tạo lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	for _, item := range req.GetServices() {
		_, err = server.store.CreateScheduleService(ctx, db.CreateScheduleServiceParams{
			AsUuid: schedule.Uuid,
			Service: sql.NullInt32{
				Int32: item.GetServiceId(),
				Valid: true,
			},
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.AppointmentScheduleResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: "Lỗi tạo dịch vụ lịch hẹn",
				Log:          errApp.Log,
			}, nil
		}
	}

	for _, item := range req.GetDrugs() {
		_, err = server.store.CreateScheduleDrug(ctx, db.CreateScheduleDrugParams{
			AsUuid:   schedule.Uuid,
			Variant:  sql.NullInt32{Int32: item.GetVariantId(), Valid: true},
			LieuDung: sql.NullString{String: item.GetLieuDung(), Valid: item.LieuDung != nil},
			Quantity: item.GetQuantity(),
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.AppointmentScheduleResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: "Lỗi tạo đơn thuốc",
				Log:          errApp.Log,
			}, nil
		}
	}

	return &pb.AppointmentScheduleResponse{
		Code:    200,
		Message: "success",
		Details: &pb.AppointmentSchedule{
			Id:   schedule.ID,
			Uuid: schedule.Uuid.String(),
		},
	}, nil
}

func (server *ServerGRPC) ScheduleList(ctx context.Context, req *pb.AppointmentScheduleListRequest) (*pb.AppointmentScheduleListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	schedules, err := server.store.GetListSchedule(ctx, db.GetListScheduleParams{
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
		return &pb.AppointmentScheduleListResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	var schedulesPb []*pb.AppointmentSchedule
	for _, item := range schedules {
		itemPb := mapper.AppointmentScheduleMapper(ctx, server.store, item)
		schedulesPb = append(schedulesPb, itemPb)
	}

	return &pb.AppointmentScheduleListResponse{
		Code:    200,
		Message: "success",
		Details: schedulesPb,
	}, nil
}

func (server *ServerGRPC) ScheduleDetail(ctx context.Context, req *pb.AppointmentSchedule) (*pb.AppointmentScheduleResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	uuidParse, _ := uuid.Parse(req.GetUuid())

	schedules, err := server.store.GetListSchedule(ctx, db.GetListScheduleParams{
		Company: sql.NullInt32{
			Int32: req.Company,
			Valid: true,
		},
		Uuid: uuid.NullUUID{UUID: uuidParse, Valid: true},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.AppointmentScheduleResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	if len(schedules) == 0 {
		return &pb.AppointmentScheduleResponse{
			Code:         404,
			MessageTrans: "Dữ liệu lịch hẹn không tồn tại",
		}, nil
	}

	itemPb := mapper.AppointmentScheduleMapper(ctx, server.store, schedules[0])

	return &pb.AppointmentScheduleResponse{
		Code:    200,
		Message: "success",
		Details: itemPb,
	}, nil
}

func (server *ServerGRPC) ScheduleUpdate(ctx context.Context, req *pb.AppointmentScheduleUpdateRequest) (*pb.AppointmentScheduleUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	uuidParse, _ := uuid.Parse(req.GetUuid())

	schedule, err := server.store.DetailSchedule(ctx, uuidParse)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.AppointmentScheduleUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi lấy dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	_, err = server.CreateMannyMediaRecord(ctx, req.GetFiles(), req.GetType().String(), tokenPayload.UserID, schedule.Customer.Int32, &uuidParse)
	if err != nil {
		errApp := common.ErrInternal(err)
		return &pb.AppointmentScheduleUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi tạo dữ liệu",
			Log:          errApp.Log,
		}, nil
	}

	_, err = server.store.UpdateSchedule(ctx, db.UpdateScheduleParams{
		Uuid: uuidParse,
		IsDone: sql.NullBool{
			Bool:  req.GetIsDone(),
			Valid: true,
		},
		Diagnostic: sql.NullString{
			String: req.GetDiagnostic(),
			Valid:  true,
		},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.AppointmentScheduleUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: "Lỗi cập nhật dữ liệu lịch hẹn",
			Log:          errApp.Log,
		}, nil
	}

	return &pb.AppointmentScheduleUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}
