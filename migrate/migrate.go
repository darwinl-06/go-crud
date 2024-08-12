package main

import (
	"github.com/darwinl-06/go_crud/initializers"
	"github.com/darwinl-06/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Game{})
}