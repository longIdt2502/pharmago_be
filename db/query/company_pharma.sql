-- name: CreateCompanyPharma :one
INSERT INTO company_pharma (
    name, code, country, address, company_pharma_type
) VALUES (
    sqlc.arg('name')::varchar, sqlc.narg('code')::varchar, sqlc.narg('country')::varchar, sqlc.narg('address')::varchar,
    sqlc.arg('company_pharma_type')::varchar
) RETURNING *;

-- name: UpdateCompanyPharma :one
UPDATE company_pharma
SET
    name = COALESCE(sqlc.narg('name')::varchar, name),
    code = COALESCE(sqlc.narg('code')::varchar, code),
    country = COALESCE(sqlc.narg('country')::varchar, country),
    address = COALESCE(sqlc.narg('address')::varchar, address)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetCompanyPharmaByName :one
SELECT * FROM company_pharma
WHERE name = $1
LIMIT 1;

-- name: GetCompanyPharmaDetail :one
SELECT * FROM company_pharma
WHERE id = $1
LIMIT 1;

-- name: DeleteCompanyPharma :one
DELETE FROM company_pharma
WHERE id = $1 RETURNING *;