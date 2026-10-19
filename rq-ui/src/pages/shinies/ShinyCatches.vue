<script setup lang="ts">

import { onMounted, ref } from 'vue';
import type { Response } from '@/types/responses';
import type { ShinyPokemon } from '@/types/ribbons';
import { useDates } from '@/composables/useDates';
import API from '@/composables/endpoints';
import { useApi } from '@/composables/useApi';
import Loading from '@/components/Loading.vue';

const loading = ref(true);
const { apiFetch } = useApi();
const { toLocal } = useDates();
const stats = ref<ShinyPokemon[] | null>(null);
async function fetchPokemonStats() {
  try {
    const response = await apiFetch(API.Shinies.GetAllShinies);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<ShinyPokemon[]>;
    stats.value = data.data;
  } catch (error) {
    console.error('Error fetching Pokemon stats:', error);
  } finally {
    loading.value = false;
  }
}
 
const headers = [
  { title: 'Dex No.', value: 'species.pokedexId' },
  { title: 'Species', key: 'Species', value: (item:ShinyPokemon) => `${item.species.name} (${item.species.form})` },
  { title: 'Level', value: 'details.level' },
  { title: 'Method', value: 'details.method.value' },
  { title: 'Nature', value: 'details.nature.value' },
  { title: 'Characteristic', value: 'details.characteristic.value' },
  { title: 'Game', value: 'details.game.value' },
  { title: 'Caught At', key: 'details.caughtAt', value: (item:ShinyPokemon) => toLocal(item.details.caughtAt) }
] 

onMounted(fetchPokemonStats);
</script>

<template>
  <Loading v-if="loading" />
  <v-data-table-virtual
    v-if="!loading && stats"
    :headers="headers"
    :items="stats"
    item-value="name"
    item-key="shinyId"
  ></v-data-table-virtual>
</template>

<style>

</style>
