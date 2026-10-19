package ribbons

import (
	"dezzles-apps/rq-server/interfaces"
	services "dezzles-apps/rq-server/services/ribbons"

	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
)

type GamesController struct {
	gameService *services.GameService
}

func (gc *GamesController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	var controller = &GamesController{
		gameService: ctx.Services().RibbonGames,
	}
	controller.registerRoutes(ctx)
}

func (gc *GamesController) registerRoutes(ctx interfaces.ControllerInitContext) {
	gameGroup := ctx.Router().Group("/api/ribbons/v1/games")
	{
		gameGroup.GET("/:game", gc.getGame)
		gameGroup.GET("/", gc.getAllGames)
	}
}

func (gc *GamesController) getAllGames(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	allGames, err := gc.gameService.GetAllGames(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": allGames})
}

func (gc *GamesController) getGame(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	game := c.Param("game")
	gameData, err := gc.gameService.GetGame(ctx, game)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": gameData})
}
