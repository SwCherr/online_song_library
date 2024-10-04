package service

import (
	"app"
	"app/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	// GetAllData(c *gin.Context)
	// GetSong(c *gin.Context)
	// DeleteSong(c *gin.Context)
	// UpdateSong(c *gin.Context)
	// PostNewSong(c *gin.Context)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
