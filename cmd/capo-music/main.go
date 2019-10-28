package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/webnator/capo-music-api/cmd/capo-music/apis"
	"github.com/webnator/capo-music-api/cmd/capo-music/config"

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
		v1.GET("/users/:id", apis.GetUser)
	}

	// config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	// if config.Config.DBErr != nil {
	// 	panic(config.Config.DBErr)
	// }

	// // config.Config.DB.AutoMigrate(&models.User{}) // This is needed for generation of schema for postgres image.

	// defer config.Config.DB.Close()

	// log.Println("Successfully connected to database")

	// log.Println(config.Config.TestVal)

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
