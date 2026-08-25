<script setup lang="ts">

import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { PokemonStats, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
import PokemonInfo from '@/components/PokemonInfo.vue';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const stats = ref<PokemonStats[] | null>(null);
async function fetchPokemonStats() {
  try {
    const response = await apiFetch(`/api/pokemon/v1/`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<PokemonStats[]>;
    stats.value = data.data;
  } catch (error) {
    console.error('Error fetching Pokemon stats:', error);
  } finally {
    loading.value = false;
  }
}


onMounted(fetchPokemonStats);
</script>

<template>
  <div class="about" v-for="pokemon in stats" :key="pokemon.pokemon" v-if="stats">
    <div class="box mb-5">
      <PokemonInfo
        :pokemon="pokemon.pokemon"
        :nickname="pokemon.nickname"
        :region="pokemon.region"
        :current="pokemon.current"
        :target="pokemon.target"
        :includeLink="true"
      />
    </div>
  </div>
</template>

<style>

</style>
