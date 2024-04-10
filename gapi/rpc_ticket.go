package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/skip2/go-qrcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) TicketCreate(ctx context.Context, req *pb.TicketCreateRequest) (*pb.TicketCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	account, err := server.store.GetAccountByUseName(ctx, tokenPayload.Username)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	ticketType, _ := server.store.GetTicketType(ctx, db.GetTicketTypeParams{
		ID: sql.NullInt32{},
		Code: sql.NullString{
			String: req.Ticket.GetType(),
			Valid:  true,
		},
	})

	ticketStatus, _ := server.store.GetTicketStatus(ctx, db.GetTicketStatusParams{
		ID: sql.NullInt32{},
		Code: sql.NullString{
			String: req.Ticket.GetStatus(),
			Valid:  true,
		},
	})

	codeTicket := req.Ticket.GetCode()
	if req.Ticket.Code == nil {
		codeTicket = fmt.Sprintf("TICKET-%s", utils.RandomString(6))
	}

	var png []byte
	png, err = qrcode.Encode(codeTicket, qrcode.Medium, 256)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create qr code: %e", err)
	}
	file, _ := utils.NewFileFromImage(png)
	_, err = server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save qr code: %e", err)
	}
	urlQr, _ := server.b2Bucket.FileURL(file.Name)

	qr, _ := server.store.CreateMedia(ctx, urlQr)

	var idAddressExportTo int32
	if req.Ticket.ExportTo == nil {
		warehouse, err := server.store.GetWarehouse(ctx, req.Ticket.GetWarehouse())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "warehouse not exists: %e", err)
			}
			return nil, status.Errorf(codes.Internal, "fail to get warehouse: %e", err)
		}
		if !warehouse.Address.Valid {
			return nil, status.Errorf(codes.InvalidArgument, "warehouse address not exists")
		}
		idAddressExportTo = warehouse.Address.Int32
	} else {
		idAddressExportTo = req.Ticket.GetExportTo()
	}

	ticket, err := server.store.CreateTicket(ctx, db.CreateTicketParams{
		Code: codeTicket,
		Type: sql.NullInt32{
			Int32: ticketType.ID,
			Valid: true,
		},
		Status: sql.NullInt32{
			Int32: ticketStatus.ID,
			Valid: true,
		},
		Note: sql.NullString{
			String: req.Ticket.GetNote(),
			Valid:  req.Ticket.Note != "",
		},
		Qr: sql.NullInt32{
			Int32: qr.ID,
			Valid: true,
		},
		ExportTo: sql.NullInt32{
			Int32: idAddressExportTo,
			Valid: true,
		},
		ImportFrom: sql.NullInt32{
			Int32: req.Ticket.GetImportFrom(),
			Valid: true,
		},
		TotalPrice:  float64(req.Ticket.GetTotalPrice()),
		Warehouse:   req.Ticket.GetWarehouse(),
		UserCreated: account.ID,
		UserUpdated: sql.NullInt32{
			Int32: account.ID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot record ticket: %e", err)
	}

	for _, value := range req.Consignment {
		code := value.GetCode()
		if value.Code == nil {
			code = fmt.Sprintf("CONSIGMENT-%s", utils.RandomString(6))
		}
		_, err = server.store.CreateConsignment(ctx, db.CreateConsignmentParams{
			Code:      code,
			Quantity:  value.GetQuantity(),
			Inventory: value.GetQuantity(),
			Ticket: sql.NullInt32{
				Int32: ticket.ID,
				Valid: true,
			},
			Variant: sql.NullInt32{
				Int32: value.GetVariant(),
				Valid: true,
			},
			ExpiredAt:   time.Unix(value.ExpiredAt.Seconds, int64(value.ExpiredAt.Nanos)),
			ProductedAt: time.Unix(value.ProducedAt.Seconds, int64(value.ProducedAt.Nanos)),
			UserCreated: sql.NullInt32{
				Int32: account.ID,
				Valid: true,
			},
			UserUpdated: sql.NullInt32{
				Int32: account.ID,
				Valid: true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot record consignment: %e", err)
		}
	}

	return &pb.TicketCreateResponse{
		Code:    200,
		Message: "success",
		Details: ticket.ID,
	}, nil
}

func (server *ServerGRPC) TicketList(ctx context.Context, req *pb.TicketListRequest) (*pb.TicketListResponse, error) {
	tickets, err := server.store.GetListTicket(ctx, db.GetListTicketParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
		Status: sql.NullString{
			String: req.GetStatus(),
			Valid:  req.Status != nil,
		},
		Type: sql.NullString{
			String: req.GetType(),
			Valid:  req.Type != nil,
		},
		Supplier: sql.NullInt32{
			Int32: req.GetSupplier(),
			Valid: req.Supplier != nil,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list tickets: %e", err)
	}

	var ticketsPb []*pb.TicketPreview
	for _, value := range tickets {
		dataPb := mapper.TicketMapper(value)
		ticketsPb = append(ticketsPb, dataPb)
	}

	newCount, _ := server.store.CountTicketByStatus(ctx, "NEW")
	processCount, _ := server.store.CountTicketByStatus(ctx, "IN_PROCESS")
	completeCount, _ := server.store.CountTicketByStatus(ctx, "COMPLETE")
	cancelCount, _ := server.store.CountTicketByStatus(ctx, "CANCEL")

	return &pb.TicketListResponse{
		Code:    200,
		Message: "success",
		Details: ticketsPb,
		Count: &pb.TicketListResponseCount{
			New:       int32(newCount),
			InProcess: int32(processCount),
			Complete:  int32(completeCount),
			Cancel:    int32(cancelCount),
		},
	}, nil
}

func (server *ServerGRPC) TicketDetail(ctx context.Context, req *pb.TicketDetailRequest) (*pb.TicketDetailResponse, error) {
	ticket, err := server.store.GetDetailTicket(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "ticket not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get ticket: %e", err)
	}

	ticketPb := mapper.TicketDetailMapper(ctx, server.store, ticket)
	return &pb.TicketDetailResponse{
		Code:    200,
		Message: "success",
		Details: ticketPb,
	}, nil
}

func (server *ServerGRPC) TicketUpdateStatus(ctx context.Context, req *pb.TicketUpdateStatusRequest) (*pb.TicketUpdateStatusResponse, error) {
	statusTicket, err := server.store.GetTicketStatus(ctx, db.GetTicketStatusParams{
		Code: sql.NullString{
			String: req.Status.String(),
			Valid:  true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get status ticket")
	}

	_, err = server.store.UpdateTicketStatus(ctx, db.UpdateTicketStatusParams{
		Status: sql.NullInt32{
			Int32: statusTicket.ID,
			Valid: true,
		},
		ID: req.Id,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "id ticket not exists: %e", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to change status ticket: %e", err)
	}

	if statusTicket.Code == "COMPLETE" {
		_, err = server.store.UpdateConsignmentByTicket(ctx, sql.NullInt32{
			Int32: req.Id,
			Valid: true,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to change status consignment")
		}
	}

	return &pb.TicketUpdateStatusResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) ConsignmentList(ctx context.Context, req *pb.ConsignmentListRequest) (*pb.ConsignmentListResponse, error) {
	consignments, err := server.store.GetConsignments(ctx, db.GetConsignmentsParams{
		Company: req.GetCompany(),
		Warehouse: sql.NullInt32{
			Int32: req.GetWarehouse(),
			Valid: req.Warehouse != nil,
		},
		Available: sql.NullBool{
			Bool:  req.GetAvailable(),
			Valid: req.Available != nil,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list consignment: %e", err)
	}

	var consignmentsPb []*pb.Consignment
	for _, value := range consignments {

		data := mapper.ConsignmentMapper(ctx, server.store, db.Consignment{
			ID:          value.ID,
			Code:        value.Code,
			Quantity:    value.Quantity,
			Inventory:   value.Inventory,
			Ticket:      value.Ticket,
			ExpiredAt:   value.ExpiredAt,
			ProductedAt: value.ProductedAt,
			IsAvailable: value.IsAvailable,
			UserCreated: value.UserCreated,
			UserUpdated: value.UserUpdated,
			UpdatedAt:   value.UpdatedAt,
			CreatedAt:   value.CreatedAt,
			Variant:     value.Variant,
		}, req.GetCompany())

		consignmentsPb = append(consignmentsPb, data)
	}

	return &pb.ConsignmentListResponse{
		Code:    200,
		Message: "success",
		Details: consignmentsPb,
	}, nil
}
