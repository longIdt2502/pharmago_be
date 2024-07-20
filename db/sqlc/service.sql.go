// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: service.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createService = `-- name: CreateService :one
INSERT INTO services (
    code, image, title, entity, staff, frequency, unit, price, description, company, user_created, user_updated, reminder_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING id, image, code, title, entity, staff, frequency, reminder_time, unit, price, description, company, user_created, user_updated, created_at, updated_at
`

type CreateServiceParams struct {
	Code         string         `json:"code"`
	Image        sql.NullInt32  `json:"image"`
	Title        string         `json:"title"`
	Entity       sql.NullString `json:"entity"`
	Staff        int32          `json:"staff"`
	Frequency    sql.NullString `json:"frequency"`
	Unit         string         `json:"unit"`
	Price        float64        `json:"price"`
	Description  sql.NullString `json:"description"`
	Company      int32          `json:"company"`
	UserCreated  int32          `json:"user_created"`
	UserUpdated  sql.NullInt32  `json:"user_updated"`
	ReminderTime sql.NullInt32  `json:"reminder_time"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, createService,
		arg.Code,
		arg.Image,
		arg.Title,
		arg.Entity,
		arg.Staff,
		arg.Frequency,
		arg.Unit,
		arg.Price,
		arg.Description,
		arg.Company,
		arg.UserCreated,
		arg.UserUpdated,
		arg.ReminderTime,
	)
	var i Service
	err := row.Scan(
		&i.ID,
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
	)
	return i, err
}

const createServiceVariant = `-- name: CreateServiceVariant :one
INSERT INTO service_variant (
    service, variant
) VALUES (
    $1, $2
) RETURNING id, service, variant
`

type CreateServiceVariantParams struct {
	Service sql.NullInt32 `json:"service"`
	Variant sql.NullInt32 `json:"variant"`
}

func (q *Queries) CreateServiceVariant(ctx context.Context, arg CreateServiceVariantParams) (ServiceVariant, error) {
	row := q.db.QueryRowContext(ctx, createServiceVariant, arg.Service, arg.Variant)
	var i ServiceVariant
	err := row.Scan(&i.ID, &i.Service, &i.Variant)
	return i, err
}

const deleteService = `-- name: DeleteService :one
DELETE FROM services
WHERE id = $1
RETURNING id, image, code, title, entity, staff, frequency, reminder_time, unit, price, description, company, user_created, user_updated, created_at, updated_at
`

func (q *Queries) DeleteService(ctx context.Context, id int32) (Service, error) {
	row := q.db.QueryRowContext(ctx, deleteService, id)
	var i Service
	err := row.Scan(
		&i.ID,
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
	)
	return i, err
}

const deleteServiceVariant = `-- name: DeleteServiceVariant :one
DELETE FROM service_variant
WHERE id = $1
RETURNING id, service, variant
`

func (q *Queries) DeleteServiceVariant(ctx context.Context, id int32) (ServiceVariant, error) {
	row := q.db.QueryRowContext(ctx, deleteServiceVariant, id)
	var i ServiceVariant
	err := row.Scan(&i.ID, &i.Service, &i.Variant)
	return i, err
}

const detailService = `-- name: DetailService :one
SELECT id, image, code, title, entity, staff, frequency, reminder_time, unit, price, description, company, user_created, user_updated, created_at, updated_at FROM services
WHERE id = $1
`

func (q *Queries) DetailService(ctx context.Context, id int32) (Service, error) {
	row := q.db.QueryRowContext(ctx, detailService, id)
	var i Service
	err := row.Scan(
		&i.ID,
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
	)
	return i, err
}

const getListService = `-- name: GetListService :many
WITH quantity_use AS (
    SELECT "service", COUNT("service") as quantity_use FROM service_order_item
    GROUP BY "service"
)
SELECT id, image, code, title, entity, staff, frequency, reminder_time, unit, price, description, company, user_created, user_updated, created_at, updated_at, service, quantity_use FROM services s
LEFT JOIN quantity_use qu ON s.id = qu.service
WHERE s.company = $1::int
AND (
    s.title ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    s.code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -s.id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type GetListServiceParams struct {
	Company sql.NullInt32  `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type GetListServiceRow struct {
	ID           int32          `json:"id"`
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
	Service      sql.NullInt32  `json:"service"`
	QuantityUse  sql.NullInt64  `json:"quantity_use"`
}

func (q *Queries) GetListService(ctx context.Context, arg GetListServiceParams) ([]GetListServiceRow, error) {
	rows, err := q.db.QueryContext(ctx, getListService,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListServiceRow{}
	for rows.Next() {
		var i GetListServiceRow
		if err := rows.Scan(
			&i.ID,
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
			&i.Service,
			&i.QuantityUse,
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

const getServicesByCustomer = `-- name: GetServicesByCustomer :many
SELECT s.id, s.image, s.code, s.title, s.entity, s.staff, s.frequency, s.reminder_time, s.unit, s.price, s.description, s.company, s.user_created, s.user_updated, s.created_at, s.updated_at, SUM(quantity) as quantity_use FROM service_order_item soi
JOIN orders o ON o.id = soi.order
JOIN services s ON s.id = soi.service
WHERE o.customer = $1::int
GROUP BY soi.service, s.id
LIMIT COALESCE($3::int, 10)
OFFSET (COALESCE($2::int, 1) - 1) * COALESCE($3::int, 10)
`

type GetServicesByCustomerParams struct {
	Customer int32         `json:"customer"`
	Page     sql.NullInt32 `json:"page"`
	Limit    sql.NullInt32 `json:"limit"`
}

type GetServicesByCustomerRow struct {
	ID           int32          `json:"id"`
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
	QuantityUse  int64          `json:"quantity_use"`
}

func (q *Queries) GetServicesByCustomer(ctx context.Context, arg GetServicesByCustomerParams) ([]GetServicesByCustomerRow, error) {
	rows, err := q.db.QueryContext(ctx, getServicesByCustomer, arg.Customer, arg.Page, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetServicesByCustomerRow{}
	for rows.Next() {
		var i GetServicesByCustomerRow
		if err := rows.Scan(
			&i.ID,
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
			&i.QuantityUse,
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

const listServiceVariant = `-- name: ListServiceVariant :many
SELECT sv.id, service, variant, v.id, name, code, barcode, decision_number, register_number, longevity, vat, product, user_created, user_updated, updated_at, created_at, initial_inventory, real_inventory FROM service_variant sv
LEFT JOIN variants v ON v.id = sv.variant
WHERE sv.service = $1::int
`

type ListServiceVariantRow struct {
	ID               int32           `json:"id"`
	Service          sql.NullInt32   `json:"service"`
	Variant          sql.NullInt32   `json:"variant"`
	ID_2             sql.NullInt32   `json:"id_2"`
	Name             sql.NullString  `json:"name"`
	Code             sql.NullString  `json:"code"`
	Barcode          sql.NullString  `json:"barcode"`
	DecisionNumber   sql.NullString  `json:"decision_number"`
	RegisterNumber   sql.NullString  `json:"register_number"`
	Longevity        sql.NullString  `json:"longevity"`
	Vat              sql.NullFloat64 `json:"vat"`
	Product          sql.NullInt32   `json:"product"`
	UserCreated      sql.NullInt32   `json:"user_created"`
	UserUpdated      sql.NullInt32   `json:"user_updated"`
	UpdatedAt        sql.NullTime    `json:"updated_at"`
	CreatedAt        sql.NullTime    `json:"created_at"`
	InitialInventory sql.NullInt32   `json:"initial_inventory"`
	RealInventory    sql.NullInt32   `json:"real_inventory"`
}

func (q *Queries) ListServiceVariant(ctx context.Context, id int32) ([]ListServiceVariantRow, error) {
	rows, err := q.db.QueryContext(ctx, listServiceVariant, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListServiceVariantRow{}
	for rows.Next() {
		var i ListServiceVariantRow
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.Variant,
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

const servicesUsedByCustomer = `-- name: ServicesUsedByCustomer :many
SELECT s.id, s.image, s.code, s.title, s.entity, s.staff, s.frequency, s.reminder_time, s.unit, s.price, s.description, s.company, s.user_created, s.user_updated, s.created_at, s.updated_at, COUNT(s.id) AS number_of_uses FROM service_order_item soi
JOIN orders o ON o.id = soi.order
JOIN services s ON s.id = soi.service
WHERE o.customer = $1::int
GROUP BY s.id
LIMIT COALESCE($3::int, 10)
OFFSET (COALESCE($2::int, 1) - 1) * COALESCE($3::int, 10)
`

type ServicesUsedByCustomerParams struct {
	Customer int32         `json:"customer"`
	Page     sql.NullInt32 `json:"page"`
	Limit    sql.NullInt32 `json:"limit"`
}

type ServicesUsedByCustomerRow struct {
	ID           int32          `json:"id"`
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
	NumberOfUses int64          `json:"number_of_uses"`
}

func (q *Queries) ServicesUsedByCustomer(ctx context.Context, arg ServicesUsedByCustomerParams) ([]ServicesUsedByCustomerRow, error) {
	rows, err := q.db.QueryContext(ctx, servicesUsedByCustomer, arg.Customer, arg.Page, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ServicesUsedByCustomerRow{}
	for rows.Next() {
		var i ServicesUsedByCustomerRow
		if err := rows.Scan(
			&i.ID,
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
			&i.NumberOfUses,
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

const updateService = `-- name: UpdateService :one
UPDATE services
SET
    image = COALESCE($1::int, image),
    title = COALESCE($2::varchar, title),
    entity = COALESCE($3::varchar, entity),
    staff = COALESCE($4::int, staff),
    frequency = COALESCE($5::varchar, frequency),
    unit = COALESCE($6::varchar, unit),
    price = COALESCE($7::float, price),
    description = COALESCE($8::varchar, description),
    user_updated = $9::int,
    updated_at = now()
WHERE id = $10
RETURNING id, image, code, title, entity, staff, frequency, reminder_time, unit, price, description, company, user_created, user_updated, created_at, updated_at
`

type UpdateServiceParams struct {
	Image       sql.NullInt32   `json:"image"`
	Title       sql.NullString  `json:"title"`
	Entity      sql.NullString  `json:"entity"`
	Staff       sql.NullInt32   `json:"staff"`
	Frequency   sql.NullString  `json:"frequency"`
	Unit        sql.NullString  `json:"unit"`
	Price       sql.NullFloat64 `json:"price"`
	Description sql.NullString  `json:"description"`
	UserUpdated sql.NullInt32   `json:"user_updated"`
	ID          int32           `json:"id"`
}

func (q *Queries) UpdateService(ctx context.Context, arg UpdateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, updateService,
		arg.Image,
		arg.Title,
		arg.Entity,
		arg.Staff,
		arg.Frequency,
		arg.Unit,
		arg.Price,
		arg.Description,
		arg.UserUpdated,
		arg.ID,
	)
	var i Service
	err := row.Scan(
		&i.ID,
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
	)
	return i, err
}
