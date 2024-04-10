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
	"github.com/skip2/go-qrcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) OrderCreate(ctx context.Context, req *pb.OrderCreateRequest) (*pb.OrderCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	var addressOrder int32
	if req.Order.Customer != nil {
		customer, err := server.store.GetCustomer(ctx, req.Order.GetCustomer())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "customer not exists")
			}
			return nil, status.Errorf(codes.Internal, "failed to get customer")
		}
		addressOrder = customer.ID
	}
	// else {
	// 	// TODO:
	// }

	warehouse, err := server.store.GetWarehouse(ctx, req.GetWarehouse())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "warehouse not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get warehouse")
	}

	paymentCode := fmt.Sprintf("PM-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	payment, err := server.store.CreatePayment(ctx, db.CreatePaymentParams{
		Code:     paymentCode,
		MustPaid: float64(req.Payment.GetMustPaid()),
		HadPaid:  float64(req.Payment.GetHadPaid()),
		NeedPay:  float64(req.Payment.GetNeedPay()),
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to record payment: %e", err)
	}

	for _, value := range req.PaymentItems {
		_, err = server.store.CreatePaymentItem(ctx, db.CreatePaymentItemParams{
			Type:    value.GetType(),
			Value:   float64(value.GetValue()),
			IsPaid:  value.GetIsPaid(),
			Payment: payment.ID,
			ExtraNote: sql.NullString{
				String: value.GetExtraNote(),
				Valid:  value.ExtraNote != nil,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to record payment item: %e", err)
		}
	}

	ticketType, _ := server.store.GetTicketType(ctx, db.GetTicketTypeParams{
		ID: sql.NullInt32{},
		Code: sql.NullString{
			String: "EXPORT",
			Valid:  true,
		},
	})
	ticketStatus, _ := server.store.GetTicketStatus(ctx, db.GetTicketStatusParams{
		ID: sql.NullInt32{},
		Code: sql.NullString{
			String: "NEW",
			Valid:  true,
		},
	})

	codeTicket := fmt.Sprintf("TICKET-%s", utils.RandomString(6))
	var png []byte
	png, err = qrcode.Encode(codeTicket, qrcode.Medium, 256)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create qr code ticket: %e", err)
	}
	file, _ := utils.NewFileFromImage(png)
	_, err = server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save qr code: %e", err)
	}
	urlQr, _ := server.b2Bucket.FileURL(file.Name)

	qr, _ := server.store.CreateMedia(ctx, urlQr)

	ticket, err := server.store.CreateTicket(ctx, db.CreateTicketParams{
		Code: codeTicket,
		Type: sql.NullInt32{
			Int32: ticketType.ID,
			Valid: true,
		},
		Status: sql.NullInt32{
			Int32: ticketStatus.ID,
			Valid: true,
		},
		Note: sql.NullString{
			String: "Phiếu xuất hàng cho đơn hàng",
			Valid:  true,
		},
		Qr: sql.NullInt32{
			Int32: qr.ID,
			Valid: true,
		},
		ExportTo: sql.NullInt32{
			Int32: addressOrder,
			Valid: req.Order.Customer != nil,
		},
		ImportFrom: sql.NullInt32{
			Int32: warehouse.Address.Int32,
			Valid: true,
		},
		TotalPrice:  float64(req.Order.GetTotalPrice()),
		Warehouse:   req.GetWarehouse(),
		UserCreated: tokenPayload.UserID,
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record ticket: %e", err)
	}

	var orderCode string
	if req.Order.Code == nil {
		orderCode = fmt.Sprintf("ORDER-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	} else {
		orderCode = req.Order.GetCode()
	}

	var pngOrder []byte
	pngOrder, err = qrcode.Encode(orderCode, qrcode.Medium, 256)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create qr code order: %e", err)
	}
	fileOrder, _ := utils.NewFileFromImage(pngOrder)
	_, err = server.b2Bucket.UploadFile(fileOrder.Name, fileOrder.Meta, fileOrder.File)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save qr code: %e", err)
	}
	urlQrOrder, _ := server.b2Bucket.FileURL(fileOrder.Name)

	qrOrder, _ := server.store.CreateMedia(ctx, urlQrOrder)

	order, err := server.store.CreateOrder(ctx, db.CreateOrderParams{
		Code:       orderCode,
		TotalPrice: float64(req.Order.GetTotalPrice()),
		Description: sql.NullString{
			String: req.Order.GetDescription(),
			Valid:  req.Order.Description != nil,
		},
		Vat:          float64(req.Order.GetVat()),
		Discount:     req.Order.GetDiscount(),
		ServicePrice: float64(req.Order.GetServicePrice()),
		MustPaid:     float64(req.Order.GetMustPaid()),
		Customer: sql.NullInt32{
			Int32: req.Order.GetCustomer(),
			Valid: req.Order.Customer != nil,
		},
		Address: sql.NullInt32{
			Int32: addressOrder,
			Valid: req.Order.Customer != nil,
		},
		Status: sql.NullString{
			String: req.Order.GetStatus(),
			Valid:  true,
		},
		Type: sql.NullString{
			String: req.Order.GetType(),
			Valid:  true,
		},
		Ticket: sql.NullInt32{
			Int32: ticket.ID,
			Valid: true,
		},
		Qr: sql.NullInt32{
			Int32: qrOrder.ID,
			Valid: true,
		},
		Company: req.Order.GetCompany(),
		Payment: payment.ID,
		UserCreated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to record order: %e", err)
	}

	for _, value := range req.GetOrderItems() {
		var consignmentLog db.ConsignmentLog
		var consignment db.Consignment
		if value.Consignment != nil {
			consignment, err = server.store.GetConsignment(ctx, db.GetConsignmentParams{
				ID: value.GetConsignment(),
				Variant: sql.NullInt32{
					Int32: value.GetVariant(),
					Valid: true,
				},
			})
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil, status.Errorf(codes.NotFound, "consignment not exists: %e", err)
				}
				return nil, status.Errorf(codes.NotFound, "consignment error: %e", err)
			}
		} else {
			consignment, err = server.store.SuggestConsignmentForVariant(ctx, db.SuggestConsignmentForVariantParams{
				Variant: sql.NullInt32{
					Int32: value.GetVariant(),
					Valid: true,
				},
				Inventory: value.GetValue(),
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to get consignment: %e", err)
			}
		}

		if consignment.Inventory < value.GetValue() {
			return nil, status.Errorf(codes.Internal, "inventory not enough")
		}

		var amount int32
		if req.Order.Type == "SELL" {
			amount = -value.GetValue()
		} else {
			amount = value.GetValue()
		}
		_, _ = server.store.UpdateConsignment(ctx, db.UpdateConsignmentParams{
			Amount: amount,
			ID:     consignment.ID,
		})

		consignmentLog, err = server.store.CreateConsignmentLog(ctx, db.CreateConsignmentLogParams{
			Consignment:  consignment.ID,
			Inventory:    consignment.Inventory,
			AmountChange: amount,
			UserCreated: sql.NullInt32{
				Int32: tokenPayload.UserID,
				Valid: true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to record consignment log: %e", err)
		}

		_, err = server.store.CreateOrderItem(ctx, db.CreateOrderItemParams{
			Order:      order.ID,
			Variant:    value.GetVariant(),
			Value:      value.GetValue(),
			TotalPrice: float64(value.GetTotalPrice()),
			Consignment: sql.NullInt32{
				Int32: consignment.ID,
				Valid: true,
			},
			ConsignmentLog: sql.NullInt32{
				Int32: consignmentLog.ID,
				Valid: true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to record order item: %e", err)
		}
	}

	return &pb.OrderCreateResponse{
		Code:    200,
		Message: "success",
		Details: order.ID,
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
		Company: sql.NullInt32{
			Int32: req.Company,
			Valid: true,
		},
		Status: sql.NullString{
			String: req.GetStatus(),
			Valid:  req.Status != nil,
		},
		Warehouse: sql.NullInt32{
			Int32: req.GetWarehouse(),
			Valid: req.Warehouse != nil,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		CreatedStart: sql.NullTime{
			Time:  time.Unix(req.CreatedStart.GetSeconds(), 0),
			Valid: req.CreatedStart != nil,
		},
		CreatedEnd: sql.NullTime{
			Time:  time.Unix(req.CreatedEnd.GetSeconds(), 0),
			Valid: req.CreatedEnd != nil,
		},
		UpdatedStart: sql.NullTime{
			Time:  time.Unix(req.UpdatedStart.GetSeconds(), 0),
			Valid: req.UpdatedStart != nil,
		},
		UpdatedEnd: sql.NullTime{
			Time:  time.Unix(req.UpdatedEnd.GetSeconds(), 0),
			Valid: req.UpdatedEnd != nil,
		},
		OrderBy: sql.NullString{
			String: req.GetOrderBy(),
			Valid:  req.OrderBy != nil,
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
		return nil, status.Errorf(codes.Internal, "failed to get orders: %e", err)
	}

	var ordersPb []*pb.OrderPreview

	for _, value := range orders {
		dataPb := mapper.OrderPreviewMapper(value)
		ordersPb = append(ordersPb, dataPb)
	}

	draftCount, _ := server.store.CountOrderByStatus(ctx, sql.NullString{
		String: "DRAFT",
		Valid:  true,
	})

	inProcessCount, _ := server.store.CountOrderByStatus(ctx, sql.NullString{
		String: "IN_PROCESS",
		Valid:  true,
	})

	completeCount, _ := server.store.CountOrderByStatus(ctx, sql.NullString{
		String: "COMPLETE",
		Valid:  true,
	})

	cancelCount, _ := server.store.CountOrderByStatus(ctx, sql.NullString{
		String: "CANCEL",
		Valid:  true,
	})

	return &pb.OrderListResponse{
		Code:    200,
		Message: "success",
		Details: ordersPb,
		Count: &pb.OrderListResponseCount{
			Draft:     int32(draftCount),
			InProcess: int32(inProcessCount),
			Complete:  int32(completeCount),
			Cancel:    int32(cancelCount),
		},
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
		return nil, status.Errorf(codes.Internal, "failed to get order: %e", err)
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
