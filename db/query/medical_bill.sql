-- name: GetListMedicalBill :many
SELECT * FROM medical_bills sch
LEFT JOIN customers c ON c.id = sch.customer
JOIN accounts a ON a.id = sch.doctor
JOIN accounts uc ON uc.id = sch.user_created
LEFT JOIN accounts uu ON uu.id = sch.user_updated
WHERE sch.company = sqlc.arg(company)
AND (
    c.full_name ILIKE '%' || COALESCE(sqlc.narg(search)::varchar, '') || '%'
)
AND (sqlc.narg(doctor)::int IS NULL OR sqlc.narg(doctor)::int = a.id)
AND (sqlc.narg(uuid)::uuid IS NULL OR sqlc.narg(uuid)::uuid = sch.uuid)
AND  ((
    sqlc.narg('created_start')::timestamp IS NULL AND sqlc.narg('created_end')::timestamp IS NULL
) OR (
    (sqlc.narg('created_start')::timestamp IS NULL OR sch.created_at >= sqlc.narg('created_start')::timestamp) AND
    (sqlc.narg('created_end')::timestamp IS NULL OR sch.created_at <= sqlc.narg('created_end')::timestamp)
))
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: DetailMedicalBill :one
SELECT * FROM medical_bills
WHERE uuid = $1;

-- name: UpdateMedicalBill :one
UPDATE medical_bills
SET
    diagnostic = COALESCE(sqlc.narg(diagnostic)::varchar, diagnostic),
    symptoms = COALESCE(sqlc.narg(symptoms)::varchar, symptoms)
WHERE uuid = sqlc.arg(uuid)::uuid
RETURNING *;

-- name: CreateMedicalBill :one
INSERT INTO medical_bills (
    uuid, code, customer, company, doctor, symptoms, diagnostic, is_done, meeting_at, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;