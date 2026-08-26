package services

import (
	"database/sql"
	"dezzles-apps/rq-server/model/dto"
	"errors"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
)

//go:embed sql/get-ribbon.sql
var getRibbonQuery string

type RibbonService struct {
	connection *cdb.Database
}

func NewRibbonService(
	connection *cdb.Database,
) *RibbonService {
	return &RibbonService{
		connection: connection,
	}
}

func (rs *RibbonService) AddRibbon(pokemon string, ribbon string) (*dto.PokemonRibbon, error) {
	r, err := rs.getRibbon(pokemon, ribbon)
	if err != nil {
		return nil, err
	}
	if r.Achieved {
		return r, nil
	}
	err = rs.addRibbon(pokemon, ribbon)
	if err != nil {
		return nil, err
	}
	return rs.getRibbon(pokemon, ribbon)
}

func (rs *RibbonService) RemoveRibbon(pokemon string, ribbon string) (*dto.PokemonRibbon, error) {
	_, err := rs.getRibbon(pokemon, ribbon)
	if err != nil {
		return nil, err
	}
	err = rs.removeRibbon(pokemon, ribbon)
	if err != nil {
		return nil, err
	}
	return rs.getRibbon(pokemon, ribbon)
}

func (rs *RibbonService) addRibbon(pokemon string, ribbon string) error {
	_, err := rs.connection.GetDB().Exec("INSERT INTO pokemon_ribbons (pokemon, ribbon_key) VALUES (?, ?)", pokemon, ribbon)
	if err != nil {
		return err
	}
	return nil
}

func (rs *RibbonService) removeRibbon(pokemon string, ribbon string) error {
	_, err := rs.connection.GetDB().Exec("DELETE FROM pokemon_ribbons WHERE pokemon = ? AND ribbon_key = ?", pokemon, ribbon)
	if err != nil {
		return err
	}
	return nil
}

func (rs *RibbonService) getRibbon(pokemon string, ribbon string) (*dto.PokemonRibbon, error) {
	ribbonData := &dto.PokemonRibbon{}
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
		return nil, err
	}

	return ribbonData, nil
}
