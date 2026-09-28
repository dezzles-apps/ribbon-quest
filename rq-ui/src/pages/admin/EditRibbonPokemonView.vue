<script setup lang="ts">

import { onMounted, ref } from 'vue';
import type { Pokemon, GameWithStats } from '@/types/ribbons';
import type { Response } from '@/types/responses';
import { useRoute } from 'vue-router';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import Loading from '@/components/Loading.vue';
const loading = ref(true);
const { apiFetch } = useApi();
const route = useRoute();
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

const pokemon = ref<Pokemon | null>(null);
const games = ref<GameWithStats[] | null>(null);
async function fetchPokemon() {
  return apiFetch(API.Ribbons.GetPokemon(route.params.pokemonId as string))
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP Error :(`);
      }
      return response.json() as Promise<Response<Pokemon>>;
    }).then(data => {
      pokemon.value = data.data;
    })
}
async function fetchGames() {
  return apiFetch(API.Ribbons.GetAllGames)
    .then(response => {
      if (!response.ok) {
        throw new Error(`Http Error :(`)
      }
      return response.json() as Promise<Response<GameWithStats[]>>;
    }).then(data => {
      games.value = data.data
    })
}

const natures = ref<string[]>([
  '', 'Adamant', 'Bashful', 'Bold', 'Brave', 'Calm', 'Careful', 'Docile',
  'Gentle', 'Hardy', 'Hasty', 'Impish', 'Jolly', 'Lax', 'Lonely', 'Mild',
  'Modest', 'Naive', 'Naughty', 'Quiet', 'Quirky', 'Rash', 'Relaxed', 'Sassy', 'Serious', 'Timid'
]);

const updatingPokemon = ref(false);
const showSnackbar = ref<SnackbarDetails>(new SnackbarDetails('', false, false));

function removeGame(gameId: String) {
  console.log(`removing game ${gameId}`)
}

function catchPokemon(pokemon: Pokemon) {
  updatingPokemon.value = true
  apiFetch(API.Ribbons.CatchPokemon(pokemon.pokemon), {
    method: 'POST'
  }).then(async response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const { data } = await response.json();
    syncDetails(pokemon, data)
    getSnackInfo(pokemon).update(`${pokemon.pokemon} caught`, false);
  })
  .catch(error => {
    getSnackInfo(pokemon).update(`Error catching ${pokemon.pokemon}`, true);
    console.error('Error catching Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value = false
  });
}

function getSnackInfo(details: Pokemon) : SnackbarDetails {
  const pokemon = details.pokemon;
  const snack = showSnackbar.value
  if (!snack) {
    return new SnackbarDetails("", false, false);
  }
  return snack;
}

function syncDetails(pokemon: Pokemon, data: Pokemon) {
  pokemon.details.caughtAt = data.details.caughtAt;
  pokemon.details.nickname = data.details.nickname ?? '';
  pokemon.details.nature = data.details.nature ?? '';
  pokemon.details.shiny = data.details.shiny?? false;
  pokemon.details.characteristic = data.details.characteristic?? '';
  pokemon.details.notes = data.details.notes?? '';
  pokemon.games = data.games;
}

function updateGame(gameKey: string) {
  if (!pokemon.value) {
    return;
  }
  const base = pokemon.value
  updatingPokemon.value = true
  apiFetch(API.Ribbons.UpdatePokemonGame(pokemon.value.pokemon, gameKey), {
    method: hasGame(gameKey) ? 'DELETE': 'PUT'
  }).then(async response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const { data } = await response.json() as Response<Pokemon>
    syncDetails(base, data)
    getSnackInfo(base).update(`Updated ${base.details.nickname}`, false);
  })
  .catch(error => {
    getSnackInfo(base).update(`Error updating ${base.details.nickname}`, true);
    console.error('Error updating Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value = false
  });
}

function updatePokemon(pokemon: Pokemon) {
  updatingPokemon.value = true
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
    const { data } = await response.json() as Response<Pokemon>
    syncDetails(pokemon, data)
    getSnackInfo(pokemon).update(`Updated ${pokemon.details.nickname}`, false);
  })
  .catch(error => {
    getSnackInfo(pokemon).update(`Error updating ${pokemon.details.nickname}`, true);
    console.error('Error updating Pokemon:', error);
  })
  .finally(() => {
    updatingPokemon.value = false
  });
}

function hasGame(gameKey: String) : boolean {
  if (pokemon.value == null) return false
  return pokemon.value.games.filter(v => v.gameKey === gameKey).length > 0
}

async function fetchData() {
  await Promise.all([
    fetchPokemon(), fetchGames()
  ]).finally(() => {
    loading.value = false
  })
}

onMounted(fetchData);
</script>

<template>
  Edit
  <Loading v-if="loading"/>
  <v-sheet v-if="pokemon" class="pa-5">
    <v-icon
      class="mr-2"
      icon="mdi-pokeball"
      :color="pokemon.details.caughtAt ? 'white': 'grey'"
    ></v-icon> {{  pokemon.pokemon }}
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
      <h3 class="subtitle">
        Games
      </h3>
      <table>
        <tr
          v-for="game in games"
        >
          <td>{{ game.name }}</td>
          <td>
            <v-btn
              variant="elevated"
              :color="hasGame(game.gameKey) ? 'error' : 'primary'"
              style="width: 100%;"
              :loading="updatingPokemon"
              @click="updateGame(game.gameKey)"
            >
              {{ hasGame(game.gameKey) ? 'Remove' : 'Add' }}
            </v-btn>
          </td>
        </tr>
      </table>
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
        :loading="updatingPokemon"
        @click="updatePokemon(pokemon)"
      >
        Update
      </v-btn>
      <v-btn
        v-if="!pokemon.details.caughtAt"
        color="orange"
        variant="flat"
        :loading="updatingPokemon"
        @click="catchPokemon(pokemon)"
      >
        Catch
      </v-btn>
    </v-card-actions>
  </v-sheet>
</template>

<style>

</style>
