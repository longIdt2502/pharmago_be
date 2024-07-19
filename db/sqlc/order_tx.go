package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/longIdt2502/pharmago_be/b2"
	"github.com/longIdt2502/pharmago_be/common"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/skip2/go-qrcode"
)

type CreateOrderTxParams struct {
	*pb.OrderCreateRequest
	B2Bucket     *b2.B2Bucket
	TokenPayload *token.Payload
}

type CreateOrderTxResult struct {
	Id int32
}

func (store *Store) CreateOrderTx(ctx context.Context, req CreateOrderTxParams) (CreateOrderTxResult, error) {
	var result CreateOrderTxResult

	opts := &sql.TxOptions{
		Isolation: 1,
		ReadOnly:  false,
	}

	err := store.execTx(ctx, opts, func(q *Queries) error {
		var err error

		var addressOrder int32
		if req.Order.Customer != nil {
			customer, err := q.GetCustomer(ctx, req.Order.GetCustomer())
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return common.ErrDBWithMsg(err, "Khách hàng chưa có trong hệ thống")
				}
				return common.ErrDB(err)
			}
			if !customer.Address.Valid {
				return common.ErrDBWithMsg(err, "Khách hàng chưa có địa chỉ")
			}
			addressOrder = customer.Address.Int32
		}
		// else {
		// 	// TODO:
		// }

		warehouse, err := q.GetWarehouse(ctx, req.GetWarehouse())
		if err != nil {
			return common.ErrDB(err)
		}

		paymentCode := fmt.Sprintf("PM-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
		payment, err := q.CreatePayment(ctx, CreatePaymentParams{
			Code:     paymentCode,
			MustPaid: float64(req.Payment.GetMustPaid()),
			HadPaid:  float64(req.Payment.GetHadPaid()),
			NeedPay:  float64(req.Payment.GetNeedPay()),
		})
		if err != nil {
			return common.ErrDBWithMsg(err, "Lỗi tạo thanh toán cho đơn hàng")
		}

		for _, value := range req.PaymentItems {
			_, err = q.CreatePaymentItem(ctx, CreatePaymentItemParams{
				Type:    value.GetType(),
				Value:   float64(value.GetValue()),
				IsPaid:  value.GetIsPaid(),
				Payment: payment.ID,
				ExtraNote: sql.NullString{
					String: value.GetExtraNote(),
					Valid:  value.ExtraNote != nil,
				},
			})
			if err != nil {
				return common.ErrDBWithMsg(err, "Lỗi tạo thanh toán chi tiết")
			}
		}

		ticketType, _ := q.GetTicketType(ctx, GetTicketTypeParams{
			Code: sql.NullString{
				String: "EXPORT",
				Valid:  true,
			},
		})
		ticketStatus, _ := q.GetTicketStatus(ctx, GetTicketStatusParams{
			Code: sql.NullString{
				String: "NEW",
				Valid:  true,
			},
		})

		codeTicket := fmt.Sprintf("TICKET-%s", utils.RandomString(6))
		var png []byte
		png, err = qrcode.Encode(codeTicket, qrcode.Medium, 256)
		if err != nil {
			return common.ErrInternalWithMsg(err, "Lỗi tạo qrcode phiếu xuất hàng")
		}
		file, _ := utils.NewFileFromImage(png)
		_, err = req.B2Bucket.UploadFile(file.Name, file.Meta, file.File)
		if err != nil {
			return common.ErrInternal(err)
		}
		urlQr, _ := req.B2Bucket.FileURL(file.Name)

		qr, _ := q.CreateMedia(ctx, urlQr)

		ticket, err := q.CreateTicket(ctx, CreateTicketParams{
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
				String: "Phiếu xuất hàng cho đơn hàng",
				Valid:  true,
			},
			Qr: sql.NullInt32{
				Int32: qr.ID,
				Valid: true,
			},
			ExportTo: sql.NullInt32{
				Int32: addressOrder,
				Valid: addressOrder != 0,
			},
			ImportFrom: sql.NullInt32{
				Int32: warehouse.Address.Int32,
				Valid: true,
			},
			TotalPrice:  float64(req.Order.GetTotalPrice()),
			Warehouse:   req.GetWarehouse(),
			UserCreated: req.TokenPayload.UserID,
			UserUpdated: sql.NullInt32{
				Int32: req.TokenPayload.UserID,
				Valid: true,
			},
		})
		if err != nil {
			return common.ErrDBWithMsg(err, "Lỗi tạo phiếu xuất hàng")
		}

		var orderCode string
		if req.Order.Code == nil {
			orderCode = fmt.Sprintf("ORDER-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
		} else {
			orderCode = req.Order.GetCode()
		}

		var pngOrder []byte
		pngOrder, err = qrcode.Encode(orderCode, qrcode.Medium, 256)
		if err != nil {
			return common.ErrInternalWithMsg(err, "Lỗi tạo qrcode phiếu xuất hàng đơn hàng")
		}
		fileOrder, _ := utils.NewFileFromImage(pngOrder)
		_, err = req.B2Bucket.UploadFile(fileOrder.Name, fileOrder.Meta, fileOrder.File)
		if err != nil {
			return common.ErrInternal(err)
		}
		urlQrOrder, _ := req.B2Bucket.FileURL(fileOrder.Name)

		qrOrder, _ := q.CreateMedia(ctx, urlQrOrder)

		order, err := q.CreateOrder(ctx, CreateOrderParams{
			Code:       orderCode,
			TotalPrice: float64(req.Order.GetTotalPrice()),
			Description: sql.NullString{
				String: req.Order.GetDescription(),
				Valid:  req.Order.Description != nil,
			},
			Vat:          float64(req.Order.GetVat()),
			Discount:     req.Order.GetDiscount(),
			ServicePrice: float64(req.Order.GetServicePrice()),
			MustPaid:     float64(req.Order.GetMustPaid()),
			Customer: sql.NullInt32{
				Int32: req.Order.GetCustomer(),
				Valid: req.Order.Customer != nil,
			},
			Address: sql.NullInt32{
				Int32: addressOrder,
				Valid: req.Order.Customer != nil,
			},
			Status: sql.NullString{
				String: req.Order.GetStatus(),
				Valid:  true,
			},
			Type: sql.NullString{
				String: req.Order.GetType(),
				Valid:  true,
			},
			Ticket: sql.NullInt32{
				Int32: ticket.ID,
				Valid: true,
			},
			Qr: sql.NullInt32{
				Int32: qrOrder.ID,
				Valid: true,
			},
			Company: req.Order.GetCompany(),
			Payment: payment.ID,
			UserCreated: sql.NullInt32{
				Int32: req.TokenPayload.UserID,
				Valid: true,
			},
			UserUpdated: sql.NullInt32{
				Int32: req.TokenPayload.UserID,
				Valid: true,
			},
		})
		if err != nil {
			return common.ErrDBWithMsg(err, "lỗi tạo đơn hàng")
		}

		for _, value := range req.GetOrderItems() {
			// var consignmentLog ConsignmentLog
			// var consignment Consignment
			// if value.Consignment != nil {
			// 	consignment, err = q.GetConsignment(ctx, GetConsignmentParams{
			// 		ID: value.GetConsignment(),
			// 		Variant: sql.NullInt32{
			// 			Int32: value.GetVariant(),
			// 			Valid: true,
			// 		},
			// 	})
			// 	if err != nil {
			// 		return common.ErrDBWithMsg(err, fmt.Sprintf("Không thể lấy lô hàng sản phẩm phù hợp: id = %d", value.Variant))
			// 	}
			// } else {
			// 	consignment, err = q.SuggestConsignmentForVariant(ctx, SuggestConsignmentForVariantParams{
			// 		Variant: sql.NullInt32{
			// 			Int32: value.GetVariant(),
			// 			Valid: true,
			// 		},
			// 		Inventory: value.GetValue(),
			// 	})
			// 	if err != nil {
			// 		return common.ErrDBWithMsg(err, fmt.Sprintf("Không thể lấy lô hàng sản phẩm phù hợp: id = %d", value.Variant))
			// 	}
			// }

			// var amount int32
			// if req.Order.Type == "SELL" {
			// 	amount = -value.GetValue()
			// } else {
			// 	amount = value.GetValue()
			// }
			// _, _ = q.UpdateConsignment(ctx, UpdateConsignmentParams{
			// 	Amount: amount,
			// 	ID:     consignment.ID,
			// })

			// consignmentLog, err = q.CreateConsignmentLog(ctx, CreateConsignmentLogParams{
			// 	Consignment:  consignment.ID,
			// 	Inventory:    consignment.Inventory,
			// 	AmountChange: amount,
			// 	UserCreated: sql.NullInt32{
			// 		Int32: req.TokenPayload.UserID,
			// 		Valid: true,
			// 	},
			// })
			// if err != nil {
			// 	return common.ErrDBWithMsg(err, "Lỗi tạo log lô hàng")
			// }

			variantDb, err := q.GetVariantById(ctx, value.Variant)
			if err != nil {
				return common.ErrDB(err)
			}

			if value.GetValue() > variantDb.RealInventory {
				return common.ErrInternalWithMsg(errors.New("invalid inventory"), "Tồn kho không đủ")
			}

			_, err = q.UpdateVariant(ctx, UpdateVariantParams{
				ID: value.Variant,
				RealInventory: sql.NullInt32{
					Int32: variantDb.RealInventory - value.GetValue(),
					Valid: true,
				},
			})
			if err != nil {
				return common.ErrDBWithMsg(err, "lỗi update tồn kho sản phẩm")
			}

			_, err = q.CreateOrderItem(ctx, CreateOrderItemParams{
				Order:      order.ID,
				Variant:    value.GetVariant(),
				Value:      value.GetValue(),
				TotalPrice: float64(value.GetTotalPrice()),
				// Consignment: sql.NullInt32{
				// 	Int32: consignment.ID,
				// 	Valid: true,
				// },
				// ConsignmentLog: sql.NullInt32{
				// 	Int32: consignmentLog.ID,
				// 	Valid: true,
				// },
			})
			if err != nil {
				return common.ErrDBWithMsg(err, "Lỗi tạo sản phẩm đơn hàng")
			}
		}

		for _, item := range req.GetServiceItems() {
			_, err = q.CreateOrderServiceItem(ctx, CreateOrderServiceItemParams{
				Order: order.ID,
				Service: sql.NullInt32{
					Int32: item.GetService(),
					Valid: true,
				},
				UnitPrice:  float64(item.GetUnitPrice()),
				TotalPrice: float64(item.GetTotalPrice()),
				Discount:   float64(item.GetDiscount()),
			})
			if err != nil {
				return common.ErrDBWithMsg(err, "Lỗi tạo dịch vụ đơn hàng") // status.Errorf(codes.Internal, "failed to record order service: %e", err)
			}
		}
		result.Id = order.ID

		return err
	})

	return result, err
}
