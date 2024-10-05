package handler

import (
	"app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/api")
	{
		auth.GET("/songs", h.GetAllData)
		auth.GET("/song", h.GetSong)
		auth.DELETE("/song", h.DeleteSong)
		auth.PATCH("/song", h.UpdateSong)
		auth.POST("/song", h.PostNewSong)
	}
	return router
}
