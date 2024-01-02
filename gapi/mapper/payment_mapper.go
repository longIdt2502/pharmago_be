package mapper

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func PaymentMapper(ctx context.Context, store *db.Store, id int32) *pb.Payment {

	paymentDb, _ := store.DetailPayment(ctx, id)

	paymentItemDb, _ := store.ListPaymentItem(ctx, id)

	var items []*pb.PaymentItem
	for _, value := range paymentItemDb {
		dataPb := &pb.PaymentItem{
			Id:        value.ID,
			Type:      nil,
			Value:     float32(value.Value),
			IsPaid:    value.IsPaid,
			ExtraNote: value.ExtraNote.String,
		}
		items = append(items, dataPb)
	}

	return &pb.Payment{
		Id:       paymentDb.ID,
		Code:     paymentDb.Code,
		MustPaid: float32(paymentDb.MustPaid),
		HadPaid:  float32(paymentDb.HadPaid),
		NeedPay:  float32(paymentDb.NeedPay),
		Items:    items,
	}
}
