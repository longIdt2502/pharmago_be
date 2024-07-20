// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: order.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const countOrder = `-- name: CountOrder :one
SELECT COALESCE(COUNT(id), 0)::int FROM orders
WHERE company = $1
`

func (q *Queries) CountOrder(ctx context.Context, company int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, countOrder, company)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const countOrderByStatus = `-- name: CountOrderByStatus :many
SELECT os.code, COALESCE(COUNT(os.code), 0)::int AS count FROM order_status os
RIGHT JOIN orders o ON os.code = o.status
WHERE o.company = $1
GROUP BY os.code
`

type CountOrderByStatusRow struct {
	Code  sql.NullString `json:"code"`
	Count int32          `json:"count"`
}

func (q *Queries) CountOrderByStatus(ctx context.Context, company int32) ([]CountOrderByStatusRow, error) {
	rows, err := q.db.QueryContext(ctx, countOrderByStatus, company)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CountOrderByStatusRow{}
	for rows.Next() {
		var i CountOrderByStatusRow
		if err := rows.Scan(&i.Code, &i.Count); err != nil {
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

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    code, total_price, description, vat, discount, service_price,
    must_paid, customer, address, status, type, ticket, qr,
    company, payment, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
) RETURNING id, code, total_price, description, vat, discount, service_price, must_paid, customer, address, status, type, ticket, qr, company, payment, user_created, user_updated, created_at, updated_at
`

type CreateOrderParams struct {
	Code         string         `json:"code"`
	TotalPrice   float64        `json:"total_price"`
	Description  sql.NullString `json:"description"`
	Vat          float64        `json:"vat"`
	Discount     string         `json:"discount"`
	ServicePrice float64        `json:"service_price"`
	MustPaid     float64        `json:"must_paid"`
	Customer     sql.NullInt32  `json:"customer"`
	Address      sql.NullInt32  `json:"address"`
	Status       sql.NullString `json:"status"`
	Type         sql.NullString `json:"type"`
	Ticket       sql.NullInt32  `json:"ticket"`
	Qr           sql.NullInt32  `json:"qr"`
	Company      int32          `json:"company"`
	Payment      int32          `json:"payment"`
	UserCreated  sql.NullInt32  `json:"user_created"`
	UserUpdated  sql.NullInt32  `json:"user_updated"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.Code,
		arg.TotalPrice,
		arg.Description,
		arg.Vat,
		arg.Discount,
		arg.ServicePrice,
		arg.MustPaid,
		arg.Customer,
		arg.Address,
		arg.Status,
		arg.Type,
		arg.Ticket,
		arg.Qr,
		arg.Company,
		arg.Payment,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.TotalPrice,
		&i.Description,
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
		&i.Company,
		&i.Payment,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (
    "order", variant, value, total_price, consignment, consignment_log
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, "order", variant, value, total_price, consignment, consignment_log
`

type CreateOrderItemParams struct {
	Order          int32         `json:"order"`
	Variant        int32         `json:"variant"`
	Value          int32         `json:"value"`
	TotalPrice     float64       `json:"total_price"`
	Consignment    sql.NullInt32 `json:"consignment"`
	ConsignmentLog sql.NullInt32 `json:"consignment_log"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.Order,
		arg.Variant,
		arg.Value,
		arg.TotalPrice,
		arg.Consignment,
		arg.ConsignmentLog,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.Order,
		&i.Variant,
		&i.Value,
		&i.TotalPrice,
		&i.Consignment,
		&i.ConsignmentLog,
	)
	return i, err
}

const createOrderServiceItem = `-- name: CreateOrderServiceItem :one
INSERT INTO service_order_item (
    "order", service, unit_price, total_price, discount
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, "order", service, unit_price, discount, total_price, quantity
`

type CreateOrderServiceItemParams struct {
	Order      int32         `json:"order"`
	Service    sql.NullInt32 `json:"service"`
	UnitPrice  float64       `json:"unit_price"`
	TotalPrice float64       `json:"total_price"`
	Discount   float64       `json:"discount"`
}

func (q *Queries) CreateOrderServiceItem(ctx context.Context, arg CreateOrderServiceItemParams) (ServiceOrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderServiceItem,
		arg.Order,
		arg.Service,
		arg.UnitPrice,
		arg.TotalPrice,
		arg.Discount,
	)
	var i ServiceOrderItem
	err := row.Scan(
		&i.ID,
		&i.Order,
		&i.Service,
		&i.UnitPrice,
		&i.Discount,
		&i.TotalPrice,
		&i.Quantity,
	)
	return i, err
}

const detailOrder = `-- name: DetailOrder :one
SELECT o.id, o.code, total_price, description, vat, discount, service_price, must_paid, customer, o.address, status, o.type, ticket, qr, company, payment, user_created, user_updated, o.created_at, updated_at, m.id, media_url, ot.id, ot.code, ot.title, os.id, os.code, os.title, a.id, a.username, a.hashed_password, a.full_name, a.email, a.type, a.is_verify, a.password_changed_at, a.created_at, a.role, a.gender, a.licence, a.dob, a.address, uu.id, uu.username, uu.hashed_password, uu.full_name, uu.email, uu.type, uu.is_verify, uu.password_changed_at, uu.created_at, uu.role, uu.gender, uu.licence, uu.dob, uu.address, m.media_url AS qr_url, ot.id AS ot_id, ot.code AS ot_code, ot.title AS ot_title,
       os.id AS os_id, os.code AS os_code, os.title AS os_title,
       a.full_name AS a_full_name FROM orders o
JOIN medias m ON o.qr = m.id
JOIN order_type ot ON o.type = ot.code
JOIN order_status os ON o.status = os.code
JOIN accounts a ON o.user_created = a.id
JOIN accounts uu ON o.user_updated = uu.id
WHERE (o.id = $1 OR o.code = $2)
`

type DetailOrderParams struct {
	ID   sql.NullInt32  `json:"id"`
	Code sql.NullString `json:"code"`
}

type DetailOrderRow struct {
	ID                  int32          `json:"id"`
	Code                string         `json:"code"`
	TotalPrice          float64        `json:"total_price"`
	Description         sql.NullString `json:"description"`
	Vat                 float64        `json:"vat"`
	Discount            string         `json:"discount"`
	ServicePrice        float64        `json:"service_price"`
	MustPaid            float64        `json:"must_paid"`
	Customer            sql.NullInt32  `json:"customer"`
	Address             sql.NullInt32  `json:"address"`
	Status              sql.NullString `json:"status"`
	Type                sql.NullString `json:"type"`
	Ticket              sql.NullInt32  `json:"ticket"`
	Qr                  sql.NullInt32  `json:"qr"`
	Company             int32          `json:"company"`
	Payment             int32          `json:"payment"`
	UserCreated         sql.NullInt32  `json:"user_created"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	ID_2                int32          `json:"id_2"`
	MediaUrl            string         `json:"media_url"`
	ID_3                int32          `json:"id_3"`
	Code_2              string         `json:"code_2"`
	Title               string         `json:"title"`
	ID_4                int32          `json:"id_4"`
	Code_3              string         `json:"code_3"`
	Title_2             string         `json:"title_2"`
	ID_5                int32          `json:"id_5"`
	Username            string         `json:"username"`
	HashedPassword      string         `json:"hashed_password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	Type_2              int32          `json:"type_2"`
	IsVerify            bool           `json:"is_verify"`
	PasswordChangedAt   time.Time      `json:"password_changed_at"`
	CreatedAt_2         time.Time      `json:"created_at_2"`
	Role                sql.NullInt32  `json:"role"`
	Gender              NullGender     `json:"gender"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address_2           sql.NullInt32  `json:"address_2"`
	ID_6                int32          `json:"id_6"`
	Username_2          string         `json:"username_2"`
	HashedPassword_2    string         `json:"hashed_password_2"`
	FullName_2          string         `json:"full_name_2"`
	Email_2             string         `json:"email_2"`
	Type_3              int32          `json:"type_3"`
	IsVerify_2          bool           `json:"is_verify_2"`
	PasswordChangedAt_2 time.Time      `json:"password_changed_at_2"`
	CreatedAt_3         time.Time      `json:"created_at_3"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_3           sql.NullInt32  `json:"address_3"`
	QrUrl               string         `json:"qr_url"`
	OtID                int32          `json:"ot_id"`
	OtCode              string         `json:"ot_code"`
	OtTitle             string         `json:"ot_title"`
	OsID                int32          `json:"os_id"`
	OsCode              string         `json:"os_code"`
	OsTitle             string         `json:"os_title"`
	AFullName           string         `json:"a_full_name"`
}

func (q *Queries) DetailOrder(ctx context.Context, arg DetailOrderParams) (DetailOrderRow, error) {
	row := q.db.QueryRowContext(ctx, detailOrder, arg.ID, arg.Code)
	var i DetailOrderRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.TotalPrice,
		&i.Description,
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
		&i.Company,
		&i.Payment,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.MediaUrl,
		&i.ID_3,
		&i.Code_2,
		&i.Title,
		&i.ID_4,
		&i.Code_3,
		&i.Title_2,
		&i.ID_5,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type_2,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt_2,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address_2,
		&i.ID_6,
		&i.Username_2,
		&i.HashedPassword_2,
		&i.FullName_2,
		&i.Email_2,
		&i.Type_3,
		&i.IsVerify_2,
		&i.PasswordChangedAt_2,
		&i.CreatedAt_3,
		&i.Role_2,
		&i.Gender_2,
		&i.Licence_2,
		&i.Dob_2,
		&i.Address_3,
		&i.QrUrl,
		&i.OtID,
		&i.OtCode,
		&i.OtTitle,
		&i.OsID,
		&i.OsCode,
		&i.OsTitle,
		&i.AFullName,
	)
	return i, err
}

const listOrder = `-- name: ListOrder :many
SELECT o.id, o.code, o.total_price, description, vat, discount, service_price, o.must_paid, customer, o.address, o.status, o.type, ticket, o.qr, o.company, payment, o.user_created, o.user_updated, o.created_at, o.updated_at, c.id, c.full_name, c.code, c.company, c.address, c.email, phone, license, birthday, c.user_created, c.user_updated, c.updated_at, c.created_at, "group", c.title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number, bank_name, bank_branch, issued_by, c.gender, t.id, t.code, t.type, t.status, note, t.qr, export_to, import_from, t.total_price, warehouse, t.user_created, t.user_updated, t.updated_at, t.created_at, os.id, os.code, os.title, a.id, username, hashed_password, a.full_name, a.email, a.type, is_verify, password_changed_at, a.created_at, role, a.gender, licence, dob, a.address, p.id, p.code, p.must_paid, had_paid, need_pay, c.full_name AS c_full_name, os.title AS os_title, os.id AS os_id, a.full_name AS a_full_name FROM orders o
JOIN customers c ON o.customer = c.id
JOIN tickets t ON o.ticket = t.id
JOIN order_status os ON os.code = o.status
JOIN accounts a ON a.id = o.user_created
JOIN payments p ON p.id = o.payment
WHERE o.company = $1::int
AND (
    $2::varchar IS NULL OR o.status = $2::varchar
)
AND (
    $3::int IS NULL OR t.warehouse = $3::int
)
AND (
    $4::int IS NULL OR o.customer = $4::int
)
AND (
    o.code ILIKE '%' || COALESCE($5::varchar, '') || '%' OR
    c.full_name ILIKE '%' || COALESCE($5::varchar, '') || '%'
)
AND  ((
    $6::timestamp IS NULL AND $7::timestamp  IS NULL
) OR (
    ($6::timestamp IS NULL OR o.created_at >= $6::timestamp) AND
    ($7::timestamp IS NULL OR o.created_at <= $7::timestamp)
))
AND ((
    $8::timestamp IS NULL AND $9::timestamp  IS NULL
) OR (
    (o.updated_at >= $8::timestamp OR $8::timestamp  IS NULL) AND
    (o.updated_at <= $9::timestamp OR $9::timestamp  IS NULL)
))
ORDER BY
    CASE WHEN $10::varchar = 'created_at' THEN o.created_at END DESC,
    CASE WHEN $10::varchar = '-created_at' THEN o.created_at END ASC,
    CASE WHEN $10::varchar = 'updated_at' THEN o.updated_at END DESC,
    CASE WHEN $10::varchar = '-updated_at' THEN o.updated_at END ASC,
    CASE WHEN $10::varchar IS NULL THEN o.id END DESC
LIMIT COALESCE($12::int, 10)
OFFSET (COALESCE($11::int, 1) - 1) * COALESCE($12::int, 10)
`

type ListOrderParams struct {
	Company      sql.NullInt32  `json:"company"`
	Status       sql.NullString `json:"status"`
	Warehouse    sql.NullInt32  `json:"warehouse"`
	Customer     sql.NullInt32  `json:"customer"`
	Search       sql.NullString `json:"search"`
	CreatedStart sql.NullTime   `json:"created_start"`
	CreatedEnd   sql.NullTime   `json:"created_end"`
	UpdatedStart sql.NullTime   `json:"updated_start"`
	UpdatedEnd   sql.NullTime   `json:"updated_end"`
	OrderBy      sql.NullString `json:"order_by"`
	Page         sql.NullInt32  `json:"page"`
	Limit        sql.NullInt32  `json:"limit"`
}

type ListOrderRow struct {
	ID                int32          `json:"id"`
	Code              string         `json:"code"`
	TotalPrice        float64        `json:"total_price"`
	Description       sql.NullString `json:"description"`
	Vat               float64        `json:"vat"`
	Discount          string         `json:"discount"`
	ServicePrice      float64        `json:"service_price"`
	MustPaid          float64        `json:"must_paid"`
	Customer          sql.NullInt32  `json:"customer"`
	Address           sql.NullInt32  `json:"address"`
	Status            sql.NullString `json:"status"`
	Type              sql.NullString `json:"type"`
	Ticket            sql.NullInt32  `json:"ticket"`
	Qr                sql.NullInt32  `json:"qr"`
	Company           int32          `json:"company"`
	Payment           int32          `json:"payment"`
	UserCreated       sql.NullInt32  `json:"user_created"`
	UserUpdated       sql.NullInt32  `json:"user_updated"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         sql.NullTime   `json:"updated_at"`
	ID_2              int32          `json:"id_2"`
	FullName          string         `json:"full_name"`
	Code_2            string         `json:"code_2"`
	Company_2         int32          `json:"company_2"`
	Address_2         sql.NullInt32  `json:"address_2"`
	Email             sql.NullString `json:"email"`
	Phone             sql.NullString `json:"phone"`
	License           sql.NullString `json:"license"`
	Birthday          sql.NullTime   `json:"birthday"`
	UserCreated_2     int32          `json:"user_created_2"`
	UserUpdated_2     sql.NullInt32  `json:"user_updated_2"`
	UpdatedAt_2       sql.NullTime   `json:"updated_at_2"`
	CreatedAt_2       time.Time      `json:"created_at_2"`
	Group             sql.NullInt32  `json:"group"`
	Title             sql.NullString `json:"title"`
	LicenseDate       sql.NullTime   `json:"license_date"`
	ContactName       sql.NullString `json:"contact_name"`
	ContactTitle      sql.NullString `json:"contact_title"`
	ContactPhone      sql.NullString `json:"contact_phone"`
	ContactEmail      sql.NullString `json:"contact_email"`
	ContactAddress    sql.NullInt32  `json:"contact_address"`
	AccountNumber     sql.NullString `json:"account_number"`
	BankName          sql.NullString `json:"bank_name"`
	BankBranch        sql.NullString `json:"bank_branch"`
	IssuedBy          sql.NullString `json:"issued_by"`
	Gender            NullGender     `json:"gender"`
	ID_3              int32          `json:"id_3"`
	Code_3            string         `json:"code_3"`
	Type_2            sql.NullInt32  `json:"type_2"`
	Status_2          sql.NullInt32  `json:"status_2"`
	Note              sql.NullString `json:"note"`
	Qr_2              sql.NullInt32  `json:"qr_2"`
	ExportTo          sql.NullInt32  `json:"export_to"`
	ImportFrom        sql.NullInt32  `json:"import_from"`
	TotalPrice_2      float64        `json:"total_price_2"`
	Warehouse         int32          `json:"warehouse"`
	UserCreated_3     int32          `json:"user_created_3"`
	UserUpdated_3     sql.NullInt32  `json:"user_updated_3"`
	UpdatedAt_3       sql.NullTime   `json:"updated_at_3"`
	CreatedAt_3       time.Time      `json:"created_at_3"`
	ID_4              int32          `json:"id_4"`
	Code_4            string         `json:"code_4"`
	Title_2           string         `json:"title_2"`
	ID_5              int32          `json:"id_5"`
	Username          string         `json:"username"`
	HashedPassword    string         `json:"hashed_password"`
	FullName_2        string         `json:"full_name_2"`
	Email_2           string         `json:"email_2"`
	Type_3            int32          `json:"type_3"`
	IsVerify          bool           `json:"is_verify"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt_4       time.Time      `json:"created_at_4"`
	Role              sql.NullInt32  `json:"role"`
	Gender_2          NullGender     `json:"gender_2"`
	Licence           sql.NullString `json:"licence"`
	Dob               sql.NullTime   `json:"dob"`
	Address_3         sql.NullInt32  `json:"address_3"`
	ID_6              int32          `json:"id_6"`
	Code_5            string         `json:"code_5"`
	MustPaid_2        float64        `json:"must_paid_2"`
	HadPaid           float64        `json:"had_paid"`
	NeedPay           float64        `json:"need_pay"`
	CFullName         string         `json:"c_full_name"`
	OsTitle           string         `json:"os_title"`
	OsID              int32          `json:"os_id"`
	AFullName         string         `json:"a_full_name"`
}

func (q *Queries) ListOrder(ctx context.Context, arg ListOrderParams) ([]ListOrderRow, error) {
	rows, err := q.db.QueryContext(ctx, listOrder,
		arg.Company,
		arg.Status,
		arg.Warehouse,
		arg.Customer,
		arg.Search,
		arg.CreatedStart,
		arg.CreatedEnd,
		arg.UpdatedStart,
		arg.UpdatedEnd,
		arg.OrderBy,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListOrderRow{}
	for rows.Next() {
		var i ListOrderRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.TotalPrice,
			&i.Description,
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
			&i.Company,
			&i.Payment,
			&i.UserCreated,
			&i.UserUpdated,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.FullName,
			&i.Code_2,
			&i.Company_2,
			&i.Address_2,
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
			&i.Code_3,
			&i.Type_2,
			&i.Status_2,
			&i.Note,
			&i.Qr_2,
			&i.ExportTo,
			&i.ImportFrom,
			&i.TotalPrice_2,
			&i.Warehouse,
			&i.UserCreated_3,
			&i.UserUpdated_3,
			&i.UpdatedAt_3,
			&i.CreatedAt_3,
			&i.ID_4,
			&i.Code_4,
			&i.Title_2,
			&i.ID_5,
			&i.Username,
			&i.HashedPassword,
			&i.FullName_2,
			&i.Email_2,
			&i.Type_3,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt_4,
			&i.Role,
			&i.Gender_2,
			&i.Licence,
			&i.Dob,
			&i.Address_3,
			&i.ID_6,
			&i.Code_5,
			&i.MustPaid_2,
			&i.HadPaid,
			&i.NeedPay,
			&i.CFullName,
			&i.OsTitle,
			&i.OsID,
			&i.AFullName,
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

const listOrderItem = `-- name: ListOrderItem :many
SELECT oi.id, "order", oi.variant, value, total_price, consignment, consignment_log, v.id, name, v.code, barcode, decision_number, register_number, longevity, vat, product, v.user_created, v.user_updated, v.updated_at, v.created_at, initial_inventory, real_inventory, c.id, c.code, quantity, inventory, ticket, expired_at, producted_at, is_available, c.user_created, c.user_updated, c.updated_at, c.created_at, c.variant, vm.id, vm.variant, media, m.id, media_url FROM order_items oi
JOIN variants v ON v.id = oi.variant
JOIN consignment c ON c.id = oi.consignment
JOIN variant_media vm ON vm.variant = v.id
JOIN medias m ON vm.media = m.id
WHERE oi.order = $1
`

type ListOrderItemRow struct {
	ID               int32           `json:"id"`
	Order            int32           `json:"order"`
	Variant          int32           `json:"variant"`
	Value            int32           `json:"value"`
	TotalPrice       float64         `json:"total_price"`
	Consignment      sql.NullInt32   `json:"consignment"`
	ConsignmentLog   sql.NullInt32   `json:"consignment_log"`
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
	ID_3             int32           `json:"id_3"`
	Code_2           string          `json:"code_2"`
	Quantity         int32           `json:"quantity"`
	Inventory        int32           `json:"inventory"`
	Ticket           sql.NullInt32   `json:"ticket"`
	ExpiredAt        time.Time       `json:"expired_at"`
	ProductedAt      time.Time       `json:"producted_at"`
	IsAvailable      bool            `json:"is_available"`
	UserCreated_2    sql.NullInt32   `json:"user_created_2"`
	UserUpdated_2    sql.NullInt32   `json:"user_updated_2"`
	UpdatedAt_2      sql.NullTime    `json:"updated_at_2"`
	CreatedAt_2      time.Time       `json:"created_at_2"`
	Variant_2        sql.NullInt32   `json:"variant_2"`
	ID_4             int32           `json:"id_4"`
	Variant_3        int32           `json:"variant_3"`
	Media            int32           `json:"media"`
	ID_5             int32           `json:"id_5"`
	MediaUrl         string          `json:"media_url"`
}

func (q *Queries) ListOrderItem(ctx context.Context, order int32) ([]ListOrderItemRow, error) {
	rows, err := q.db.QueryContext(ctx, listOrderItem, order)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListOrderItemRow{}
	for rows.Next() {
		var i ListOrderItemRow
		if err := rows.Scan(
			&i.ID,
			&i.Order,
			&i.Variant,
			&i.Value,
			&i.TotalPrice,
			&i.Consignment,
			&i.ConsignmentLog,
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
			&i.Code_2,
			&i.Quantity,
			&i.Inventory,
			&i.Ticket,
			&i.ExpiredAt,
			&i.ProductedAt,
			&i.IsAvailable,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
			&i.Variant_2,
			&i.ID_4,
			&i.Variant_3,
			&i.Media,
			&i.ID_5,
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

const listOrderServiceItem = `-- name: ListOrderServiceItem :many
SELECT soi.id, "order", service, unit_price, discount, total_price, quantity, s.id, image, code, title, entity, staff, frequency, reminder_time, unit, price, description, company, user_created, user_updated, created_at, updated_at FROM service_order_item soi
JOIN services s ON s.id = soi.service
WHERE soi.order = $1
`

type ListOrderServiceItemRow struct {
	ID           int32          `json:"id"`
	Order        int32          `json:"order"`
	Service      sql.NullInt32  `json:"service"`
	UnitPrice    float64        `json:"unit_price"`
	Discount     float64        `json:"discount"`
	TotalPrice   float64        `json:"total_price"`
	Quantity     sql.NullInt32  `json:"quantity"`
	ID_2         int32          `json:"id_2"`
	Image        sql.NullInt32  `json:"image"`
	Code         string         `json:"code"`
	Title        string         `json:"title"`
	Entity       sql.NullString `json:"entity"`
	Staff        int32          `json:"staff"`
	Frequency    sql.NullString `json:"frequency"`
	ReminderTime sql.NullInt32  `json:"reminder_time"`
	Unit         string         `json:"unit"`
	Price        float64        `json:"price"`
	Description  sql.NullString `json:"description"`
	Company      int32          `json:"company"`
	UserCreated  int32          `json:"user_created"`
	UserUpdated  sql.NullInt32  `json:"user_updated"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

func (q *Queries) ListOrderServiceItem(ctx context.Context, order int32) ([]ListOrderServiceItemRow, error) {
	rows, err := q.db.QueryContext(ctx, listOrderServiceItem, order)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListOrderServiceItemRow{}
	for rows.Next() {
		var i ListOrderServiceItemRow
		if err := rows.Scan(
			&i.ID,
			&i.Order,
			&i.Service,
			&i.UnitPrice,
			&i.Discount,
			&i.TotalPrice,
			&i.Quantity,
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

const updateStatusOrder = `-- name: UpdateStatusOrder :one
UPDATE orders
SET status = $1::varchar
WHERE id = $2::int
RETURNING id, code, total_price, description, vat, discount, service_price, must_paid, customer, address, status, type, ticket, qr, company, payment, user_created, user_updated, created_at, updated_at
`

type UpdateStatusOrderParams struct {
	Status string `json:"status"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdateStatusOrder(ctx context.Context, arg UpdateStatusOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateStatusOrder, arg.Status, arg.ID)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.TotalPrice,
		&i.Description,
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
		&i.Company,
		&i.Payment,
		&i.UserCreated,
		&i.UserUpdated,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
