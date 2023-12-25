package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func BrandMapper(data db.ProductBrand) *pb.Brand {

	return &pb.Brand{
		Id:      data.ID,
		Name:    data.Name,
		Code:    data.Code,
		Company: data.Company,
	}
}
