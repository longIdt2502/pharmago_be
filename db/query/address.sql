-- name: GetProvince :many
SELECT * FROM provinces;

-- name: GetProvinceByCode :one
SELECT * FROM provinces
WHERE code = $1 LIMIT 1;

-- name: GetDistrict :many
SELECT * FROM districts
WHERE province_code = sqlc.narg(province_code);

-- name: GetDistrictByCode :one
SELECT * FROM districts
WHERE code = $1 LIMIT 1;

-- name: GetWard :many
SELECT * FROM wards
WHERE district_code = sqlc.narg(district_code);

-- name: GetWardByCode :one
SELECT * FROM wards
WHERE code = $1 LIMIT 1;

-- name: CreateAddress :one
INSERT INTO address (
    lat, lng, province, district, ward, title, user_created
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;