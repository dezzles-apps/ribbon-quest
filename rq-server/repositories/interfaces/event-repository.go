package interfaces

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"
	"time"

	cmodel "github.com/dezzles-apps/go-common/model"
)

type EventRepository interface {
	GetEvents(ctx *cmodel.DAContext, date *time.Time, category string, pokemon string) ([]dto.Event, errors.Error)
}
