// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: price_list.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createProductPriceList = `-- name: CreateProductPriceList :one
INSERT INTO price_list (
    variant_code, variant_name, unit, price_import, price_sell, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, variant_code, variant_name, price_import, price_sell, unit, user_created, user_updated, updated_at, created_at
`

type CreateProductPriceListParams struct {
	VariantCode string        `json:"variant_code"`
	VariantName string        `json:"variant_name"`
	Unit        int32         `json:"unit"`
	PriceImport float64       `json:"price_import"`
	PriceSell   float64       `json:"price_sell"`
	UserCreated int32         `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
}

func (q *Queries) CreateProductPriceList(ctx context.Context, arg CreateProductPriceListParams) (PriceList, error) {
	row := q.db.QueryRowContext(ctx, createProductPriceList,
		arg.VariantCode,
		arg.VariantName,
		arg.Unit,
		arg.PriceImport,
		arg.PriceSell,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i PriceList
	err := row.Scan(
		&i.ID,
		&i.VariantCode,
		&i.VariantName,
		&i.PriceImport,
		&i.PriceSell,
		&i.Unit,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const detailPriceList = `-- name: DetailPriceList :one
SELECT pl.id, variant_code, variant_name, price_import, price_sell, unit, pl.user_created, pl.user_updated, pl.updated_at, pl.created_at, u.id, u.name, sell_price, import_price, weight, weight_unit, u.user_created, u.user_updated, u.updated_at, u.created_at, v.id, v.name, code, barcode, decision_number, register_number, longevity, vat, product, v.user_created, v.user_updated, v.updated_at, v.created_at, initial_inventory, real_inventory, a.id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, a.created_at, role, gender, licence, dob, address, vm.id, variant, media, m.id, media_url, u.name AS unit_name, m.media_url AS variant_media,
       a.full_name AS user_created_name FROM price_list pl
JOIN units u ON pl.unit = u.id
JOIN variants v ON pl.variant_code = v.code
JOIN accounts a ON a.id = pl.user_created
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
WHERE pl.id = $1::int
`

type DetailPriceListRow struct {
	ID                int32           `json:"id"`
	VariantCode       string          `json:"variant_code"`
	VariantName       string          `json:"variant_name"`
	PriceImport       float64         `json:"price_import"`
	PriceSell         float64         `json:"price_sell"`
	Unit              int32           `json:"unit"`
	UserCreated       int32           `json:"user_created"`
	UserUpdated       sql.NullInt32   `json:"user_updated"`
	UpdatedAt         sql.NullTime    `json:"updated_at"`
	CreatedAt         time.Time       `json:"created_at"`
	ID_2              int32           `json:"id_2"`
	Name              string          `json:"name"`
	SellPrice         float64         `json:"sell_price"`
	ImportPrice       float64         `json:"import_price"`
	Weight            sql.NullFloat64 `json:"weight"`
	WeightUnit        sql.NullString  `json:"weight_unit"`
	UserCreated_2     int32           `json:"user_created_2"`
	UserUpdated_2     sql.NullInt32   `json:"user_updated_2"`
	UpdatedAt_2       sql.NullTime    `json:"updated_at_2"`
	CreatedAt_2       time.Time       `json:"created_at_2"`
	ID_3              int32           `json:"id_3"`
	Name_2            string          `json:"name_2"`
	Code              string          `json:"code"`
	Barcode           sql.NullString  `json:"barcode"`
	DecisionNumber    sql.NullString  `json:"decision_number"`
	RegisterNumber    sql.NullString  `json:"register_number"`
	Longevity         sql.NullString  `json:"longevity"`
	Vat               sql.NullFloat64 `json:"vat"`
	Product           int32           `json:"product"`
	UserCreated_3     int32           `json:"user_created_3"`
	UserUpdated_3     sql.NullInt32   `json:"user_updated_3"`
	UpdatedAt_3       sql.NullTime    `json:"updated_at_3"`
	CreatedAt_3       time.Time       `json:"created_at_3"`
	InitialInventory  int32           `json:"initial_inventory"`
	RealInventory     int32           `json:"real_inventory"`
	ID_4              int32           `json:"id_4"`
	Username          string          `json:"username"`
	HashedPassword    string          `json:"hashed_password"`
	FullName          string          `json:"full_name"`
	Email             string          `json:"email"`
	Type              int32           `json:"type"`
	IsVerify          bool            `json:"is_verify"`
	PasswordChangedAt time.Time       `json:"password_changed_at"`
	CreatedAt_4       time.Time       `json:"created_at_4"`
	Role              sql.NullInt32   `json:"role"`
	Gender            NullGender      `json:"gender"`
	Licence           sql.NullString  `json:"licence"`
	Dob               sql.NullTime    `json:"dob"`
	Address           sql.NullInt32   `json:"address"`
	ID_5              sql.NullInt32   `json:"id_5"`
	Variant           sql.NullInt32   `json:"variant"`
	Media             sql.NullInt32   `json:"media"`
	ID_6              sql.NullInt32   `json:"id_6"`
	MediaUrl          sql.NullString  `json:"media_url"`
	UnitName          string          `json:"unit_name"`
	VariantMedia      sql.NullString  `json:"variant_media"`
	UserCreatedName   string          `json:"user_created_name"`
}

func (q *Queries) DetailPriceList(ctx context.Context, id int32) (DetailPriceListRow, error) {
	row := q.db.QueryRowContext(ctx, detailPriceList, id)
	var i DetailPriceListRow
	err := row.Scan(
		&i.ID,
		&i.VariantCode,
		&i.VariantName,
		&i.PriceImport,
		&i.PriceSell,
		&i.Unit,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.ID_2,
		&i.Name,
		&i.SellPrice,
		&i.ImportPrice,
		&i.Weight,
		&i.WeightUnit,
		&i.UserCreated_2,
		&i.UserUpdated_2,
		&i.UpdatedAt_2,
		&i.CreatedAt_2,
		&i.ID_3,
		&i.Name_2,
		&i.Code,
		&i.Barcode,
		&i.DecisionNumber,
		&i.RegisterNumber,
		&i.Longevity,
		&i.Vat,
		&i.Product,
		&i.UserCreated_3,
		&i.UserUpdated_3,
		&i.UpdatedAt_3,
		&i.CreatedAt_3,
		&i.InitialInventory,
		&i.RealInventory,
		&i.ID_4,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt_4,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
		&i.ID_5,
		&i.Variant,
		&i.Media,
		&i.ID_6,
		&i.MediaUrl,
		&i.UnitName,
		&i.VariantMedia,
		&i.UserCreatedName,
	)
	return i, err
}

const getPriceLists = `-- name: GetPriceLists :many
SELECT pl.id, variant_code, variant_name, price_import, price_sell, pl.unit, pl.user_created, pl.user_updated, pl.updated_at, pl.created_at, u.id, u.name, sell_price, import_price, weight, weight_unit, u.user_created, u.user_updated, u.updated_at, u.created_at, v.id, v.name, v.code, barcode, decision_number, register_number, longevity, vat, product, v.user_created, v.user_updated, v.updated_at, v.created_at, initial_inventory, real_inventory, a.id, username, hashed_password, full_name, email, a.type, is_verify, password_changed_at, a.created_at, role, gender, licence, dob, address, vm.id, variant, media, m.id, media_url, p.id, p.name, p.code, product_category, p.type, brand, p.unit, ta_duoc, nong_do, lieu_dung, chi_dinh, chong_chi_dinh, cong_dung, tac_dung_phu, than_trong, tuong_tac, bao_quan, dong_goi, phan_loai, dang_bao_che, tieu_chuan_sx, cong_ty_sx, cong_ty_dk, active, company, p.user_created, p.user_updated, p.updated_at, p.created_at, u.name AS unit_name, p.company AS company, m.media_url AS variant_media,
       a.full_name AS user_created_name FROM price_list pl
JOIN units u ON pl.unit = u.id
JOIN variants v ON pl.variant_code = v.code
JOIN accounts a ON a.id = pl.user_created
LEFT JOIN variant_media vm ON vm.variant = v.id
LEFT JOIN medias m ON m.id = vm.media
JOIN products p ON p.id = v.product
WHERE p.company = $1::int
AND (
    variant_name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    variant_code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
AND (
    ($3::float IS NULL AND $4::float IS NULL)
    OR (pl.price_import BETWEEN $3::float AND $4::float)
    OR ($5::float IS NULL AND $6::float IS NULL)
    OR (pl.price_sell BETWEEN $5::float AND $6::float)
)
ORDER BY -pl.id
LIMIT COALESCE($8::int, 10)
OFFSET (COALESCE($7::int, 1) - 1) * COALESCE($8::int, 10)
`

type GetPriceListsParams struct {
	Company        sql.NullInt32   `json:"company"`
	Search         sql.NullString  `json:"search"`
	MinPriceImport sql.NullFloat64 `json:"min_price_import"`
	MaxPriceImport sql.NullFloat64 `json:"max_price_import"`
	MinPriceSell   sql.NullFloat64 `json:"min_price_sell"`
	MaxPriceSell   sql.NullFloat64 `json:"max_price_sell"`
	Page           sql.NullInt32   `json:"page"`
	Limit          sql.NullInt32   `json:"limit"`
}

type GetPriceListsRow struct {
	ID                int32           `json:"id"`
	VariantCode       string          `json:"variant_code"`
	VariantName       string          `json:"variant_name"`
	PriceImport       float64         `json:"price_import"`
	PriceSell         float64         `json:"price_sell"`
	Unit              int32           `json:"unit"`
	UserCreated       int32           `json:"user_created"`
	UserUpdated       sql.NullInt32   `json:"user_updated"`
	UpdatedAt         sql.NullTime    `json:"updated_at"`
	CreatedAt         time.Time       `json:"created_at"`
	ID_2              int32           `json:"id_2"`
	Name              string          `json:"name"`
	SellPrice         float64         `json:"sell_price"`
	ImportPrice       float64         `json:"import_price"`
	Weight            sql.NullFloat64 `json:"weight"`
	WeightUnit        sql.NullString  `json:"weight_unit"`
	UserCreated_2     int32           `json:"user_created_2"`
	UserUpdated_2     sql.NullInt32   `json:"user_updated_2"`
	UpdatedAt_2       sql.NullTime    `json:"updated_at_2"`
	CreatedAt_2       time.Time       `json:"created_at_2"`
	ID_3              int32           `json:"id_3"`
	Name_2            string          `json:"name_2"`
	Code              string          `json:"code"`
	Barcode           sql.NullString  `json:"barcode"`
	DecisionNumber    sql.NullString  `json:"decision_number"`
	RegisterNumber    sql.NullString  `json:"register_number"`
	Longevity         sql.NullString  `json:"longevity"`
	Vat               sql.NullFloat64 `json:"vat"`
	Product           int32           `json:"product"`
	UserCreated_3     int32           `json:"user_created_3"`
	UserUpdated_3     sql.NullInt32   `json:"user_updated_3"`
	UpdatedAt_3       sql.NullTime    `json:"updated_at_3"`
	CreatedAt_3       time.Time       `json:"created_at_3"`
	InitialInventory  int32           `json:"initial_inventory"`
	RealInventory     int32           `json:"real_inventory"`
	ID_4              int32           `json:"id_4"`
	Username          string          `json:"username"`
	HashedPassword    string          `json:"hashed_password"`
	FullName          string          `json:"full_name"`
	Email             string          `json:"email"`
	Type              int32           `json:"type"`
	IsVerify          bool            `json:"is_verify"`
	PasswordChangedAt time.Time       `json:"password_changed_at"`
	CreatedAt_4       time.Time       `json:"created_at_4"`
	Role              sql.NullInt32   `json:"role"`
	Gender            NullGender      `json:"gender"`
	Licence           sql.NullString  `json:"licence"`
	Dob               sql.NullTime    `json:"dob"`
	Address           sql.NullInt32   `json:"address"`
	ID_5              sql.NullInt32   `json:"id_5"`
	Variant           sql.NullInt32   `json:"variant"`
	Media             sql.NullInt32   `json:"media"`
	ID_6              sql.NullInt32   `json:"id_6"`
	MediaUrl          sql.NullString  `json:"media_url"`
	ID_7              int32           `json:"id_7"`
	Name_3            string          `json:"name_3"`
	Code_2            string          `json:"code_2"`
	ProductCategory   sql.NullInt32   `json:"product_category"`
	Type_2            sql.NullInt32   `json:"type_2"`
	Brand             sql.NullInt32   `json:"brand"`
	Unit_2            int32           `json:"unit_2"`
	TaDuoc            sql.NullString  `json:"ta_duoc"`
	NongDo            sql.NullString  `json:"nong_do"`
	LieuDung          sql.NullString  `json:"lieu_dung"`
	ChiDinh           sql.NullString  `json:"chi_dinh"`
	ChongChiDinh      sql.NullString  `json:"chong_chi_dinh"`
	CongDung          sql.NullString  `json:"cong_dung"`
	TacDungPhu        sql.NullString  `json:"tac_dung_phu"`
	ThanTrong         sql.NullString  `json:"than_trong"`
	TuongTac          sql.NullString  `json:"tuong_tac"`
	BaoQuan           sql.NullString  `json:"bao_quan"`
	DongGoi           sql.NullString  `json:"dong_goi"`
	PhanLoai          sql.NullString  `json:"phan_loai"`
	DangBaoChe        sql.NullString  `json:"dang_bao_che"`
	TieuChuanSx       sql.NullString  `json:"tieu_chuan_sx"`
	CongTySx          sql.NullInt32   `json:"cong_ty_sx"`
	CongTyDk          sql.NullInt32   `json:"cong_ty_dk"`
	Active            bool            `json:"active"`
	Company           int32           `json:"company"`
	UserCreated_4     int32           `json:"user_created_4"`
	UserUpdated_4     sql.NullInt32   `json:"user_updated_4"`
	UpdatedAt_4       sql.NullTime    `json:"updated_at_4"`
	CreatedAt_5       time.Time       `json:"created_at_5"`
	UnitName          string          `json:"unit_name"`
	Company_2         int32           `json:"company_2"`
	VariantMedia      sql.NullString  `json:"variant_media"`
	UserCreatedName   string          `json:"user_created_name"`
}

func (q *Queries) GetPriceLists(ctx context.Context, arg GetPriceListsParams) ([]GetPriceListsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPriceLists,
		arg.Company,
		arg.Search,
		arg.MinPriceImport,
		arg.MaxPriceImport,
		arg.MinPriceSell,
		arg.MaxPriceSell,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPriceListsRow{}
	for rows.Next() {
		var i GetPriceListsRow
		if err := rows.Scan(
			&i.ID,
			&i.VariantCode,
			&i.VariantName,
			&i.PriceImport,
			&i.PriceSell,
			&i.Unit,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.ID_2,
			&i.Name,
			&i.SellPrice,
			&i.ImportPrice,
			&i.Weight,
			&i.WeightUnit,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
			&i.ID_3,
			&i.Name_2,
			&i.Code,
			&i.Barcode,
			&i.DecisionNumber,
			&i.RegisterNumber,
			&i.Longevity,
			&i.Vat,
			&i.Product,
			&i.UserCreated_3,
			&i.UserUpdated_3,
			&i.UpdatedAt_3,
			&i.CreatedAt_3,
			&i.InitialInventory,
			&i.RealInventory,
			&i.ID_4,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.Type,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt_4,
			&i.Role,
			&i.Gender,
			&i.Licence,
			&i.Dob,
			&i.Address,
			&i.ID_5,
			&i.Variant,
			&i.Media,
			&i.ID_6,
			&i.MediaUrl,
			&i.ID_7,
			&i.Name_3,
			&i.Code_2,
			&i.ProductCategory,
			&i.Type_2,
			&i.Brand,
			&i.Unit_2,
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
			&i.UserCreated_4,
			&i.UserUpdated_4,
			&i.UpdatedAt_4,
			&i.CreatedAt_5,
			&i.UnitName,
			&i.Company_2,
			&i.VariantMedia,
			&i.UserCreatedName,
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

const updatePriceList = `-- name: UpdatePriceList :one
UPDATE price_list
SET
    price_import = $1::float,
    price_sell = $2:: float
WHERE id = $3
RETURNING id, variant_code, variant_name, price_import, price_sell, unit, user_created, user_updated, updated_at, created_at
`

type UpdatePriceListParams struct {
	PriceImport float64 `json:"price_import"`
	PriceSell   float64 `json:"price_sell"`
	ID          int32   `json:"id"`
}

func (q *Queries) UpdatePriceList(ctx context.Context, arg UpdatePriceListParams) (PriceList, error) {
	row := q.db.QueryRowContext(ctx, updatePriceList, arg.PriceImport, arg.PriceSell, arg.ID)
	var i PriceList
	err := row.Scan(
		&i.ID,
		&i.VariantCode,
		&i.VariantName,
		&i.PriceImport,
		&i.PriceSell,
		&i.Unit,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
