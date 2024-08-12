package main

import (
	"github.com/darwinl-06/go_crud/controllers"
	"github.com/darwinl-06/go_crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/games", controllers.GamesCreate)
	r.PUT("/games/:id", controllers.GamesUpdate)
	r.GET("/games", controllers.GamesGet)
	r.GET("/games/:id", controllers.GamesGetById)
	r.DELETE("/games/:id", controllers.GamesDelete)

	r.Run()
}
