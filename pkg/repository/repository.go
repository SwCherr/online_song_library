package repository

import (
	"app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
