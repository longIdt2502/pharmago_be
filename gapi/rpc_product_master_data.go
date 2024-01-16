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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *ServerGRPC) ClassifyList(ctx context.Context, req *pb.ClassifyListRequest) (*pb.ClassifyListResponse, error) {
	classifies, err := server.store.GetListClassify(ctx, db.GetListClassifyParams{
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
		return nil, status.Errorf(codes.Internal, "failed to get list classify: %w", err)
	}

	var classifiesPb []*pb.SimpleData
	for _, value := range classifies {
		dataPb := &pb.SimpleData{
			Id:   value.ID,
			Name: value.Name,
			Code: value.Code,
		}
		classifiesPb = append(classifiesPb, dataPb)
	}

	return &pb.ClassifyListResponse{
		Code:    200,
		Message: "success",
		Details: classifiesPb,
	}, nil
}

func (server *ServerGRPC) ProductionStandardList(ctx context.Context, req *pb.ProductionStandardListRequest) (*pb.ProductionStandardListResponse, error) {
	productionStandard, err := server.store.ListProductionStandard(ctx, db.ListProductionStandardParams{
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
		Company: req.GetCompany(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list production standard: %w", err)
	}

	var productionStandardPb []*pb.SimpleData
	for _, value := range productionStandard {
		var userCreatedName *string
		if value.FullName.Valid {
			name := value.FullName.String
			userCreatedName = &name
		}

		quantity := value.Quantity

		var description *string
		if value.Description.Valid {
			data := value.Description.String
			description = &data
		}

		dataPb := &pb.SimpleData{
			Id:              value.ID,
			Name:            value.Name,
			Code:            value.Code,
			UserCreatedName: userCreatedName,
			CreatedAt:       timestamppb.New(value.CreatedAt),
			ValueExtra:      &quantity,
			Description:     description,
		}
		productionStandardPb = append(productionStandardPb, dataPb)
	}

	return &pb.ProductionStandardListResponse{
		Code:    200,
		Message: "success",
		Details: productionStandardPb,
	}, nil
}

func (server *ServerGRPC) ProductionStandardCreate(ctx context.Context, req *pb.ProductionStandardCreateRequest) (*pb.ProductionStandardCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("PS-%s-%d", utils.RandomString(3), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}
	data, err := server.store.CreateProductionStandard(ctx, db.CreateProductionStandardParams{
		Code: code,
		Name: req.GetName(),
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
		UserCreated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		UserUpdated: sql.NullInt32{},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record production standard: ", err.Error())
	}

	return &pb.ProductionStandardCreateResponse{
		Code:    200,
		Message: "success",
		Details: data.ID,
	}, nil
}

func (server *ServerGRPC) ProductionStandardDetail(ctx context.Context, req *pb.ProductionStandardDetailRequest) (*pb.ProductionStandardDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	data, err := server.store.DetailProductionStandard(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get detail production standard: ", err.Error())
	}

	var userCreatedName *string
	if data.FullName.Valid {
		name := data.FullName.String
		userCreatedName = &name
	}

	var userUpdatedName *string
	if data.UserUpdatedName.Valid {
		nameUd := data.UserUpdatedName.String
		userUpdatedName = &nameUd
	}

	var description *string
	if data.Description.Valid {
		data := data.Description.String
		description = &data
	}

	return &pb.ProductionStandardDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.SimpleData{
			Id:              data.ID,
			Name:            data.Name,
			Code:            data.Code,
			UserCreatedName: userCreatedName,
			CreatedAt:       timestamppb.New(data.CreatedAt),
			UserUpdatedName: userUpdatedName,
			UpdatedAt:       timestamppb.New(data.UpdatedAt.Time),
			ValueExtra:      nil,
			Description:     description,
		},
	}, nil
}

func (server *ServerGRPC) ProductionStandardUpdate(ctx context.Context, req *pb.ProductionStandardUpdateRequest) (*pb.ProductionStandardUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	ps, err := server.store.DetailProductionStandard(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "production standard not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get production standard: ", err.Error())
	}
	if !ps.UserCreated.Valid {
		return nil, status.Errorf(codes.PermissionDenied, "failed to update production standard: ", err.Error())
	}

	data, err := server.store.UpdateProductionStandard(ctx, db.UpdateProductionStandardParams{
		Name: req.GetName(),
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  req.Code != nil,
		},
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		ID: req.GetId(),
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update production standard: ", err.Error())
	}

	return &pb.ProductionStandardUpdateResponse{
		Code:    200,
		Message: "success",
		Details: data.ID,
	}, nil
}

func (server *ServerGRPC) ProductionStandardDelete(ctx context.Context, req *pb.ProductionStandardDeleteRequest) (*pb.ProductionStandardDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteProductionStandard(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "production standard not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete production standard: ", err.Error())
	}

	return &pb.ProductionStandardDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) PreparationTypeList(ctx context.Context, req *pb.PreparationTypeListRequest) (*pb.PreparationTypeListResponse, error) {
	preparationType, err := server.store.GetListPreparation(ctx, db.GetListPreparationParams{
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
		return nil, status.Errorf(codes.Internal, "failed to get list preparation type: %w", err)
	}

	var preparationTypePb []*pb.SimpleData
	for _, value := range preparationType {
		dataPb := &pb.SimpleData{
			Id:   value.ID,
			Name: value.Name,
			Code: value.Code,
		}
		preparationTypePb = append(preparationTypePb, dataPb)
	}

	return &pb.PreparationTypeListResponse{
		Code:    200,
		Message: "success",
		Details: preparationTypePb,
	}, nil
}

func (server *ServerGRPC) CompanyPharmaList(ctx context.Context, req *pb.CompanyPharmaListRequest) (*pb.CompanyPharmaListResponse, error) {
	companyPharma, err := server.store.GetListCompanyPharma(ctx, db.GetListCompanyPharmaParams{
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
		Type: req.GetType(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list preparation type: %w", err)
	}

	var companyPharmaPb []*pb.CompanyPharma
	for _, value := range companyPharma {
		companyPharmaPb = append(companyPharmaPb, mapper.CompanyPharmaMapper(value))
	}

	return &pb.CompanyPharmaListResponse{
		Code:    200,
		Message: "success",
		Details: companyPharmaPb,
	}, nil
}
