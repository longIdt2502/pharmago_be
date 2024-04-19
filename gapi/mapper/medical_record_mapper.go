package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MedicalRecordMapper(item db.MedicalRecord) *pb.MedicalRecord {

	var weight *float32
	if item.Weight.Valid {
		w := float32(item.Weight.Float64)
		weight = &w
	}

	var long *float32
	if item.Long.Valid {
		w := float32(item.Long.Float64)
		long = &w
	}

	return &pb.MedicalRecord{
		Id:            item.ID,
		Code:          item.Code,
		Customer:      item.Customer,
		Weight:        weight,
		Long:          long,
		Symptom:       item.Symptom,
		Diagnostic:    item.Diagnostic,
		Result:        item.Result,
		ReExamination: &item.ReExamination,
		Note:          item.Note.String,
		CreatedAt:     timestamppb.New(item.CreatedAt),
		UpdatedAt:     timestamppb.New(item.UpdatedAt.Time),
	}
}
