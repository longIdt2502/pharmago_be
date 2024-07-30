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
		Id:             item.ID,
		Code:           item.Code,
		Title:          item.Title,
		Entity:         &item.Entity.String,
		Staff:          &pb.Account{},
		Frequency:      &item.Frequency.String,
		Unit:           item.Unit,
		Price:          float32(item.Price),
		Description:    &item.Description.String,
		Company:        item.Company,
		Variants:       []*pb.Variant{},
		ReminderTime:   time,
		Used:           0,
		ChiDinh:        &item.ChiDinh.String,
		ChongChiDinh:   &item.ChongChiDinh.String,
		CongDung:       &item.CongDung.String,
		Caution:        &item.Caution.String,
		HinhThuc:       &item.HinhThuc.String,
		TacDungPhu:     &item.TacDungPhu.String,
		NumberRegister: &item.NumberRegister.String,
		NumberDecision: &item.NumberDecision.String,
		CongTyDk:       &item.CongTyDk.String,
		Message:        &item.Message.String,
		Brand:          &pb.SimpleData{},
		ActionTime:     &item.ActionTime.String,
	}
}

func ServiceDetailRowMapper(item db.DetailServiceRow) *pb.Service {

	var time *int32
	if item.ReminderTime.Valid {
		time = &(item.ReminderTime.Int32)
	}

	return &pb.Service{
		Id:     item.ID,
		Code:   item.Code,
		Title:  item.Title,
		Entity: &item.Entity.String,
		Staff: &pb.Account{
			Id:       item.ID_2.Int32,
			Username: item.Username.String,
			FullName: item.FullName.String,
		},
		Frequency:      &item.Frequency.String,
		Unit:           item.Unit,
		Price:          float32(item.Price),
		Description:    &item.Description.String,
		Company:        item.Company,
		Variants:       []*pb.Variant{},
		ReminderTime:   time,
		Used:           0,
		ChiDinh:        &item.ChiDinh.String,
		ChongChiDinh:   &item.ChongChiDinh.String,
		CongDung:       &item.CongDung.String,
		Caution:        &item.Caution.String,
		HinhThuc:       &item.HinhThuc.String,
		TacDungPhu:     &item.TacDungPhu.String,
		NumberRegister: &item.NumberRegister.String,
		NumberDecision: &item.NumberDecision.String,
		CongTyDk:       &item.CongTyDk.String,
		Message:        &item.Message.String,
		Brand: &pb.SimpleData{
			Id:   item.ID_3.Int32,
			Name: item.Name.String,
			Code: item.Code_2.String,
		},
		ActionTime: &item.ActionTime.String,
	}
}

func ServiceGetListServiceRowMapper(item db.GetListServiceRow) *pb.Service {

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
		Used:         int32(item.QuantityUse.Int64),
	}
}

func ServiceByCustomerMapper(item db.GetServicesByCustomerRow) *pb.Service {

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
		Used:         int32(item.QuantityUse),
	}
}
