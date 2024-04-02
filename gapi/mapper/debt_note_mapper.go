package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DebtNoteMapper(data db.DebtNote, items *[]db.DebtRepayment) *pb.DebtNote {
	var itemsPb []*pb.DebtRepayment
	if items != nil {
		for _, item := range *items {
			itemsPb = append(itemsPb, DebtRepaymentMapper(item))
		}
	}
	return &pb.DebtNote{
		Id:          data.ID,
		Code:        data.Code,
		Company:     data.Company,
		Title:       data.Title.String,
		Entity:      data.Entity,
		Money:       float32(data.Money),
		Paymented:   float32(data.Paymented),
		Note:        data.Note.String,
		Type:        data.Type,
		Status:      data.Status,
		UserCreated: data.UserCreated,
		Exprise:     timestamppb.New(data.Exprise),
		DabtNoteAt:  timestamppb.New(data.DabtNoteAt.Time),
		Repayments:  itemsPb,
	}
}

func DebtRepaymentMapper(data db.DebtRepayment) *pb.DebtRepayment {
	return &pb.DebtRepayment{
		Id:          data.ID,
		Code:        data.Code,
		Money:       float32(data.Money),
		Debt:        data.Debt,
		UserCreated: data.UserCreated,
		CreatedAt:   timestamppb.New(data.CreatedAt.Time),
	}
}