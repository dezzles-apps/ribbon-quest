package shinies

import (
	"dezzles-apps/rq-server/model"
	"time"
)

// ShinyStatsResponse is the top-level response for the shiny stats endpoint
type ShinyStatsResponse struct {
	DateRange      DateRange               `json:"dateRange"`
	MonthlyByGame  map[string][]MonthStats `json:"monthlyByGame"`
	GameStats      []GameStat              `json:"gameStats"`
	OverallStats   OverallStats            `json:"overallStats"`
	MonthlyTotals  []MonthlyTotal          `json:"monthlyTotals"`
	Trends         Trends                  `json:"trends"`
}

// DateRange represents the requested date range
type DateRange struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

// MonthStats represents the monthly catch statistics for a single game
type MonthStats struct {
	Year       int    `json:"year"`
	Month      int    `json:"month"`
	MonthName  string `json:"monthName"`
	Count      int    `json:"count"`
	Cumulative int    `json:"cumulative"`
}

// GameStat represents statistics for a single game across the date range
type GameStat struct {
	Game              *model.KeyValuePair `json:"game"`
	TotalCatches      int                 `json:"totalCatches"`
	MonthlyBreakdown  []MonthStats        `json:"monthlyBreakdown"`
	AverageCatchesPerMonth float64         `json:"averageCatchesPerMonth"`
}

// OverallStats represents aggregate statistics across all games
type OverallStats struct {
	TotalShiniesCaught int                     `json:"totalShiniesCaught"`
	AveragePerMonth    float64                 `json:"averagePerMonth"`
	CatchesPerGame     []CatchPerGameBreakdown `json:"catchesPerGame"`
	TotalGamesPlayed   int                     `json:"totalGamesPlayed"`
}

// CatchPerGameBreakdown represents a game's catch count in the overall stats
type CatchPerGameBreakdown struct {
	Game   *model.KeyValuePair `json:"game"`
	Rank   int                 `json:"rank"`
	Catches int                `json:"catches"`
}

// MonthlyTotal represents aggregate statistics for a single month
type MonthlyTotal struct {
	Month            int                 `json:"month"`
	MonthName        string              `json:"monthName"`
	TotalCatches     int                 `json:"totalCatches"`
	GamesContributing int                `json:"gamesContributing"`
	TopGame          *model.KeyValuePair `json:"topGame"`
	TopGameCount     int                 `json:"topGameCount"`
}

// Trends represents trend analysis across games
type Trends struct {
	UpwardTrend    []GameTrend         `json:"upwardTrend"`
	DownwardTrend  []GameTrend         `json:"downwardTrend"`
	MostActive     *model.KeyValuePair `json:"mostActive"`
	MostActiveCount int                `json:"mostActiveCount"`
	LeastActive    *model.KeyValuePair `json:"leastActive"`
	LeastActiveCount int               `json:"leastActiveCount"`
	Streaks        []GameStreak        `json:"streaks"`
}

// GameTrend represents a game's trend analysis
type GameTrend struct {
	Game       *model.KeyValuePair `json:"game"`
	Momentum   float64             `json:"momentum"`
	StartMonth int                 `json:"startMonth"`
	EndMonth   int                 `json:"endMonth"`
	StartCount int                 `json:"startCount"`
	EndCount   int                 `json:"endCount"`
}

// GameStreak represents a game's consecutive months of catches
type GameStreak struct {
	Game              *model.KeyValuePair `json:"game"`
	StreakLength      int                 `json:"streakLength"`
	StartMonth        int                 `json:"startMonth"`
	EndMonth          int                 `json:"endMonth"`
	TotalInStreak     int                 `json:"totalInStreak"`
}
