package pokedexes

type Pokedex struct {
	Name    string         `json:"name"`
	Entries []PokedexEntry `json:"entries"`
	Stats   PokedexStats   `json:"stats"`
	Image   string         `json:"image"`
	Notes   string         `json:"notes"`
}

type PokedexEntry struct {
	PokedexNo string `json:"pokedexNo"`
	Pokemon   string `json:"pokemon"`
	Caught    bool   `json:"caught"`
	SpriteRef string `json:"spriteRef"`
}

type PokedexStats struct {
	Caught int `json:"caught"`
	Total  int `json:"total"`
}
