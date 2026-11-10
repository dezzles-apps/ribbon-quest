<script setup lang="ts">
import type { PokemonRibbon } from '@/types/ribbons';
import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import type { RibbonFilter } from '@/types/filters';
import Ribbon from './common/Ribbon.vue';
import { RibbonState, type RibbonTheme } from '@/types/enums';
const authStore = useAuthStore();
const api = useApi();

const props = defineProps<{
  title?: string;
  pokemon: string;
  ribbons: PokemonRibbon[];
  filter?: RibbonFilter
}>()

const loadingRibbons = ref(new Map<string, boolean>());

function toggleRibbon(ribbon: PokemonRibbon) {
  if (!authStore.isAuthenticated) {
    return;
  }

  const isLoading = loadingRibbons.value.get(ribbon.ribbonKey) || false;
  if (isLoading) {
    return;
  }

  loadingRibbons.value.set(ribbon.ribbonKey, true);
  let method = ribbon.achieved ? 'DELETE' : 'POST';
  api.apiFetch(API.Ribbons.UpdateRibbon(props.pokemon, ribbon.ribbonKey), {
    method: method,
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(async response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      let newRibbon = await response.json();
      ribbon.achieved = newRibbon.data.achieved;
      ribbon.achievedAt = newRibbon.data.achievedAt;
    })
    .catch(error => {
      console.error('Error toggling ribbon:', error);
    })
    .finally(() => {
      loadingRibbons.value.set(ribbon.ribbonKey, false);
    });
}

function showRibbon(ribbon : PokemonRibbon) {
  const filter = props.filter;
  if (!filter)
    return true;
  if (filter.hideCompleteRibbons) {
    return !ribbon.achieved;
  }
  return true;
}

function ribbonState(ribbon : PokemonRibbon) : RibbonState {
  if (authStore.isAuthenticated) {
    if (loadingRibbons.value.get(ribbon.ribbonKey)) {
      return RibbonState.Loading;
    } else {
      return RibbonState.Clickable;
    }
  }
  return RibbonState.Standard;
}

</script>

<template>
  <div class="box mb-4">
    <slot name="title">
      <h2 class="title is-4">{{ props.title }}</h2>
    </slot>
    <div style="display: inline-flex; flex-wrap: wrap;">
      <Ribbon
        v-for="ribbon in props.ribbons.filter(showRibbon)"
        :theme="ribbon.category as RibbonTheme"
        icon="mdi-seal"
        :bottom-text="ribbon.name"
        :achieved="ribbon.achieved"
        :state="ribbonState(ribbon)"
        :onClick="() => toggleRibbon(ribbon)"
      />
    </div>
  </div>
</template>

<style scoped>

</style>