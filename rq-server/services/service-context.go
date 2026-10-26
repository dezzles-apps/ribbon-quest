package services

import (
	"dezzles-apps/rq-server/repositories"
	"dezzles-apps/rq-server/services/data"
	"dezzles-apps/rq-server/services/pokedexes"
	"dezzles-apps/rq-server/services/ribbons"
	"dezzles-apps/rq-server/services/shinies"

	"github.com/dezzles-apps/go-common/db"
)

type ServiceContext struct {
	Users  *UserService
	Events *EventService
	Config *ConfigService

	CommonData  *data.CommonDataService
	PokemonData *data.PokemonDataService
	Pokedex     *pokedexes.PokedexService

	RibbonGames   *ribbons.GameService
	RibbonPokemon *ribbons.PokemonService
	Ribbons       *ribbons.RibbonService
	RibbonStats   *ribbons.RibbonStatsService

	Shinies *shinies.ShinyService
}

func (ctx *ServiceContext) Initialise(database *db.Database, repos repositories.RepositoryContext) {
	ctx.Config = NewConfigService(database)
	ctx.Users = NewUserService(database, ctx.Config)
	ctx.Events = NewEventService(database)

	ctx.CommonData = data.NewCommonDataService(repos.CommonData)
	ctx.PokemonData = data.NewPokemonDataService(database)
	ctx.Pokedex = pokedexes.NewPokedexService(database)

	ctx.RibbonGames = ribbons.NewGameService(database)
	ctx.RibbonPokemon = ribbons.NewPokemonService(
		database,
		ctx.RibbonGames,
		ctx.PokemonData,
	)
	ctx.Ribbons = ribbons.NewRibbonService(
		database,
		ctx.RibbonPokemon,
		repos.Ribbons,
	)
	ctx.RibbonStats = ribbons.NewRibbonStatsService(
		database,
		ctx.RibbonPokemon,
	)
	ctx.Shinies = shinies.NewShinyService(
		ctx.CommonData,
		ctx.PokemonData,
		repos.Shiny,
	)
}
