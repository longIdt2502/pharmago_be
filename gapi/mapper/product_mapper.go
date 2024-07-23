package mapper

import (
	"context"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/rs/zerolog/log"
)

func ProductMapper(ctx context.Context, store *db.Store, data db.Product) *pb.Product {
	image, err := store.GetProductMedia(ctx, data.ID)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get image product")
	}

	var name []string
	for _, value := range image {
		name = append(name, value.MediaUrl)
	}

	unit, _ := store.DetailUnit(ctx, data.Unit)

	return &pb.Product{
		Id:           data.ID,
		Name:         data.Name,
		Code:         data.Code,
		Category:     config.NewInt32Optional(data.ProductCategory),
		Type:         config.NewInt32Optional(data.Type),
		TaDuoc:       config.NewStringOptional(data.TaDuoc),
		NongDo:       config.NewStringOptional(data.NongDo),
		LieuDung:     data.LieuDung.String,
		ChiDinh:      data.ChiDinh.String,
		ChongChiDinh: config.NewStringOptional(data.ChongChiDinh),
		CongDung:     data.CongDung.String,
		TacDungPhu:   data.TacDungPhu.String,
		ThanTrong:    data.ThanTrong.String,
		TuongTac:     config.NewStringOptional(data.TuongTac),
		BaoQuan:      data.BaoQuan.String,
		DongGoi:      data.DongGoi.String,
		CongTySx:     data.CongTySx.Int32,
		CongTyDk:     data.CongTyDk.Int32,
		Image:        name,
		Unit: &pb.Unit{
			Id:          unit.ID,
			Name:        unit.Name,
			SellPrice:   float32(unit.SellPrice),
			ImportPrice: float32(unit.ImportPrice),
			Default:     true,
		},
	}
}
