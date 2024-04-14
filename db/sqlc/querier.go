// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddNewMusic(ctx context.Context, arg AddNewMusicParams) (Music, error)
	DeleteMusic(ctx context.Context, musicID uuid.UUID) error
	GetMusicAlbumsById(ctx context.Context, albumID uuid.UUID) (Musicalbum, error)
	GetMusicByAlbumId(ctx context.Context, albumID uuid.UUID) ([]Musicview, error)
	GetMusicById(ctx context.Context, musicID uuid.UUID) (Music, error)
	MusiciansById(ctx context.Context, musicianID uuid.UUID) (Musician, error)
	UpdateMusic(ctx context.Context, arg UpdateMusicParams) (Music, error)
}

var _ Querier = (*Queries)(nil)
