package mapper

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func VariantMapper(ctx context.Context, store *db.Store, data db.Variant) *pb.Variant {

	media, _ := store.GetMediaVariant(ctx, data.ID)

	return &pb.Variant{
		Id:             data.ID,
		Code:           data.Code,
		Name:           data.Name,
		Barcode:        data.Name,
		DecisionNumber: data.DecisionNumber,
		RegisterNumber: data.RegisterNumber,
		Longevity:      data.Longevity,
		Vat:            float32(data.Vat),
		Product:        data.Product,
		Media:          media.MediaUrl,
	}
}
