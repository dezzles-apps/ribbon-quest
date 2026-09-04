package controllers

import (
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/services"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	pokemonService *services.PokemonService
	ribbonService  *services.RibbonService
}

func NewPokemonController(
	router *gin.Engine,
	pokemonService *services.PokemonService,
	ribbonService *services.RibbonService,
	authMiddleware *middleware.AuthMiddleware,
) *PokemonController {
	var controller = &PokemonController{
		pokemonService: pokemonService,
		ribbonService:  ribbonService,
	}
	controller.registerRoutes(router, authMiddleware)
	return controller
}

func (pc *PokemonController) registerRoutes(router *gin.Engine, authMiddleware *middleware.AuthMiddleware) {
	pokemonGroup := router.Group("/api/ribbons/v1/pokemon")
	{
		pokemonGroup.GET("/", pc.getAllPokemon)
		pokemonGroup.GET("/:pokemon", pc.getPokemon)
		pokemonGroup.POST("/:pokemon/ribbons/:ribbon", authMiddleware.ValidateUser, pc.addRibbon)
		pokemonGroup.DELETE("/:pokemon/ribbons/:ribbon", authMiddleware.ValidateUser, pc.removeRibbon)
		pokemonGroup.POST("/:pokemon/catch", authMiddleware.ValidateUser, pc.catchPokemon)
		pokemonGroup.PUT("/:pokemon", authMiddleware.ValidateUser, pc.updatePokemon)
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

func (pc *PokemonController) addRibbon(c *gin.Context) {
	pokemon := c.Param("pokemon")
	ribbon := c.Param("ribbon")

	ribbonData, err := pc.ribbonService.AddRibbon(pokemon, ribbon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": ribbonData})
}

func (pc *PokemonController) removeRibbon(c *gin.Context) {
	pokemon := c.Param("pokemon")
	ribbon := c.Param("ribbon")

	ribbonData, err := pc.ribbonService.RemoveRibbon(pokemon, ribbon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": ribbonData})
}

func (pc *PokemonController) catchPokemon(c *gin.Context) {
	pokemon := c.Param("pokemon")

	pokemonData, err := pc.pokemonService.CatchPokemon(pokemon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokemonData})
}

func (pc *PokemonController) updatePokemon(c *gin.Context) {
	pokemon := c.Param("pokemon")

	var updateData dto.UpdatePokemon
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedPokemon, err := pc.pokemonService.UpdatePokemon(pokemon, updateData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedPokemon})
}
