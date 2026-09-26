package ribbons

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	dataservices "dezzles-apps/rq-server/services/data"
	_ "embed"
	"log"
	"strings"

	cdb "github.com/dezzles-apps/go-common/db"
	"github.com/dezzles-apps/go-common/model/responses"
	"go.uber.org/zap"
)

//go:embed sql/get-pokemon-ribbons.sql
var getPokemonRibbons string

//go:embed sql/get-pokemon-games.sql
var getPokemonGames string

//go:embed sql/get-pokemon-info.sql
var getPokemonInfo string

//go:embed sql/get-all-pokemon.sql
var getAllPokemon string

type PokemonService struct {
	connection         *cdb.Database
	gameService        *GameService
	pokemonDataService *dataservices.PokemonDataService
}

func NewPokemonService(
	connection *cdb.Database,
	gameService *GameService,
	pokemonDataService *dataservices.PokemonDataService,
) *PokemonService {
	return &PokemonService{
		connection:         connection,
		gameService:        gameService,
		pokemonDataService: pokemonDataService,
	}
}

func (ps *PokemonService) GetPokemon(ctx *model.RQContext, pokemonName string) (*dto.Pokemon, error) {
	Pokemon, err := ps.getPokemon(ctx, pokemonName)
	if err != nil {
		return nil, err
	}
	games, err := ps.getPokemonGames(ctx, pokemonName)
	if err != nil {
		return nil, err
	}
	Pokemon.Games = games
	ribbons, err := ps.getPokemonRibbons(ctx, pokemonName)
	if err != nil {
		return nil, err
	}
	Pokemon.Ribbons = ribbons
	return Pokemon, nil
}

func (ps *PokemonService) getPokemon(ctx *model.RQContext, pokemonName string) (*dto.Pokemon, error) {
	Pokemon := &dto.Pokemon{}
	ctx.Logger.Info("Getting ribbon pokemon", zap.String("pokemonId", pokemonName))
	row := ps.connection.GetDB().QueryRow(getPokemonInfo, pokemonName)
	var caughtAt sql.NullTime
	var nature sql.NullString
	var notes sql.NullString
	var characteristic sql.NullString
	err := row.Scan(
		&Pokemon.Pokemon,
		&Pokemon.Species.PokedexId,
		&Pokemon.Species.Name,
		&Pokemon.Species.Form,
		&Pokemon.Species.ImageRef,
		&Pokemon.Details.Nickname,
		&caughtAt,
		&nature,
		&characteristic,
		&notes,
		&Pokemon.Details.Shiny,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Logger.Info("Ribbon pokemon not found", zap.String("pokemonId", pokemonName))
			return nil, model.PokemonNotFound
		}
		ctx.Logger.Error("Error retrieving ribbon pokemon", zap.String("pokemonId", pokemonName), zap.Error(err))
		return nil, err
	}
	if caughtAt.Valid {
		Pokemon.Details.CaughtAt = &caughtAt.Time
	}
	if nature.Valid {
		Pokemon.Details.Nature = nature.String
	}
	if characteristic.Valid {
		Pokemon.Details.Characteristic = characteristic.String
	}
	if notes.Valid {
		Pokemon.Details.Notes = notes.String
	}
	ctx.Logger.Info("Retrieved ribbon pokemon", zap.String("pokemonId", pokemonName))
	return Pokemon, nil
}

func (ps *PokemonService) getPokemonGames(ctx *model.RQContext, pokemonName string) ([]dto.PokemonGame, error) {
	var gamesMap map[string]*dto.PokemonGame = make(map[string]*dto.PokemonGame)
	rows, err := ps.connection.GetDB().Query(getPokemonGames, pokemonName)
	if err != nil {
		ctx.Logger.Error("Error getting games for ribbon pokemon", zap.String("pokemonId", pokemonName), zap.Error(err))
		return nil, model.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var gameKey string
		var gameName string
		var viewOrder int
		var ribbonKey string
		err := rows.Scan(&gameKey, &gameName, &viewOrder, &ribbonKey)
		if err != nil {
			ctx.Logger.Error("Error scanning games for ribbon pokemon", zap.String("pokemonId", pokemonName), zap.Error(err))
			return nil, model.InternalServerError
		}
		if game, exists := gamesMap[gameKey]; exists {
			game.Ribbons = append(game.Ribbons, ribbonKey)
		} else {
			game := &dto.PokemonGame{
				GameKey:   gameKey,
				Name:      gameName,
				ViewOrder: viewOrder,
				Ribbons:   []string{ribbonKey},
			}
			gamesMap[gameKey] = game
		}
	}
	var games []dto.PokemonGame
	for _, game := range gamesMap {
		games = append(games, *game)
	}
	ctx.Logger.Info("Retrieved games for ribbon pokemon", zap.String("pokemonId", pokemonName))
	return games, nil
}

func (ps *PokemonService) getPokemonRibbons(ctx *model.RQContext, pokemonName string) ([]dto.PokemonRibbon, error) {
	var ribbons []dto.PokemonRibbon
	ctx.Logger.Info("Getting ribbons for Pokemon", zap.String("pokemonId", pokemonName))
	rows, err := ps.connection.GetDB().Query(getPokemonRibbons, pokemonName)
	if err != nil {
		ctx.Logger.Error("Error retrieving ribbons for ribbon pokemon", zap.String("pokemonId", pokemonName), zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ribbonKey string
		var name string
		var achieved sql.NullBool
		var achieved_at sql.NullTime
		var category string
		var order int
		err := rows.Scan(&ribbonKey, &name, &achieved, &achieved_at, &category, &order)
		if err != nil {
			ctx.Logger.Error("Error scanning ribbons for ribbon pokemon", zap.String("pokemonId", pokemonName), zap.Error(err))
			return nil, err
		}
		if !achieved.Valid {
			achieved.Bool = false
		}
		ribbon := dto.PokemonRibbon{
			RibbonKey: ribbonKey,
			Name:      name,
			Achieved:  achieved.Bool,
			Category:  category,
		}
		if achieved_at.Valid {
			ribbon.AchievedAt = &achieved_at.Time
		}
		ribbons = append(ribbons, ribbon)
	}
	ctx.Logger.Info("Successfully retrieved ribbons for ribbon pokemon", zap.String("pokemonId", pokemonName))

	return ribbons, nil
}

func (ps *PokemonService) GetAllPokemon(ctx *model.RQContext) ([]dto.AllPokemon, error) {
	ctx.Logger.Info("Retrieving all ribbon pokemon")

	var allPokemon []dto.AllPokemon
	rows, err := ps.connection.GetDB().Query(getAllPokemon)
	if err != nil {
		ctx.Logger.Error("Error Getting all ribbon pokemon", zap.Error(err))
		return nil, model.InternalServerError
	}
	var pokemonMap map[string]*dto.AllPokemon = make(map[string]*dto.AllPokemon)
	var pokemonList []*dto.AllPokemon = make([]*dto.AllPokemon, 0)
	defer rows.Close()
	for rows.Next() {
		var pokemon string
		var nickname string
		var pokedexId int
		var species string
		var form string
		var spriteRef string
		var caughtAt sql.NullTime
		var nature sql.NullString
		var characteristic sql.NullString
		var notes sql.NullString
		var shiny sql.NullBool
		var achieved sql.NullBool
		var count int
		err := rows.Scan(&pokemon, &pokedexId, &species, &form, &spriteRef, &nickname, &caughtAt, &nature, &characteristic, &notes, &shiny, &achieved, &count)
		if err != nil {
			ctx.Logger.Error("Error scanning all ribbon pokemon", zap.Error(err))
			return nil, model.InternalServerError
		}
		if !achieved.Valid {
			achieved.Bool = false
		}
		if p, exists := pokemonMap[pokemon]; exists {
			if achieved.Bool {
				p.Current += count
			}
			p.Total += count
		} else {
			p := &dto.AllPokemon{
				Pokemon: pokemon,
				Current: 0,
				Total:   0,
			}
			p.Species.PokedexId = pokedexId
			p.Species.Name = species
			p.Species.Form = form
			p.Species.ImageRef = spriteRef

			p.Details.Nickname = nickname
			p.Details.Nature = nature.String
			p.Details.Characteristic = characteristic.String
			p.Details.Notes = notes.String
			p.Details.Shiny = shiny.Bool

			p.Species.Form = form

			if caughtAt.Valid {
				p.Details.CaughtAt = &caughtAt.Time
			}
			if achieved.Bool {
				p.Current = count
			}
			p.Total = count
			pokemonMap[pokemon] = p
			pokemonList = append(pokemonList, p)
		}
	}
	ctx.Logger.Info("Retrieved all ribbon pokemon")
	for _, p := range pokemonList {
		allPokemon = append(allPokemon, *p)
	}
	return allPokemon, nil
}

func (ps *PokemonService) CatchPokemon(ctx *model.RQContext, pokemon string) (*dto.Pokemon, error) {
	ctx.Logger.Info("Marking pokemon as caught", zap.String("pokemonId", pokemon))
	details, err := ps.getPokemon(ctx, pokemon)
	if err != nil {
		return nil, err
	}
	if details.Details.CaughtAt != nil {
		return nil, model.PokemonAlreadyCaught
	}
	_, err = ps.connection.GetDB().Exec("UPDATE ribbon_pokemon SET caught_at = CURRENT_TIMESTAMP WHERE pokemon = ?", pokemon)
	if err != nil {
		ctx.Logger.Error("Failed to mark pokemon as caught", zap.String("pokemonId", pokemon), zap.Error(err))
		return nil, err
	}
	return ps.GetPokemon(ctx, pokemon)
}

func (ps *PokemonService) UpdatePokemon(ctx *model.RQContext, pokemon string, updateData dto.UpdatePokemon) (*dto.Pokemon, error) {
	details, err := ps.getPokemon(ctx, pokemon)
	if err != nil {
		return nil, err
	}
	if updateData.Nickname != "" {
		details.Details.Nickname = updateData.Nickname
	}
	if updateData.Nature != "" {
		details.Details.Nature = updateData.Nature
	}
	if updateData.Characteristic != "" {
		details.Details.Characteristic = updateData.Characteristic
	}
	if updateData.Shiny != nil {
		details.Details.Shiny = *updateData.Shiny
	}

	if updateData.Notes != "" {
		details.Details.Notes = updateData.Notes
	}

	_, err = ps.connection.GetDB().Exec(
		"UPDATE ribbon_pokemon SET nickname = ?, nature = ?, characteristic = ?, notes = ?, shiny = ? WHERE pokemon = ?",
		details.Details.Nickname,
		details.Details.Nature,
		details.Details.Characteristic,
		details.Details.Notes,
		details.Details.Shiny,
		pokemon,
	)
	if err != nil {
		ctx.Logger.Error("Error updating pokemon data",
			zap.String("pokemonId", pokemon),
			zap.Error(err),
		)
		return nil, err
	}
	return ps.GetPokemon(ctx, pokemon)
}

func (ps *PokemonService) CreatePokemon(ctx *model.RQContext, pokemon *dto.AddNewRibbonPokemon) (*dto.Pokemon, error) {
	viewOrder, err := ps.getNextViewOrder(ctx)
	if err != nil {
		return nil, err
	}
	formId, err := ps.pokemonDataService.GetFormId(ctx, pokemon.PokedexId, pokemon.Form)
	if err != nil {
		return nil, err
	}

	_, err = ps.connection.GetDB().Exec(
		"INSERT INTO ribbon_pokemon (pokemon, nickname, pokedex_id, form_id, view_order) VALUES(?, ?, ?, ?, ?)",
		pokemon.Pokemon, pokemon.Pokemon, pokemon.PokedexId, formId, viewOrder,
	)
	if err != nil {
		ctx.Logger.Error("Failed to create new pokemon",
			zap.String("pokemonId", pokemon.Pokemon),
			zap.Int("pokedexId", pokemon.PokedexId),
			zap.Int("formId", formId),
			zap.Int("viewOrder", viewOrder),
			zap.Error(err),
		)
		return nil, model.InternalServerError
	}

	err = ps.saveGames(ctx, pokemon)
	if err != nil {
		return nil, err
	}

	return ps.getPokemon(ctx, pokemon.Pokemon)
}

func (ps *PokemonService) ValidateCreate(ctx *model.RQContext, pokemon *dto.AddNewRibbonPokemon) ([]responses.Error, error) {
	var errors []responses.Error

	// Trim everything
	pokemon.Pokemon = strings.TrimSpace(pokemon.Pokemon)
	pokemon.Form = strings.TrimSpace(pokemon.Form)
	validation, err := ps.validatePokemonId(ctx, pokemon.Pokemon)
	if err != nil {
		return nil, err
	}
	if validation != nil {
		errors = append(errors, *validation)
	}
	validation, err = ps.validatePokemon(ctx, pokemon.PokedexId, pokemon.Form)
	if err != nil {
		return nil, err
	}
	if validation != nil {
		errors = append(errors, *validation)
	}
	validation, err = ps.validateGames(ctx, pokemon.Games)
	if err != nil {
		return nil, err
	}
	if validation != nil {
		errors = append(errors, *validation)
	}
	if len(errors) > 0 {
		return errors, nil
	}
	return nil, nil
}

func (ps *PokemonService) validatePokemonId(ctx *model.RQContext, pokemonId string) (*responses.Error, error) {
	if pokemonId == "" {
		e := responses.CreateError("pokemon", "PokemonId cannot be empty")
		return &e, nil
	} else {
		existing, err := ps.getPokemon(ctx, pokemonId)
		if existing != nil {
			e := responses.CreateError("pokemon", "PokemonId already in use")
			return &e, nil
		}
		if err != nil && err != model.PokemonNotFound {
			return nil, err
		}
	}
	return nil, nil
}

func (ps *PokemonService) validatePokemon(ctx *model.RQContext, pokedexNo int, form string) (*responses.Error, error) {
	pokemon, err := ps.pokemonDataService.GetPokemon(ctx, pokedexNo)
	log.Print(pokemon)
	if err != nil {
		return nil, err
	}
	if pokemon == nil {
		e := responses.CreateError("pokedexId", "A pokemon must be selected")
		return &e, nil
	}

	formFound := false

	for _, f := range pokemon.Forms {
		if form == f.FormName {
			formFound = true
		}
	}

	if !formFound {
		e := responses.CreateError("form", "Invalid form")
		return &e, nil
	}

	return nil, nil
}

func (ps *PokemonService) validateGames(ctx *model.RQContext, games []string) (*responses.Error, error) {
	log.Print("validateGames: getting games")
	allGames, err := ps.gameService.GetAllGames(ctx)
	if err != nil {
		return nil, err
	}
	var gameMap = make(map[string]*dto.GameWithStats)
	for _, game := range allGames {
		gameMap[game.GameKey] = game
	}
	if len(games) == 0 {
		var r = responses.CreateError("games", "At least one game must be selected")
		return &r, nil
	}
	log.Print(gameMap)

	for _, g := range games {
		if _, exists := gameMap[g]; !exists {
			var r = responses.CreateError("games", "Invalid game: "+g)
			return &r, nil
		}
	}

	log.Print("validateGames: games validated")
	return nil, nil
}

func (ps *PokemonService) getNextViewOrder(ctx *model.RQContext) (int, error) {
	var number int
	row := ps.connection.GetDB().QueryRow("SELECT view_order FROM ribbon_pokemon ORDER BY view_order DESC LIMIT 1")
	err := row.Scan(&number)
	if err != nil {
		ctx.Logger.Error("Error getting next view order", zap.Error(err))
		return 0, model.InternalServerError
	}
	return number + 1, nil
}

func (ps *PokemonService) saveGames(ctx *model.RQContext, pokemon *dto.AddNewRibbonPokemon) error {
	ribbonPokemonId, err := ps.GetPokemonId(ctx, pokemon.Pokemon)
	if err != nil {
		return err
	}
	ctx.Logger.Info("saveGames: Adding games to ribbon pokemon",
		zap.Strings("games", pokemon.Games),
		zap.Int("ribbonPokemonId", ribbonPokemonId),
	)
	log.Printf("saveGames: Adding %s to %d", pokemon.Games, ribbonPokemonId)
	var query = "INSERT INTO ribbon_pokemon_games (ribbon_pokemon_id, game_key) VALUES "
	var params []interface{}
	var first = true
	for _, game := range pokemon.Games {
		if !first {
			query += ", "
		}
		query += "(?, ?)"
		params = append(params, ribbonPokemonId)
		params = append(params, game)

		first = false
	}
	_, err = ps.connection.GetDB().Exec(query, params...)
	if err != nil {
		ctx.Logger.Error(
			"Failed to save games to ribbon pokemon",
			zap.Strings("games", pokemon.Games),
			zap.Int("ribbonPokemonId", ribbonPokemonId),
			zap.Error(err),
		)
	}
	return err
}

func (ps *PokemonService) GetPokemonId(ctx *model.RQContext, pokemon string) (int, error) {
	ctx.Logger.Info("Retrieving pokemon", zap.String("pokemonId", pokemon))
	var query = "SELECT ribbon_pokemon_id FROM ribbon_pokemon WHERE pokemon = ?"
	var number int
	row := ps.connection.GetDB().QueryRow(query, pokemon)
	err := row.Scan(&number)
	if err != nil {
		ctx.Logger.Error("Failed to retrieve pokemon", zap.String("pokemonId", pokemon), zap.Error(err))
		return 0, model.InternalServerError
	}
	return number, nil
}
