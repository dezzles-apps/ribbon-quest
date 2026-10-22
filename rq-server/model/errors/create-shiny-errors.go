package errors

var InvalidFormError = CreateError("form", "Invalid form")
var InvalidLevel = CreateError("level", "Level must be between 1 and 100")
