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
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *ServerGRPC) BrandList(ctx context.Context, req *pb.BrandListRequest) (*pb.BrandListResponse, error) {
	brands, err := server.store.GetListBrand(ctx, db.GetListBrandParams{
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
		return nil, status.Errorf(codes.Internal, "failed to get list brands: %e", err)
	}

	var brandsPb []*pb.Brand
	for _, value := range brands {
		brandsPb = append(brandsPb, mapper.BrandMapper(value))
	}

	return &pb.BrandListResponse{
		Code:    200,
		Message: "success",
		Details: brandsPb,
	}, nil
}

func (server *ServerGRPC) BrandCreate(ctx context.Context, req *pb.BrandCreateRequest) (*pb.BrandCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("BRAND-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	brand, err := server.store.CreateBrand(ctx, db.CreateBrandParams{
		Code: code,
		Name: req.GetName(),
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		Company:     req.GetCompany(),
		UserCreated: tokenPayload.UserID,
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create brand: %v", err))
	}

	for _, item := range req.GetProducts() {
		_, err = server.store.UpdateProduct(ctx, db.UpdateProductParams{
			Brand: sql.NullInt32{
				Int32: brand.ID,
				Valid: true,
			},
			ID: item,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update product: %v", err))
		}
	}

	return &pb.BrandCreateResponse{
		Code:    200,
		Message: "success",
		Details: brand.ID,
	}, nil
}

func (server *ServerGRPC) BrandUpdate(ctx context.Context, req *pb.BrandUpdateRequest) (*pb.BrandUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("BRAND-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	brand, err := server.store.UpdateBrand(ctx, db.UpdateBrandParams{
		Code: code,
		Name: req.GetName(),
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		ID: req.GetId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update brand: %v", err))
	}

	products, _ := server.store.GetProducts(ctx, db.GetProductsParams{
		Company: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Brand: sql.NullInt32{
			Int32: req.GetId(),
			Valid: true,
		},
	})

	for _, item := range products {
		if !funk.ContainsInt32(req.GetProducts(), item.ID) {
			_, err := server.store.UpdateProduct(ctx, db.UpdateProductParams{
				ID: item.ID,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update product")
			}
		}
	}

	return &pb.BrandUpdateResponse{
		Code:    200,
		Message: "success",
		Details: brand.ID,
	}, nil
}

func (server *ServerGRPC) BrandDetail(ctx context.Context, req *pb.BrandDetailRequest) (*pb.BrandDetailResponse, error) {
	brand, err := server.store.DetailBrand(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "brand note exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get brand: %v", err))
	}

	return &pb.BrandDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.SimpleData{
			Id:              brand.ID,
			Name:            brand.Name,
			Code:            brand.Code,
			UserCreatedName: &(brand.CreatedName),
			CreatedAt:       timestamppb.New(brand.CreatedAt),
			UserUpdatedName: &(brand.UpdatedName.String),
			UpdatedAt:       timestamppb.New(brand.UpdatedAt.Time),
			ValueExtra:      nil,
			Description:     &(brand.Description.String),
		},
	}, nil
}

func (server *ServerGRPC) BrandDelete(ctx context.Context, req *pb.BrandDeleteRequest) (*pb.BrandDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteBrand(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "brand note exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete brand: %v", err))
	}

	return &pb.BrandDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
