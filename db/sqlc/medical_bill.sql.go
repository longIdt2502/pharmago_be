// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: medical_bill.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createMedicalBill = `-- name: CreateMedicalBill :one
INSERT INTO medical_bills (
    uuid, code, customer, company, doctor, symptoms, diagnostic, is_done, meeting_at, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, uuid, code, customer, company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, user_created, user_updated, created_at, updated_at, prescription
`

type CreateMedicalBillParams struct {
	Uuid        uuid.UUID      `json:"uuid"`
	Code        string         `json:"code"`
	Customer    sql.NullInt32  `json:"customer"`
	Company     sql.NullInt32  `json:"company"`
	Doctor      sql.NullInt32  `json:"doctor"`
	Symptoms    sql.NullString `json:"symptoms"`
	Diagnostic  sql.NullString `json:"diagnostic"`
	IsDone      bool           `json:"is_done"`
	MeetingAt   time.Time      `json:"meeting_at"`
	UserCreated int32          `json:"user_created"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
}

func (q *Queries) CreateMedicalBill(ctx context.Context, arg CreateMedicalBillParams) (MedicalBill, error) {
	row := q.db.QueryRowContext(ctx, createMedicalBill,
		arg.Uuid,
		arg.Code,
		arg.Customer,
		arg.Company,
		arg.Doctor,
		arg.Symptoms,
		arg.Diagnostic,
		arg.IsDone,
		arg.MeetingAt,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i MedicalBill
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Code,
		&i.Customer,
		&i.Company,
		&i.Doctor,
		&i.Symptoms,
		&i.Diagnostic,
		&i.QrCodeUrl,
		&i.IsDone,
		&i.MeetingAt,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Prescription,
	)
	return i, err
}

const detailMedicalBill = `-- name: DetailMedicalBill :one
SELECT id, uuid, code, customer, company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, user_created, user_updated, created_at, updated_at, prescription FROM medical_bills
WHERE uuid = $1
`

func (q *Queries) DetailMedicalBill(ctx context.Context, argUuid uuid.UUID) (MedicalBill, error) {
	row := q.db.QueryRowContext(ctx, detailMedicalBill, argUuid)
	var i MedicalBill
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Code,
		&i.Customer,
		&i.Company,
		&i.Doctor,
		&i.Symptoms,
		&i.Diagnostic,
		&i.QrCodeUrl,
		&i.IsDone,
		&i.MeetingAt,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Prescription,
	)
	return i, err
}

const getListMedicalBill = `-- name: GetListMedicalBill :many
SELECT sch.id, uuid, sch.code, customer, sch.company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, sch.user_created, sch.user_updated, sch.created_at, sch.updated_at, prescription, c.id, c.full_name, c.code, c.company, c.address, c.email, phone, license, birthday, c.user_created, c.user_updated, c.updated_at, c.created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, c.gender, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, a.role, a.gender, a.licence, a.dob, a.address, uc.id, uc.username, uc.hashed_password, uc.full_name, uc.email, uc.type, uc.is_verify, uc.password_changed_at, uc.created_at, uc.role, uc.gender, uc.licence, uc.dob, uc.address, uu.id, uu.username, uu.hashed_password, uu.full_name, uu.email, uu.type, uu.is_verify, uu.password_changed_at, uu.created_at, uu.role, uu.gender, uu.licence, uu.dob, uu.address FROM medical_bills sch
LEFT JOIN customers c ON c.id = sch.customer
JOIN accounts a ON a.id = sch.doctor
JOIN accounts uc ON uc.id = sch.user_created
LEFT JOIN accounts uu ON uu.id = sch.user_updated
WHERE sch.company = $1
AND (
    c.full_name ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
AND ($3::int IS NULL OR $3::int = a.id)
AND ($4::uuid IS NULL OR $4::uuid = sch.uuid)
AND  ((
    $5::timestamp IS NULL AND $6::timestamp IS NULL
) OR (
    ($5::timestamp IS NULL OR sch.created_at >= $5::timestamp) AND
    ($6::timestamp IS NULL OR sch.created_at <= $6::timestamp)
))
LIMIT COALESCE($8::int, 10)
OFFSET (COALESCE($7::int, 1) - 1) * COALESCE($8::int, 10)
`

type GetListMedicalBillParams struct {
	Company      sql.NullInt32  `json:"company"`
	Search       sql.NullString `json:"search"`
	Doctor       sql.NullInt32  `json:"doctor"`
	Uuid         uuid.NullUUID  `json:"uuid"`
	CreatedStart sql.NullTime   `json:"created_start"`
	CreatedEnd   sql.NullTime   `json:"created_end"`
	Page         sql.NullInt32  `json:"page"`
	Limit        sql.NullInt32  `json:"limit"`
}

type GetListMedicalBillRow struct {
	ID                  int32          `json:"id"`
	Uuid                uuid.UUID      `json:"uuid"`
	Code                string         `json:"code"`
	Customer            sql.NullInt32  `json:"customer"`
	Company             sql.NullInt32  `json:"company"`
	Doctor              sql.NullInt32  `json:"doctor"`
	Symptoms            sql.NullString `json:"symptoms"`
	Diagnostic          sql.NullString `json:"diagnostic"`
	QrCodeUrl           sql.NullString `json:"qr_code_url"`
	IsDone              bool           `json:"is_done"`
	MeetingAt           time.Time      `json:"meeting_at"`
	UserCreated         int32          `json:"user_created"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	Prescription        uuid.NullUUID  `json:"prescription"`
	ID_2                sql.NullInt32  `json:"id_2"`
	FullName            sql.NullString `json:"full_name"`
	Code_2              sql.NullString `json:"code_2"`
	Company_2           sql.NullInt32  `json:"company_2"`
	Address             sql.NullInt32  `json:"address"`
	Email               sql.NullString `json:"email"`
	Phone               sql.NullString `json:"phone"`
	License             sql.NullString `json:"license"`
	Birthday            sql.NullTime   `json:"birthday"`
	UserCreated_2       sql.NullInt32  `json:"user_created_2"`
	UserUpdated_2       sql.NullInt32  `json:"user_updated_2"`
	UpdatedAt_2         sql.NullTime   `json:"updated_at_2"`
	CreatedAt_2         sql.NullTime   `json:"created_at_2"`
	Group               sql.NullInt32  `json:"group"`
	Title               sql.NullString `json:"title"`
	LicenseDate         sql.NullTime   `json:"license_date"`
	ContactName         sql.NullString `json:"contact_name"`
	ContactTitle        sql.NullString `json:"contact_title"`
	ContactPhone        sql.NullString `json:"contact_phone"`
	ContactEmail        sql.NullString `json:"contact_email"`
	ContactAddress      sql.NullInt32  `json:"contact_address"`
	AccountNumber       sql.NullString `json:"account_number"`
	BankName            sql.NullString `json:"bank_name"`
	BankBranch          sql.NullString `json:"bank_branch"`
	IssuedBy            sql.NullString `json:"issued_by"`
	Gender              NullGender     `json:"gender"`
	ID_3                int32          `json:"id_3"`
	Username            string         `json:"username"`
	HashedPassword      string         `json:"hashed_password"`
	FullName_2          string         `json:"full_name_2"`
	Email_2             string         `json:"email_2"`
	Type                int32          `json:"type"`
	IsVerify            bool           `json:"is_verify"`
	PasswordChangedAt   time.Time      `json:"password_changed_at"`
	CreatedAt_3         time.Time      `json:"created_at_3"`
	Role                sql.NullInt32  `json:"role"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address_2           sql.NullInt32  `json:"address_2"`
	ID_4                int32          `json:"id_4"`
	Username_2          string         `json:"username_2"`
	HashedPassword_2    string         `json:"hashed_password_2"`
	FullName_3          string         `json:"full_name_3"`
	Email_3             string         `json:"email_3"`
	Type_2              int32          `json:"type_2"`
	IsVerify_2          bool           `json:"is_verify_2"`
	PasswordChangedAt_2 time.Time      `json:"password_changed_at_2"`
	CreatedAt_4         time.Time      `json:"created_at_4"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_3            NullGender     `json:"gender_3"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_3           sql.NullInt32  `json:"address_3"`
	ID_5                sql.NullInt32  `json:"id_5"`
	Username_3          sql.NullString `json:"username_3"`
	HashedPassword_3    sql.NullString `json:"hashed_password_3"`
	FullName_4          sql.NullString `json:"full_name_4"`
	Email_4             sql.NullString `json:"email_4"`
	Type_3              sql.NullInt32  `json:"type_3"`
	IsVerify_3          sql.NullBool   `json:"is_verify_3"`
	PasswordChangedAt_3 sql.NullTime   `json:"password_changed_at_3"`
	CreatedAt_5         sql.NullTime   `json:"created_at_5"`
	Role_3              sql.NullInt32  `json:"role_3"`
	Gender_4            NullGender     `json:"gender_4"`
	Licence_3           sql.NullString `json:"licence_3"`
	Dob_3               sql.NullTime   `json:"dob_3"`
	Address_4           sql.NullInt32  `json:"address_4"`
}

func (q *Queries) GetListMedicalBill(ctx context.Context, arg GetListMedicalBillParams) ([]GetListMedicalBillRow, error) {
	rows, err := q.db.QueryContext(ctx, getListMedicalBill,
		arg.Company,
		arg.Search,
		arg.Doctor,
		arg.Uuid,
		arg.CreatedStart,
		arg.CreatedEnd,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListMedicalBillRow{}
	for rows.Next() {
		var i GetListMedicalBillRow
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Code,
			&i.Customer,
			&i.Company,
			&i.Doctor,
			&i.Symptoms,
			&i.Diagnostic,
			&i.QrCodeUrl,
			&i.IsDone,
			&i.MeetingAt,
			&i.UserCreated,
			&i.UserUpdated,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Prescription,
			&i.ID_2,
			&i.FullName,
			&i.Code_2,
			&i.Company_2,
			&i.Address,
			&i.Email,
			&i.Phone,
			&i.License,
			&i.Birthday,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
			&i.Group,
			&i.Title,
			&i.LicenseDate,
			&i.ContactName,
			&i.ContactTitle,
			&i.ContactPhone,
			&i.ContactEmail,
			&i.ContactAddress,
			&i.AccountNumber,
			&i.BankName,
			&i.BankBranch,
			&i.IssuedBy,
			&i.Gender,
			&i.ID_3,
			&i.Username,
			&i.HashedPassword,
			&i.FullName_2,
			&i.Email_2,
			&i.Type,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt_3,
			&i.Role,
			&i.Gender_2,
			&i.Licence,
			&i.Dob,
			&i.Address_2,
			&i.ID_4,
			&i.Username_2,
			&i.HashedPassword_2,
			&i.FullName_3,
			&i.Email_3,
			&i.Type_2,
			&i.IsVerify_2,
			&i.PasswordChangedAt_2,
			&i.CreatedAt_4,
			&i.Role_2,
			&i.Gender_3,
			&i.Licence_2,
			&i.Dob_2,
			&i.Address_3,
			&i.ID_5,
			&i.Username_3,
			&i.HashedPassword_3,
			&i.FullName_4,
			&i.Email_4,
			&i.Type_3,
			&i.IsVerify_3,
			&i.PasswordChangedAt_3,
			&i.CreatedAt_5,
			&i.Role_3,
			&i.Gender_4,
			&i.Licence_3,
			&i.Dob_3,
			&i.Address_4,
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

const updateMedicalBill = `-- name: UpdateMedicalBill :one
UPDATE medical_bills
SET
    diagnostic = COALESCE($1::varchar, diagnostic),
    symptoms = COALESCE($2::varchar, symptoms)
WHERE uuid = $3::uuid
RETURNING id, uuid, code, customer, company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, user_created, user_updated, created_at, updated_at, prescription
`

type UpdateMedicalBillParams struct {
	Diagnostic sql.NullString `json:"diagnostic"`
	Symptoms   sql.NullString `json:"symptoms"`
	Uuid       uuid.UUID      `json:"uuid"`
}

func (q *Queries) UpdateMedicalBill(ctx context.Context, arg UpdateMedicalBillParams) (MedicalBill, error) {
	row := q.db.QueryRowContext(ctx, updateMedicalBill, arg.Diagnostic, arg.Symptoms, arg.Uuid)
	var i MedicalBill
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Code,
		&i.Customer,
		&i.Company,
		&i.Doctor,
		&i.Symptoms,
		&i.Diagnostic,
		&i.QrCodeUrl,
		&i.IsDone,
		&i.MeetingAt,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Prescription,
	)
	return i, err
}
