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
	productionStandard, err := server.store.GetListProductionStandard(ctx, db.GetListProductionStandardParams{
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
		return nil, status.Errorf(codes.Internal, "failed to get list production standard: %w", err)
	}

	var productionStandardPb []*pb.SimpleData
	for _, value := range productionStandard {
		dataPb := &pb.SimpleData{
			Id:   value.ID,
			Name: value.Name,
			Code: value.Code,
		}
		productionStandardPb = append(productionStandardPb, dataPb)
	}

	return &pb.ProductionStandardListResponse{
		Code:    200,
		Message: "success",
		Details: productionStandardPb,
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
