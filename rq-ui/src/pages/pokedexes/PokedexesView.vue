<script setup lang="ts">

import { onMounted, ref } from 'vue';
import type { Response } from '@/types/responses';
import type { Pokedex } from '@/types/pokedex';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import PokedexInfo from '@/components/pokedexes/PokedexInfo.vue';
import Loading from '@/components/Loading.vue';
const loading = ref(true);
const { apiFetch } = useApi();

const stats = ref<Pokedex[] | null>(null);
async function fetchPokemonStats() {
  try {
    const response = await apiFetch(API.Pokedexes.GetAllPokedexes);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<Pokedex[]>;
    stats.value = data.data;
  } catch (error) {
    console.error('Error fetching Pokedexes stats:', error);
  } finally {
    loading.value = false;
  }
}

onMounted(fetchPokemonStats);
</script>

<template>
  <Loading v-if="loading" />
  <div class="about" v-for="dex in stats" :key="dex.name" v-if="stats">
    <div class="box mb-5">
      <PokedexInfo
        :details="dex"
        :includeLink="true"
      />
    </div>
  </div>
</template>

<style>

</style>
