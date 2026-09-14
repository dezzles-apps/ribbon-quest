package ribbons

import (
	"dezzles-apps/rq-server/model/dto"
	"errors"

	"github.com/dezzles-apps/go-common/db"
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

func (rss *RibbonStatsService) GetStats() (*dto.RibbonStats, error) {
	pokemon, err := rss.pokemonService.GetAllPokemon()
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
