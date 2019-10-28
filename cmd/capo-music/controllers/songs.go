package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/webnator/capo-music-api/cmd/capo-music/daos"
	"github.com/webnator/capo-music-api/cmd/capo-music/services"
)

// GetSongs godoc
// @Summary Retrieves songs based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /songs/{id} [get]
func GetSongs(context *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		context.JSON(http.StatusOK, user)
	}
}
