-- name: CreatePrescriptionItem :one
INSERT INTO prescription_item (
    prescription_uuid, variant, lieu_dung, quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreatePrescription :one
INSERT INTO prescriptions (
    uuid, code, symptoms, diagnostic, customer, doctor, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: ListPrescriptionItem :many
SELECT * FROM prescription_item pi 
JOIN variants v ON pi.variant = v.id
WHERE pi.prescription_uuid = $1;

-- name: DetailPrescription :one
SELECT * FROM prescriptions p
JOIN customers c ON c.id = p.customer
JOIN accounts a ON a.id = p.doctor
JOIN accounts uc ON uc.id = p.user_created
LEFT JOIN accounts uu ON uu.id = p.user_updated
WHERE p.uuid = $1;

-- name: ListPrescription :many
SELECT * FROM prescriptions p
JOIN customers c ON c.id = p.customer
JOIN accounts a ON a.id = p.doctor
JOIN accounts uc ON uc.id = p.user_created
LEFT JOIN accounts uu ON uu.id = p.user_updated
WHERE p.company = sqlc.narg('company')::int
AND (
    p.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    c.full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 1);

-- name: UpdatePrescription :one
UPDATE prescriptions
SET 
    code = COALESCE(sqlc.narg(code)::varchar, code),
    diagnostic = COALESCE(sqlc.narg(diagnostic)::varchar, diagnostic),
    customer = COALESCE(sqlc.narg(customer)::int, customer)
WHERE uuid = sqlc.arg(uuid)::uuid
RETURNING *;

-- name: UpdatePrescriptionItem :one
UPDATE prescription_item
SET 
    lieu_dung = COALESCE(sqlc.narg(lieu_dung)::varchar, lieu_dung),
    quantity = COALESCE(sqlc.narg(quantity)::int, quantity)
WHERE id = sqlc.arg(id)::int
RETURNING *;

-- name: DeletePrescriptionItem :one
DELETE FROM prescription_item
WHERE id = $1 RETURNING *;

