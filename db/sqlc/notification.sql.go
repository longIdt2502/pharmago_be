// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: notification.sql

package db

import (
	"context"
	"database/sql"
)

const countNotification = `-- name: CountNotification :many
SELECT COUNT(*), is_read FROM notification
WHERE company = $1
GROUP BY is_read
`

type CountNotificationRow struct {
	Count  int64 `json:"count"`
	IsRead bool  `json:"is_read"`
}

func (q *Queries) CountNotification(ctx context.Context, company sql.NullInt32) ([]CountNotificationRow, error) {
	rows, err := q.db.QueryContext(ctx, countNotification, company)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CountNotificationRow{}
	for rows.Next() {
		var i CountNotificationRow
		if err := rows.Scan(&i.Count, &i.IsRead); err != nil {
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

const createNotification = `-- name: CreateNotification :one
INSERT INTO notification (
    type, topic, title, content, is_read, data, company
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, type, topic, title, content, is_read, data, company, created_at
`

type CreateNotificationParams struct {
	Type    string         `json:"type"`
	Topic   string         `json:"topic"`
	Title   string         `json:"title"`
	Content string         `json:"content"`
	IsRead  bool           `json:"is_read"`
	Data    sql.NullString `json:"data"`
	Company sql.NullInt32  `json:"company"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error) {
	row := q.db.QueryRowContext(ctx, createNotification,
		arg.Type,
		arg.Topic,
		arg.Title,
		arg.Content,
		arg.IsRead,
		arg.Data,
		arg.Company,
	)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Topic,
		&i.Title,
		&i.Content,
		&i.IsRead,
		&i.Data,
		&i.Company,
		&i.CreatedAt,
	)
	return i, err
}

const listNotification = `-- name: ListNotification :many
SELECT id, type, topic, title, content, is_read, data, company, created_at FROM notification 
WHERE company = $1::int
AND (
    title ILIKE '%' || COALESCE($2::varchar, '') || '%' OR
    content ILIKE '%' || COALESCE($2::varchar, '') || '%'
)
ORDER BY -id
LIMIT COALESCE($4::int, 10)
OFFSET (COALESCE($3::int, 1) - 1) * COALESCE($4::int, 10)
`

type ListNotificationParams struct {
	Company int32          `json:"company"`
	Search  sql.NullString `json:"search"`
	Page    sql.NullInt32  `json:"page"`
	Limit   sql.NullInt32  `json:"limit"`
}

func (q *Queries) ListNotification(ctx context.Context, arg ListNotificationParams) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, listNotification,
		arg.Company,
		arg.Search,
		arg.Page,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Notification{}
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Topic,
			&i.Title,
			&i.Content,
			&i.IsRead,
			&i.Data,
			&i.Company,
			&i.CreatedAt,
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