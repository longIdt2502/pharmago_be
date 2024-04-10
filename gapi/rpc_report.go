package gapi

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) HomeData(ctx context.Context, req *pb.HomeDataRequest) (*pb.HomeDataResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	company, err := server.store.GetCompanyById(ctx, req.GetCompany())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "company not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get company: %e", err)
	}

	companyPb := mapper.CompanyMapper(ctx, server.store, company)

	revenue, err := server.store.GetRevenueCompany(ctx, req.GetCompany())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get revenue: %e", err)
	}

	orders, _ := server.store.ListOrder(ctx, db.ListOrderParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
		Status: sql.NullString{
			String: "COMPLETE",
			Valid:  true,
		},
	})

	variants, err := server.store.GetVariantBestSale(ctx, req.GetCompany())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get variants: %e", err)
	}

	var variantsPb []*pb.Variant
	for _, value := range variants {
		var revenue *float32
		f64 := float32(value.Revenue)
		revenue = &f64
		dataPb := &pb.Variant{
			Id:              value.ID,
			Code:            value.Code,
			Name:            value.Name,
			Barcode:         value.Barcode,
			DecisionNumber:  value.DecisionNumber,
			RegisterNumber:  value.RegisterNumber,
			Longevity:       value.Longevity,
			Vat:             float32(value.Vat),
			Product:         value.Product,
			Media:           value.Imageurl,
			QuantityInStock: nil,
			Units:           nil,
			PriceSell:       0,
			PriceImport:     0,
			Revenue:         revenue,
		}
		variantsPb = append(variantsPb, dataPb)
	}

	return &pb.HomeDataResponse{
		Company:         companyPb,
		Revenue:         float32(revenue),
		OrderComplete:   int32(len(orders)),
		VariantBestSale: variantsPb,
	}, nil
}
