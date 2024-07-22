package mapper

import (
	"context"
	"database/sql"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func VariantMapper(ctx context.Context, store *db.Store, data db.GetVariantsRow) *pb.Variant {

	media, _ := store.GetMediaVariant(ctx, data.ID)

	var totalInventory *int32
	inventory, _ := store.GetInventoryVariant(ctx, sql.NullInt32{
		Int32: data.ID,
		Valid: true,
	})
	inventory32 := inventory
	totalInventory = &inventory32

	var units []*pb.Unit
	units = append(units, &pb.Unit{
		Id:          data.UnitID,
		Name:        data.UnitName,
		Value:       1,
		SellPrice:   float32(data.UnitSellPrice),
		ImportPrice: 0,
		Weight:      float32(data.UnitWeight.Float64),
		WeightUnit:  data.UnitWeightUnit.String,
		Default:     true,
	})
	unitChangeDb, _ := store.GetListUnitChange(ctx, data.UnitID)
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

	var vat *float32
	if data.Vat.Valid {
		vat32 := float32(data.Vat.Float64)
		vat = &vat32
	}

	return &pb.Variant{
		Id:               data.ID,
		Code:             data.Code,
		Name:             data.Name,
		Barcode:          &data.Barcode.String,
		DecisionNumber:   &data.DecisionNumber.String,
		RegisterNumber:   &data.RegisterNumber.String,
		Longevity:        &data.Longevity.String,
		Product:          data.Product,
		Media:            media.MediaUrl,
		QuantityInStock:  totalInventory,
		Units:            units,
		Vat:              vat,
		PriceSell:        float32(data.PlPriceSell.Float64),
		PriceImport:      float32(data.PlPriceImport.Float64),
		InitialInventory: data.InitialInventory,
		RealInventory:    data.RealInventory,
	}
}

func VariantPreviewMapper(ctx context.Context, store *db.Store, data db.Variant) *pb.Variant {

	media, _ := store.GetMediaVariant(ctx, data.ID)

	var vat *float32
	if data.Vat.Valid {
		vat32 := float32(data.Vat.Float64)
		vat = &vat32
	}

	return &pb.Variant{
		Id:               data.ID,
		Code:             data.Code,
		Name:             data.Name,
		Barcode:          &data.Barcode.String,
		DecisionNumber:   &data.DecisionNumber.String,
		RegisterNumber:   &data.RegisterNumber.String,
		Longevity:        &data.Longevity.String,
		Vat:              vat,
		Product:          data.Product,
		Media:            media.MediaUrl,
		InitialInventory: data.InitialInventory,
		RealInventory:    data.RealInventory,
	}
}

func VariantCustomerBuyMapper(ctx context.Context, store *db.Store, data db.VariantsCustomerBuyRow) *pb.Variant {

	media, _ := store.GetMediaVariant(ctx, data.ID)

	var vat *float32
	if data.Vat.Valid {
		vat32 := float32(data.Vat.Float64)
		vat = &vat32
	}

	return &pb.Variant{
		Id:               data.ID,
		Code:             data.Code,
		Name:             data.Name,
		Barcode:          &data.Barcode.String,
		DecisionNumber:   &data.DecisionNumber.String,
		RegisterNumber:   &data.RegisterNumber.String,
		Longevity:        &data.Longevity.String,
		Vat:              vat,
		Product:          data.Product,
		Media:            media.MediaUrl,
		InitialInventory: data.InitialInventory,
		RealInventory:    data.RealInventory,
		QuantityBuy:      int32(data.QuantityBuy),
		Units: []*pb.Unit{
			{
				Id:          data.ID_2,
				Name:        data.Name_2,
				SellPrice:   float32(data.SellPrice),
				ImportPrice: float32(data.ImportPrice),
				Default:     true,
			},
		},
	}
}
