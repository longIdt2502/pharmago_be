package mapper

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func AddressMapper(ctx context.Context, store *db.Store, data db.Address) *pb.Address {

	var provincePb *pb.AddressItem
	if data.Province.Valid {
		province, _ := store.GetProvinceByCode(ctx, data.Province.String)
		provincePb = &pb.AddressItem{
			Code:       province.Code,
			Name:       province.Name,
			NameEn:     province.NameEn,
			FullName:   province.FullName,
			FullNameEn: province.FullNameEn,
		}
	}

	var districtPb *pb.AddressItem
	if data.District.Valid {
		district, _ := store.GetDistrictByCode(ctx, data.District.String)
		districtPb = &pb.AddressItem{
			Code:       district.Code,
			Name:       district.Name,
			NameEn:     district.NameEn,
			FullName:   district.FullName,
			FullNameEn: district.FullNameEn,
		}
	}

	var wardPb *pb.AddressItem
	if data.Ward.Valid {
		ward, _ := store.GetWardByCode(ctx, data.Ward.String)
		wardPb = &pb.AddressItem{
			Code:       ward.Code,
			Name:       ward.Name,
			NameEn:     ward.NameEn,
			FullName:   ward.FullName,
			FullNameEn: ward.FullNameEn,
		}
	}

	addressPb := &pb.Address{
		Id:       data.ID,
		Lat:      float32(data.Lat),
		Lng:      float32(data.Lng),
		Province: provincePb,
		District: districtPb,
		Ward:     wardPb,
		Title:    data.Title,
	}

	return addressPb
}
