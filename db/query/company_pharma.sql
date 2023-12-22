-- name: CreateCompanyPharma :one
INSERT INTO company_pharma (
    name, code, country, address, company_pharma_type
) VALUES (
    sqlc.arg('name')::varchar, sqlc.narg('code')::varchar, sqlc.narg('country')::varchar, sqlc.narg('address')::varchar,
    sqlc.arg('company_pharma_type')::varchar
) RETURNING *;