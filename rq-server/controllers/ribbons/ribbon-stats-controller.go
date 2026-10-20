package ribbons

import (
	"dezzles-apps/rq-server/interfaces"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/services/ribbons"

	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
)

type RibbonStatsController struct {
	ribbonStatsService *ribbons.RibbonStatsService
}

func (rsc *RibbonStatsController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	controller := RibbonStatsController{
		ribbonStatsService: ctx.Services().RibbonStats,
	}
	controller.createRoutes(ctx)
}

func (rsc *RibbonStatsController) createRoutes(ctx interfaces.ControllerInitContext) {
	ctx.Router().GET("/api/ribbons/v1/stats", rsc.getStats)
}

func (rsc *RibbonStatsController) getStats(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	stats, err := rsc.ribbonStatsService.GetStats(ctx)
	if err != nil {
		c.JSON(200, gin.H{"errors": model.InternalServerError})
		return
	}
	c.JSON(200, gin.H{"data": stats})
}
