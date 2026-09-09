package ribbons

import (
	"dezzles-apps/rq-server/services/ribbons"

	"github.com/gin-gonic/gin"
)

type RibbonStatsController struct {
	ribbonStatsService *ribbons.RibbonStatsService
}

func NewRibbonStatsController(
	router *gin.Engine,
	service *ribbons.RibbonStatsService,
) *RibbonStatsController {
	controller := RibbonStatsController{
		ribbonStatsService: service,
	}
	controller.createRoutes(router)
	return &controller
}

func (rsc *RibbonStatsController) createRoutes(router *gin.Engine) {
	router.GET("/api/ribbons/v1/stats", rsc.getStats)
}

func (rsc *RibbonStatsController) getStats(c *gin.Context) {
	stats, err := rsc.ribbonStatsService.GetStats()
	if err != nil {
		c.JSON(200, gin.H{"errors": "Something went wrong"})
	}
	c.JSON(200, gin.H{"data": stats})
}
