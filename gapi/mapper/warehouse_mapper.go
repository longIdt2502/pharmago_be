package mapper

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/pkg/errors"
)

func WarehouseMapper(ctx context.Context, store *db.Store, data db.Warehouse) (*pb.Warehouse, error) {
	print(data.Address.Int32)
	address, err := store.GetAddress(ctx, data.Address.Int32)
	if err != nil {
		return nil, errors.Wrap(err, "address err")
	}
	addressPb := AddressMapper(ctx, store, address)

	company, err := store.GetCompanyById(ctx, data.Companies.Int32)
	if err != nil {
		return nil, errors.Wrap(err, "company err")
	}

	companyPb := CompanyMapper(ctx, store, company)

	return &pb.Warehouse{
		Id:      data.ID,
		Name:    data.Name,
		Code:    data.Code,
		Address: addressPb,
		Company: companyPb,
	}, nil
}
