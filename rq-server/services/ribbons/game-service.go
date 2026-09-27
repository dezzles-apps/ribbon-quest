package ribbons

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"
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

func (gs *GameService) GetGame(ctx *cmodel.DAContext, gameName string) (*dto.Game, error) {
	return gs.getGame(ctx, gameName)
}

func (gs *GameService) getGame(ctx *cmodel.DAContext, gameName string) (*dto.Game, error) {
	game := &dto.Game{}
	ctx.Logger.Info("Retrieving game data", zap.String("gameName", gameName))
	row := gs.connection.GetDB().QueryRow(getGameInfoQuery, gameName)
	viewOrder := 0
	err := row.Scan(&game.GameKey, &game.Name, &viewOrder)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.GameNotFound
		}
		ctx.Logger.Error("Error retrieving game data", zap.String("gameName", gameName), zap.Error(err))
		return nil, model.InternalServerError
	}
	pokemon, err := gs.getPokemonByGame(ctx, gameName)
	if err != nil {
		return nil, err
	}
	ctx.Logger.Info("Retrieved all data for game", zap.String("gameName", gameName))
	game.Pokemon = pokemon
	return game, nil
}

func (gs *GameService) getPokemonByGame(ctx *cmodel.DAContext, gameName string) ([]*dto.GamePokemon, error) {
	pokemonList := []*dto.GamePokemon{}
	ctx.Logger.Info("Retrieving pokemon for game", zap.String("gameName", gameName))
	rows, err := gs.connection.GetDB().Query(getGamePokemonQuery, gameName)
	if err != nil {
		ctx.Logger.Error("Error retrieving pokemon for game", zap.String("gameName", gameName), zap.Error(err))
		return nil, model.InternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var caughtAt sql.NullTime
		var nature sql.NullString
		var characteristic sql.NullString
		var notes sql.NullString
		var shiny sql.NullBool
		pokemon := &dto.GamePokemon{}
		err := rows.Scan(
			&pokemon.Pokemon,
			&pokemon.Species.PokedexId,
			&pokemon.Species.Name,
			&pokemon.Species.Form,
			&pokemon.Species.ImageRef,
			&pokemon.Details.Nickname,
			&caughtAt,
			&nature,
			&characteristic,
			&notes,
			&shiny,
		)
		if err != nil {
			ctx.Logger.Error("Error retrieving pokemon row for game", zap.String("gameName", gameName), zap.Error(err))
			return nil, model.InternalServerError
		}
		if caughtAt.Valid {
			pokemon.Details.CaughtAt = &caughtAt.Time
		}
		if nature.Valid {
			pokemon.Details.Nature = nature.String
		}
		if characteristic.Valid {
			pokemon.Details.Characteristic = characteristic.String
		}
		if shiny.Valid {
			pokemon.Details.Shiny = shiny.Bool
		}
		if notes.Valid {
			pokemon.Details.Notes = notes.String
		}
		pokemonList = append(pokemonList, pokemon)
	}
	if err = rows.Err(); err != nil {
		ctx.Logger.Error("Error retrieving pokemon rows for game", zap.String("gameName", gameName), zap.Error(err))
		return nil, model.InternalServerError
	}
	err = gs.loadRibbonsForPokemon(ctx, gameName, pokemonList)
	if err != nil {
		return nil, err
	}
	ctx.Logger.Info("Retrieved pokemon for game", zap.String("gameName", gameName))
	return pokemonList, nil
}

func (gs *GameService) loadRibbonsForPokemon(ctx *cmodel.DAContext, gameName string, pokemon []*dto.GamePokemon) error {
	rows, err := gs.connection.GetDB().Query(getGameRibbons, gameName)
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
			ctx.Logger.Error("Error retrieving pokemon ribbons for game", zap.String("gameName", gameName), zap.Error(err))
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
		ctx.Logger.Error("Error retrieving pokemon ribbons for game", zap.String("gameName", gameName), zap.Error(err))
		return model.InternalServerError
	}
	ctx.Logger.Info("Retrieved pokemon ribbons for game", zap.String("gameName", gameName))
	return nil
}

func (gs *GameService) GetAllGames(ctx *cmodel.DAContext) ([]*dto.GameWithStats, error) {
	games := []*dto.GameWithStats{}
	rows, err := gs.connection.GetDB().Query(getAllGamesQuery)
	if err != nil {
		ctx.Logger.Error("Error retrieving all games", zap.Error(err))
		return nil, model.InternalServerError
	}
	defer rows.Close()
	var allGames map[string]*dto.GameWithStats = make(map[string]*dto.GameWithStats)

	for rows.Next() {
		game := &dto.GameWithStats{}
		var achievedRow bool
		var total int
		err := rows.Scan(&game.GameKey, &game.Name, &achievedRow, &total)
		if err != nil {
			ctx.Logger.Error("GetAllGames: error scanning games", zap.Error(err))
			return nil, err
		}
		if _, exists := allGames[game.GameKey]; !exists {
			allGames[game.GameKey] = game
			games = append(games, game)
		}
		existingGame := allGames[game.GameKey]
		if achievedRow {
			existingGame.Achieved += total
		}
		existingGame.Total += total

	}
	if err = rows.Err(); err != nil {
		ctx.Logger.Error("Error retrieving all games", zap.Error(err))
		return nil, model.InternalServerError
	}
	ctx.Logger.Info("Successfully retrieved all games")
	return games, nil
}
