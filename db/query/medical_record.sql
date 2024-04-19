-- name: ListMedicalRecord :many
SELECT * FROM medical_records
WHERE customer = sqlc.arg(customer)::int
AND (
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CreateMedicalRecord :one
INSERT INTO medical_records (
    code, customer, weight, long, symptom, diagnostic, result, doctor, re_examination, note, user_created 
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: DetailMedicalRecord :one
SELECT * FROM medical_records
WHERE id = $1;