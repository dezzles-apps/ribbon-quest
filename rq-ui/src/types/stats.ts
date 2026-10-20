import type { KeyValuePair } from './data'

interface DateRange {
  startDate: string;
  endDate: string;
}

interface MonthStats {
  year: number;
  month: number;
  monthName: string;
  count: number;
  cumulative: number;
}

interface GameStat {
  game: KeyValuePair;
  totalCatches: number;
  monthlyBreakdown: MonthStats[];
  averageCatchesPerMonth: number;
}

interface CatchPerGameBreakdown {
  game: KeyValuePair;
  rank: number;
  catches: number;
}

interface OverallStats {
  totalShiniesCaught: number;
  averagePerMonth: number;
  catchesPerGame: CatchPerGameBreakdown[];
  totalGamesPlayed: number;
}

interface MonthlyTotal {
  month: number;
  monthName: string;
  totalCatches: number;
  gamesContributing: number;
  topGame: KeyValuePair;
  topGameCount: number;
}

interface GameTrend {
  game: KeyValuePair;
  momentum: number;
  startMonth: number;
  endMonth: number;
  startCount: number;
  endCount: number;
}

interface GameStreak {
  game: KeyValuePair;
  streakLength: number;
  startMonth: number;
  endMonth: number;
  totalInStreak: number;
}

interface Trends {
  upwardTrend: GameTrend[];
  downwardTrend: GameTrend[];
  mostActive: KeyValuePair;
  mostActiveCount: number;
  leastActive: KeyValuePair;
  leastActiveCount: number;
  streaks: GameStreak[];
}

interface ShinyStatsResponse {
  dateRange: DateRange;
  monthlyByGame: Record<string, MonthStats[]>;
  gameStats: GameStat[];
  overallStats: OverallStats;
  monthlyTotals: MonthlyTotal[];
  trends: Trends;
}

export type {
  DateRange,
  MonthStats,
  GameStat,
  CatchPerGameBreakdown,
  OverallStats,
  MonthlyTotal,
  GameTrend,
  GameStreak,
  Trends,
  ShinyStatsResponse
}