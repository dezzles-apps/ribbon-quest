<script setup lang="ts">
import Box from '@/components/pokedexes/Box.vue';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { Response } from '@/types/responses';
import type { Pokedex, PokedexEntry } from '@/types/pokedex';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import Loading from '@/components/Loading.vue';
import PokedexInfo from '@/components/pokedexes/PokedexInfo.vue';
import { entryVisible } from '@/composables/box';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const pokedex = ref<Pokedex | null>(null);
const entries = ref<PokedexEntry[][]>([]);
const filter = ref('');

function createEntries(input: PokedexEntry[]) {
  let result = [] as PokedexEntry[][]
  let current : PokedexEntry[] = []
  for (let idx = 0; idx < input.length; ++idx) {
    if (idx % 30 == 0) {
      current = []
      result.push(current)
    }
    current.push(input[idx])
  }
  entries.value = result
}

async function fetchPokedex() {
  try {
    const response = await apiFetch(API.Pokedexes.GetPokedex(route.params.pokedex as string));
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<Pokedex>;
    createEntries(data.data.entries)
    pokedex.value = data.data;
  } catch (error) {
    console.error('Error fetching Pokedex:', error);
  } finally {
    loading.value = false;
  }
}

function isVisible(box : PokedexEntry[]) {
  let search = filter.value.trim()
  if (!search) {
    return true
  }
  return box.map((v : PokedexEntry) => v.pokemon).filter(entryVisible(search)).length > 0
}

onMounted(fetchPokedex);
</script>



<template>
  <Loading v-if="loading" />
  <div v-else-if="pokedex">
    <section>
      <PokedexInfo
        :details="pokedex"
        :includeLink="false"
      />
    </section>
    <v-text-field
      label="Filter"
      prepend-inner-icon="mdi-map-marker"
      v-model="filter"
    ></v-text-field>
    <div
      v-for="box in entries"
    >
      <Box
        v-if="isVisible(box)"
        :entries="box"
        :pokedex="pokedex.name"
        :filter="filter"
      />
    </div>

  </div>
  <div v-else>
    <p>Pokedex not found.</p>
  </div>
</template>

