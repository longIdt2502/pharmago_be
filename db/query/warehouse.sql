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

-- name: GetListTicket :many
SELECT *, w.name AS w_name, a.full_name AS a_full_name, m.media_url AS qr_url,
    tt.id AS tt_id, tt.code AS tt_code, tt.title AS tt_title,
    ts.id AS ts_id, ts.code AS ts_code, ts.title AS ts_title,
    COALESCE(SUM(c.quantity), 0)::int AS total_products
FROM tickets t
JOIN warehouses w ON t.warehouse = w.id
JOIN accounts a ON t.user_created = a.id
JOIN medias m ON t.qr = m.id
JOIN ticket_type tt ON t.type = tt.id
JOIN ticket_status ts ON t.status = ts.id
LEFT JOIN consignment c ON t.id = c.ticket
WHERE w.companies = sqlc.arg('company')
AND (
    t.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
GROUP BY
    t.id, t.code, t.type, t.status, t.note, t.qr, t.total_price, t.warehouse, t.user_created, t.created_at,
    w.id, a.id, m.id, tt.id, ts.id, c.ticket, c.id,
    w.name, a.full_name, m.media_url, tt.id, tt.code, tt.title, ts.id, ts.code, ts.title
ORDER BY -t.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: GetDetailTicket :one
SELECT * , m.media_url AS qr_url, a_uc.full_name AS user_created_name, a_uu.full_name AS user_updated_name,
       w.id AS w_id, w.name AS w_name, w.code AS w_code, w.address AS w_address,
        c.id AS c_id, c.full_name AS c_name, c.code AS c_code, c.address AS c_address, c.email AS c_email, c.phone AS c_phone, c.company AS c_company,
        s.id AS s_id, s.code AS s_code, s.name AS s_name, s.deputy_name AS s_deputy, s.phone AS s_phone, s.email AS s_email, s.address AS s_address, s.company AS s_company,
       tt.id AS tt_id, tt.code AS tt_code, tt.title AS tt_title,
       ts.id AS ts_id, ts.code AS ts_code, ts.title AS ts_title
FROM tickets t
LEFT JOIN suplier s ON t.import_from = s.address AND t.type = 1
LEFT JOIN customers c ON t.export_to = c.address AND t.type = 2
JOIN warehouses w ON w.id = t.warehouse
JOIN ticket_type tt ON t.type = tt.id
JOIN ticket_status ts ON t.status = ts.id
JOIN medias m ON t.qr = m.id
JOIN accounts a_uc ON a_uc.id = t.user_created
JOIN accounts a_uu ON a_uu.id = t.user_updated
WHERE t.id = $1;

-- name: GetItemsTicket :many
SELECT * FROM consignment
WHERE ticket = $1;

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
AND (sqlc.narg('available')::bool IS NULL OR c.is_available = sqlc.narg('available')::bool)
AND (
    sqlc.narg('warehouse')::int IS NULL OR
    w.id = sqlc.narg('warehouse')::int
)
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
WHERE c.variant = $1
AND is_available = true
AND inventory > $2
ORDER BY ABS(EXTRACT(EPOCH FROM (expired_at - NOW()))) ASC
LIMIT 1;

-- name: UpdateConsignmentByTicket :many
UPDATE consignment
SET is_available = true
WHERE ticket = $1
RETURNING *;

-- name: GetListSupplier :many
SELECT * FROM suplier
WHERE company = sqlc.arg('company')::int
AND (
    name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);
