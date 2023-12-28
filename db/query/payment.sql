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