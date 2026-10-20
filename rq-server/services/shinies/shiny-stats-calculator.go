package shinies

import (
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/dto"
	shiniesDTO "dezzles-apps/rq-server/model/dto/shinies"
	"sort"
	"time"
)

// StatsCalculator handles all calculations for shiny statistics
type StatsCalculator struct {
	shinies []dto.ShinyPokemon
	months  []time.Time
}

// NewStatsCalculator creates a new stats calculator for the given date range
func NewStatsCalculator(shinies []dto.ShinyPokemon, startDate, endDate time.Time) *StatsCalculator {
	calculator := &StatsCalculator{
		shinies: shinies,
		months:  generateMonths(startDate, endDate),
	}
	return calculator
}

// CalculateStats generates the complete stats response
func (sc *StatsCalculator) CalculateStats(startDate, endDate time.Time) *shiniesDTO.ShinyStatsResponse {
	response := &shiniesDTO.ShinyStatsResponse{
		DateRange: shiniesDTO.DateRange{
			StartDate: startDate,
			EndDate:   endDate,
		},
		MonthlyByGame: sc.calculateMonthlyByGame(),
		GameStats:     sc.calculateGameStats(),
		OverallStats:  sc.calculateOverallStats(),
		MonthlyTotals: sc.calculateMonthlyTotals(),
		Trends:        sc.calculateTrends(),
	}
	return response
}

// calculateMonthlyByGame groups shinies by game and month
func (sc *StatsCalculator) calculateMonthlyByGame() map[string][]shiniesDTO.MonthStats {
	result := make(map[string][]shiniesDTO.MonthStats)

	// Group shinies by game
	gameMap := make(map[string][]dto.ShinyPokemon)
	for _, shiny := range sc.shinies {
		if shiny.Details.Game != nil {
			gameMap[shiny.Details.Game.Value] = append(gameMap[shiny.Details.Game.Value], shiny)
		}
	}

	// For each game, calculate monthly stats
	for gameName, gameShinies := range gameMap {
		monthlyData := make([]shiniesDTO.MonthStats, len(sc.months))

		// Initialize monthly data
		for i := range monthlyData {
			monthlyData[i] = shiniesDTO.MonthStats{
				Year:       sc.months[i].Year(),
				Month:      int(sc.months[i].Month()),
				MonthName:  sc.months[i].Format("January"),
				Count:      0,
				Cumulative: 0,
			}
		}

		// Count shinies per month
		for _, shiny := range gameShinies {
			if shiny.Details.CaughtAt != nil {
				monthIdx := sc.getMonthIndex(*shiny.Details.CaughtAt)
				if monthIdx >= 0 && monthIdx < len(monthlyData) {
					monthlyData[monthIdx].Count++
				}
			}
		}

		// Calculate cumulative values
		cumulative := 0
		for i := range monthlyData {
			cumulative += monthlyData[i].Count
			monthlyData[i].Cumulative = cumulative
		}

		result[gameName] = monthlyData
	}

	return result
}

// calculateGameStats creates per-game summary statistics
func (sc *StatsCalculator) calculateGameStats() []shiniesDTO.GameStat {
	monthlyByGame := sc.calculateMonthlyByGame()
	gameStats := make([]shiniesDTO.GameStat, 0, len(monthlyByGame))

	for gameName, monthlyData := range monthlyByGame {
		totalCatches := 0
		if len(monthlyData) > 0 {
			totalCatches = monthlyData[len(monthlyData)-1].Cumulative
		}

		avgPerMonth := 0.0
		if len(sc.months) > 0 {
			avgPerMonth = float64(totalCatches) / float64(len(sc.months))
		}

		stat := shiniesDTO.GameStat{
			Game: &model.KeyValuePair{
				Key:   sc.getGameKey(gameName),
				Value: gameName,
			},
			TotalCatches:           totalCatches,
			MonthlyBreakdown:       monthlyData,
			AverageCatchesPerMonth: avgPerMonth,
		}
		gameStats = append(gameStats, stat)
	}

	// Sort by total catches descending
	sort.Slice(gameStats, func(i, j int) bool {
		return gameStats[i].TotalCatches > gameStats[j].TotalCatches
	})

	return gameStats
}

// calculateOverallStats creates aggregate statistics
func (sc *StatsCalculator) calculateOverallStats() shiniesDTO.OverallStats {
	gameStats := sc.calculateGameStats()

	totalCatches := 0
	gamesPlayed := 0
	catchesPerGame := make([]shiniesDTO.CatchPerGameBreakdown, 0)

	for rank, stat := range gameStats {
		if stat.TotalCatches > 0 {
			gamesPlayed++
		}
		totalCatches += stat.TotalCatches
		catchesPerGame = append(catchesPerGame, shiniesDTO.CatchPerGameBreakdown{
			Game:    stat.Game,
			Rank:    rank + 1,
			Catches: stat.TotalCatches,
		})
	}

	avgPerMonth := 0.0
	if len(sc.months) > 0 {
		avgPerMonth = float64(totalCatches) / float64(len(sc.months))
	}

	return shiniesDTO.OverallStats{
		TotalShiniesCaught: totalCatches,
		AveragePerMonth:    avgPerMonth,
		CatchesPerGame:     catchesPerGame,
		TotalGamesPlayed:   gamesPlayed,
	}
}

// calculateMonthlyTotals creates cross-game monthly aggregates
func (sc *StatsCalculator) calculateMonthlyTotals() []shiniesDTO.MonthlyTotal {
	monthlyByGame := sc.calculateMonthlyByGame()
	monthlyTotals := make([]shiniesDTO.MonthlyTotal, len(sc.months))

	for monthIdx := range sc.months {
		totalCatches := 0
		gamesContributing := 0
		var topGame *model.KeyValuePair
		topGameCount := 0

		for gameName, monthlyData := range monthlyByGame {
			if monthIdx < len(monthlyData) && monthlyData[monthIdx].Count > 0 {
				gamesContributing++
				totalCatches += monthlyData[monthIdx].Count

				if monthlyData[monthIdx].Count > topGameCount {
					topGameCount = monthlyData[monthIdx].Count
					topGame = &model.KeyValuePair{
						Key:   sc.getGameKey(gameName),
						Value: gameName,
					}
				}
			}
		}

		monthlyTotals[monthIdx] = shiniesDTO.MonthlyTotal{
			Month:             monthIdx,
			MonthName:         sc.months[monthIdx].Format("January"),
			TotalCatches:      totalCatches,
			GamesContributing: gamesContributing,
			TopGame:           topGame,
			TopGameCount:      topGameCount,
		}
	}

	return monthlyTotals
}

// calculateTrends identifies trend patterns
func (sc *StatsCalculator) calculateTrends() shiniesDTO.Trends {
	gameStats := sc.calculateGameStats()
	monthlyByGame := sc.calculateMonthlyByGame()

	trends := shiniesDTO.Trends{
		UpwardTrend:   make([]shiniesDTO.GameTrend, 0),
		DownwardTrend: make([]shiniesDTO.GameTrend, 0),
		Streaks:       make([]shiniesDTO.GameStreak, 0),
	}

	var mostActive, leastActive *model.KeyValuePair
	mostActiveCount := 0
	leastActiveCount := -1

	// Analyze each game
	for _, gameStat := range gameStats {
		if gameStat.TotalCatches == 0 {
			continue
		}

		// Track most/least active
		if gameStat.TotalCatches > mostActiveCount {
			mostActiveCount = gameStat.TotalCatches
			mostActive = gameStat.Game
		}
		if leastActiveCount < 0 || (gameStat.TotalCatches > 0 && gameStat.TotalCatches < leastActiveCount) {
			leastActiveCount = gameStat.TotalCatches
			leastActive = gameStat.Game
		}

		// Calculate trend
		monthlyData := monthlyByGame[gameStat.Game.Value]
		if len(monthlyData) >= 2 {
			firstMonth := monthlyData[0].Count
			lastMonth := monthlyData[len(monthlyData)-1].Count
			momentum := 0.0
			if firstMonth > 0 {
				momentum = float64(lastMonth-firstMonth) / float64(firstMonth)
			}

			if momentum > 0.1 { // 10% increase
				trends.UpwardTrend = append(trends.UpwardTrend, shiniesDTO.GameTrend{
					Game:       gameStat.Game,
					Momentum:   momentum,
					StartMonth: 0,
					EndMonth:   len(monthlyData) - 1,
					StartCount: firstMonth,
					EndCount:   lastMonth,
				})
			} else if momentum < -0.1 { // 10% decrease
				trends.DownwardTrend = append(trends.DownwardTrend, shiniesDTO.GameTrend{
					Game:       gameStat.Game,
					Momentum:   momentum,
					StartMonth: 0,
					EndMonth:   len(monthlyData) - 1,
					StartCount: firstMonth,
					EndCount:   lastMonth,
				})
			}
		}

		// Find longest streak
		streak := sc.findLongestStreak(monthlyData, gameStat.Game)
		if streak != nil {
			trends.Streaks = append(trends.Streaks, *streak)
		}
	}

	trends.MostActive = mostActive
	trends.MostActiveCount = mostActiveCount
	trends.LeastActive = leastActive
	trends.LeastActiveCount = leastActiveCount

	return trends
}

// findLongestStreak finds the longest consecutive months with catches
func (sc *StatsCalculator) findLongestStreak(monthlyData []shiniesDTO.MonthStats, game *model.KeyValuePair) *shiniesDTO.GameStreak {
	if len(monthlyData) == 0 {
		return nil
	}

	maxStreakLength := 0
	maxStreakStart := -1
	maxStreakEnd := -1
	maxStreakTotal := 0

	currentStreakLength := 0
	currentStreakStart := -1
	currentStreakTotal := 0

	for i, month := range monthlyData {
		if month.Count > 0 {
			if currentStreakLength == 0 {
				currentStreakStart = i
			}
			currentStreakLength++
			currentStreakTotal += month.Count
		} else {
			if currentStreakLength > maxStreakLength {
				maxStreakLength = currentStreakLength
				maxStreakStart = currentStreakStart
				maxStreakEnd = i - 1
				maxStreakTotal = currentStreakTotal
			}
			currentStreakLength = 0
			currentStreakStart = -1
			currentStreakTotal = 0
		}
	}

	// Check final streak
	if currentStreakLength > maxStreakLength {
		maxStreakLength = currentStreakLength
		maxStreakStart = currentStreakStart
		maxStreakEnd = len(monthlyData) - 1
		maxStreakTotal = currentStreakTotal
	}

	if maxStreakLength == 0 {
		return nil
	}

	return &shiniesDTO.GameStreak{
		Game:          game,
		StreakLength:  maxStreakLength,
		StartMonth:    maxStreakStart,
		EndMonth:      maxStreakEnd,
		TotalInStreak: maxStreakTotal,
	}
}

// getMonthIndex returns the index of the month for a given date
func (sc *StatsCalculator) getMonthIndex(date time.Time) int {
	for i, month := range sc.months {
		if month.Year() == date.Year() && month.Month() == date.Month() {
			return i
		}
	}
	return -1
}

// getGameKey converts a game name to its database key (simplified reverse mapping)
// This is a helper; ideally would use the repository's game data
func (sc *StatsCalculator) getGameKey(gameName string) string {
	// Map common game names to keys - can be enhanced with a proper lookup
	gameKeyMap := map[string]string{
		"X":               "X",
		"Y":               "Y",
		"Omega Ruby":      "OMEGA_RUBY",
		"Alpha Sapphire":  "ALPHA_SAPPHIRE",
		"Sun":             "SUN",
		"Moon":            "MOON",
		"Ultra Sun":       "ULTRA_SUN",
		"Ultra Moon":      "ULTRA_MOON",
		"Brilliant Diamond": "BRILLIANT_DIAMOND",
		"Shining Pearl":   "SHINING_PEARL",
		"Legends: Arceus": "LEGENDS_ARCEUS",
		"Sword":           "SWORD",
		"Shield":          "SHIELD",
		"Scarlet":         "SCARLET",
		"Violet":          "VIOLET",
		"Legends: Z-A":    "LZA",
	}
	if key, exists := gameKeyMap[gameName]; exists {
		return key
	}
	return gameName // fallback to name itself
}

// generateMonths creates a list of month start dates for the given range
func generateMonths(startDate, endDate time.Time) []time.Time {
	months := make([]time.Time, 0)

	// Normalize to first day of month
	current := time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location())
	end := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, endDate.Location())

	for current.Before(end) || current.Equal(end) {
		months = append(months, current)
		current = current.AddDate(0, 1, 0)
	}

	return months
}
