package model

import (
	"errors"
)

var PokemonNotFound = errors.New("No pokemon found")
var InternalServerError = errors.New("Internal Server Error")
