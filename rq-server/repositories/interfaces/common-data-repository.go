package interfaces

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/errors"
)

type CommonDataRepository interface {
	GetCharacteristics() ([]model.KeyValuePair, errors.Error)
	GetNatures() ([]model.KeyValuePair, errors.Error)
	GetGames() ([]model.KeyValuePair, errors.Error)
	GetMethods() ([]model.KeyValuePair, errors.Error)
	GetGenders() ([]model.KeyValuePair, errors.Error)
	GetCharacteristic(key string) (*model.KeyValuePair, errors.Error)
	GetNature(key string) (*model.KeyValuePair, errors.Error)
	GetGame(key string) (*model.KeyValuePair, errors.Error)
	GetMethod(key string) (*model.KeyValuePair, errors.Error)
	GetGender(key string) (*model.KeyValuePair, errors.Error)
}
