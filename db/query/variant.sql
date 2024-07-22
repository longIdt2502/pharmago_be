-- name: GetVariantById :one
SELECT * FROM variants
WHERE id = $1
LIMIT 1;

-- name: GetVariants :many
WITH revenue AS (
    SELECT variant, SUM("value") AS total_buy FROM order_items
    GROUP BY variant
)
SELECT *, m.media_url AS media,
       u.id AS unit_id, u.name AS unit_name, u.sell_price AS unit_sell_price, u.weight AS unit_weight, u.weight_unit AS unit_weight_unit,
       pl.price_import AS pl_price_import, pl.price_sell AS pl_price_sell
FROM variants v
JOIN products p ON v.product = p.id
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
JOIN units u ON u.id = p.unit
LEFT JOIN price_list pl ON pl.variant_code = v.code
LEFT JOIN revenue r ON r.variant = v.id
WHERE (p.company = sqlc.arg(company)::int OR v.product = sqlc.arg(product)::int)
AND (
    v.name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    v.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    v.barcode ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
) AND (
    sqlc.narg('id')::int IS NULL OR v.id = sqlc.narg('id')::int
) AND (
    sqlc.narg('filter')::varchar IS NULL OR (r.total_buy IS NOT NULL AND sqlc.narg('filter') = 'POPULAR')
)
ORDER BY 
    CASE
        WHEN sqlc.narg('filter') = 'POPULAR' THEN r.total_buy
    END, -v.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: GetVariantsByCode :one
SELECT *, m.media_url AS media,
       u.id AS unit_id, u.name AS unit_name, u.sell_price AS unit_sell_price, u.weight AS unit_weight, u.weight_unit AS unit_weight_unit,
       pl.price_import AS pl_price_import, pl.price_sell AS pl_price_sell
FROM variants v
         JOIN products p ON v.product = p.id
         LEFT JOIN variant_media vm ON vm.variant = v.id
         LEFT JOIN medias m ON m.id = vm.media
         JOIN units u ON u.id = p.unit
         JOIN price_list pl ON pl.variant_code = v.code
WHERE p.company = sqlc.arg(company)::int
AND v.barcode = sqlc.arg('code')::varchar
LIMIT 1;

-- name: GetInventoryVariant :one
SELECT COALESCE(SUM(inventory), 0)::int AS total_inventory
FROM consignment
WHERE variant = $1 AND is_available = true;

-- name: GetVariantsByProduct :many
SELECT * FROM variants
WHERE product = $1;

-- name: VariantsCustomerBuy :many
SELECT v.*, u.*, SUM(value) AS quantity_buy FROM order_items oi
JOIN variants v ON v.id = oi.variant
JOIN orders o ON o.id = oi.order
JOIN products p ON p.id = v.product
JOIN units u ON u.id = p.unit
WHERE o.customer = sqlc.arg(customer)::int
GROUP BY oi.variant, v.id, u.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);
