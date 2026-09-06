package controllers

import (
	"dezzles-apps/rq-server/services"

	"github.com/gin-gonic/gin"
)

type EventsController struct {
	eventService *services.EventService
}

func NewEventsController(
	router *gin.Engine,
	eventService services.EventService,
) {
	controller := EventsController{
		eventService: &eventService,
	}
	controller.registerRoutes(router)
}

func (ec *EventsController) registerRoutes(router *gin.Engine) {
	router.GET("/api/events/v1", ec.getEvents)
}

func (ec *EventsController) getEvents(c *gin.Context) {
	events, err := ec.eventService.GetLatestEvents()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": events})
}
