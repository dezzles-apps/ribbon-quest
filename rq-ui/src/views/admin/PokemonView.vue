<script setup lang="ts">

import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { PokemonStats, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
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

const natures = ref<string[]>([
  '', 'Adamant', 'Bashful', 'Bold', 'Brave', 'Calm', 'Careful', 'Docile',
  'Gentle', 'Hardy', 'Hasty', 'Impish', 'Jolly', 'Lax', 'Lonely', 'Mild',
  'Modest', 'Naive', 'Naughty', 'Quiet', 'Quirky', 'Rash', 'Relaxed', 'Sassy', 'Serious', 'Timid'
]);


onMounted(fetchPokemonStats);
</script>

<template>
  <div class="about" v-for="pokemon in stats" :key="pokemon.pokemon" v-if="stats">
    <div class="box mb-5">
      <h1 class="title is-3">{{ pokemon.pokemon }}</h1>
      <div v-if="pokemon.caughtAt">
        <div class="field">
          <p class="control has-icons-left has-icons-right">
            <input class="input" type="text" placeholder="Nickname" v-model="pokemon.nickname">
            <span class="icon is-small is-left">
              <i class="mdi mdi-account"></i>
            </span>
            <span class="icon is-small is-right">
              <i class="mdi mdi-check"></i>
            </span>
          </p>
        </div>
        <div class="field">
          <p class="control">
            <div class="select">
              <select placeholder="Nature" v-model="pokemon.nature">
                <option v-for="nature in natures" :key="nature" :value="nature">{{ nature }}</option>
              </select>
            </div>
          </p>
        </div>
        {{ viewError }}
        <div class="field">
          <p class="control">
            <button class="button is-success" @click="handleLogin" :class="{ 'is-loading': loading }">
              Login
            </button>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<style>

</style>
