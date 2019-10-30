package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/webnator/capo-music-api/cmd/capo-music/config"
	"github.com/webnator/capo-music-api/cmd/capo-music/controllers"
	"github.com/webnator/capo-music-api/cmd/capo-music/services/db"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/webnator/capo-music-api/cmd/capo-music/docs"
)

// @title capo-music Swagger API
// @version 1.0
// @description Swagger API for Golang Project capo-music.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @BasePath /api/v1
func main() {
	// load application configurations
	if err := config.LoadConfig(); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/songs", controllers.GetSongs)
	}
	db.Connect(db.ConnOptions{
		DBName: config.Config.DBConfig.DBName,
		Host:   config.Config.DBConfig.Host,
		Port:   config.Config.DBConfig.Port,
	})

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
