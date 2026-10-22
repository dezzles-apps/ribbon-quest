package model

import (
	"dezzles-apps/rq-server/model/errors"
)

var PokemonNotFound = errors.CreateError("", "No pokemon found")
var InternalServerError = errors.CreateError("", "Internal Server Error")
var InvalidUsernameOrPassword = errors.CreateError("", "Invalid Username or Password")
var PokemonAlreadyCaught = errors.CreateError("", "Pokemon already caught")
var GameNotFound = errors.CreateError("", "Game not found")
var PokedexNotFound = errors.CreateError("", "Pokedex not found")
var InvalidPokedexNo = errors.CreateError("", "Invalid PokedexNo")
