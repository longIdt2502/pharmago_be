package gapi

import (
	"context"
	"database/sql"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ProductTypeList(ctx context.Context, req *pb.ProductTypeListRequest) (*pb.ProductTypeListResponse, error) {
	productType, err := server.store.GetListProductType(ctx, db.GetListProductTypeParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: req.Company != nil,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
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
		return nil, status.Errorf(codes.Internal, "failed to get list brands: %w", err)
	}

	var productTypePb []*pb.ProductType
	for _, value := range productType {
		productTypePb = append(productTypePb, mapper.ProductTypeMapper(value))
	}

	return &pb.ProductTypeListResponse{
		Code:    200,
		Message: "success",
		Details: productTypePb,
	}, nil
}
