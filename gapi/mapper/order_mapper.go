package mapper

import (
	"context"
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
	customer := CustomerMapper(customerDb)

	var address *pb.Address
	if data.Address.Valid {
		addressDb, _ := store.GetAddress(ctx, data.Address.Int32)
		address = AddressMapper(ctx, store, addressDb)
	}

	var payment *pb.Payment
	payment = PaymentMapper(ctx, store, data.Payment)

	var orderItems []*pb.OrderItem
	orderItemsDb, _ := store.ListOrderItem(ctx, data.ID)
	for _, value := range orderItemsDb {
		orderItems = append(orderItems, &pb.OrderItem{
			Id: value.ID,
			Variant: &pb.Variant{
				Id:              value.ID_2,
				Code:            value.Code,
				Name:            value.Name,
				Barcode:         value.Barcode,
				DecisionNumber:  value.DecisionNumber,
				RegisterNumber:  value.RegisterNumber,
				Longevity:       value.Longevity,
				Vat:             float32(value.Vat),
				Product:         value.Product,
				Media:           value.MediaUrl,
				QuantityInStock: nil,
				Units:           nil,
				PriceSell:       0,
				PriceImport:     0,
			},
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
		Qr:          data.QrUrl,
		Company:     data.Company,
		UserCreated: data.Username,
		UserUpdated: "",
		CreatedAt:   timestamppb.New(data.CreatedAt),
		UpdatedAt:   nil,
		Payment:     payment,
		Items:       orderItems,
	}
}
