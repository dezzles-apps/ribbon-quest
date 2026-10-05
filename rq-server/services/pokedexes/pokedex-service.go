package pokedexes

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"

	errs "dezzles-apps/rq-server/model"
	model "dezzles-apps/rq-server/model/dto/pokedexes"

	cdb "github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"
)

//go:embed sql/get-all-pokedexes.sql
var getAllPokedexesQuery string

//go:embed sql/get-pokedex-entries.sql
var getPokedexEntries string

//go:embed sql/get-pokedex-entry.sql
var getPokedexEntry string

//go:embed sql/uncatch-pokemon.sql
var uncatchPokemon string

//go:embed sql/catch-pokemon.sql
var catchPokemon string

type PokedexService struct {
	connection *cdb.Database
}

func NewPokedexService(
	connection *cdb.Database,
) *PokedexService {
	return &PokedexService{
		connection: connection,
	}
}

func (ps *PokedexService) GetPokedexes(ctx *cmodel.DAContext) ([]model.Pokedex, error) {
	ctx.Logger.Info("Retrieving all pokedexes")
	rows, err := ps.connection.GetDB().Query(getAllPokedexesQuery)
	if err != nil {
		ctx.Logger.Error("Retrieving pokedexes failed", zap.Error(err))
		return nil, errs.InternalServerError
	}
	defer rows.Close()
	var result []model.Pokedex = []model.Pokedex{}
	for rows.Next() {
		dex := model.Pokedex{}
		var total sql.NullInt32
		var caught sql.NullInt32
		err = rows.Scan(&dex.Name, &caught, &total, &dex.Image, &dex.Notes)
		if err != nil {
			ctx.Logger.Error("Scanning pokedexes failed", zap.Error(err))
			return nil, errs.InternalServerError
		}
		if total.Valid {
			dex.Stats.Total = int(total.Int32)
		}
		if caught.Valid {
			dex.Stats.Caught = int(caught.Int32)
		}

		result = append(result, dex)
	}
	return result, nil
}

func (ps *PokedexService) GetPokedex(ctx *cmodel.DAContext, name string) (*model.Pokedex, error) {
	result, prefix, padding, err := ps.getPokedex(ctx, name)
	if err != nil {
		return nil, err
	}
	if err := ps.getPokedexEntries(ctx, result, prefix, padding); err != nil {
		return nil, err
	}
	return result, nil
}

func (ps *PokedexService) getPokedex(ctx *cmodel.DAContext, name string) (*model.Pokedex, string, int, error) {
	var pokedex model.Pokedex
	var prefix string
	var padding int
	row := ps.connection.GetDB().QueryRow("SELECT name, prefix, padding, image, notes FROM pokedexes WHERE name = ?", name)
	err := row.Scan(&pokedex.Name, &prefix, &padding, &pokedex.Image, &pokedex.Notes)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", 0, errs.PokedexNotFound
		}
		ctx.Logger.Info("Error retrieving pokedex", zap.String("pokedex", name), zap.Error(err))
		return nil, "", 0, errs.InternalServerError
	}

	return &pokedex, prefix, padding, nil
}

func (ps *PokedexService) getPokedexEntry(ctx *cmodel.DAContext, pokedex string, inputNo string, pokedexNo int) (*model.PokedexEntry, error) {
	ctx.Logger.Info("Getting pokedexEntry", zap.String("pokedex", pokedex), zap.Int("pokedexNo", pokedexNo))
	row := ps.connection.GetDB().QueryRow(getPokedexEntry, pokedex, pokedexNo)
	if row.Err() != nil {
		ctx.Logger.Info("Error getting pokedex entry", zap.String("pokedex", pokedex), zap.Int("pokedexNo", pokedexNo), zap.Error(row.Err()))
		return nil, errs.InternalServerError
	}
	var entry model.PokedexEntry
	entry.PokedexNo = inputNo
	var tempNo int
	err := row.Scan(&tempNo, &entry.Pokemon, &entry.SpriteRef, &entry.Caught)
	if err != nil {
		ctx.Logger.Info("Error scanning pokedex entry", zap.String("pokedex", pokedex), zap.Int("pokedexNo", pokedexNo), zap.Error(row.Err()))
		return nil, errs.InternalServerError
	}
	return &entry, nil
}

func (ps *PokedexService) getPokedexEntries(ctx *cmodel.DAContext, pokedex *model.Pokedex, prefix string, padding int) error {
	rows, err := ps.connection.GetDB().Query(getPokedexEntries, pokedex.Name)
	if err != nil {
		ctx.Logger.Info("Error getting pokedex entries", zap.String("pokedex", pokedex.Name), zap.Error(err))
		return errs.InternalServerError
	}
	defer rows.Close()
	var format string = "%s-%0" + strconv.Itoa(padding) + "d"
	for rows.Next() {
		var entry model.PokedexEntry
		var pokedexNo int
		err = rows.Scan(&pokedexNo, &entry.Pokemon, &entry.SpriteRef, &entry.Caught)
		if err != nil {
			ctx.Logger.Info("Error scanning pokedex entries", zap.String("pokedex", pokedex.Name), zap.Error(err))
			return errs.InternalServerError
		}
		entry.PokedexNo = fmt.Sprintf(format, prefix, pokedexNo)
		if entry.Caught {
			pokedex.Stats.Caught += 1
		}
		pokedex.Stats.Total += 1
		pokedex.Entries = append(pokedex.Entries, entry)
	}
	return nil
}

func (ps *PokedexService) UncatchPokemon(ctx *cmodel.DAContext, pokedex string, pokedexNo string) (*model.PokedexEntry, error) {
	_, number, ok := strings.Cut(pokedexNo, "-")
	if !ok {
		return nil, errors.New("Invalid pokedexNo")
	}
	actualNo, err := strconv.Atoi(number)
	if err != nil {
		ctx.Logger.Error("Failed to get PokedexNumber", zap.String("pokedexNo", pokedexNo), zap.Error(err))
		return nil, errs.InvalidPokedexNo
	}
	ctx.Logger.Info("Uncatching Pokemon", zap.String("pokedex", pokedex), zap.Int("pokedexNo", actualNo))
	_, err = ps.connection.GetDB().Exec(uncatchPokemon, pokedex, actualNo)
	if err != nil {
		ctx.Logger.Error("Failed to update Entry", zap.String("pokedex", pokedex), zap.String("pokedexNo", pokedexNo), zap.Error(err))
		return nil, errs.InternalServerError
	}
	ctx.Logger.Info("Uncaught Pokemon", zap.String("pokedex", pokedex), zap.Int("pokedexNo", actualNo))
	return ps.getPokedexEntry(ctx, pokedex, pokedexNo, actualNo)
}

func (ps *PokedexService) CatchPokemon(ctx *cmodel.DAContext, pokedex string, pokedexNo string) (*model.PokedexEntry, error) {
	_, number, ok := strings.Cut(pokedexNo, "-")
	if !ok {
		return nil, errors.New("Invalid pokedexNo")
	}
	actualNo, err := strconv.Atoi(number)
	if err != nil {
		ctx.Logger.Error("Failed to get PokedexNumber", zap.String("pokedexNo", pokedexNo), zap.Error(err))
	}
	ctx.Logger.Info("Catching Pokemon", zap.String("pokedex", pokedex), zap.Int("pokedexNo", actualNo))
	_, err = ps.connection.GetDB().Exec(catchPokemon, pokedex, actualNo)
	if err != nil {
		ctx.Logger.Error("Failed to update Entry", zap.String("pokedex", pokedex), zap.String("pokedexNo", pokedexNo), zap.Error(err))
		return nil, errs.InternalServerError
	}
	ctx.Logger.Info("Caught Pokemon", zap.String("pokedex", pokedex), zap.Int("pokedexNo", actualNo))
	return ps.getPokedexEntry(ctx, pokedex, pokedexNo, actualNo)
}
