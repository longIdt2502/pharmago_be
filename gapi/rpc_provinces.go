package gapi

import (
	"context"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListProvinces(ctx context.Context, req *pb.ProvincesRequest) (*pb.ProvincesResponse, error) {
	province, err := server.store.GetProvince(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get provinces")
	}
	var provinceData []*pb.Province
	for _, value := range province {
		item := mapper.ProvinceMapper(value)
		provinceData = append(provinceData, item)
	}

	rsp := &pb.ProvincesResponse{
		Code:    200,
		Message: "success",
		Details: provinceData,
		Count:   int32(len(provinceData)),
	}

	return rsp, nil
}
