package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func CategoryMapper(data db.ProductCategory) *pb.Category {

	return &pb.Category{
		Id:      data.ID,
		Name:    data.Name,
		Code:    data.Code,
		Company: data.Company,
	}
}
