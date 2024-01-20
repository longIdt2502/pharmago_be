-- name: GetRevenueCompany :one
SELECT COALESCE(SUM(total_price), 0)::float AS total_sum
FROM orders
WHERE status = 'COMPLETE'
AND type = 'SELL'
AND company = $1;

-- name: GetVariantBestSale :many
WITH variant_total AS (
    SELECT v.id AS variant_id,
           COALESCE(SUM(oi.total_price), 0)::float AS total_revenue
    FROM variants v
    LEFT JOIN order_items oi ON v.id = oi.variant
    LEFT JOIN orders o ON o.id = oi.order
    WHERE o.status = 'COMPLETE'
    GROUP BY v.id
)
SELECT v.*, vt.total_revenue AS revenue, m.media_url AS imageUrl FROM variant_total vt
JOIN variants v ON v.id = vt.variant_id
JOIN products p ON v.product = p.id
JOIN variant_media vm ON vm.variant = v.id
JOIN medias m ON vm.media = m.id
WHERE p.company = $1
ORDER BY vt.total_revenue DESC
LIMIT 3;