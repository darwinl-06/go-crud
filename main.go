package main

import (
	"github.com/darwinl-06/go_crud/controllers"
	"github.com/darwinl-06/go_crud/initializers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
	
    r.Use(cors.New(config))

	r.POST("/games", controllers.GamesCreate)
	r.PUT("/games/:id", controllers.GamesUpdate)
	r.GET("/games", controllers.GamesGet)
	r.GET("/games/:id", controllers.GamesGetById)
	r.DELETE("/games/:id", controllers.GamesDelete)

	r.Run()
}
