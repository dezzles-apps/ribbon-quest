package shinies

import (
	"fmt"
)

type CreateShinyRequest struct {
	Form           string `json:"form"`
	Level          int    `json:"level"`
	Game           string `json:"game"`
	Method         string `json:"method"`
	Characteristic string `json:"characteristic"`
	Nature         string `json:"nature"`
	Notes          string `json:"notes"`
}

func (c *CreateShinyRequest) String() string {
	return fmt.Sprintf(
		"CreateShinyRequest{Form: %q, Level: %d, Game: %q, Method: %q, Characteristic: %q, Nature: %q, Notes: %q}",
		c.Form, c.Level, c.Game, c.Method, c.Characteristic, c.Nature, c.Notes,
	)
}
