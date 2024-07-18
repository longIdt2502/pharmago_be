package gapi

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/thoas/go-funk"
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
			Id:               value.ID,
			Code:             value.Code,
			Name:             value.Name,
			Barcode:          &value.Barcode.String,
			DecisionNumber:   &value.DecisionNumber.String,
			RegisterNumber:   &value.RegisterNumber.String,
			Longevity:        &value.Longevity.String,
			Product:          value.Product,
			Media:            value.Imageurl,
			QuantityInStock:  nil,
			Units:            nil,
			PriceSell:        0,
			PriceImport:      0,
			Revenue:          revenue,
			InitialInventory: value.InitialInventory,
			RealInventory:    value.RealInventory,
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

func (server *ServerGRPC) ReportOrder(ctx context.Context, req *pb.ReportOrderRequest) (*pb.ReportOrderResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	report, err := server.store.TotalOrderByMonth(ctx, req.GetCompany())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get report order: %e", err)
	}

	var reportPb []*pb.ReportItem
	for _, item := range report {
		reportPb = append(reportPb, &pb.ReportItem{
			Title: item.Date.Format("2006-01-02T15:04:05"),
			Value: float32(item.Count),
		})
	}

	var currentValue db.TotalOrderByMonthRow
	var lastValue db.TotalOrderByMonthRow

	currentValue = funk.Find(report, func(item db.TotalOrderByMonthRow) bool {
		return time.Unix(item.Date.Unix(), 0).Month() == time.Now().Month()
	}).(db.TotalOrderByMonthRow)

	lastValue = funk.Find(report, func(item db.TotalOrderByMonthRow) bool {
		return time.Unix(item.Date.Unix(), 0).Month() == time.Now().AddDate(0, -1, 0).Month()
	}).(db.TotalOrderByMonthRow)

	return &pb.ReportOrderResponse{
		Code:         200,
		Message:      "success",
		Details:      reportPb,
		CurrentValue: float32(currentValue.Count),
		LastValue:    float32(lastValue.Count),
	}, nil
}

func (server *ServerGRPC) ReportCustomer(ctx context.Context, req *pb.ReportCustomerRequest) (*pb.ReportCustomerResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	customer, err := server.store.TotalCustomerByMonth(ctx, req.GetCompany())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get report customer: %e", err)
	}

	var currentValue db.TotalCustomerByMonthRow
	var lastValue db.TotalCustomerByMonthRow

	currentValue = funk.Find(customer, func(item db.TotalCustomerByMonthRow) bool {
		return time.Unix(item.Date.Unix(), 0).Month() == time.Now().Month()
	}).(db.TotalCustomerByMonthRow)

	lastValue = funk.Find(customer, func(item db.TotalCustomerByMonthRow) bool {
		return time.Unix(item.Date.Unix(), 0).Month() == time.Now().AddDate(0, -1, 0).Month()
	}).(db.TotalCustomerByMonthRow)

	return &pb.ReportCustomerResponse{
		Code:         200,
		Message:      "success",
		CurrentValue: float32(currentValue.Count),
		LastValue:    float32(lastValue.Count),
	}, nil
}

func (server *ServerGRPC) ReportCustomerRevenue(ctx context.Context, req *pb.ReportCustomerRevenueRequest) (*pb.ReportCustomerRevenueResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	orderBy := "quantity"
	if req.OrderBy.String() == "REVENUE" {
		orderBy = "revenue"
	}

	reports, err := server.store.ReportCustomerRevenue(ctx, db.ReportCustomerRevenueParams{
		Company: req.GetCompany(),
		OrderBy: orderBy,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get report customer: %e", err)
	}

	var reportPb []*pb.CustomerReportItem
	for _, item := range reports {

		reportPb = append(reportPb, &pb.CustomerReportItem{
			Id:       item.ID.Int32,
			FullName: item.FullName.String,
			// Image: item.,
			Quantity: int32(item.CountOrder),
			Revenue:  float32(item.TotalPrice),
		})
	}

	totalOrder, _ := server.store.CountOrder(ctx, req.GetCompany())

	totalCustomer, _ := server.store.CountCustomer(ctx, req.GetCompany())

	return &pb.ReportCustomerRevenueResponse{
		Code:    200,
		Message: "success",
		Details: reportPb,
		Total:   int32(totalCustomer),
		Average: float32(totalOrder / int32(totalCustomer)),
	}, nil
}
