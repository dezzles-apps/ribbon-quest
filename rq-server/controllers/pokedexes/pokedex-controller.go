package pokedexes

import (
	"dezzles-apps/rq-server/interfaces"
	services "dezzles-apps/rq-server/services/pokedexes"

	cmodel "github.com/dezzles-apps/go-common/model"

	"github.com/gin-gonic/gin"
)

type PokedexController struct {
	pokedexService *services.PokedexService
}

func (pc *PokedexController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	var controller = &PokedexController{
		ctx.Services().Pokedex,
	}
	controller.registerRoutes(ctx)
}

func (pc *PokedexController) registerRoutes(ctx interfaces.ControllerInitContext) {
	group := ctx.Router().Group("/api/pokedexes/v1")
	{
		group.GET("/", pc.getPokedexes)
		group.GET("/:pokedex", pc.getPokedex)
		group.POST("/:pokedex/entries/:pokedexNo", ctx.AuthMiddleware().ValidateUser, pc.catchPokemon)
		group.DELETE("/:pokedex/entries/:pokedexNo", ctx.AuthMiddleware().ValidateUser, pc.uncatchPokemon)
	}
}

func (pc *PokedexController) getPokedex(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	dex := c.Param("pokedex")
	pokedexes, err := pc.pokedexService.GetPokedex(ctx, dex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokedexes})
}

func (pc *PokedexController) getPokedexes(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokedexes, err := pc.pokedexService.GetPokedexes(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokedexes})
}

func (pc *PokedexController) uncatchPokemon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokedex := c.Param("pokedex")
	pokedexNo := c.Param("pokedexNo")
	res, err := pc.pokedexService.UncatchPokemon(ctx, pokedex, pokedexNo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": res})
}

func (pc *PokedexController) catchPokemon(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokedex := c.Param("pokedex")
	pokedexNo := c.Param("pokedexNo")
	res, err := pc.pokedexService.CatchPokemon(ctx, pokedex, pokedexNo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": res})
}
