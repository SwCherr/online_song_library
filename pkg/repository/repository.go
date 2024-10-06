package repository

import (
	"app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetAllData(page int, sizePage int, song app.Song) (app.Song, error)
	GetSong(id int) (string, error)
	DeleteSong(song app.Song) error
	UpdateSong(song app.Song) (int, error)
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
