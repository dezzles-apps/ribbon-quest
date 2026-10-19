package repositories

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"
	_ "embed"
	"fmt"

	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"

	"github.com/dezzles-apps/go-common/db"
	"github.com/dezzles-apps/go-common/model/responses"
)

//go:embed "sql/get-all-shinies.sql"
var getAllShinies string

type ShinyRepository interface {
	GetAllShinies(ctx *cmodel.DAContext) ([]dto.ShinyPokemon, *responses.Error)
}

type DatabaseShinyRepository struct {
	connection *db.Database
	common     CommonDataRepository
}

func NewDatabaseShinyRepository(
	connection *db.Database,
	common CommonDataRepository,
) ShinyRepository {
	return &DatabaseShinyRepository{
		connection: connection,
		common:     common,
	}
}

func (repo *DatabaseShinyRepository) GetAllShinies(ctx *cmodel.DAContext) ([]dto.ShinyPokemon, *responses.Error) {
	shinies := make([]dto.ShinyPokemon, 0)
	rows, err := repo.connection.GetDB().Query(getAllShinies)
	if err != nil {
		ctx.Logger.Info("Error getting shinies", zap.Error(err))
		return nil, &errors.InternalServerError
	}
	if rows.Err() != nil {
		ctx.Logger.Info("Error getting shiny rows", zap.Error(err))
		return nil, &errors.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var entry dto.ShinyPokemon
		var pokedexNo int
		var format string = "#%04d"
		var gameKey string
		var natureKey string
		var characteristicKey string
		var methodKey string
		var genderKey string
		err = rows.Scan(&entry.ShinyId, &pokedexNo, &entry.Species.Name, &entry.Species.Form, &entry.Species.ImageRef,
			&entry.Details.CaughtAt, &natureKey, &characteristicKey, &entry.Details.Level,
			&entry.Details.Notes, &gameKey, &methodKey, &genderKey,
		)
		if err != nil {
			ctx.Logger.Info("Error getting shinies", zap.Error(err))
			return nil, &errors.InternalServerError
		}

		entry.Species.PokedexId = fmt.Sprintf(format, pokedexNo)

		// Map Game
		if game, err := repo.common.GetGame(gameKey); err != nil {
			ctx.Logger.Warn("Failed to map game", zap.String("gameKey", gameKey), zap.Error(err))
		} else {
			entry.Details.Game = game
		}

		// Map Nature
		if nature, err := repo.common.GetNature(natureKey); err != nil {
			ctx.Logger.Warn("Failed to map nature", zap.String("natureKey", natureKey), zap.Error(err))
		} else {
			entry.Details.Nature = *nature
		}

		// Map Characteristic
		if characteristic, err := repo.common.GetCharacteristic(characteristicKey); err != nil {
			ctx.Logger.Warn("Failed to map characteristic", zap.String("characteristicKey", characteristicKey), zap.Error(err))
		} else {
			entry.Details.Characteristic = *characteristic
		}

		// Map Method
		if method, err := repo.common.GetMethod(methodKey); err != nil {
			ctx.Logger.Warn("Failed to map method", zap.String("methodKey", methodKey), zap.Error(err))
		} else {
			entry.Details.Method = method
		}

		// Map Gender
		if gender, err := repo.common.GetGender(genderKey); err != nil {
			ctx.Logger.Warn("Failed to map gender", zap.String("genderKey", genderKey), zap.Error(err))
		} else {
			entry.Details.Gender = gender
		}

		shinies = append(shinies, entry)
	}
	return shinies, nil
}
