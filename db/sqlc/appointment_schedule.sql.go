// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: appointment_schedule.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO appointment_schedules (
    uuid, code, customer, company, doctor, symptoms, diagnostic, is_done, meeting_at, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, uuid, code, customer, company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, user_created, user_updated, created_at, updated_at
`

type CreateScheduleParams struct {
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

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (AppointmentSchedule, error) {
	row := q.db.QueryRowContext(ctx, createSchedule,
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
	var i AppointmentSchedule
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
	)
	return i, err
}

const createScheduleDrug = `-- name: CreateScheduleDrug :one
INSERT INTO appointment_schedule_drug (
    as_uuid, variant, lieu_dung, quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING id, as_uuid, variant, lieu_dung, quantity
`

type CreateScheduleDrugParams struct {
	AsUuid   uuid.UUID      `json:"as_uuid"`
	Variant  sql.NullInt32  `json:"variant"`
	LieuDung sql.NullString `json:"lieu_dung"`
	Quantity int32          `json:"quantity"`
}

func (q *Queries) CreateScheduleDrug(ctx context.Context, arg CreateScheduleDrugParams) (AppointmentScheduleDrug, error) {
	row := q.db.QueryRowContext(ctx, createScheduleDrug,
		arg.AsUuid,
		arg.Variant,
		arg.LieuDung,
		arg.Quantity,
	)
	var i AppointmentScheduleDrug
	err := row.Scan(
		&i.ID,
		&i.AsUuid,
		&i.Variant,
		&i.LieuDung,
		&i.Quantity,
	)
	return i, err
}

const createScheduleService = `-- name: CreateScheduleService :one
INSERT INTO appointment_schedule_service (
    as_uuid, "service", order_service
) VALUES (
    $1, $2, $3
) RETURNING id, as_uuid, service, order_service
`

type CreateScheduleServiceParams struct {
	AsUuid       uuid.UUID     `json:"as_uuid"`
	Service      sql.NullInt32 `json:"service"`
	OrderService sql.NullInt32 `json:"order_service"`
}

func (q *Queries) CreateScheduleService(ctx context.Context, arg CreateScheduleServiceParams) (AppointmentScheduleService, error) {
	row := q.db.QueryRowContext(ctx, createScheduleService, arg.AsUuid, arg.Service, arg.OrderService)
	var i AppointmentScheduleService
	err := row.Scan(
		&i.ID,
		&i.AsUuid,
		&i.Service,
		&i.OrderService,
	)
	return i, err
}

const createScheduleUrl = `-- name: CreateScheduleUrl :one
INSERT INTO appointment_schedule_url (
    as_uuid, url, name_doc
) VALUES (
    $1, $2, $3
) RETURNING id, as_uuid, url, name_doc
`

type CreateScheduleUrlParams struct {
	AsUuid  uuid.UUID      `json:"as_uuid"`
	Url     sql.NullString `json:"url"`
	NameDoc sql.NullString `json:"name_doc"`
}

func (q *Queries) CreateScheduleUrl(ctx context.Context, arg CreateScheduleUrlParams) (AppointmentScheduleUrl, error) {
	row := q.db.QueryRowContext(ctx, createScheduleUrl, arg.AsUuid, arg.Url, arg.NameDoc)
	var i AppointmentScheduleUrl
	err := row.Scan(
		&i.ID,
		&i.AsUuid,
		&i.Url,
		&i.NameDoc,
	)
	return i, err
}

const detailSchedule = `-- name: DetailSchedule :one
SELECT id, uuid, code, customer, company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, user_created, user_updated, created_at, updated_at FROM appointment_schedules
WHERE uuid = $1
`

func (q *Queries) DetailSchedule(ctx context.Context, argUuid uuid.UUID) (AppointmentSchedule, error) {
	row := q.db.QueryRowContext(ctx, detailSchedule, argUuid)
	var i AppointmentSchedule
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
	)
	return i, err
}

const getListSchedule = `-- name: GetListSchedule :many
SELECT sch.id, uuid, sch.code, customer, sch.company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, sch.user_created, sch.user_updated, sch.created_at, sch.updated_at, c.id, c.full_name, c.code, c.company, c.address, c.email, phone, license, birthday, c.user_created, c.user_updated, c.updated_at, c.created_at, "group", title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, c.gender, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, a.role, a.gender, a.licence, a.dob, a.address, uc.id, uc.username, uc.hashed_password, uc.full_name, uc.email, uc.type, uc.is_verify, uc.password_changed_at, uc.created_at, uc.role, uc.gender, uc.licence, uc.dob, uc.address, uu.id, uu.username, uu.hashed_password, uu.full_name, uu.email, uu.type, uu.is_verify, uu.password_changed_at, uu.created_at, uu.role, uu.gender, uu.licence, uu.dob, uu.address FROM appointment_schedules sch
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

type GetListScheduleParams struct {
	Company      sql.NullInt32  `json:"company"`
	Search       sql.NullString `json:"search"`
	Doctor       sql.NullInt32  `json:"doctor"`
	Uuid         uuid.NullUUID  `json:"uuid"`
	CreatedStart sql.NullTime   `json:"created_start"`
	CreatedEnd   sql.NullTime   `json:"created_end"`
	Page         sql.NullInt32  `json:"page"`
	Limit        sql.NullInt32  `json:"limit"`
}

type GetListScheduleRow struct {
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

func (q *Queries) GetListSchedule(ctx context.Context, arg GetListScheduleParams) ([]GetListScheduleRow, error) {
	rows, err := q.db.QueryContext(ctx, getListSchedule,
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
	items := []GetListScheduleRow{}
	for rows.Next() {
		var i GetListScheduleRow
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

const getListScheduleDrug = `-- name: GetListScheduleDrug :many
SELECT asd.id, as_uuid, asd.variant, lieu_dung, quantity, v.id, name, code, barcode, decision_number, register_number, longevity, vat, product, user_created, user_updated, updated_at, created_at, initial_inventory, real_inventory, vm.id, vm.variant, media, m.id, media_url FROM appointment_schedule_drug asd
JOIN variants v ON v.id = asd.variant
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
WHERE asd.as_uuid = $1
`

type GetListScheduleDrugRow struct {
	ID               int32           `json:"id"`
	AsUuid           uuid.UUID       `json:"as_uuid"`
	Variant          sql.NullInt32   `json:"variant"`
	LieuDung         sql.NullString  `json:"lieu_dung"`
	Quantity         int32           `json:"quantity"`
	ID_2             int32           `json:"id_2"`
	Name             string          `json:"name"`
	Code             string          `json:"code"`
	Barcode          sql.NullString  `json:"barcode"`
	DecisionNumber   sql.NullString  `json:"decision_number"`
	RegisterNumber   sql.NullString  `json:"register_number"`
	Longevity        sql.NullString  `json:"longevity"`
	Vat              sql.NullFloat64 `json:"vat"`
	Product          int32           `json:"product"`
	UserCreated      int32           `json:"user_created"`
	UserUpdated      sql.NullInt32   `json:"user_updated"`
	UpdatedAt        sql.NullTime    `json:"updated_at"`
	CreatedAt        time.Time       `json:"created_at"`
	InitialInventory int32           `json:"initial_inventory"`
	RealInventory    int32           `json:"real_inventory"`
	ID_3             sql.NullInt32   `json:"id_3"`
	Variant_2        sql.NullInt32   `json:"variant_2"`
	Media            sql.NullInt32   `json:"media"`
	ID_4             sql.NullInt32   `json:"id_4"`
	MediaUrl         sql.NullString  `json:"media_url"`
}

func (q *Queries) GetListScheduleDrug(ctx context.Context, asUuid uuid.UUID) ([]GetListScheduleDrugRow, error) {
	rows, err := q.db.QueryContext(ctx, getListScheduleDrug, asUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListScheduleDrugRow{}
	for rows.Next() {
		var i GetListScheduleDrugRow
		if err := rows.Scan(
			&i.ID,
			&i.AsUuid,
			&i.Variant,
			&i.LieuDung,
			&i.Quantity,
			&i.ID_2,
			&i.Name,
			&i.Code,
			&i.Barcode,
			&i.DecisionNumber,
			&i.RegisterNumber,
			&i.Longevity,
			&i.Vat,
			&i.Product,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.InitialInventory,
			&i.RealInventory,
			&i.ID_3,
			&i.Variant_2,
			&i.Media,
			&i.ID_4,
			&i.MediaUrl,
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

const getListScheduleService = `-- name: GetListScheduleService :many
SELECT ass.id, as_uuid, service, order_service, s.id, image, s.code, title, entity, staff, frequency, reminder_time, unit, price, s.description, s.company, s.user_created, s.user_updated, s.created_at, s.updated_at, os.id, os.code, total_price, os.description, vat, discount, service_price, must_paid, customer, address, status, type, ticket, qr, os.company, payment, os.user_created, os.user_updated, os.created_at, os.updated_at FROM appointment_schedule_service ass
JOIN services s ON s.id = ass.service
LEFT JOIN orders os ON os.id = ass.order_service
WHERE ass.as_uuid = $1
`

type GetListScheduleServiceRow struct {
	ID            int32           `json:"id"`
	AsUuid        uuid.UUID       `json:"as_uuid"`
	Service       sql.NullInt32   `json:"service"`
	OrderService  sql.NullInt32   `json:"order_service"`
	ID_2          int32           `json:"id_2"`
	Image         sql.NullInt32   `json:"image"`
	Code          string          `json:"code"`
	Title         string          `json:"title"`
	Entity        sql.NullString  `json:"entity"`
	Staff         int32           `json:"staff"`
	Frequency     sql.NullString  `json:"frequency"`
	ReminderTime  sql.NullInt32   `json:"reminder_time"`
	Unit          string          `json:"unit"`
	Price         float64         `json:"price"`
	Description   sql.NullString  `json:"description"`
	Company       int32           `json:"company"`
	UserCreated   int32           `json:"user_created"`
	UserUpdated   sql.NullInt32   `json:"user_updated"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     sql.NullTime    `json:"updated_at"`
	ID_3          sql.NullInt32   `json:"id_3"`
	Code_2        sql.NullString  `json:"code_2"`
	TotalPrice    sql.NullFloat64 `json:"total_price"`
	Description_2 sql.NullString  `json:"description_2"`
	Vat           sql.NullFloat64 `json:"vat"`
	Discount      sql.NullString  `json:"discount"`
	ServicePrice  sql.NullFloat64 `json:"service_price"`
	MustPaid      sql.NullFloat64 `json:"must_paid"`
	Customer      sql.NullInt32   `json:"customer"`
	Address       sql.NullInt32   `json:"address"`
	Status        sql.NullString  `json:"status"`
	Type          sql.NullString  `json:"type"`
	Ticket        sql.NullInt32   `json:"ticket"`
	Qr            sql.NullInt32   `json:"qr"`
	Company_2     sql.NullInt32   `json:"company_2"`
	Payment       sql.NullInt32   `json:"payment"`
	UserCreated_2 sql.NullInt32   `json:"user_created_2"`
	UserUpdated_2 sql.NullInt32   `json:"user_updated_2"`
	CreatedAt_2   sql.NullTime    `json:"created_at_2"`
	UpdatedAt_2   sql.NullTime    `json:"updated_at_2"`
}

func (q *Queries) GetListScheduleService(ctx context.Context, asUuid uuid.UUID) ([]GetListScheduleServiceRow, error) {
	rows, err := q.db.QueryContext(ctx, getListScheduleService, asUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListScheduleServiceRow{}
	for rows.Next() {
		var i GetListScheduleServiceRow
		if err := rows.Scan(
			&i.ID,
			&i.AsUuid,
			&i.Service,
			&i.OrderService,
			&i.ID_2,
			&i.Image,
			&i.Code,
			&i.Title,
			&i.Entity,
			&i.Staff,
			&i.Frequency,
			&i.ReminderTime,
			&i.Unit,
			&i.Price,
			&i.Description,
			&i.Company,
			&i.UserCreated,
			&i.UserUpdated,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_3,
			&i.Code_2,
			&i.TotalPrice,
			&i.Description_2,
			&i.Vat,
			&i.Discount,
			&i.ServicePrice,
			&i.MustPaid,
			&i.Customer,
			&i.Address,
			&i.Status,
			&i.Type,
			&i.Ticket,
			&i.Qr,
			&i.Company_2,
			&i.Payment,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
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

const getListScheduleUrl = `-- name: GetListScheduleUrl :many
SELECT id, as_uuid, url, name_doc FROM appointment_schedule_url
WHERE as_uuid = $1
`

func (q *Queries) GetListScheduleUrl(ctx context.Context, asUuid uuid.UUID) ([]AppointmentScheduleUrl, error) {
	rows, err := q.db.QueryContext(ctx, getListScheduleUrl, asUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AppointmentScheduleUrl{}
	for rows.Next() {
		var i AppointmentScheduleUrl
		if err := rows.Scan(
			&i.ID,
			&i.AsUuid,
			&i.Url,
			&i.NameDoc,
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

const updateSchedule = `-- name: UpdateSchedule :one
UPDATE appointment_schedules
SET
    is_done = COALESCE($1::bool, is_done),
    diagnostic = COALESCE($2::varchar, diagnostic)
WHERE uuid = $3::uuid
RETURNING id, uuid, code, customer, company, doctor, symptoms, diagnostic, qr_code_url, is_done, meeting_at, user_created, user_updated, created_at, updated_at
`

type UpdateScheduleParams struct {
	IsDone     sql.NullBool   `json:"is_done"`
	Diagnostic sql.NullString `json:"diagnostic"`
	Uuid       uuid.UUID      `json:"uuid"`
}

func (q *Queries) UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) (AppointmentSchedule, error) {
	row := q.db.QueryRowContext(ctx, updateSchedule, arg.IsDone, arg.Diagnostic, arg.Uuid)
	var i AppointmentSchedule
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
	)
	return i, err
}