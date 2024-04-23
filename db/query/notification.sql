-- name: ListNotification :many
SELECT * FROM notification 
WHERE company = sqlc.arg('company')::int
AND (
    title ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    content ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateNotification :one
INSERT INTO notification (
    type, topic, title, content, is_read, data, company
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: CountNotification :many
SELECT COUNT(*), is_read FROM notification
WHERE company = $1
GROUP BY is_read;

-- name: DetailNotification :one
SELECT * FROM notification
WHERE id = $1;

-- name: UpdateNotification :one
UPDATE notification 
SET
    is_read = COALESCE(sqlc.narg(is_read), is_read)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SeenAllNotification :many
UPDATE notification
SET
    is_read = true
WHERE 
    company = sqlc.arg(company) AND
    is_read = false
RETURNING *;

