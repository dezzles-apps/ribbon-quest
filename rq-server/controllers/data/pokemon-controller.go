package data

import (
	"dezzles-apps/rq-server/interfaces"
	"dezzles-apps/rq-server/model"
	services "dezzles-apps/rq-server/services/data"

	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	pokemonService    *services.PokemonDataService
	commonDataService *services.CommonDataService
}

func (pc *PokemonController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	var controller = &PokemonController{
		pokemonService:    ctx.Services().PokemonData,
		commonDataService: ctx.Services().CommonData,
	}
	controller.registerRoutes(ctx)
}

func (pc *PokemonController) registerRoutes(ctx interfaces.ControllerInitContext) {
	ctx.Router().GET("/api/data/v1/pokemon", pc.getPokemon)
	ctx.Router().GET("/api/data/v1/common", pc.getCommonData)
}

func (pc *PokemonController) getPokemon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	result, err := pc.pokemonService.GetAllPokemon(ctx)
	if err != nil {
		c.JSON(400, gin.H{"errors": model.InternalServerError})
		return
	}
	c.JSON(200, gin.H{"data": result})
}

func (pc *PokemonController) getCommonData(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	result, err := pc.commonDataService.GetAllData(ctx)
	if err != nil {
		c.JSON(500, gin.H{"errors": model.InternalServerError})
		return
	}
	c.JSON(200, gin.H{"data": result})
}
