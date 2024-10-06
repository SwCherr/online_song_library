package service

import (
	"app"
	"app/pkg/repository"
	"errors"
	"strings"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GetAllData(page int, sizePage int, song app.Song) (app.Song, []string, error) {
	info, err := s.repo.GetAllData(page, sizePage, song)
	if err != nil {
		return app.Song{}, []string{}, err
	}
	text, err := s.paginate(info.Text, page, sizePage)
	if err != nil {
		return app.Song{}, []string{}, err
	}

	return info, text, nil
}

func (s *AuthService) GetSong(id int, page int, sizePage int) ([]string, error) {
	text, err := s.repo.GetSong(id)
	if err != nil {
		return []string{}, err
	}
	return s.paginate(text, page, sizePage)
}

func (s *AuthService) paginate(text string, page int, sizePage int) ([]string, error) {
	if page < 1 || sizePage < 1 {
		return []string{}, errors.New("incorrect size page")
	}
	couplets := strings.Split(text, "\n\n")
	start := (page - 1) * sizePage
	end := (page-1)*sizePage + sizePage

	if start >= len(couplets) {
		return []string{}, errors.New("incorrect size page")
	}

	if end > len(couplets) {
		end = len(couplets)
	}

	return couplets[start:end], nil
}

func (s *AuthService) DeleteSong(song app.Song) error {
	return s.repo.DeleteSong(song)
}

func (s *AuthService) UpdateSong(song app.Song) (int, error) {
	return s.repo.UpdateSong(song)
}

func (s *AuthService) PostNewSong(song app.Song) (int, error) {
	return s.repo.PostNewSong(song)
}

// func (s *AuthService) PostNewSong(song app.Song) (int, error) {
// 	return s.repo.PostNewSong(song)
// }
