package handler

import (
	"app"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AllInput struct {
	ID       int `json:"id" binding:"required"`
	Page     int `json:"page" binding:"required"`
	SizePage int `json:"sizePage" binding:"required"`
	app.Song
}

type Input struct {
	ID       int `json:"id" binding:"required"`
	Page     int `json:"page" binding:"required"`
	SizePage int `json:"sizePage" binding:"required"`
}

// @Summary get Filter Data Paginate
// @Description Получение всех данных песни с пагинацией по куплетам
// @Tags songs
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Param sizePage query int true "Number of items per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /songs [get]
func (h *Handler) getFilterDataPaginate(c *gin.Context) {
	req := c.Request.URL.Query()

	page, err := strconv.Atoi(req.Get("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	sizePage, err := strconv.Atoi(req.Get("sizePage"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	var input app.Song
	input.Group = req.Get("group")
	input.Song = req.Get("song")
	input.Text = req.Get("text")
	input.Link = req.Get("link")
	input.ReleaseDate = req.Get("releaseDate")

	songs, err := h.service.GetFilterDataPaginate(page, sizePage, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"page":  req["page"],
		"songs": songs,
	})
}

// @Summary get Text Song Paginate
// @Description Получение текста песни с пагинацией по куплетам
// @Tags song
// @Accept  json
// @Produce  json
// @Param id query int true "Song ID"
// @Param page query int true "Page number"
// @Param sizePage query int true "Number of couplets per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [get]
func (h *Handler) getTextSongPaginate(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	couplets, err := h.service.GetTextSongPaginate(input.ID, input.Page, input.SizePage)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"page":     input.Page,
		"couplets": couplets,
	})
}

// @Summary deleteSong
// @Description Удаление песни по названию группы и песни
// @Tags song
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Param sizePage query int true "Number of items per page"
// @Param song body app.Song.Info true "Song main information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [delete]
func (h *Handler) deleteSong(c *gin.Context) {
	var song app.Song
	if err := c.BindJSON(&song.Info); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.service.DeleteSong(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary updateSong
// @Description Обновление данных песни
// @Tags song
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Param sizePage query int true "Number of items per page"
// @Param song body app.Song.Info true "Song main information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [patch]
func (h *Handler) updateSong(c *gin.Context) {
	var song app.Song
	if err := c.BindJSON(&song.Info); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.UpdateSong(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary postNewSong
// @Description Добавление новой песни с использованием стороннего API для обогащения данных
// @Tags song
// @Accept  json
// @Produce  json
// @Param song body app.Song true "New Song Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [post]
func (h *Handler) postNewSong(c *gin.Context) {
	var song app.Song
	if err := c.BindJSON(&song.Info); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := getFullInfo(&song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.service.PostNewSong(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func getFullInfo(song *app.Song) error {
	// --- uncomment for release ---
	// str, err := requestThirdPartyAPI(song)
	// if err != nil {
	// 	return err
	// }

	// --- use mock data: comment for release ---
	str, err := requestMockData()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(str, &song.Detail); err != nil {
		return err
	}
	return nil
}

func requestMockData() ([]byte, error) {
	file, err := os.Open("mockInfoSong.txt")
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	str, err := io.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}
	return str, nil
}

func requestThirdPartyAPI(song *app.Song) ([]byte, error) {
	client, err := app.NewClient()
	if err != nil {
		return []byte{}, err
	}

	group := strings.ReplaceAll(song.Info.Group, " ", "+")
	namesong := strings.ReplaceAll(song.Info.Song, " ", "+")
	finalURL := "https://" + client.Host + "/info?group=" + group + "&song=" + namesong

	resp, err := http.Get(finalURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	str, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return str, nil
}
