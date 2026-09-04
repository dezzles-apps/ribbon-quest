<script setup lang="ts">

import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { PokemonStats, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
import { Ribbons } from '@/composables/endpoints';
import PokemonInfo from '@/components/PokemonInfo.vue';
import Loading from '@/components/Loading.vue';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const stats = ref<PokemonStats[] | null>(null);
async function fetchPokemonStats() {
  try {
    const response = await apiFetch(Ribbons.GetAllPokemon);
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
  <Loading v-if="loading" />
  <div class="about" v-for="pokemon in stats" :key="pokemon.pokemon" v-if="stats">
    <div class="box mb-5">
      <PokemonInfo
        :details="pokemon"
        :current="pokemon.current"
        :total="pokemon.total"
        :includeLink="true"
      />
    </div>
  </div>
</template>

<style>

</style>
