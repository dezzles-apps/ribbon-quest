<script setup lang="ts">
import type { PokedexEntry } from '@/types/pokedex';
import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import { isEntryVisible } from '@/composables/box';

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

function getRibbonClass(entry: PokedexEntry): string[] {
  const classes: string[] = [];
  let mod = ''
  if (!isEntryVisible(props.filter, entry.pokemon)) {
    mod = '-low-opacity'
  }
  if (entry.caught) {
    classes.push('ribbon-achieved' + mod);
  } else {
    classes.push('ribbon-not-achieved' + mod);
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
  const isLoading = loadingRibbons.value.get(entry.pokedexNo) || false;
  if (isLoading) {
    return;
  }

  loadingRibbons.value.set(entry.pokedexNo, true);
  let method = entry.caught ? 'DELETE' : 'POST';
  api.apiFetch(API.Pokedexes.UpdatePokedex(props.pokedex, entry.pokedexNo), {
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
      loadingRibbons.value.set(entry.pokedexNo, false);
    });
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
        :style="{ backgroundImage: `url('${getPokemonImage(entry)}')` }"
        :key="entry.pokedexNo"
        @click="toggleEntry(entry)"
      >
        <div class="ribbon-name">
          {{ entry.pokedexNo }}
        </div>

        <div class="ribbon-pokemon-name">
          {{ entry.pokemon }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ribbon {
  margin: 10px;
  position: relative;
  align-items: center;
  text-align: center;
  width: 110px;
  height: 100px;
  outline-width: 5px;
  outline-style: solid;
  border-radius: 10px;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 50% auto;
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

.ribbon-pokemon-name {
  position: absolute;
  bottom: 4px;
  left: 0;
  font-weight: bold;
  margin-right: 10px;
  width: 100%;
  color: darkslategray;
}

.ribbon-achieved-low-opacity {
  opacity: 0.5;
}

.ribbon-not-achieved {
  opacity: 0.5;
}

.ribbon-not-achieved-low-opacity {
  opacity: 0.20;
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
.ribbon-low-opacity {
  opacity: 10%;
}
</style>