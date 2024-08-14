package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/woker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) OrderCreate(ctx context.Context, req *pb.OrderCreateRequest) (*pb.OrderCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	result, err := server.store.CreateOrderTx(ctx, db.CreateOrderTxParams{
		OrderCreateRequest: req,
		B2Bucket:           server.b2Bucket,
		TokenPayload:       tokenPayload,
	})
	if err != nil {
		errApp := err.(*common.AppError)
		return &pb.OrderCreateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: errApp.MessageTrans,
			Log:          errApp.Log,
		}, nil
	}

	company, _ := server.store.GetCompanyById(ctx, req.Order.Company)
	order, _ := server.store.DetailOrder(ctx, db.DetailOrderParams{ID: sql.NullInt32{
		Int32: result.Id,
		Valid: true,
	}})

	customer, _ := server.store.DetailCustomer(ctx, req.Order.GetCustomer())

	payload := &woker.PayloadZNS{
		OaID: company.OaID.String,
		Data: woker.PayloadZNSData{
			Name:       customer.FullName,
			Status:     "Chờ xác nhận",
			CreatedAt:  time.Now().Format("15:04:05 02/01/2006"),
			Total:      strconv.FormatFloat(float64(req.Order.GetTotalPrice()), 'f', -1, 64),
			Phone:      customer.Phone.String,
			Code:       order.Code,
			OrderItems: "Thuốc",
		},
		Phone: customer.Phone.String,
		Mode:  "production",
		Type:  "confirm_order",
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(1 * time.Second),
		asynq.Queue(woker.QueueCritical),
	}

	_ = server.taskDistributor.DistributorTaskSendOrderZns(ctx, payload, opts...)

	if len(req.ServiceItems) != 0 {
		services, err := server.store.ListOrderServiceItem(ctx, result.Id)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get service item")
		}

		for _, item := range services {
			if item.ReminderTime.Valid {
				date := time.Unix(time.Now().Unix()+int64(item.ReminderTime.Int32), 0)
				payloadTaskFcm := &woker.PayloadSendFcm{
					To:      fmt.Sprintf("/topics/COMPANY_%s", company.Code),
					Title:   "Thông báo dịch vụ",
					Body:    fmt.Sprintf("%s: có lịch dịch vụ %s với khách hàng (%s)", company.Name, item.Title, date.Format("02/01")),
					Company: company.ID,
					Data: &woker.DataNoti{
						Order:   &result.Id,
						Service: &item.ID_2,
					},
				}

				opts := []asynq.Option{
					asynq.MaxRetry(10),
					asynq.ProcessIn(time.Duration(item.ReminderTime.Int32) * time.Second),
					asynq.Queue(woker.QueueCritical),
				}

				_ = server.taskDistributor.DistributorTaskSendFcm(ctx, payloadTaskFcm, opts...)
			}
		}
	}

	return &pb.OrderCreateResponse{
		Code:    200,
		Message: "success",
		Details: result.Id,
	}, nil
}

func (server *ServerGRPC) OrderList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}
	a := time.Unix(req.CreatedStart.GetSeconds(), 0)
	print(a.String())
	orders, err := server.store.ListOrder(ctx, db.ListOrderParams{
		Company:      sql.NullInt32{Int32: req.Company, Valid: true},
		Status:       sql.NullString{String: req.GetStatus(), Valid: req.Status != nil},
		Warehouse:    sql.NullInt32{Int32: req.GetWarehouse(), Valid: req.Warehouse != nil},
		Type:         sql.NullString{String: req.GetType().String(), Valid: req.Type != nil},
		Customer:     sql.NullInt32{Int32: req.GetCustomer(), Valid: req.Customer != nil},
		Search:       sql.NullString{String: req.GetSearch(), Valid: req.Search != nil},
		CreatedStart: sql.NullTime{Time: time.Unix(req.CreatedStart.GetSeconds(), 0), Valid: req.CreatedStart != nil},
		CreatedEnd:   sql.NullTime{Time: time.Unix(req.CreatedEnd.GetSeconds(), 0), Valid: req.CreatedEnd != nil},
		UpdatedStart: sql.NullTime{Time: time.Unix(req.UpdatedStart.GetSeconds(), 0), Valid: req.UpdatedStart != nil},
		UpdatedEnd:   sql.NullTime{Time: time.Unix(req.UpdatedEnd.GetSeconds(), 0), Valid: req.UpdatedEnd != nil},
		OrderBy:      sql.NullString{String: req.GetOrderBy(), Valid: req.OrderBy != nil},
		Page:         sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:        sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get orders: %e", err)
	}

	var ordersPb []*pb.OrderPreview

	for _, value := range orders {
		dataPb := mapper.OrderPreviewMapper(ctx, server.store, value)
		ordersPb = append(ordersPb, dataPb)
	}

	countData, _ := server.store.CountOrderByType(ctx, req.Company)

	sellCount := 0
	serviceCount := 0
	for _, item := range countData {
		switch item.Code.String {
		case "SELL":
			sellCount = int(item.Count)
		case "SERVICE":
			serviceCount = int(item.Count)
		}
	}

	return &pb.OrderListResponse{
		Code:    200,
		Message: "success",
		Details: ordersPb,
		Count: &pb.OrderListResponseCount{
			Sell:    int32(sellCount),
			Service: int32(serviceCount),
		},
	}, nil
}

func (server *ServerGRPC) OrderListByMedicalBill(ctx context.Context, req *pb.OrdersByMedicalBillRequest) (*pb.OrderListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	uuidParse, _ := uuid.Parse(req.GetMbUuid())
	orders, err := server.store.ListByMedicalBill(ctx, db.ListByMedicalBillParams{
		Uuid:   uuidParse,
		Status: sql.NullString{String: req.GetStatus(), Valid: req.Status != nil},
		Type:   sql.NullString{String: req.GetType().String(), Valid: req.Type != nil},
		Search: sql.NullString{String: req.GetSearch(), Valid: req.Search != nil},
		Page:   sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:  sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get orders: %e", err)
	}

	var ordersPb []*pb.OrderPreview

	for _, value := range orders {
		dataPb := mapper.OrderMedicalBillMapper(value)
		ordersPb = append(ordersPb, dataPb)
	}

	// countData, _ := server.store.CountOrderByType(ctx, req.Company)

	// sellCount := 0
	// serviceCount := 0
	// for _, item := range countData {
	// 	switch item.Code.String {
	// 	case "SELL":
	// 		sellCount = int(item.Count)
	// 	case "SERVICE":
	// 		serviceCount = int(item.Count)
	// 	}
	// }

	return &pb.OrderListResponse{
		Code:    200,
		Message: "success",
		Details: ordersPb,
		// Count: &pb.OrderListResponseCount{
		// 	Sell:    int32(sellCount),
		// 	Service: int32(serviceCount),
		// },
	}, nil
}

func (server *ServerGRPC) OrderDetail(ctx context.Context, req *pb.OrderDetailRequest) (*pb.OrderDetailResponse, error) {
	orderDb, err := server.store.DetailOrder(ctx, db.DetailOrderParams{
		ID: sql.NullInt32{
			Int32: req.GetId(),
			Valid: true,
		},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "order not exists")
		}
		return nil, common.ErrDB(err) // status.Errorf(codes.Internal, "failed to get order: %e", err)
	}

	orderPb := mapper.OrderDetailMapper(ctx, server.store, orderDb)

	return &pb.OrderDetailResponse{
		Code:    200,
		Message: "success",
		Details: orderPb,
	}, nil
}

func (server *ServerGRPC) OrderUpdateStatus(ctx context.Context, req *pb.OrderUpdateStatusRequest) (*pb.OrderUpdateStatusResponse, error) {
	order, err := server.store.DetailOrder(ctx, db.DetailOrderParams{
		ID: sql.NullInt32{
			Int32: req.GetId(),
			Valid: true,
		},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "order not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get order: %e", err)
	}

	if order.OsCode == "COMPLETE" || order.OsCode == "CANCEL" {
		return nil, status.Errorf(codes.InvalidArgument, "can't update status")
	}

	// if order.OsCode == "COMPLETE" {
	// 	// TODO: edit order create, when status complete so minus quantity consignment
	// }

	_, err = server.store.UpdateStatusOrder(ctx, db.UpdateStatusOrderParams{
		ID:     req.Id,
		Status: req.Code.String(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update status order: %e", err)
	}

	return &pb.OrderUpdateStatusResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) OrderScan(ctx context.Context, req *pb.OrderScanRequest) (*pb.OrderScanResponse, error) {
	orderDb, err := server.store.DetailOrder(ctx, db.DetailOrderParams{
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  true,
		},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "order not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get order: %e", err)
	}

	if orderDb.OtCode != "PRESCRIPTION" {
		return nil, status.Errorf(codes.NotFound, "order not exists")
	}

	orderPb := mapper.OrderDetailMapper(ctx, server.store, orderDb)

	return &pb.OrderScanResponse{
		Code:    200,
		Message: "success",
		Details: orderPb,
	}, nil
}

func (server *ServerGRPC) CreatePaymentItemOrder(ctx context.Context, req *pb.PaymentItemOrderRequest) (*pb.PaymentItemOrderResponse, error) {
	orderDb, err := server.store.DetailOrder(ctx, db.DetailOrderParams{
		ID: sql.NullInt32{Int32: req.GetOrderId(), Valid: true},
	})
	if err != nil {
		appErr := common.ErrDBWithMsg(err, "lỗi lấy dữ liệu đơn hàng")
		return &pb.PaymentItemOrderResponse{
			Code:         int32(appErr.StatusCode),
			Message:      appErr.Message,
			MessageTrans: appErr.MessageTrans,
			Log:          appErr.Log,
		}, nil
	}

	payment, err := server.store.DetailPayment(ctx, orderDb.Payment)
	if err != nil {
		appErr := common.ErrDBWithMsg(err, "lỗi lấy dữ liệu thanh toán")
		return &pb.PaymentItemOrderResponse{
			Code:         int32(appErr.StatusCode),
			Message:      appErr.Message,
			MessageTrans: appErr.MessageTrans,
			Log:          appErr.Log,
		}, nil
	}

	if req.GetValue() > float32(payment.NeedPay) {
		appErr := common.ErrInvalidRequest(errors.New("giá trị không hợp lệ"))
		return &pb.PaymentItemOrderResponse{
			Code:         int32(appErr.StatusCode),
			Message:      appErr.Message,
			MessageTrans: appErr.MessageTrans,
			Log:          appErr.Log,
		}, nil
	}

	_, err = server.store.CreatePaymentItem(ctx, db.CreatePaymentItemParams{
		Type:      req.GetType().String(),
		Value:     float64(req.GetValue()),
		IsPaid:    true,
		Payment:   orderDb.Payment,
		ExtraNote: sql.NullString{},
	})
	if err != nil {
		appErr := common.ErrDBWithMsg(err, "lỗi tạo thanh toán")
		return &pb.PaymentItemOrderResponse{
			Code:         int32(appErr.StatusCode),
			Message:      appErr.Message,
			MessageTrans: appErr.MessageTrans,
			Log:          appErr.Log,
		}, nil
	}

	_, err = server.store.UpdatePayment(ctx, db.UpdatePaymentParams{
		HadPaid: sql.NullFloat64{Float64: payment.HadPaid + float64(req.GetValue()), Valid: true},
		NeedPay: sql.NullFloat64{Float64: payment.NeedPay - float64(req.GetValue()), Valid: true},
		ID:      payment.ID,
	})

	if payment.NeedPay-float64(req.GetValue()) <= 0 {
		_, err = server.store.UpdateStatusOrder(ctx, db.UpdateStatusOrderParams{
			ID:     req.GetOrderId(),
			Status: "COMPLETE",
		})
		if err != nil {
			appErr := common.ErrDBWithMsg(err, "lỗi cập nhật thông tin đơn hàng")
			return &pb.PaymentItemOrderResponse{
				Code:         int32(appErr.StatusCode),
				Message:      appErr.Message,
				MessageTrans: appErr.MessageTrans,
				Log:          appErr.Log,
			}, nil
		}
	} else if 0 < payment.NeedPay-float64(req.GetValue()) && payment.NeedPay-float64(req.GetValue()) < payment.MustPaid {
		_, err = server.store.UpdateStatusOrder(ctx, db.UpdateStatusOrderParams{
			ID:     req.GetOrderId(),
			Status: "IN_PROCESS",
		})
		if err != nil {
			appErr := common.ErrDBWithMsg(err, "lỗi cập nhật thông tin đơn hàng")
			return &pb.PaymentItemOrderResponse{
				Code:         int32(appErr.StatusCode),
				Message:      appErr.Message,
				MessageTrans: appErr.MessageTrans,
				Log:          appErr.Log,
			}, nil
		}
	}

	if err != nil {
		appErr := common.ErrDBWithMsg(err, "lỗi cập nhật thông tin thanh toán")
		return &pb.PaymentItemOrderResponse{
			Code:         int32(appErr.StatusCode),
			Message:      appErr.Message,
			MessageTrans: appErr.MessageTrans,
			Log:          appErr.Log,
		}, nil
	}

	return &pb.PaymentItemOrderResponse{
		Code:    200,
		Message: "success",
	}, nil
}
