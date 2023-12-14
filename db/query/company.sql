-- name: CreateCompany :one
INSERT INTO companies (
    name, code, tax_code, phone, description, address, owner
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;