-- name: GetListService :many
WITH quantity_use AS (
    SELECT "service", COUNT("service") as quantity_use FROM service_order_item
    GROUP BY "service"
)
SELECT * FROM services s
LEFT JOIN quantity_use qu ON s.id = qu.service
WHERE s.company = sqlc.narg(company)::int
AND (
    s.title ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    s.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -s.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: GetServicesByCustomer :many
SELECT s.*, SUM(quantity) as quantity_use FROM service_order_item soi
JOIN orders o ON o.id = soi.order
JOIN services s ON s.id = soi.service
WHERE o.customer = sqlc.arg(customer)::int
GROUP BY soi.service, s.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: ServicesUsedByCustomer :many
SELECT s.*, COUNT(s.id) AS number_of_uses FROM service_order_item soi
JOIN orders o ON o.id = soi.order
JOIN services s ON s.id = soi.service
WHERE o.customer = sqlc.arg('customer')::int
GROUP BY s.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateService :one
INSERT INTO services (
    code, image, title, entity, staff, frequency, unit, price, description, company, user_created, user_updated, reminder_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: UpdateService :one
UPDATE services
SET
    image = COALESCE(sqlc.narg(image)::int, image),
    title = COALESCE(sqlc.narg(title)::varchar, title),
    entity = COALESCE(sqlc.narg(entity)::varchar, entity),
    staff = COALESCE(sqlc.narg(staff)::int, staff),
    frequency = COALESCE(sqlc.narg(frequency)::varchar, frequency),
    unit = COALESCE(sqlc.narg(unit)::varchar, unit),
    price = COALESCE(sqlc.narg(price)::float, price),
    description = COALESCE(sqlc.narg(description)::varchar, description),
    user_updated = sqlc.narg(user_updated)::int,
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DetailService :one
SELECT * FROM services
WHERE id = $1;

-- name: DeleteService :one
DELETE FROM services
WHERE id = $1
RETURNING *;

-- name: CreateServiceVariant :one
INSERT INTO service_variant (
    service, variant
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListServiceVariant :many
SELECT * FROM service_variant sv
LEFT JOIN variants v ON v.id = sv.variant
WHERE sv.service = sqlc.arg(id)::int;

-- name: DeleteServiceVariant :one
DELETE FROM service_variant
WHERE id = $1
RETURNING *;
