package service

import (
	"app"
	"app/pkg/repository"
	"time"
)

const (
	salt            = "ahdbvjdccjdn"
	siginkey        = "djdjdjbvhdn3d^&*("
	accessTokenITL  = 30 * time.Minute
	refreshTokenITL = 720 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
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

// func (s *AuthService) PostNewSong(song app.Song) (int, error) {
// 	return s.repo.PostNewSong(song)
// }

// func (s *AuthService) generateHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))
// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }
