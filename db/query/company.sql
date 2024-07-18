-- name: CreateCompany :one
INSERT INTO companies (
    name, code, type, tax_code, phone, description, address, owner
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetCompanies :many
SELECT * FROM companies
WHERE owner = sqlc.narg('owner')::int AND
    (name ILIKE COALESCE(sqlc.narg('search')::varchar, '%') OR
    phone ILIKE COALESCE(sqlc.narg('search')::varchar, '%'))
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: GetCompanyById :one
SELECT * FROM companies
WHERE id = $1
LIMIT 1;

-- name: GetCompanyByPhone :one
SELECT * FROM companies
WHERE phone = $1
LIMIT 1;
