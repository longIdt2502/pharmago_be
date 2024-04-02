// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: debt_note.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createDebtNote = `-- name: CreateDebtNote :one
INSERT INTO debt_note (
    code, title, entity, money, paymented, note, type, status, company, user_created, exprise, dabt_note_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id, code, title, entity, money, paymented, note, type, status, company, user_created, exprise, dabt_note_at
`

type CreateDebtNoteParams struct {
	Code        string         `json:"code"`
	Title       sql.NullString `json:"title"`
	Entity      string         `json:"entity"`
	Money       float64        `json:"money"`
	Paymented   float64        `json:"paymented"`
	Note        sql.NullString `json:"note"`
	Type        string         `json:"type"`
	Status      string         `json:"status"`
	Company     int32          `json:"company"`
	UserCreated int32          `json:"user_created"`
	Exprise     time.Time      `json:"exprise"`
	DabtNoteAt  sql.NullTime   `json:"dabt_note_at"`
}

func (q *Queries) CreateDebtNote(ctx context.Context, arg CreateDebtNoteParams) (DebtNote, error) {
	row := q.db.QueryRowContext(ctx, createDebtNote,
		arg.Code,
		arg.Title,
		arg.Entity,
		arg.Money,
		arg.Paymented,
		arg.Note,
		arg.Type,
		arg.Status,
		arg.Company,
		arg.UserCreated,
		arg.Exprise,
		arg.DabtNoteAt,
	)
	var i DebtNote
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Entity,
		&i.Money,
		&i.Paymented,
		&i.Note,
		&i.Type,
		&i.Status,
		&i.Company,
		&i.UserCreated,
		&i.Exprise,
		&i.DabtNoteAt,
	)
	return i, err
}

const createRepayment = `-- name: CreateRepayment :one
INSERT INTO debt_repayment (
    code, money, debt, user_created
) VALUES (
    $1, $2, $3, $4
) RETURNING id, code, money, created_at, debt, user_created
`

type CreateRepaymentParams struct {
	Code        string  `json:"code"`
	Money       float64 `json:"money"`
	Debt        int32   `json:"debt"`
	UserCreated int32   `json:"user_created"`
}

func (q *Queries) CreateRepayment(ctx context.Context, arg CreateRepaymentParams) (DebtRepayment, error) {
	row := q.db.QueryRowContext(ctx, createRepayment,
		arg.Code,
		arg.Money,
		arg.Debt,
		arg.UserCreated,
	)
	var i DebtRepayment
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Money,
		&i.CreatedAt,
		&i.Debt,
		&i.UserCreated,
	)
	return i, err
}

const detailDebtNote = `-- name: DetailDebtNote :one
SELECT id, code, title, entity, money, paymented, note, type, status, company, user_created, exprise, dabt_note_at FROM debt_note
WHERE id = $1
`

func (q *Queries) DetailDebtNote(ctx context.Context, id int32) (DebtNote, error) {
	row := q.db.QueryRowContext(ctx, detailDebtNote, id)
	var i DebtNote
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Entity,
		&i.Money,
		&i.Paymented,
		&i.Note,
		&i.Type,
		&i.Status,
		&i.Company,
		&i.UserCreated,
		&i.Exprise,
		&i.DabtNoteAt,
	)
	return i, err
}

const getListDebtNote = `-- name: GetListDebtNote :many
SELECT id, code, title, entity, money, paymented, note, type, status, company, user_created, exprise, dabt_note_at FROM debt_note dn
WHERE dn.company = $1::int
AND (
    $2::varchar IS NULL OR dn.status = $2::varchar
)
AND (
    dn.code ILIKE '%' || COALESCE($3::varchar, '') || '%' OR
    dn.title ILIKE '%' || COALESCE($3::varchar, '') || '%'
)
ORDER BY
    CASE WHEN $4::varchar = 'exprise' THEN dn.exprise END DESC,
    CASE WHEN $4::varchar = '-exprise' THEN dn.exprise END ASC,
    CASE WHEN $4::varchar = 'dabt_note_at' THEN dn.dabt_note_at END DESC,
    CASE WHEN $4::varchar = '-dabt_note_at' THEN dn.dabt_note_at END ASC,
    CASE WHEN $4::varchar IS NULL THEN dn.id END DESC
LIMIT COALESCE($6::int, 10)
OFFSET (COALESCE($5::int, 1) - 1) * COALESCE($6::int, 10)
`

type GetListDebtNoteParams struct {
	Company sql.NullInt32  `json:"company"`
	Status  sql.NullString `json:"status"`
	Search  sql.NullString `json:"search"`
	OrderBy sql.NullString `json:"order_by"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

// WITH total_repayment AS (
//
//	SELECT debt, COALESCE(SUM(money), 0)::float AS total_money
//	FROM debt_repayment
//	GROUP BY debt
//
// )
// AND  ((
//
//	sqlc.narg('created_start')::timestamp IS NULL AND sqlc.narg('created_end')::timestamp  IS NULL
//
// ) OR (
//
//	(sqlc.narg('created_start')::timestamp IS NULL OR o.created_at >= sqlc.narg('created_start')::timestamp) AND
//	(sqlc.narg('created_end')::timestamp IS NULL OR o.created_at <= sqlc.narg('created_end')::timestamp)
//
// ))
// AND ((
//
//	sqlc.narg('updated_start')::timestamp IS NULL AND sqlc.narg('updated_end')::timestamp  IS NULL
//
// ) OR (
//
//	(o.updated_at >= sqlc.narg('updated_start')::timestamp OR sqlc.narg('updated_start')::timestamp  IS NULL) AND
//	(o.updated_at <= sqlc.narg('updated_end')::timestamp OR sqlc.narg('updated_end')::timestamp  IS NULL)
//
// ))
func (q *Queries) GetListDebtNote(ctx context.Context, arg GetListDebtNoteParams) ([]DebtNote, error) {
	rows, err := q.db.QueryContext(ctx, getListDebtNote,
		arg.Company,
		arg.Status,
		arg.Search,
		arg.OrderBy,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DebtNote{}
	for rows.Next() {
		var i DebtNote
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Title,
			&i.Entity,
			&i.Money,
			&i.Paymented,
			&i.Note,
			&i.Type,
			&i.Status,
			&i.Company,
			&i.UserCreated,
			&i.Exprise,
			&i.DabtNoteAt,
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

const listRepayment = `-- name: ListRepayment :many
SELECT id, code, money, created_at, debt, user_created FROM debt_repayment
WHERE debt = $1
`

func (q *Queries) ListRepayment(ctx context.Context, debt int32) ([]DebtRepayment, error) {
	rows, err := q.db.QueryContext(ctx, listRepayment, debt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DebtRepayment{}
	for rows.Next() {
		var i DebtRepayment
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Money,
			&i.CreatedAt,
			&i.Debt,
			&i.UserCreated,
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