// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: category.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO product_categories (
    code, name, description, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, code, name, user_created, created_at, company, user_updated, updated_at, description
`

type CreateCategoryParams struct {
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Company     int32          `json:"company"`
	UserCreated int32          `json:"user_created"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.Code,
		arg.Name,
		arg.Description,
		arg.Company,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.UserCreated,
		&i.CreatedAt,
		&i.Company,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :one
DELETE FROM product_categories
WHERE id = $1 RETURNING id, code, name, user_created, created_at, company, user_updated, updated_at, description
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, deleteCategory, id)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.UserCreated,
		&i.CreatedAt,
		&i.Company,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const detailCategory = `-- name: DetailCategory :one
SELECT pb.id, code, name, user_created, pb.created_at, company, user_updated, updated_at, description, ac.id, ac.username, ac.hashed_password, ac.full_name, ac.email, ac.type, ac.is_verify, ac.password_changed_at, ac.created_at, ac.role, ac.gender, ac.licence, ac.dob, ac.address, au.id, au.username, au.hashed_password, au.full_name, au.email, au.type, au.is_verify, au.password_changed_at, au.created_at, au.role, au.gender, au.licence, au.dob, au.address, ac.full_name AS created_name, au.full_name AS updated_name FROM product_categories pb
JOIN accounts ac ON ac.id = pb.user_created
LEFT JOIN accounts au ON au.id = pb.user_updated
WHERE pb.id = $1
`

type DetailCategoryRow struct {
	ID                  int32          `json:"id"`
	Code                string         `json:"code"`
	Name                string         `json:"name"`
	UserCreated         int32          `json:"user_created"`
	CreatedAt           time.Time      `json:"created_at"`
	Company             int32          `json:"company"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	Description         sql.NullString `json:"description"`
	ID_2                int32          `json:"id_2"`
	Username            string         `json:"username"`
	HashedPassword      string         `json:"hashed_password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	Type                int32          `json:"type"`
	IsVerify            bool           `json:"is_verify"`
	PasswordChangedAt   time.Time      `json:"password_changed_at"`
	CreatedAt_2         time.Time      `json:"created_at_2"`
	Role                sql.NullInt32  `json:"role"`
	Gender              NullGender     `json:"gender"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address             sql.NullInt32  `json:"address"`
	ID_3                sql.NullInt32  `json:"id_3"`
	Username_2          sql.NullString `json:"username_2"`
	HashedPassword_2    sql.NullString `json:"hashed_password_2"`
	FullName_2          sql.NullString `json:"full_name_2"`
	Email_2             sql.NullString `json:"email_2"`
	Type_2              sql.NullInt32  `json:"type_2"`
	IsVerify_2          sql.NullBool   `json:"is_verify_2"`
	PasswordChangedAt_2 sql.NullTime   `json:"password_changed_at_2"`
	CreatedAt_3         sql.NullTime   `json:"created_at_3"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_2           sql.NullInt32  `json:"address_2"`
	CreatedName         string         `json:"created_name"`
	UpdatedName         sql.NullString `json:"updated_name"`
}

func (q *Queries) DetailCategory(ctx context.Context, id int32) (DetailCategoryRow, error) {
	row := q.db.QueryRowContext(ctx, detailCategory, id)
	var i DetailCategoryRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.UserCreated,
		&i.CreatedAt,
		&i.Company,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.Description,
		&i.ID_2,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt_2,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address,
		&i.ID_3,
		&i.Username_2,
		&i.HashedPassword_2,
		&i.FullName_2,
		&i.Email_2,
		&i.Type_2,
		&i.IsVerify_2,
		&i.PasswordChangedAt_2,
		&i.CreatedAt_3,
		&i.Role_2,
		&i.Gender_2,
		&i.Licence_2,
		&i.Dob_2,
		&i.Address_2,
		&i.CreatedName,
		&i.UpdatedName,
	)
	return i, err
}

const getListCategory = `-- name: GetListCategory :many
SELECT id, code, name, user_created, created_at, company, user_updated, updated_at, description FROM product_categories
WHERE company = $1::int
AND (
    name ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    code ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -id
    LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type GetListCategoryParams struct {
	Company sql.NullInt32  `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

func (q *Queries) GetListCategory(ctx context.Context, arg GetListCategoryParams) ([]ProductCategory, error) {
	rows, err := q.db.QueryContext(ctx, getListCategory,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductCategory{}
	for rows.Next() {
		var i ProductCategory
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.UserCreated,
			&i.CreatedAt,
			&i.Company,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.Description,
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

const updateCategory = `-- name: UpdateCategory :one
UPDATE product_categories
SET
    code = $1,
    name = $2,
    description = $3,
    user_updated = $4
WHERE id = $5
    RETURNING id, code, name, user_created, created_at, company, user_updated, updated_at, description
`

type UpdateCategoryParams struct {
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, updateCategory,
		arg.Code,
		arg.Name,
		arg.Description,
		arg.UserUpdated,
		arg.ID,
	)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.UserCreated,
		&i.CreatedAt,
		&i.Company,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}
