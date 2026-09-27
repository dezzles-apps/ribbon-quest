package ribbons

import (
	"dezzles-apps/rq-server/model/dto"
	"errors"

	"github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
)

type RibbonStatsService struct {
	connection     *db.Database
	pokemonService *PokemonService
}

func NewRibbonStatsService(
	connection *db.Database,
	pokemonService *PokemonService,
) *RibbonStatsService {
	return &RibbonStatsService{
		connection:     connection,
		pokemonService: pokemonService,
	}
}

func (rss *RibbonStatsService) GetStats(ctx *cmodel.DAContext) (*dto.RibbonStats, error) {
	pokemon, err := rss.pokemonService.GetAllPokemon(ctx)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}
	response := dto.RibbonStats{
		Pokemon: pokemon,
	}
	for _, p := range pokemon {
		response.Current += p.Current
		response.Total += p.Total
	}

	return &response, nil
}
