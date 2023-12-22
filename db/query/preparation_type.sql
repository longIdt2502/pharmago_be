-- name: CreatePreparationType :one
INSERT INTO preparation_type (
    code, name
) VALUES (
    $1, $2
) RETURNING *;