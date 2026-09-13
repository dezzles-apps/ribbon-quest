package dto

type Game struct {
	GameKey string         `json:"gameKey"`
	Name    string         `json:"name"`
	Pokemon []*GamePokemon `json:"pokemon"`
}

type GamePokemon struct {
	Pokemon string          `json:"pokemon"`
	Species PokemonSpecies  `json:"species"`
	Details CatchDetails    `json:"details"`
	Ribbons []PokemonRibbon `json:"ribbons"`
}

type GameWithStats struct {
	GameKey  string `json:"gameKey"`
	Name     string `json:"name"`
	Achieved int    `json:"achieved"`
	Total    int    `json:"total"`
}
