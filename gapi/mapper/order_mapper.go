package mapper

import (
	"context"
	"database/sql"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OrderPreviewMapper(ctx context.Context, store *db.Store, order db.ListOrderRow) *pb.OrderPreview {
	var firstVariant *pb.Variant
	var firstService *pb.Service
	var countItems int

	if order.Type.String == "SELL" {
		orderItemsDb, _ := store.ListOrderItem(ctx, order.ID)
		countItems = len(orderItemsDb)
		if countItems > 0 {
			firstVariant = &pb.Variant{
				Id:          orderItemsDb[0].ID,
				Code:        orderItemsDb[0].Code,
				Name:        orderItemsDb[0].Name,
				Media:       orderItemsDb[0].MediaUrl.String,
				QuantityBuy: orderItemsDb[0].Quantity.Int32,
				PriceSell:   float32(orderItemsDb[0].TotalPrice),
			}
		}

	} else {
		orderServiceItemsDb, _ := store.ListOrderServiceItem(ctx, order.ID)
		countItems = len(orderServiceItemsDb)
		if countItems > 0 {
			firstService = &pb.Service{
				Id:    orderServiceItemsDb[0].ID,
				Code:  orderServiceItemsDb[0].Code,
				Title: orderServiceItemsDb[0].Title,
			}
		}
	}

	return &pb.OrderPreview{
		Id:           order.ID,
		Code:         order.Code,
		TotalPrice:   float32(order.TotalPrice),
		Status:       &pb.SimpleData{Id: order.OsID, Name: order.OsTitle, Code: order.Status.String},
		Type:         &pb.SimpleData{Name: order.Title_3, Code: order.Code_6},
		Description:  order.Description.String,
		Payment:      &pb.Payment{Id: order.ID_6, Code: order.Code_5, MustPaid: float32(order.MustPaid_2), HadPaid: float32(order.HadPaid), NeedPay: float32(order.NeedPay)},
		CustomerName: order.CFullName,
		UserCreated:  order.AFullName,
		CreatedAt:    timestamppb.New(order.CreatedAt),
		FirstVariant: firstVariant,
		FirstService: firstService,
		CountItems:   int32(countItems),
	}
}

func OrderMedicalBillMapper(order db.ListByMedicalBillRow) *pb.OrderPreview {

	return &pb.OrderPreview{
		Id:         order.ID,
		Code:       order.Code,
		TotalPrice: float32(order.TotalPrice),
		Status: &pb.SimpleData{
			Id:   order.OsID,
			Name: order.OsTitle,
			Code: order.Status.String,
		},
		Description:  order.Description.String,
		CustomerName: order.CFullName,
		UserCreated:  order.AFullName,
		CreatedAt:    timestamppb.New(order.CreatedAt),
		Payment: &pb.Payment{
			Id:       order.ID_6,
			Code:     order.Code_5,
			MustPaid: float32(order.MustPaid_2),
			HadPaid:  float32(order.HadPaid),
			NeedPay:  float32(order.NeedPay),
		},
		Type: &pb.SimpleData{
			Name: order.Title_3,
			Code: order.Code_6,
		},
	}
}

func OrderDetailMapper(ctx context.Context, store *db.Store, data db.DetailOrderRow) *pb.Order {

	customerDb, _ := store.GetCustomer(ctx, data.Customer.Int32)
	customer, _ := CustomerDetailMapper(ctx, store, customerDb)

	var address *pb.Address
	if data.Address.Valid {
		addressDb, _ := store.GetAddress(ctx, data.Address.Int32)
		address = AddressMapper(ctx, store, addressDb)
	}

	payment := PaymentMapper(ctx, store, data.Payment)

	var orderItems []*pb.OrderItem
	orderItemsDb, _ := store.ListOrderItem(ctx, data.ID)
	for _, value := range orderItemsDb {

		variant, _ := store.GetVariants(ctx, db.GetVariantsParams{
			Company: data.Company,
			ID: sql.NullInt32{
				Int32: value.Variant,
				Valid: true,
			},
		})
		variantDb := VariantMapper(ctx, store, variant[0])
		orderItems = append(orderItems, &pb.OrderItem{
			Id:         value.ID,
			Variant:    variantDb,
			Value:      value.Value,
			TotalPrice: float32(value.TotalPrice),
			// Consignment: &pb.Consignment{
			// 	Id:          value.ID_3,
			// 	Code:        value.Code_2,
			// 	Quantity:    value.Quantity,
			// 	Inventory:   value.Inventory,
			// 	Variant:     nil,
			// 	ExpiredAt:   timestamppb.New(value.ExpiredAt),
			// 	ProducedAt:  timestamppb.New(value.ProductedAt),
			// 	IsAvailable: value.IsAvailable,
			// },
		})
	}

	var orderServiceItems []*pb.OrderServiceItem
	orderServiceItemsDb, _ := store.ListOrderServiceItem(ctx, data.ID)
	for _, item := range orderServiceItemsDb {
		orderServiceItems = append(orderServiceItems, &pb.OrderServiceItem{
			Id:         item.ID,
			UnitPrice:  float32(item.UnitPrice),
			TotalPrice: float32(item.TotalPrice),
			Discount:   float32(item.Discount),
			Service: ServiceMapper(db.Service{
				ID:    item.ID_2,
				Title: item.Title,
				Code:  item.Code,
			}),
			Quantity: item.Quantity.Int32,
		})
	}

	return &pb.Order{
		Id:           data.ID,
		Code:         data.Code,
		TotalPrice:   float32(data.TotalPrice),
		Description:  data.Description.String,
		Vat:          float32(data.Vat),
		Discount:     data.Discount,
		ServicePrice: float32(data.ServicePrice),
		MustPaid:     float32(data.MustPaid),
		Customer:     customer,
		Address:      address,

		Type: &pb.SimpleData{
			Id:   data.OtID,
			Name: data.OtTitle,
			Code: data.OtCode,
		},
		Status: &pb.SimpleData{
			Id:   data.OsID,
			Name: data.OsTitle,
			Code: data.OsCode,
		},
		Qr:           data.QrUrl,
		Company:      data.Company,
		UserCreated:  data.AFullName,
		UserUpdated:  data.FullName_2,
		CreatedAt:    timestamppb.New(data.CreatedAt),
		UpdatedAt:    timestamppb.New(data.UpdatedAt.Time),
		Payment:      payment,
		Items:        orderItems,
		ServiceItems: orderServiceItems,
	}
}
