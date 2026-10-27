package ribbons

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/interfaces"

	_ "embed"

	"github.com/dezzles-apps/go-common/model"
	cmodel "github.com/dezzles-apps/go-common/model"
)

type GameService struct {
	ribbonGameRepo interfaces.RibbonGameRepository
}

func NewGameService(
	ribbonGameRepo interfaces.RibbonGameRepository,
) *GameService {
	return &GameService{
		ribbonGameRepo: ribbonGameRepo,
	}
}

func (gs *GameService) GetGame(ctx *cmodel.DAContext, gameName string) (*dto.Game, errors.Error) {
	return gs.ribbonGameRepo.GetGame(ctx, gameName)
}

func (gs *GameService) GetAllGames(ctx *model.DAContext) ([]*dto.GameWithStats, errors.Error) {
	return gs.ribbonGameRepo.GetAllGames(ctx)
}
