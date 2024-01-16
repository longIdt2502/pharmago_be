// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: production_standard.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createProductionStandard = `-- name: CreateProductionStandard :one
INSERT INTO production_standard (
    code, name, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, code, name, company, user_created, user_updated, created_at, updated_at, description
`

type CreateProductionStandardParams struct {
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Company     sql.NullInt32 `json:"company"`
	UserCreated sql.NullInt32 `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
}

func (q *Queries) CreateProductionStandard(ctx context.Context, arg CreateProductionStandardParams) (ProductionStandard, error) {
	row := q.db.QueryRowContext(ctx, createProductionStandard,
		arg.Code,
		arg.Name,
		arg.Company,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i ProductionStandard
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const deleteProductionStandard = `-- name: DeleteProductionStandard :one
DELETE FROM production_standard
WHERE id = $1 RETURNING id, code, name, company, user_created, user_updated, created_at, updated_at, description
`

func (q *Queries) DeleteProductionStandard(ctx context.Context, id int32) (ProductionStandard, error) {
	row := q.db.QueryRowContext(ctx, deleteProductionStandard, id)
	var i ProductionStandard
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const detailProductionStandard = `-- name: DetailProductionStandard :one
SELECT ps.id, ps.code, ps.name, ps.company, ps.user_created, ps.user_updated, ps.created_at, ps.updated_at, ps.description, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, au.full_name AS user_updated_name FROM production_standard ps
LEFT JOIN accounts a ON a.id = ps.user_created
LEFT JOIN accounts au ON au.id = ps.user_updated
WHERE ps.id = $1
`

type DetailProductionStandardRow struct {
	ID                int32          `json:"id"`
	Code              string         `json:"code"`
	Name              string         `json:"name"`
	Company           sql.NullInt32  `json:"company"`
	UserCreated       sql.NullInt32  `json:"user_created"`
	UserUpdated       sql.NullInt32  `json:"user_updated"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	Description       sql.NullString `json:"description"`
	ID_2              sql.NullInt32  `json:"id_2"`
	Username          sql.NullString `json:"username"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	Type              sql.NullInt32  `json:"type"`
	IsVerify          sql.NullBool   `json:"is_verify"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	CreatedAt_2       sql.NullTime   `json:"created_at_2"`
	UserUpdatedName   sql.NullString `json:"user_updated_name"`
}

func (q *Queries) DetailProductionStandard(ctx context.Context, id int32) (DetailProductionStandardRow, error) {
	row := q.db.QueryRowContext(ctx, detailProductionStandard, id)
	var i DetailProductionStandardRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.ID_2,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt_2,
		&i.UserUpdatedName,
	)
	return i, err
}

const listProductionStandard = `-- name: ListProductionStandard :many
WITH production_standard_quantity AS (
    SELECT ps.id AS production_standard_id,
           COALESCE(COUNT(p.id), 0)::int AS total_quantity
    FROM production_standard ps
    LEFT JOIN products p ON p.tieu_chuan_sx = ps.code
    GROUP BY ps.id
)
SELECT ps.id, ps.code, ps.name, ps.company, ps.user_created, ps.user_updated, ps.created_at, ps.updated_at, ps.description, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, psq.total_quantity AS quantity FROM production_standard_quantity psq
JOIN production_standard ps ON psq.production_standard_id = ps.id
LEFT JOIN accounts a ON a.id = ps.user_created
WHERE (
    ps.company IS NULL OR
    ps.company = $1::int
)
AND (
    ps.name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    ps.code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -ps.id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type ListProductionStandardParams struct {
	Company int32          `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type ListProductionStandardRow struct {
	ID                int32          `json:"id"`
	Code              string         `json:"code"`
	Name              string         `json:"name"`
	Company           sql.NullInt32  `json:"company"`
	UserCreated       sql.NullInt32  `json:"user_created"`
	UserUpdated       sql.NullInt32  `json:"user_updated"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	Description       sql.NullString `json:"description"`
	ID_2              sql.NullInt32  `json:"id_2"`
	Username          sql.NullString `json:"username"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	Type              sql.NullInt32  `json:"type"`
	IsVerify          sql.NullBool   `json:"is_verify"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	CreatedAt_2       sql.NullTime   `json:"created_at_2"`
	Quantity          int32          `json:"quantity"`
}

func (q *Queries) ListProductionStandard(ctx context.Context, arg ListProductionStandardParams) ([]ListProductionStandardRow, error) {
	rows, err := q.db.QueryContext(ctx, listProductionStandard,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductionStandardRow{}
	for rows.Next() {
		var i ListProductionStandardRow
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
			&i.ID_2,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.Type,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt_2,
			&i.Quantity,
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

const updateProductionStandard = `-- name: UpdateProductionStandard :one
UPDATE production_standard
SET
    name = $1,
    code = COALESCE($2, code),
    description = COALESCE($3, description),
    user_updated = $4,
    updated_at = now()
WHERE id = $5
RETURNING id, code, name, company, user_created, user_updated, created_at, updated_at, description
`

type UpdateProductionStandardParams struct {
	Name        string         `json:"name"`
	Code        sql.NullString `json:"code"`
	Description sql.NullString `json:"description"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateProductionStandard(ctx context.Context, arg UpdateProductionStandardParams) (ProductionStandard, error) {
	row := q.db.QueryRowContext(ctx, updateProductionStandard,
		arg.Name,
		arg.Code,
		arg.Description,
		arg.UserUpdated,
		arg.ID,
	)
	var i ProductionStandard
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}
