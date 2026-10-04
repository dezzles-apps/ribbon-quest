package pokedexes

import (
	services "dezzles-apps/rq-server/services/pokedexes"

	"dezzles-apps/rq-server/middleware"

	cmodel "github.com/dezzles-apps/go-common/model"

	"github.com/gin-gonic/gin"
)

type PokedexController struct {
	pokedexService *services.PokedexService
}

func NewPokedexController(
	router *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	pokedexService *services.PokedexService,
) *PokedexController {
	var controller = &PokedexController{
		pokedexService: pokedexService,
	}
	controller.registerRoutes(router)
	return controller
}

func (pc *PokedexController) registerRoutes(router *gin.Engine) {
	group := router.Group("/api/pokedexes/v1")
	{
		group.GET("/", pc.getPokedexes)
		group.GET("/:pokedex", pc.getPokedex)
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
