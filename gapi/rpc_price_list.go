package gapi

import (
	"context"
	"database/sql"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) GetPriceList(ctx context.Context, req *pb.PriceListRequest) (*pb.PriceListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	priceList, err := server.store.GetPriceLists(ctx, db.GetPriceListsParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: req.Company != nil,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		MinPriceImport: sql.NullFloat64{
			Float64: float64(req.GetMinPriceImport()),
			Valid:   req.MinPriceImport != nil,
		},
		MaxPriceImport: sql.NullFloat64{
			Float64: float64(req.GetMaxPriceImport()),
			Valid:   req.MaxPriceImport != nil,
		},
		MinPriceSell: sql.NullFloat64{
			Float64: float64(req.GetMinPriceSell()),
			Valid:   req.MinPriceSell != nil,
		},
		MaxPriceSell: sql.NullFloat64{
			Float64: float64(req.GetMaxPriceSell()),
			Valid:   req.MaxPriceSell != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get price list: %e", err)
	}

	var priceListPb []*pb.PriceList
	for _, value := range priceList {
		dataPb := mapper.PriceListMapper(value)
		priceListPb = append(priceListPb, dataPb)
	}

	return &pb.PriceListResponse{
		Code:    200,
		Message: "success",
		Details: priceListPb,
	}, nil
}
