-- name: GetMedia :one
SELECT * FROM medias
WHERE id = $1 LIMIT 1;

-- name: CreateMedia :one
INSERT INTO medias (
    media_url
) VALUES ($1) RETURNING *;