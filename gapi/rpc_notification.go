package gapi

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/woker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListNotification(ctx context.Context, req *pb.ListNotificationRequest) (*pb.ListNotificationResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	notifications, err := server.store.ListNotification(ctx, db.ListNotificationParams{
		Company: int32(req.GetCompany()),
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
		return nil, status.Errorf(codes.Internal, "failed to get record noti: %e", err)
	}

	var notificationsPb []*pb.Notification
	for _, item := range notifications {
		notificationsPb = append(notificationsPb, mapper.NotificationMapper(item))
	}

	count, err := server.store.CountNotification(ctx, sql.NullInt32{
		Int32: int32(req.GetCompany()),
		Valid: true,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get count notification: %e", err)
	}

	countNotSeen := 0
	countSeen := 0

	for _, item := range count {
		if item.IsRead {
			countSeen = int(item.Count)
		} else {
			countNotSeen = int(item.Count)
		}
	}

	return &pb.ListNotificationResponse{
		Code:         200,
		Message:      "success",
		Details:      notificationsPb,
		CountNotSeen: int32(countNotSeen),
		Count:        int32(countSeen + countNotSeen),
	}, nil
}

func (server *ServerGRPC) DetailNotification(ctx context.Context, req *pb.DetailNotificationRequest) (*pb.DetailNotificationResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	noti, err := server.store.DetailNotification(ctx, int32(req.GetId()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "noti not exitst")
		}
		return nil, status.Errorf(codes.Internal, "failed to get record noti")
	}

	notiPb := mapper.NotificationMapper(noti)

	if !noti.IsRead {
		server.store.UpdateNotification(ctx, db.UpdateNotificationParams{
			IsRead: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
			ID: noti.ID,
		})
	}

	var orderPb *pb.Order
	var servicePb *pb.Service
	if noti.Data.Valid {
		var dataNoti woker.DataNoti
		if err = json.Unmarshal([]byte(noti.Data.String), &dataNoti); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal data noti")
		}

		if dataNoti.Order != nil {
			order, err := server.store.DetailOrder(ctx, db.DetailOrderParams{
				ID: sql.NullInt32{
					Int32: *dataNoti.Order,
					Valid: true,
				},
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to get order")
			}
			orderPb = mapper.OrderDetailMapper(ctx, server.store, order)
		}

		if dataNoti.Service != nil {
			service, err := server.store.DetailService(ctx, *dataNoti.Service)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to get service")
			}
			servicePb = mapper.ServiceMapper(service)
		}
	}

	return &pb.DetailNotificationResponse{
		Code:    200,
		Message: "success",
		Details: notiPb,
		Order:   orderPb,
		Service: servicePb,
	}, nil
}

func (server *ServerGRPC) SeenAllNoti(ctx context.Context, req *pb.SeenAllNotiRequest) (*pb.SeenAllNotiResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.SeenAllNotification(ctx, sql.NullInt32{
		Int32: int32(req.GetCompany()),
		Valid: true,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to seen all notification")
	}

	return &pb.SeenAllNotiResponse{
		Code:    200,
		Message: "success",
	}, nil
}
