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