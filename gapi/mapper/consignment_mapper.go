package mapper

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConsignmentMapper(ctx context.Context, store *db.Store, data db.Consignment) *pb.Consignment {

	variant, _ := store.GetVariantById(ctx, data.Variant.Int32)
	variantPb := VariantPreviewMapper(ctx, store, variant)

	return &pb.Consignment{
		Id:          data.ID,
		Code:        data.Code,
		Quantity:    data.Quantity,
		Inventory:   data.Inventory,
		Variant:     variantPb,
		ExpiredAt:   timestamppb.New(data.ExpiredAt),
		ProducedAt:  timestamppb.New(data.ProductedAt),
		IsAvailable: data.IsAvailable,
	}
}
