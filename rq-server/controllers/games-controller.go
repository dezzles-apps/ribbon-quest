package controllers

import (
	"dezzles-apps/rq-server/services"

	"github.com/gin-gonic/gin"
)

type GamesController struct {
	gameService *services.GameService
}

func NewGamesController(
	router *gin.Engine,
	gameService *services.GameService,
) *GamesController {
	var controller = &GamesController{
		gameService: gameService,
	}
	controller.registerRoutes(router)
	return controller
}

func (gc *GamesController) registerRoutes(router *gin.Engine) {
	gameGroup := router.Group("/api/games/v1")
	{
		gameGroup.GET("/:game", gc.getGame)
		gameGroup.GET("/", gc.getAllGames)
	}
}

func (gc *GamesController) getAllGames(c *gin.Context) {
	allGames, err := gc.gameService.GetAllGames()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": allGames})
}

func (gc *GamesController) getGame(c *gin.Context) {
	game := c.Param("game")
	gameData, err := gc.gameService.GetGame(game)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": gameData})
}
