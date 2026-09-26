package ribbons

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	"errors"
	"log"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
)

//go:embed sql/get-ribbon.sql
var getRibbonQuery string

type RibbonService struct {
	connection     *cdb.Database
	pokemonService PokemonService
}

func NewRibbonService(
	connection *cdb.Database,
	PokemonService *PokemonService,
) *RibbonService {
	return &RibbonService{
		connection:     connection,
		pokemonService: *PokemonService,
	}
}

func (rs *RibbonService) AddRibbon(ctx *model.RQContext, pokemon string, ribbon string) (*dto.PokemonRibbon, error) {
	log.Printf("AddRibbon: Adding ribbon %s to %s", ribbon, pokemon)
	pokemonId, err := rs.pokemonService.GetPokemonId(ctx, pokemon)
	if err != nil {
		log.Printf("AddRibbon: Error adding ribbon %s to %s. %s", ribbon, pokemon, err.Error())
		return nil, model.InternalServerError
	}
	r, err := rs.getRibbon(ctx, pokemonId, ribbon)
	if err != nil {
		return nil, err
	}
	if r.Achieved {
		return r, nil
	}
	err = rs.addRibbon(ctx, pokemonId, ribbon)
	if err != nil {
		return nil, err
	}
	return rs.getRibbon(ctx, pokemonId, ribbon)
}

func (rs *RibbonService) RemoveRibbon(ctx *model.RQContext, pokemon string, ribbon string) (*dto.PokemonRibbon, error) {
	pokemonId, err := rs.pokemonService.GetPokemonId(ctx, pokemon)
	if err != nil {
		return nil, err
	}

	_, err = rs.getRibbon(ctx, pokemonId, ribbon)
	if err != nil {
		return nil, err
	}
	err = rs.removeRibbon(ctx, pokemonId, ribbon)
	if err != nil {
		return nil, err
	}
	return rs.getRibbon(ctx, pokemonId, ribbon)
}

func (rs *RibbonService) addRibbon(ctx *model.RQContext, pokemon int, ribbon string) error {
	_, err := rs.connection.GetDB().Exec("INSERT INTO ribbons_earned (ribbon_pokemon_id, ribbon_key) VALUES (?, ?)", pokemon, ribbon)
	if err != nil {
		return err
	}
	return nil
}

func (rs *RibbonService) removeRibbon(ctx *model.RQContext, pokemon int, ribbon string) error {
	_, err := rs.connection.GetDB().Exec("DELETE FROM ribbons_earned WHERE ribbon_pokemon_id = ? AND ribbon_key = ?", pokemon, ribbon)
	if err != nil {
		return err
	}
	return nil
}

func (rs *RibbonService) getRibbon(ctx *model.RQContext, pokemon int, ribbon string) (*dto.PokemonRibbon, error) {
	ribbonData := &dto.PokemonRibbon{}
	log.Printf("getRibbon: Retrieving ribbon %d:%s", pokemon, ribbon)
	err := rs.connection.GetDB().QueryRow(getRibbonQuery, pokemon, ribbon).Scan(
		&ribbonData.RibbonKey,
		&ribbonData.Name,
		&ribbonData.Achieved,
		&ribbonData.Category,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid ribbon combination")
		}
		log.Printf("getRibbon: Error retrieving ribbon %d:%s. %s", pokemon, ribbon, err.Error())
		return nil, model.InternalServerError
	}

	return ribbonData, nil
}
