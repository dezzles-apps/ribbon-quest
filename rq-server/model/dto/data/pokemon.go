package data

import (
	"dezzles-apps/rq-server/model"
)

type PokemonData struct {
	PokedexNo int    `json:"pokedexNo"`
	Species   string `json:"species"`
	Forms     []Form `json:"forms"`
}

type Form struct {
	FormId   string `json:"formId"`
	FormName string `json:"form"`
	ImageRef string `json:"imageRef"`
}

type CommonData struct {
	Games           []model.KeyValuePair `json:"games"`
	Genders         []model.KeyValuePair `json:"genders"`
	Characteristics []model.KeyValuePair `json:"characteristics"`
	Natures         []model.KeyValuePair `json:"natures"`
	Methods         []model.KeyValuePair `json:"methods"`
}
