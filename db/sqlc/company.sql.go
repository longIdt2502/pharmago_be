// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: company.sql

package db

import (
	"context"
	"database/sql"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO companies (
    name, code, tax_code, phone, description, address, owner
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, name, code, tax_code, phone, description, created_at, owner, address
`

type CreateCompanyParams struct {
	Name        string         `json:"name"`
	Code        string         `json:"code"`
	TaxCode     sql.NullString `json:"tax_code"`
	Phone       sql.NullString `json:"phone"`
	Description sql.NullString `json:"description"`
	Address     int64          `json:"address"`
	Owner       int64          `json:"owner"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, createCompany,
		arg.Name,
		arg.Code,
		arg.TaxCode,
		arg.Phone,
		arg.Description,
		arg.Address,
		arg.Owner,
	)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.TaxCode,
		&i.Phone,
		&i.Description,
		&i.CreatedAt,
		&i.Owner,
		&i.Address,
	)
	return i, err
}
