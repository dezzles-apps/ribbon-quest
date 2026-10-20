package data

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/errors"
	repos "dezzles-apps/rq-server/repositories"

	"github.com/dezzles-apps/go-common/model/responses"
)

type CommonDataService struct {
	repository repos.CommonDataRepository
}

func NewCommonDataService(repository repos.CommonDataRepository) *CommonDataService {
	cds := &CommonDataService{
		repository: repository,
	}
	return cds
}

func (cds *CommonDataService) GetCharacteristics() ([]model.KeyValuePair, error) {
	return cds.repository.GetCharacteristics()
}

func (cds *CommonDataService) GetGames() ([]model.KeyValuePair, error) {
	return cds.repository.GetGames()
}

func (cds *CommonDataService) GetNatures() ([]model.KeyValuePair, error) {
	return cds.repository.GetNatures()
}

func (cds *CommonDataService) GetMethods() ([]model.KeyValuePair, error) {
	return cds.repository.GetMethods()
}

func (cds *CommonDataService) GetGenders() ([]model.KeyValuePair, error) {
	return cds.repository.GetGenders()
}

func (cds *CommonDataService) ValidGame(game string) *responses.Error {
	games, err := cds.GetGames()
	if err != nil {
		return nil
	}
	for _, g := range games {
		if g.Key == game {
			return nil
		}
	}
	return &errors.InvalidGame
}

func (cds *CommonDataService) ValidMethod(method string) *responses.Error {
	methods, err := cds.GetMethods()
	if err != nil {
		return nil
	}
	for _, m := range methods {
		if m.Key == method {
			return nil
		}
	}
	return &errors.InvalidMethod
}

func (cds *CommonDataService) ValidCharacteristic(characteristic string) *responses.Error {
	characteristics, err := cds.GetCharacteristics()
	if err != nil {
		return nil
	}
	for _, c := range characteristics {
		if c.Key == characteristic {
			return nil
		}
	}
	return &errors.InvalidCharacteristic
}

func (cds *CommonDataService) ValidNature(nature string) *responses.Error {
	natures, err := cds.GetNatures()
	if err != nil {
		return nil
	}
	for _, n := range natures {
		if n.Key == nature {
			return nil
		}
	}
	return &errors.InvalidNature
}

func (cds *CommonDataService) ValidGender(gender string) *responses.Error {
	genders, err := cds.GetGenders()
	if err != nil {
		return nil
	}
	for _, n := range genders {
		if n.Key == gender {
			return nil
		}
	}
	return &errors.InvalidGender
}
