package shinies

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	"testing"
	"time"
)

func TestCalculateMonthlyByGame(t *testing.T) {
	startDate := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, time.March, 31, 0, 0, 0, 0, time.UTC)

	// Create test data
	shinies := []dto.ShinyPokemon{
		{
			ShinyId: 1,
			Species: dto.PokemonSpecies{Name: "Pikachu", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"},
				Level: 10,
				Shiny: true,
			},
		},
		{
			ShinyId: 2,
			Species: dto.PokemonSpecies{Name: "Pikachu", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"},
				Level: 15,
				Shiny: true,
			},
		},
		{
			ShinyId: 3,
			Species: dto.PokemonSpecies{Name: "Charmander", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.February, 10, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"},
				Level: 12,
				Shiny: true,
			},
		},
		{
			ShinyId: 4,
			Species: dto.PokemonSpecies{Name: "Squirtle", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.February, 15, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"},
				Level: 8,
				Shiny: true,
			},
		},
	}

	calculator := NewStatsCalculator(shinies, startDate, endDate)
	monthly := calculator.calculateMonthlyByGame()

	// Verify results
	if len(monthly) != 2 {
		t.Errorf("Expected 2 games, got %d", len(monthly))
	}

	// Check Scarlet
	if scarletData, ok := monthly["Scarlet"]; ok {
		if len(scarletData) != 3 {
			t.Errorf("Expected 3 months of data for Scarlet, got %d", len(scarletData))
		}
		// January should have 2 catches
		if scarletData[0].Count != 2 {
			t.Errorf("Expected 2 catches in January for Scarlet, got %d", scarletData[0].Count)
		}
		if scarletData[0].Cumulative != 2 {
			t.Errorf("Expected cumulative 2 in January for Scarlet, got %d", scarletData[0].Cumulative)
		}
		// February should have 1 catch
		if scarletData[1].Count != 1 {
			t.Errorf("Expected 1 catch in February for Scarlet, got %d", scarletData[1].Count)
		}
		if scarletData[1].Cumulative != 3 {
			t.Errorf("Expected cumulative 3 in February for Scarlet, got %d", scarletData[1].Cumulative)
		}
	} else {
		t.Error("Scarlet data not found")
	}

	// Check Violet
	if violetData, ok := monthly["Violet"]; ok {
		if len(violetData) != 3 {
			t.Errorf("Expected 3 months of data for Violet, got %d", len(violetData))
		}
		// February should have 1 catch
		if violetData[1].Count != 1 {
			t.Errorf("Expected 1 catch in February for Violet, got %d", violetData[1].Count)
		}
	} else {
		t.Error("Violet data not found")
	}
}

func TestCalculateOverallStats(t *testing.T) {
	startDate := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, time.March, 31, 0, 0, 0, 0, time.UTC)

	shinies := []dto.ShinyPokemon{
		{
			ShinyId: 1,
			Species: dto.PokemonSpecies{Name: "Pikachu", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"},
				Level: 10,
				Shiny: true,
			},
		},
		{
			ShinyId: 2,
			Species: dto.PokemonSpecies{Name: "Charmander", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"},
				Level: 12,
				Shiny: true,
			},
		},
	}

	calculator := NewStatsCalculator(shinies, startDate, endDate)
	overall := calculator.calculateOverallStats()

	if overall.TotalShiniesCaught != 2 {
		t.Errorf("Expected total 2 shinies, got %d", overall.TotalShiniesCaught)
	}

	if overall.TotalGamesPlayed != 2 {
		t.Errorf("Expected 2 games played, got %d", overall.TotalGamesPlayed)
	}

	if len(overall.CatchesPerGame) != 2 {
		t.Errorf("Expected 2 game breakdowns, got %d", len(overall.CatchesPerGame))
	}

	expectedAvg := 2.0 / 3.0 // 2 shinies across 3 months
	if overall.AveragePerMonth != expectedAvg {
		t.Errorf("Expected average %.4f, got %.4f", expectedAvg, overall.AveragePerMonth)
	}
}

func TestCalculateMonthlyTotals(t *testing.T) {
	startDate := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, time.February, 28, 0, 0, 0, 0, time.UTC)

	shinies := []dto.ShinyPokemon{
		{
			ShinyId: 1,
			Species: dto.PokemonSpecies{Name: "Pikachu", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"},
				Level: 10,
				Shiny: true,
			},
		},
		{
			ShinyId: 2,
			Species: dto.PokemonSpecies{Name: "Charmander", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"},
				Level: 12,
				Shiny: true,
			},
		},
		{
			ShinyId: 3,
			Species: dto.PokemonSpecies{Name: "Squirtle", Form: "Normal"},
			Details: dto.CatchDetails{
				CaughtAt: &[]time.Time{time.Date(2025, time.February, 10, 0, 0, 0, 0, time.UTC)}[0],
				Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"},
				Level: 8,
				Shiny: true,
			},
		},
	}

	calculator := NewStatsCalculator(shinies, startDate, endDate)
	monthly := calculator.calculateMonthlyTotals()

	if len(monthly) != 2 {
		t.Errorf("Expected 2 months of data, got %d", len(monthly))
	}

	// January should have 2 total catches from 2 games
	if monthly[0].TotalCatches != 2 {
		t.Errorf("Expected 2 total catches in January, got %d", monthly[0].TotalCatches)
	}
	if monthly[0].GamesContributing != 2 {
		t.Errorf("Expected 2 games contributing in January, got %d", monthly[0].GamesContributing)
	}

	// February should have 1 total catch from 1 game
	if monthly[1].TotalCatches != 1 {
		t.Errorf("Expected 1 total catch in February, got %d", monthly[1].TotalCatches)
	}
	if monthly[1].GamesContributing != 1 {
		t.Errorf("Expected 1 game contributing in February, got %d", monthly[1].GamesContributing)
	}
}

func TestValidateDateRange(t *testing.T) {
	svc := &ShinyService{}

	tests := []struct {
		name      string
		startDate string
		endDate   string
		shouldErr bool
	}{
		{
			name:      "valid dates",
			startDate: "2025-01-01",
			endDate:   "2025-12-31",
			shouldErr: false,
		},
		{
			name:      "empty dates should default",
			startDate: "",
			endDate:   "",
			shouldErr: false,
		},
		{
			name:      "invalid start date format",
			startDate: "invalid",
			endDate:   "2025-12-31",
			shouldErr: true,
		},
		{
			name:      "start after end",
			startDate: "2025-12-31",
			endDate:   "2025-01-01",
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := svc.validateAndParseDates(tt.startDate, tt.endDate)
			if (err != nil) != tt.shouldErr {
				t.Errorf("Expected error: %v, got: %v", tt.shouldErr, err)
			}
		})
	}
}
