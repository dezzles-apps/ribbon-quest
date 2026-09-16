package data

type PokemonData struct {
	PokedexNo int    `json:"pokedexNo"`
	Species   string `json:"species"`
	Forms     []Form `json:"forms"`
}

type Form struct {
	FormName string `json:"form"`
	ImageRef string `json:"imageRef"`
}
