package mapper

import (
	"context"
	"database/sql"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TicketMapper(data db.GetListTicketRow) *pb.TicketPreview {
	note := ""
	if data.Note.Valid {
		note = data.Note.String
	}

	return &pb.TicketPreview{
		Id:   data.ID,
		Code: data.Code,
		Type: &pb.SimpleData{
			Id:   data.TtID,
			Name: data.TtTitle,
			Code: data.TtCode,
		},
		Status: &pb.SimpleData{
			Id:   data.TsID,
			Name: data.TsTitle,
			Code: data.TsCode,
		},
		Note:           note,
		Qr:             data.QrUrl,
		TotalItems:     data.TotalProducts,
		TotalItemsType: 0,
		TotalPrice:     float32(data.TotalPrice),
		WarehouseName:  data.WName,
		UserCreated:    data.AFullName,
		CreatedAt:      timestamppb.New(data.CreatedAt),
	}
}

func TicketDetailMapper(ctx context.Context, store *db.Store, data db.GetDetailTicketRow) *pb.Ticket {

	var supplier *pb.Supplier
	if data.SID.Valid {
		addressDb, _ := store.GetAddress(ctx, data.SAddress.Int32)
		addressPb := AddressMapper(ctx, store, addressDb)
		supplier = &pb.Supplier{
			Id:         data.SID.Int32,
			Code:       data.SCode.String,
			Name:       data.SName.String,
			DeputyName: data.SDeputy.String,
			Phone:      data.SPhone.String,
			Email:      data.SEmail.String,
			Address:    addressPb,
			Company:    data.SCompany.Int32,
		}
	}

	var customer *pb.Customer
	if data.CID.Valid {
		addressDb, _ := store.GetAddress(ctx, data.CAddress.Int32)
		addressPb := AddressMapper(ctx, store, addressDb)
		var email *string
		if data.CEmail.Valid {
			email = &data.CEmail.String
		}
		customer = &pb.Customer{
			Id:       data.CID.Int32,
			Code:     data.CCode.String,
			FullName: data.CName.String,
			Company:  data.CCompany.Int32,
			Address:  addressPb,
			Phone:    data.CPhone.String,
			Email:    email,
		}
	}

	var warehouse *pb.Warehouse
	warehouseDb, _ := store.GetWarehouse(ctx, data.Warehouse)
	warehouse, _ = WarehouseMapper(ctx, store, warehouseDb)

	var consignments []*pb.Consignment
	consignmentsDb, _ := store.GetItemsTicket(ctx, sql.NullInt32{
		Int32: data.ID,
		Valid: true,
	})
	for _, value := range consignmentsDb {
		dataPb := ConsignmentMapper(ctx, store, value, warehouseDb.Companies.Int32)
		consignments = append(consignments, dataPb)
	}

	return &pb.Ticket{
		Id:   data.ID,
		Code: data.Code,
		Type: &pb.SimpleData{
			Id:   data.TtID,
			Name: data.TtTitle,
			Code: data.TtCode,
		},
		Status: &pb.SimpleData{
			Id:   data.TsID,
			Name: data.TsTitle,
			Code: data.TsCode,
		},
		Note:         data.Note.String,
		Qr:           data.QrUrl,
		TotalPrice:   float32(data.TotalPrice),
		Supplier:     supplier,
		Customer:     customer,
		Consignments: consignments,
		Warehouse:    warehouse,
		UserCreated:  data.UserCreatedName,
		UserUpdated:  data.UserUpdatedName,
		CreatedAt:    timestamppb.New(data.CreatedAt),
	}
}
