package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func WardMapper(ward db.Ward) *pb.Ward {

	var administrativeUnit *int32
	if ward.AdministrativeUnitID.Valid {
		administrativeUnitValue := ward.AdministrativeUnitID.Int32
		administrativeUnit = &administrativeUnitValue
	}

	var provinceCode *string
	if ward.DistrictCode.Valid {
		provinceCode = &ward.DistrictCode.String
	}

	return &pb.Ward{
		Code:               ward.Code,
		Name:               ward.Name,
		NameEn:             ward.NameEn,
		FullName:           ward.FullName,
		FullNameEn:         ward.FullNameEn,
		CodeName:           ward.CodeName,
		DistrictCode:       provinceCode,
		AdministrativeUnit: administrativeUnit,
	}
}
