package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
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

func (server *ServerGRPC) SupplierCreate(ctx context.Context, req *pb.SupplierCreateRequest) (*pb.SupplierCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

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
		return nil, status.Errorf(codes.Internal, "failed to record address: ", err.Error())
	}

	code := fmt.Sprintf("SUPPLIER-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	supplier, err := server.store.CreateSupplier(ctx, db.CreateSupplierParams{
		Code:       code,
		Name:       req.GetName(),
		DeputyName: req.GetDeputy(),
		Phone:      req.GetPhone(),
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Address: sql.NullInt32{
			Int32: address.ID,
			Valid: true,
		},
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record address: ", err.Error())
	}

	return &pb.SupplierCreateResponse{
		Code:    200,
		Message: "success",
		Details: supplier.ID,
	}, nil
}

func (server *ServerGRPC) SupplierDetail(ctx context.Context, req *pb.SupplierDetailRequest) (*pb.SupplierDetailResponse, error) {
	supplier, err := server.store.DetailSupplier(ctx, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "supplier not exists")
		}
		return nil, status.Errorf(codes.Internal, "")
	}

	supplierPb := mapper.SupplierMapper(ctx, server.store, supplier)
	return &pb.SupplierDetailResponse{
		Code:    200,
		Message: "success",
		Details: supplierPb,
	}, nil
}

func (server *ServerGRPC) SupplierUpdate(ctx context.Context, req *pb.SupplierUpdateRequest) (*pb.SupplierUpdateResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	supplier, err := server.store.DetailSupplier(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "supplier not exists")
		}
		return nil, status.Errorf(codes.Internal, "")
	}

	if supplier.Address.Valid {
		_, err = server.store.UpdateAddress(ctx, db.UpdateAddressParams{
			Lat: sql.NullFloat64{
				Float64: float64(req.Address.GetLat()),
				Valid:   true,
			},
			Lng: sql.NullFloat64{
				Float64: float64(req.Address.GetLat()),
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
				Valid:  true,
			},
			Title: sql.NullString{
				String: req.Address.GetTitle(),
				Valid:  true,
			},
			ID: supplier.Address.Int32,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update address: ", err.Error())
		}
	}

	_, err = server.store.UpdateSupplier(ctx, db.UpdateSupplierParams{
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  true,
		},
		DeputyName: sql.NullString{
			String: req.GetDeputy(),
			Valid:  true,
		},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  true,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  true,
		},
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update supplier: ", err.Error())
	}

	return &pb.SupplierUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}
