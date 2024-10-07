package repository

import (
	"app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetFilterData(song app.Song) ([]app.Song, error)
	GetTextSong(id int) (string, error)
	DeleteSongByID(id int) error
	UpdateSongByID(song app.Song) error
	PostNewSong(song app.Song) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
