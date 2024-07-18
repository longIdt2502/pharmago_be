package mapper

import (
	"context"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AppointmentScheduleMapper(ctx context.Context, store *db.Store, item db.GetListScheduleRow) *pb.AppointmentSchedule {
	services, _ := store.GetListScheduleService(ctx, item.Uuid)
	var servicesPb []*pb.AppointmentScheduleService
	for _, item := range services {
		servicesPb = append(servicesPb, AppointmentScheduleServiceMapper(item))
	}

	drugs, _ := store.GetListScheduleDrug(ctx, item.Uuid)
	var drugsPb []*pb.AppointmentScheduleDrug
	for _, item := range drugs {
		drugsPb = append(drugsPb, AppointmentScheduleDrugMapper(item))
	}

	return &pb.AppointmentSchedule{
		Id:         item.ID,
		Uuid:       item.Uuid.String(),
		Code:       item.Code,
		CustomerId: item.Customer.Int32,
		Customer: &pb.Account{
			Id:       item.ID_2.Int32,
			FullName: item.FullName.String,
		},
		Company:  item.Company.Int32,
		DoctorId: item.ID_3,
		Doctor: &pb.Account{
			Id:       item.ID_3,
			FullName: item.FullName_2,
		},
		Symptoms:      &item.Symptoms.String,
		Diagnostic:    &item.Diagnostic.String,
		QrCodeUrl:     &item.QrCodeUrl.String,
		IsDone:        item.IsDone,
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
		MeetingAt: timestamppb.New(item.MeetingAt),
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt.Time),
		Services:  servicesPb,
		Urls:      []*pb.AppointmentScheduleUrl{},
		Drugs:     drugsPb,
	}
}

func AppointmentScheduleServiceMapper(item db.GetListScheduleServiceRow) *pb.AppointmentScheduleService {
	var orderId *int32
	if item.OrderService.Valid {
		orderId = &item.OrderService.Int32
	}

	return &pb.AppointmentScheduleService{
		Id:     item.ID,
		AsUuid: item.AsUuid.String(),
		Service: &pb.Service{
			Id:           item.ID_2,
			Code:         item.Code,
			Title:        item.Title,
			Entity:       &item.Entity.String,
			Frequency:    &item.Frequency.String,
			Unit:         item.Unit,
			Price:        float32(item.Price),
			Description:  &item.Description.String,
			Company:      item.Company,
			ReminderTime: &item.ReminderTime.Int32,
		},
		ServiceId:    item.Service.Int32,
		OrderService: &pb.Order{},
		OrderId:      orderId,
	}
}

func AppointmentScheduleDrugMapper(item db.GetListScheduleDrugRow) *pb.AppointmentScheduleDrug {
	return &pb.AppointmentScheduleDrug{
		Id:        item.ID,
		AsUuid:    item.AsUuid.String(),
		VariantId: item.Variant.Int32,
		Variant: &pb.Variant{
			Id:               item.ID_2,
			Code:             item.Code,
			Name:             item.Name,
			Product:          item.Product,
			Media:            item.MediaUrl.String,
			InitialInventory: item.InitialInventory,
			RealInventory:    item.RealInventory,
		},
		LieuDung: &item.LieuDung.String,
		Quantity: item.Quantity,
	}
}
