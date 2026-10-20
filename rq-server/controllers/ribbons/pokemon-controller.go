package ribbons

import (
	"dezzles-apps/rq-server/interfaces"
	"dezzles-apps/rq-server/model/dto"
	services "dezzles-apps/rq-server/services/ribbons"

	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PokemonController struct {
	pokemonService *services.PokemonService
	ribbonService  *services.RibbonService
}

func (pc *PokemonController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	var controller = &PokemonController{
		pokemonService: ctx.Services().RibbonPokemon,
		ribbonService:  ctx.Services().Ribbons,
	}
	controller.registerRoutes(ctx)
}

func (pc *PokemonController) registerRoutes(ctx interfaces.ControllerInitContext) {
	pokemonGroup := ctx.Router().Group("/api/ribbons/v1/pokemon")
	{
		pokemonGroup.GET("/", pc.getAllPokemon)
		pokemonGroup.GET("/:pokemon", pc.getPokemon)
		pokemonGroup.POST("/:pokemon/ribbons/:ribbon", ctx.AuthMiddleware().ValidateUser, pc.addRibbon)
		pokemonGroup.DELETE("/:pokemon/ribbons/:ribbon", ctx.AuthMiddleware().ValidateUser, pc.removeRibbon)
		pokemonGroup.POST("/:pokemon/catch", ctx.AuthMiddleware().ValidateUser, pc.catchPokemon)
		pokemonGroup.PUT("/:pokemon", ctx.AuthMiddleware().ValidateUser, pc.updatePokemon)
		pokemonGroup.POST("/", ctx.AuthMiddleware().ValidateUser, pc.createPokemon)

		pokemonGroup.PUT("/:pokemon/games/:gameKey", ctx.AuthMiddleware().ValidateUser, pc.addPokemonGame)
		pokemonGroup.DELETE("/:pokemon/games/:gameKey", ctx.AuthMiddleware().ValidateUser, pc.removePokemonGame)
	}
}

func (pc *PokemonController) getAllPokemon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	allPokemon, err := pc.pokemonService.GetAllPokemon(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": allPokemon})
}

func (pc *PokemonController) getPokemon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokemon := c.Param("pokemon")
	pokemonData, err := pc.pokemonService.GetPokemon(ctx, pokemon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokemonData})
}

func (pc *PokemonController) addRibbon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
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
	ctx := cmodel.GetContext(c)
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
	ctx := cmodel.GetContext(c)
	pokemon := c.Param("pokemon")

	pokemonData, err := pc.pokemonService.CatchPokemon(ctx, pokemon)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokemonData})
}

func (pc *PokemonController) updatePokemon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
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
	ctx := cmodel.GetContext(c)
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

func (pc *PokemonController) addPokemonGame(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokemon := c.Param("pokemon")
	gameKey := c.Param("gameKey")

	updatedPokemon, err := pc.pokemonService.AddPokemonGame(ctx, pokemon, gameKey)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedPokemon})
}

func (pc *PokemonController) removePokemonGame(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokemon := c.Param("pokemon")
	gameKey := c.Param("gameKey")

	updatedPokemon, err := pc.pokemonService.RemovePokemonGame(ctx, pokemon, gameKey)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedPokemon})
}
