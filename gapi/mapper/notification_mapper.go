package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NotificationMapper(item db.Notification) *pb.Notification {
	return &pb.Notification{
		Id:        item.ID,
		Type:      item.Type,
		Topic:     item.Topic,
		Title:     item.Title,
		Content:   item.Content,
		IsRead:    item.IsRead,
		Data:      item.Data.String,
		Company:   item.Company.Int32,
		CreatedAt: timestamppb.New(item.CreatedAt),
	}
}
