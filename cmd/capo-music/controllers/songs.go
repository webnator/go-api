package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webnator/capo-music-api/cmd/capo-music/songs"
)

// GetSongs godoc
// @Summary Retrieves songs based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /songs/{id} [get]
func GetSongs(context *gin.Context) {
	if songs, err := songs.GetAll(); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		context.JSON(http.StatusOK, songs)
	}
}
