package controllers

import (
	"github.com/darwinl-06/go_crud/initializers"
	"github.com/darwinl-06/go_crud/models"
	"github.com/gin-gonic/gin"
)

func GamesCreate(c *gin.Context) {

	// Get data off req body
	var body struct {
		Name string
		Type string
		Year string
	}

	c.Bind(&body)

	// Create a post
	game := models.Game{Name: body.Name, Type: body.Type, Year: body.Year}

	result := initializers.DB.Create(&game)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"game": game,
	})
}

func GamesGet(c *gin.Context) {
	// Get games
	var games []models.Game

	initializers.DB.Find(&games)

	// Return games
	c.JSON(200, gin.H{
		"games": games,
	})
}

func GamesDelete(c *gin.Context) {
	// Get the id
	id := c.Param("id")

	// Delete the game
	initializers.DB.Delete(&models.Game{}, id)

	// Return
	c.Status(200)
}

func GamesUpdate(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Get data from body
	var body struct {
		Name    string
		Type	string
		Year 	string
	}

	c.Bind(&body)

	// Finf the game were updating
	var game models.Game

	initializers.DB.First(&game, id)

	if game.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Not found",
		})
	} else {
		// Update it
		initializers.DB.Model(&game).Updates(models.Game{
			Name:	 body.Name,
			Type:	 body.Type,
			Year: 	 body.Year,
		})

		// Return it
		c.JSON(200, gin.H{
			"game": game,
		})
	}
}

func GamesGetById(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Get the game
	var game models.Game

	// al manipular con la DB usando la variable
	// esta se modifica con el resultado de la operación
	initializers.DB.First(&game, id)

	// Return
	if game.ID == 0 {
		c.JSON(400, gin.H{
			"message": "Not found",
		})
	} else {
		c.JSON(200, gin.H{
			"game": game,
		})
	}
}
