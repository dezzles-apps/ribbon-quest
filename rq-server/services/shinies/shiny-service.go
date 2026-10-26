package shinies

import (
	"dezzles-apps/rq-server/model/dto"
	request "dezzles-apps/rq-server/model/dto/shinies"
	"dezzles-apps/rq-server/model/errors"
	"dezzles-apps/rq-server/repositories/interfaces"
	data "dezzles-apps/rq-server/services/data"
	"time"

	cmodel "github.com/dezzles-apps/go-common/model"
)

type ShinyService struct {
	commonDataService  *data.CommonDataService
	pokemonDataService *data.PokemonDataService
	shinyRepository    interfaces.ShinyRepository
}

func NewShinyService(
	commonDataService *data.CommonDataService,
	pokemonDataService *data.PokemonDataService,
	shinyRepository interfaces.ShinyRepository,
) *ShinyService {
	return &ShinyService{
		commonDataService:  commonDataService,
		pokemonDataService: pokemonDataService,
		shinyRepository:    shinyRepository,
	}
}

func (svc *ShinyService) GetShinies(ctx *cmodel.DAContext) ([]dto.ShinyPokemon, errors.Error) {
	return svc.shinyRepository.GetAllShinies(ctx)
}

func (svc *ShinyService) GetShinyStats(ctx *cmodel.DAContext, startDate, endDate string) (*request.ShinyStatsResponse, errors.Error) {
	// Parse and validate dates
	start, end, err := svc.validateAndParseDates(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Fetch shinies for date range
	shinies, fetchErr := svc.shinyRepository.GetShiniesByDateRange(ctx, start.Format("2006-01-02"), end.Format("2006-01-02"))
	if fetchErr != nil {
		return nil, fetchErr
	}

	// Calculate stats
	calculator := NewStatsCalculator(shinies, start, end)
	stats := calculator.CalculateStats(start, end)

	return stats, nil
}

// validateAndParseDates parses date strings and applies defaults
func (svc *ShinyService) validateAndParseDates(startDate, endDate string) (time.Time, time.Time, errors.Error) {
	now := time.Now().UTC()
	var start, end time.Time
	var err error

	// Parse or default start date (12 months ago)
	if startDate == "" {
		start = now.AddDate(-1, 0, 0)
	} else {
		start, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			e := errors.CreateError("startDate", "Invalid date format. Use YYYY-MM-DD")
			return time.Time{}, time.Time{}, e
		}
	}

	// Parse or default end date (today)
	if endDate == "" {
		end = now
	} else {
		end, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			e := errors.CreateError("endDate", "Invalid date format. Use YYYY-MM-DD")
			return time.Time{}, time.Time{}, e
		}
	}

	// Validate order
	if start.After(end) {
		e := errors.CreateError("dateRange", "Start date must be before end date")
		return time.Time{}, time.Time{}, e
	}

	return start, end, nil
}

func (svc *ShinyService) ValidateRequest(ctx *cmodel.DAContext, request request.CreateShinyRequest) errors.Error {
	err := svc.validateRequest(request)
	_, e := svc.pokemonDataService.GetFormIdByName(ctx, request.Form)
	if e != nil {
		if e.Fail() {
			return e
		} else {
			res := &errors.ValidationErrors{}
			res.AddError(errors.InvalidFormError)
			for _, e := range err.Errors() {
				res.AddError(e)
			}
			return res
		}
	}
	if err != nil && err.Fail() {
		return err
	}

	return nil
}

func (svc *ShinyService) validateRequest(request request.CreateShinyRequest) errors.Error {
	result := errors.ValidationErrors{}
	if request.Level < 1 || request.Level > 100 {
		result.AddError(errors.InvalidLevel)
	}

	gameErr := svc.commonDataService.ValidGame(request.Game)
	if gameErr != nil {
		if gameErr.Fail() {
			return gameErr
		}
		result.AddError(gameErr)
	}

	methodErr := svc.commonDataService.ValidMethod(request.Method)
	if methodErr != nil {
		if methodErr.Fail() {
			return methodErr
		}
		result.AddError(methodErr)
	}

	characteristicErr := svc.commonDataService.ValidCharacteristic(request.Characteristic)
	if characteristicErr != nil {
		if characteristicErr.Fail() {
			return characteristicErr
		}
		result.AddError(characteristicErr)
	}

	natureErr := svc.commonDataService.ValidNature(request.Nature)
	if natureErr != nil {
		if natureErr.Fail() {
			return natureErr
		}
		result.AddError(natureErr)
	}

	genderErr := svc.commonDataService.ValidGender(request.Gender)
	if genderErr != nil {
		if genderErr.Fail() {
			return genderErr
		}
		result.AddError(genderErr)
	}
	if result.Fail() {
		return &result
	}
	return nil
}

func (svc *ShinyService) CreateShiny(ctx *cmodel.DAContext, request request.CreateShinyRequest) errors.Error {
	errs := svc.ValidateRequest(ctx, request)
	if errs != nil {
		return errs
	}
	return svc.shinyRepository.SaveShiny(ctx, request)
}
