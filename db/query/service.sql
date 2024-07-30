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
    code, image, title, entity, staff, frequency, unit, price, description, 
    company, user_created, user_updated, reminder_time,
    brand, action_time, chi_dinh, chong_chi_dinh, cong_dung, caution, hinh_thuc,
    tac_dung_phu, number_register, number_decision, cong_ty_dk, message
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, 
    $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
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
    brand = COALESCE(sqlc.narg(brand)::varchar, brand),
    action_time = COALESCE(sqlc.narg(action_time)::varchar, action_time),
    chi_dinh = COALESCE(sqlc.narg(chi_dinh)::varchar, chi_dinh),
    chong_chi_dinh = COALESCE(sqlc.narg(chong_chi_dinh)::varchar, chong_chi_dinh),
    cong_dung = COALESCE(sqlc.narg(cong_dung)::varchar, cong_dung),
    caution = COALESCE(sqlc.narg(caution)::varchar, caution),
    hinh_thuc = COALESCE(sqlc.narg(hinh_thuc)::varchar, hinh_thuc),
    tac_dung_phu = COALESCE(sqlc.narg(tac_dung_phu)::varchar, tac_dung_phu),
    number_register = COALESCE(sqlc.narg(number_register)::varchar, number_register),
    number_decision = COALESCE(sqlc.narg(number_decision)::varchar, number_decision),
    cong_ty_dk = COALESCE(sqlc.narg(cong_ty_dk)::varchar, cong_ty_dk),
    message = COALESCE(sqlc.narg(message)::varchar, message),
    user_updated = sqlc.narg(user_updated)::int,
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DetailService :one
SELECT * FROM services s
LEFT JOIN accounts a ON a.id = s.staff
LEFT JOIN product_brand pb ON pb.id = s.brand
WHERE s.id = $1;

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
