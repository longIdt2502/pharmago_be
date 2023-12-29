package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TicketMapper(data db.GetListTicketRow) *pb.TicketPreview {
	note := ""
	if data.Note.Valid {
		note = data.Note.String
	}

	return &pb.TicketPreview{
		Id:   data.ID,
		Code: data.Code,
		Type: &pb.SimpleData{
			Id:   data.TtID,
			Name: data.TtTitle,
			Code: data.TtCode,
		},
		Status: &pb.SimpleData{
			Id:   data.TsID,
			Name: data.TsTitle,
			Code: data.TsCode,
		},
		Note:           note,
		Qr:             data.QrUrl,
		TotalItems:     data.TotalProducts,
		TotalItemsType: 0,
		TotalPrice:     float32(data.TotalPrice),
		WarehouseName:  data.WName,
		UserCreated:    data.AFullName,
		CreatedAt:      timestamppb.New(data.CreatedAt),
	}
}
