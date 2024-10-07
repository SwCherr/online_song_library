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

func (s *AuthService) GetFilterDataPaginate(page int, sizePage int, song app.Song) ([]app.Song, error) {
	info, err := s.repo.GetFilterData(song)
	if err != nil {
		return []app.Song{}, err
	}

	start, end, err := s.calculationStartEndPage(len(info), page, sizePage)
	if err != nil {
		return []app.Song{}, err
	}
	return info[start:end], nil
}

func (s *AuthService) GetTextSongPaginate(id int, page int, sizePage int) ([]string, error) {
	text, err := s.repo.GetTextSong(id)
	if err != nil {
		return []string{}, err
	}
	couplets := strings.Split(text, "\n\n")
	start, end, err := s.calculationStartEndPage(len(couplets), page, sizePage)
	if err != nil {
		return []string{}, err
	}
	return couplets[start:end], nil
}

func (s *AuthService) DeleteSongByID(id int) error {
	return s.repo.DeleteSongByID(id)
}

func (s *AuthService) UpdateSongByID(song app.Song) error {
	return s.repo.UpdateSongByID(song)
}

func (s *AuthService) PostNewSong(song app.Song) (int, error) {
	return s.repo.PostNewSong(song)
}

func (s *AuthService) calculationStartEndPage(len_arr int, page int, sizePage int) (int, int, error) {
	if page < 1 || sizePage < 1 {
		return 0, 0, errors.New("incorrect size page")
	}

	start := (page - 1) * sizePage
	end := (page-1)*sizePage + sizePage

	if start >= len_arr {
		return 0, 0, errors.New("incorrect size page")
	}

	if end > len_arr {
		end = len_arr
	}

	return start, end, nil
}
