package shinies

import (
	"dezzles-apps/rq-server/model/dto/shinies"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/static"
	dataServices "dezzles-apps/rq-server/services/data"
	"testing"
)

type CreateShinyTest struct {
	Description string
	Request     shinies.CreateShinyRequest
	Expected    errors.Error
}

func setupTestService() *ShinyService {
	return NewShinyService(
		dataServices.NewCommonDataService(
			static.NewCommonDataService(),
		),
		nil,
		nil,
	)
}

func TestValidateCreateShiny(t *testing.T) {
	testService := setupTestService()
	tests := []CreateShinyTest{
		{
			Description: "Returns error when level less than 1",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          -1,
				Game:           "X",
				Method:         "RANDOM",
				Characteristic: "MISCHIEVOUS",
				Nature:         "MODEST",
				Gender:         "F",
			},
			Expected: errors.InvalidLevel,
		},
		{
			Description: "Returns error when level greater than 100",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          101,
				Game:           "X",
				Method:         "RANDOM",
				Characteristic: "MISCHIEVOUS",
				Nature:         "MODEST",
				Gender:         "U",
			},
			Expected: errors.InvalidLevel,
		},
		{
			Description: "Returns error when invalid game",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          100,
				Game:           "Z",
				Method:         "RANDOM",
				Characteristic: "MISCHIEVOUS",
				Nature:         "MODEST",
				Gender:         "M",
			},
			Expected: errors.InvalidGame,
		},
		{
			Description: "Returns error when invalid method",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          100,
				Game:           "OMEGA_RUBY",
				Method:         "NA",
				Characteristic: "MISCHIEVOUS",
				Nature:         "MODEST",
				Gender:         "M",
			},
			Expected: errors.InvalidMethod,
		},
		{
			Description: "Returns error when invalid characteristic",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          100,
				Game:           "OMEGA_RUBY",
				Method:         "RANDOM",
				Characteristic: "IS_JEFF",
				Nature:         "MODEST",
				Gender:         "M",
			},
			Expected: errors.InvalidCharacteristic,
		},
		{
			Description: "Returns error when invalid nature",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          100,
				Game:           "OMEGA_RUBY",
				Method:         "RANDOM",
				Characteristic: "MISCHIEVOUS",
				Nature:         "SNEAKY",
				Gender:         "M",
			},
			Expected: errors.InvalidNature,
		},
		{
			Description: "Returns error when invalid nature",
			Request: shinies.CreateShinyRequest{
				Form:           "psyduck",
				Level:          100,
				Game:           "OMEGA_RUBY",
				Method:         "RANDOM",
				Characteristic: "MISCHIEVOUS",
				Nature:         "ADAMANT",
				Gender:         "Z",
			},
			Expected: errors.InvalidGender,
		},
	}
	for _, test := range tests {
		actual := testService.validateRequest(test.Request).Errors()
		if len(actual) != 1 {
			t.Errorf("%s: expected 1 error, got %d", test.Description, len(actual))
			continue
		}
		if actual[0] != test.Expected {
			t.Errorf("%s: Expected %s, got but got %s", test.Description, test.Expected.Error(), actual[0].Error())
		}
	}
}
