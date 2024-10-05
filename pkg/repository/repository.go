package repository

import (
	"app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	// GetAllData(song app.Song) (int, error)
	// GetSong(song app.Song) (int, error)
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
