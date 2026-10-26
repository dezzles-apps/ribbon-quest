package services

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/repositories/interfaces"
	"time"

	_ "embed"

	"dezzles-apps/rq-server/model/errors"

	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"
)

type EventService struct {
	eventRepository interfaces.EventRepository
}

func NewEventService(eventRepository interfaces.EventRepository) *EventService {
	return &EventService{
		eventRepository: eventRepository,
	}
}

func (es *EventService) GetEvents(ctx *cmodel.DAContext, dateStr string, category string, pokemon string) ([]dto.Event, errors.Error) {
	var date *time.Time
	if dateStr != "" {
		parsedDate, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			ctx.Logger.Info("Invalid date provided", zap.String("date", dateStr))
			return nil, errors.InvalidKey
		}
		date = &parsedDate
	}
	return es.eventRepository.GetEvents(ctx, date, category, pokemon)
}
