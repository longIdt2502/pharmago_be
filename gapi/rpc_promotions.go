package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/thoas/go-funk"
)

func (server *ServerGRPC) PromotionByProduct(ctx context.Context, req *pb.PromotionByProductRequest) (*pb.PromotionByProductResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	promotions, err := server.store.GetByVariantOrService(ctx, db.GetByVariantOrServiceParams{
		Variant: sql.NullInt32{
			Int32: req.GetVariant(),
			Valid: req.Variant != nil,
		},
		Service: sql.NullInt32{
			Int32: req.GetService(),
			Valid: req.Service != nil,
		},
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.PromotionByProductResponse{
			Code:    int32(errApp.StatusCode),
			RootErr: errApp.RootError().Error(),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	var promotionsPb []*pb.Promotion
	for _, promotion := range promotions {
		index := funk.IndexOf(promotionsPb, func(value *pb.Promotion) bool {
			return value.Id == promotion.ID_2.UUID.String()
		})
		if index == -1 {
			promotionsPb = append(promotionsPb, mapper.PromotionMapperGetByProductRow(
				promotion,
			))
		} else {
			promotionsPb[index].Items = append(promotionsPb[index].Items, mapper.PromotionItemMapperGetByProductRow(promotion))
		}
	}

	return &pb.PromotionByProductResponse{
		Code:    200,
		Message: "success",
		Details: promotionsPb,
	}, nil
}

func (server *ServerGRPC) PromotionCheck(ctx context.Context, req *pb.PromotionCheckRequest) (*pb.PromotionCheckResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	promotions, err := server.store.GetPromotionByPriceOrder(ctx, db.GetPromotionByPriceOrderParams{
		Price:   float64(req.GetTotalPrice()),
		Company: req.GetCompany(),
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.PromotionCheckResponse{
			Code:    int32(errApp.StatusCode),
			RootErr: errApp.RootError().Error(),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	var promotionsPb []*pb.Promotion
	for _, item := range promotions {
		promotionsPb = append(promotionsPb, mapper.PromotionMapper(item))
	}

	return &pb.PromotionCheckResponse{
		Code:    200,
		Message: "success",
		Details: promotionsPb,
	}, nil
}

func (server *ServerGRPC) PromotionCreate(ctx context.Context, req *pb.PromotionCreateRequest) (*pb.PromotionCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("PROMO-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))

	promotion, err := server.store.CreatePromotion(ctx, db.CreatePromotionParams{
		ID:   uuid.New(),
		Code: code,
		Type: sql.NullString{
			String: req.GetType(),
			Valid:  true,
		},
		Title: req.GetTitle(),
		ConditionsText: sql.NullString{
			String: req.ConditionsText,
			Valid:  true,
		},
		ConditionsPointCustomer: sql.NullInt32{
			Int32: req.GetConditionsPointCustomer(),
			Valid: true,
		},
		MinValue:          float64(req.MinValue),
		IsDiscountPercent: req.IsDiscountPercent,
		ValueDiscount:     float64(req.ValueDiscount),
		MaxDiscount:       float64(req.MaxDiscount),
		TimeApply: sql.NullInt32{
			Int32: req.GetTimeApply(),
			Valid: req.TimeApply != nil,
		},
		DateStart: sql.NullTime{
			Time:  time.Unix(req.DateStart.GetSeconds(), 0),
			Valid: req.DateStart.IsValid(),
		},
		DateEnd: sql.NullTime{
			Time:  time.Unix(req.DateEnd.GetSeconds(), 0),
			Valid: req.DateEnd.IsValid(),
		},
		ApplyMultipleTimes:  req.ApplyMultipleTimes,
		ApplySimultaneously: req.ApplySimultaneously,
		Status:              req.Status,
		Company:             req.Company,
		UserCreated:         tokenPayload.UserID,
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.PromotionCreateResponse{
			Code:    int32(errApp.StatusCode),
			RootErr: errApp.RootErr.Error(),
			Message: errApp.Message,
			Log:     errApp.Log,
		}, nil
	}

	if req.Type == "GIFT" {
		for _, item := range req.GetItems() {
			_, err = server.store.CreatePromotionItem(ctx, db.CreatePromotionItemParams{
				ID:         uuid.New(),
				MinBuy:     item.GetMinBuy(),
				AmountGift: item.GetAmountGift(),
				Promotions: promotion.ID,
				Variant: sql.NullInt32{
					Int32: item.GetVariant(),
					Valid: item.Variant != nil,
				},
				Service: sql.NullInt32{
					Int32: item.GetService(),
					Valid: item.Service != nil,
				},
				ApplicableVariant: sql.NullInt32{
					Int32: item.GetApplicableVariant(),
					Valid: item.ApplicableVariant != nil,
				},
				ApplicableService: sql.NullInt32{
					Int32: item.GetApplicableService(),
					Valid: item.ApplicableService != nil,
				},
			})
			if err != nil {
				errApp := common.ErrDB(err)
				return &pb.PromotionCreateResponse{
					Code:    int32(errApp.StatusCode),
					RootErr: errApp.RootErr.Error(),
					Message: errApp.Message,
					Log:     errApp.Log,
				}, nil
			}
		}
	}

	return &pb.PromotionCreateResponse{
		Code:    200,
		Message: "success",
		Details: promotion.ID.String(),
	}, nil
}
