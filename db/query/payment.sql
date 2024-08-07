-- name: CreatePayment :one
INSERT INTO payments (
    code, must_paid, had_paid, need_pay
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: CreatePaymentItem :one
INSERT INTO payment_items (
    type, value, is_paid, payment, extra_note
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListPaymentItem :many
SELECT * FROM payment_items pi
JOIN payment_item_types pit ON pi.type = pit.code
WHERE payment = $1;

-- name: DetailPayment :one
SELECT * FROM payments
WHERE id = $1
LIMIT 1;

-- name: PaymentOrderByMedicalBill :one
SELECT COALESCE(SUM(p.must_paid), 0)::float AS total_must_paid, 
        COALESCE(SUM(p.had_paid), 0)::float AS total_had_paid, 
        COALESCE(SUM(p.need_pay), 0)::float AS total_need_pay
    FROM medical_bill_order_sell mbos
JOIN orders o ON o.id = mbos.order
JOIN payments p ON p.id = o.payment
WHERE mbos.uuid = sqlc.arg(uuid)::uuid
GROUP BY mbos.uuid;

-- name: PaymentOrderServiceByMedicalBill :one
SELECT COALESCE(SUM(p.must_paid), 0)::float AS total_must_paid, 
        COALESCE(SUM(p.had_paid), 0)::float AS total_had_paid, 
        COALESCE(SUM(p.need_pay), 0)::float AS total_need_pay
    FROM appointment_schedule_service ass
JOIN orders o ON o.id = ass.order_service
JOIN payments p ON p.id = o.payment
WHERE ass.mb_uuid = sqlc.arg(uuid)::uuid
GROUP BY ass.mb_uuid;

-- name: UpdatePayment :one
UPDATE payments
SET had_paid = COALESCE(sqlc.narg(had_paid)::float, had_paid),
    need_pay = COALESCE(sqlc.narg(need_pay)::float, need_pay)
WHERE id = sqlc.arg(id)::int
RETURNING *;