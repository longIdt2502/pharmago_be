package mapper

import (
	"context"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func CompanyMapper(ctx context.Context, store *db.Store, data db.Company) *pb.Company {
	var taxCode *string
	if data.TaxCode.Valid {
		taxCode = &data.TaxCode.String
	}

	var phone *string
	if data.Phone.Valid {
		phone = &data.Phone.String
	}

	var addressPb *pb.Address
	if data.Address.Valid {
		address, _ := store.GetAddress(ctx, data.Address.Int32)

		var provincePb *pb.AddressItem
		if address.Province.Valid {
			province, _ := store.GetProvinceByCode(ctx, address.Province.String)
			provincePb = &pb.AddressItem{
				Code:       province.Code,
				Name:       province.Name,
				NameEn:     province.NameEn,
				FullName:   province.FullName,
				FullNameEn: province.FullNameEn,
			}
		}

		var districtPb *pb.AddressItem
		if address.District.Valid {
			district, _ := store.GetDistrictByCode(ctx, address.District.String)
			districtPb = &pb.AddressItem{
				Code:       district.Code,
				Name:       district.Name,
				NameEn:     district.NameEn,
				FullName:   district.FullName,
				FullNameEn: district.FullNameEn,
			}
		}

		var wardPb *pb.AddressItem
		if address.Ward.Valid {
			ward, _ := store.GetWardByCode(ctx, address.Ward.String)
			wardPb = &pb.AddressItem{
				Code:       ward.Code,
				Name:       ward.Name,
				NameEn:     ward.NameEn,
				FullName:   ward.FullName,
				FullNameEn: ward.FullNameEn,
			}
		}

		addressPb = &pb.Address{
			Id:       int32(address.ID),
			Lat:      float32(address.Lat),
			Lng:      float32(address.Lng),
			Province: provincePb,
			District: districtPb,
			Ward:     wardPb,
			Title:    address.Title,
		}
	}

	var timeOpen *string
	if data.TimeOpen.Valid {
		x := data.TimeOpen.Time.Format("15:04:05")
		timeOpen = &x
	}

	var timeClose *string
	if data.TimeClose.Valid {
		x := data.TimeClose.Time.Format("15:04:05")
		timeClose = &x
	}

	totalEmployee, _ := store.CountEmployee(ctx, data.ID)

	return &pb.Company{
		Id:            int32(data.ID),
		Name:          data.Name,
		Code:          data.Code,
		Type:          data.Type,
		TaxCode:       taxCode,
		Phone:         phone,
		Description:   nil,
		Address:       addressPb,
		Owner:         int32(data.Owner),
		Manager:       &pb.Account{},
		OaId:          &data.OaID.String,
		TimeOpen:      timeOpen,
		TimeClose:     timeClose,
		TotalEmployee: int32(totalEmployee),
		UserCreated:   &pb.Account{},
		UserUpdated:   &pb.Account{},
	}
}

func CompanyDetailMapper(ctx context.Context, store *db.Store, data db.DetailCompanyRow) *pb.Company {
	var taxCode *string
	if data.TaxCode.Valid {
		taxCode = &data.TaxCode.String
	}

	var phone *string
	if data.Phone.Valid {
		phone = &data.Phone.String
	}

	var addressPb *pb.Address
	if data.Address.Valid {
		address, _ := store.GetAddress(ctx, data.Address.Int32)

		var provincePb *pb.AddressItem
		if address.Province.Valid {
			province, _ := store.GetProvinceByCode(ctx, address.Province.String)
			provincePb = &pb.AddressItem{
				Code:       province.Code,
				Name:       province.Name,
				NameEn:     province.NameEn,
				FullName:   province.FullName,
				FullNameEn: province.FullNameEn,
			}
		}

		var districtPb *pb.AddressItem
		if address.District.Valid {
			district, _ := store.GetDistrictByCode(ctx, address.District.String)
			districtPb = &pb.AddressItem{
				Code:       district.Code,
				Name:       district.Name,
				NameEn:     district.NameEn,
				FullName:   district.FullName,
				FullNameEn: district.FullNameEn,
			}
		}

		var wardPb *pb.AddressItem
		if address.Ward.Valid {
			ward, _ := store.GetWardByCode(ctx, address.Ward.String)
			wardPb = &pb.AddressItem{
				Code:       ward.Code,
				Name:       ward.Name,
				NameEn:     ward.NameEn,
				FullName:   ward.FullName,
				FullNameEn: ward.FullNameEn,
			}
		}

		addressPb = &pb.Address{
			Id:       int32(address.ID),
			Lat:      float32(address.Lat),
			Lng:      float32(address.Lng),
			Province: provincePb,
			District: districtPb,
			Ward:     wardPb,
			Title:    address.Title,
		}
	}

	var timeOpen *string
	if data.TimeOpen.Valid {
		x := data.TimeOpen.Time.Format("15:04:05")
		timeOpen = &x
	}

	var timeClose *string
	if data.TimeClose.Valid {
		x := data.TimeClose.Time.Format("15:04:05")
		timeClose = &x
	}

	totalEmployee, _ := store.CountEmployee(ctx, data.ID)

	return &pb.Company{
		Id:          int32(data.ID),
		Name:        data.Name,
		Code:        data.Code,
		Type:        data.Title,
		TaxCode:     taxCode,
		Phone:       phone,
		Description: nil,
		Address:     addressPb,
		Owner:       int32(data.Owner),
		Manager: &pb.Account{
			Id:       data.ID_2.Int32,
			FullName: data.FullName.String,
			Username: data.Username.String,
		},
		OaId:          &data.OaID.String,
		TimeOpen:      timeOpen,
		TimeClose:     timeClose,
		TotalEmployee: int32(totalEmployee),
		UserCreated: &pb.Account{
			Id:       data.ID_3.Int32,
			FullName: data.FullName_2.String,
			Username: data.Username_2.String,
		},
		UserUpdated: &pb.Account{
			Id:       data.ID_4.Int32,
			FullName: data.FullName_3.String,
			Username: data.Username_3.String,
		},
	}
}
