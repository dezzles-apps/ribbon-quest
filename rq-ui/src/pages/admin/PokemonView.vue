<script setup lang="ts">

import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { PokemonStats, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
const route = useRoute();
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
const natures = ref<string[]>([
  '', 'Adamant', 'Bashful', 'Bold', 'Brave', 'Calm', 'Careful', 'Docile',
  'Gentle', 'Hardy', 'Hasty', 'Impish', 'Jolly', 'Lax', 'Lonely', 'Mild',
  'Modest', 'Naive', 'Naughty', 'Quiet', 'Quirky', 'Rash', 'Relaxed', 'Sassy', 'Serious', 'Timid'
]);

const updatingPokemon = ref<Map<String, boolean>>(new Map());

function catchPokemon(pokemon: PokemonStats) {
  updatingPokemon.value.set(pokemon.pokemon, true);
  apiFetch(API.Ribbons.CatchPokemon(pokemon.pokemon), {
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
  apiFetch(API.Ribbons.UpdatePokemon(pokemon.pokemon), {
    method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        shiny: pokemon.shiny,
        nature: pokemon.nature,
        characteristic: pokemon.characteristic,
        nickname: pokemon.nickname,
        notes: pokemon.notes
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
    pokemon.notes = data.notes?? '';
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
      <v-card
        elevation="2"
        rounded
        style="padding: 10px;"
        :title="pokemon.pokemon"
      >
        <div v-if="pokemon.caughtAt">
          <v-text-field
            label="Nickname"
            v-model="pokemon.nickname"
            type="text"
          ></v-text-field>
          <v-select
            label="Nature"
            v-model="pokemon.nature"
            :items="natures"
          ></v-select>
          <v-text-field
            label="Characteristic"
            v-model="pokemon.characteristic"
            type="text"
          ></v-text-field>
          <v-textarea
            label="Notes"
            v-model="pokemon.notes"
            type="text"
          ></v-textarea>
          <v-checkbox
            label="Shiny"
            v-model="pokemon.shiny"
          ></v-checkbox>
        </div>
        <div v-else>
          <p>This Pokémon has not been caught yet.</p>
        </div>
        <v-card-actions>
          <v-btn
            v-if="pokemon.caughtAt"
            color="orange"
            variant="flat"
            :loading="updatingPokemon.get(pokemon.pokemon)"
            @click="updatePokemon(pokemon)"
          >
            Update
          </v-btn>
          <v-btn
            v-if="!pokemon.caughtAt"
            color="orange"
            variant="flat"
            :loading="updatingPokemon.get(pokemon.pokemon)"
            @click="catchPokemon(pokemon)"
          >
            Catch
          </v-btn>
        </v-card-actions>
      </v-card>
    </div>
  </div>
</template>

<style>

</style>
