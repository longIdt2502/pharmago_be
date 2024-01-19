package mapper

import (
	"context"
	"database/sql"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConsignmentMapper(ctx context.Context, store *db.Store, data db.Consignment, company int32) *pb.Consignment {

	variant, _ := store.GetVariants(ctx, db.GetVariantsParams{
		ID: sql.NullInt32{
			Int32: data.Variant.Int32,
			Valid: true,
		},
		Company: company,
	})

	var variantPb *pb.Variant
	variantPb = VariantMapper(ctx, store, variant[0])

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
