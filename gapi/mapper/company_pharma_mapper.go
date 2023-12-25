package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func CompanyPharmaMapper(data db.CompanyPharma) *pb.CompanyPharma {
	var code *string
	if data.Code.Valid {
		code = &data.Code.String
	}

	var country *string
	if data.Country.Valid {
		country = &data.Country.String
	}

	var address *string
	if data.Address.Valid {
		address = &data.Address.String
	}

	var cpType *string
	if data.CompanyPharmaType.Valid {
		cpType = &data.CompanyPharmaType.String
	}
	return &pb.CompanyPharma{
		Id:      int32(data.ID),
		Name:    data.Name,
		Code:    code,
		Country: country,
		Address: address,
		Type:    cpType,
	}
}
