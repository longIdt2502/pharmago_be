package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
)

type CreateServiceTxParams struct {
	*pb.ServiceCreateRequest
	TokenPayload *token.Payload
}

type CreateServiceTxResult struct {
	Id int32
}

func (store *Store) CreateServiceTx(ctx context.Context, req CreateServiceTxParams) (CreateServiceTxResult, error) {
	var result CreateServiceTxResult

	opts := &sql.TxOptions{
		Isolation: 1,
		ReadOnly:  false,
	}

	err := store.execTx(ctx, opts, func(q *Queries) error {
		var err error

		code := fmt.Sprintf("SER-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
		if req.Code != nil {
			code = req.GetCode()
		}
		serviceDb, err := q.CreateService(ctx, CreateServiceParams{
			Code:  code,
			Title: req.GetTitle(),
			Entity: sql.NullString{
				String: req.GetEntity(),
				Valid:  req.Entity != nil,
			},
			Staff: sql.NullInt32{
				Int32: req.GetStaff(),
				Valid: req.Staff != nil,
			},
			Frequency: sql.NullString{
				String: req.GetFrequency(),
				Valid:  req.Frequency != nil,
			},
			Unit:  req.GetUnit(),
			Price: req.GetPrice(),
			Description: sql.NullString{
				String: req.GetDescription(),
				Valid:  req.Description != nil,
			},
			Company:     req.GetCompany(),
			UserCreated: req.TokenPayload.UserID,
			ReminderTime: sql.NullInt32{
				Int32: req.GetReminderTime(),
				Valid: req.ReminderTime != nil,
			},
		})
		if err != nil {
			return err
		}

		for _, item := range req.GetVariants() {
			_, err = q.CreateServiceVariant(ctx, CreateServiceVariantParams{
				Variant: sql.NullInt32{
					Int32: item,
					Valid: true,
				},
				Service: sql.NullInt32{
					Int32: serviceDb.ID,
					Valid: true,
				},
			})
		}

		result.Id = serviceDb.ID

		return err
	})

	return result, err
}
