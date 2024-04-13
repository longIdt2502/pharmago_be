package mapper

import (
	"context"
	"database/sql"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OrderPreviewMapper(order db.ListOrderRow) *pb.OrderPreview {

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
			},
		})
		variantDb := VariantMapper(ctx, store, variant[0])
		orderItems = append(orderItems, &pb.OrderItem{
			Id:         value.ID,
			Variant:    variantDb,
			Value:      value.Value,
			TotalPrice: float32(value.TotalPrice),
			Consignment: &pb.Consignment{
				Id:          value.ID_3,
				Code:        value.Code_2,
				Quantity:    value.Quantity,
				Inventory:   value.Inventory,
				Variant:     nil,
				ExpiredAt:   timestamppb.New(value.ExpiredAt),
				ProducedAt:  timestamppb.New(value.ProductedAt),
				IsAvailable: value.IsAvailable,
			},
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
		UserCreated:  data.Username,
		UserUpdated:  "",
		CreatedAt:    timestamppb.New(data.CreatedAt),
		UpdatedAt:    nil,
		Payment:      payment,
		Items:        orderItems,
		ServiceItems: orderServiceItems,
	}
}
