package data

import (
	"dezzles-apps/rq-server/model"
	services "dezzles-apps/rq-server/services/data"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	pokemonService *services.PokemonDataService
}

func NewPokemonController(
	router *gin.Engine,
	pokemonService *services.PokemonDataService,
) {
	var controller = &PokemonController{
		pokemonService: pokemonService,
	}
	controller.registerRoutes(router)
}

func (pc *PokemonController) registerRoutes(router *gin.Engine) {
	router.GET("/api/data/v1/pokemon", pc.getPokemon)
}

func (pc *PokemonController) getPokemon(c *gin.Context) {
	ctx := model.GetContext(c)
	result, err := pc.pokemonService.GetAllPokemon(ctx)
	if err != nil {
		c.JSON(200, gin.H{"errors": model.InternalServerError})
	}
	c.JSON(200, gin.H{"data": result})
}
