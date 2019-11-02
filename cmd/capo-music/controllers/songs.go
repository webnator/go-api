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
	params := map[string]string{
		"search":   context.Query("search"),
		"category": context.Query("category"),
	}
	if songs, err := songService.Find(params); err != nil {
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

// GetCategories godoc
// @Summary Retrieves song categories
// @Produce json
// @Success 200 {object} models.User
// @Router /categories [get]
func GetCategories(context *gin.Context) {
	songService := songs.NewSongService(songs.Deps{DB: db.Service()})
	if categories, err := songService.GetSongCategories(); err != nil || categories == nil {
		context.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		context.JSON(http.StatusOK, categories)
	}
}

// AddSong godoc
// @Summary Adds a new song
// @Produce json
// @Success 201 {object} models.User
// @Router /songs [post]
func AddSong(context *gin.Context) {
	songService := songs.NewSongService(songs.Deps{DB: db.Service()})

	var song songs.SongModel
	if err := context.ShouldBindJSON(&song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := songService.AddSong(song); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		context.JSON(http.StatusCreated, gin.H{})
	}
}

// UpdateSong godoc
// @Summary Updates an existing song
// @Produce json
// @Success 201 {object} models.User
// @Router /songs/:slug [patch]
func UpdateSong(context *gin.Context) {
	songService := songs.NewSongService(songs.Deps{DB: db.Service()})

	var song songs.SongUpdateModel
	if err := context.ShouldBindJSON(&song); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if song.Viewed != 0 {
		songService.IncreaseViewCount(context.Param("slug"))
		song.Viewed = 0
	}

	if err := songService.UpdateSong(context.Param("slug"), song); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		context.JSON(http.StatusCreated, gin.H{})
	}
}

// ResetViews godoc
// @Summary Updates an existing song
// @Produce json
// @Success 201 {object} models.User
// @Router /songs/:slug [patch]
func ResetViews(context *gin.Context) {
	songService := songs.NewSongService(songs.Deps{DB: db.Service()})

	var resetParams songs.ResetViewsModel
	if err := context.ShouldBindJSON(&resetParams); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := songService.ResetViews(resetParams); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		context.JSON(http.StatusCreated, gin.H{})
	}
}
