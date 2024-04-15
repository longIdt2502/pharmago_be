-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

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
INSERT INTO accounts (username, hashed_password, full_name, email, type, role, gender, licence, dob, address)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

-- name: CreateAccountCompany :one
INSERT INTO account_company (account, company) VALUES ($1, $2) RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET
    is_verify = COALESCE(sqlc.narg(is_verify), is_verify)
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
WHERE ac.company = sqlc.arg(company)::int
AND (
    a.full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    a.username ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
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

