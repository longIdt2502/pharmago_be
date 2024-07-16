package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PromotionMapper(data db.Promotion) *pb.Promotion {
	return &pb.Promotion{
		Id:                      data.ID.String(),
		Code:                    data.Code,
		Type:                    &pb.SimpleData{},
		Title:                   data.Title.String,
		ConditionsText:          data.ConditionsText.String,
		ConditionsPointCustomer: data.ConditionsPointCustomer.Int32,
		MinValue:                float32(data.MinValue.Float64),
		IsDiscountPercent:       data.IsDiscountPercent,
		ValueDiscount:           float32(data.ValueDiscount),
		MaxDiscount:             float32(data.MaxDiscount.Float64),
		TimeApply:               &data.TimeApply.Int32,
		DateStart:               timestamppb.New(data.DateStart.Time),
		DateEnd:                 timestamppb.New(data.DateEnd.Time),
		ApplyMultipleTimes:      data.ApplyMultipleTimes,
		ApplySimultaneously:     data.ApplySimultaneously,
		Status:                  data.Status,
		UserCreated:             &pb.Account{},
		UserUpdated:             &pb.Account{},
		CreatedAt:               timestamppb.New(data.CreatedAt),
		UpdatedAt:               timestamppb.New(data.UpdatedAt.Time),
		Items:                   []*pb.PromotionItem{},
		Company:                 data.Company.Int32,
	}
}

func PromotionMapperGetByProductRow(data db.GetByVariantOrServiceRow) *pb.Promotion {
	return &pb.Promotion{
		Id:   data.ID_2.UUID.String(),
		Code: data.Code.String,
		Type: &pb.SimpleData{
			Name: data.Title_2.String,
			Code: data.Code_3.String,
		},
		Title:                   data.Title.String,
		ConditionsText:          data.ConditionsText.String,
		ConditionsPointCustomer: data.ConditionsPointCustomer.Int32,
		MinValue:                float32(data.MinValue.Float64),
		IsDiscountPercent:       data.IsDiscountPercent.Bool,
		ValueDiscount:           float32(data.ValueDiscount.Float64),
		MaxDiscount:             float32(data.MaxDiscount.Float64),
		TimeApply:               &data.TimeApply.Int32,
		DateStart:               timestamppb.New(data.DateStart.Time),
		DateEnd:                 timestamppb.New(data.DateEnd.Time),
		ApplyMultipleTimes:      data.ApplyMultipleTimes.Bool,
		ApplySimultaneously:     data.ApplySimultaneously.Bool,
		Status:                  data.Status.Bool,
		UserCreated:             &pb.Account{},
		UserUpdated:             &pb.Account{},
		CreatedAt:               timestamppb.New(data.CreatedAt.Time),
		UpdatedAt:               timestamppb.New(data.UpdatedAt.Time),
		Items: []*pb.PromotionItem{
			PromotionItemMapperGetByProductRow(data),
		},
		Company: data.Company.Int32,
	}
}

func PromotionItemMapperGetByProductRow(data db.GetByVariantOrServiceRow) *pb.PromotionItem {

	var service *pb.Service
	var variant *pb.Variant

	if data.VName.Valid {
		variant = &pb.Variant{
			Name: data.VName.String,
			Code: data.VCode.String,
		}
	}

	if data.SName.Valid {
		service = &pb.Service{
			Title: data.SName.String,
			Code:  data.SCode.String,
		}
	}

	return &pb.PromotionItem{
		Id:         data.ID.String(),
		MinBuy:     data.MinBuy.Int32,
		AmountGift: data.AmountGift.Int32,
		Promotions: data.ID_2.UUID.String(),
		Variant:    variant,
		Service:    service,
	}
}
