-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = sqlc.arg('id')
LIMIT 1;

-- name: ListCustomer :many
SELECT * FROM customers
WHERE company = sqlc.arg(company)::int
AND (
    full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateCustomer :one
INSERT INTO customers (
    full_name, code, company, address, email, phone ,license, birthday, user_updated, user_created
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: DetailCustomer :one
SELECT * FROM customers
WHERE id = $1
LIMIT 1;

-- name: UpdateCustomer :one
UPDATE customers
SET
    full_name = COALESCE(sqlc.narg(full_name)::varchar, full_name),
    code = COALESCE(sqlc.narg(code)::varchar, code),
    email = COALESCE(sqlc.narg(email)::varchar, email),
    phone = COALESCE(sqlc.narg(phone)::varchar, phone),
    license = COALESCE(sqlc.narg(license)::varchar, license),
    birthday = COALESCE(sqlc.narg(birthday)::timestamp, birthday),
    user_updated = COALESCE(sqlc.narg(user_updated)::int, user_updated)
WHERE id = sqlc.arg(id)
RETURNING *;