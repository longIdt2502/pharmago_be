package gapi

import (
	"context"
	"database/sql"
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	products, err := server.store.GetProducts(ctx, db.GetProductsParams{
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
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get product: %e", err))
	}

	var productsPb []*pb.Product
	for _, value := range products {
		productsPb = append(productsPb, mapper.ProductMapper(ctx, server.store, value))
	}

	return &pb.ListProductResponse{
		Code:    200,
		Message: "success",
		Details: productsPb,
	}, nil
}
