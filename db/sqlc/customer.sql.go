// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: customer.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const countCustomer = `-- name: CountCustomer :one
SELECT COUNT(*) FROM customers
WHERE company = $1::int
`

func (q *Queries) CountCustomer(ctx context.Context, company int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, countCustomer, company)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
    full_name, code, company, address, email, phone, gender, license, issued_by, birthday, user_updated, user_created, "group", 
    title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number,
    bank_name, bank_branch
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
) RETURNING id, full_name, code, company, address, email, phone, license, birthday, user_created, user_updated, updated_at, created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, gender
`

type CreateCustomerParams struct {
	FullName       string         `json:"full_name"`
	Code           string         `json:"code"`
	Company        int32          `json:"company"`
	Address        sql.NullInt32  `json:"address"`
	Email          sql.NullString `json:"email"`
	Phone          sql.NullString `json:"phone"`
	Gender         NullGender     `json:"gender"`
	License        sql.NullString `json:"license"`
	IssuedBy       sql.NullString `json:"issued_by"`
	Birthday       sql.NullTime   `json:"birthday"`
	UserUpdated    sql.NullInt32  `json:"user_updated"`
	UserCreated    int32          `json:"user_created"`
	Group          sql.NullInt32  `json:"group"`
	Title          sql.NullString `json:"title"`
	LicenseDate    sql.NullTime   `json:"license_date"`
	ContactName    sql.NullString `json:"contact_name"`
	ContactTitle   sql.NullString `json:"contact_title"`
	ContactPhone   sql.NullString `json:"contact_phone"`
	ContactEmail   sql.NullString `json:"contact_email"`
	ContactAddress sql.NullInt32  `json:"contact_address"`
	AccountNumber  sql.NullString `json:"account_number"`
	BankName       sql.NullString `json:"bank_name"`
	BankBranch     sql.NullString `json:"bank_branch"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.FullName,
		arg.Code,
		arg.Company,
		arg.Address,
		arg.Email,
		arg.Phone,
		arg.Gender,
		arg.License,
		arg.IssuedBy,
		arg.Birthday,
		arg.UserUpdated,
		arg.UserCreated,
		arg.Group,
		arg.Title,
		arg.LicenseDate,
		arg.ContactName,
		arg.ContactTitle,
		arg.ContactPhone,
		arg.ContactEmail,
		arg.ContactAddress,
		arg.AccountNumber,
		arg.BankName,
		arg.BankBranch,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Code,
		&i.Company,
		&i.Address,
		&i.Email,
		&i.Phone,
		&i.License,
		&i.Birthday,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
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
	)
	return i, err
}

const createCustomerGroup = `-- name: CreateCustomerGroup :one
INSERT INTO customer_group (
    code, name, company, note, user_created
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, code, name, company, note, user_created, user_updated, updated_at, created_at
`

type CreateCustomerGroupParams struct {
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Company     int32          `json:"company"`
	Note        sql.NullString `json:"note"`
	UserCreated int32          `json:"user_created"`
}

func (q *Queries) CreateCustomerGroup(ctx context.Context, arg CreateCustomerGroupParams) (CustomerGroup, error) {
	row := q.db.QueryRowContext(ctx, createCustomerGroup,
		arg.Code,
		arg.Name,
		arg.Company,
		arg.Note,
		arg.UserCreated,
	)
	var i CustomerGroup
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.Note,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const createMedicalRecordLink = `-- name: CreateMedicalRecordLink :one
INSERT INTO medical_record_link (
    uuid, "type", title, url, customer, appointment_schedule, medical_bill, user_created, size
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, uuid, type, title, url, customer, appointment_schedule, user_created, created_at, medical_bill, size
`

type CreateMedicalRecordLinkParams struct {
	Uuid                uuid.UUID             `json:"uuid"`
	Type                MedicalRecordLinkType `json:"type"`
	Title               sql.NullString        `json:"title"`
	Url                 string                `json:"url"`
	Customer            sql.NullInt32         `json:"customer"`
	AppointmentSchedule uuid.NullUUID         `json:"appointment_schedule"`
	MedicalBill         uuid.NullUUID         `json:"medical_bill"`
	UserCreated         sql.NullInt32         `json:"user_created"`
	Size                sql.NullInt32         `json:"size"`
}

func (q *Queries) CreateMedicalRecordLink(ctx context.Context, arg CreateMedicalRecordLinkParams) (MedicalRecordLink, error) {
	row := q.db.QueryRowContext(ctx, createMedicalRecordLink,
		arg.Uuid,
		arg.Type,
		arg.Title,
		arg.Url,
		arg.Customer,
		arg.AppointmentSchedule,
		arg.MedicalBill,
		arg.UserCreated,
		arg.Size,
	)
	var i MedicalRecordLink
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Type,
		&i.Title,
		&i.Url,
		&i.Customer,
		&i.AppointmentSchedule,
		&i.UserCreated,
		&i.CreatedAt,
		&i.MedicalBill,
		&i.Size,
	)
	return i, err
}

const deleteCustomerGroup = `-- name: DeleteCustomerGroup :one
DELETE FROM customer_group
WHERE id = $1 RETURNING id, code, name, company, note, user_created, user_updated, updated_at, created_at
`

func (q *Queries) DeleteCustomerGroup(ctx context.Context, id int32) (CustomerGroup, error) {
	row := q.db.QueryRowContext(ctx, deleteCustomerGroup, id)
	var i CustomerGroup
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.Note,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMedicalRecordLink = `-- name: DeleteMedicalRecordLink :one
DELETE FROM medical_record_link 
WHERE uuid = $1 RETURNING id, uuid, type, title, url, customer, appointment_schedule, user_created, created_at, medical_bill, size
`

func (q *Queries) DeleteMedicalRecordLink(ctx context.Context, argUuid uuid.UUID) (MedicalRecordLink, error) {
	row := q.db.QueryRowContext(ctx, deleteMedicalRecordLink, argUuid)
	var i MedicalRecordLink
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Type,
		&i.Title,
		&i.Url,
		&i.Customer,
		&i.AppointmentSchedule,
		&i.UserCreated,
		&i.CreatedAt,
		&i.MedicalBill,
		&i.Size,
	)
	return i, err
}

const detailCustomer = `-- name: DetailCustomer :one
SELECT id, full_name, code, company, address, email, phone, license, birthday, user_created, user_updated, updated_at, created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, gender FROM customers
WHERE id = $1
LIMIT 1
`

func (q *Queries) DetailCustomer(ctx context.Context, id int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, detailCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Code,
		&i.Company,
		&i.Address,
		&i.Email,
		&i.Phone,
		&i.License,
		&i.Birthday,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
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
	)
	return i, err
}

const detailCustomerGroup = `-- name: DetailCustomerGroup :one
SELECT cg.id, code, name, company, note, user_created, user_updated, updated_at, cg.created_at, ac.id, ac.username, ac.hashed_password, ac.full_name, ac.email, ac.type, ac.is_verify, ac.password_changed_at, ac.created_at, ac.role, ac.gender, ac.licence, ac.dob, ac.address, au.id, au.username, au.hashed_password, au.full_name, au.email, au.type, au.is_verify, au.password_changed_at, au.created_at, au.role, au.gender, au.licence, au.dob, au.address FROM customer_group cg
LEFT JOIN accounts ac ON ac.id = cg.user_created
LEFT JOIN accounts au ON au.id = cg.user_updated
WHERE cg.id = $1
`

type DetailCustomerGroupRow struct {
	ID                  int32          `json:"id"`
	Code                string         `json:"code"`
	Name                string         `json:"name"`
	Company             int32          `json:"company"`
	Note                sql.NullString `json:"note"`
	UserCreated         int32          `json:"user_created"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	CreatedAt           time.Time      `json:"created_at"`
	ID_2                sql.NullInt32  `json:"id_2"`
	Username            sql.NullString `json:"username"`
	HashedPassword      sql.NullString `json:"hashed_password"`
	FullName            sql.NullString `json:"full_name"`
	Email               sql.NullString `json:"email"`
	Type                sql.NullInt32  `json:"type"`
	IsVerify            sql.NullBool   `json:"is_verify"`
	PasswordChangedAt   sql.NullTime   `json:"password_changed_at"`
	CreatedAt_2         sql.NullTime   `json:"created_at_2"`
	Role                sql.NullInt32  `json:"role"`
	Gender              NullGender     `json:"gender"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address             sql.NullInt32  `json:"address"`
	ID_3                sql.NullInt32  `json:"id_3"`
	Username_2          sql.NullString `json:"username_2"`
	HashedPassword_2    sql.NullString `json:"hashed_password_2"`
	FullName_2          sql.NullString `json:"full_name_2"`
	Email_2             sql.NullString `json:"email_2"`
	Type_2              sql.NullInt32  `json:"type_2"`
	IsVerify_2          sql.NullBool   `json:"is_verify_2"`
	PasswordChangedAt_2 sql.NullTime   `json:"password_changed_at_2"`
	CreatedAt_3         sql.NullTime   `json:"created_at_3"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_2           sql.NullInt32  `json:"address_2"`
}

func (q *Queries) DetailCustomerGroup(ctx context.Context, id int32) (DetailCustomerGroupRow, error) {
	row := q.db.QueryRowContext(ctx, detailCustomerGroup, id)
	var i DetailCustomerGroupRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.Note,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
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
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
		&i.ID_3,
		&i.Username_2,
		&i.HashedPassword_2,
		&i.FullName_2,
		&i.Email_2,
		&i.Type_2,
		&i.IsVerify_2,
		&i.PasswordChangedAt_2,
		&i.CreatedAt_3,
		&i.Role_2,
		&i.Gender_2,
		&i.Licence_2,
		&i.Dob_2,
		&i.Address_2,
	)
	return i, err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, full_name, code, company, address, email, phone, license, birthday, user_created, user_updated, updated_at, created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, gender FROM customers
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, id int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Code,
		&i.Company,
		&i.Address,
		&i.Email,
		&i.Phone,
		&i.License,
		&i.Birthday,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
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
	)
	return i, err
}

const listCustomer = `-- name: ListCustomer :many
WITH revenue AS (
    SELECT customer,
    COALESCE(SUM(total_price), 0)::float AS total_revenue,
    COALESCE(COUNT(id), 0)::int AS total_orders
    FROM orders
    GROUP BY customer
)
SELECT id, full_name, code, company, address, email, phone, license, birthday, user_created, user_updated, updated_at, created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, gender, customer, total_revenue, total_orders, r.total_revenue, r.total_orders FROM customers c
LEFT JOIN revenue r ON c.id = r.customer 
WHERE c.company = $1::int
AND (
    c.full_name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    c.code ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    c.phone ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -c.id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type ListCustomerParams struct {
	Company int32          `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type ListCustomerRow struct {
	ID             int32           `json:"id"`
	FullName       string          `json:"full_name"`
	Code           string          `json:"code"`
	Company        int32           `json:"company"`
	Address        sql.NullInt32   `json:"address"`
	Email          sql.NullString  `json:"email"`
	Phone          sql.NullString  `json:"phone"`
	License        sql.NullString  `json:"license"`
	Birthday       sql.NullTime    `json:"birthday"`
	UserCreated    int32           `json:"user_created"`
	UserUpdated    sql.NullInt32   `json:"user_updated"`
	UpdatedAt      sql.NullTime    `json:"updated_at"`
	CreatedAt      time.Time       `json:"created_at"`
	Group          sql.NullInt32   `json:"group"`
	Title          sql.NullString  `json:"title"`
	LicenseDate    sql.NullTime    `json:"license_date"`
	ContactName    sql.NullString  `json:"contact_name"`
	ContactTitle   sql.NullString  `json:"contact_title"`
	ContactPhone   sql.NullString  `json:"contact_phone"`
	ContactEmail   sql.NullString  `json:"contact_email"`
	ContactAddress sql.NullInt32   `json:"contact_address"`
	AccountNumber  sql.NullString  `json:"account_number"`
	BankName       sql.NullString  `json:"bank_name"`
	BankBranch     sql.NullString  `json:"bank_branch"`
	IssuedBy       sql.NullString  `json:"issued_by"`
	Gender         NullGender      `json:"gender"`
	Customer       sql.NullInt32   `json:"customer"`
	TotalRevenue   sql.NullFloat64 `json:"total_revenue"`
	TotalOrders    sql.NullInt32   `json:"total_orders"`
	TotalRevenue_2 sql.NullFloat64 `json:"total_revenue_2"`
	TotalOrders_2  sql.NullInt32   `json:"total_orders_2"`
}

func (q *Queries) ListCustomer(ctx context.Context, arg ListCustomerParams) ([]ListCustomerRow, error) {
	rows, err := q.db.QueryContext(ctx, listCustomer,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCustomerRow{}
	for rows.Next() {
		var i ListCustomerRow
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Code,
			&i.Company,
			&i.Address,
			&i.Email,
			&i.Phone,
			&i.License,
			&i.Birthday,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
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
			&i.Customer,
			&i.TotalRevenue,
			&i.TotalOrders,
			&i.TotalRevenue_2,
			&i.TotalOrders_2,
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

const listCustomerGroup = `-- name: ListCustomerGroup :many
SELECT cg.id, code, name, company, note, user_created, user_updated, updated_at, cg.created_at, ac.id, ac.username, ac.hashed_password, ac.full_name, ac.email, ac.type, ac.is_verify, ac.password_changed_at, ac.created_at, ac.role, ac.gender, ac.licence, ac.dob, ac.address, au.id, au.username, au.hashed_password, au.full_name, au.email, au.type, au.is_verify, au.password_changed_at, au.created_at, au.role, au.gender, au.licence, au.dob, au.address FROM customer_group cg
LEFT JOIN accounts ac ON ac.id = cg.user_created
LEFT JOIN accounts au ON au.id = cg.user_updated
WHERE cg.company = $1::int
AND (
    cg.name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    cg.code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -cg.id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type ListCustomerGroupParams struct {
	Company int32          `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type ListCustomerGroupRow struct {
	ID                  int32          `json:"id"`
	Code                string         `json:"code"`
	Name                string         `json:"name"`
	Company             int32          `json:"company"`
	Note                sql.NullString `json:"note"`
	UserCreated         int32          `json:"user_created"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	CreatedAt           time.Time      `json:"created_at"`
	ID_2                sql.NullInt32  `json:"id_2"`
	Username            sql.NullString `json:"username"`
	HashedPassword      sql.NullString `json:"hashed_password"`
	FullName            sql.NullString `json:"full_name"`
	Email               sql.NullString `json:"email"`
	Type                sql.NullInt32  `json:"type"`
	IsVerify            sql.NullBool   `json:"is_verify"`
	PasswordChangedAt   sql.NullTime   `json:"password_changed_at"`
	CreatedAt_2         sql.NullTime   `json:"created_at_2"`
	Role                sql.NullInt32  `json:"role"`
	Gender              NullGender     `json:"gender"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address             sql.NullInt32  `json:"address"`
	ID_3                sql.NullInt32  `json:"id_3"`
	Username_2          sql.NullString `json:"username_2"`
	HashedPassword_2    sql.NullString `json:"hashed_password_2"`
	FullName_2          sql.NullString `json:"full_name_2"`
	Email_2             sql.NullString `json:"email_2"`
	Type_2              sql.NullInt32  `json:"type_2"`
	IsVerify_2          sql.NullBool   `json:"is_verify_2"`
	PasswordChangedAt_2 sql.NullTime   `json:"password_changed_at_2"`
	CreatedAt_3         sql.NullTime   `json:"created_at_3"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_2           sql.NullInt32  `json:"address_2"`
}

func (q *Queries) ListCustomerGroup(ctx context.Context, arg ListCustomerGroupParams) ([]ListCustomerGroupRow, error) {
	rows, err := q.db.QueryContext(ctx, listCustomerGroup,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCustomerGroupRow{}
	for rows.Next() {
		var i ListCustomerGroupRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.Company,
			&i.Note,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
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
			&i.Gender,
			&i.Licence,
			&i.Dob,
			&i.Address,
			&i.ID_3,
			&i.Username_2,
			&i.HashedPassword_2,
			&i.FullName_2,
			&i.Email_2,
			&i.Type_2,
			&i.IsVerify_2,
			&i.PasswordChangedAt_2,
			&i.CreatedAt_3,
			&i.Role_2,
			&i.Gender_2,
			&i.Licence_2,
			&i.Dob_2,
			&i.Address_2,
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

const listMedicalRecordLink = `-- name: ListMedicalRecordLink :many
SELECT id, uuid, type, title, url, customer, appointment_schedule, user_created, created_at, medical_bill, size FROM medical_record_link
WHERE ($1::int IS NULL OR $1::int = customer)
AND ($2::medical_record_link_type IS NULL OR $2::medical_record_link_type = "type")
AND ($3::uuid IS NULL OR $3::uuid = appointment_schedule)
AND ($4::uuid IS NULL OR $4::uuid = medical_bill)
`

type ListMedicalRecordLinkParams struct {
	Customer    sql.NullInt32             `json:"customer"`
	TypeMrl     NullMedicalRecordLinkType `json:"type_mrl"`
	Schedule    uuid.NullUUID             `json:"schedule"`
	MedicalBill uuid.NullUUID             `json:"medical_bill"`
}

func (q *Queries) ListMedicalRecordLink(ctx context.Context, arg ListMedicalRecordLinkParams) ([]MedicalRecordLink, error) {
	rows, err := q.db.QueryContext(ctx, listMedicalRecordLink,
		arg.Customer,
		arg.TypeMrl,
		arg.Schedule,
		arg.MedicalBill,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedicalRecordLink{}
	for rows.Next() {
		var i MedicalRecordLink
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Type,
			&i.Title,
			&i.Url,
			&i.Customer,
			&i.AppointmentSchedule,
			&i.UserCreated,
			&i.CreatedAt,
			&i.MedicalBill,
			&i.Size,
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

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE customers
SET
    full_name = COALESCE($1::varchar, full_name),
    code = COALESCE($2::varchar, code),
    email = COALESCE($3::varchar, email),
    phone = COALESCE($4::varchar, phone),
    license = COALESCE($5::varchar, license),
    birthday = COALESCE($6::timestamp, birthday),
    user_updated = COALESCE($7::int, user_updated),
    "group" = COALESCE($8::int, "group"),
    title = COALESCE($9::varchar, title),
    gender = COALESCE($10::gender, gender),
    license_date = COALESCE($11::timestamp, license_date),
    contact_name = COALESCE($12::varchar, contact_name),
    contact_title = COALESCE($13::varchar, contact_title),
    contact_phone = COALESCE($14::varchar, contact_phone),
    contact_email = COALESCE($15::varchar, contact_email),
    contact_address = COALESCE($16::int, contact_address),
    account_number = COALESCE($17::varchar, account_number),
    bank_name = COALESCE($18::varchar, bank_name),
    bank_branch = COALESCE($19::varchar, bank_branch)
WHERE id = $20
RETURNING id, full_name, code, company, address, email, phone, license, birthday, user_created, user_updated, updated_at, created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, gender
`

type UpdateCustomerParams struct {
	FullName       sql.NullString `json:"full_name"`
	Code           sql.NullString `json:"code"`
	Email          sql.NullString `json:"email"`
	Phone          sql.NullString `json:"phone"`
	License        sql.NullString `json:"license"`
	Birthday       sql.NullTime   `json:"birthday"`
	UserUpdated    sql.NullInt32  `json:"user_updated"`
	Group          sql.NullInt32  `json:"group"`
	Title          sql.NullString `json:"title"`
	Gender         NullGender     `json:"gender"`
	LicenseDate    sql.NullTime   `json:"license_date"`
	ContactName    sql.NullString `json:"contact_name"`
	ContactTitle   sql.NullString `json:"contact_title"`
	ContactPhone   sql.NullString `json:"contact_phone"`
	ContactEmail   sql.NullString `json:"contact_email"`
	ContactAddress sql.NullInt32  `json:"contact_address"`
	AccountNumber  sql.NullString `json:"account_number"`
	BankName       sql.NullString `json:"bank_name"`
	BankBranch     sql.NullString `json:"bank_branch"`
	ID             int32          `json:"id"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.FullName,
		arg.Code,
		arg.Email,
		arg.Phone,
		arg.License,
		arg.Birthday,
		arg.UserUpdated,
		arg.Group,
		arg.Title,
		arg.Gender,
		arg.LicenseDate,
		arg.ContactName,
		arg.ContactTitle,
		arg.ContactPhone,
		arg.ContactEmail,
		arg.ContactAddress,
		arg.AccountNumber,
		arg.BankName,
		arg.BankBranch,
		arg.ID,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Code,
		&i.Company,
		&i.Address,
		&i.Email,
		&i.Phone,
		&i.License,
		&i.Birthday,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
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
	)
	return i, err
}

const updateCustomerGroup = `-- name: UpdateCustomerGroup :one
UPDATE customer_group
SET
    name = COALESCE($1, name),
    code = COALESCE($2, code),
    note = COALESCE($3, note),
    user_updated = $4,
    updated_at = now()
WHERE id = $5
RETURNING id, code, name, company, note, user_created, user_updated, updated_at, created_at
`

type UpdateCustomerGroupParams struct {
	Name        sql.NullString `json:"name"`
	Code        sql.NullString `json:"code"`
	Note        sql.NullString `json:"note"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateCustomerGroup(ctx context.Context, arg UpdateCustomerGroupParams) (CustomerGroup, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerGroup,
		arg.Name,
		arg.Code,
		arg.Note,
		arg.UserUpdated,
		arg.ID,
	)
	var i CustomerGroup
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.Company,
		&i.Note,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
