package services

import (
	"database/sql"
	"dezzles-apps/rq-server/model/dto"
	"errors"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
)

//go:embed sql/get-game-info.sql
var getGameInfoQuery string

//go:embed sql/get-all-games.sql
var getAllGamesQuery string

//go:embed sql/get-game-pokemon.sql
var getGamePokemonQuery string

//go:embed sql/get-game-ribbons.sql
var getGameRibbons string

type GameService struct {
	connection *cdb.Database
}

func NewGameService(
	connection *cdb.Database,
) *GameService {
	return &GameService{
		connection: connection,
	}
}

func (gs *GameService) GetGame(gameName string) (*dto.Game, error) {
	game, err := gs.getGame(gameName)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (gs *GameService) getGame(gameName string) (*dto.Game, error) {
	game := &dto.Game{}
	row := gs.connection.GetDB().QueryRow(getGameInfoQuery, gameName)
	viewOrder := 0
	err := row.Scan(&game.GameKey, &game.Name, &viewOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("game not found")
		}
		return nil, err
	}
	pokemon, err := gs.getPokemonByGame(gameName)
	if err != nil {
		return nil, err
	}
	game.Pokemon = pokemon
	return game, nil
}

func (gs *GameService) getPokemonByGame(gameName string) ([]*dto.GamePokemon, error) {
	pokemonList := []*dto.GamePokemon{}
	rows, err := gs.connection.GetDB().Query(getGamePokemonQuery, gameName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var caughtAt sql.NullTime
	var nature sql.NullString
	var characteristic sql.NullString
	var shiny sql.NullBool
	for rows.Next() {
		pokemon := &dto.GamePokemon{}
		err := rows.Scan(
			&pokemon.Pokemon,
			&pokemon.Nickname,
			&pokemon.Region,
			&caughtAt,
			&nature,
			&characteristic,
			&shiny,
		)
		if err != nil {
			return nil, err
		}
		if caughtAt.Valid {
			pokemon.CaughtAt = &caughtAt.Time
		}
		if nature.Valid {
			pokemon.Nature = nature.String
		}
		if characteristic.Valid {
			pokemon.Characteristic = characteristic.String
		}
		if shiny.Valid {
			pokemon.Shiny = shiny.Bool
		}
		pokemonList = append(pokemonList, pokemon)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = gs.loadRibbonsForPokemon(gameName, pokemonList)
	if err != nil {
		return nil, err
	}
	return pokemonList, nil
}

func (gs *GameService) loadRibbonsForPokemon(game string, pokemon []*dto.GamePokemon) error {
	rows, err := gs.connection.GetDB().Query(getGameRibbons, game)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonName string
		var ribbon dto.PokemonRibbon
		var achievedAt sql.NullTime
		err := rows.Scan(&pokemonName, &ribbon.RibbonKey, &ribbon.Name, &ribbon.Achieved, &achievedAt, &ribbon.Category)
		if err != nil {
			return err
		}
		if achievedAt.Valid {
			ribbon.AchievedAt = &achievedAt.Time
		}
		for _, p := range pokemon {
			if p.Pokemon == pokemonName {
				p.Ribbons = append(p.Ribbons, ribbon)
				break
			}
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (gs *GameService) GetAllGames() ([]*dto.GameWithStats, error) {
	games := []*dto.GameWithStats{}
	rows, err := gs.connection.GetDB().Query(getAllGamesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allGames map[string]*dto.GameWithStats = make(map[string]*dto.GameWithStats)

	for rows.Next() {
		game := &dto.GameWithStats{}
		err := rows.Scan(&game.GameKey, &game.Name, &game.Achieved, &game.Total)
		if err != nil {
			return nil, err
		}
		if _, exists := allGames[game.GameKey]; !exists {
			allGames[game.GameKey] = game
			games = append(games, game)
		} else {
			existingGame := allGames[game.GameKey]
			existingGame.Achieved += game.Achieved
			existingGame.Total += game.Total
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}
