package controllers

import (
	"dezzles-apps/rq-server/services"
	"log"
	"time"

	"dezzles-apps/rq-server/model/dto"

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
	date := c.Query("key")
	events, err := ec.eventService.GetLatestEvents(date)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var result = dto.Page[dto.Event]{
		Items: events,
	}
	if len(events) > 0 {
		next := events[len(events)-1].EventTime.Format(time.RFC3339)
		result.Next = &next
	}
	log.Printf("Event Count: %d", len(events))
	c.JSON(200, gin.H{"data": result})
}
