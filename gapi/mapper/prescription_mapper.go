package mapper

import (
	"context"

	"github.com/google/uuid"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PrescriptionMapper(ctx context.Context, store *db.Store, item db.DetailPrescriptionRow) *pb.Prescription {
	var itemsPb []*pb.PrescriptionItem

	itemsDb, _ := store.ListPrescriptionItem(ctx, uuid.NullUUID{UUID: item.Uuid, Valid: true})
	if len(itemsDb) != 0 {
		for _, item := range itemsDb {
			itemsPb = append(itemsPb, &pb.PrescriptionItem{
				Id:        item.ID,
				VariantId: item.Variant.Int32,
				Variant: &pb.Variant{
					Name: item.Name,
				},
				LieuDung: &item.LieuDung.String,
				Quantity: item.Quantity,
			})
		}
	}

	return &pb.Prescription{
		Id:         item.ID,
		Uuid:       item.Uuid.String(),
		Code:       item.Code,
		CustomerId: &item.ID_2,
		Customer: &pb.Account{
			Id:       item.ID_2,
			FullName: item.FullName,
			Username: item.Phone.String,
		},
		Company:  item.Company,
		DoctorId: item.ID,
		Doctor: &pb.Account{
			Id:       item.ID_3,
			FullName: item.FullName_2,
		},
		Symptoms:      &item.Symptoms.String,
		Diagnostic:    &item.Diagnostic.String,
		Items:         itemsPb,
		Payment:       []*pb.Payment{},
		UserCreatedId: item.ID_4,
		UserCreated: &pb.Account{
			Id:       item.ID_4,
			FullName: item.FullName_3,
		},
		UserUpdatedId: &item.ID_5.Int32,
		UserUpdated: &pb.Account{
			Id:       item.ID_5.Int32,
			FullName: item.FullName_4.String,
		},
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt.Time),
	}
}

func PrescriptionListItemMapper(ctx context.Context, store *db.Store, item db.ListPrescriptionRow) *pb.Prescription {
	var itemsPb []*pb.PrescriptionItem

	itemsDb, _ := store.ListPrescriptionItem(ctx, uuid.NullUUID{UUID: item.Uuid, Valid: true})
	if len(itemsDb) != 0 {
		for _, item := range itemsDb {
			itemsPb = append(itemsPb, &pb.PrescriptionItem{
				Id:        item.ID,
				VariantId: item.Variant.Int32,
				Variant: &pb.Variant{
					Name: item.Name,
				},
				LieuDung: &item.LieuDung.String,
				Quantity: item.Quantity,
			})
		}
	}

	return &pb.Prescription{
		Id:         item.ID,
		Uuid:       item.Uuid.String(),
		Code:       item.Code,
		CustomerId: &item.ID_2,
		Customer: &pb.Account{
			Id:       item.ID_2,
			FullName: item.FullName,
		},
		Company:  item.Company,
		DoctorId: item.ID,
		Doctor: &pb.Account{
			Id:       item.ID_3,
			FullName: item.FullName_2,
		},
		Symptoms:      &item.Symptoms.String,
		Diagnostic:    &item.Diagnostic.String,
		Items:         itemsPb,
		Payment:       []*pb.Payment{},
		UserCreatedId: item.ID_4,
		UserCreated: &pb.Account{
			Id:       item.ID_4,
			FullName: item.FullName_3,
		},
		UserUpdatedId: &item.ID_5.Int32,
		UserUpdated: &pb.Account{
			Id:       item.ID_5.Int32,
			FullName: item.FullName_4.String,
		},
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt.Time),
	}
}
