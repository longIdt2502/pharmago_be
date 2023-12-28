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

func (server *ServerGRPC) ListVariant(ctx context.Context, req *pb.ListVariantRequest) (*pb.ListVariantResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	variants, err := server.store.GetVariants(ctx, db.GetVariantsParams{
		Company: req.GetCompany(),
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
		return nil, status.Errorf(codes.Internal, "failed to get variant record: ", err)
	}

	var variantsPb []*pb.Variant
	for _, value := range variants {
		data := mapper.VariantMapper(ctx, server.store, value)
		variantsPb = append(variantsPb, data)
	}

	return &pb.ListVariantResponse{
		Code:    200,
		Message: "success",
		Details: variantsPb,
	}, nil
}
