package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
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
		return nil, status.Errorf(codes.Internal, "failed to get customer: %e", err)
	}

	var customersPb []*pb.CustomerDetail
	for _, value := range customers {
		dataPb, _ := mapper.CustomerDetailMapper(ctx, server.store, value)
		customersPb = append(customersPb, dataPb)
	}

	return &pb.CustomerListResponse{
		Code:    200,
		Message: "success",
		Details: customersPb,
	}, nil
}

func (server *ServerGRPC) CustomerCreate(ctx context.Context, req *pb.CustomerCreateRequest) (*pb.CustomerCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("CUSTOMER-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	var addressId int32
	if req.Address != nil {
		address, err := server.store.CreateAddress(ctx, db.CreateAddressParams{
			Lat: float64(req.Address.GetLat()),
			Lng: float64(req.Address.GetLng()),
			Province: sql.NullString{
				String: req.Address.GetProvince(),
				Valid:  true,
			},
			District: sql.NullString{
				String: req.Address.GetDistrict(),
				Valid:  true,
			},
			Ward: sql.NullString{
				String: req.Address.GetWard(),
				Valid:  req.Address.Ward != nil,
			},
			Title:       req.Address.GetTitle(),
			UserCreated: tokenPayload.UserID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record address: %e", err)
		}

		addressId = address.ID
	}

	customer, err := server.store.CreateCustomer(ctx, db.CreateCustomerParams{
		FullName: req.GetName(),
		Code:     code,
		Company:  req.GetCompany(),
		Address: sql.NullInt32{
			Int32: addressId,
			Valid: req.Address != nil,
		},
		Email: sql.NullString{},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  true,
		},
		License: sql.NullString{},
		Birthday: sql.NullTime{
			Time:  time.Unix(req.GetBirthday().GetSeconds(), 0),
			Valid: req.Birthday != nil,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		UserCreated: tokenPayload.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record customer: %e", err)
	}

	return &pb.CustomerCreateResponse{
		Code:    200,
		Message: "success",
		Details: customer.ID,
	}, nil
}

func (server *ServerGRPC) CustomerDetail(ctx context.Context, req *pb.CustomerDetailRequest) (*pb.CustomerDetailResponse, error) {
	customer, err := server.store.DetailCustomer(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get customer: %e", err)
	}

	customerPb, err := mapper.CustomerDetailMapper(ctx, server.store, customer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to mapper customer: %e", err)
	}

	return &pb.CustomerDetailResponse{
		Code:    200,
		Message: "success",
		Details: customerPb,
	}, nil
}

func (server *ServerGRPC) CustomerUpdate(ctx context.Context, req *pb.CustomerUpdateRequest) (*pb.CustomerUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	customerDb, err := server.store.DetailCustomer(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get customer: %e", err)
	}

	if req.Address != nil {
		if customerDb.Address.Valid {
			_, err = server.store.UpdateAddress(ctx, db.UpdateAddressParams{
				Lat: sql.NullFloat64{
					Float64: float64(req.Address.GetLat()),
					Valid:   true,
				},
				Lng: sql.NullFloat64{
					Float64: float64(req.Address.GetLng()),
					Valid:   true,
				},
				Province: sql.NullString{
					String: req.Address.GetProvince(),
					Valid:  true,
				},
				District: sql.NullString{
					String: req.Address.GetDistrict(),
					Valid:  true,
				},
				Ward: sql.NullString{
					String: req.Address.GetWard(),
					Valid:  req.Address.Ward != nil,
				},
				Title: sql.NullString{
					String: req.Address.GetTitle(),
					Valid:  true,
				},
				ID: customerDb.Address.Int32,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update address: %e", err)
			}
		}
		// else {
		// 	// TODO: create new address
		// }
	}

	_, err = server.store.UpdateCustomer(ctx, db.UpdateCustomerParams{
		FullName: sql.NullString{
			String: req.GetName(),
			Valid:  true,
		},
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  true,
		},
		Email: sql.NullString{},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  true,
		},
		License: sql.NullString{},
		Birthday: sql.NullTime{
			Time:  time.Unix(req.GetBirthday().GetSeconds(), 0),
			Valid: req.Birthday != nil,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update customer: %e", err)
	}

	return &pb.CustomerUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}
