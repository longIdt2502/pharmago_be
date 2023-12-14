package gapi

import (
	"context"
	"database/sql"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListWards(ctx context.Context, req *pb.WardsRequest) (*pb.WardsResponse, error) {
	code := sql.NullString{
		String: req.GetDistrict(),
		Valid:  true,
	}
	wards, err := server.store.GetWard(ctx, code)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get provinces")
	}
	var wardData []*pb.Ward
	for _, value := range wards {
		item := mapper.WardMapper(value)
		wardData = append(wardData, item)
	}
	rsp := &pb.WardsResponse{
		Code:    200,
		Message: "success",
		Details: wardData,
		Count:   int32(len(wardData)),
	}

	return rsp, nil
}
