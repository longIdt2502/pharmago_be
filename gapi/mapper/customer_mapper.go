package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func CustomerMapper(data db.Customer) *pb.Customer {
	return &pb.Customer{
		Id:       data.ID,
		Code:     data.Code,
		FullName: data.FullName,
		Company:  data.Company,
		Address:  data.Address.Int32,
		Phone:    data.Phone.String,
		Email:    nil,
	}
}
