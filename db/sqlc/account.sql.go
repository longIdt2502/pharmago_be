// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (username, hashed_password, full_name, email, type, role, gender, licence, dob, address, is_verify)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address
`

type CreateAccountParams struct {
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	FullName       string         `json:"full_name"`
	Email          string         `json:"email"`
	Type           int32          `json:"type"`
	Role           sql.NullInt32  `json:"role"`
	Gender         NullGender     `json:"gender"`
	Licence        sql.NullString `json:"licence"`
	Dob            sql.NullTime   `json:"dob"`
	Address        sql.NullInt32  `json:"address"`
	IsVerify       bool           `json:"is_verify"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.Type,
		arg.Role,
		arg.Gender,
		arg.Licence,
		arg.Dob,
		arg.Address,
		arg.IsVerify,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}

const createAccountCompany = `-- name: CreateAccountCompany :one
INSERT INTO account_company (account, company) VALUES ($1, $2) RETURNING id, account, company
`

type CreateAccountCompanyParams struct {
	Account int32 `json:"account"`
	Company int32 `json:"company"`
}

func (q *Queries) CreateAccountCompany(ctx context.Context, arg CreateAccountCompanyParams) (AccountCompany, error) {
	row := q.db.QueryRowContext(ctx, createAccountCompany, arg.Account, arg.Company)
	var i AccountCompany
	err := row.Scan(&i.ID, &i.Account, &i.Company)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}

const getAccountByMail = `-- name: GetAccountByMail :one
SELECT id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address FROM accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetAccountByMail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByMail, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}

const getAccountByPhone = `-- name: GetAccountByPhone :one
SELECT id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address FROM accounts
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetAccountByPhone(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByPhone, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}

const getAccountByUseName = `-- name: GetAccountByUseName :one
SELECT id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address FROM accounts
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetAccountByUseName(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUseName, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}

const listAccount = `-- name: ListAccount :many
SELECT a.id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address, ac.id, account, company FROM accounts a
LEFT JOIN account_company ac ON ac.account = a.id
WHERE ac.company = $1::int
AND (
    a.full_name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    a.username ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
AND (
    $3::int IS NULL OR a.type = $3::int
)
AND (
    $4::int IS NULL OR a.role = $4::int
    
)
ORDER BY -a.id
LIMIT COALESCE($6::int, 10)
OFFSET (COALESCE($5::int, 1) - 1) * COALESCE($6::int, 10)
`

type ListAccountParams struct {
	Company int32          `json:"company"`
	Search  sql.NullString `json:"search"`
	Type    sql.NullInt32  `json:"type"`
	Role    sql.NullInt32  `json:"role"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

type ListAccountRow struct {
	ID                int32          `json:"id"`
	Username          string         `json:"username"`
	HashedPassword    string         `json:"hashed_password"`
	FullName          string         `json:"full_name"`
	Email             string         `json:"email"`
	Type              int32          `json:"type"`
	IsVerify          bool           `json:"is_verify"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
	Role              sql.NullInt32  `json:"role"`
	Gender            NullGender     `json:"gender"`
	Licence           sql.NullString `json:"licence"`
	Dob               sql.NullTime   `json:"dob"`
	Address           sql.NullInt32  `json:"address"`
	ID_2              sql.NullInt32  `json:"id_2"`
	Account           sql.NullInt32  `json:"account"`
	Company           sql.NullInt32  `json:"company"`
}

func (q *Queries) ListAccount(ctx context.Context, arg ListAccountParams) ([]ListAccountRow, error) {
	rows, err := q.db.QueryContext(ctx, listAccount,
		arg.Company,
		arg.Search,
		arg.Type,
		arg.Role,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAccountRow{}
	for rows.Next() {
		var i ListAccountRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.Type,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.Role,
			&i.Gender,
			&i.Licence,
			&i.Dob,
			&i.Address,
			&i.ID_2,
			&i.Account,
			&i.Company,
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

const resetPassword = `-- name: ResetPassword :one
UPDATE accounts
SET
    hashed_password = COALESCE($1, hashed_password)
WHERE
    email = $2
RETURNING id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address
`

type ResetPasswordParams struct {
	HashedPassword sql.NullString `json:"hashed_password"`
	Email          sql.NullString `json:"email"`
}

func (q *Queries) ResetPassword(ctx context.Context, arg ResetPasswordParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, resetPassword, arg.HashedPassword, arg.Email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET
    is_verify = COALESCE($1, is_verify),
    hashed_password = COALESCE($2::varchar, hashed_password),
    full_name = COALESCE($3::varchar, full_name),
    email = COALESCE($4::varchar, email),
    type = COALESCE($5::int, type),
    role = COALESCE($6::int, role),
    gender = COALESCE($7::gender, gender),
    licence = COALESCE($8::varchar, licence),
    dob = COALESCE($9::timestamp, dob)
WHERE
    id = $10
    OR username = $11
RETURNING id, username, hashed_password, full_name, email, type, is_verify, password_changed_at, created_at, role, gender, licence, dob, address
`

type UpdateAccountParams struct {
	IsVerify sql.NullBool   `json:"is_verify"`
	Password sql.NullString `json:"password"`
	FullName sql.NullString `json:"full_name"`
	Email    sql.NullString `json:"email"`
	Type     sql.NullInt32  `json:"type"`
	Role     sql.NullInt32  `json:"role"`
	Gender   NullGender     `json:"gender"`
	Licence  sql.NullString `json:"licence"`
	Dob      sql.NullTime   `json:"dob"`
	ID       sql.NullInt32  `json:"id"`
	Username sql.NullString `json:"username"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.IsVerify,
		arg.Password,
		arg.FullName,
		arg.Email,
		arg.Type,
		arg.Role,
		arg.Gender,
		arg.Licence,
		arg.Dob,
		arg.ID,
		arg.Username,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
	)
	return i, err
}
