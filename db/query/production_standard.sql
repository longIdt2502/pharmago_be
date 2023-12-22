-- name: CreateProductionStandard :one
INSERT INTO production_standard (
    code, name
) VALUES (
    $1, $2
) RETURNING *;