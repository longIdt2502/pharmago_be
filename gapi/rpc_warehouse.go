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

func (server *ServerGRPC) WarehouseCreate(ctx context.Context, req *pb.WarehouseCreateRequest) (*pb.WarehouseCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	address, err := server.store.CreateAddress(ctx, db.CreateAddressParams{
		Lat: float64(req.Address.GetLat()),
		Lng: float64(req.Address.GetLng()),
		Province: sql.NullString{
			String: req.Address.Province,
			Valid:  true,
		},
		District: sql.NullString{
			String: req.Address.District,
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
		return nil, status.Errorf(codes.Internal, "failed to record address data: %e", err)
	}

	codeW := fmt.Sprintf("WAREHOUSE-%s%d", utils.RandomString(3), utils.RandomInt(100, 999))
	if req.Code != nil {
		codeW = req.GetCode()
	}

	warehouse, err := server.store.CreateWarehouse(ctx, db.CreateWarehouseParams{
		Name: req.GetName(),
		Code: codeW,
		Address: sql.NullInt32{
			Int32: address.ID,
			Valid: true,
		},
		Companies: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record warehouse data: %e", err)
	}

	return &pb.WarehouseCreateResponse{
		Code:    200,
		Message: "success",
		Details: warehouse.ID,
	}, nil
}

func (server *ServerGRPC) WarehouseList(ctx context.Context, req *pb.WarehouseListRequest) (*pb.WarehouseListResponse, error) {
	warehouses, err := server.store.ListWarehouse(ctx, db.ListWarehouseParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
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
			Valid: req.Page != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get warehouse list: %e", err)
	}

	var warehousesPb []*pb.Warehouse
	for _, value := range warehouses {
		dataPb, err := mapper.WarehouseMapper(ctx, server.store, value)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to mapper warehouse: %e", err)
		}
		warehousesPb = append(warehousesPb, dataPb)
	}

	return &pb.WarehouseListResponse{
		Code:    200,
		Message: "success",
		Details: warehousesPb,
	}, nil
}

func (server *ServerGRPC) WarehouseDetail(ctx context.Context, req *pb.WarehouseDetailRequest) (*pb.WarehouseDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	warehouse, err := server.store.DetailWarehouse(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "warehouse not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get warehouse list: %e", err)
	}

	warehousesPb, err := mapper.WarehouseMapper(ctx, server.store, warehouse)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to mapper warehouse: %e", err)
	}

	return &pb.WarehouseDetailResponse{
		Code:    200,
		Message: "success",
		Details: warehousesPb,
	}, nil
}

func (server *ServerGRPC) WarehouseUpdate(ctx context.Context, req *pb.WarehouseUpdateRequest) (*pb.WarehouseUpdateResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	warehouse, err := server.store.DetailWarehouse(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "warehouse not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get warehouse: %e", err)
	}

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
		ID: warehouse.Address.Int32,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update address: %e", err)
	}

	_, err = server.store.UpdateWarehouse(ctx, db.UpdateWarehouseParams{
		ID: req.GetId(),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  req.Code != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update warehouse: %e", err)
	}

	return &pb.WarehouseUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) WarehouseDelete(ctx context.Context, req *pb.WarehouseDeleteRequest) (*pb.WarehouseDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteWarehouse(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "warehouse not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete warehouse: %e", err)
	}

	return &pb.WarehouseDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
