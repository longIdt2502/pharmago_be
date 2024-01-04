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

-- name: CreateSupplier :one
INSERT INTO suplier (
    code, name, deputy_name, phone, email, address, company
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: DetailSupplier :one
SELECT * FROM suplier
WHERE id = $1;

-- name: UpdateSupplier :one
UPDATE suplier
SET
    name = COALESCE(sqlc.narg(name), name),
    deputy_name = COALESCE(sqlc.narg(deputy_name), deputy_name),
    phone = COALESCE(sqlc.narg(phone), phone),
    email = COALESCE(sqlc.narg(email), email)
WHERE id = sqlc.arg(id)
RETURNING *;
