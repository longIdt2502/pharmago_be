-- name: GetListDebtNote :many
-- WITH total_repayment AS (
--     SELECT debt, COALESCE(SUM(money), 0)::float AS total_money
--     FROM debt_repayment
--     GROUP BY debt
-- )
SELECT * FROM debt_note dn
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

-- name: DetailDebtNote :one
SELECT * FROM debt_note
WHERE id = $1;

-- name: ListRepayment :many
SELECT * FROM debt_repayment
WHERE debt = $1;

-- name: CreateRepayment :one
INSERT INTO debt_repayment (
    code, money, debt, user_created
) VALUES (
    $1, $2, $3, $4
) RETURNING *;