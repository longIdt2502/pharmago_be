package gapi

import (
	"context"
	"database/sql"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListDistricts(ctx context.Context, req *pb.DistrictsRequest) (*pb.DistrictsResponse, error) {
	code := sql.NullString{
		String: req.GetProvince(),
		Valid:  true,
	}
	district, err := server.store.GetDistrict(ctx, code)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get provinces")
	}
	var districtData []*pb.District
	for _, value := range district {
		item := mapper.DistrictMapper(value)
		districtData = append(districtData, item)
	}
	rsp := &pb.DistrictsResponse{
		Code:    200,
		Message: "success",
		Details: districtData,
		Count:   int32(len(districtData)),
	}

	return rsp, nil
}
