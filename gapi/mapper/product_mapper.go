package mapper

import (
	"context"
	"database/sql"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/rs/zerolog/log"
)

func ProductMapper(ctx context.Context, store *db.Store, data db.Product) *pb.Product {
	image, err := store.GetProductMedia(ctx, sql.NullInt64{
		Int64: data.ID,
		Valid: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get image product")
	}

	var name []string
	for _, value := range image {
		name = append(name, value.MediaUrl)
	}

	return &pb.Product{
		Id:           int32(data.ID),
		Name:         data.Name,
		Code:         data.Code,
		Category:     config.NewInt32Optional(data.ProductCategory),
		Type:         config.NewInt32Optional(data.Type),
		TaDuoc:       config.NewStringOptional(data.Taduoc),
		NongDo:       config.NewStringOptional(data.Nongdo),
		LieuDung:     data.Lieudung,
		ChiDinh:      data.Chidinh,
		ChongChiDinh: config.NewStringOptional(data.Chongchidinh),
		CongDung:     data.Congdung,
		TacDungPhu:   data.Tacdungphu,
		ThanTrong:    data.Thantrong,
		TuongTac:     config.NewStringOptional(data.Tuongtac),
		BaoQuan:      data.Baoquan,
		DongGoi:      data.Donggoi,
		NoiSx:        data.Noisx,
		CongTySx:     data.Congtysx,
		CongTyDk:     data.Congtydk,
		Image:        name,
	}
}
