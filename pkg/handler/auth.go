package handler

import (
	"app"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input app.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllData(c *gin.Context) {

}

func (h *Handler) GetSong(c *gin.Context) {

}

func (h *Handler) DeleteSong(c *gin.Context) {

}

func (h *Handler) UpdateSong(c *gin.Context) {

}

func (h *Handler) PostNewSong(c *gin.Context) {

}
