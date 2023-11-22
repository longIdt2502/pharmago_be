// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID                int64          `json:"id"`
	Username          sql.NullString `json:"username"`
	HashedPassword    string         `json:"hashed_password"`
	FullName          string         `json:"full_name"`
	Email             string         `json:"email"`
	Type              int64          `json:"type"`
	Media             sql.NullInt64  `json:"media"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
}

type AccountType struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Address struct {
	ID          int64     `json:"id"`
	Lat         string    `json:"lat"`
	Lng         string    `json:"lng"`
	Title       string    `json:"title"`
	UserCreated int64     `json:"user_created"`
	CreatedAt   time.Time `json:"created_at"`
}

type Company struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	TaxCode     sql.NullString `json:"tax_code"`
	Phone       sql.NullString `json:"phone"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	Owner       sql.NullInt64  `json:"owner"`
}

type Customer struct {
	ID          int64          `json:"id"`
	FullName    string         `json:"full_name"`
	Code        string         `json:"code"`
	Company     sql.NullInt64  `json:"company"`
	Address     sql.NullInt64  `json:"address"`
	Email       sql.NullString `json:"email"`
	Birthday    sql.NullTime   `json:"birthday"`
	UserCreated int64          `json:"user_created"`
	UserUpdated sql.NullInt64  `json:"user_updated"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Media struct {
	ID       int64  `json:"id"`
	MediaUrl string `json:"media_url"`
}

type Order struct {
	ID          int64          `json:"id"`
	Vat         string         `json:"vat"`
	TotalPrice  string         `json:"total_price"`
	Description sql.NullString `json:"description"`
	Customer    sql.NullInt64  `json:"customer"`
	Status      int64          `json:"status"`
	Type        sql.NullInt64  `json:"type"`
	Ticket      sql.NullInt64  `json:"ticket"`
	Qr          sql.NullInt64  `json:"qr"`
}

type OrderItem struct {
	ID      int64         `json:"id"`
	Order   sql.NullInt64 `json:"order"`
	Variant sql.NullInt64 `json:"variant"`
	Value   int32         `json:"value"`
}

type OrderStatus struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type OrderType struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Product struct {
	ID              int64         `json:"id"`
	Name            string        `json:"name"`
	Code            string        `json:"code"`
	ProductCategory sql.NullInt64 `json:"product_category"`
	Type            sql.NullInt64 `json:"type"`
	Unit            int64         `json:"unit"`
	Company         sql.NullInt64 `json:"company"`
	UserCreated     int64         `json:"user_created"`
	UserUpdated     sql.NullInt64 `json:"user_updated"`
	UpdatedAt       sql.NullTime  `json:"updated_at"`
	CreatedAt       time.Time     `json:"created_at"`
}

type ProductCategory struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	UserCreated int64         `json:"user_created"`
	UserUpdated sql.NullInt64 `json:"user_updated"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	CreatedAt   time.Time     `json:"created_at"`
}

type ProductMedium struct {
	ID      int64         `json:"id"`
	Product sql.NullInt64 `json:"product"`
	Media   sql.NullInt64 `json:"media"`
}

type ProductType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Ticket struct {
	ID          int64          `json:"id"`
	Code        string         `json:"code"`
	Type        int64          `json:"type"`
	Status      int64          `json:"status"`
	Note        sql.NullString `json:"note"`
	Qr          sql.NullInt64  `json:"qr"`
	ExportFrom  int64          `json:"export_from"`
	ImportTo    int64          `json:"import_to"`
	UserCreated int64          `json:"user_created"`
	UserUpdated sql.NullInt64  `json:"user_updated"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

type TicketStatus struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type TicketType struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Unit struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	SellPrice   string         `json:"sell_price"`
	ImportPrice string         `json:"import_price"`
	Weight      sql.NullString `json:"weight"`
	WeightUnit  sql.NullString `json:"weight_unit"`
	UserCreated int64          `json:"user_created"`
	UserUpdated sql.NullInt64  `json:"user_updated"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	CreatedAt   time.Time      `json:"created_at"`
}

type UnitChange struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Value       int64         `json:"value"`
	SellPrice   string        `json:"sell_price"`
	Unit        sql.NullInt64 `json:"unit"`
	UserCreated int64         `json:"user_created"`
	UserUpdated sql.NullInt64 `json:"user_updated"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
	CreatedAt   time.Time     `json:"created_at"`
}

type Variant struct {
	ID             int64         `json:"id"`
	Name           string        `json:"name"`
	Code           string        `json:"code"`
	Barcode        string        `json:"barcode"`
	DecisionNumber int64         `json:"decision_number"`
	RegisterNumber int64         `json:"register_number"`
	Discount       string        `json:"discount"`
	Vat            string        `json:"vat"`
	Product        sql.NullInt64 `json:"product"`
	UserCreated    int64         `json:"user_created"`
	UserUpdated    sql.NullInt64 `json:"user_updated"`
	UpdatedAt      sql.NullTime  `json:"updated_at"`
	CreatedAt      time.Time     `json:"created_at"`
}

type Warehouse struct {
	ID        int64         `json:"id"`
	Address   sql.NullInt64 `json:"address"`
	Companies sql.NullInt64 `json:"companies"`
}