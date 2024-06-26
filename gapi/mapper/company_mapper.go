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

	return &pb.Company{
		Id:          int32(data.ID),
		Name:        data.Name,
		Code:        data.Code,
		Type:        data.Type,
		TaxCode:     taxCode,
		Phone:       phone,
		Description: nil,
		Address:     addressPb,
		Owner:       int32(data.Owner),
		OaId:        &data.OaID.String,
	}
}
