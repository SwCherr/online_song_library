package service

import (
	"app"
	"app/pkg/repository"
	"crypto/sha1"
	"fmt"
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

func (s *AuthService) CreateUser(user app.User) (int, error) {
	user.Password = s.generateHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
