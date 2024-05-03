-- name: GetRevenueCompany :one
SELECT
    COALESCE(SUM(total_price), 0) :: float AS total_sum
FROM
    orders
WHERE
    status = 'COMPLETE'
    AND type = 'SELL'
    AND company = $1;

-- name: GetVariantBestSale :many
WITH variant_total AS (
    SELECT
        v.id AS variant_id,
        COALESCE(SUM(oi.total_price), 0) :: float AS total_revenue
    FROM
        variants v
        LEFT JOIN order_items oi ON v.id = oi.variant
        LEFT JOIN orders o ON o.id = oi.order
    WHERE
        o.status = 'COMPLETE'
    GROUP BY
        v.id
)
SELECT
    v.*,
    vt.total_revenue AS revenue,
    m.media_url AS imageUrl
FROM
    variant_total vt
    JOIN variants v ON v.id = vt.variant_id
    JOIN products p ON v.product = p.id
    JOIN variant_media vm ON vm.variant = v.id
    JOIN medias m ON vm.media = m.id
WHERE
    p.company = $1
ORDER BY
    vt.total_revenue DESC
LIMIT
    3;

-- name: GetReportRevenue :many
WITH date_series AS (
    SELECT
        generate_series(
            date_trunc(
                CASE
                    WHEN COALESCE(sqlc.arg('filter') :: varchar, 'day') = 'day' THEN 'month'
                    WHEN COALESCE(sqlc.arg('filter') :: varchar, 'day') = 'month' THEN 'year'
                    ELSE 'month'
                END,
                current_date
            ),
            current_date,
            CASE
                WHEN COALESCE(sqlc.arg('filter') :: varchar, 'day') = 'day' THEN interval '1 day'
                WHEN COALESCE(sqlc.arg('filter') :: varchar, 'day') = 'month' THEN interval '1 month'
                ELSE interval '1 day'
            END
        ) :: timestamp AS date
),
order_series As (
    SELECT
        *
    FROM
        orders
    WHERE
        company = sqlc.arg(company) :: int
        AND (
            sqlc.narg(status) :: varchar IS NULL
            OR status = sqlc.narg(status) :: varchar
        )
)
SELECT
    ds.*,
    COALESCE(SUM(o1.total_price), 0)::float AS current_revenue,
    COALESCE(SUM(o2.total_price), 0)::float AS last_revenue
FROM
    date_series ds
    LEFT JOIN order_series o1 ON date_trunc(
        COALESCE(sqlc.arg('filter') :: varchar, 'day'),
        o1.created_at
    ) = ds.date
    LEFT JOIN order_series o2 ON date_trunc(
        COALESCE(sqlc.arg('filter') :: varchar, 'day'),
        o2.created_at
    ) = date_trunc(
        COALESCE(sqlc.arg('filter') :: varchar, 'day'),
        ds.date - CASE
            WHEN COALESCE(sqlc.arg('filter') :: varchar, 'day') = 'day' THEN interval '1 month'
            WHEN COALESCE(sqlc.arg('filter') :: varchar, 'day') = 'month' THEN interval '1 year'
            ELSE interval '1 day'
        END
    )
GROUP BY
    ds.date;

-- name: TotalRevenue :one
SELECT COALESCE(SUM(total_price), 0)::float AS value FROM orders
WHERE company = sqlc.arg(company) :: int AND date_trunc(
    COALESCE(sqlc.arg('filter') :: varchar, 'month'),
    created_at
) = date_trunc(COALESCE(sqlc.arg('filter') :: varchar, 'month'), current_date - 
    CASE 
        WHEN COALESCE(sqlc.arg('filter') :: varchar, 'month') = 'month' THEN MAKE_INTERVAL(
            months => COALESCE(sqlc.arg('interval') :: int, 0)
        )::interval
        WHEN COALESCE(sqlc.arg('filter') :: varchar, 'month') = 'year' THEN MAKE_INTERVAL(
            years => COALESCE(sqlc.arg('interval') :: int, 0)
        )::interval
        ELSE interval '0 month'    
    END
    );