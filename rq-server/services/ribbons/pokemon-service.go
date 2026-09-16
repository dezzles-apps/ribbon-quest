package ribbons

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	dataservices "dezzles-apps/rq-server/services/data"
	_ "embed"
	"errors"
	"log"
	"strings"

	cdb "github.com/dezzles-apps/go-common/db"
	"github.com/dezzles-apps/go-common/model/responses"
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

func (ps *PokemonService) GetPokemon(pokemonName string) (*dto.Pokemon, error) {
	Pokemon, err := ps.getPokemon(pokemonName)
	if err != nil {
		return nil, err
	}
	games, err := ps.getPokemonGames(pokemonName)
	if err != nil {
		return nil, err
	}
	Pokemon.Games = games
	ribbons, err := ps.getPokemonRibbons(pokemonName)
	if err != nil {
		return nil, err
	}
	Pokemon.Ribbons = ribbons
	return Pokemon, nil
}

func (ps *PokemonService) getPokemon(pokemonName string) (*dto.Pokemon, error) {
	Pokemon := &dto.Pokemon{}
	row := ps.connection.GetDB().QueryRow(getPokemonInfo, pokemonName)
	var caughtAt sql.NullTime
	var nature sql.NullString
	var notes sql.NullString
	var characteristic sql.NullString
	log.Println("Reading getPokemon")
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
			return nil, model.PokemonNotFound
		}
		log.Printf("Err getPokemon %s", err.Error())
		return nil, err
	}
	log.Println("Read getPokemon")
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
	return Pokemon, nil
}

func (ps *PokemonService) getPokemonGames(pokemonName string) ([]dto.PokemonGame, error) {
	var gamesMap map[string]*dto.PokemonGame = make(map[string]*dto.PokemonGame)
	rows, err := ps.connection.GetDB().Query(getPokemonGames, pokemonName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gameKey string
		var gameName string
		var viewOrder int
		var ribbonKey string
		err := rows.Scan(&gameKey, &gameName, &viewOrder, &ribbonKey)
		if err != nil {
			return nil, err
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

	return games, nil
}

func (ps *PokemonService) getPokemonRibbons(pokemonName string) ([]dto.PokemonRibbon, error) {
	var ribbons []dto.PokemonRibbon
	rows, err := ps.connection.GetDB().Query(getPokemonRibbons, pokemonName)
	if err != nil {
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
	return ribbons, nil
}

func (ps *PokemonService) GetAllPokemon() ([]dto.AllPokemon, error) {
	log.Println("Retrieving all pokemon")
	var allPokemon []dto.AllPokemon
	rows, err := ps.connection.GetDB().Query(getAllPokemon)
	if err != nil {
		return nil, err
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
			return nil, err
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
	log.Print("Retrieved all pokemon")
	for _, p := range pokemonList {
		allPokemon = append(allPokemon, *p)
	}
	return allPokemon, nil
}

func (ps *PokemonService) CatchPokemon(pokemon string) (*dto.Pokemon, error) {
	details, err := ps.getPokemon(pokemon)
	if err != nil {
		return nil, err
	}
	if details.Details.CaughtAt != nil {
		return nil, errors.New("Pokemon already caught")
	}
	_, err = ps.connection.GetDB().Exec("UPDATE pokemon SET caught_at = CURRENT_TIMESTAMP WHERE pokemon = ?", pokemon)
	if err != nil {
		return nil, err
	}
	return ps.GetPokemon(pokemon)
}

func (ps *PokemonService) UpdatePokemon(pokemon string, updateData dto.UpdatePokemon) (*dto.Pokemon, error) {
	details, err := ps.getPokemon(pokemon)
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
		"UPDATE pokemon SET nickname = ?, nature = ?, characteristic = ?, notes = ?, shiny = ? WHERE pokemon = ?",
		details.Details.Nickname,
		details.Details.Nature,
		details.Details.Characteristic,
		details.Details.Notes,
		details.Details.Shiny,
		pokemon,
	)
	if err != nil {
		return nil, err
	}
	return ps.GetPokemon(pokemon)
}

func (ps *PokemonService) CreatePokemon(pokemon *dto.AddNewRibbonPokemon) (*dto.Pokemon, error) {
	viewOrder, err := ps.getNextViewOrder()
	if err != nil {
		return nil, err
	}
	formId, err := ps.pokemonDataService.GetFormId(pokemon.PokedexId, pokemon.Form)
	if err != nil {
		return nil, err
	}

	_, err = ps.connection.GetDB().Exec(
		"INSERT INTO pokemon (pokemon, nickname, pokedex_id, form_id, view_order) VALUES(?, ?, ?, ?, ?)",
		pokemon.Pokemon, pokemon.Pokemon, pokemon.PokedexId, formId, viewOrder,
	)
	if err != nil {
		return nil, err
	}

	err = ps.saveGames(pokemon)
	if err != nil {
		return nil, err
	}

	return ps.getPokemon(pokemon.Pokemon)
}

func (ps *PokemonService) ValidateCreate(pokemon *dto.AddNewRibbonPokemon) ([]responses.Error, error) {
	var errors []responses.Error

	// Trim everything
	pokemon.Pokemon = strings.TrimSpace(pokemon.Pokemon)
	pokemon.Form = strings.TrimSpace(pokemon.Form)
	validation, err := ps.validatePokemonId(pokemon.Pokemon)
	if err != nil {
		return nil, err
	}
	if validation != nil {
		errors = append(errors, *validation)
	}
	validation, err = ps.validatePokemon(pokemon.PokedexId, pokemon.Form)
	if err != nil {
		return nil, err
	}
	if validation != nil {
		errors = append(errors, *validation)
	}
	validation, err = ps.validateGames(pokemon.Games)
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

func (ps *PokemonService) validatePokemonId(pokemonId string) (*responses.Error, error) {
	if pokemonId == "" {
		e := responses.CreateError("pokemon", "PokemonId cannot be empty")
		return &e, nil
	} else {
		existing, err := ps.getPokemon(pokemonId)
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

func (ps *PokemonService) validatePokemon(pokedexNo int, form string) (*responses.Error, error) {
	pokemon, err := ps.pokemonDataService.GetPokemon(pokedexNo)
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

func (ps *PokemonService) validateGames(games []string) (*responses.Error, error) {
	log.Print("validateGames: getting games")
	allGames, err := ps.gameService.GetAllGames()
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

func (ps *PokemonService) getNextViewOrder() (int, error) {
	var number int
	row := ps.connection.GetDB().QueryRow("SELECT view_order FROM pokemon ORDER BY view_order DESC LIMIT 1")
	err := row.Scan(&number)
	if err != nil {
		return 0, err
	}
	return number + 1, nil
}

func (ps *PokemonService) saveGames(pokemon *dto.AddNewRibbonPokemon) error {
	var query = "INSERT INTO pokemon_games (pokemon, game_key) VALUES "
	var params []interface{}
	var first = true
	for _, game := range pokemon.Games {
		if !first {
			query += ", "
		}
		query += "(?, ?)"
		params = append(params, pokemon.Pokemon)
		params = append(params, game)

		first = false
	}
	_, err := ps.connection.GetDB().Exec(query, params...)
	return err

}
