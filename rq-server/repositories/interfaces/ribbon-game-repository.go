package interfaces

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"

	cmodel "github.com/dezzles-apps/go-common/model"
)

type RibbonGameRepository interface {
	GetGame(ctx *cmodel.DAContext, gameName string) (*dto.Game, errors.Error)
	GetPokemonByGame(ctx *cmodel.DAContext, gameName string) ([]*dto.GamePokemon, errors.Error)
	GetAllGames(ctx *cmodel.DAContext) ([]*dto.GameWithStats, errors.Error)
}
