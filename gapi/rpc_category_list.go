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

func (server *ServerGRPC) CategoryList(ctx context.Context, req *pb.CategoryListRequest) (*pb.CategoryListResponse, error) {
	categories, err := server.store.GetListCategory(ctx, db.GetListCategoryParams{
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

	var brandsPb []*pb.Category
	for _, value := range categories {
		brandsPb = append(brandsPb, mapper.CategoryMapper(value))
	}

	return &pb.CategoryListResponse{
		Code:    200,
		Message: "success",
		Details: brandsPb,
	}, nil
}
