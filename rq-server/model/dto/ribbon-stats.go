package dto

type RibbonStats struct {
	Current int          `json:"current"`
	Total   int          `json:"total"`
	Pokemon []AllPokemon `json:"pokemon"`
}
