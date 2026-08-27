<script setup lang="ts">

import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { PokemonStats, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
const route = useRoute();
const loading = ref(true);
const { apiFetch, apiFetchWithAuth } = useApi();

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
const natures = ref<string[]>([
  '', 'Adamant', 'Bashful', 'Bold', 'Brave', 'Calm', 'Careful', 'Docile',
  'Gentle', 'Hardy', 'Hasty', 'Impish', 'Jolly', 'Lax', 'Lonely', 'Mild',
  'Modest', 'Naive', 'Naughty', 'Quiet', 'Quirky', 'Rash', 'Relaxed', 'Sassy', 'Serious', 'Timid'
]);

const updatingPokemon = ref<Map<String, boolean>>(new Map());

function catchPokemon(pokemon: PokemonStats) {
  updatingPokemon.value.set(pokemon.pokemon, true);
  apiFetchWithAuth(`/api/pokemon/v1/${pokemon.pokemon}/catch`, {
    method: 'POST'
  }).then(async response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const { data } = await response.json();
    pokemon.caughtAt = data.caughtAt;
    pokemon.nickname = data.nickname ?? '';
    pokemon.nature = data.nature ?? '';
  })
  .catch(error => {
    console.error('Error catching Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value.set(pokemon.pokemon, false);
  });
}

function updatePokemon(pokemon: PokemonStats) {
  updatingPokemon.value.set(pokemon.pokemon, true);
  apiFetchWithAuth(`/api/pokemon/v1/${pokemon.pokemon}`, {
    method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        shiny: pokemon.shiny,
        nature: pokemon.nature,
        characteristic: pokemon.characteristic,
        nickname: pokemon.nickname
      })
  }).then(async response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const { data } = await response.json();
    pokemon.caughtAt = data.caughtAt;
    pokemon.nickname = data.nickname ?? '';
    pokemon.nature = data.nature ?? '';
    pokemon.shiny = data.shiny?? false;
    pokemon.characteristic = data.characteristic?? '';
  })
  .catch(error => {
    console.error('Error catching Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value.set(pokemon.pokemon, false);
  });
}

onMounted(fetchPokemonStats);
</script>

<template>
  <div class="about" v-for="(pokemon, idx) in stats" :key="pokemon.pokemon" v-if="stats">
    <div class="box mb-5">
      <h1 class="title is-3">{{ pokemon.pokemon }}</h1>
      <div v-if="pokemon.caughtAt">
        <div class="field">
          <p class="control has-icons-left has-icons">
            <input class="input" type="text" placeholder="Nickname" v-model="pokemon.nickname">
            <span class="icon is-small is-left">
              <i class="mdi mdi-account"></i>
            </span>
          </p>
        </div>
        <div class="field">
          <p class="control select">
            <select placeholder="Nature" v-model="pokemon.nature">
              <option v-for="nature in natures" :key="nature" :value="nature">{{ nature }}</option>
            </select>
          </p>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons">
            <input class="input" type="text" placeholder="Characteristic" v-model="pokemon.characteristic">
            <span class="icon is-small is-left">
              <i class="mdi mdi-account"></i>
            </span>
          </p>
        </div>
        <div class="field">
          <p class="control has-icons-left has-icons">
            <label class="checkbox">
              <input type="checkbox" v-model="pokemon.shiny" />
              Shiny
            </label>
          </p>
        </div>

        <div class="field">
          <p class="control">
            <button class="button is-success" @click="updatePokemon(pokemon)" :class="{ 'is-loading': updatingPokemon.get(pokemon.pokemon) }">
              Update
            </button>
          </p>
        </div>
      </div>
      <div v-else>
        <p>This Pokémon has not been caught yet.</p>
        <div class="field">
          <p class="control">
            <button class="button is-success" @click="catchPokemon(pokemon)" :class="{ 'is-loading': updatingPokemon.get(pokemon.pokemon) }">
              Catch
            </button>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
