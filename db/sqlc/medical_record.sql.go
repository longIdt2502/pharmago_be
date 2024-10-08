// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: medical_record.sql

package db

import (
	"context"
	"database/sql"
)

const createMedicalRecord = `-- name: CreateMedicalRecord :one
INSERT INTO medical_records (
    code, customer, weight, long, symptom, diagnostic, result, doctor, re_examination, note, user_created 
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, code, customer, weight, long, symptom, diagnostic, result, doctor, re_examination, note, created_at, updated_at, user_created, user_updated
`

type CreateMedicalRecordParams struct {
	Code          string          `json:"code"`
	Customer      int32           `json:"customer"`
	Weight        sql.NullFloat64 `json:"weight"`
	Long          sql.NullFloat64 `json:"long"`
	Symptom       string          `json:"symptom"`
	Diagnostic    string          `json:"diagnostic"`
	Result        string          `json:"result"`
	Doctor        sql.NullInt32   `json:"doctor"`
	ReExamination int32           `json:"re_examination"`
	Note          sql.NullString  `json:"note"`
	UserCreated   sql.NullInt32   `json:"user_created"`
}

func (q *Queries) CreateMedicalRecord(ctx context.Context, arg CreateMedicalRecordParams) (MedicalRecord, error) {
	row := q.db.QueryRowContext(ctx, createMedicalRecord,
		arg.Code,
		arg.Customer,
		arg.Weight,
		arg.Long,
		arg.Symptom,
		arg.Diagnostic,
		arg.Result,
		arg.Doctor,
		arg.ReExamination,
		arg.Note,
		arg.UserCreated,
	)
	var i MedicalRecord
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Customer,
		&i.Weight,
		&i.Long,
		&i.Symptom,
		&i.Diagnostic,
		&i.Result,
		&i.Doctor,
		&i.ReExamination,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserCreated,
		&i.UserUpdated,
	)
	return i, err
}

const detailMedicalRecord = `-- name: DetailMedicalRecord :one
SELECT id, code, customer, weight, long, symptom, diagnostic, result, doctor, re_examination, note, created_at, updated_at, user_created, user_updated FROM medical_records
WHERE id = $1
`

func (q *Queries) DetailMedicalRecord(ctx context.Context, id int32) (MedicalRecord, error) {
	row := q.db.QueryRowContext(ctx, detailMedicalRecord, id)
	var i MedicalRecord
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Customer,
		&i.Weight,
		&i.Long,
		&i.Symptom,
		&i.Diagnostic,
		&i.Result,
		&i.Doctor,
		&i.ReExamination,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserCreated,
		&i.UserUpdated,
	)
	return i, err
}

const listMedicalRecord = `-- name: ListMedicalRecord :many
SELECT id, code, customer, weight, long, symptom, diagnostic, result, doctor, re_examination, note, created_at, updated_at, user_created, user_updated FROM medical_records
WHERE customer = $1::int
AND (
    code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type ListMedicalRecordParams struct {
	Customer int32          `json:"customer"`
	Search   sql.NullString `json:"search"`
	Page     sql.NullInt32  `json:"page"`
	Limit    sql.NullInt32  `json:"limit"`
}

func (q *Queries) ListMedicalRecord(ctx context.Context, arg ListMedicalRecordParams) ([]MedicalRecord, error) {
	rows, err := q.db.QueryContext(ctx, listMedicalRecord,
		arg.Customer,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedicalRecord{}
	for rows.Next() {
		var i MedicalRecord
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Customer,
			&i.Weight,
			&i.Long,
			&i.Symptom,
			&i.Diagnostic,
			&i.Result,
			&i.Doctor,
			&i.ReExamination,
			&i.Note,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserCreated,
			&i.UserUpdated,
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
