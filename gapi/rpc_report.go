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

func (server *ServerGRPC) ReportRevenue(ctx context.Context, req *pb.ReportRevenueRequest) (*pb.ReportRevenueResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	filter := "day"
	switch req.Filter {
	case pb.ReportRevenueRequest_YEAR:
		filter = "month"
	}

	reports, err := server.store.GetReportRevenue(ctx, db.GetReportRevenueParams{
		Filter:  filter,
		Company: req.GetCompany(),
		Status:  sql.NullString{},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get data: %e", err)
	}

	var reportsPb []*pb.ReportItem
	for _, item := range reports {
		reportsPb = append(reportsPb, &pb.ReportItem{
			Title:      item.Date.Format("2006-01-02T15:04:05"),
			Value:      float32(item.CurrentRevenue),
			ValueExtra: float32(item.LastRevenue),
		})
	}

	switch req.Filter {
	case pb.ReportRevenueRequest_YEAR:
		filter = "year"
	case pb.ReportRevenueRequest_MONTH:
		filter = "month"
	}

	currentValue, _ := server.store.TotalRevenue(ctx, db.TotalRevenueParams{
		Company:  req.GetCompany(),
		Filter:   filter,
		Interval: 0,
	})
	lastValue, _ := server.store.TotalRevenue(ctx, db.TotalRevenueParams{
		Company:  req.GetCompany(),
		Filter:   filter,
		Interval: 1,
	})

	return &pb.ReportRevenueResponse{
		Code:         200,
		Message:      "success",
		Details:      reportsPb,
		CurrentValue: float32(currentValue),
		LastValue:    float32(lastValue),
	}, nil
}
