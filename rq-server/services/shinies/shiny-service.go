package shinies

import (
	"dezzles-apps/rq-server/model/dto"
	request "dezzles-apps/rq-server/model/dto/shinies"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories"
	data "dezzles-apps/rq-server/services/data"

	cmodel "github.com/dezzles-apps/go-common/model"

	"github.com/dezzles-apps/go-common/model/responses"
)

type ShinyService struct {
	commonDataService  *data.CommonDataService
	pokemonDataService *data.PokemonDataService
	shinyRepository    repositories.ShinyRepository
}

func NewShinyService(
	commonDataService *data.CommonDataService,
	pokemonDataService *data.PokemonDataService,
	shinyRepository repositories.ShinyRepository,
) *ShinyService {
	return &ShinyService{
		commonDataService:  commonDataService,
		pokemonDataService: pokemonDataService,
		shinyRepository:    shinyRepository,
	}
}

func (svc *ShinyService) GetShinies(ctx *cmodel.DAContext) ([]dto.ShinyPokemon, *responses.Error) {
	return svc.shinyRepository.GetAllShinies(ctx)
}

func (svc *ShinyService) ValidateRequest(request request.CreateShinyRequest) []responses.Error {
	results := make([]responses.Error, 0)

	if request.Level < 1 || request.Level > 100 {
		results = append(results, errors.InvalidLevel)
	}

	gameErr := svc.commonDataService.ValidGame(request.Game)
	if gameErr != nil {
		if gameErr == &errors.InternalServerError {
			return []responses.Error{*gameErr}
		}
		results = append(results, *gameErr)
	}

	methodErr := svc.commonDataService.ValidMethod(request.Method)
	if methodErr != nil {
		if methodErr == &errors.InternalServerError {
			return []responses.Error{*methodErr}
		}
		results = append(results, *methodErr)
	}

	characteristicErr := svc.commonDataService.ValidCharacteristic(request.Characteristic)
	if characteristicErr != nil {
		if characteristicErr == &errors.InternalServerError {
			return []responses.Error{*characteristicErr}
		}
		results = append(results, *characteristicErr)
	}

	natureErr := svc.commonDataService.ValidNature(request.Nature)
	if natureErr != nil {
		if natureErr == &errors.InternalServerError {
			return []responses.Error{*natureErr}
		}
		results = append(results, *natureErr)
	}

	return results
}
