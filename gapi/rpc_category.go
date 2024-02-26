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
		return nil, status.Errorf(codes.Internal, "failed to get list categorys: %w", err)
	}

	var categorysPb []*pb.Category
	for _, value := range categories {
		categorysPb = append(categorysPb, mapper.CategoryMapper(value))
	}

	return &pb.CategoryListResponse{
		Code:    200,
		Message: "success",
		Details: categorysPb,
	}, nil
}

func (server *ServerGRPC) CategoryCreate(ctx context.Context, req *pb.CategoryCreateRequest) (*pb.CategoryCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("CATEGORY-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	category, err := server.store.CreateCategory(ctx, db.CreateCategoryParams{
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
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create category: %v", err))
	}

	for _, item := range req.GetProducts() {
		_, err = server.store.UpdateProduct(ctx, db.UpdateProductParams{
			ProductCategory: sql.NullInt32{
				Int32: category.ID,
				Valid: true,
			},
			ID: item,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update product: %v", err))
		}
	}

	return &pb.CategoryCreateResponse{
		Code:    200,
		Message: "success",
		Details: category.ID,
	}, nil
}

func (server *ServerGRPC) CategoryUpdate(ctx context.Context, req *pb.CategoryUpdateRequest) (*pb.CategoryUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("category-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	category, err := server.store.UpdateCategory(ctx, db.UpdateCategoryParams{
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
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to update category: %v", err))
	}

	products, err := server.store.GetProducts(ctx, db.GetProductsParams{
		Company: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		ProductCategory: sql.NullInt32{
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

	return &pb.CategoryUpdateResponse{
		Code:    200,
		Message: "success",
		Details: category.ID,
	}, nil
}

func (server *ServerGRPC) CategoryDetail(ctx context.Context, req *pb.CategoryDetailRequest) (*pb.CategoryDetailResponse, error) {
	category, err := server.store.DetailCategory(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "category note exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get category: %v", err))
	}

	return &pb.CategoryDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.SimpleData{
			Id:              category.ID,
			Name:            category.Name,
			Code:            category.Code,
			UserCreatedName: &(category.CreatedName),
			CreatedAt:       timestamppb.New(category.CreatedAt),
			UserUpdatedName: &(category.UpdatedName.String),
			UpdatedAt:       timestamppb.New(category.UpdatedAt.Time),
			ValueExtra:      nil,
			Description:     &(category.Description.String),
		},
	}, nil
}

func (server *ServerGRPC) CategoryDelete(ctx context.Context, req *pb.CategoryDeleteRequest) (*pb.CategoryDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteCategory(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "category note exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete category: %v", err))
	}

	return &pb.CategoryDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
