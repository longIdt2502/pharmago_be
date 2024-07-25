// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: role.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles (
    code, title, note, company, user_created, user_updated
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, code, title, note, company, user_created, user_updated, updated_at, created_at
`

type CreateRoleParams struct {
	Code        string         `json:"code"`
	Title       string         `json:"title"`
	Note        sql.NullString `json:"note"`
	Company     sql.NullInt32  `json:"company"`
	UserCreated int32          `json:"user_created"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole,
		arg.Code,
		arg.Title,
		arg.Note,
		arg.Company,
		arg.UserCreated,
		arg.UserUpdated,
	)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Note,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const createRoleItem = `-- name: CreateRoleItem :one
INSERT INTO role_item (
    roles, app, value
) VALUES (
    $1, $2, $3
) RETURNING id, roles, app, value
`

type CreateRoleItemParams struct {
	Roles int32        `json:"roles"`
	App   string       `json:"app"`
	Value sql.NullBool `json:"value"`
}

func (q *Queries) CreateRoleItem(ctx context.Context, arg CreateRoleItemParams) (RoleItem, error) {
	row := q.db.QueryRowContext(ctx, createRoleItem, arg.Roles, arg.App, arg.Value)
	var i RoleItem
	err := row.Scan(
		&i.ID,
		&i.Roles,
		&i.App,
		&i.Value,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :one
DELETE FROM roles
WHERE id = $1 RETURNING id, code, title, note, company, user_created, user_updated, updated_at, created_at
`

func (q *Queries) DeleteRole(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRowContext(ctx, deleteRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Note,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listApp = `-- name: ListApp :many
SELECT id, title, code, parent, level FROM apps
WHERE (parent = $1::varchar
OR level = $2::int)
`

type ListAppParams struct {
	Parent sql.NullString `json:"parent"`
	Level  sql.NullInt32  `json:"level"`
}

func (q *Queries) ListApp(ctx context.Context, arg ListAppParams) ([]App, error) {
	rows, err := q.db.QueryContext(ctx, listApp, arg.Parent, arg.Level)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []App{}
	for rows.Next() {
		var i App
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Code,
			&i.Parent,
			&i.Level,
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

const listRole = `-- name: ListRole :many
SELECT r.id, r.code, title, note, company, r.user_created, r.user_updated, r.updated_at, r.created_at, c.id, name, c.code, tax_code, phone, description, c.address, oa_id, c.created_at, owner, c.type, time_open, time_close, parent, is_active, manager, c.user_created, c.user_updated, c.updated_at, ac.id, ac.username, ac.hashed_password, ac.full_name, ac.email, ac.type, ac.is_verify, ac.password_changed_at, ac.created_at, ac.role, ac.gender, ac.licence, ac.dob, ac.address, au.id, au.username, au.hashed_password, au.full_name, au.email, au.type, au.is_verify, au.password_changed_at, au.created_at, au.role, au.gender, au.licence, au.dob, au.address, ac.full_name AS created_name, au.full_name AS updated_name FROM roles r
LEFT JOIN companies c ON c.id = r.company
JOIN accounts ac ON ac.id = r.user_created
LEFT JOIN accounts au ON au.id = r.user_updated
WHERE (($1::int IS NULL AND r.company IS NULL) OR r.company = $1::int)
AND (
    r.code ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    r.title ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
AND  ((
        $3::timestamp IS NULL AND $4::timestamp  IS NULL
    ) OR (
        ($3::timestamp IS NULL OR r.created_at >= $3::timestamp) AND
        ($4::timestamp IS NULL OR r.created_at <= $4::timestamp)
))
AND ((
       $5::timestamp IS NULL AND $6::timestamp  IS NULL
   ) OR (
       (r.updated_at >= $5::timestamp OR $5::timestamp  IS NULL) AND
       (r.updated_at <= $6::timestamp OR $6::timestamp  IS NULL)
))
LIMIT COALESCE($8::int, 10)
OFFSET (COALESCE($7::int, 1) - 1) * COALESCE($8::int, 1)
`

type ListRoleParams struct {
	Company      sql.NullInt32  `json:"company"`
	Search       sql.NullString `json:"search"`
	CreatedStart sql.NullTime   `json:"created_start"`
	CreatedEnd   sql.NullTime   `json:"created_end"`
	UpdatedStart sql.NullTime   `json:"updated_start"`
	UpdatedEnd   sql.NullTime   `json:"updated_end"`
	Page         sql.NullInt32  `json:"page"`
	Limit        sql.NullInt32  `json:"limit"`
}

type ListRoleRow struct {
	ID                  int32          `json:"id"`
	Code                string         `json:"code"`
	Title               string         `json:"title"`
	Note                sql.NullString `json:"note"`
	Company             sql.NullInt32  `json:"company"`
	UserCreated         int32          `json:"user_created"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	CreatedAt           time.Time      `json:"created_at"`
	ID_2                sql.NullInt32  `json:"id_2"`
	Name                sql.NullString `json:"name"`
	Code_2              sql.NullString `json:"code_2"`
	TaxCode             sql.NullString `json:"tax_code"`
	Phone               sql.NullString `json:"phone"`
	Description         sql.NullString `json:"description"`
	Address             sql.NullInt32  `json:"address"`
	OaID                sql.NullString `json:"oa_id"`
	CreatedAt_2         sql.NullTime   `json:"created_at_2"`
	Owner               sql.NullInt32  `json:"owner"`
	Type                sql.NullString `json:"type"`
	TimeOpen            sql.NullTime   `json:"time_open"`
	TimeClose           sql.NullTime   `json:"time_close"`
	Parent              sql.NullInt32  `json:"parent"`
	IsActive            sql.NullBool   `json:"is_active"`
	Manager             sql.NullInt32  `json:"manager"`
	UserCreated_2       sql.NullInt32  `json:"user_created_2"`
	UserUpdated_2       sql.NullInt32  `json:"user_updated_2"`
	UpdatedAt_2         sql.NullTime   `json:"updated_at_2"`
	ID_3                int32          `json:"id_3"`
	Username            string         `json:"username"`
	HashedPassword      string         `json:"hashed_password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	Type_2              int32          `json:"type_2"`
	IsVerify            bool           `json:"is_verify"`
	PasswordChangedAt   time.Time      `json:"password_changed_at"`
	CreatedAt_3         time.Time      `json:"created_at_3"`
	Role                sql.NullInt32  `json:"role"`
	Gender              NullGender     `json:"gender"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address_2           sql.NullInt32  `json:"address_2"`
	ID_4                sql.NullInt32  `json:"id_4"`
	Username_2          sql.NullString `json:"username_2"`
	HashedPassword_2    sql.NullString `json:"hashed_password_2"`
	FullName_2          sql.NullString `json:"full_name_2"`
	Email_2             sql.NullString `json:"email_2"`
	Type_3              sql.NullInt32  `json:"type_3"`
	IsVerify_2          sql.NullBool   `json:"is_verify_2"`
	PasswordChangedAt_2 sql.NullTime   `json:"password_changed_at_2"`
	CreatedAt_4         sql.NullTime   `json:"created_at_4"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_3           sql.NullInt32  `json:"address_3"`
	CreatedName         string         `json:"created_name"`
	UpdatedName         sql.NullString `json:"updated_name"`
}

func (q *Queries) ListRole(ctx context.Context, arg ListRoleParams) ([]ListRoleRow, error) {
	rows, err := q.db.QueryContext(ctx, listRole,
		arg.Company,
		arg.Search,
		arg.CreatedStart,
		arg.CreatedEnd,
		arg.UpdatedStart,
		arg.UpdatedEnd,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRoleRow{}
	for rows.Next() {
		var i ListRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Title,
			&i.Note,
			&i.Company,
			&i.UserCreated,
			&i.UserUpdated,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.ID_2,
			&i.Name,
			&i.Code_2,
			&i.TaxCode,
			&i.Phone,
			&i.Description,
			&i.Address,
			&i.OaID,
			&i.CreatedAt_2,
			&i.Owner,
			&i.Type,
			&i.TimeOpen,
			&i.TimeClose,
			&i.Parent,
			&i.IsActive,
			&i.Manager,
			&i.UserCreated_2,
			&i.UserUpdated_2,
			&i.UpdatedAt_2,
			&i.ID_3,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.Type_2,
			&i.IsVerify,
			&i.PasswordChangedAt,
			&i.CreatedAt_3,
			&i.Role,
			&i.Gender,
			&i.Licence,
			&i.Dob,
			&i.Address_2,
			&i.ID_4,
			&i.Username_2,
			&i.HashedPassword_2,
			&i.FullName_2,
			&i.Email_2,
			&i.Type_3,
			&i.IsVerify_2,
			&i.PasswordChangedAt_2,
			&i.CreatedAt_4,
			&i.Role_2,
			&i.Gender_2,
			&i.Licence_2,
			&i.Dob_2,
			&i.Address_3,
			&i.CreatedName,
			&i.UpdatedName,
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

const listRoleItem = `-- name: ListRoleItem :many
SELECT ri.id, roles, app, value, a.id, title, code, parent, level FROM role_item ri
JOIN apps a ON ri.app = a.code
WHERE ri.roles = $1
`

type ListRoleItemRow struct {
	ID     int32          `json:"id"`
	Roles  int32          `json:"roles"`
	App    string         `json:"app"`
	Value  sql.NullBool   `json:"value"`
	ID_2   int32          `json:"id_2"`
	Title  string         `json:"title"`
	Code   string         `json:"code"`
	Parent sql.NullString `json:"parent"`
	Level  sql.NullInt32  `json:"level"`
}

func (q *Queries) ListRoleItem(ctx context.Context, roles int32) ([]ListRoleItemRow, error) {
	rows, err := q.db.QueryContext(ctx, listRoleItem, roles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRoleItemRow{}
	for rows.Next() {
		var i ListRoleItemRow
		if err := rows.Scan(
			&i.ID,
			&i.Roles,
			&i.App,
			&i.Value,
			&i.ID_2,
			&i.Title,
			&i.Code,
			&i.Parent,
			&i.Level,
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

const roleDetail = `-- name: RoleDetail :one
SELECT r.id, r.code, title, note, company, r.user_created, r.user_updated, r.updated_at, r.created_at, c.id, name, c.code, tax_code, phone, description, c.address, oa_id, c.created_at, owner, c.type, time_open, time_close, parent, is_active, manager, c.user_created, c.user_updated, c.updated_at, ac.id, ac.username, ac.hashed_password, ac.full_name, ac.email, ac.type, ac.is_verify, ac.password_changed_at, ac.created_at, ac.role, ac.gender, ac.licence, ac.dob, ac.address, au.id, au.username, au.hashed_password, au.full_name, au.email, au.type, au.is_verify, au.password_changed_at, au.created_at, au.role, au.gender, au.licence, au.dob, au.address, ac.full_name AS created_name, au.full_name AS updated_name FROM roles r
LEFT JOIN companies c ON c.id = r.company
JOIN accounts ac ON ac.id = r.user_created
LEFT JOIN accounts au ON ac.id = r.user_updated
WHERE r.id = $1
`

type RoleDetailRow struct {
	ID                  int32          `json:"id"`
	Code                string         `json:"code"`
	Title               string         `json:"title"`
	Note                sql.NullString `json:"note"`
	Company             sql.NullInt32  `json:"company"`
	UserCreated         int32          `json:"user_created"`
	UserUpdated         sql.NullInt32  `json:"user_updated"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
	CreatedAt           time.Time      `json:"created_at"`
	ID_2                sql.NullInt32  `json:"id_2"`
	Name                sql.NullString `json:"name"`
	Code_2              sql.NullString `json:"code_2"`
	TaxCode             sql.NullString `json:"tax_code"`
	Phone               sql.NullString `json:"phone"`
	Description         sql.NullString `json:"description"`
	Address             sql.NullInt32  `json:"address"`
	OaID                sql.NullString `json:"oa_id"`
	CreatedAt_2         sql.NullTime   `json:"created_at_2"`
	Owner               sql.NullInt32  `json:"owner"`
	Type                sql.NullString `json:"type"`
	TimeOpen            sql.NullTime   `json:"time_open"`
	TimeClose           sql.NullTime   `json:"time_close"`
	Parent              sql.NullInt32  `json:"parent"`
	IsActive            sql.NullBool   `json:"is_active"`
	Manager             sql.NullInt32  `json:"manager"`
	UserCreated_2       sql.NullInt32  `json:"user_created_2"`
	UserUpdated_2       sql.NullInt32  `json:"user_updated_2"`
	UpdatedAt_2         sql.NullTime   `json:"updated_at_2"`
	ID_3                int32          `json:"id_3"`
	Username            string         `json:"username"`
	HashedPassword      string         `json:"hashed_password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	Type_2              int32          `json:"type_2"`
	IsVerify            bool           `json:"is_verify"`
	PasswordChangedAt   time.Time      `json:"password_changed_at"`
	CreatedAt_3         time.Time      `json:"created_at_3"`
	Role                sql.NullInt32  `json:"role"`
	Gender              NullGender     `json:"gender"`
	Licence             sql.NullString `json:"licence"`
	Dob                 sql.NullTime   `json:"dob"`
	Address_2           sql.NullInt32  `json:"address_2"`
	ID_4                sql.NullInt32  `json:"id_4"`
	Username_2          sql.NullString `json:"username_2"`
	HashedPassword_2    sql.NullString `json:"hashed_password_2"`
	FullName_2          sql.NullString `json:"full_name_2"`
	Email_2             sql.NullString `json:"email_2"`
	Type_3              sql.NullInt32  `json:"type_3"`
	IsVerify_2          sql.NullBool   `json:"is_verify_2"`
	PasswordChangedAt_2 sql.NullTime   `json:"password_changed_at_2"`
	CreatedAt_4         sql.NullTime   `json:"created_at_4"`
	Role_2              sql.NullInt32  `json:"role_2"`
	Gender_2            NullGender     `json:"gender_2"`
	Licence_2           sql.NullString `json:"licence_2"`
	Dob_2               sql.NullTime   `json:"dob_2"`
	Address_3           sql.NullInt32  `json:"address_3"`
	CreatedName         string         `json:"created_name"`
	UpdatedName         sql.NullString `json:"updated_name"`
}

func (q *Queries) RoleDetail(ctx context.Context, id int32) (RoleDetailRow, error) {
	row := q.db.QueryRowContext(ctx, roleDetail, id)
	var i RoleDetailRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Note,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.ID_2,
		&i.Name,
		&i.Code_2,
		&i.TaxCode,
		&i.Phone,
		&i.Description,
		&i.Address,
		&i.OaID,
		&i.CreatedAt_2,
		&i.Owner,
		&i.Type,
		&i.TimeOpen,
		&i.TimeClose,
		&i.Parent,
		&i.IsActive,
		&i.Manager,
		&i.UserCreated_2,
		&i.UserUpdated_2,
		&i.UpdatedAt_2,
		&i.ID_3,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Type_2,
		&i.IsVerify,
		&i.PasswordChangedAt,
		&i.CreatedAt_3,
		&i.Role,
		&i.Gender,
		&i.Licence,
		&i.Dob,
		&i.Address_2,
		&i.ID_4,
		&i.Username_2,
		&i.HashedPassword_2,
		&i.FullName_2,
		&i.Email_2,
		&i.Type_3,
		&i.IsVerify_2,
		&i.PasswordChangedAt_2,
		&i.CreatedAt_4,
		&i.Role_2,
		&i.Gender_2,
		&i.Licence_2,
		&i.Dob_2,
		&i.Address_3,
		&i.CreatedName,
		&i.UpdatedName,
	)
	return i, err
}

const updateRole = `-- name: UpdateRole :one
UPDATE roles
SET
    code = COALESCE($1, code),
    title = COALESCE($2, title),
    note = COALESCE($3, note),
    user_updated = COALESCE($4, user_updated)
WHERE id = $5
RETURNING id, code, title, note, company, user_created, user_updated, updated_at, created_at
`

type UpdateRoleParams struct {
	Code        sql.NullString `json:"code"`
	Title       sql.NullString `json:"title"`
	Note        sql.NullString `json:"note"`
	UserUpdated sql.NullInt32  `json:"user_updated"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, updateRole,
		arg.Code,
		arg.Title,
		arg.Note,
		arg.UserUpdated,
		arg.ID,
	)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Note,
		&i.Company,
		&i.UserCreated,
		&i.UserUpdated,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateRoleItem = `-- name: UpdateRoleItem :one
UPDATE role_item
SET
    value = COALESCE($1, value)
WHERE roles = $2 AND app = $3
RETURNING id, roles, app, value
`

type UpdateRoleItemParams struct {
	Value sql.NullBool `json:"value"`
	Roles int32        `json:"roles"`
	App   string       `json:"app"`
}

func (q *Queries) UpdateRoleItem(ctx context.Context, arg UpdateRoleItemParams) (RoleItem, error) {
	row := q.db.QueryRowContext(ctx, updateRoleItem, arg.Value, arg.Roles, arg.App)
	var i RoleItem
	err := row.Scan(
		&i.ID,
		&i.Roles,
		&i.App,
		&i.Value,
	)
	return i, err
}
