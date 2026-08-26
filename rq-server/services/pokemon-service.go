package services

import (
	"database/sql"
	"dezzles-apps/rq-server/model/dto"
	"errors"

	_ "embed"

	cdb "github.com/dezzles-apps/go-common/db"
)

//go:embed sql/get-pokemon-ribbons.sql
var getPokemonRibbons string

//go:embed sql/get-pokemon-games.sql
var getPokemonGames string

//go:embed sql/get-pokemon-info.sql
var getPokemonInfo string

//go:embed sql/get-all-pokemon.sql
var getAllPokemon string

type PokemonService struct {
	connection *cdb.Database
}

func NewPokemonService(
	connection *cdb.Database,
) *PokemonService {
	return &PokemonService{
		connection: connection,
	}
}

func (ps *PokemonService) GetPokemon(pokemonName string) (*dto.Pokemon, error) {
	Pokemon, err := ps.getPokemon(pokemonName)
	if err != nil {
		return nil, err
	}
	games, err := ps.getPokemonGames(pokemonName)
	if err != nil {
		return nil, err
	}
	Pokemon.Games = games
	ribbons, err := ps.getPokemonRibbons(pokemonName)
	if err != nil {
		return nil, err
	}
	Pokemon.Ribbons = ribbons
	return Pokemon, nil
}

func (ps *PokemonService) getPokemon(pokemonName string) (*dto.Pokemon, error) {
	Pokemon := &dto.Pokemon{}
	row := ps.connection.GetDB().QueryRow(getPokemonInfo, pokemonName)
	var caughtAt sql.NullString
	var nature sql.NullString
	var characteristic sql.NullString
	err := row.Scan(&Pokemon.Pokemon, &Pokemon.Nickname, &Pokemon.Region, &caughtAt, &nature, &characteristic, &Pokemon.Shiny)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No pokemon found")
		}
		return nil, err
	}
	if caughtAt.Valid {
		Pokemon.CaughtAt = caughtAt.String
	}
	if nature.Valid {
		Pokemon.Nature = nature.String
	}
	if characteristic.Valid {
		Pokemon.Characteristic = characteristic.String
	}
	return Pokemon, nil
}

func (ps *PokemonService) getPokemonGames(pokemonName string) ([]dto.PokemonGame, error) {
	var gamesMap map[string]*dto.PokemonGame = make(map[string]*dto.PokemonGame)
	rows, err := ps.connection.GetDB().Query(getPokemonGames, pokemonName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gameKey string
		var gameName string
		var viewOrder int
		var ribbonKey string
		err := rows.Scan(&gameKey, &gameName, &viewOrder, &ribbonKey)
		if err != nil {
			return nil, err
		}
		if game, exists := gamesMap[gameKey]; exists {
			game.Ribbons = append(game.Ribbons, ribbonKey)
		} else {
			game := &dto.PokemonGame{
				GameKey:   gameKey,
				Name:      gameName,
				ViewOrder: viewOrder,
				Ribbons:   []string{ribbonKey},
			}
			gamesMap[gameKey] = game
		}
	}
	var games []dto.PokemonGame
	for _, game := range gamesMap {
		games = append(games, *game)
	}

	return games, nil
}

func (ps *PokemonService) getPokemonRibbons(pokemonName string) ([]dto.PokemonRibbon, error) {
	var ribbons []dto.PokemonRibbon
	rows, err := ps.connection.GetDB().Query(getPokemonRibbons, pokemonName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ribbonKey string
		var name string
		var achieved sql.NullBool
		var category string
		var order int
		err := rows.Scan(&ribbonKey, &name, &achieved, &category, &order)
		if err != nil {
			return nil, err
		}
		if !achieved.Valid {
			achieved.Bool = false
		}
		ribbon := dto.PokemonRibbon{
			RibbonKey: ribbonKey,
			Name:      name,
			Achieved:  achieved.Bool,
			Category:  category,
		}
		ribbons = append(ribbons, ribbon)
	}
	return ribbons, nil
}

func (ps *PokemonService) GetAllPokemon() ([]dto.AllPokemon, error) {
	var allPokemon []dto.AllPokemon
	rows, err := ps.connection.GetDB().Query(getAllPokemon)
	if err != nil {
		return nil, err
	}
	var pokemonMap map[string]*dto.AllPokemon = make(map[string]*dto.AllPokemon)
	var pokemonList []*dto.AllPokemon = make([]*dto.AllPokemon, 0)
	defer rows.Close()
	for rows.Next() {
		var pokemon string
		var nickname string
		var region string
		var achieved sql.NullBool
		var count int
		err := rows.Scan(&pokemon, &nickname, &region, &achieved, &count)
		if err != nil {
			return nil, err
		}
		if !achieved.Valid {
			achieved.Bool = false
		}
		if p, exists := pokemonMap[pokemon]; exists {
			if achieved.Bool {
				p.Current += count
			}
			p.Target += count
		} else {
			p := &dto.AllPokemon{
				Pokemon:  pokemon,
				Nickname: nickname,
				Region:   region,
				Current:  0,
				Target:   0,
			}
			if achieved.Bool {
				p.Current = count
			}
			p.Target = count
			pokemonMap[pokemon] = p
			pokemonList = append(pokemonList, p)
		}
	}
	for _, p := range pokemonList {
		allPokemon = append(allPokemon, *p)
	}
	return allPokemon, nil
}

func (ps *PokemonService) CatchPokemon(pokemon string) (*dto.Pokemon, error) {
	details, err := ps.getPokemon(pokemon)
	if err != nil {
		return nil, err
	}
	if details.CaughtAt != "" {
		return nil, errors.New("Pokemon already caught")
	}
	_, err = ps.connection.GetDB().Exec("UPDATE pokemon SET caught_at = CURRENT_TIMESTAMP WHERE pokemon = ?", pokemon)
	if err != nil {
		return nil, err
	}
	return ps.GetPokemon(pokemon)
}

func (ps *PokemonService) UpdatePokemon(pokemon string, updateData dto.UpdatePokemon) (*dto.Pokemon, error) {
	details, err := ps.getPokemon(pokemon)
	if err != nil {
		return nil, err
	}
	if updateData.Nickname != "" {
		details.Nickname = updateData.Nickname
	}
	if updateData.Nature != "" {
		details.Nature = updateData.Nature
	}
	if updateData.Characteristic != "" {
		details.Characteristic = updateData.Characteristic
	}
	_, err = ps.connection.GetDB().Exec(
		"UPDATE pokemon SET nickname = ?, nature = ?, characteristic = ?, shiny = ? WHERE pokemon = ?",
		details.Nickname,
		details.Nature,
		details.Characteristic,
		details.Shiny,
		pokemon,
	)
	if err != nil {
		return nil, err
	}
	return ps.GetPokemon(pokemon)
}
