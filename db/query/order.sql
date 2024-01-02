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
    sqlc.narg('status')::varchar IS NULL OR o.status = sqlc.narg('status')::varchar
)
AND (
    sqlc.narg('warehouse')::int IS NULL OR t.warehouse = sqlc.narg('warehouse')::int
)
AND (
    o.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    c.full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
AND  ((
    sqlc.narg('created_start')::timestamp  IS NULL AND sqlc.narg('created_end')::timestamp  IS NULL
) OR (
    (sqlc.narg('created_start')::timestamp  IS NULL OR o.created_at >= sqlc.narg('created_start')::timestamp) AND
    (sqlc.narg('created_end')::timestamp  IS NULL OR o.created_at <= sqlc.narg('created_end')::timestamp)
))
AND ((
    sqlc.narg('updated_start')::timestamp  IS NULL AND sqlc.narg('updated_end')::timestamp  IS NULL
) OR (
    (o.updated_at >= sqlc.narg('updated_start')::timestamp OR sqlc.narg('updated_start')::timestamp  IS NULL) AND
    (o.updated_at <= sqlc.narg('updated_end')::timestamp OR sqlc.narg('updated_end')::timestamp  IS NULL)
))
ORDER BY
    CASE WHEN sqlc.narg('order_by')::varchar = 'created_at' THEN o.created_at END DESC,
    CASE WHEN sqlc.narg('order_by')::varchar = '-created_at' THEN o.created_at END ASC,
    CASE WHEN sqlc.narg('order_by')::varchar = 'updated_at' THEN o.updated_at END DESC,
    CASE WHEN sqlc.narg('order_by')::varchar = '-updated_at' THEN o.updated_at END ASC,
    CASE WHEN sqlc.narg('order_by')::varchar IS NULL THEN o.id END DESC
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: DetailOrder :one
SELECT *, m.media_url AS qr_url, ot.id AS ot_id, ot.code AS ot_code, ot.title AS ot_title,
       os.id AS os_id, os.code AS os_code, os.title AS os_title,
       a.full_name AS a_full_name FROM orders o
JOIN medias m ON o.qr = m.id
JOIN order_type ot ON o.type = ot.code
JOIN order_status os ON o.status = os.code
JOIN accounts a ON o.user_created = a.id
WHERE o.id = $1;

-- name: ListOrderItem :many
SELECT * FROM order_items oi
JOIN variants v ON v.id = oi.variant
JOIN consignment c ON c.id = oi.consignment
JOIN variant_media vm ON vm.variant = v.id
JOIN medias m ON vm.media = m.id
WHERE oi.order = $1;

-- name: UpdateStatusOrder :one
UPDATE orders
SET status = sqlc.arg('status')::varchar
WHERE id = sqlc.arg('id')::int
RETURNING *;
