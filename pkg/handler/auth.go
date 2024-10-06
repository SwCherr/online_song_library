package handler

import (
	"app"
	"encoding/json"
	"io"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AllInput struct {
	Page     int `json:"page" binding:"required"`
	SizePage int `json:"sizePage" binding:"required"`
	app.Song
}

type Input struct {
	ID       int `json:"id" binding:"required"`
	Page     int `json:"page" binding:"required"`
	SizePage int `json:"sizePage" binding:"required"`
}

func (h *Handler) GetAllData(c *gin.Context) {
	var input AllInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	song, text, err := h.service.GetAllData(input.Page, input.SizePage, input.Song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"page": input.Page,
		"song": map[string]interface{}{
			"id":          song.Id,
			"group":       song.Group,
			"song":        song.Song,
			"releaseDate": song.ReleaseDate,
			"text":        text,
			"link":        song.Link,
		},
	})
}

func (h *Handler) GetSong(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	couplets, err := h.service.GetSong(input.ID, input.Page, input.SizePage)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"page":     input.Page,
		"couplets": couplets,
	})
}

func (h *Handler) DeleteSong(c *gin.Context) {
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

func (h *Handler) UpdateSong(c *gin.Context) {
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

func (h *Handler) PostNewSong(c *gin.Context) {
	var song app.Song
	if err := c.BindJSON(&song.Info); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	// get full info in SongDetail
	err := getFullInfo(&song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// post new song
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

// func requestThirdPartyAPI(song *app.Song) ([]byte, error) {
// 	client, err := app.NewClient()
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	group := strings.ReplaceAll(song.Info.Group, " ", "+")
// 	namesong := strings.ReplaceAll(song.Info.Song, " ", "+")
// 	finalURL := "https://" + client.Host + "/info?group=" + group + "&song=" + namesong

// 	resp, err := http.Get(finalURL)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	defer resp.Body.Close()

// 	str, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	return str, nil
// }
