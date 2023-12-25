package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func ProductTypeMapper(data db.ProductType) *pb.ProductType {

	return &pb.ProductType{
		Id:      data.ID,
		Name:    data.Name,
		Code:    data.Code,
		Company: data.Company,
	}
}
