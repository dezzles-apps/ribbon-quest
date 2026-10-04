<script setup lang="ts">
import type { PokedexEntry } from '@/types/pokedex';
import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';

const authStore = useAuthStore();
const api = useApi();

const props = defineProps<{
  entries: PokedexEntry[]
}>()



function getPokemonImage(entry: PokedexEntry): string {
  const folder = false ? 'shiny' : 'normal'
  return `/sprites/${folder}/${entry.spriteRef}.png`;
}

const loadingRibbons = ref(new Map<string, boolean>());

function getRibbonClass(entry: PokedexEntry): string[] {
  const classes: string[] = [];
  if (entry.caught) {
    classes.push('ribbon-achieved');
  } else {
    classes.push('ribbon-not-achieved');
  }
  if (authStore.isAuthenticated) {
    if (loadingRibbons.value.get(entry.pokedexNo)) {
      classes.push('ribbon-loading');
    } else {
      classes.push('ribbon-clickable');
    }
  }
  classes.push(`ribbon-stats`);

  return classes;
}

function toggleEntry(entry: PokedexEntry) {
  if (!authStore.isAuthenticated) {
    return;
  }
  /*
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
    */
}


</script>

<template>
  <div class="box ma-auto text-center" style="width: 100%;">
    <slot name="title">
      <h2 class="title is-4">{{ entries[0].pokedexNo }} - {{ entries[entries.length -1].pokedexNo }}</h2>
    </slot>
    <div style="display: inline-flex; flex-wrap: wrap; max-width: 780px;" class="ma-auto">
      <div
        class="ribbon"
        :class="getRibbonClass(entry)"
        v-for="entry in props.entries"
        :key="entry.pokedexNo"
        @click="toggleEntry(entry)"
      >
        <div class="ribbon-name">
          {{ entry.pokedexNo }}
        </div>
        <img
          :src="getPokemonImage(entry)"
          :alt="entry.pokemon"
          class="pokemon-image"
        />
        <div class="ribbon-name">
          {{ entry.pokemon }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ribbon {
  margin: 10px;
  align-items: center;
  text-align: center;
  width: 110px;
  height: 100px;
  outline-width: 5px;
  outline-style: solid;
  border-radius: 10px;
}

.ribbon-loading {
  cursor: wait;
}

.ribbon-clickable {
  cursor: pointer;
}

.ribbon-name {
  font-weight: bold;
  margin-right: 10px;
  width: 100%;
  color: darkslategray;
}

.ribbon-not-achieved {
  opacity: 0.5;
}

.ribbon-julie {
  background-color: #DCB0F2;
  outline-color: #a35ec9;
}
.ribbon-champion {
  background-color: #66C5CC;
  outline-color: #358d95;
}
.ribbon-battle {
  background-color: #F89C74;
  outline-color: #754834;
}
.ribbon-contest {
  background-color: #FE88B1;
  outline-color: #8e3d5a;
}
.ribbon-stats {
  background-color: #9EB9F3;
  outline-color: #62749b;
}
.ribbon-shopping {
  background-color: #F6CF71;
  outline-color: #d5a32c;
}
</style>