// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                int32     `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Type              int32     `json:"type"`
	IsVerify          bool      `json:"is_verify"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type AccountCompany struct {
	ID      int32 `json:"id"`
	Account int32 `json:"account"`
	Company int32 `json:"company"`
}

type AccountMedium struct {
	ID      int32 `json:"id"`
	Account int32 `json:"account"`
	Media   int32 `json:"media"`
}

type AccountType struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Address struct {
	ID          int32          `json:"id"`
	Lat         float64        `json:"lat"`
	Lng         float64        `json:"lng"`
	Province    sql.NullString `json:"province"`
	District    sql.NullString `json:"district"`
	Ward        sql.NullString `json:"ward"`
	Title       string         `json:"title"`
	UserCreated int32          `json:"user_created"`
	CreatedAt   time.Time      `json:"created_at"`
}

type AdministrativeRegion struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	NameEn     string `json:"name_en"`
	CodeName   string `json:"code_name"`
	CodeNameEn string `json:"code_name_en"`
}

type AdministrativeUnit struct {
	ID          int32  `json:"id"`
	FullName    string `json:"full_name"`
	FullNameEn  string `json:"full_name_en"`
	ShortName   string `json:"short_name"`
	ShortNameEn string `json:"short_name_en"`
	CodeName    string `json:"code_name"`
	CodeNameEn  string `json:"code_name_en"`
}

type Classify struct {
	ID   int32  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Company struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Code        string         `json:"code"`
	TaxCode     sql.NullString `json:"tax_code"`
	Phone       sql.NullString `json:"phone"`
	Description sql.NullString `json:"description"`
	Address     sql.NullInt32  `json:"address"`
	CreatedAt   time.Time      `json:"created_at"`
	Owner       int32          `json:"owner"`
}

type CompanyPharma struct {
	ID                int32          `json:"id"`
	Name              string         `json:"name"`
	Code              sql.NullString `json:"code"`
	Country           sql.NullString `json:"country"`
	Address           sql.NullString `json:"address"`
	CompanyPharmaType sql.NullString `json:"company_pharma_type"`
	CreatedAt         time.Time      `json:"created_at"`
}

type CompanyPharmaType struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Consignment struct {
	ID          int32         `json:"id"`
	Code        string        `json:"code"`
	Quantity    int32         `json:"quantity"`
	Inventory   int32         `json:"inventory"`
	Ticket      sql.NullInt32 `json:"ticket"`
	ExpiredAt   time.Time     `json:"expired_at"`
	ProductedAt time.Time     `json:"producted_at"`
	IsAvailable bool          `json:"is_available"`
	UserCreated sql.NullInt32 `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	CreatedAt   time.Time     `json:"created_at"`
	Variant     sql.NullInt32 `json:"variant"`
}

type ConsignmentLog struct {
	ID           int32         `json:"id"`
	Consignment  int32         `json:"consignment"`
	Inventory    int32         `json:"inventory"`
	AmountChange int32         `json:"amount_change"`
	UserCreated  sql.NullInt32 `json:"user_created"`
	CreatedAt    time.Time     `json:"created_at"`
}

type Customer struct {
	ID          int32          `json:"id"`
	FullName    string         `json:"full_name"`
	Code        string         `json:"code"`
	Company     int32          `json:"company"`
	Address     sql.NullInt32  `json:"address"`
	Email       sql.NullString `json:"email"`
	Phone       sql.NullString `json:"phone"`
	License     sql.NullString `json:"license"`
	Birthday    sql.NullTime   `json:"birthday"`
	UserCreated int32          `json:"user_created"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

type District struct {
	Code                 string         `json:"code"`
	Name                 string         `json:"name"`
	NameEn               string         `json:"name_en"`
	FullName             string         `json:"full_name"`
	FullNameEn           string         `json:"full_name_en"`
	CodeName             string         `json:"code_name"`
	ProvinceCode         sql.NullString `json:"province_code"`
	AdministrativeUnitID sql.NullInt32  `json:"administrative_unit_id"`
}

type Ingredient struct {
	ID      int32   `json:"id"`
	Name    string  `json:"name"`
	Weight  float64 `json:"weight"`
	Unit    string  `json:"unit"`
	Product int32   `json:"product"`
}

type Media struct {
	ID       int32  `json:"id"`
	MediaUrl string `json:"media_url"`
}

type Order struct {
	ID           int32          `json:"id"`
	Code         string         `json:"code"`
	TotalPrice   float64        `json:"total_price"`
	Description  sql.NullString `json:"description"`
	Vat          float64        `json:"vat"`
	Discount     string         `json:"discount"`
	ServicePrice float64        `json:"service_price"`
	MustPaid     float64        `json:"must_paid"`
	Customer     sql.NullInt32  `json:"customer"`
	Address      sql.NullInt32  `json:"address"`
	Status       sql.NullString `json:"status"`
	Type         sql.NullString `json:"type"`
	Ticket       sql.NullInt32  `json:"ticket"`
	Qr           sql.NullInt32  `json:"qr"`
	Company      int32          `json:"company"`
	Payment      int32          `json:"payment"`
	UserCreated  sql.NullInt32  `json:"user_created"`
	UserUpdated  sql.NullInt32  `json:"user_updated"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

type OrderItem struct {
	ID             int32         `json:"id"`
	Order          int32         `json:"order"`
	Variant        int32         `json:"variant"`
	Value          int32         `json:"value"`
	TotalPrice     float64       `json:"total_price"`
	Consignment    sql.NullInt32 `json:"consignment"`
	ConsignmentLog sql.NullInt32 `json:"consignment_log"`
}

type OrderStatus struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type OrderType struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Payment struct {
	ID       int32   `json:"id"`
	Code     string  `json:"code"`
	MustPaid float64 `json:"must_paid"`
	HadPaid  float64 `json:"had_paid"`
	NeedPay  float64 `json:"need_pay"`
}

type PaymentItem struct {
	ID        int32          `json:"id"`
	Type      string         `json:"type"`
	Value     float64        `json:"value"`
	IsPaid    bool           `json:"is_paid"`
	Payment   int32          `json:"payment"`
	ExtraNote sql.NullString `json:"extra_note"`
}

type PaymentItemType struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type PreparationType struct {
	ID   int32  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type PriceList struct {
	ID          int32         `json:"id"`
	VariantCode string        `json:"variant_code"`
	VariantName string        `json:"variant_name"`
	PriceImport float64       `json:"price_import"`
	PriceSell   float64       `json:"price_sell"`
	Unit        int32         `json:"unit"`
	UserCreated int32         `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	CreatedAt   time.Time     `json:"created_at"`
}

type PriceListLog struct {
	ID             int32         `json:"id"`
	OldPriceImport float64       `json:"old_price_import"`
	NewPriceImport float64       `json:"new_price_import"`
	OldPriceSell   float64       `json:"old_price_sell"`
	NewPriceSell   float64       `json:"new_price_sell"`
	PriceList      int32         `json:"price_list"`
	UserUpdated    sql.NullInt32 `json:"user_updated"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type Product struct {
	ID              int32          `json:"id"`
	Name            string         `json:"name"`
	Code            string         `json:"code"`
	ProductCategory sql.NullInt32  `json:"product_category"`
	Type            sql.NullInt32  `json:"type"`
	Brand           sql.NullInt32  `json:"brand"`
	Unit            int32          `json:"unit"`
	TaDuoc          sql.NullString `json:"ta_duoc"`
	NongDo          sql.NullString `json:"nong_do"`
	LieuDung        string         `json:"lieu_dung"`
	ChiDinh         string         `json:"chi_dinh"`
	ChongChiDinh    sql.NullString `json:"chong_chi_dinh"`
	CongDung        string         `json:"cong_dung"`
	TacDungPhu      string         `json:"tac_dung_phu"`
	ThanTrong       string         `json:"than_trong"`
	TuongTac        sql.NullString `json:"tuong_tac"`
	BaoQuan         string         `json:"bao_quan"`
	DongGoi         string         `json:"dong_goi"`
	PhanLoai        sql.NullString `json:"phan_loai"`
	DangBaoChe      string         `json:"dang_bao_che"`
	TieuChuanSx     string         `json:"tieu_chuan_sx"`
	CongTySx        int32          `json:"cong_ty_sx"`
	CongTyDk        int32          `json:"cong_ty_dk"`
	Active          bool           `json:"active"`
	Company         int32          `json:"company"`
	UserCreated     int32          `json:"user_created"`
	UserUpdated     sql.NullInt32  `json:"user_updated"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
	CreatedAt       time.Time      `json:"created_at"`
}

type ProductBrand struct {
	ID          int32     `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	UserCreated int32     `json:"user_created"`
	CreatedAt   time.Time `json:"created_at"`
	Company     int32     `json:"company"`
}

type ProductCategory struct {
	ID          int32     `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	UserCreated int32     `json:"user_created"`
	CreatedAt   time.Time `json:"created_at"`
	Company     int32     `json:"company"`
}

type ProductMedium struct {
	ID      int32 `json:"id"`
	Product int32 `json:"product"`
	Media   int32 `json:"media"`
}

type ProductType struct {
	ID          int32     `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	UserCreated int32     `json:"user_created"`
	CreatedAt   time.Time `json:"created_at"`
	Company     int32     `json:"company"`
}

type ProductionStandard struct {
	ID   int32  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProductsBank struct {
	ID           int32          `json:"id"`
	Name         string         `json:"name"`
	Code         string         `json:"code"`
	TaDuoc       sql.NullString `json:"ta_duoc"`
	NongDo       sql.NullString `json:"nong_do"`
	LieuDung     string         `json:"lieu_dung"`
	ChiDinh      string         `json:"chi_dinh"`
	ChongChiDinh sql.NullString `json:"chong_chi_dinh"`
	CongDung     string         `json:"cong_dung"`
	TacDungPhu   string         `json:"tac_dung_phu"`
	ThanTrong    string         `json:"than_trong"`
	TuongTac     sql.NullString `json:"tuong_tac"`
	BaoQuan      string         `json:"bao_quan"`
	DongGoi      string         `json:"dong_goi"`
	PhanLoai     sql.NullString `json:"phan_loai"`
	DangBaoChe   string         `json:"dang_bao_che"`
	TieuChuanSx  string         `json:"tieu_chuan_sx"`
	CongTySx     int32          `json:"cong_ty_sx"`
	CongTyDk     int32          `json:"cong_ty_dk"`
}

type Province struct {
	Code                   string        `json:"code"`
	Name                   string        `json:"name"`
	NameEn                 string        `json:"name_en"`
	FullName               string        `json:"full_name"`
	FullNameEn             string        `json:"full_name_en"`
	CodeName               string        `json:"code_name"`
	AdministrativeUnitID   sql.NullInt32 `json:"administrative_unit_id"`
	AdministrativeRegionID sql.NullInt32 `json:"administrative_region_id"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Suplier struct {
	ID         int32          `json:"id"`
	Code       string         `json:"code"`
	Name       string         `json:"name"`
	DeputyName string         `json:"deputy_name"`
	Phone      string         `json:"phone"`
	Email      sql.NullString `json:"email"`
	Address    sql.NullInt32  `json:"address"`
	Company    sql.NullInt32  `json:"company"`
}

type Ticket struct {
	ID          int32          `json:"id"`
	Code        string         `json:"code"`
	Type        sql.NullInt32  `json:"type"`
	Status      sql.NullInt32  `json:"status"`
	Note        sql.NullString `json:"note"`
	Qr          sql.NullInt32  `json:"qr"`
	ExportTo    sql.NullInt32  `json:"export_to"`
	ImportFrom  sql.NullInt32  `json:"import_from"`
	TotalPrice  float64        `json:"total_price"`
	Warehouse   int32          `json:"warehouse"`
	UserCreated int32          `json:"user_created"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

type TicketStatus struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type TicketType struct {
	ID    int32  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Unit struct {
	ID          int32           `json:"id"`
	Name        string          `json:"name"`
	SellPrice   float64         `json:"sell_price"`
	ImportPrice float64         `json:"import_price"`
	Weight      sql.NullFloat64 `json:"weight"`
	WeightUnit  sql.NullString  `json:"weight_unit"`
	UserCreated int32           `json:"user_created"`
	UserUpdated sql.NullInt32   `json:"user_updated"`
	UpdatedAt   sql.NullTime    `json:"updated_at"`
	CreatedAt   time.Time       `json:"created_at"`
}

type UnitChange struct {
	ID          int32         `json:"id"`
	Name        string        `json:"name"`
	Value       int64         `json:"value"`
	SellPrice   float64       `json:"sell_price"`
	Unit        int32         `json:"unit"`
	UserCreated int32         `json:"user_created"`
	UserUpdated sql.NullInt32 `json:"user_updated"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	CreatedAt   time.Time     `json:"created_at"`
}

type Variant struct {
	ID             int32         `json:"id"`
	Name           string        `json:"name"`
	Code           string        `json:"code"`
	Barcode        string        `json:"barcode"`
	DecisionNumber string        `json:"decision_number"`
	RegisterNumber string        `json:"register_number"`
	Longevity      string        `json:"longevity"`
	Vat            float64       `json:"vat"`
	Product        int32         `json:"product"`
	UserCreated    int32         `json:"user_created"`
	UserUpdated    sql.NullInt32 `json:"user_updated"`
	UpdatedAt      sql.NullTime  `json:"updated_at"`
	CreatedAt      time.Time     `json:"created_at"`
}

type VariantMedium struct {
	ID      int32 `json:"id"`
	Variant int32 `json:"variant"`
	Media   int32 `json:"media"`
}

type Verify struct {
	ID         int32     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

type Ward struct {
	Code                 string         `json:"code"`
	Name                 string         `json:"name"`
	NameEn               string         `json:"name_en"`
	FullName             string         `json:"full_name"`
	FullNameEn           string         `json:"full_name_en"`
	CodeName             string         `json:"code_name"`
	DistrictCode         sql.NullString `json:"district_code"`
	AdministrativeUnitID sql.NullInt32  `json:"administrative_unit_id"`
}

type Warehouse struct {
	ID        int32         `json:"id"`
	Address   sql.NullInt32 `json:"address"`
	Companies sql.NullInt32 `json:"companies"`
	Name      string        `json:"name"`
	Code      string        `json:"code"`
}
