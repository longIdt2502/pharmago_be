// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: product_master_data.sql

package db

import (
	"context"
	"database/sql"
)

const getListClassify = `-- name: GetListClassify :many
SELECT id, code, name FROM classify
WHERE (
    name ILIKE '%' || COALESCE($1::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE($1::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE($3::int, 10)
OFFSET (COALESCE($2::int, 1) - 1) * COALESCE($3::int, 10)
`

type GetListClassifyParams struct {
	Search sql.NullString `json:"search"`
	Page   sql.NullInt32  `json:"page"`
	Limit  sql.NullInt32  `json:"limit"`
}

func (q *Queries) GetListClassify(ctx context.Context, arg GetListClassifyParams) ([]Classify, error) {
	rows, err := q.db.QueryContext(ctx, getListClassify, arg.Search, arg.Page, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Classify{}
	for rows.Next() {
		var i Classify
		if err := rows.Scan(&i.ID, &i.Code, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListCompanyPharma = `-- name: GetListCompanyPharma :many
SELECT id, name, code, country, address, company_pharma_type, created_at FROM company_pharma
WHERE company_pharma_type = $1::varchar
AND (
    name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -id
    LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type GetListCompanyPharmaParams struct {
	Type   string         `json:"type"`
	Search sql.NullString `json:"search"`
	Page   sql.NullInt32  `json:"page"`
	Limit  sql.NullInt32  `json:"limit"`
}

func (q *Queries) GetListCompanyPharma(ctx context.Context, arg GetListCompanyPharmaParams) ([]CompanyPharma, error) {
	rows, err := q.db.QueryContext(ctx, getListCompanyPharma,
		arg.Type,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CompanyPharma{}
	for rows.Next() {
		var i CompanyPharma
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Country,
			&i.Address,
			&i.CompanyPharmaType,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListPreparation = `-- name: GetListPreparation :many
SELECT id, code, name, company, user_created, user_updated, created_at, updated_at, description FROM preparation_type
WHERE (
    name ILIKE '%' || COALESCE($1::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE($1::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE($3::int, 10)
OFFSET (COALESCE($2::int, 1) - 1) * COALESCE($3::int, 10)
`

type GetListPreparationParams struct {
	Search sql.NullString `json:"search"`
	Page   sql.NullInt32  `json:"page"`
	Limit  sql.NullInt32  `json:"limit"`
}

func (q *Queries) GetListPreparation(ctx context.Context, arg GetListPreparationParams) ([]PreparationType, error) {
	rows, err := q.db.QueryContext(ctx, getListPreparation, arg.Search, arg.Page, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PreparationType{}
	for rows.Next() {
		var i PreparationType
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.Company,
			&i.UserCreated,
			&i.UserUpdated,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListProductionStandard = `-- name: GetListProductionStandard :many
SELECT id, code, name, company, user_created, user_updated, created_at, updated_at, description FROM production_standard
WHERE (
    name ILIKE '%' || COALESCE($1::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE($1::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE($3::int, 10)
OFFSET (COALESCE($2::int, 1) - 1) * COALESCE($3::int, 10)
`

type GetListProductionStandardParams struct {
	Search sql.NullString `json:"search"`
	Page   sql.NullInt32  `json:"page"`
	Limit  sql.NullInt32  `json:"limit"`
}

func (q *Queries) GetListProductionStandard(ctx context.Context, arg GetListProductionStandardParams) ([]ProductionStandard, error) {
	rows, err := q.db.QueryContext(ctx, getListProductionStandard, arg.Search, arg.Page, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductionStandard{}
	for rows.Next() {
		var i ProductionStandard
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.Company,
			&i.UserCreated,
			&i.UserUpdated,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
