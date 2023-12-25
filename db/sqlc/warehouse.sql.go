// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: warehouse.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createConsignment = `-- name: CreateConsignment :one
INSERT INTO consignment (
    code, quantity, inventory, ticket, variant, expired_at, producted_at, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING id, code, quantity, inventory, ticket, expired_at, producted_at, is_available, user_created, user_updated, updated_at, created_at, variant
`

type CreateConsignmentParams struct {
	Code        string        `json:"code"`
	Quantity    int32         `json:"quantity"`
	Inventory   int32         `json:"inventory"`
	Ticket      sql.NullInt32 `json:"ticket"`
	Variant     sql.NullInt32 `json:"variant"`
	ExpiredAt   time.Time     `json:"expired_at"`
	ProductedAt time.Time     `json:"producted_at"`
	UserCreated sql.NullInt32 `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
}

func (q *Queries) CreateConsignment(ctx context.Context, arg CreateConsignmentParams) (Consignment, error) {
	row := q.db.QueryRowContext(ctx, createConsignment,
		arg.Code,
		arg.Quantity,
		arg.Inventory,
		arg.Ticket,
		arg.Variant,
		arg.ExpiredAt,
		arg.ProductedAt,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i Consignment
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Quantity,
		&i.Inventory,
		&i.Ticket,
		&i.ExpiredAt,
		&i.ProductedAt,
		&i.IsAvailable,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Variant,
	)
	return i, err
}

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets (
    code, type, status, note, qr, export_to, import_from, total_price, warehouse, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, code, type, status, note, qr, export_to, import_from, total_price, warehouse, user_created, user_updated, updated_at, created_at
`

type CreateTicketParams struct {
	Code        string         `json:"code"`
	Type        sql.NullInt32  `json:"type"`
	Status      sql.NullInt32  `json:"status"`
	Note        sql.NullString `json:"note"`
	Qr          sql.NullInt32  `json:"qr"`
	ExportTo    sql.NullInt32  `json:"export_to"`
	ImportFrom  sql.NullInt32  `json:"import_from"`
	TotalPrice  float64        `json:"total_price"`
	Warehouse   int32          `json:"warehouse"`
	UserCreated int32          `json:"user_created"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, createTicket,
		arg.Code,
		arg.Type,
		arg.Status,
		arg.Note,
		arg.Qr,
		arg.ExportTo,
		arg.ImportFrom,
		arg.TotalPrice,
		arg.Warehouse,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Type,
		&i.Status,
		&i.Note,
		&i.Qr,
		&i.ExportTo,
		&i.ImportFrom,
		&i.TotalPrice,
		&i.Warehouse,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const createWarehouse = `-- name: CreateWarehouse :one
INSERT INTO warehouses (
    name, code, address, companies
) VALUES (
    $1, $2, $3, $4
) RETURNING id, address, companies, name, code
`

type CreateWarehouseParams struct {
	Name      string        `json:"name"`
	Code      string        `json:"code"`
	Address   sql.NullInt32 `json:"address"`
	Companies sql.NullInt32 `json:"companies"`
}

func (q *Queries) CreateWarehouse(ctx context.Context, arg CreateWarehouseParams) (Warehouse, error) {
	row := q.db.QueryRowContext(ctx, createWarehouse,
		arg.Name,
		arg.Code,
		arg.Address,
		arg.Companies,
	)
	var i Warehouse
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Companies,
		&i.Name,
		&i.Code,
	)
	return i, err
}

const getConsignments = `-- name: GetConsignments :many
SELECT c.id, c.code, quantity, inventory, ticket, expired_at, producted_at, is_available, c.user_created, c.user_updated, c.updated_at, c.created_at, variant, t.id, t.code, type, status, note, qr, export_to, import_from, total_price, warehouse, t.user_created, t.user_updated, t.updated_at, t.created_at, w.id, address, companies, name, w.code FROM consignment c
JOIN tickets t ON c.ticket = t.id
JOIN warehouses w ON t.warehouse = w.id
WHERE w.companies = $1::int
AND w.id = $2::int
AND (
    c.code ILIKE '%' || COALESCE($3::varchar, '') || '%'
)
ORDER BY -c.id
LIMIT COALESCE($5::int, 10)
OFFSET (COALESCE($4::int, 1) - 1) * COALESCE($5::int, 10)
`

type GetConsignmentsParams struct {
	Company   int32          `json:"company"`
	Warehouse int32          `json:"warehouse"`
	Search    sql.NullString `json:"search"`
	Page      sql.NullInt32  `json:"page"`
	Limit     sql.NullInt32  `json:"limit"`
}

type GetConsignmentsRow struct {
	ID            int32          `json:"id"`
	Code          string         `json:"code"`
	Quantity      int32          `json:"quantity"`
	Inventory     int32          `json:"inventory"`
	Ticket        sql.NullInt32  `json:"ticket"`
	ExpiredAt     time.Time      `json:"expired_at"`
	ProductedAt   time.Time      `json:"producted_at"`
	IsAvailable   bool           `json:"is_available"`
	UserCreated   sql.NullInt32  `json:"user_created"`
	UserUpdated   sql.NullInt32  `json:"user_updated"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	CreatedAt     time.Time      `json:"created_at"`
	Variant       sql.NullInt32  `json:"variant"`
	ID_2          int32          `json:"id_2"`
	Code_2        string         `json:"code_2"`
	Type          sql.NullInt32  `json:"type"`
	Status        sql.NullInt32  `json:"status"`
	Note          sql.NullString `json:"note"`
	Qr            sql.NullInt32  `json:"qr"`
	ExportTo      sql.NullInt32  `json:"export_to"`
	ImportFrom    sql.NullInt32  `json:"import_from"`
	TotalPrice    float64        `json:"total_price"`
	Warehouse     int32          `json:"warehouse"`
	UserCreated_2 int32          `json:"user_created_2"`
	UserUpdated_2 sql.NullInt32  `json:"user_updated_2"`
	UpdatedAt_2   sql.NullTime   `json:"updated_at_2"`
	CreatedAt_2   time.Time      `json:"created_at_2"`
	ID_3          int32          `json:"id_3"`
	Address       sql.NullInt32  `json:"address"`
	Companies     sql.NullInt32  `json:"companies"`
	Name          string         `json:"name"`
	Code_3        string         `json:"code_3"`
}

func (q *Queries) GetConsignments(ctx context.Context, arg GetConsignmentsParams) ([]GetConsignmentsRow, error) {
	rows, err := q.db.QueryContext(ctx, getConsignments,
		arg.Company,
		arg.Warehouse,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetConsignmentsRow{}
	for rows.Next() {
		var i GetConsignmentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Quantity,
			&i.Inventory,
			&i.Ticket,
			&i.ExpiredAt,
			&i.ProductedAt,
			&i.IsAvailable,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Variant,
			&i.ID_2,
			&i.Code_2,
			&i.Type,
			&i.Status,
			&i.Note,
			&i.Qr,
			&i.ExportTo,
			&i.ImportFrom,
			&i.TotalPrice,
			&i.Warehouse,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
			&i.ID_3,
			&i.Address,
			&i.Companies,
			&i.Name,
			&i.Code_3,
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

const getTicketStatus = `-- name: GetTicketStatus :one
SELECT id, code, title FROM ticket_status
WHERE id = $1 OR code = $2
`

type GetTicketStatusParams struct {
	ID   sql.NullInt32  `json:"id"`
	Code sql.NullString `json:"code"`
}

func (q *Queries) GetTicketStatus(ctx context.Context, arg GetTicketStatusParams) (TicketStatus, error) {
	row := q.db.QueryRowContext(ctx, getTicketStatus, arg.ID, arg.Code)
	var i TicketStatus
	err := row.Scan(&i.ID, &i.Code, &i.Title)
	return i, err
}

const getTicketType = `-- name: GetTicketType :one
SELECT id, code, title FROM ticket_type
WHERE id = $1 OR code = $2
`

type GetTicketTypeParams struct {
	ID   sql.NullInt32  `json:"id"`
	Code sql.NullString `json:"code"`
}

func (q *Queries) GetTicketType(ctx context.Context, arg GetTicketTypeParams) (TicketType, error) {
	row := q.db.QueryRowContext(ctx, getTicketType, arg.ID, arg.Code)
	var i TicketType
	err := row.Scan(&i.ID, &i.Code, &i.Title)
	return i, err
}

const listWarehouse = `-- name: ListWarehouse :many
SELECT id, address, companies, name, code FROM warehouses
WHERE companies = $1::int AND (
    name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type ListWarehouseParams struct {
	Company sql.NullInt32  `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

func (q *Queries) ListWarehouse(ctx context.Context, arg ListWarehouseParams) ([]Warehouse, error) {
	rows, err := q.db.QueryContext(ctx, listWarehouse,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Warehouse{}
	for rows.Next() {
		var i Warehouse
		if err := rows.Scan(
			&i.ID,
			&i.Address,
			&i.Companies,
			&i.Name,
			&i.Code,
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

const updateConsignmentByTicket = `-- name: UpdateConsignmentByTicket :many
UPDATE consignment
SET is_available = true
WHERE ticket = $1
RETURNING id, code, quantity, inventory, ticket, expired_at, producted_at, is_available, user_created, user_updated, updated_at, created_at, variant
`

func (q *Queries) UpdateConsignmentByTicket(ctx context.Context, ticket sql.NullInt32) ([]Consignment, error) {
	rows, err := q.db.QueryContext(ctx, updateConsignmentByTicket, ticket)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Consignment{}
	for rows.Next() {
		var i Consignment
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Quantity,
			&i.Inventory,
			&i.Ticket,
			&i.ExpiredAt,
			&i.ProductedAt,
			&i.IsAvailable,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Variant,
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

const updateTicketStatus = `-- name: UpdateTicketStatus :one
UPDATE tickets
SET status = $1
WHERE id = $2
RETURNING id, code, type, status, note, qr, export_to, import_from, total_price, warehouse, user_created, user_updated, updated_at, created_at
`

type UpdateTicketStatusParams struct {
	Status sql.NullInt32 `json:"status"`
	ID     int32         `json:"id"`
}

func (q *Queries) UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, updateTicketStatus, arg.Status, arg.ID)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Type,
		&i.Status,
		&i.Note,
		&i.Qr,
		&i.ExportTo,
		&i.ImportFrom,
		&i.TotalPrice,
		&i.Warehouse,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
