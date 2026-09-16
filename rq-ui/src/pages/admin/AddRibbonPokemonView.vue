<script setup lang="ts">

import { ref, onMounted, shallowRef, computed } from 'vue';
import { useRouter } from 'vue-router'
import { useApi } from '@/composables/useApi'
import API from '@/composables/endpoints'
import type { PokemonData } from '@/types/data'
import type { Game, NewRibbonPokemonRequest } from '@/types/ribbons'
import type { Response, Error } from '@/types/responses'
const router = useRouter()
const { apiFetch } = useApi()
const menu = shallowRef(false)
const loading = ref(true)
const pokemonId = ref('')
const selectedForm = ref('')
const allGames = ref([] as Game[])
const allPokemon = ref([] as PokemonData[])
const selectedPokemon = ref('' as string)
const errors = ref([] as Error[])
const selectedGames = ref([] as string[])
const saving = ref(false)
function getPokemonData() {
  return apiFetch(API.Data.GetPokemonData)
    .then(async response => {
      if (!response.ok) {
        throw new Error("Something went wrong")
      }
      const result = await response.json() as Response<PokemonData[]>
      allPokemon.value = result.data
    })
}
function getGames() {
  return apiFetch(API.Ribbons.GetAllGames)
    .then(async response => {
      if (!response.ok) {
        throw new Error("Something went wrong")
      }
      const result = await response.json() as Response<Game[]>
      allGames.value = result.data
    })
}

function load() {
  return Promise.all([
    getGames(),
    getPokemonData()
  ]).then(v => {
    loading.value = false
  })
}

const currentPokemon = computed(getCurrentPokemon)
function getCurrentPokemon() {
  if (!selectedPokemon) return null
  return allPokemon.value.filter(v => v.species === selectedPokemon.value)[0]

}
function updateForm() {
  selectedForm.value = currentPokemon.value?.forms[0].form ?? ''
}

function getRequestBody() : NewRibbonPokemonRequest {
  let result = {
    pokemon: pokemonId.value,
  } as NewRibbonPokemonRequest
  let current = getCurrentPokemon()
  if (current) {
    result.pokedexId = current.pokedexNo
  }
  result.form = selectedForm.value
  result.games = selectedGames.value
  return result
}

function savePokemon() {
  saving.value = true
  errors.value = []
  const request = getRequestBody()
  return apiFetch(API.Ribbons.GetAllPokemon, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(request)
  }).then(async resp => {
    saving.value = false
    const body = await resp.json() as Response<PokemonData>
    if (body.errors) {
      errors.value = body.errors
      return
    }
    router.push(`/ribbons/pokemon/${pokemonId.value}`)
    console.log(body)
  })
}

function getErrorMessage(field : string) : string {
  if (!errors.value) {
    return ''
  }
  let fieldErrors = errors.value.filter(e => e.field === field)
  if (fieldErrors.length === 0) {
    return ''
  }
  return fieldErrors[0].message
}

onMounted(load)

</script>

<template>
  <v-card
    class="mx-auto"
    color="surface-variant"
    title="New Ribbon Pokemon"
    variant="tonal"
  >
    {{  errors }}
    <v-form>
      <v-container v-if="!loading">
        <v-text-field
          label="PokemonId"
          v-model="pokemonId"
          :error-messages="getErrorMessage('pokemon')"
        ></v-text-field>
        <v-select
          v-model="selectedPokemon"
          v-model:menu="menu"
          :items="allPokemon"
          item-title="species"
          v-on:update:model-value="updateForm"
          :error-messages="getErrorMessage('pokedexId')"
        >
          <template v-slot:menu-header="{ search, filteredItems }">
            <div class="pa-2 border-b">
              <v-text-field
                v-model="search.value"
                :error="!!search.value && !filteredItems.length"
                density="compact"
                placeholder="Search..."
                prepend-inner-icon="mdi-magnify"
                variant="outlined"
                clearable
                hide-details
              ></v-text-field>
            </div>
          </template>
        </v-select>
        <v-select
          v-if="currentPokemon"
          v-model="selectedForm"
          :items="currentPokemon.forms"
          item-title="form"
          :error-messages="getErrorMessage('form')"
        >
        </v-select>
      </v-container>
      <v-container fluid>
        Games
        <v-row>
          <v-col
            v-for="game in allGames"
            cols="12"
            md="3"
            sm="3"
          >
            <v-checkbox
              v-model="selectedGames"
              :label="game.name"
              :value="game.gameKey"
              density="compact"
            ></v-checkbox>
          </v-col>
        </v-row>
            <div class="error">
              {{ getErrorMessage('games') }}
            </div>
        <v-btn
          :loading="saving"
          color="primary"
          :to="{ name: 'admin-add-ribbon-pokemon' }"
          @click="savePokemon"
        >
          Add New Pokemon
        </v-btn>
      </v-container>
    </v-form>

  </v-card>
</template>


<style lang="css" scoped>
.error {
  font-size: 16px;
  color: red;
}
</style>