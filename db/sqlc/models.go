// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Albummusician struct {
	AlbumID    uuid.UUID `json:"album_id"`
	MusicianID uuid.UUID `json:"musician_id"`
}

type Albummusiciansview struct {
	AlbumID      uuid.UUID      `json:"album_id"`
	AlbumName    string         `json:"album_name"`
	ReleaseDate  pgtype.Date    `json:"release_date"`
	Genre        pgtype.Text    `json:"genre"`
	Price        pgtype.Numeric `json:"price"`
	Description  pgtype.Text    `json:"description"`
	MusicianID   uuid.UUID      `json:"musician_id"`
	MusicianName string         `json:"musician_name"`
	MusicianType string         `json:"musician_type"`
}

type Musicalbum struct {
	AlbumID     uuid.UUID      `json:"album_id"`
	AlbumName   string         `json:"album_name"`
	ReleaseDate pgtype.Date    `json:"release_date"`
	Genre       pgtype.Text    `json:"genre"`
	Price       pgtype.Numeric `json:"price"`
	Description pgtype.Text    `json:"description"`
}

type Musician struct {
	MusicianID   uuid.UUID `json:"musician_id"`
	MusicianName string    `json:"musician_name"`
	MusicianType string    `json:"musician_type"`
}

type Musiciansview struct {
	MusicianID   uuid.UUID      `json:"musician_id"`
	MusicianName string         `json:"musician_name"`
	MusicianType string         `json:"musician_type"`
	AlbumID      uuid.UUID      `json:"album_id"`
	AlbumName    string         `json:"album_name"`
	ReleaseDate  pgtype.Date    `json:"release_date"`
	Genre        pgtype.Text    `json:"genre"`
	Price        pgtype.Numeric `json:"price"`
	Description  pgtype.Text    `json:"description"`
}