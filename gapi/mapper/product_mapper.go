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

	return &pb.Product{
		Id:           data.ID,
		Name:         data.Name,
		Code:         data.Code,
		Category:     config.NewInt32Optional(data.ProductCategory),
		Type:         config.NewInt32Optional(data.Type),
		TaDuoc:       config.NewStringOptional(data.TaDuoc),
		NongDo:       config.NewStringOptional(data.NongDo),
		LieuDung:     data.LieuDung,
		ChiDinh:      data.ChiDinh,
		ChongChiDinh: config.NewStringOptional(data.ChongChiDinh),
		CongDung:     data.CongDung,
		TacDungPhu:   data.TacDungPhu,
		ThanTrong:    data.ThanTrong,
		TuongTac:     config.NewStringOptional(data.TuongTac),
		BaoQuan:      data.BaoQuan,
		DongGoi:      data.DongGoi,
		CongTySx:     data.CongTySx,
		CongTyDk:     data.CongTyDk,
		Image:        name,
	}
}
