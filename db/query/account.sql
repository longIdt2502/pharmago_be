-- name: GetAccount :one
SELECT * FROM accounts a
LEFT JOIN account_company ac ON ac.account = a.id
LEFT JOIN companies c ON ac.company = c.id
WHERE a.id = $1 LIMIT 1;

-- name: GetAccountByUseName :one
SELECT * FROM accounts
WHERE username = $1 LIMIT 1;

-- name: GetAccountByMail :one
SELECT * FROM accounts
WHERE email = $1 LIMIT 1;

-- name: GetAccountByPhone :one
SELECT * FROM accounts
WHERE username = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (username, hashed_password, full_name, email, type, role, gender, licence, dob, address, is_verify)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *;

-- name: CreateAccountCompany :one
INSERT INTO account_company (account, company, company_parent) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET
    is_verify = COALESCE(sqlc.narg(is_verify), is_verify),
    hashed_password = COALESCE(sqlc.narg(password)::varchar, hashed_password),
    full_name = COALESCE(sqlc.narg(full_name)::varchar, full_name),
    email = COALESCE(sqlc.narg(email)::varchar, email),
    type = COALESCE(sqlc.narg(type)::int, type),
    role = COALESCE(sqlc.narg(role)::int, role),
    gender = COALESCE(sqlc.narg(gender)::gender, gender),
    licence = COALESCE(sqlc.narg(licence)::varchar, licence),
    dob = COALESCE(sqlc.narg(dob)::timestamp, dob),
    address = COALESCE(sqlc.narg(address)::int, address)
WHERE
    id = sqlc.narg(id)
    OR username = sqlc.narg(username)
RETURNING *;

-- name: ResetPassword :one
UPDATE accounts
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password)
WHERE
    email = sqlc.narg(email)
RETURNING *;

-- name: ListAccount :many
SELECT * FROM accounts a
LEFT JOIN account_company ac ON ac.account = a.id
LEFT JOIN companies c ON c.id = ac.company
LEFT JOIN account_type at ON at.id = a.type 
LEFT JOIN roles r ON r.id = a.role 
WHERE (ac.company = sqlc.arg(company)::int OR ac.company_parent = sqlc.narg(company))
AND (
    a.full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    a.username ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
AND (
    sqlc.narg(is_verify)::bool IS NULL OR a.is_verify = sqlc.narg(is_verify)::bool
)
AND (
    sqlc.narg(type)::int IS NULL OR a.type = sqlc.narg(type)::int
)
AND (
    sqlc.narg(role)::int IS NULL OR a.role = sqlc.narg(role)::int
)
ORDER BY -a.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: AssignEmployee :one
UPDATE account_company 
SET company = sqlc.narg(company)
WHERE account = sqlc.arg(account)
RETURNING *;

-- name: CountAccountByStatus :many
SELECT a.is_verify ,COUNT(a.id) as "count" FROM accounts a
LEFT JOIN account_company ac ON ac.account = a.id
WHERE (ac.company = sqlc.arg(company)::int OR ac.company_parent = sqlc.arg(company)::int)
GROUP BY a.is_verify;

-- name: DeleteEmployee :one
DELETE FROM accounts
WHERE id = $1 RETURNING *;
