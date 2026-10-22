package shinies

import (
	"dezzles-apps/rq-server/interfaces"
	dto "dezzles-apps/rq-server/model/dto/shinies"
	"dezzles-apps/rq-server/model/errors"
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
		shinies.POST("/", ctx.AuthMiddleware().ValidateUser, pc.createShiny)
	}
}

func (pc *ShinyController) getShinies(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	pokedexes, err := pc.shinyService.GetShinies(ctx)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"errors": err.Errors()})
		return
	}

	c.JSON(200, gin.H{"data": pokedexes})
}

func (pc *ShinyController) createShiny(c *gin.Context) {
	var input dto.CreateShinyRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"errors": errors.InvalidRequestBody.Errors()})
		return
	}
	ctx := cmodel.GetContext(c)
	err := pc.shinyService.CreateShiny(ctx, input)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"errors": err.Errors()})
		return
	}

	c.JSON(201, gin.H{"data": nil})
}

func (pc *ShinyController) getShinyStats(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	stats, err := pc.shinyService.GetShinyStats(ctx, startDate, endDate)
	if err != nil {
		c.JSON(err.StatusCode(), gin.H{"errors": err.Errors()})
		return
	}

	c.JSON(200, gin.H{"data": stats})
}
