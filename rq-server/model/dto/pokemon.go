package dto

import (
	"time"
)

type Pokemon struct {
	Pokemon        string          `json:"pokemon"`
	Nickname       string          `json:"nickname"`
	Region         string          `json:"region"`
	CaughtAt       *time.Time      `json:"caughtAt"`
	Nature         string          `json:"nature"`
	Characteristic string          `json:"characteristic"`
	Shiny          bool            `json:"shiny"`
	Ribbons        []PokemonRibbon `json:"ribbons"`
	Games          []PokemonGame   `json:"games"`
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
	Pokemon        string     `json:"pokemon"`
	Nickname       string     `json:"nickname"`
	Region         string     `json:"region"`
	CaughtAt       *time.Time `json:"caughtAt"`
	Nature         string     `json:"nature"`
	Characteristic string     `json:"characteristic"`
	Shiny          bool       `json:"shiny"`
	Current        int        `json:"current"`
	Target         int        `json:"target"`
}
