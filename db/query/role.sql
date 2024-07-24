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
LEFT JOIN companies c ON c.id = r.company
JOIN accounts ac ON ac.id = r.user_created
LEFT JOIN accounts au ON au.id = r.user_updated
WHERE ((sqlc.narg(company)::int IS NULL AND r.company IS NULL) OR r.company = sqlc.narg(company)::int)
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

-- name: RoleDetail :one
SELECT *, ac.full_name AS created_name, au.full_name AS updated_name FROM roles r
JOIN companies c ON c.id = r.company
JOIN accounts ac ON ac.id = r.user_created
JOIN accounts au ON ac.id = r.user_updated
WHERE r.id = $1;

-- name: ListRoleItem :many
SELECT * FROM role_item ri
JOIN apps a ON ri.app = a.code
WHERE ri.roles = $1;

-- name: UpdateRole :one
UPDATE roles
SET
    code = COALESCE(sqlc.narg(code), code),
    title = COALESCE(sqlc.narg(title), title),
    note = COALESCE(sqlc.narg(note), note),
    user_updated = COALESCE(sqlc.narg(user_updated), user_updated)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: UpdateRoleItem :one
UPDATE role_item
SET
    value = COALESCE(sqlc.narg(value), value)
WHERE roles = sqlc.arg(roles) AND app = sqlc.arg(app)
RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles
WHERE id = $1 RETURNING *;
