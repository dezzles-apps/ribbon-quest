<script setup lang="ts">
import type { PokedexEntry } from '@/types/pokedex';
import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import { isEntryHighlighted } from '@/composables/box';
import Ribbon from '@/components/common/Ribbon.vue';
import { RibbonState, RibbonTheme } from '@/types/enums';
const authStore = useAuthStore();
const api = useApi();

const props = defineProps<{
  pokedex: string
  entries: PokedexEntry[]
  filter: string
}>()



function getPokemonImage(entry: PokedexEntry): string {
  const folder = false ? 'shiny' : 'normal'
  return `/sprites/${folder}/${entry.spriteRef}.png`;
}

const loadingRibbons = ref(new Map<string, boolean>());

function toggleEntry(entry: PokedexEntry) {
  if (!authStore.isAuthenticated) {
    return;
  }
  const isLoading = loadingRibbons.value.get(entry.pokedexNo.text) || false;
  if (isLoading) {
    return;
  }

  loadingRibbons.value.set(entry.pokedexNo.text, true);
  let method = entry.caught ? 'DELETE' : 'POST';
  api.apiFetch(API.Pokedexes.UpdatePokedex(props.pokedex, entry.pokedexNo.text), {
    method: method,
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(async response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      let newData = await response.json();
      entry.caught = newData.data.caught;
    })
    .catch(error => {
      console.error('Error toggling catch:', error);
    })
    .finally(() => {
      loadingRibbons.value.set(entry.pokedexNo.text, false);
    });
}

function ribbonState(entry: PokedexEntry) : RibbonState {
  if (authStore.isAuthenticated) {
    if (loadingRibbons.value.get(entry.pokedexNo.text)) {
      return RibbonState.Loading;
    } else {
      return RibbonState.Clickable;
    }
  }
  return RibbonState.Standard;
}

</script>

<template>
  <div class="box ma-auto text-center" style="width: 100%;">
    <slot name="title">
      <h2 class="title is-4">{{ entries[0].pokedexNo.text }} - {{ entries[entries.length -1].pokedexNo.text }}</h2>
    </slot>
    <div style="display: inline-flex; flex-wrap: wrap; max-width: 780px;" class="ma-auto">
      <Ribbon
        v-for="entry in props.entries"
        :key="entry.pokedexNo.text"
        :theme="RibbonTheme.Stats"
        :top-text="entry.pokedexNo.text"
        :bottom-text="entry.pokemon"
        :image="getPokemonImage(entry)"
        :state="ribbonState(entry)"
        :achieved="entry.caught"
        :highlighted="isEntryHighlighted(props.filter, entry.pokemon)"
        :onClick="() => toggleEntry(entry)"
      />
    </div>
  </div>
</template>
