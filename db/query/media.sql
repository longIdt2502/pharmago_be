-- name: GetMedia :one
SELECT * FROM medias
WHERE id = $1 LIMIT 1;