package gapi

import (
	"context"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) DetailPriceList(ctx context.Context, req *pb.DetailPriceListRequest) (*pb.DetailPriceListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	priceList, err := server.store.DetailPriceList(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get detail price list: ", err)
	}

	priceListPb := mapper.PriceListDetailMapper(priceList)

	return &pb.DetailPriceListResponse{
		Code:    200,
		Message: "success",
		Details: priceListPb,
	}, nil
}
