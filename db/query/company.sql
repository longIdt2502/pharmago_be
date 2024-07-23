-- name: CreateCompany :one
INSERT INTO companies (
    name, code, type, tax_code, phone, description, address, owner, time_open, time_close, user_created, user_updated, parent, manager
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING *;

-- name: GetCompanies :many
SELECT * FROM companies
WHERE owner = sqlc.narg('owner')::int AND
    (name ILIKE COALESCE(sqlc.narg('search')::varchar, '%') OR
    phone ILIKE COALESCE(sqlc.narg('search')::varchar, '%'))
AND (sqlc.narg(parent)::int IS NULL OR parent = sqlc.narg(parent)::int)
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

-- name: CountEmployee :one
SELECT COUNT(id) AS total FROM account_company
WHERE company = $1
GROUP BY id;

-- name: UpdateCompany :one
UPDATE companies
SET 
    name = COALESCE(sqlc.narg(name)::varchar, name),
    type = COALESCE(sqlc.narg(type)::varchar, type),
    manager = COALESCE(sqlc.narg(manager)::int, manager),
    is_active = COALESCE(sqlc.narg(is_active)::bool, is_active),
    time_open = COALESCE(sqlc.narg(time_open)::time, time_open),
    time_close = COALESCE(sqlc.narg(time_close)::time, time_close),
    user_updated = COALESCE(sqlc.narg(user_updated)::int, user_updated)
WHERE id = sqlc.arg(id)::int
RETURNING *;

-- name: DetailCompany :one
WITH employee AS (
    SELECT *, COUNT(id) AS total_employee FROM account_company
    WHERE company = $1
    GROUP BY id
)
SELECT * FROM companies c
JOIN company_type ct ON ct.code = c.type
LEFT JOIN accounts am ON am.id = c.manager
LEFT JOIN accounts ac ON ac.id = c.user_created
LEFT JOIN accounts au ON au.id = c.user_updated
LEFT JOIN employee e ON e.company = c.id
WHERE c.id = $1
LIMIT 1;