<script setup lang="ts">
import { onMounted, ref, computed } from 'vue';
import { Line } from 'vue-chartjs';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Filler } from 'chart.js';
import type { ShinyStatsResponse } from '@/types/stats';
import type { Response } from '@/types/responses';
import { useApi } from '@/composables/useApi';
import { useDates } from '@/composables/useDates';
import API from '@/composables/endpoints';
import Loading from '@/components/Loading.vue';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Filler);



const loading = ref(true);
const { apiFetch } = useApi();
const { toLocal } = useDates();
const stats = ref<ShinyStatsResponse | null>(null);
const error = ref<string | null>(null);

async function fetchStats() {
  try {
    loading.value = true;
    error.value = null;
    const response = await apiFetch(API.Shinies.GetStats);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<ShinyStatsResponse>;
    if (data.error) {
      throw new Error(data.error);
    }
    stats.value = data.data;
  } catch (err) {
    console.error('Error fetching shiny stats:', err);
    error.value = err instanceof Error ? err.message : 'Failed to load stats';
  } finally {
    loading.value = false;
  }
}

const gameRankings = computed(() => {
  return stats.value?.overallStats.catchesPerGame ?? [];
});

const monthlyChartData = computed(() => {
  const gameStats = stats.value?.gameStats ?? [];
  if (gameStats.length === 0) {
    return null;
  }
  
  // Get all unique months from all games, sorted chronologically
  const monthsMap = new Map<string, { year: number; month: number; name: string }>();
  gameStats.forEach(game => {
    game.monthlyBreakdown?.forEach(month => {
      const key = `${month.year}-${month.month}`;
      if (!monthsMap.has(key)) {
        monthsMap.set(key, { year: month.year, month: month.month, name: month.monthName });
      }
    });
  });
  
  const sortedMonths = Array.from(monthsMap.values()).sort((a, b) => {
    if (a.year !== b.year) return a.year - b.year;
    return a.month - b.month;
  });
  
  const labels = sortedMonths.map(m => m.name);
  
  // Color palette for games
  const colors = [
    '#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF',
    '#FF9F40', '#FF6384', '#C9CBCF', '#4BC0C0', '#FF6384',
    '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF', '#FF9F40'
  ];
  
  const datasets = gameStats.map((game, idx) => {
    const gameData = sortedMonths.map(month => {
      const monthStat = game.monthlyBreakdown?.find(
        m => m.year === month.year && m.month === month.month
      );
      return monthStat?.count ?? 0;
    });
    
    return {
      label: game.game.value,
      data: gameData,
      borderColor: colors[idx % colors.length],
      backgroundColor: colors[idx % colors.length],
      fill: true,
      tension: 0.3,
      pointRadius: 4,
      pointBackgroundColor: colors[idx % colors.length],
      pointBorderColor: '#fff',
      pointBorderWidth: 2,
    };
  });
  
  return {
    labels,
    datasets
  };
});

const monthlyChartOptions = computed(() => ({
  responsive: true,
  maintainAspectRatio: true,
  plugins: {
    legend: {
      display: true,
      position: 'top' as const,
    },
    title: {
      display: false,
    }
  },
  scales: {
    x: {
      stacked: false,
    },
    y: {
      stacked: true,
      beginAtZero: true,
      ticks: {
        stepSize: 1,
      }
    }
  }
}));

const trendingUp = computed(() => {
  return stats.value?.trends.upwardTrend ?? [];
});

const trendingDown = computed(() => {
  return stats.value?.trends.downwardTrend ?? [];
});

const streaks = computed(() => {
  return stats.value?.trends.streaks ?? [];
});

onMounted(fetchStats);
</script>

<template>
  <div class="shiny-stats">
    <Loading v-if="loading" />
    
    <div v-if="error" class="error-container">
      <v-alert type="error" :title="'Error Loading Stats'" :text="error" closable />
    </div>

    <div v-if="!loading && stats" class="stats-content">
      <!-- Header Stats -->
      <v-row>
        <v-col cols="12" sm="6" md="3">
          <v-card>
            <v-card-title class="text-h6">Total Shinies Caught</v-card-title>
            <v-card-text class="text-h4 font-weight-bold text-green">
              {{ stats.overallStats.totalShiniesCaught }}
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card>
            <v-card-title class="text-h6">Games Played</v-card-title>
            <v-card-text class="text-h4 font-weight-bold text-blue">
              {{ stats.overallStats.totalGamesPlayed }}
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card>
            <v-card-title class="text-h6">Avg Per Month</v-card-title>
            <v-card-text class="text-h4 font-weight-bold text-orange">
              {{ stats.overallStats.averagePerMonth.toFixed(2) }}
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card>
            <v-card-title class="text-h6">Most Active</v-card-title>
            <v-card-text class="text-base font-weight-bold">
              <span v-if="stats.trends.mostActive">{{ stats.trends.mostActive.value }} ({{ stats.trends.mostActiveCount }})</span>
              <span v-else>—</span>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Monthly Totals Chart -->
      <v-row class="mt-4">
        <v-col cols="12">
          <v-card title="Monthly Catches by Game">
            <v-card-text>
              <Line
                v-if="monthlyChartData"
                :data="monthlyChartData"
                :options="monthlyChartOptions"
                :style="{ height: '300px', position: 'relative', margin: 'auto' }"
              />
              <div v-else class="text-center text-gray py-8">
                No monthly data available
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Game Rankings -->
      <v-row class="mt-4">
        <v-col cols="12">
          <v-card title="Game Rankings">
            <v-table>
              <thead>
                <tr>
                  <th>Rank</th>
                  <th>Game</th>
                  <th class="text-right">Catches</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="game in gameRankings" :key="game.game.key">
                  <td>{{ game.rank }}</td>
                  <td>{{ game.game.value }}</td>
                  <td class="text-right font-weight-bold">{{ game.catches }}</td>
                </tr>
              </tbody>
            </v-table>
          </v-card>
        </v-col>
      </v-row>

      <!-- Streaks -->
      <v-row class="mt-4">
        <v-col cols="12" md="6">
          <v-card title="Catching Streaks">
            <v-list v-if="streaks.length > 0">
              <v-list-item v-for="streak in streaks" :key="streak.game.key">
                <v-list-item-title>{{ streak.game.value }}</v-list-item-title>
                <template #subtitle>
                  {{ streak.streakLength }} months | Total: {{ streak.totalInStreak }}
                </template>
              </v-list-item>
            </v-list>
            <v-card-text v-else class="text-center text-gray">
              No streaks recorded
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Trends -->
      <v-row class="mt-4">
        <v-col cols="12" md="6">
          <v-card title="Upward Trends 📈">
            <v-list v-if="trendingUp.length > 0">
              <v-list-item v-for="trend in trendingUp" :key="trend.game.key">
                <v-list-item-title>{{ trend.game.value }}</v-list-item-title>
                <template #subtitle>
                  Momentum: {{ trend.momentum.toFixed(2) }} | {{ trend.startCount }} → {{ trend.endCount }}
                </template>
              </v-list-item>
            </v-list>
            <v-card-text v-else class="text-center text-gray">
              No upward trends
            </v-card-text>
          </v-card>
        </v-col>

        <v-col cols="12" md="6">
          <v-card title="Downward Trends 📉">
            <v-list v-if="trendingDown.length > 0">
              <v-list-item v-for="trend in trendingDown" :key="trend.game.key">
                <v-list-item-title>{{ trend.game.value }}</v-list-item-title>
                <template #subtitle>
                  Momentum: {{ trend.momentum.toFixed(2) }} | {{ trend.startCount }} → {{ trend.endCount }}
                </template>
              </v-list-item>
            </v-list>
            <v-card-text v-else class="text-center text-gray">
              No downward trends
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Least Active -->
      <v-row class="mt-4">
        <v-col cols="12" md="6">
          <v-card>
            <v-card-title class="text-h6">Least Active Game</v-card-title>
            <v-card-text>
              <div v-if="stats.trends.leastActive" class="text-h5 font-weight-bold">{{ stats.trends.leastActive.value }}</div>
              <div v-if="stats.trends.leastActive" class="text-caption text-gray">{{ stats.trends.leastActiveCount }} catches</div>
              <div v-else class="text-caption text-gray">No data</div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" md="6">
          <v-card>
            <v-card-title class="text-h6">Date Range</v-card-title>
            <v-card-text>
              <div class="text-caption">From: {{ toLocal(stats.dateRange.startDate) }}</div>
              <div class="text-caption">To: {{ toLocal(stats.dateRange.endDate) }}</div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<style scoped>
.shiny-stats {
  padding: 16px;
}

.stats-content {
  width: 100%;
}

.error-container {
  margin-bottom: 16px;
}

.text-gray {
  color: #999;
}

v-table {
  width: 100%;
}
</style>