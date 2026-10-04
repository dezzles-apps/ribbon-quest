package initialisers

import (
	"github.com/dezzles-apps/go-common/db"
	"github.com/gin-gonic/gin"

	controllers "dezzles-apps/rq-server/controllers/pokedexes"
	"dezzles-apps/rq-server/middleware"
	services "dezzles-apps/rq-server/services/pokedexes"
)

func InitialisePokedex(
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	database *db.Database,
) {
	pokedexService := services.NewPokedexService(database)
	controllers.NewPokedexController(
		router,
		authMiddleware,
		pokedexService,
	)
}
