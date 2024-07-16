-- name: GetByVariantOrService :many
SELECT *, v.name AS v_name, v.code AS v_code, s.title AS s_name, s.code AS s_code FROM promotion_item pi
LEFT JOIN promotions p ON p.id = pi.promotions
LEFT JOIN variants v ON pi.variant = v.id
LEFT JOIN services s ON pi.service = s.id
LEFT JOIN promotion_type pt ON p.type = pt.code 
WHERE (pi.applicable_variant = sqlc.narg('variant')::int OR pi.applicable_service = sqlc.narg('service')::int)
AND p.type = 'GIFT';

-- name: CreatePromotion :one
INSERT INTO promotions (
  id, code, type, title, conditions_text, conditions_point_customer, min_value, 
  is_discount_percent, value_discount, max_discount, time_apply, date_start, date_end, 
  apply_multiple_times, apply_simultaneously, status, company, user_created, 
  user_updated
) VALUES (
  sqlc.arg(id)::uuid, sqlc.arg(code)::varchar, sqlc.narg(type)::varchar, sqlc.arg(title)::varchar, 
  sqlc.narg(conditions_text)::varchar, sqlc.narg(conditions_point_customer)::int, sqlc.arg(min_value)::float, 
  sqlc.arg(is_discount_percent)::bool, sqlc.arg(value_discount)::float, sqlc.arg(max_discount)::float, sqlc.narg(time_apply)::int, 
  sqlc.narg(date_start)::timestamp, sqlc.narg(date_end)::timestamp, sqlc.arg(apply_multiple_times)::bool, 
  sqlc.arg(apply_simultaneously)::bool, sqlc.arg(status)::bool, sqlc.arg(company)::int, 
  sqlc.arg(user_created)::int, sqlc.narg(user_updated)::int
) RETURNING *;

-- name: CreatePromotionItem :one
INSERT INTO promotion_item (
  id, min_buy, amount_gift, promotions, variant, service, applicable_variant, applicable_service
) VALUES (
  sqlc.arg(id)::uuid, sqlc.arg(min_buy)::int, sqlc.arg(amount_gift)::int, sqlc.arg(promotions)::uuid, 
  sqlc.narg(variant)::int, sqlc.narg(service)::int, sqlc.narg(applicable_variant)::int, sqlc.narg(applicable_service)::int
) RETURNING *;

-- name: GetPromotionByPriceOrder :many
SELECT * FROM promotions
WHERE min_value >= sqlc.arg('price')::float
AND company = sqlc.arg(company)::int
AND "type" = 'DISCOUNT';