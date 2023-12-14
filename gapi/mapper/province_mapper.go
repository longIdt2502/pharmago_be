package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func ProvinceMapper(province db.Province) *pb.Province {

	var administrativeUnit *int32
	if province.AdministrativeUnitID.Valid {
		administrativeUnitValue := int32(province.AdministrativeUnitID.Int64)
		administrativeUnit = &administrativeUnitValue
	}

	var administrativeRegion *int32
	if province.AdministrativeRegionID.Valid {
		administrativeRegionValue := int32(province.AdministrativeRegionID.Int64)
		administrativeRegion = &administrativeRegionValue
	}

	return &pb.Province{
		Code:                 province.Code,
		Name:                 province.Name,
		NameEn:               province.NameEn,
		FullName:             province.FullName,
		FullNameEn:           province.FullNameEn,
		CodeName:             province.CodeName,
		AdministrativeUnit:   administrativeUnit,
		AdministrativeRegion: administrativeRegion,
	}
}
