package model

import (
	"errors"
)

var PokemonNotFound = errors.New("No pokemon found")
var InternalServerError = errors.New("Internal Server Error")
var InvalidUsernameOrPassword = errors.New("Invalid Username or Password")
var PokemonAlreadyCaught = errors.New("Pokemon already caught")
var GameNotFound = errors.New("Game not found")
var PokedexNotFound = errors.New("Pokedex not found")
