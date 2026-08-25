<script setup lang="ts">
import Ribbons from '@/components/Ribbons.vue';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { Pokemon, PokemonRibbon, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
import PokemonInfo from '@/components/PokemonInfo.vue';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const pokemon = ref<Pokemon | null>(null);
const ribbonMap = ref<Map<string, PokemonRibbon>>(new Map<string, PokemonRibbon>())
async function fetchPokemon() {
  try {
    const response = await apiFetch(`/api/pokemon/v1/${route.params.pokemon}`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<Pokemon>;
    pokemon.value = data.data;
    pokemon.value.games.sort((a, b) => a.viewOrder - b.viewOrder);
    ribbonMap.value.clear();
    data.data.ribbons.forEach(ribbon => {
      ribbonMap.value.set(ribbon.ribbonKey, ribbon);
    });
  } catch (error) {
    console.error('Error fetching Pokemon:', error);
  } finally {
    loading.value = false;
  }
}

function getRibbonClass(ribbon: PokemonRibbon): string[] {
  const classes: string[] = [];
  if (ribbon.achieved) {
    classes.push('ribbon-achieved');
  } else {
    classes.push('ribbon-not-achieved');
  }
  classes.push(`ribbon-${ribbon.category.toLowerCase()}`);

  return classes;
}


onMounted(fetchPokemon);
</script>



<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="pokemon">
    <section>
      <PokemonInfo
        :pokemon="pokemon.pokemon"
        :nickname="pokemon.nickname"
        :region="pokemon.region"
        :includeLink="false"
      />
    </section>
    <Ribbons title="All Ribbons" :ribbons="pokemon.ribbons" />
    <Ribbons
      v-for="game in pokemon.games"
      :key="game.gameKey"
      :title="game.name"
      :ribbons="game.ribbons.map(ribbonKey => ribbonMap.get(ribbonKey)!).filter(ribbon => ribbon !== undefined)"
    />
  </div>
  <div v-else>
    <p>Pokemon not found.</p>
  </div>
</template>

