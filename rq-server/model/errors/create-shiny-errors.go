package errors

import (
	"github.com/dezzles-apps/go-common/model/responses"
)

var InvalidFormError = responses.CreateError("form", "Invalid form")
var InvalidLevel = responses.CreateError("level", "Level must be between 1 and 100")
