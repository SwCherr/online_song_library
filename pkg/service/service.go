package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	// GetAllData(song app.Song) (int, error)
	// GetSong(song app.Song) (int, error)
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
