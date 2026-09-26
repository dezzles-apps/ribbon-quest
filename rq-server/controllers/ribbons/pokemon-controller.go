package ribbons

import (
	"dezzles-apps/rq-server/middleware"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	services "dezzles-apps/rq-server/services/ribbons"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		pokemonGroup.POST("/", authMiddleware.ValidateUser, pc.createPokemon)
	}
}

func (pc *PokemonController) getAllPokemon(c *gin.Context) {
	ctx := model.GetContext(c)
	allPokemon, err := pc.pokemonService.GetAllPokemon(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": allPokemon})
}

func (pc *PokemonController) getPokemon(c *gin.Context) {
	ctx := model.GetContext(c)
	pokemon := c.Param("pokemon")
	pokemonData, err := pc.pokemonService.GetPokemon(ctx, pokemon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokemonData})
}

func (pc *PokemonController) addRibbon(c *gin.Context) {
	ctx := model.GetContext(c)
	pokemon := c.Param("pokemon")
	ribbon := c.Param("ribbon")

	ribbonData, err := pc.ribbonService.AddRibbon(ctx, pokemon, ribbon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": ribbonData})
}

func (pc *PokemonController) removeRibbon(c *gin.Context) {
	ctx := model.GetContext(c)
	pokemon := c.Param("pokemon")
	ribbon := c.Param("ribbon")

	ribbonData, err := pc.ribbonService.RemoveRibbon(ctx, pokemon, ribbon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": ribbonData})
}

func (pc *PokemonController) catchPokemon(c *gin.Context) {
	ctx := model.GetContext(c)
	pokemon := c.Param("pokemon")

	pokemonData, err := pc.pokemonService.CatchPokemon(ctx, pokemon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokemonData})
}

func (pc *PokemonController) updatePokemon(c *gin.Context) {
	ctx := model.GetContext(c)
	pokemon := c.Param("pokemon")
	var updateData dto.UpdatePokemon
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedPokemon, err := pc.pokemonService.UpdatePokemon(ctx, pokemon, updateData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedPokemon})
}

func (pc *PokemonController) createPokemon(c *gin.Context) {
	ctx := model.GetContext(c)
	var updateData dto.AddNewRibbonPokemon
	ctx.Logger.Info("Creating new pokemon")
	if err := c.ShouldBindJSON(&updateData); err != nil {
		ctx.Logger.Info("Creating new pokemon failed", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	validations, err := pc.pokemonService.ValidateCreate(ctx, &updateData)
	if err != nil {
		c.JSON(500, gin.H{"errors": err.Error()})
		return
	}
	if validations != nil {
		c.JSON(400, gin.H{"errors": validations})
		return
	}

	updatedPokemon, err := pc.pokemonService.CreatePokemon(ctx, &updateData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedPokemon})
}
