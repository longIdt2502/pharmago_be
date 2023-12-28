package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/skip2/go-qrcode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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

	ticketType, err := server.store.GetTicketType(ctx, db.GetTicketTypeParams{
		ID: sql.NullInt32{},
		Code: sql.NullString{
			String: req.Ticket.GetType(),
			Valid:  true,
		},
	})

	ticketStatus, err := server.store.GetTicketStatus(ctx, db.GetTicketStatusParams{
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
		return nil, status.Errorf(codes.Internal, "failed to create qr code: %w", err)
	}
	file, _ := utils.NewFileFromImage(png)
	_, err = server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save qr code: %w", err)
	}
	urlQr, _ := server.b2Bucket.FileURL(file.Name)

	qr, err := server.store.CreateMedia(ctx, urlQr)

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
			Int32: req.Ticket.GetExportTo(),
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
		return nil, status.Errorf(codes.InvalidArgument, "cannot record ticket: %w", err)
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
			return nil, status.Errorf(codes.InvalidArgument, "cannot record consignment: %w", err)
		}
	}

	return &pb.TicketCreateResponse{
		Code:    200,
		Message: "success",
		Details: ticket.ID,
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
			return nil, status.Errorf(codes.NotFound, "id ticket not exists: ", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to change status ticket: ", err)
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
		Company:   req.GetCompany(),
		Warehouse: req.GetWarehouse(),
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
		return nil, status.Errorf(codes.Internal, "failed to get list consignment: %w", err)
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
		})

		consignmentsPb = append(consignmentsPb, data)
	}

	return &pb.ConsignmentListResponse{
		Code:    200,
		Message: "success",
		Details: consignmentsPb,
	}, nil
}
