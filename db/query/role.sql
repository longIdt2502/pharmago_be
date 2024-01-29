-- name: CreateRole :one
INSERT INTO roles (
    code, title, note, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: CreateRoleItem :one
INSERT INTO role_item (
    roles, app, value
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListRole :many
SELECT *, ac.full_name AS created_name, au.full_name AS updated_name FROM roles r
JOIN companies c ON c.id = r.company
JOIN accounts ac ON ac.id = r.user_created
JOIN accounts au ON ac.id = r.user_updated
WHERE company = sqlc.arg(company)
AND (
    r.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    r.title ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
AND  ((
        sqlc.narg('created_start')::timestamp IS NULL AND sqlc.narg('created_end')::timestamp  IS NULL
    ) OR (
        (sqlc.narg('created_start')::timestamp IS NULL OR r.created_at >= sqlc.narg('created_start')::timestamp) AND
        (sqlc.narg('created_end')::timestamp IS NULL OR r.created_at <= sqlc.narg('created_end')::timestamp)
))
AND ((
       sqlc.narg('updated_start')::timestamp IS NULL AND sqlc.narg('updated_end')::timestamp  IS NULL
   ) OR (
       (r.updated_at >= sqlc.narg('updated_start')::timestamp OR sqlc.narg('updated_start')::timestamp  IS NULL) AND
       (r.updated_at <= sqlc.narg('updated_end')::timestamp OR sqlc.narg('updated_end')::timestamp  IS NULL)
))
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 1);

-- name: ListApp :many
SELECT * FROM apps
WHERE (parent = sqlc.narg(parent)::varchar
OR level = sqlc.narg(level)::int);