-- name: CreateProduct :one
INSERT INTO products (
    name, code, product_category, type, brand, unit, ta_duoc, nong_do, lieu_dung, chi_dinh, chong_chi_dinh,
    cong_dung, tac_dung_phu, than_trong, tuong_tac, bao_quan, dong_goi, cong_ty_sx, cong_ty_dk,
    company, user_created, user_updated, phan_loai, dang_bao_che, tieu_chuan_sx
) values (
    sqlc.arg(name)::varchar, sqlc.arg(code)::varchar, sqlc.narg(product_category)::int, sqlc.narg(type)::int, sqlc.narg(brand)::int,
    sqlc.arg(unit)::int, sqlc.narg(taDuoc)::varchar, sqlc.narg(nongDo)::varchar, sqlc.narg(lieuDung)::varchar,
    sqlc.narg(chiDinh)::varchar, sqlc.narg(chongChiDinh)::varchar, sqlc.narg(congDung)::varchar, sqlc.narg(tacDungPhu)::varchar,
    sqlc.narg(thanTrong)::varchar, sqlc.narg(tuongTac)::varchar, sqlc.narg(baoQuan)::varchar, sqlc.narg(dongGoi)::varchar,
    sqlc.narg(congTySx)::int, sqlc.narg(congTyDk)::int, sqlc.arg(company)::int,
    sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int, sqlc.narg(phanLoai)::varchar, sqlc.narg(dangBaoche)::varchar, sqlc.narg(tieuChuanSx)::varchar
) RETURNING *;

-- name: CreateVariant :one
INSERT INTO variants (
    name, code, barcode, vat, decision_number, register_number, longevity, product, user_created, user_updated, initial_inventory, real_inventory
) values (
    sqlc.arg(name)::varchar, sqlc.arg(code)::varchar, sqlc.narg(barcode)::varchar, sqlc.narg(vat)::float,
    sqlc.narg(decision_number)::varchar, sqlc.narg(register_number)::varchar, sqlc.narg(longevity)::varchar,
    sqlc.arg(product)::int, sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int, sqlc.arg(initial_inventory)::int, sqlc.arg(real_inventory)::int
) RETURNING *;

-- name: UpdateVariant :one
UPDATE variants 
SET
    real_inventory = COALESCE(sqlc.narg(real_inventory)::int, real_inventory)
WHERE id = $1
RETURNING *;

-- name: GetVariantByProduct :one
SELECT * FROM variants
WHERE product = $1;

-- name: CreateUnit :one
INSERT INTO units (
    name, sell_price, import_price, weight, weight_unit, user_created, user_updated
) values (
    sqlc.arg(name)::varchar, sqlc.arg(sell_price)::float, sqlc.arg(import_price)::float, sqlc.narg(weight)::float,
    sqlc.narg(weight_unit)::varchar, sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int
) RETURNING *;

-- name: CreateUnitChange :one
INSERT INTO unit_changes (
    name, value, sell_price, unit, user_created, user_updated
) values (
    sqlc.arg(name)::varchar, sqlc.arg(value)::int, sqlc.arg(sell_price)::float, sqlc.narg(unit)::int,
    sqlc.arg(user_created)::int, sqlc.arg(user_updated)::int
) RETURNING *;

-- name: CreateIngredient :one
INSERT INTO ingredient (
    name, weight, unit, product
) values (
    sqlc.arg(name)::varchar, sqlc.arg(weight)::float, sqlc.arg(unit)::varchar, sqlc.narg(product)::int
) RETURNING *;

-- name: ListIngredient :many
SELECT * FROM ingredient
WHERE product = $1;

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
    code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    (sqlc.narg(brand)::int IS NULL OR brand = sqlc.narg(brand)::int) OR
    (sqlc.narg(product_category)::int IS NULL OR product_category = sqlc.narg(product_category)::int)
) AND (
    sqlc.narg(active)::bool IS NULL OR sqlc.narg(active)::bool = active
)
ORDER BY -id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: GetProductsByCode :one
SELECT * FROM products
WHERE code = $1;

-- name: DetailProduct :one
SELECT * FROM products p
LEFT JOIN product_categories pc ON pc.id = p.product_category
LEFT JOIN product_type pt ON pt.id = p.type
LEFT JOIN product_brand pb ON pb.id = p.brand
LEFT JOIN units u ON u.id = p.unit
LEFT JOIN production_standard ps ON ps.code = p.tieu_chuan_sx
LEFT JOIN preparation_type pret ON pret.code = p.dang_bao_che
LEFT JOIN classify cl ON cl.code = p.phan_loai
LEFT JOIN company_pharma cp1 ON cp1.id = p.cong_ty_dk
LEFT JOIN company_pharma cp2 ON cp1.id = p.cong_ty_sx
WHERE p.id = $1;

-- name: UpdateProduct :one
UPDATE products
SET
    brand = COALESCE(sqlc.narg(brand), product_category),
    product_category = COALESCE(sqlc.narg(product_category), product_category),
    active = COALESCE(sqlc.narg(active)::bool, active)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: CountProduct :many
SELECT active, COUNT(active) AS total FROM products
WHERE company = $1
GROUP BY active;
