package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webnator/capo-music-api/cmd/capo-music/services/db"
	"github.com/webnator/capo-music-api/cmd/capo-music/songs"
)

// GetSongs godoc
// @Summary Retrieves songs based on given ID
// @Produce json
// @Success 200 {object} models.User
// @Router /songs [get]
func GetSongs(context *gin.Context) {
	songService := songs.NewSongService(songs.Deps{DB: db.Service()})
	if songs, err := songService.Find(context.Query("search")); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		context.JSON(http.StatusOK, songs)
	}
}

// GetSong godoc
// @Summary Retrieves songs based on given ID
// @Produce json
// @Param id path integer true "Song slug"
// @Success 200 {object} models.User
// @Router /songs/{id} [get]
func GetSong(context *gin.Context) {
	songService := songs.NewSongService(songs.Deps{DB: db.Service()})
	if song, err := songService.FindSongBySlug(context.Param("slug")); err != nil || song == nil {
		context.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		context.JSON(http.StatusOK, song)
	}
}
