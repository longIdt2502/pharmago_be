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

func (server *ServerGRPC) CustomerList(ctx context.Context, req *pb.CustomerListRequest) (*pb.CustomerListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	customers, err := server.store.ListCustomer(ctx, db.ListCustomerParams{
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
		return nil, status.Errorf(codes.Internal, "failed to get customer: ", err.Error())
	}

	var customersPb []*pb.Customer
	for _, value := range customers {
		dataPb := mapper.CustomerMapper(value)
		customersPb = append(customersPb, dataPb)
	}

	return &pb.CustomerListResponse{
		Code:    200,
		Message: "success",
		Details: customersPb,
	}, nil
}
