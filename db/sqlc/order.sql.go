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

const countOrderByStatus = `-- name: CountOrderByStatus :one
SELECT COUNT(*) FROM orders
WHERE status = $1
`

func (q *Queries) CountOrderByStatus(ctx context.Context, status sql.NullString) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOrderByStatus, status)
	var count int64
	err := row.Scan(&count)
	return count, err
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

const detailOrder = `-- name: DetailOrder :one
SELECT o.id, o.code, total_price, description, vat, discount, service_price, must_paid, customer, address, status, o.type, ticket, qr, company, payment, user_created, user_updated, o.created_at, updated_at, m.id, media_url, ot.id, ot.code, ot.title, os.id, os.code, os.title, a.id, username, hashed_password, full_name, email, a.type, is_verify, password_changed_at, a.created_at, m.media_url AS qr_url, ot.id AS ot_id, ot.code AS ot_code, ot.title AS ot_title,
       os.id AS os_id, os.code AS os_code, os.title AS os_title,
       a.full_name AS a_full_name FROM orders o
JOIN medias m ON o.qr = m.id
JOIN order_type ot ON o.type = ot.code
JOIN order_status os ON o.status = os.code
JOIN accounts a ON o.user_created = a.id
WHERE o.id = $1
`

type DetailOrderRow struct {
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
	MediaUrl          string         `json:"media_url"`
	ID_3              int32          `json:"id_3"`
	Code_2            string         `json:"code_2"`
	Title             string         `json:"title"`
	ID_4              int32          `json:"id_4"`
	Code_3            string         `json:"code_3"`
	Title_2           string         `json:"title_2"`
	ID_5              int32          `json:"id_5"`
	Username          string         `json:"username"`
	HashedPassword    string         `json:"hashed_password"`
	FullName          string         `json:"full_name"`
	Email             string         `json:"email"`
	Type_2            int32          `json:"type_2"`
	IsVerify          bool           `json:"is_verify"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt_2       time.Time      `json:"created_at_2"`
	QrUrl             string         `json:"qr_url"`
	OtID              int32          `json:"ot_id"`
	OtCode            string         `json:"ot_code"`
	OtTitle           string         `json:"ot_title"`
	OsID              int32          `json:"os_id"`
	OsCode            string         `json:"os_code"`
	OsTitle           string         `json:"os_title"`
	AFullName         string         `json:"a_full_name"`
}

func (q *Queries) DetailOrder(ctx context.Context, id int32) (DetailOrderRow, error) {
	row := q.db.QueryRowContext(ctx, detailOrder, id)
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
SELECT o.id, o.code, o.total_price, description, vat, discount, service_price, must_paid, customer, o.address, o.status, o.type, ticket, o.qr, o.company, payment, o.user_created, o.user_updated, o.created_at, o.updated_at, c.id, c.full_name, c.code, c.company, c.address, c.email, phone, license, birthday, c.user_created, c.user_updated, c.updated_at, c.created_at, t.id, t.code, t.type, t.status, note, t.qr, export_to, import_from, t.total_price, warehouse, t.user_created, t.user_updated, t.updated_at, t.created_at, os.id, os.code, title, a.id, username, hashed_password, a.full_name, a.email, a.type, is_verify, password_changed_at, a.created_at, c.full_name AS c_full_name, os.title AS os_title, os.id AS os_id, a.full_name AS a_full_name FROM orders o
JOIN customers c ON o.customer = c.id
JOIN tickets t ON o.ticket = t.id
JOIN order_status os ON os.code = o.status
JOIN accounts a ON a.id = o.user_created
WHERE o.company = $1::int
AND (
    $2::varchar IS NULL OR o.status = $2::varchar
)
AND (
    $3::int IS NULL OR t.warehouse = $3::int
)
AND (
    o.code ILIKE '%' || COALESCE($4::varchar, '') || '%' OR
    c.full_name ILIKE '%' || COALESCE($4::varchar, '') || '%'
)
AND  ((
    $5::timestamp IS NULL AND $6::timestamp  IS NULL
) OR (
    ($5::timestamp IS NULL OR o.created_at >= $5::timestamp) AND
    ($6::timestamp IS NULL OR o.created_at <= $6::timestamp)
))
AND ((
    $7::timestamp IS NULL AND $8::timestamp  IS NULL
) OR (
    (o.updated_at >= $7::timestamp OR $7::timestamp  IS NULL) AND
    (o.updated_at <= $8::timestamp OR $8::timestamp  IS NULL)
))
ORDER BY
    CASE WHEN $9::varchar = 'created_at' THEN o.created_at END DESC,
    CASE WHEN $9::varchar = '-created_at' THEN o.created_at END ASC,
    CASE WHEN $9::varchar = 'updated_at' THEN o.updated_at END DESC,
    CASE WHEN $9::varchar = '-updated_at' THEN o.updated_at END ASC,
    CASE WHEN $9::varchar IS NULL THEN o.id END DESC
LIMIT COALESCE($11::int, 10)
OFFSET (COALESCE($10::int, 1) - 1) * COALESCE($11::int, 10)
`

type ListOrderParams struct {
	Company      sql.NullInt32  `json:"company"`
	Status       sql.NullString `json:"status"`
	Warehouse    sql.NullInt32  `json:"warehouse"`
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
	Title             string         `json:"title"`
	ID_5              int32          `json:"id_5"`
	Username          string         `json:"username"`
	HashedPassword    string         `json:"hashed_password"`
	FullName_2        string         `json:"full_name_2"`
	Email_2           string         `json:"email_2"`
	Type_3            int32          `json:"type_3"`
	IsVerify          bool           `json:"is_verify"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt_4       time.Time      `json:"created_at_4"`
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
			&i.Title,
			&i.ID_5,
			&i.Username,
			&i.HashedPassword,
			&i.FullName_2,
			&i.Email_2,
			&i.Type_3,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt_4,
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
SELECT oi.id, "order", oi.variant, value, total_price, consignment, consignment_log, v.id, name, v.code, barcode, decision_number, register_number, longevity, vat, product, v.user_created, v.user_updated, v.updated_at, v.created_at, c.id, c.code, quantity, inventory, ticket, expired_at, producted_at, is_available, c.user_created, c.user_updated, c.updated_at, c.created_at, c.variant, vm.id, vm.variant, media, m.id, media_url FROM order_items oi
JOIN variants v ON v.id = oi.variant
JOIN consignment c ON c.id = oi.consignment
JOIN variant_media vm ON vm.variant = v.id
JOIN medias m ON vm.media = m.id
WHERE oi.order = $1
`

type ListOrderItemRow struct {
	ID             int32         `json:"id"`
	Order          int32         `json:"order"`
	Variant        int32         `json:"variant"`
	Value          int32         `json:"value"`
	TotalPrice     float64       `json:"total_price"`
	Consignment    sql.NullInt32 `json:"consignment"`
	ConsignmentLog sql.NullInt32 `json:"consignment_log"`
	ID_2           int32         `json:"id_2"`
	Name           string        `json:"name"`
	Code           string        `json:"code"`
	Barcode        string        `json:"barcode"`
	DecisionNumber string        `json:"decision_number"`
	RegisterNumber string        `json:"register_number"`
	Longevity      string        `json:"longevity"`
	Vat            float64       `json:"vat"`
	Product        int32         `json:"product"`
	UserCreated    int32         `json:"user_created"`
	UserUpdated    sql.NullInt32 `json:"user_updated"`
	UpdatedAt      sql.NullTime  `json:"updated_at"`
	CreatedAt      time.Time     `json:"created_at"`
	ID_3           int32         `json:"id_3"`
	Code_2         string        `json:"code_2"`
	Quantity       int32         `json:"quantity"`
	Inventory      int32         `json:"inventory"`
	Ticket         sql.NullInt32 `json:"ticket"`
	ExpiredAt      time.Time     `json:"expired_at"`
	ProductedAt    time.Time     `json:"producted_at"`
	IsAvailable    bool          `json:"is_available"`
	UserCreated_2  sql.NullInt32 `json:"user_created_2"`
	UserUpdated_2  sql.NullInt32 `json:"user_updated_2"`
	UpdatedAt_2    sql.NullTime  `json:"updated_at_2"`
	CreatedAt_2    time.Time     `json:"created_at_2"`
	Variant_2      sql.NullInt32 `json:"variant_2"`
	ID_4           int32         `json:"id_4"`
	Variant_3      int32         `json:"variant_3"`
	Media          int32         `json:"media"`
	ID_5           int32         `json:"id_5"`
	MediaUrl       string        `json:"media_url"`
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
