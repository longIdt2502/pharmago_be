package db

import (
	"context"
	"database/sql"
	"github.com/longIdt2502/pharmago_be/pb"
)

type CreateRoleTxParams struct {
	CreateRoleParams
	Items []*pb.RoleCreateItem
}

type CreateRoleTxResult struct {
	Id int32
}

func (store *Store) CreateRoleTx(ctx context.Context, params CreateRoleTxParams) (CreateRoleTxResult, error) {
	var result CreateRoleTxResult

	opts := &sql.TxOptions{
		Isolation: 1,
		ReadOnly:  false,
	}

	err := store.execTx(ctx, opts, func(q *Queries) error {
		var err error

		role, err := q.CreateRole(ctx, params.CreateRoleParams)
		if err != nil {
			return err
		}

		for _, item := range params.Items {
			_, err = q.CreateRoleItem(ctx, CreateRoleItemParams{
				Roles: role.ID,
				App:   item.GetAppCode(),
				Value: sql.NullBool{
					Bool:  item.GetChecked(),
					Valid: true,
				},
			})
			if err != nil {
				return err
			}
		}

		result.Id = role.ID

		return err
	})

	return result, err
}
