-- name: CreatePrescriptionItem :one
INSERT INTO prescription_item (
    prescription_uuid, variant, lieu_dung, quantity
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreatePrescription :one
INSERT INTO prescriptions (
    uuid, code, symptoms, diagnostic, doctor, user_created
) VALUES (
    $1, $2, $3, $4, $5, $6
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