package mapper

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func SupplierMapper(ctx context.Context, store *db.Store, data db.Suplier) *pb.Supplier {
	addressDb, _ := store.GetAddress(ctx, data.Address.Int32)
	addressPb := AddressMapper(ctx, store, addressDb)

	return &pb.Supplier{
		Id:         data.ID,
		Code:       data.Code,
		Name:       data.Name,
		DeputyName: data.DeputyName,
		Phone:      data.Phone,
		Email:      data.Email.String,
		Address:    addressPb,
		Company:    data.Company.Int32,
	}
}
