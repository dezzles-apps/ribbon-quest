package data

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto/data"
	"dezzles-apps/rq-server/model/errors"
	repos "dezzles-apps/rq-server/repositories/interfaces"

	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"
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

func (cds *CommonDataService) GetAllData(ctx *cmodel.DAContext) (*data.CommonData, errors.Error) {
	result := &data.CommonData{}
	ctx.Logger.Info("Getting common data: Characteristics")
	kvp, err := cds.GetCharacteristics()
	ctx.Logger.Info("Got common data: Characteristics")
	if err != nil {
		ctx.Logger.Error("WAT", zap.Error(err))
		return nil, err
	}
	result.Characteristics = kvp
	ctx.Logger.Info("Getting common data: Games")
	kvp, err = cds.GetGames()
	if err != nil {
		return nil, err
	}
	result.Games = kvp
	ctx.Logger.Info("Getting common data: Methods")
	kvp, err = cds.GetMethods()
	if err != nil {
		return nil, err
	}
	result.Methods = kvp

	ctx.Logger.Info("Getting common data: Natures")
	kvp, err = cds.GetNatures()
	if err != nil {
		return nil, err
	}
	result.Natures = kvp

	kvp, err = cds.GetGenders()
	if err != nil {
		return nil, err
	}
	result.Genders = kvp
	return result, nil
}

func (cds *CommonDataService) GetCharacteristics() ([]model.KeyValuePair, errors.Error) {
	return cds.repository.GetCharacteristics()
}

func (cds *CommonDataService) GetGames() ([]model.KeyValuePair, errors.Error) {
	return cds.repository.GetGames()
}

func (cds *CommonDataService) GetNatures() ([]model.KeyValuePair, errors.Error) {
	return cds.repository.GetNatures()
}

func (cds *CommonDataService) GetMethods() ([]model.KeyValuePair, errors.Error) {
	return cds.repository.GetMethods()
}

func (cds *CommonDataService) GetGenders() ([]model.KeyValuePair, errors.Error) {
	return cds.repository.GetGenders()
}

func (cds *CommonDataService) ValidGame(game string) errors.Error {
	games, err := cds.GetGames()
	if err != nil {
		return nil
	}
	for _, g := range games {
		if g.Key == game {
			return nil
		}
	}
	return errors.InvalidGame
}

func (cds *CommonDataService) ValidMethod(method string) errors.Error {
	methods, err := cds.GetMethods()
	if err != nil {
		return nil
	}
	for _, m := range methods {
		if m.Key == method {
			return nil
		}
	}
	return errors.InvalidMethod
}

func (cds *CommonDataService) ValidCharacteristic(characteristic string) errors.Error {
	characteristics, err := cds.GetCharacteristics()
	if err != nil {
		return nil
	}
	for _, c := range characteristics {
		if c.Key == characteristic {
			return nil
		}
	}
	return errors.InvalidCharacteristic
}

func (cds *CommonDataService) ValidNature(nature string) errors.Error {
	natures, err := cds.GetNatures()
	if err != nil {
		return nil
	}
	for _, n := range natures {
		if n.Key == nature {
			return nil
		}
	}
	return errors.InvalidNature
}

func (cds *CommonDataService) ValidGender(gender string) errors.Error {
	genders, err := cds.GetGenders()
	if err != nil {
		return nil
	}
	for _, n := range genders {
		if n.Key == gender {
			return nil
		}
	}
	return errors.InvalidGender
}
