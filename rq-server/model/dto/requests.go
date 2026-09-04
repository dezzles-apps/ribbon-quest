package dto

type UpdatePokemon struct {
	Nickname       string `json:"nickname,omitempty"`
	Nature         string `json:"nature,omitempty"`
	Characteristic string `json:"characteristic,omitempty"`
	Notes          string `json:"notes,omitempty"`
	Shiny          *bool  `json:"shiny,omitempty"`
}
