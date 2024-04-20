package gapi

import (
	"context"
	"database/sql"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
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

	return &pb.ListNotificationResponse{
		Code:    200,
		Message: "success",
		Details: notificationsPb,
	}, nil
}
