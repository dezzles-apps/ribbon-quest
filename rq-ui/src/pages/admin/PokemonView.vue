<script setup lang="ts">

import { onMounted, ref } from 'vue';
import type { PokemonStats } from '@/types/ribbons';
import type { Response } from '@/types/responses';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import Loading from '@/components/Loading.vue';
const loading = ref(true);
const { apiFetch } = useApi();

const stats = ref<PokemonStats[] | null>(null);
async function fetchPokemonStats() {
  try {
    const response = await apiFetch(API.Ribbons.GetAllPokemon);
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
  <Loading v-if="loading"/>
  <v-btn
    color="primary"
    :to="{ name: 'admin-add-ribbon-pokemon' }"
  >
    Add New Pokemon
  </v-btn>
  <v-list>
    <v-list-subheader>Pokemon</v-list-subheader>

    <v-list-item
      v-for="pokemon in stats"
      :key="pokemon.pokemon"
      :value="pokemon.details.nickname"
      :to="{ name: 'admin-edit-ribbon-pokemon', params: { pokemonId: pokemon.pokemon }}"
    >
      <template v-slot:prepend>
        <v-icon
          class="mr-2"
          icon="mdi-pokeball"
          :color="pokemon.details.caughtAt ? 'white': 'grey'"
        ></v-icon>
      </template>

      <v-list-item-title v-text="pokemon.pokemon"></v-list-item-title>
    </v-list-item>
  </v-list>
</template>

<style>

</style>
