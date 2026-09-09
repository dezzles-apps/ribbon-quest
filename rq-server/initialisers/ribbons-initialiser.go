package initialisers

import (
	"github.com/dezzles-apps/go-common/db"
	"github.com/gin-gonic/gin"

	controllers "dezzles-apps/rq-server/controllers/ribbons"
	"dezzles-apps/rq-server/middleware"
	services "dezzles-apps/rq-server/services/ribbons"
)

func InitialiseRibbons(
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	database *db.Database,
) {
	pokemonService := services.NewPokemonService(database)
	ribbonService := services.NewRibbonService(database)
	gameService := services.NewGameService(database)
	ribbonStatsService := services.NewRibbonStatsService(
		database,
		pokemonService,
	)

	controllers.NewRibbonStatsController(
		router,
		ribbonStatsService,
	)
	controllers.NewPokemonController(router, pokemonService, ribbonService, authMiddleware)
	controllers.NewGamesController(router, gameService)
}
