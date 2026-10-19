package controllers

import (
	"dezzles-apps/rq-server/interfaces"
	"dezzles-apps/rq-server/services"
	"time"

	"dezzles-apps/rq-server/model/dto"

	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EventsController struct {
	eventService *services.EventService
}

func (ec *EventsController) Initialise(
	ctx interfaces.ControllerInitContext,
) {
	controller := EventsController{
		eventService: ctx.Services().Events,
	}
	controller.registerRoutes(ctx)
}

func (ec *EventsController) registerRoutes(ctx interfaces.ControllerInitContext) {
	ctx.Router().GET("/api/events/v1", ec.getEvents)
}

func (ec *EventsController) getEvents(c *gin.Context) {
	ctx := cmodel.GetContext(c)
	date := c.Query("key")
	category := c.Query("category")
	pokemon := c.Query("pokemon")
	events, err := ec.eventService.GetEvents(ctx, date, category, pokemon)
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
	ctx.Logger.Info("Retrieved events", zap.Int("eventCount", len(events)))
	c.JSON(200, gin.H{"data": result})
}
