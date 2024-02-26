-- name: GetListCategory :many
SELECT * FROM product_categories
WHERE company = sqlc.narg(company)::int
AND (
    name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
    LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateCategory :one
INSERT INTO product_categories (
    code, name, description, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateCategory :one
UPDATE product_categories
SET
    code = sqlc.arg(code),
    name = sqlc.arg(name),
    description = sqlc.arg(description),
    user_updated = sqlc.arg(user_updated)
WHERE id = sqlc.arg(id)
    RETURNING *;

-- name: DetailCategory :one
SELECT *, ac.full_name AS created_name, au.full_name AS updated_name FROM product_categories pb
JOIN accounts ac ON ac.id = pb.user_created
LEFT JOIN accounts au ON au.id = pb.user_updated
WHERE pb.id = $1;

-- name: DeleteCategory :one
DELETE FROM product_categories
WHERE id = $1 RETURNING *;