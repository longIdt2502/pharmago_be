package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PriceListMapper(data db.GetPriceListsRow) *pb.PriceList {

	var media *string
	if data.VariantMedia.Valid {
		media = &data.VariantMedia.String
	}

	return &pb.PriceList{
		Id:              int32(data.ID),
		VariantCode:     data.VariantCode,
		VariantName:     data.VariantName,
		VariantMedia:    media,
		PriceImport:     float32(data.PriceImport),
		PriceSell:       float32(data.PriceSell),
		Unit:            int32(data.Unit),
		UnitName:        data.Name_2,
		UserCreated:     int32(data.UserCreated),
		UserCreatedName: data.UserCreatedName,
		CreatedAt:       timestamppb.New(data.CreatedAt),
	}
}
