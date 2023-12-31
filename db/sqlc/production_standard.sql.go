// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: production_standard.sql

package db

import (
	"context"
)

const createProductionStandard = `-- name: CreateProductionStandard :one
INSERT INTO production_standard (
    code, name
) VALUES (
    $1, $2
) RETURNING id, code, name
`

type CreateProductionStandardParams struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (q *Queries) CreateProductionStandard(ctx context.Context, arg CreateProductionStandardParams) (ProductionStandard, error) {
	row := q.db.QueryRowContext(ctx, createProductionStandard, arg.Code, arg.Name)
	var i ProductionStandard
	err := row.Scan(&i.ID, &i.Code, &i.Name)
	return i, err
}
