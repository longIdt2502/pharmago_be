// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: media.sql

package db

import (
	"context"
)

const createMedia = `-- name: CreateMedia :one
INSERT INTO medias (
    media_url
) VALUES ($1) RETURNING id, media_url
`

func (q *Queries) CreateMedia(ctx context.Context, mediaUrl string) (Media, error) {
	row := q.db.QueryRowContext(ctx, createMedia, mediaUrl)
	var i Media
	err := row.Scan(&i.ID, &i.MediaUrl)
	return i, err
}

const getMedia = `-- name: GetMedia :one
SELECT id, media_url FROM medias
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMedia(ctx context.Context, id int32) (Media, error) {
	row := q.db.QueryRowContext(ctx, getMedia, id)
	var i Media
	err := row.Scan(&i.ID, &i.MediaUrl)
	return i, err
}

const getMediaVariant = `-- name: GetMediaVariant :one
SELECT vm.id, variant, media, m.id, media_url, m.media_url FROM variant_media vm
JOIN medias m ON m.id = vm.media
WHERE variant = $1 LIMIT 1
`

type GetMediaVariantRow struct {
	ID         int32  `json:"id"`
	Variant    int32  `json:"variant"`
	Media      int32  `json:"media"`
	ID_2       int32  `json:"id_2"`
	MediaUrl   string `json:"media_url"`
	MediaUrl_2 string `json:"media_url_2"`
}

func (q *Queries) GetMediaVariant(ctx context.Context, variant int32) (GetMediaVariantRow, error) {
	row := q.db.QueryRowContext(ctx, getMediaVariant, variant)
	var i GetMediaVariantRow
	err := row.Scan(
		&i.ID,
		&i.Variant,
		&i.Media,
		&i.ID_2,
		&i.MediaUrl,
		&i.MediaUrl_2,
	)
	return i, err
}