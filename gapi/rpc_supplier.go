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

func (server *ServerGRPC) SupplierList(ctx context.Context, req *pb.SupplierListRequest) (*pb.SupplierListResponse, error) {
	suppliers, err := server.store.GetListSupplier(ctx, db.GetListSupplierParams{
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
		return nil, status.Errorf(codes.Internal, "failed to get suppliers: ", err.Error())
	}

	var suppliersPb []*pb.Supplier
	for _, value := range suppliers {
		dataPb := mapper.SupplierMapper(ctx, server.store, value)
		suppliersPb = append(suppliersPb, dataPb)
	}

	return &pb.SupplierListResponse{
		Code:    200,
		Message: "success",
		Details: suppliersPb,
	}, nil
}
