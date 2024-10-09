package repository

import (
	"online_music/base"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetFilterData(song base.Song) ([]base.Song, error)
	GetTextSong(id int) (string, error)
	DeleteSongByID(id int) error
	UpdateSongByID(song base.Song) error
	PostNewSong(song base.Song) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
