package data

import (
	dto "dezzles-apps/rq-server/model/dto/data"
	"errors"
	"log"

	_ "embed"

	"github.com/dezzles-apps/go-common/db"
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

func (pc *PokemonDataService) GetAllPokemon() ([]*dto.PokemonData, error) {
	var pokemonMap = make(map[int]*dto.PokemonData)
	var results = []*dto.PokemonData{}

	var rows, err = pc.connection.GetDB().Query(getPokemonDataQuery)
	if err != nil {
		log.Printf("Error retrieving all pokemon: %s", err.Error())
		return nil, errors.New("Internal server error")
	}
	defer rows.Close()
	for rows.Next() {
		var pokedexId int
		var species string
		var formName string
		var formSprite string
		err := rows.Scan(&pokedexId, &species, &formName, &formSprite)
		if err != nil {
			log.Printf("Error scannig row: %s", err.Error())
			return nil, errors.New("Internal server error")
		}
		form := dto.Form{
			FormName: formName,
			ImageRef: formSprite,
		}
		data, exists := pokemonMap[pokedexId]
		if !exists {
			pokemonData := &dto.PokemonData{
				PokedexNo: pokedexId,
				Species:   species,
			}
			pokemonData.Forms = append(pokemonData.Forms, form)
			results = append(results, pokemonData)
			pokemonMap[pokedexId] = pokemonData
		} else {
			data.Forms = append(data.Forms, form)

		}
	}
	return results, nil
}

func (pc *PokemonDataService) GetPokemon(pokedexNo int) (*dto.PokemonData, error) {
	var result *dto.PokemonData = nil
	var rows, err = pc.connection.GetDB().Query(getPokemonDataQuery+" AND pf.pokedex_id = ?", pokedexNo)
	if err != nil {
		log.Printf("Error retrieving all pokemon: %s", err.Error())
		return nil, errors.New("Internal server error")
	}
	defer rows.Close()
	for rows.Next() {
		if result == nil {
			result = &dto.PokemonData{}
		}
		var formName string
		var formSprite string
		err := rows.Scan(&result.PokedexNo, &result.Species, &formName, &formSprite)
		if err != nil {
			log.Printf("Error scannig row: %s", err.Error())
			return nil, errors.New("Internal server error")
		}
		form := dto.Form{
			FormName: formName,
			ImageRef: formSprite,
		}

		result.Forms = append(result.Forms, form)
	}
	return result, nil
}

func (pds *PokemonDataService) GetFormId(pokedexId int, form string) (int, error) {
	var number int
	row := pds.connection.GetDB().QueryRow(
		"SELECT form_id FROM pokemon_forms pf WHERE pf.pokedex_id = ? AND pf.form_name = ?",
		pokedexId,
		form,
	)
	err := row.Scan(&number)
	if err != nil {
		return 0, err
	}
	return number, nil
}
