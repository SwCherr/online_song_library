package handler

import (
	"encoding/json"
	"io"
	"online_music/base"
	"os"
	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AllInput struct {
	ID       int `json:"id" binding:"required"`
	Page     int `json:"page" binding:"required"`
	SizePage int `json:"sizePage" binding:"required"`
	base.Song
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
// @Param group query string false "Song group"
// @Param song query string false "Song name"
// @Param releaseDate query string false "Song releaseDate"
// @Param text query string false "Song text"
// @Param link query string false "Song link"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /songs [get]
func (h *Handler) getFilterDataPaginate(c *gin.Context) {
	req := c.Request.URL.Query()

	page, err := strconv.Atoi(req.Get("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read page: ", err)
		return
	}

	sizePage, err := strconv.Atoi(req.Get("sizePage"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read size page: ", err)
		return
	}

	var input base.Song
	input.Group = req.Get("group")
	input.Song = req.Get("song")
	input.Text = req.Get("text")
	input.Link = req.Get("link")
	input.ReleaseDate = req.Get("releaseDate")

	songs, err := h.service.GetFilterDataPaginate(page, sizePage, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error("error get filter data paginate: ", err)
		return
	}

	logrus.Info("get data by filters paginate: OK")

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
	req := c.Request.URL.Query()

	page, err := strconv.Atoi(req.Get("page"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read page: ", err)
		return
	}

	sizePage, err := strconv.Atoi(req.Get("sizePage"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read size page: ", err)
		return
	}

	id, err := strconv.Atoi(req.Get("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read id: ", err)
		return
	}

	couplets, err := h.service.GetTextSongPaginate(id, page, sizePage)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error("error get text song paginate: ", err)
		return
	}

	logrus.Info("get text song paginate: OK")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       id,
		"page":     page,
		"couplets": couplets,
	})
}

// @Summary deleteSong
// @Description Удаление песни по названию группы и песни
// @Tags song
// @Accept  json
// @Produce  json
// @Param id query int true "id song"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [delete]
func (h *Handler) deleteSongByID(c *gin.Context) {
	req := c.Request.URL.Query()
	id, err := strconv.Atoi(req.Get("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read id: ", err)
		return
	}

	if err := h.service.DeleteSongByID(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error("error ddelete song by id: ", err)
		return
	}

	logrus.Info("delete song by id: OK")
	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary updateSong
// @Description Обновление данных песни
// @Tags song
// @Accept  json
// @Produce  json
// @Param song body base.Song true "Song Errorrmation for update"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [patch]
func (h *Handler) updateSongByID(c *gin.Context) {
	var song base.Song
	if err := c.BindJSON(&song); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read Error song: ", err)
		return
	}

	if err := h.service.UpdateSongByID(song); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error("error update song by id: ", err)
		return
	}

	logrus.Info("update song by id: OK")
	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary postNewSong
// @Description Добавление новой песни с использованием стороннего API для обогащения данных
// @Tags song
// @Accept  json
// @Produce  json
// @Param song body base.Song true "New Song Data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid input body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /song [post]
func (h *Handler) postNewSong(c *gin.Context) {
	var song base.Song
	if err := c.BindJSON(&song); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error("error read Error song: ", err)
		return
	}

	err := getFullError(&song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error("error get full Error from third party API: ", err)
		return
	}

	id, err := h.service.PostNewSong(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		logrus.Error("error post song by id: ", err)
		return
	}

	logrus.Info("post new song: OK")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func getFullError(song *base.Song) error {
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

func requestThirdPartyAPI(song *base.Song) ([]byte, error) {
	client, err := base.NewClient()
	if err != nil {
		return []byte{}, err
	}

	group := strings.ReplaceAll(song.Group, " ", "+")
	namesong := strings.ReplaceAll(song.Song, " ", "+")
	finalURL := "https://" + client.Host + "/Error?group=" + group + "&song=" + namesong

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
