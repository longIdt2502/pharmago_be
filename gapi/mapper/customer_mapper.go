package mapper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CustomerMapper(data db.Customer) *pb.Customer {

	return &pb.Customer{
		Id:       data.ID,
		Code:     data.Code,
		FullName: data.FullName,
		Company:  data.Company,
		Phone:    data.Phone.String,
		Email:    nil,
	}
}

func CustomerDetailMapper(ctx context.Context, store *db.Store, data db.Customer) (*pb.CustomerDetail, error) {
	var addressPb *pb.Address
	if data.Address.Valid {
		address, err := store.GetAddress(ctx, data.Address.Int32)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "address not exists")
			}
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get address: %e", err))
		}
		addressPb = AddressMapper(ctx, store, address)
	}

	var contactAddressPb *pb.Address
	if data.ContactAddress.Valid {
		address, err := store.GetAddress(ctx, data.ContactAddress.Int32)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "address not exists")
			}
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get address: %e", err))
		}
		contactAddressPb = AddressMapper(ctx, store, address)
	}

	var birthday *timestamp.Timestamp
	if data.Birthday.Valid {
		birthday = timestamppb.New(data.Birthday.Time)
	}

	var licenseDate *timestamp.Timestamp
	if data.LicenseDate.Valid {
		licenseDate = timestamppb.New(data.LicenseDate.Time)
	}

	return &pb.CustomerDetail{
		Id:             data.ID,
		Code:           data.Code,
		FullName:       data.FullName,
		Company:        data.Company,
		Address:        addressPb,
		Phone:          data.Phone.String,
		Email:          nil,
		Birthday:       birthday,
		Title:          &data.Title.String,
		LicenseDate:    licenseDate,
		ContactName:    &data.ContactName.String,
		ContactTitle:   &data.ContactTitle.String,
		ContactPhone:   &data.ContactPhone.String,
		ContactEmail:   &data.ContactEmail.String,
		ContactAddress: contactAddressPb,
		AccountNumber:  &data.AccountNumber.String,
		BankName:       &data.BankName.String,
		BankBranch:     &data.BankBranch.String,
	}, nil
}
