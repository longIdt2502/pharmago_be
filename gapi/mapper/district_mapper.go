package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func DistrictMapper(district db.District) *pb.District {

	var administrativeUnit *int32
	if district.AdministrativeUnitID.Valid {
		administrativeUnitValue := district.AdministrativeUnitID.Int32
		administrativeUnit = &administrativeUnitValue
	}

	var provinceCode *string
	if district.ProvinceCode.Valid {
		provinceCode = &district.ProvinceCode.String
	}

	return &pb.District{
		Code:               district.Code,
		Name:               district.Name,
		NameEn:             district.NameEn,
		FullName:           district.FullName,
		FullNameEn:         district.FullNameEn,
		CodeName:           district.CodeName,
		ProvinceCode:       provinceCode,
		AdministrativeUnit: administrativeUnit,
	}
}
