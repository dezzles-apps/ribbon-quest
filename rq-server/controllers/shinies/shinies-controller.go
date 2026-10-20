package shinies

import (
	"dezzles-apps/rq-server/interfaces"
	services "dezzles-apps/rq-server/services/shinies"

	cmodel "github.com/dezzles-apps/go-common/model"

	"github.com/gin-gonic/gin"
)

type ShinyController struct {
	shinyService *services.ShinyService
}

func (sc *ShinyController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	var controller = &ShinyController{
		shinyService: ctx.Services().Shinies,
	}
	controller.registerRoutes(ctx)
}

func (pc *ShinyController) registerRoutes(ctx interfaces.ControllerInitContext) {
	shinies := ctx.Router().Group("/api/shinies/v1")
	{
		shinies.GET("/", pc.getShinies)
		shinies.GET("/stats", pc.getShinyStats)
	}
}

func (pc *ShinyController) getShinies(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokedexes, err := pc.shinyService.GetShinies(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pokedexes})
}

func (pc *ShinyController) getShinyStats(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	stats, err := pc.shinyService.GetShinyStats(ctx, startDate, endDate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": stats})
}
