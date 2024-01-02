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