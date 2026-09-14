<script setup lang="ts">

import { onMounted, ref } from 'vue';
import type { PokemonStats, PokemonDetails } from '@/types/ribbons';
import type { Response } from '@/types/responses';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import Loading from '@/components/Loading.vue';
const loading = ref(true);
const { apiFetch } = useApi();

class SnackbarDetails {
  message: string
  show: boolean
  isError: boolean

  constructor(message: string, show: boolean, isError: boolean) {
    this.message = message;
    this.show = show;
    this.isError = isError;
  }

  update(message : string, isError : boolean) {
    this.message = message;
    this.isError = isError;
    this.show = true;
  }
}

const stats = ref<PokemonStats[] | null>(null);
async function fetchPokemonStats() {
  try {
    const response = await apiFetch(API.Ribbons.GetAllPokemon);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<PokemonStats[]>;
    stats.value = data.data;
    data.data.forEach(pokemon => {
      showSnackbar.value.set(pokemon.pokemon, new SnackbarDetails('', false, false));
    })
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
const showSnackbar = ref<Map<string, SnackbarDetails>>(new Map());

function catchPokemon(pokemon: PokemonStats) {
  updatingPokemon.value.set(pokemon.pokemon, true);
  apiFetch(API.Ribbons.CatchPokemon(pokemon.pokemon), {
    method: 'POST'
  }).then(async response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const { data } = await response.json();
    pokemon.details.caughtAt = data.caughtAt;
    pokemon.details.nickname = data.nickname ?? '';
    pokemon.details.nature = data.nature ?? '';
    getSnackInfo(pokemon).update(`${pokemon.pokemon} caught`, false);
  })
  .catch(error => {
    getSnackInfo(pokemon).update(`Error catching ${pokemon.pokemon}`, true);
    console.error('Error catching Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value.set(pokemon.pokemon, false);
  });
}

function getSnackInfo(details: PokemonStats) : SnackbarDetails {
  const pokemon = details.pokemon;
  const snack = showSnackbar.value.get(pokemon)
  if (!snack) {
    return new SnackbarDetails("", false, false);
  }
  return snack;
}

function updatePokemon(pokemon: PokemonStats) {
  updatingPokemon.value.set(pokemon.pokemon, true);
  apiFetch(API.Ribbons.UpdatePokemon(pokemon.pokemon), {
    method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        shiny: pokemon.details.shiny,
        nature: pokemon.details.nature,
        characteristic: pokemon.details.characteristic,
        nickname: pokemon.details.nickname,
        notes: pokemon.details.notes
      })
  }).then(async response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const { data } = await response.json() as Response<PokemonDetails>
    pokemon.details.caughtAt = data.details.caughtAt;
    pokemon.details.nickname = data.details.nickname ?? '';
    pokemon.details.nature = data.details.nature ?? '';
    pokemon.details.shiny = data.details.shiny?? false;
    pokemon.details.characteristic = data.details.characteristic?? '';
    pokemon.details.notes = data.details.notes?? '';
    getSnackInfo(pokemon).update(`Updated ${pokemon.details.nickname}`, false);
  })
  .catch(error => {
    getSnackInfo(pokemon).update(`Error updating ${pokemon.details.nickname}`, true);
    console.error('Error updating Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value.set(pokemon.pokemon, false);
  });
}

onMounted(fetchPokemonStats);
</script>

<template>
  <Loading v-if="loading"/>
  <v-expansion-panels>
    <v-expansion-panel v-for="(pokemon, idx) in stats" :key="pokemon.pokemon" :title="pokemon.pokemon" v-if="stats">
      <v-expansion-panel-text>
        <div v-if="pokemon.details.caughtAt">
          <v-text-field
            label="Nickname"
            v-model="pokemon.details.nickname"
            type="text"
          ></v-text-field>
          <v-select
            label="Nature"
            v-model="pokemon.details.nature"
            :items="natures"
          ></v-select>
          <v-text-field
            label="Characteristic"
            v-model="pokemon.details.characteristic"
            type="text"
          ></v-text-field>
          <v-textarea
            label="Notes"
            v-model="pokemon.details.notes"
            type="text"
          ></v-textarea>
          <v-checkbox
            label="Shiny"
            v-model="pokemon.details.shiny"
          ></v-checkbox>
        </div>
        <div v-else>
          <p>This Pokémon has not been caught yet.</p>
        </div>
        <v-snackbar
          v-model="getSnackInfo(pokemon).show"
          :timeout="2000"
          :color="getSnackInfo(pokemon).isError ? 'warning': 'success'"
        >
          {{ getSnackInfo(pokemon).message }}
        </v-snackbar>
        <v-card-actions>
          <v-btn
            v-if="pokemon.details.caughtAt"
            color="orange"
            variant="flat"
            :loading="updatingPokemon.get(pokemon.pokemon)"
            @click="updatePokemon(pokemon)"
          >
            Update
          </v-btn>
          <v-btn
            v-if="!pokemon.details.caughtAt"
            color="orange"
            variant="flat"
            :loading="updatingPokemon.get(pokemon.pokemon)"
            @click="catchPokemon(pokemon)"
          >
            Catch
          </v-btn>
        </v-card-actions>
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<style>

</style>
