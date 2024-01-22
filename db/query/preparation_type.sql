-- name: CreatePreparationType :one
INSERT INTO preparation_type (
    code, name, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListPreparationType :many
WITH preparation_type_quantity AS (
    SELECT ps.id AS preparation_type_id,
           COALESCE(COUNT(p.id), 0)::int AS total_quantity
    FROM preparation_type ps
             LEFT JOIN products p ON p.tieu_chuan_sx = ps.code
    GROUP BY ps.id
)
SELECT ps.*, a.*, psq.total_quantity AS quantity FROM preparation_type_quantity psq
JOIN preparation_type ps ON psq.preparation_type_id = ps.id
LEFT JOIN accounts a ON a.id = ps.user_created
WHERE (
    ps.company IS NULL OR
    ps.company = sqlc.arg('company')::int
)
AND (
    ps.name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    ps.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -ps.id
    LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: DetailPreparationType :one
SELECT ps.*, a.*, au.full_name AS user_updated_name FROM preparation_type ps
LEFT JOIN accounts a ON a.id = ps.user_created
LEFT JOIN accounts au ON au.id = ps.user_updated
WHERE ps.id = $1;

-- name: UpdatePreparationType :one
UPDATE preparation_type
SET
    name = sqlc.arg(name),
    code = COALESCE(sqlc.narg(code), code),
    description = COALESCE(sqlc.narg(description), description),
    user_updated = sqlc.arg(user_updated),
    updated_at = now()
WHERE id = sqlc.arg(id)
    RETURNING *;

-- name: DeletePreparationType :one
DELETE FROM preparation_type
WHERE id = $1 RETURNING *;