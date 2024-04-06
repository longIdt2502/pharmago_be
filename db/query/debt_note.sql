-- name: GetListDebtNote :many
-- WITH total_repayment AS (
--     SELECT debt, COALESCE(SUM(money), 0)::float AS total_money
--     FROM debt_repayment
--     GROUP BY debt
-- )
SELECT dn.*, a.full_name AS a_name, c.full_name AS c_name, s.name AS s_name FROM debt_note dn
LEFT JOIN accounts a ON a.id = dn.user_created
LEFT JOIN customers c ON c.code = dn.entity
LEFT JOIN suplier s ON s.code = dn.entity
WHERE dn.company = sqlc.narg(company)::int
AND (
    sqlc.narg('status')::varchar IS NULL OR dn.status = sqlc.narg('status')::varchar
)
AND (
    sqlc.narg('type')::varchar IS NULL OR dn.type = sqlc.narg('type')::varchar
)
AND (
    dn.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    dn.title ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
-- AND  ((
--     sqlc.narg('created_start')::timestamp IS NULL AND sqlc.narg('created_end')::timestamp  IS NULL
-- ) OR (
--     (sqlc.narg('created_start')::timestamp IS NULL OR o.created_at >= sqlc.narg('created_start')::timestamp) AND
--     (sqlc.narg('created_end')::timestamp IS NULL OR o.created_at <= sqlc.narg('created_end')::timestamp)
-- ))
-- AND ((
--     sqlc.narg('updated_start')::timestamp IS NULL AND sqlc.narg('updated_end')::timestamp  IS NULL
-- ) OR (
--     (o.updated_at >= sqlc.narg('updated_start')::timestamp OR sqlc.narg('updated_start')::timestamp  IS NULL) AND
--     (o.updated_at <= sqlc.narg('updated_end')::timestamp OR sqlc.narg('updated_end')::timestamp  IS NULL)
-- ))
ORDER BY
    CASE WHEN sqlc.narg('order_by')::varchar = 'exprise' THEN dn.exprise END DESC,
    CASE WHEN sqlc.narg('order_by')::varchar = '-exprise' THEN dn.exprise END ASC,
    CASE WHEN sqlc.narg('order_by')::varchar = 'dabt_note_at' THEN dn.dabt_note_at END DESC,
    CASE WHEN sqlc.narg('order_by')::varchar = '-dabt_note_at' THEN dn.dabt_note_at END ASC,
    CASE WHEN sqlc.narg('order_by')::varchar IS NULL THEN dn.id END DESC
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateDebtNote :one
INSERT INTO debt_note (
    code, title, entity, money, paymented, note, type, status, company, user_created, exprise, dabt_note_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: UpdateDebtNote :one
UPDATE debt_note
SET
    status = COALESCE(sqlc.narg(status)::varchar, status),
    paymented = COALESCE(sqlc.narg(paymented)::float, paymented)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DetailDebtNote :one
SELECT *, a.full_name AS a_name, c.full_name AS c_name, s.name AS s_name FROM debt_note dn
LEFT JOIN accounts a ON a.id = dn.user_created
LEFT JOIN customers c ON c.code = dn.entity
LEFT JOIN suplier s ON s.code = dn.entity
WHERE dn.id = $1;

-- name: ListRepayment :many
SELECT *, a.full_name AS a_name FROM debt_repayment dr
LEFT JOIN accounts a ON a.id = dr.user_created
WHERE dr.debt = $1;

-- name: CreateRepayment :one
INSERT INTO debt_repayment (
    code, money, debt, user_created
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ReportChartDebtNote :many
WITH date_series AS (
    SELECT generate_series(current_date - interval '7 days', current_date, interval '1 day')::timestamp AS date
)
SELECT date_series.date AS truncated_date,
       COALESCE(SUM(dn.money), 0)::float AS total_money,
       COALESCE(COUNT(dn.*), 0)::int AS ticket
FROM date_series
LEFT JOIN debt_note dn 
ON DATE_TRUNC('day', dn.dabt_note_at) = date_series.date
    AND (sqlc.narg('status')::varchar IS NULL OR dn.status = sqlc.narg('status')::varchar)
    AND (sqlc.narg('type')::varchar IS NULL OR dn.type = sqlc.narg('type')::varchar)
    AND dn.company = sqlc.arg('company')::int
GROUP BY date_series.date
ORDER BY date_series.date;

-- name: ReportRevenueDebtNote :many
SELECT ds.code, COALESCE(SUM(money), 0)::float AS money, COUNT(dn.*) AS ticket 
FROM debt_note_status ds
LEFT JOIN debt_note dn 
    ON  ds.code = dn.status
    AND (sqlc.narg('type')::varchar IS NULL OR dn.type = sqlc.narg('type')::varchar)
    AND dn.company = sqlc.arg('company')::int
GROUP BY ds.code;
