package shinies

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	"encoding/json"
	"testing"
	"time"
)

// TestFullStatsResponse verifies the complete stats response structure
func TestFullStatsResponse(t *testing.T) {
	startDate := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC)  // Extended to include all test data

	// Create representative test data matching the original API structure
	shinies := []dto.ShinyPokemon{
		// Scarlet catches
		{ShinyId: 1, Species: dto.PokemonSpecies{Name: "Sprigatito", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.January, 5, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},
		{ShinyId: 2, Species: dto.PokemonSpecies{Name: "Pawmi", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},
		{ShinyId: 3, Species: dto.PokemonSpecies{Name: "Fidough", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.February, 8, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},
		{ShinyId: 4, Species: dto.PokemonSpecies{Name: "Tandemaus", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.March, 15, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},
		{ShinyId: 5, Species: dto.PokemonSpecies{Name: "Lechonk", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.April, 3, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},

		// Violet catches
		{ShinyId: 6, Species: dto.PokemonSpecies{Name: "Quaxly", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.August, 12, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"}, Shiny: true}},
		{ShinyId: 7, Species: dto.PokemonSpecies{Name: "Paldean Wooper", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.September, 5, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"}, Shiny: true}},
		{ShinyId: 8, Species: dto.PokemonSpecies{Name: "Tadbulb", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.October, 1, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"}, Shiny: true}},

		// Ultra Sun catches
		{ShinyId: 9, Species: dto.PokemonSpecies{Name: "Litten", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.February, 20, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "ULTRA_SUN", Value: "Ultra Sun"}, Shiny: true}},

		// Legends: Z-A catches
		{ShinyId: 10, Species: dto.PokemonSpecies{Name: "Sprigatito", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.November, 10, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "LZA", Value: "Legends: Z-A"}, Shiny: true}},
		{ShinyId: 11, Species: dto.PokemonSpecies{Name: "Pawmi", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.December, 15, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "LZA", Value: "Legends: Z-A"}, Shiny: true}},
		{ShinyId: 12, Species: dto.PokemonSpecies{Name: "Fidough", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.December, 20, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "LZA", Value: "Legends: Z-A"}, Shiny: true}},
	}

	calculator := NewStatsCalculator(shinies, startDate, endDate)
	stats := calculator.CalculateStats(startDate, endDate)

	// Verify response structure
	if stats.DateRange.StartDate != startDate {
		t.Errorf("DateRange.StartDate mismatch")
	}
	if stats.DateRange.EndDate != endDate {
		t.Errorf("DateRange.EndDate mismatch")
	}

	// Verify monthlyByGame
	if len(stats.MonthlyByGame) != 4 {
		t.Errorf("Expected 4 games, got %d", len(stats.MonthlyByGame))
	}

	// Verify gameStats
	if len(stats.GameStats) != 4 {
		t.Errorf("Expected 4 game stats, got %d", len(stats.GameStats))
	}

	// Verify overall stats
	if stats.OverallStats.TotalShiniesCaught != 12 {
		t.Errorf("Expected 12 total shinies, got %d", stats.OverallStats.TotalShiniesCaught)
	}
	if stats.OverallStats.TotalGamesPlayed != 4 {
		t.Errorf("Expected 4 games played, got %d", stats.OverallStats.TotalGamesPlayed)
	}

	// Verify monthly totals (should have 16 months: Oct 2024 - Jan 2026)
	expectedMonths := 16
	if len(stats.MonthlyTotals) != expectedMonths {
		t.Errorf("Expected %d months, got %d", expectedMonths, len(stats.MonthlyTotals))
	}

	// Verify first month (October 2024) has no catches
	if stats.MonthlyTotals[0].TotalCatches != 0 {
		t.Errorf("Expected 0 catches in first month, got %d", stats.MonthlyTotals[0].TotalCatches)
	}

	// Verify January 2025 (month index 3) has 2 Scarlet catches
	if stats.MonthlyTotals[3].TotalCatches != 2 {
		t.Errorf("Expected 2 catches in January 2025, got %d", stats.MonthlyTotals[3].TotalCatches)
	}
	if stats.MonthlyTotals[3].GamesContributing != 1 {
		t.Errorf("Expected 1 game in January 2025, got %d", stats.MonthlyTotals[3].GamesContributing)
	}

	// Verify trends
	if len(stats.Trends.Streaks) > 0 {
		t.Logf("Found %d streaks in data", len(stats.Trends.Streaks))
	}

	// Verify serialization to JSON works
	jsonData, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Failed to marshal response to JSON: %v", err)
	}

	if len(jsonData) == 0 {
		t.Errorf("JSON serialization produced empty output")
	}

	t.Logf("✅ Response structure verified. JSON size: %d bytes", len(jsonData))
}

// TestResponseJSON demonstrates the actual JSON structure that will be returned
func TestResponseJSON(t *testing.T) {
	startDate := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, time.March, 31, 0, 0, 0, 0, time.UTC)

	shinies := []dto.ShinyPokemon{
		{ShinyId: 1, Species: dto.PokemonSpecies{Name: "Pikachu", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},
		{ShinyId: 2, Species: dto.PokemonSpecies{Name: "Charmander", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.February, 10, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "SCARLET", Value: "Scarlet"}, Shiny: true}},
		{ShinyId: 3, Species: dto.PokemonSpecies{Name: "Squirtle", Form: "Normal"}, Details: dto.CatchDetails{CaughtAt: &[]time.Time{time.Date(2025, time.March, 5, 0, 0, 0, 0, time.UTC)}[0], Game: &model.KeyValuePair{Key: "VIOLET", Value: "Violet"}, Shiny: true}},
	}

	calculator := NewStatsCalculator(shinies, startDate, endDate)
	stats := calculator.CalculateStats(startDate, endDate)

	jsonData, _ := json.MarshalIndent(stats, "", "  ")
	t.Logf("Generated Stats Response:\n%s", string(jsonData))
}
