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
	start, end, err := calculationStartEndPage(page, sizePage)
	if err != nil {
		return []app.Song{}, err
	}

	info, err := s.repo.GetFilterData(song)
	if err != nil {
		return []app.Song{}, err
	}

	if start >= len(info) {
		return []app.Song{}, errors.New("incorrect size page")
	}

	if end > len(info) {
		end = len(info)
	}

	return info[start:end], nil
}

func (s *AuthService) GetTextSongPaginate(id int, page int, sizePage int) ([]string, error) {
	text, err := s.repo.GetTextSong(id)
	if err != nil {
		return []string{}, err
	}
	return s.paginateTextSong(text, page, sizePage)
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

func calculationStartEndPage(page int, sizePage int) (int, int, error) {
	if page < 1 || sizePage < 1 {
		return 0, 0, errors.New("incorrect size page")
	}
	start := (page - 1) * sizePage
	end := (page-1)*sizePage + sizePage
	return start, end, nil
}

func (s *AuthService) paginateTextSong(text string, page int, sizePage int) ([]string, error) {
	start, end, err := calculationStartEndPage(page, sizePage)
	if err != nil {
		return []string{}, err
	}

	couplets := strings.Split(text, "\n\n")
	if start >= len(couplets) {
		return []string{}, errors.New("incorrect size page")
	}

	if end > len(couplets) {
		end = len(couplets)
	}

	return couplets[start:end], nil
}
