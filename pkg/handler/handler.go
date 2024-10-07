package handler

import (
	"app/pkg/service"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/api")
	{
		auth.GET("/songs", h.getFilterDataPaginate)
		auth.GET("/song", h.getTextSongPaginate)
		auth.DELETE("/song", h.deleteSong)
		auth.PATCH("/song", h.updateSong)
		auth.POST("/song", h.postNewSong)
	}
	return router
}
