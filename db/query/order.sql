-- name: CreateOrder :one
INSERT INTO orders (
    code, total_price, description, vat, discount, service_price,
    must_paid, customer, address, status, type, ticket, qr,
    company, payment, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (
    "order", variant, value, total_price, consignment, consignment_log
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: ListOrder :many
SELECT *, c.full_name AS c_full_name, os.title AS os_title, os.id AS os_id, a.full_name AS a_full_name FROM orders o
JOIN customers c ON o.customer = c.id
JOIN tickets t ON o.ticket = t.id
JOIN order_status os ON os.code = o.status
JOIN accounts a ON a.id = o.user_created
WHERE o.company = sqlc.narg(company)::int
AND (
    sqlc.narg('warehouse')::int IS NULL OR t.warehouse = sqlc.narg('warehouse')::int
)
AND (
    o.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    c.full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -o.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);