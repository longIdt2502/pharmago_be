-- name: GetListSchedule :many
SELECT * FROM appointment_schedules sch
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

-- name: DetailSchedule :one
SELECT * FROM appointment_schedules
WHERE uuid = $1;

-- name: UpdateSchedule :one
UPDATE appointment_schedules
SET
    is_done = COALESCE(sqlc.narg(is_done)::bool, is_done),
    diagnostic = COALESCE(sqlc.narg(diagnostic)::varchar, diagnostic)
WHERE uuid = sqlc.arg(uuid)::uuid
RETURNING *;

-- name: GetListScheduleService :many
SELECT * FROM appointment_schedule_service ass
JOIN services s ON s.id = ass.service
LEFT JOIN orders os ON os.id = ass.order_service
WHERE ass.as_uuid = $1 OR ass.mb_uuid = $2;

-- name: GetListScheduleUrl :many
SELECT * FROM appointment_schedule_url
WHERE as_uuid = $1 OR mb_uuid = $2;

-- name: GetListScheduleDrug :many
SELECT * FROM appointment_schedule_drug asd
JOIN variants v ON v.id = asd.variant
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
WHERE asd.as_uuid = $1 OR asd.mb_uuid = $2;

-- name: CreateSchedule :one
INSERT INTO appointment_schedules (
    uuid, code, customer, company, doctor, symptoms, diagnostic, is_done, meeting_at, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: CreateScheduleService :one
INSERT INTO appointment_schedule_service (
    as_uuid, mb_uuid, "service", order_service
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateScheduleUrl :one
INSERT INTO appointment_schedule_url (
    as_uuid, mb_uuid, url, name_doc
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreateScheduleDrug :one
INSERT INTO appointment_schedule_drug (
    as_uuid, mb_uuid, variant, lieu_dung, quantity
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;
