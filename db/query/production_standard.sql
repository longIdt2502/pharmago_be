-- name: CreateProductionStandard :one
INSERT INTO production_standard (
    code, name, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListProductionStandard :many
WITH production_standard_quantity AS (
    SELECT ps.id AS production_standard_id,
           COALESCE(COUNT(p.id), 0)::int AS total_quantity
    FROM production_standard ps
    LEFT JOIN products p ON p.tieu_chuan_sx = ps.code
    GROUP BY ps.id
)
SELECT ps.*, a.*, psq.total_quantity AS quantity FROM production_standard_quantity psq
JOIN production_standard ps ON psq.production_standard_id = ps.id
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

-- name: DetailProductionStandard :one
SELECT ps.*, a.*, au.full_name AS user_updated_name FROM production_standard ps
LEFT JOIN accounts a ON a.id = ps.user_created
LEFT JOIN accounts au ON au.id = ps.user_updated
WHERE ps.id = $1;

-- name: UpdateProductionStandard :one
UPDATE production_standard
SET
    name = sqlc.arg(name),
    code = COALESCE(sqlc.narg(code), code),
    description = COALESCE(sqlc.narg(description), description),
    user_updated = sqlc.arg(user_updated),
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteProductionStandard :one
DELETE FROM production_standard
WHERE id = $1 RETURNING *;