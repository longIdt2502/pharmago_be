// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: variant.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getInventoryVariant = `-- name: GetInventoryVariant :one
SELECT COALESCE(SUM(inventory), 0)::int AS total_inventory
FROM consignment
WHERE variant = $1 AND is_available = true
`

func (q *Queries) GetInventoryVariant(ctx context.Context, variant sql.NullInt32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getInventoryVariant, variant)
	var total_inventory int32
	err := row.Scan(&total_inventory)
	return total_inventory, err
}

const getVariantById = `-- name: GetVariantById :one
SELECT id, name, code, barcode, decision_number, register_number, longevity, vat, product, user_created, user_updated, updated_at, created_at, initial_inventory, real_inventory FROM variants
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetVariantById(ctx context.Context, id int32) (Variant, error) {
	row := q.db.QueryRowContext(ctx, getVariantById, id)
	var i Variant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Barcode,
		&i.DecisionNumber,
		&i.RegisterNumber,
		&i.Longevity,
		&i.Vat,
		&i.Product,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.InitialInventory,
		&i.RealInventory,
	)
	return i, err
}

const getVariants = `-- name: GetVariants :many
SELECT v.id, v.name, v.code, barcode, decision_number, register_number, longevity, vat, product, v.user_created, v.user_updated, v.updated_at, v.created_at, initial_inventory, real_inventory, p.id, p.name, p.code, product_category, type, brand, p.unit, ta_duoc, nong_do, lieu_dung, chi_dinh, chong_chi_dinh, cong_dung, tac_dung_phu, than_trong, tuong_tac, bao_quan, dong_goi, phan_loai, dang_bao_che, tieu_chuan_sx, cong_ty_sx, cong_ty_dk, active, company, p.user_created, p.user_updated, p.updated_at, p.created_at, vm.id, variant, media, m.id, media_url, u.id, u.name, sell_price, import_price, weight, weight_unit, u.user_created, u.user_updated, u.updated_at, u.created_at, pl.id, variant_code, variant_name, price_import, price_sell, pl.unit, pl.user_created, pl.user_updated, pl.updated_at, pl.created_at, m.media_url AS media,
       u.id AS unit_id, u.name AS unit_name, u.sell_price AS unit_sell_price, u.weight AS unit_weight, u.weight_unit AS unit_weight_unit,
       pl.price_import AS pl_price_import, pl.price_sell AS pl_price_sell
FROM variants v
JOIN products p ON v.product = p.id
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
JOIN units u ON u.id = p.unit
LEFT JOIN price_list pl ON pl.variant_code = v.code
WHERE (p.company = $1::int OR v.product = $2::int)
AND (
    v.name ILIKE '%' || COALESCE($3::varchar, '') || '%' OR
    v.code ILIKE '%' || COALESCE($3::varchar, '') || '%' OR
    v.barcode ILIKE '%' || COALESCE($3::varchar, '') || '%'
) AND (
    $4::int IS NULL OR v.id = $4::int
)
ORDER BY -v.id
LIMIT COALESCE($6::int, 10)
OFFSET (COALESCE($5::int, 1) - 1) * COALESCE($6::int, 10)
`

type GetVariantsParams struct {
	Company int32          `json:"company"`
	Product int32          `json:"product"`
	Search  sql.NullString `json:"search"`
	ID      sql.NullInt32  `json:"id"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type GetVariantsRow struct {
	ID               int32           `json:"id"`
	Name             string          `json:"name"`
	Code             string          `json:"code"`
	Barcode          sql.NullString  `json:"barcode"`
	DecisionNumber   sql.NullString  `json:"decision_number"`
	RegisterNumber   sql.NullString  `json:"register_number"`
	Longevity        sql.NullString  `json:"longevity"`
	Vat              sql.NullFloat64 `json:"vat"`
	Product          int32           `json:"product"`
	UserCreated      int32           `json:"user_created"`
	UserUpdated      sql.NullInt32   `json:"user_updated"`
	UpdatedAt        sql.NullTime    `json:"updated_at"`
	CreatedAt        time.Time       `json:"created_at"`
	InitialInventory int32           `json:"initial_inventory"`
	RealInventory    int32           `json:"real_inventory"`
	ID_2             int32           `json:"id_2"`
	Name_2           string          `json:"name_2"`
	Code_2           string          `json:"code_2"`
	ProductCategory  sql.NullInt32   `json:"product_category"`
	Type             sql.NullInt32   `json:"type"`
	Brand            sql.NullInt32   `json:"brand"`
	Unit             int32           `json:"unit"`
	TaDuoc           sql.NullString  `json:"ta_duoc"`
	NongDo           sql.NullString  `json:"nong_do"`
	LieuDung         sql.NullString  `json:"lieu_dung"`
	ChiDinh          sql.NullString  `json:"chi_dinh"`
	ChongChiDinh     sql.NullString  `json:"chong_chi_dinh"`
	CongDung         sql.NullString  `json:"cong_dung"`
	TacDungPhu       sql.NullString  `json:"tac_dung_phu"`
	ThanTrong        sql.NullString  `json:"than_trong"`
	TuongTac         sql.NullString  `json:"tuong_tac"`
	BaoQuan          sql.NullString  `json:"bao_quan"`
	DongGoi          sql.NullString  `json:"dong_goi"`
	PhanLoai         sql.NullString  `json:"phan_loai"`
	DangBaoChe       sql.NullString  `json:"dang_bao_che"`
	TieuChuanSx      sql.NullString  `json:"tieu_chuan_sx"`
	CongTySx         sql.NullInt32   `json:"cong_ty_sx"`
	CongTyDk         sql.NullInt32   `json:"cong_ty_dk"`
	Active           bool            `json:"active"`
	Company          int32           `json:"company"`
	UserCreated_2    int32           `json:"user_created_2"`
	UserUpdated_2    sql.NullInt32   `json:"user_updated_2"`
	UpdatedAt_2      sql.NullTime    `json:"updated_at_2"`
	CreatedAt_2      time.Time       `json:"created_at_2"`
	ID_3             sql.NullInt32   `json:"id_3"`
	Variant          sql.NullInt32   `json:"variant"`
	Media            sql.NullInt32   `json:"media"`
	ID_4             sql.NullInt32   `json:"id_4"`
	MediaUrl         sql.NullString  `json:"media_url"`
	ID_5             int32           `json:"id_5"`
	Name_3           string          `json:"name_3"`
	SellPrice        float64         `json:"sell_price"`
	ImportPrice      float64         `json:"import_price"`
	Weight           sql.NullFloat64 `json:"weight"`
	WeightUnit       sql.NullString  `json:"weight_unit"`
	UserCreated_3    int32           `json:"user_created_3"`
	UserUpdated_3    sql.NullInt32   `json:"user_updated_3"`
	UpdatedAt_3      sql.NullTime    `json:"updated_at_3"`
	CreatedAt_3      time.Time       `json:"created_at_3"`
	ID_6             sql.NullInt32   `json:"id_6"`
	VariantCode      sql.NullString  `json:"variant_code"`
	VariantName      sql.NullString  `json:"variant_name"`
	PriceImport      sql.NullFloat64 `json:"price_import"`
	PriceSell        sql.NullFloat64 `json:"price_sell"`
	Unit_2           sql.NullInt32   `json:"unit_2"`
	UserCreated_4    sql.NullInt32   `json:"user_created_4"`
	UserUpdated_4    sql.NullInt32   `json:"user_updated_4"`
	UpdatedAt_4      sql.NullTime    `json:"updated_at_4"`
	CreatedAt_4      sql.NullTime    `json:"created_at_4"`
	Media_2          sql.NullString  `json:"media_2"`
	UnitID           int32           `json:"unit_id"`
	UnitName         string          `json:"unit_name"`
	UnitSellPrice    float64         `json:"unit_sell_price"`
	UnitWeight       sql.NullFloat64 `json:"unit_weight"`
	UnitWeightUnit   sql.NullString  `json:"unit_weight_unit"`
	PlPriceImport    sql.NullFloat64 `json:"pl_price_import"`
	PlPriceSell      sql.NullFloat64 `json:"pl_price_sell"`
}

func (q *Queries) GetVariants(ctx context.Context, arg GetVariantsParams) ([]GetVariantsRow, error) {
	rows, err := q.db.QueryContext(ctx, getVariants,
		arg.Company,
		arg.Product,
		arg.Search,
		arg.ID,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetVariantsRow{}
	for rows.Next() {
		var i GetVariantsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Barcode,
			&i.DecisionNumber,
			&i.RegisterNumber,
			&i.Longevity,
			&i.Vat,
			&i.Product,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.InitialInventory,
			&i.RealInventory,
			&i.ID_2,
			&i.Name_2,
			&i.Code_2,
			&i.ProductCategory,
			&i.Type,
			&i.Brand,
			&i.Unit,
			&i.TaDuoc,
			&i.NongDo,
			&i.LieuDung,
			&i.ChiDinh,
			&i.ChongChiDinh,
			&i.CongDung,
			&i.TacDungPhu,
			&i.ThanTrong,
			&i.TuongTac,
			&i.BaoQuan,
			&i.DongGoi,
			&i.PhanLoai,
			&i.DangBaoChe,
			&i.TieuChuanSx,
			&i.CongTySx,
			&i.CongTyDk,
			&i.Active,
			&i.Company,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
			&i.ID_3,
			&i.Variant,
			&i.Media,
			&i.ID_4,
			&i.MediaUrl,
			&i.ID_5,
			&i.Name_3,
			&i.SellPrice,
			&i.ImportPrice,
			&i.Weight,
			&i.WeightUnit,
			&i.UserCreated_3,
			&i.UserUpdated_3,
			&i.UpdatedAt_3,
			&i.CreatedAt_3,
			&i.ID_6,
			&i.VariantCode,
			&i.VariantName,
			&i.PriceImport,
			&i.PriceSell,
			&i.Unit_2,
			&i.UserCreated_4,
			&i.UserUpdated_4,
			&i.UpdatedAt_4,
			&i.CreatedAt_4,
			&i.Media_2,
			&i.UnitID,
			&i.UnitName,
			&i.UnitSellPrice,
			&i.UnitWeight,
			&i.UnitWeightUnit,
			&i.PlPriceImport,
			&i.PlPriceSell,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVariantsByCode = `-- name: GetVariantsByCode :one
SELECT v.id, v.name, v.code, barcode, decision_number, register_number, longevity, vat, product, v.user_created, v.user_updated, v.updated_at, v.created_at, initial_inventory, real_inventory, p.id, p.name, p.code, product_category, type, brand, p.unit, ta_duoc, nong_do, lieu_dung, chi_dinh, chong_chi_dinh, cong_dung, tac_dung_phu, than_trong, tuong_tac, bao_quan, dong_goi, phan_loai, dang_bao_che, tieu_chuan_sx, cong_ty_sx, cong_ty_dk, active, company, p.user_created, p.user_updated, p.updated_at, p.created_at, vm.id, variant, media, m.id, media_url, u.id, u.name, sell_price, import_price, weight, weight_unit, u.user_created, u.user_updated, u.updated_at, u.created_at, pl.id, variant_code, variant_name, price_import, price_sell, pl.unit, pl.user_created, pl.user_updated, pl.updated_at, pl.created_at, m.media_url AS media,
       u.id AS unit_id, u.name AS unit_name, u.sell_price AS unit_sell_price, u.weight AS unit_weight, u.weight_unit AS unit_weight_unit,
       pl.price_import AS pl_price_import, pl.price_sell AS pl_price_sell
FROM variants v
         JOIN products p ON v.product = p.id
         LEFT JOIN variant_media vm ON vm.variant = v.id
         LEFT JOIN medias m ON m.id = vm.media
         JOIN units u ON u.id = p.unit
         JOIN price_list pl ON pl.variant_code = v.code
WHERE p.company = $1::int
AND v.barcode = $2::varchar
LIMIT 1
`

type GetVariantsByCodeParams struct {
	Company int32  `json:"company"`
	Code    string `json:"code"`
}

type GetVariantsByCodeRow struct {
	ID               int32           `json:"id"`
	Name             string          `json:"name"`
	Code             string          `json:"code"`
	Barcode          sql.NullString  `json:"barcode"`
	DecisionNumber   sql.NullString  `json:"decision_number"`
	RegisterNumber   sql.NullString  `json:"register_number"`
	Longevity        sql.NullString  `json:"longevity"`
	Vat              sql.NullFloat64 `json:"vat"`
	Product          int32           `json:"product"`
	UserCreated      int32           `json:"user_created"`
	UserUpdated      sql.NullInt32   `json:"user_updated"`
	UpdatedAt        sql.NullTime    `json:"updated_at"`
	CreatedAt        time.Time       `json:"created_at"`
	InitialInventory int32           `json:"initial_inventory"`
	RealInventory    int32           `json:"real_inventory"`
	ID_2             int32           `json:"id_2"`
	Name_2           string          `json:"name_2"`
	Code_2           string          `json:"code_2"`
	ProductCategory  sql.NullInt32   `json:"product_category"`
	Type             sql.NullInt32   `json:"type"`
	Brand            sql.NullInt32   `json:"brand"`
	Unit             int32           `json:"unit"`
	TaDuoc           sql.NullString  `json:"ta_duoc"`
	NongDo           sql.NullString  `json:"nong_do"`
	LieuDung         sql.NullString  `json:"lieu_dung"`
	ChiDinh          sql.NullString  `json:"chi_dinh"`
	ChongChiDinh     sql.NullString  `json:"chong_chi_dinh"`
	CongDung         sql.NullString  `json:"cong_dung"`
	TacDungPhu       sql.NullString  `json:"tac_dung_phu"`
	ThanTrong        sql.NullString  `json:"than_trong"`
	TuongTac         sql.NullString  `json:"tuong_tac"`
	BaoQuan          sql.NullString  `json:"bao_quan"`
	DongGoi          sql.NullString  `json:"dong_goi"`
	PhanLoai         sql.NullString  `json:"phan_loai"`
	DangBaoChe       sql.NullString  `json:"dang_bao_che"`
	TieuChuanSx      sql.NullString  `json:"tieu_chuan_sx"`
	CongTySx         sql.NullInt32   `json:"cong_ty_sx"`
	CongTyDk         sql.NullInt32   `json:"cong_ty_dk"`
	Active           bool            `json:"active"`
	Company          int32           `json:"company"`
	UserCreated_2    int32           `json:"user_created_2"`
	UserUpdated_2    sql.NullInt32   `json:"user_updated_2"`
	UpdatedAt_2      sql.NullTime    `json:"updated_at_2"`
	CreatedAt_2      time.Time       `json:"created_at_2"`
	ID_3             sql.NullInt32   `json:"id_3"`
	Variant          sql.NullInt32   `json:"variant"`
	Media            sql.NullInt32   `json:"media"`
	ID_4             sql.NullInt32   `json:"id_4"`
	MediaUrl         sql.NullString  `json:"media_url"`
	ID_5             int32           `json:"id_5"`
	Name_3           string          `json:"name_3"`
	SellPrice        float64         `json:"sell_price"`
	ImportPrice      float64         `json:"import_price"`
	Weight           sql.NullFloat64 `json:"weight"`
	WeightUnit       sql.NullString  `json:"weight_unit"`
	UserCreated_3    int32           `json:"user_created_3"`
	UserUpdated_3    sql.NullInt32   `json:"user_updated_3"`
	UpdatedAt_3      sql.NullTime    `json:"updated_at_3"`
	CreatedAt_3      time.Time       `json:"created_at_3"`
	ID_6             int32           `json:"id_6"`
	VariantCode      string          `json:"variant_code"`
	VariantName      string          `json:"variant_name"`
	PriceImport      float64         `json:"price_import"`
	PriceSell        float64         `json:"price_sell"`
	Unit_2           int32           `json:"unit_2"`
	UserCreated_4    int32           `json:"user_created_4"`
	UserUpdated_4    sql.NullInt32   `json:"user_updated_4"`
	UpdatedAt_4      sql.NullTime    `json:"updated_at_4"`
	CreatedAt_4      time.Time       `json:"created_at_4"`
	Media_2          sql.NullString  `json:"media_2"`
	UnitID           int32           `json:"unit_id"`
	UnitName         string          `json:"unit_name"`
	UnitSellPrice    float64         `json:"unit_sell_price"`
	UnitWeight       sql.NullFloat64 `json:"unit_weight"`
	UnitWeightUnit   sql.NullString  `json:"unit_weight_unit"`
	PlPriceImport    float64         `json:"pl_price_import"`
	PlPriceSell      float64         `json:"pl_price_sell"`
}

func (q *Queries) GetVariantsByCode(ctx context.Context, arg GetVariantsByCodeParams) (GetVariantsByCodeRow, error) {
	row := q.db.QueryRowContext(ctx, getVariantsByCode, arg.Company, arg.Code)
	var i GetVariantsByCodeRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Barcode,
		&i.DecisionNumber,
		&i.RegisterNumber,
		&i.Longevity,
		&i.Vat,
		&i.Product,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.InitialInventory,
		&i.RealInventory,
		&i.ID_2,
		&i.Name_2,
		&i.Code_2,
		&i.ProductCategory,
		&i.Type,
		&i.Brand,
		&i.Unit,
		&i.TaDuoc,
		&i.NongDo,
		&i.LieuDung,
		&i.ChiDinh,
		&i.ChongChiDinh,
		&i.CongDung,
		&i.TacDungPhu,
		&i.ThanTrong,
		&i.TuongTac,
		&i.BaoQuan,
		&i.DongGoi,
		&i.PhanLoai,
		&i.DangBaoChe,
		&i.TieuChuanSx,
		&i.CongTySx,
		&i.CongTyDk,
		&i.Active,
		&i.Company,
		&i.UserCreated_2,
		&i.UserUpdated_2,
		&i.UpdatedAt_2,
		&i.CreatedAt_2,
		&i.ID_3,
		&i.Variant,
		&i.Media,
		&i.ID_4,
		&i.MediaUrl,
		&i.ID_5,
		&i.Name_3,
		&i.SellPrice,
		&i.ImportPrice,
		&i.Weight,
		&i.WeightUnit,
		&i.UserCreated_3,
		&i.UserUpdated_3,
		&i.UpdatedAt_3,
		&i.CreatedAt_3,
		&i.ID_6,
		&i.VariantCode,
		&i.VariantName,
		&i.PriceImport,
		&i.PriceSell,
		&i.Unit_2,
		&i.UserCreated_4,
		&i.UserUpdated_4,
		&i.UpdatedAt_4,
		&i.CreatedAt_4,
		&i.Media_2,
		&i.UnitID,
		&i.UnitName,
		&i.UnitSellPrice,
		&i.UnitWeight,
		&i.UnitWeightUnit,
		&i.PlPriceImport,
		&i.PlPriceSell,
	)
	return i, err
}

const getVariantsByProduct = `-- name: GetVariantsByProduct :many
SELECT id, name, code, barcode, decision_number, register_number, longevity, vat, product, user_created, user_updated, updated_at, created_at, initial_inventory, real_inventory FROM variants
WHERE product = $1
`

func (q *Queries) GetVariantsByProduct(ctx context.Context, product int32) ([]Variant, error) {
	rows, err := q.db.QueryContext(ctx, getVariantsByProduct, product)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Variant{}
	for rows.Next() {
		var i Variant
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Barcode,
			&i.DecisionNumber,
			&i.RegisterNumber,
			&i.Longevity,
			&i.Vat,
			&i.Product,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.InitialInventory,
			&i.RealInventory,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const variantsCustomerBuy = `-- name: VariantsCustomerBuy :many
SELECT v.id, v.name, v.code, v.barcode, v.decision_number, v.register_number, v.longevity, v.vat, v.product, v.user_created, v.user_updated, v.updated_at, v.created_at, v.initial_inventory, v.real_inventory, u.id, u.name, u.sell_price, u.import_price, u.weight, u.weight_unit, u.user_created, u.user_updated, u.updated_at, u.created_at, SUM(value) AS quantity_buy FROM order_items oi
JOIN variants v ON v.id = oi.variant
JOIN orders o ON o.id = oi.order
JOIN products p ON p.id = v.product
JOIN units u ON u.id = p.unit
WHERE o.customer = $1::int
GROUP BY oi.variant, v.id, u.id
LIMIT COALESCE($3::int, 10)
OFFSET (COALESCE($2::int, 1) - 1) * COALESCE($3::int, 10)
`

type VariantsCustomerBuyParams struct {
	Customer int32         `json:"customer"`
	Page     sql.NullInt32 `json:"page"`
	Limit    sql.NullInt32 `json:"limit"`
}

type VariantsCustomerBuyRow struct {
	ID               int32           `json:"id"`
	Name             string          `json:"name"`
	Code             string          `json:"code"`
	Barcode          sql.NullString  `json:"barcode"`
	DecisionNumber   sql.NullString  `json:"decision_number"`
	RegisterNumber   sql.NullString  `json:"register_number"`
	Longevity        sql.NullString  `json:"longevity"`
	Vat              sql.NullFloat64 `json:"vat"`
	Product          int32           `json:"product"`
	UserCreated      int32           `json:"user_created"`
	UserUpdated      sql.NullInt32   `json:"user_updated"`
	UpdatedAt        sql.NullTime    `json:"updated_at"`
	CreatedAt        time.Time       `json:"created_at"`
	InitialInventory int32           `json:"initial_inventory"`
	RealInventory    int32           `json:"real_inventory"`
	ID_2             int32           `json:"id_2"`
	Name_2           string          `json:"name_2"`
	SellPrice        float64         `json:"sell_price"`
	ImportPrice      float64         `json:"import_price"`
	Weight           sql.NullFloat64 `json:"weight"`
	WeightUnit       sql.NullString  `json:"weight_unit"`
	UserCreated_2    int32           `json:"user_created_2"`
	UserUpdated_2    sql.NullInt32   `json:"user_updated_2"`
	UpdatedAt_2      sql.NullTime    `json:"updated_at_2"`
	CreatedAt_2      time.Time       `json:"created_at_2"`
	QuantityBuy      int64           `json:"quantity_buy"`
}

func (q *Queries) VariantsCustomerBuy(ctx context.Context, arg VariantsCustomerBuyParams) ([]VariantsCustomerBuyRow, error) {
	rows, err := q.db.QueryContext(ctx, variantsCustomerBuy, arg.Customer, arg.Page, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VariantsCustomerBuyRow{}
	for rows.Next() {
		var i VariantsCustomerBuyRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Barcode,
			&i.DecisionNumber,
			&i.RegisterNumber,
			&i.Longevity,
			&i.Vat,
			&i.Product,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.InitialInventory,
			&i.RealInventory,
			&i.ID_2,
			&i.Name_2,
			&i.SellPrice,
			&i.ImportPrice,
			&i.Weight,
			&i.WeightUnit,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
			&i.QuantityBuy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
