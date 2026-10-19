package errors

import (
	"github.com/dezzles-apps/go-common/model/responses"
)

var InvalidGame = responses.CreateError("game", "Invalid game")
var InvalidMethod = responses.CreateError("method", "Invalid method")
var InvalidCharacteristic = responses.CreateError("characteristic", "Invalid characteristic")
var InvalidNature = responses.CreateError("nature", "Invalid nature")

var InternalServerError = responses.CreateError("Internal Server Error", "Something went wrong")
