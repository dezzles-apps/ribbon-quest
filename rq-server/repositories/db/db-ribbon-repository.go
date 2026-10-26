package db

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/interfaces"
	_ "embed"

	"github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"
)

//go:embed ribbons/get-ribbon.sql
var getRibbonQuery string

type DBRibbonRepository struct {
	connection *db.Database
	common     interfaces.CommonDataRepository
}

func NewDatabaseRibbonRepository(
	connection *db.Database,
	common interfaces.CommonDataRepository,
) interfaces.RibbonRepository {
	return &DBRibbonRepository{
		connection: connection,
		common:     common,
	}
}

func (r *DBRibbonRepository) AddRibbon(ctx *cmodel.DAContext, ribbonPokemonId int, ribbon string) (*dto.PokemonRibbon, errors.Error) {
	ctx.Logger.Info("Adding earned ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon))
	_, err := r.connection.GetDB().Exec("INSERT INTO ribbons_earned (ribbon_pokemon_id, ribbon_key) VALUES (?, ?)", ribbonPokemonId, ribbon)
	if err != nil {
		ctx.Logger.Error("Error adding ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon), zap.Error(err))
		return nil, model.InternalServerError
	}
	ctx.Logger.Info("Added earned ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon))
	return r.GetRibbon(ctx, ribbonPokemonId, ribbon)
}

func (r *DBRibbonRepository) GetRibbon(ctx *cmodel.DAContext, ribbonPokemonId int, ribbon string) (*dto.PokemonRibbon, errors.Error) {
	ctx.Logger.Info("Retrieving earned ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon))
	ribbonData := &dto.PokemonRibbon{}
	err := r.connection.GetDB().QueryRow(getRibbonQuery, ribbonPokemonId, ribbon).Scan(
		&ribbonData.RibbonKey,
		&ribbonData.Name,
		&ribbonData.Achieved,
		&ribbonData.Category,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.InvalidRibbonCombination
		}
		ctx.Logger.Error("Error retrieving ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon), zap.Error(err))
		return nil, model.InternalServerError
	}
	ctx.Logger.Info("Retrieved earned ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon))

	return ribbonData, nil
}

func (r *DBRibbonRepository) RemoveRibbon(ctx *cmodel.DAContext, ribbonPokemonId int, ribbon string) (*dto.PokemonRibbon, errors.Error) {
	ctx.Logger.Info("Removing earned ribbon", zap.Int("ribbonPokemonId", ribbonPokemonId), zap.String("ribbon", ribbon))
	_, err := r.connection.GetDB().Exec("DELETE FROM ribbons_earned WHERE ribbon_pokemon_id = ? AND ribbon_key = ?", ribbonPokemonId, ribbon)
	if err != nil {
		ctx.Logger.Error("Error removing ribbon", zap.Int("pokemonId", ribbonPokemonId), zap.String("ribbon", ribbon), zap.Error(err))
		return nil, errors.InternalServerError
	}
	ctx.Logger.Info("Removing earned ribbon", zap.Int("pokemonId", ribbonPokemonId), zap.String("ribbon", ribbon))
	return r.GetRibbon(ctx, ribbonPokemonId, ribbon)
}
