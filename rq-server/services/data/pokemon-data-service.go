package data

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	dto "dezzles-apps/rq-server/model/dto/data"
	derrors "dezzles-apps/rq-server/model/errors"

	_ "embed"

	"github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"github.com/dezzles-apps/go-common/model/responses"
	"go.uber.org/zap"
)

//go:embed sql/get-pokemon-data.sql
var getPokemonDataQuery string

type PokemonDataService struct {
	connection *db.Database
}

func NewPokemonDataService(
	connection *db.Database,
) *PokemonDataService {
	return &PokemonDataService{
		connection: connection,
	}
}

func (pc *PokemonDataService) GetAllPokemon(ctx *cmodel.DAContext) ([]*dto.PokemonData, error) {
	var pokemonMap = make(map[int]*dto.PokemonData)
	var results = []*dto.PokemonData{}

	var rows, err = pc.connection.GetDB().Query(getPokemonDataQuery)
	if err != nil {
		ctx.Logger.Error("Error retrieving all pokemon", zap.Error(err))
		return nil, model.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var pokedexId int
		var species string
		var formName string
		var formSprite string
		err := rows.Scan(&pokedexId, &species, &formName, &formSprite)
		if err != nil {
			ctx.Logger.Error("Error scanning pokemon row", zap.Error(err))
			return nil, model.InternalServerError
		}
		form := dto.Form{
			FormId:   formSprite,
			FormName: formName,
			ImageRef: formSprite,
		}
		data, exists := pokemonMap[pokedexId]
		if !exists {
			pokemonData := &dto.PokemonData{
				PokedexNo: dto.PokedexNo{
					Number: pokedexId,
				},
				Species: species,
			}
			pokemonData.Forms = append(pokemonData.Forms, form)
			results = append(results, pokemonData)
			pokemonMap[pokedexId] = pokemonData
		} else {
			data.Forms = append(data.Forms, form)

		}
	}
	ctx.Logger.Info("Successfully retrieved pokemon data")
	return results, nil
}

func (pc *PokemonDataService) GetPokemon(ctx *cmodel.DAContext, pokedexNo int) (*dto.PokemonData, error) {
	var result *dto.PokemonData = nil
	ctx.Logger.Info("Retrieving Pokemon", zap.Int("pokedexNo", pokedexNo))
	var rows, err = pc.connection.GetDB().Query(getPokemonDataQuery+" AND pf.pokedex_id = ?", pokedexNo)
	if err != nil {
		ctx.Logger.Error("Error retrieving Pokemon", zap.Int("pokedexNo", pokedexNo), zap.Error(err))
		return nil, model.InternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		if result == nil {
			result = &dto.PokemonData{}
		}
		var formName string
		var formSprite string
		err := rows.Scan(&result.PokedexNo.Number, &result.Species, &formName, &formSprite)
		if err != nil {
			ctx.Logger.Error("Error retrieving Pokemon", zap.Int("pokedexNo", pokedexNo), zap.Error(err))
			return nil, model.InternalServerError
		}
		form := dto.Form{
			FormId:   formSprite,
			FormName: formName,
			ImageRef: formSprite,
		}

		result.Forms = append(result.Forms, form)
	}
	ctx.Logger.Info("Successfully retrieved Pokemon", zap.Int("pokedexNo", pokedexNo))
	return result, nil
}

func (pds *PokemonDataService) GetFormIdByName(ctx *cmodel.DAContext, formName string) (int, *responses.Error) {
	var formId int
	ctx.Logger.Info("Retrieving Pokemon form id", zap.String("formName", formName))
	row := pds.connection.GetDB().QueryRow(
		"SELECT form_id FROM pokemon_forms pf WHERE pf.sprite_ref = ?",
		formName,
	)
	err := row.Scan(&formId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Logger.Error("No form found", zap.String("formName", formName), zap.Error(err))
			return 0, &derrors.InvalidFormError
		} else {
			ctx.Logger.Error("Error retrieving Pokemon form", zap.String("formName", formName), zap.Error(err))
			return 0, &derrors.InternalServerError
		}
	}
	ctx.Logger.Info("Retrieved Pokemon form", zap.String("formName", formName), zap.Int("formId", formId))
	return formId, nil
}

func (pds *PokemonDataService) GetFormId(ctx *cmodel.DAContext, pokedexNo int, form string) (int, error) {
	var formId int
	ctx.Logger.Info("Retrieving Pokemon form id", zap.Int("pokedexNo", pokedexNo), zap.String("form", form))
	row := pds.connection.GetDB().QueryRow(
		"SELECT form_id FROM pokemon_forms pf WHERE pf.pokedex_id = ? AND pf.form_name = ?",
		pokedexNo,
		form,
	)
	err := row.Scan(&formId)
	if err != nil {
		ctx.Logger.Error("Error retrieving Pokemon form", zap.Int("pokedexNo", pokedexNo), zap.String("form", form), zap.Error(err))
		return 0, model.InternalServerError
	}
	ctx.Logger.Info("Retrieved Pokemon form", zap.Int("pokedexNo", pokedexNo), zap.String("form", form), zap.Int("formId", formId))
	return formId, nil
}
