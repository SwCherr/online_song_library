package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	GetFilterDataPaginate(page int, sizePage int, song app.Song) ([]app.Song, error)
	GetTextSongPaginate(id int, page int, sizePage int) ([]string, error)
	DeleteSong(song app.Song) error
	UpdateSong(song app.Song) (int, error)
	PostNewSong(song app.Song) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
