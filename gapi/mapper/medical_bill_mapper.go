package mapper

import (
	"context"

	"github.com/google/uuid"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MedicalBillMapper(ctx context.Context, store *db.Store, item db.GetListMedicalBillRow) *pb.MedicalBill {
	services, _ := store.GetListScheduleService(ctx, db.GetListScheduleServiceParams{
		MbUuid: uuid.NullUUID{
			UUID:  item.Uuid,
			Valid: true,
		},
	})
	var servicesPb []*pb.AppointmentScheduleService
	for _, item := range services {
		servicesPb = append(servicesPb, AppointmentScheduleServiceMapper(item))
	}

	// drugs, _ := store.GetListScheduleDrug(ctx, db.GetListScheduleDrugParams{
	// 	MbUuid: uuid.NullUUID{
	// 		UUID:  item.Uuid,
	// 		Valid: true,
	// 	},
	// })
	// var drugsPb []*pb.AppointmentScheduleDrug
	// for _, item := range drugs {
	// 	drugsPb = append(drugsPb, AppointmentScheduleDrugMapper(item))
	// }

	var paymentsPB []*pb.Payment
	paymentSell, _ := store.PaymentOrderByMedicalBill(ctx, item.Uuid)
	paymentsPB = append(paymentsPB, &pb.Payment{
		Code:     "SELL",
		MustPaid: float32(paymentSell.TotalMustPaid),
		HadPaid:  float32(paymentSell.TotalHadPaid),
		NeedPay:  float32(paymentSell.TotalNeedPay),
	})

	paymentService, _ := store.PaymentOrderServiceByMedicalBill(ctx, item.Uuid)
	paymentsPB = append(paymentsPB, &pb.Payment{
		Code:     "SERVICE",
		MustPaid: float32(paymentService.TotalMustPaid),
		HadPaid:  float32(paymentService.TotalHadPaid),
		NeedPay:  float32(paymentService.TotalNeedPay),
	})

	return &pb.MedicalBill{
		Id:            item.ID,
		Uuid:          item.Uuid.String(),
		Code:          item.Code,
		CustomerId:    item.Customer.Int32,
		Customer:      &pb.Account{Id: item.ID_2.Int32, FullName: item.FullName.String},
		Company:       item.Company.Int32,
		DoctorId:      item.ID_3,
		Doctor:        &pb.Account{Id: item.ID_3, FullName: item.FullName_2, Username: item.Username},
		Symptoms:      &item.Symptoms.String,
		Diagnostic:    &item.Diagnostic.String,
		QrCodeUrl:     &item.QrCodeUrl.String,
		IsDone:        item.IsDone,
		UserCreatedId: item.ID_4,
		UserCreated:   &pb.Account{Id: item.ID_4, FullName: item.FullName_3},
		UserUpdatedId: &item.ID_5.Int32,
		UserUpdated:   &pb.Account{Id: item.ID_5.Int32, FullName: item.FullName_4.String},
		MeetingAt:     timestamppb.New(item.MeetingAt),
		CreatedAt:     timestamppb.New(item.CreatedAt),
		UpdatedAt:     timestamppb.New(item.UpdatedAt.Time),
		Services:      servicesPb,
		Urls:          []*pb.AppointmentScheduleUrl{},
		Payment:       paymentsPB,
	}
}
