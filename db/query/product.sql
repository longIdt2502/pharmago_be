-- name: CreateProduct :one
INSERT INTO products (
    name, code, product_category, type, unit, taDuoc, nongDo, lieuDung, chiDinh, chongChiDinh,
    congDung, tacDungPhu, thanTrong, tuongTac, baoQuan, dongGoi, noiSx, congTySx, congTyDk,
    company, user_created, user_updated
) values (
    sqlc.arg(name)::varchar, sqlc.arg(code)::varchar, sqlc.narg(product_category)::int, sqlc.narg(type)::int,
    sqlc.arg(unit)::int, sqlc.narg(taDuoc)::varchar, sqlc.narg(nongDo)::varchar, sqlc.arg(lieuDung)::varchar,
    sqlc.arg(chiDinh)::varchar, sqlc.narg(chongChiDinh)::varchar, sqlc.arg(congDung)::varchar, sqlc.arg(tacDungPhu)::varchar,
    sqlc.arg(thanTrong)::varchar, sqlc.narg(tuongTac)::varchar, sqlc.arg(baoQuan)::varchar, sqlc.arg(dongGoi)::varchar,
    sqlc.arg(noiSx)::varchar, sqlc.arg(congTySx)::varchar, sqlc.arg(congTyDk)::varchar, sqlc.arg(company)::int,
    sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int
) RETURNING *;

-- name: CreateVariant :one
INSERT INTO variants (
    name, code, barcode, vat, decision_number, register_number, longevity, product, user_created, user_updated
) values (
    sqlc.arg(name)::varchar, sqlc.arg(code)::varchar, sqlc.arg(barcode)::varchar, sqlc.narg(vat)::float,
    sqlc.arg(decision_number)::varchar, sqlc.arg(register_number)::varchar, sqlc.arg(longevity)::varchar,
    sqlc.arg(product)::int, sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int
) RETURNING *;

-- name: CreateUnit :one
INSERT INTO units (
    name, sell_price, import_price, weight, weight_unit, user_created, user_updated
) values (
    sqlc.arg(name)::varchar, sqlc.arg(sell_price)::float, sqlc.arg(import_price)::float, sqlc.narg(weight)::float,
    sqlc.narg(weight_unit)::varchar, sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int
) RETURNING *;

-- name: CreateProductMedia :one
INSERT INTO product_media (
    product, media
) VALUES ($1, $2) RETURNING *;

-- name: CreateVariantMedia :one
INSERT INTO variant_media (
    variant, media
) VALUES ($1, $2) RETURNING *;

-- name: GetProductMedia :many
SELECT pm.product, pm.media, m.media_url FROM product_media pm
JOIN medias m ON pm.media = m.id
WHERE product = sqlc.arg(product);

-- name: GetProducts :many
SELECT * FROM products
WHERE company = sqlc.narg(company)::int AND (
    name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);
