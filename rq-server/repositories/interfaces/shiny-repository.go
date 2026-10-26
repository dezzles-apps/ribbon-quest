package interfaces

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/dto/shinies"
	"dezzles-apps/rq-server/model/errors"

	cmodel "github.com/dezzles-apps/go-common/model"
)

type ShinyRepository interface {
	GetAllShinies(ctx *cmodel.DAContext) ([]dto.ShinyPokemon, errors.Error)
	GetShiniesByDateRange(ctx *cmodel.DAContext, startDate string, endDate string) ([]dto.ShinyPokemon, errors.Error)
	SaveShiny(ctx *cmodel.DAContext, request shinies.CreateShinyRequest) errors.Error
}
