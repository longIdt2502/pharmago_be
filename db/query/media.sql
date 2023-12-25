-- name: GetMedia :one
SELECT * FROM medias
WHERE id = $1 LIMIT 1;

-- name: CreateMedia :one
INSERT INTO medias (
    media_url
) VALUES ($1) RETURNING *;

-- name: GetMediaVariant :one
SELECT *, m.media_url FROM variant_media vm
JOIN medias m ON m.id = vm.media
WHERE variant = $1 LIMIT 1;