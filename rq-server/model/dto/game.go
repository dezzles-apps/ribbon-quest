package dto

import (
	"time"
)

type Game struct {
	GameKey string         `json:"gameKey"`
	Name    string         `json:"name"`
	Pokemon []*GamePokemon `json:"pokemon"`
}

type GamePokemon struct {
	Pokemon        string          `json:"pokemon"`
	Nickname       string          `json:"nickname"`
	Region         string          `json:"region"`
	CaughtAt       *time.Time      `json:"caughtAt"`
	Nature         string          `json:"nature"`
	Characteristic string          `json:"characteristic"`
	Shiny          bool            `json:"shiny"`
	Ribbons        []PokemonRibbon `json:"ribbons"`
}

type GameWithStats struct {
	GameKey  string `json:"gameKey"`
	Name     string `json:"name"`
	Achieved int    `json:"achieved"`
	Total    int    `json:"total"`
}
