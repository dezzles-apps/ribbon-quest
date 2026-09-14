package dto

import (
	"time"
)

type PokemonSpecies struct {
	PokedexId int    `json:"pokedexId"`
	Name      string `json:"name"`
	Form      string `json:"form"`
	ImageRef  string `json:"imageRef"`
}

type CatchDetails struct {
	Nickname       string     `json:"nickname"`
	CaughtAt       *time.Time `json:"caughtAt"`
	Nature         string     `json:"nature"`
	Characteristic string     `json:"characteristic"`
	Notes          string     `json:"notes"`
	Shiny          bool       `json:"shiny"`
}

type Pokemon struct {
	Pokemon string          `json:"pokemon"`
	Species PokemonSpecies  `json:"species"`
	Details CatchDetails    `json:"details"`
	Ribbons []PokemonRibbon `json:"ribbons"`
	Games   []PokemonGame   `json:"games"`
}

type PokemonRibbon struct {
	RibbonKey  string     `json:"ribbonKey"`
	Name       string     `json:"name"`
	Achieved   bool       `json:"achieved"`
	AchievedAt *time.Time `json:"achievedAt"`
	Category   string     `json:"category"`
}

type PokemonGame struct {
	GameKey   string   `json:"gameKey"`
	Name      string   `json:"name"`
	ViewOrder int      `json:"viewOrder"`
	Ribbons   []string `json:"ribbons"`
}

type AllPokemon struct {
	Pokemon string         `json:"pokemon"`
	Species PokemonSpecies `json:"species"`
	Details CatchDetails   `json:"details"`
	Current int            `json:"current"`
	Total   int            `json:"total"`
}
