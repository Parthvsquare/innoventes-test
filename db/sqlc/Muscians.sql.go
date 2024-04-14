// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: Muscians.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addNewMusician = `-- name: AddNewMusician :one
INSERT INTO Musicians (musician_name, musician_type) VALUES ($1, $2) RETURNING musician_id, musician_name, musician_type
`

type AddNewMusicianParams struct {
	MusicianName string `json:"musician_name"`
	MusicianType string `json:"musician_type"`
}

func (q *Queries) AddNewMusician(ctx context.Context, arg AddNewMusicianParams) (Musician, error) {
	row := q.db.QueryRow(ctx, addNewMusician, arg.MusicianName, arg.MusicianType)
	var i Musician
	err := row.Scan(&i.MusicianID, &i.MusicianName, &i.MusicianType)
	return i, err
}

const deleteMusician = `-- name: DeleteMusician :exec
DELETE FROM Musicians WHERE musician_id = $1
`

func (q *Queries) DeleteMusician(ctx context.Context, musicianID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteMusician, musicianID)
	return err
}

const getMusicByMusicianId = `-- name: GetMusicByMusicianId :many
SELECT music_id, music_name, music_price, music_description, album_id, album_name, release_date, genre, price, description, musician_id, musician_name, musician_type FROM album_music_musician_view WHERE musician_id = $1
`

func (q *Queries) GetMusicByMusicianId(ctx context.Context, musicianID uuid.UUID) ([]AlbumMusicMusicianView, error) {
	rows, err := q.db.Query(ctx, getMusicByMusicianId, musicianID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AlbumMusicMusicianView{}
	for rows.Next() {
		var i AlbumMusicMusicianView
		if err := rows.Scan(
			&i.MusicID,
			&i.MusicName,
			&i.MusicPrice,
			&i.MusicDescription,
			&i.AlbumID,
			&i.AlbumName,
			&i.ReleaseDate,
			&i.Genre,
			&i.Price,
			&i.Description,
			&i.MusicianID,
			&i.MusicianName,
			&i.MusicianType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMusicianByAlbumId = `-- name: GetMusicianByAlbumId :one
SELECT album_id, album_name, release_date, genre, price, description, musician_id, musician_name, musician_type FROM album_musician_view WHERE album_id = $1
`

func (q *Queries) GetMusicianByAlbumId(ctx context.Context, albumID uuid.UUID) (AlbumMusicianView, error) {
	row := q.db.QueryRow(ctx, getMusicianByAlbumId, albumID)
	var i AlbumMusicianView
	err := row.Scan(
		&i.AlbumID,
		&i.AlbumName,
		&i.ReleaseDate,
		&i.Genre,
		&i.Price,
		&i.Description,
		&i.MusicianID,
		&i.MusicianName,
		&i.MusicianType,
	)
	return i, err
}

const getMusiciansById = `-- name: GetMusiciansById :one
SELECT musician_id, musician_name, musician_type FROM Musicians WHERE musician_id = $1
`

func (q *Queries) GetMusiciansById(ctx context.Context, musicianID uuid.UUID) (Musician, error) {
	row := q.db.QueryRow(ctx, getMusiciansById, musicianID)
	var i Musician
	err := row.Scan(&i.MusicianID, &i.MusicianName, &i.MusicianType)
	return i, err
}

const updateMusician = `-- name: UpdateMusician :one
UPDATE
  Musicians
SET
  musician_name = COALESCE($1, musician_name),
  musician_type = COALESCE($2, musician_type)
WHERE musician_id = $3
RETURNING musician_id, musician_name, musician_type
`

type UpdateMusicianParams struct {
	MusicianName string    `json:"musician_name"`
	MusicianType string    `json:"musician_type"`
	MusicianID   uuid.UUID `json:"musician_id"`
}

func (q *Queries) UpdateMusician(ctx context.Context, arg UpdateMusicianParams) (Musician, error) {
	row := q.db.QueryRow(ctx, updateMusician, arg.MusicianName, arg.MusicianType, arg.MusicianID)
	var i Musician
	err := row.Scan(&i.MusicianID, &i.MusicianName, &i.MusicianType)
	return i, err
}
