package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OrderPreviewMapper(order db.ListOrderRow) *pb.OrderPreview {

	return &pb.OrderPreview{
		Id:         order.ID,
		Code:       order.Code,
		TotalPrice: float32(order.TotalPrice),
		Status: &pb.SimpleData{
			Id:   order.OsID,
			Name: order.OsTitle,
			Code: order.Status.String,
		},
		Description:  order.Description.String,
		CustomerName: order.CFullName,
		UserCreated:  order.AFullName,
		CreatedAt:    timestamppb.New(order.CreatedAt),
	}
}
