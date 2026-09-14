package initialisers

import (
	controllers "dezzles-apps/rq-server/controllers/data"
	services "dezzles-apps/rq-server/services/data"

	"github.com/dezzles-apps/go-common/db"
	"github.com/gin-gonic/gin"
)

func InitialiseData(
	router *gin.Engine,
	connection *db.Database,
) {
	pokemonService := services.NewPokemonDataService(connection)

	controllers.NewPokemonController(
		router,
		pokemonService,
	)
}
