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
	// if err != nil {
	// 	return nil, err
	// }

	return &pb.OrderCreateResponse{
		Code:    200,
		Message: fmt.Sprintf("success: %e", err),
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
