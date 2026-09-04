<script setup lang="ts">

import { onMounted, ref } from 'vue';
import type { Response, GameWithStats } from '@/types/api';
import { useApi } from '@/composables/useApi';
import { Ribbons } from '@/composables/endpoints';
import GameInfo from '@/components/GameInfo.vue';

const loading = ref(true);
const { apiFetch } = useApi();

const stats = ref<GameWithStats[] | null>(null);
async function fetchGameStats() {
  try {
    const response = await apiFetch(Ribbons.GetAllGames);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<GameWithStats[]>;
    stats.value = data.data;
  } catch (error) {
    console.error('Error fetching game stats:', error);
  } finally {
    loading.value = false;
  }
}


onMounted(fetchGameStats);
</script>

<template>
  <div class="about" v-for="game in stats" :key="game.gameKey" v-if="stats">
    <div class="box mb-5">
      <GameInfo
        :gameKey="game.gameKey"
        :name="game.name"
        :current="game.achieved"
        :total="game.total"
        :includeLink="true"
      />
    </div>
  </div>
</template>

<style>

</style>
