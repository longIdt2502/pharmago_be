-- name: CreateWarehouse :one
INSERT INTO warehouses (
    name, code, address, companies
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetWarehouse :one
SELECT * FROM warehouses
WHERE id = $1
LIMIT 1;

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

-- name: CreateConsignmentLog :one
INSERT INTO consignment_log (
    consignment, inventory, amount_change, user_created
) VALUES (
     $1, $2, $3, $4
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

-- name: GetConsignment :one
SELECT * FROM consignment
WHERE id = $1 AND variant = $2
LIMIT 1;

-- name: UpdateConsignment :one
UPDATE consignment
SET inventory = inventory + sqlc.arg('amount')::int
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: SuggestConsignmentForVariant :one
SELECT * FROM consignment c
WHERE c.variant = $1 AND inventory = (SELECT MIN(inventory) FROM consignment) AND is_available = true
ORDER BY ABS(EXTRACT(EPOCH FROM (expired_at - NOW()))) ASC
LIMIT 1;

-- name: UpdateConsignmentByTicket :many
UPDATE consignment
SET is_available = true
WHERE ticket = $1
RETURNING *;
