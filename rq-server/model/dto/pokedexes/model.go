package pokedexes

type Pokedex struct {
	Name    string         `json:"name"`
	Entries []PokedexEntry `json:"entries"`
	Stats   PokedexStats   `json:"stats"`
}

type PokedexEntry struct {
	PokedexNo string `json:"pokedexNo"`
	Pokemon   string `json:"pokemon"`
	SpriteRef string `json:"spriteRef"`
	Caught    bool   `json:"caught"`
}

type PokedexStats struct {
	Caught int `json:"caught"`
	Total  int `json:"total"`
}
