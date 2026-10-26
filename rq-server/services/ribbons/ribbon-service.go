package ribbons

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/interfaces"
	"log"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
)

type RibbonService struct {
	connection       *cdb.Database
	pokemonService   PokemonService
	ribbonRepository interfaces.RibbonRepository
}

func NewRibbonService(
	connection *cdb.Database,
	PokemonService *PokemonService,
	ribbonRepository interfaces.RibbonRepository,
) *RibbonService {
	return &RibbonService{
		connection:       connection,
		pokemonService:   *PokemonService,
		ribbonRepository: ribbonRepository,
	}
}

func (rs *RibbonService) AddRibbon(ctx *cmodel.DAContext, pokemon string, ribbon string) (*dto.PokemonRibbon, errors.Error) {
	log.Printf("AddRibbon: Adding ribbon %s to %s", ribbon, pokemon)
	pokemonId, err := rs.pokemonService.GetPokemonId(ctx, pokemon)
	if err != nil {
		log.Printf("AddRibbon: Error adding ribbon %s to %s. %s", ribbon, pokemon, err.Error())
		return nil, model.InternalServerError
	}
	r, err := rs.ribbonRepository.GetRibbon(ctx, pokemonId, ribbon)
	if err != nil {
		return nil, err
	}
	if r.Achieved {
		return r, nil
	}
	return rs.ribbonRepository.AddRibbon(ctx, pokemonId, ribbon)
}

func (rs *RibbonService) RemoveRibbon(ctx *cmodel.DAContext, pokemon string, ribbon string) (*dto.PokemonRibbon, errors.Error) {
	pokemonId, err := rs.pokemonService.GetPokemonId(ctx, pokemon)
	if err != nil {
		return nil, err
	}

	_, err = rs.ribbonRepository.GetRibbon(ctx, pokemonId, ribbon)
	if err != nil {
		return nil, err
	}

	return rs.ribbonRepository.RemoveRibbon(ctx, pokemonId, ribbon)

}
