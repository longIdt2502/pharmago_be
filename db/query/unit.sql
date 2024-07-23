-- name: GetListUnitChange :many
SELECT * FROM unit_changes
WHERE unit = $1;

-- name: DetailUnit :one
SELECT * FROM units
WHERE id = $1;