package dto

type Pokemon struct {
	Pokemon  string          `json:"pokemon"`
	Nickname string          `json:"nickname"`
	Region   string          `json:"region"`
	Ribbons  []PokemonRibbon `json:"ribbons"`
	Games    []PokemonGame   `json:"games"`
}

type PokemonRibbon struct {
	RibbonKey  string `json:"ribbonKey"`
	Name       string `json:"name"`
	Achieved   bool   `json:"achieved"`
	AchievedAt string `json:"achievedAt"`
	Category   string `json:"category"`
}

type PokemonGame struct {
	GameKey   string   `json:"gameKey"`
	Name      string   `json:"name"`
	ViewOrder int      `json:"viewOrder"`
	Ribbons   []string `json:"ribbons"`
}

type AllPokemon struct {
	Pokemon  string `json:"pokemon"`
	Nickname string `json:"nickname"`
	Region   string `json:"region"`
	Current  int    `json:"current"`
	Target   int    `json:"target"`
}
