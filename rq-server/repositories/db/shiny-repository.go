package db

import (
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/dto/data"
	"dezzles-apps/rq-server/model/dto/shinies"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/interfaces"
	_ "embed"
	"fmt"

	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"

	"github.com/dezzles-apps/go-common/db"
)

//go:embed "shinies/get-all-shinies.sql"
var getAllShinies string

//go:embed "shinies/get-shinies-by-date-range.sql"
var getShiniesByDateRange string

type DatabaseShinyRepository struct {
	connection *db.Database
	common     interfaces.CommonDataRepository
}

func NewDatabaseShinyRepository(
	connection *db.Database,
	common interfaces.CommonDataRepository,
) interfaces.ShinyRepository {
	return &DatabaseShinyRepository{
		connection: connection,
		common:     common,
	}
}

func (repo *DatabaseShinyRepository) SaveShiny(ctx *cmodel.DAContext, request shinies.CreateShinyRequest) errors.Error {
	// Get form_id from the form sprite_ref
	var formId int
	row := repo.connection.GetDB().QueryRow(
		"SELECT form_id FROM pokemon_forms WHERE sprite_ref = ?",
		request.Form,
	)
	err := row.Scan(&formId)
	if err != nil {
		ctx.Logger.Error("Error getting form ID", zap.String("form", request.Form), zap.Error(err))
		return errors.InternalServerError
	}

	// Insert shiny into database
	result, err := repo.connection.GetDB().Exec(
		`INSERT INTO shinies (form_id, level, game, method, characteristic, nature, gender, caught_at, notes) 
		 VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, ?)`,
		formId,
		request.Level,
		request.Game,
		request.Method,
		request.Characteristic,
		request.Nature,
		request.Gender,
		request.Notes,
	)

	if err != nil {
		ctx.Logger.Error("Error saving shiny",
			zap.String("form", request.Form),
			zap.Int("formId", formId),
			zap.Int("level", request.Level),
			zap.String("game", request.Game),
			zap.Error(err),
		)
		return errors.InternalServerError
	}

	shinyId, err := result.LastInsertId()
	if err != nil {
		ctx.Logger.Error("Error retrieving shiny ID after insert", zap.Error(err))
		return errors.InternalServerError
	}

	ctx.Logger.Info("Shiny successfully saved", zap.Int64("shinyId", shinyId), zap.String("form", request.Form))
	return nil
}

func (repo *DatabaseShinyRepository) GetAllShinies(ctx *cmodel.DAContext) ([]dto.ShinyPokemon, errors.Error) {
	shinies := make([]dto.ShinyPokemon, 0)
	rows, err := repo.connection.GetDB().Query(getAllShinies)
	if err != nil {
		ctx.Logger.Info("Error getting shinies", zap.Error(err))
		return nil, errors.InternalServerError
	}
	if rows.Err() != nil {
		ctx.Logger.Info("Error getting shiny rows", zap.Error(err))
		return nil, errors.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var entry dto.ShinyPokemon
		entry.Details.Shiny = true
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
			return nil, errors.InternalServerError
		}

		entry.Species.PokedexNo.Number = pokedexNo
		entry.Species.PokedexNo.Text = fmt.Sprintf(format, pokedexNo)

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
	if err := rows.Err(); err != nil {
		ctx.Logger.Info("Error getting shiny rows", zap.Error(err))
		return nil, errors.InternalServerError
	}
	return shinies, nil
}

func (repo *DatabaseShinyRepository) GetShiniesByDateRange(ctx *cmodel.DAContext, startDate string, endDate string) ([]dto.ShinyPokemon, errors.Error) {
	shinies := make([]dto.ShinyPokemon, 0)
	rows, err := repo.connection.GetDB().Query(getShiniesByDateRange, startDate, endDate)
	if err != nil {
		ctx.Logger.Info("Error getting shinies by date range", zap.String("startDate", startDate), zap.String("endDate", endDate), zap.Error(err))
		return nil, errors.InternalServerError
	}
	if rows.Err() != nil {
		ctx.Logger.Info("Error getting shiny rows", zap.Error(err))
		return nil, errors.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var entry dto.ShinyPokemon
		entry.Details.Shiny = true
		var pokedexNo int
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
			ctx.Logger.Info("Error scanning shiny from date range query", zap.Error(err))
			return nil, errors.InternalServerError
		}

		entry.Species.PokedexNo = data.CreatePokedexNumber(pokedexNo, "#", 4)

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
	if err := rows.Err(); err != nil {
		ctx.Logger.Info("Error getting shiny rows", zap.Error(err))
		return nil, errors.InternalServerError
	}
	return shinies, nil
}
