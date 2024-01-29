// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: preparation_type.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPreparationType = `-- name: CreatePreparationType :one
INSERT INTO preparation_type (
    code, name, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, code, name, company, user_created, user_updated, created_at, updated_at, description
`

type CreatePreparationTypeParams struct {
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Company     sql.NullInt32 `json:"company"`
	UserCreated sql.NullInt32 `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
}

func (q *Queries) CreatePreparationType(ctx context.Context, arg CreatePreparationTypeParams) (PreparationType, error) {
	row := q.db.QueryRowContext(ctx, createPreparationType,
		arg.Code,
		arg.Name,
		arg.Company,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i PreparationType
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

const deletePreparationType = `-- name: DeletePreparationType :one
DELETE FROM preparation_type
WHERE id = $1 RETURNING id, code, name, company, user_created, user_updated, created_at, updated_at, description
`

func (q *Queries) DeletePreparationType(ctx context.Context, id int32) (PreparationType, error) {
	row := q.db.QueryRowContext(ctx, deletePreparationType, id)
	var i PreparationType
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

const detailPreparationType = `-- name: DetailPreparationType :one
SELECT ps.id, ps.code, ps.name, ps.company, ps.user_created, ps.user_updated, ps.created_at, ps.updated_at, ps.description, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, a.role, au.full_name AS user_updated_name FROM preparation_type ps
LEFT JOIN accounts a ON a.id = ps.user_created
LEFT JOIN accounts au ON au.id = ps.user_updated
WHERE ps.id = $1
`

type DetailPreparationTypeRow struct {
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
	Role              sql.NullInt32  `json:"role"`
	UserUpdatedName   sql.NullString `json:"user_updated_name"`
}

func (q *Queries) DetailPreparationType(ctx context.Context, id int32) (DetailPreparationTypeRow, error) {
	row := q.db.QueryRowContext(ctx, detailPreparationType, id)
	var i DetailPreparationTypeRow
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
		&i.Role,
		&i.UserUpdatedName,
	)
	return i, err
}

const listPreparationType = `-- name: ListPreparationType :many
WITH preparation_type_quantity AS (
    SELECT ps.id AS preparation_type_id,
           COALESCE(COUNT(p.id), 0)::int AS total_quantity
    FROM preparation_type ps
             LEFT JOIN products p ON p.tieu_chuan_sx = ps.code
    GROUP BY ps.id
)
SELECT ps.id, ps.code, ps.name, ps.company, ps.user_created, ps.user_updated, ps.created_at, ps.updated_at, ps.description, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, a.role, psq.total_quantity AS quantity FROM preparation_type_quantity psq
JOIN preparation_type ps ON psq.preparation_type_id = ps.id
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

type ListPreparationTypeParams struct {
	Company int32          `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type ListPreparationTypeRow struct {
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
	Role              sql.NullInt32  `json:"role"`
	Quantity          int32          `json:"quantity"`
}

func (q *Queries) ListPreparationType(ctx context.Context, arg ListPreparationTypeParams) ([]ListPreparationTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, listPreparationType,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPreparationTypeRow{}
	for rows.Next() {
		var i ListPreparationTypeRow
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
			&i.Role,
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

const updatePreparationType = `-- name: UpdatePreparationType :one
UPDATE preparation_type
SET
    name = $1,
    code = COALESCE($2, code),
    description = COALESCE($3, description),
    user_updated = $4,
    updated_at = now()
WHERE id = $5
    RETURNING id, code, name, company, user_created, user_updated, created_at, updated_at, description
`

type UpdatePreparationTypeParams struct {
	Name        string         `json:"name"`
	Code        sql.NullString `json:"code"`
	Description sql.NullString `json:"description"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdatePreparationType(ctx context.Context, arg UpdatePreparationTypeParams) (PreparationType, error) {
	row := q.db.QueryRowContext(ctx, updatePreparationType,
		arg.Name,
		arg.Code,
		arg.Description,
		arg.UserUpdated,
		arg.ID,
	)
	var i PreparationType
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
