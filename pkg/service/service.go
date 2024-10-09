package service

import (
	"online_music/base"
	"online_music/pkg/repository"
)

type Authorization interface {
	GetFilterDataPaginate(page int, sizePage int, song base.Song) ([]base.Song, error)
	GetTextSongPaginate(id int, page int, sizePage int) ([]string, error)
	DeleteSongByID(id int) error
	UpdateSongByID(song base.Song) error
	PostNewSong(song base.Song) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
