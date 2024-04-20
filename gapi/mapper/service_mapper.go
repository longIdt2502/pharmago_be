package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
)

func ServiceMapper(item db.Service) *pb.Service {

	var time *int32
	if item.ReminderTime.Valid {
		time = &(item.ReminderTime.Int32)
	}

	return &pb.Service{
		Id:           item.ID,
		Code:         item.Code,
		Title:        item.Title,
		Entity:       &item.Entity.String,
		Frequency:    &item.Frequency.String,
		Unit:         item.Unit,
		Price:        float32(item.Price),
		Description:  &item.Description.String,
		Company:      item.Company,
		ReminderTime: time,
	}
}
