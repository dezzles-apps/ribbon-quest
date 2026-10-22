package data

import (
	"dezzles-apps/rq-server/model"
	"fmt"
	"strconv"
)

type PokedexNo struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func CreatePokedexNumber(number int, lead string, digits int) PokedexNo {
	f := "%s%0" + strconv.Itoa(digits) + "d"
	return PokedexNo{
		Number: number,
		Text:   fmt.Sprintf(f, lead, number),
	}
}

type PokemonData struct {
	PokedexNo PokedexNo `json:"pokedexNo"`
	Species   string    `json:"species"`
	Forms     []Form    `json:"forms"`
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
