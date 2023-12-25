-- name: GetVariantById :one
SELECT * FROM variants
WHERE id = $1
LIMIT 1;