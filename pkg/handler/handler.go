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
		auth.GET("/signup", h.GetAllData)
		auth.GET("/signup", h.GetSong)
		auth.DELETE("/signup", h.DeleteSong)
		auth.PATCH("/signup", h.UpdateSong)
		auth.POST("/signup", h.PostNewSong)
	}
	return router
}
