-- name: CreateWarehouse :one
INSERT INTO warehouses (
    name, code, address, companies
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ListWarehouse :many
SELECT * FROM warehouses
WHERE companies = sqlc.narg(company)::int AND (
    name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateTicket :one
INSERT INTO tickets (
    code, type, status, note, qr, export_to, import_from, total_price, warehouse, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: CreateConsignment :one
INSERT INTO consignment (
    code, quantity, inventory, ticket, variant, expired_at, producted_at, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetTicketType :one
SELECT * FROM ticket_type
WHERE id = sqlc.narg('id') OR code = sqlc.narg('code');

-- name: GetTicketStatus :one
SELECT * FROM ticket_status
WHERE id = sqlc.narg('id') OR code = sqlc.narg('code');

-- name: UpdateTicketStatus :one
UPDATE tickets
SET status = $1
WHERE id = $2
RETURNING *;

-- name: GetConsignments :many
SELECT * FROM consignment c
JOIN tickets t ON c.ticket = t.id
JOIN warehouses w ON t.warehouse = w.id
WHERE w.companies = sqlc.arg('company')::int
AND w.id = sqlc.arg('warehouse')::int
AND (
    c.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -c.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: UpdateConsignmentByTicket :many
UPDATE consignment
SET is_available = true
WHERE ticket = $1
RETURNING *;
