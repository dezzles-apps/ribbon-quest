<script setup lang="ts">

import { ref, onMounted, shallowRef, computed } from 'vue';
import { useRouter } from 'vue-router'
import { useApi } from '@/composables/useApi'
import API from '@/composables/endpoints'
import type { PokemonData, CommonData } from '@/types/data'
import type { CreateShinyRequest} from '@/types/requests'
import type { Response, Error } from '@/types/responses'
const router = useRouter()
const request = ref({level: 1} as CreateShinyRequest)
const { apiFetch } = useApi()
const menu = shallowRef(false)
const loading = ref(true)
const loadError = ref('')
const lookupsLoaded = ref(false)
const commonData = ref({} as CommonData)
const allPokemon = ref([] as PokemonData[])
const selectedPokemon = ref('' as string)
const errors = ref([] as Error[])
const saving = ref(false)
const submitError = ref('')
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
function getCommonData() {
  return apiFetch(API.Data.GetCommonData)
    .then(async response => {
      if (!response.ok) {
        throw new Error("Something went wrong")
      }
      const result = await response.json() as Response<CommonData>
      commonData.value = result.data
    })
}

function load() {
  loading.value = true
  loadError.value = ''
  lookupsLoaded.value = false
  return Promise.all([
    getCommonData(),
    getPokemonData()
  ]).then(() => {
    lookupsLoaded.value = true
  }).catch(() => {
    loadError.value = 'Unable to load the data required to add a shiny Pokémon.'
  }).finally(() => {
    loading.value = false
  })
}

const currentPokemon = computed(getCurrentPokemon)
function getCurrentPokemon() {
  if (!selectedPokemon) return null
  return allPokemon.value.filter(v => v.species === selectedPokemon.value)[0]

}
function updateForm() {
  request.value.form = currentPokemon.value?.forms[0].formId || ''
}

async function savePokemon() {
  if (loading.value || !lookupsLoaded.value || saving.value) {
    return
  }
  saving.value = true
  errors.value = []
  submitError.value = ''
  try {
    const resp = await apiFetch(API.Shinies.GetAllShinies, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(request.value)
    })
    const body = await resp.json() as Response<PokemonData>
    if (body.errors?.length) {
      errors.value = body.errors
      return
    }
    if (!resp.ok) {
      submitError.value = body.error || 'Unable to add the shiny Pokémon.'
      return
    }
    router.push(`/shinies/catches`)
  } catch {
    submitError.value = 'Unable to add the shiny Pokémon.'
  } finally {
    saving.value = false
  }
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

function getPokemonErrorMessages() {
  return [
    getErrorMessage('pokedexId'),
    selectedPokemon.value === '' ? getErrorMessage('form') : '',
  ].filter(Boolean)
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
    <v-alert v-if="submitError" class="ma-4" type="error">
      {{ submitError }}
    </v-alert>

    <v-alert v-if="loadError" class="ma-4" type="error">
      {{ loadError }}
      <v-btn class="ml-2" @click="load">Retry</v-btn>
    </v-alert>

    <v-form>
      <v-container v-if="!loading">
        <v-select
          v-model="selectedPokemon"
          label="Pokemon"
          v-model:menu="menu"
          :items="allPokemon"
          item-title="species"
          v-on:update:model-value="updateForm"
          :error-messages="getPokemonErrorMessages()"
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
          label="Form"
          v-model="request.form"
          :items="currentPokemon.forms"
          item-title="form"
          :error-messages="getErrorMessage('form')"
          item-value="formId"
        >
        </v-select>
        <v-select
          v-if="commonData.genders"
          v-model="request.gender"
          label="Gender"
          :items="commonData.genders"
          item-title="value"
          :error-messages="getErrorMessage('gender')"
          item-value="key"
        >
        </v-select>
        <v-number-input
          :reverse="false"
          controlVariant="default"
          v-model="request.level"
          label="Level"
          :min="1"
          :max="100"
          :hideInput="false"
          :inset="false"
        ></v-number-input>
        <v-select
          label="Nature"
          v-if="commonData.natures"
          v-model="request.nature"
          :items="commonData.natures"
          item-title="value"
          :error-messages="getErrorMessage('nature')"
          item-value="key"
        >
        </v-select>
        <v-select
          label="Characteristic"
          v-if="commonData.characteristics"
          v-model="request.characteristic"
          :items="commonData.characteristics"
          item-title="value"
          :error-messages="getErrorMessage('characteristic')"
          item-value="key"
        >
        </v-select>
        <v-select
          label="Game"
          v-if="commonData.games"
          v-model="request.game"
          :items="commonData.games"
          item-title="value"
          :error-messages="getErrorMessage('game')"
          item-value="key"
        >
        </v-select>
        <v-select
          label="Method"
          v-if="commonData.methods"
          v-model="request.method"
          :items="commonData.methods"
          item-title="value"
          :error-messages="getErrorMessage('method')"
          item-value="key"
        >
        </v-select>
        <v-textarea
          label="Notes"
          v-model="request.notes"
          :error-messages="getErrorMessage('notes')"
        ></v-textarea>
      </v-container>
      <v-container>
        <v-btn
          :disabled="loading || !lookupsLoaded || saving"
          color="primary"
          :loading="saving"
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