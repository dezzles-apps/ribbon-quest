package controllers

import (
	"dezzles-apps/rq-server/services"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	pokemonService *services.PokemonService
}

func NewPokemonController(
	router *gin.Engine,
	pokemonService *services.PokemonService,
) *PokemonController {
	var controller = &PokemonController{
		pokemonService: pokemonService,
	}
	controller.registerRoutes(router)
	return controller
}

func (pc *PokemonController) registerRoutes(router *gin.Engine) {
	pokemonGroup := router.Group("/api/pokemon/v1")
	{
		pokemonGroup.GET("/:pokemon", pc.getPokemon)
		pokemonGroup.GET("/", pc.getAllPokemon)
	}
}

func (pc *PokemonController) getAllPokemon(c *gin.Context) {
	allPokemon, err := pc.pokemonService.GetAllPokemon()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": allPokemon})
}

func (pc *PokemonController) getPokemon(c *gin.Context) {
	pokemon := c.Param("pokemon")
	pokemonData, err := pc.pokemonService.GetPokemon(pokemon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokemonData})
}
