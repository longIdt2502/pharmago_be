package mapper

import (
	"context"
	"database/sql"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/rs/zerolog/log"
)

func VariantMapper(ctx context.Context, store *db.Store, data db.GetVariantsRow) *pb.Variant {

	media, err := store.GetMediaVariant(ctx, data.ID)
	log.Print(err)

	var totalInventory *int32
	inventory, err := store.GetInventoryVariant(ctx, sql.NullInt32{
		Int32: data.ID,
		Valid: true,
	})
	log.Print(err)
	inventory32 := int32(inventory)
	totalInventory = &inventory32

	var units []*pb.Unit
	units = append(units, &pb.Unit{
		Id:          data.UnitID,
		Name:        data.UnitName,
		Value:       0,
		SellPrice:   float32(data.UnitSellPrice),
		ImportPrice: 0,
		Weight:      float32(data.UnitWeight.Float64),
		WeightUnit:  data.UnitWeightUnit.String,
		Default:     true,
	})
	unitChangeDb, err := store.GetListUnitChange(ctx, data.UnitID)
	log.Print(err)
	for _, value := range unitChangeDb {
		units = append(units, &pb.Unit{
			Id:          value.ID,
			Name:        value.Name,
			Value:       int32(value.Value),
			SellPrice:   float32(value.SellPrice),
			ImportPrice: 0,
			Weight:      0,
			WeightUnit:  "",
			Default:     false,
		})
	}

	return &pb.Variant{
		Id:              data.ID,
		Code:            data.Code,
		Name:            data.Name,
		Barcode:         data.Name,
		DecisionNumber:  data.DecisionNumber,
		RegisterNumber:  data.RegisterNumber,
		Longevity:       data.Longevity,
		Vat:             float32(data.Vat),
		Product:         data.Product,
		Media:           media.MediaUrl,
		QuantityInStock: totalInventory,
		Units:           units,
		PriceSell:       float32(data.PlPriceSell),
		PriceImport:     float32(data.PlPriceImport),
	}
}

func VariantPreviewMapper(ctx context.Context, store *db.Store, data db.Variant) *pb.Variant {

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
