-- name: CreateProductPriceList :one
INSERT INTO price_list (
    variant_code, variant_name, unit, price_import, price_sell, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetPriceLists :many
SELECT *, u.name AS unit_name, p.company AS company, m.media_url AS variant_media,
       a.full_name AS user_created_name FROM price_list pl
JOIN units u ON pl.unit = u.id
JOIN variants v ON pl.variant_code = v.code
JOIN accounts a ON a.id = pl.user_created
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
JOIN products p ON p.id = v.product
WHERE p.company = sqlc.narg(company)::int
AND (
    variant_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    variant_code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
AND (
    (sqlc.narg('min_price_import')::float IS NULL AND sqlc.narg('max_price_import')::float IS NULL)
    OR (pl.price_import BETWEEN sqlc.narg('min_price_import')::float AND sqlc.narg('max_price_import')::float)
    OR (sqlc.narg('min_price_sell')::float IS NULL AND sqlc.narg('max_price_sell')::float IS NULL)
    OR (pl.price_sell BETWEEN sqlc.narg('min_price_sell')::float AND sqlc.narg('max_price_sell')::float)
)
ORDER BY -pl.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: DetailPriceList :one
SELECT *, u.name AS unit_name, m.media_url AS variant_media,
       a.full_name AS user_created_name FROM price_list pl
JOIN units u ON pl.unit = u.id
JOIN variants v ON pl.variant_code = v.code
JOIN accounts a ON a.id = pl.user_created
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
WHERE pl.id = sqlc.arg('id')::int;

-- name: UpdatePriceList :one
UPDATE price_list
SET
    price_import = sqlc.arg('price_import')::float,
    price_sell = sqlc.arg('price_sell'):: float
WHERE id = sqlc.arg('id')
RETURNING *;

























