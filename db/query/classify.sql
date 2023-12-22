-- name: CreateClassify :one
INSERT INTO classify (
    code, name
) VALUES (
    $1, $2
) RETURNING *;