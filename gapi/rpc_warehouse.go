package gapi

import (
	"context"
	"database/sql"
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
		return nil, status.Errorf(codes.Internal, "failed to record address data: %w", err)
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
		return nil, status.Errorf(codes.Internal, "failed to record warehouse data: %w", err)
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
		return nil, status.Errorf(codes.Internal, "failed to get warehouse list: %w", err)
	}

	var warehousesPb []*pb.Warehouse
	for _, value := range warehouses {
		dataPb, err := mapper.WarehouseMapper(ctx, server.store, value)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to mapper warehouse: %w", err)
		}
		warehousesPb = append(warehousesPb, dataPb)
	}

	return &pb.WarehouseListResponse{
		Code:    200,
		Message: "success",
		Details: warehousesPb,
	}, nil
}
