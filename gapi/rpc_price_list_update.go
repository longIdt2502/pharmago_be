package gapi

import (
	"context"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) UpdatePriceList(ctx context.Context, req *pb.UpdatePriceListRequest) (*pb.UpdatePriceListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.UpdatePriceList(ctx, db.UpdatePriceListParams{
		PriceImport: float64(req.GetPriceImport()),
		PriceSell:   float64(req.GetPriceSell()),
		ID:          req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update price list: %e", err)
	}

	priceList, err := server.store.DetailPriceList(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get detail price list: %e", err)
	}

	priceListPb := mapper.PriceListDetailMapper(priceList)

	return &pb.UpdatePriceListResponse{
		Code:    200,
		Message: "success",
		Details: priceListPb,
	}, nil
}
