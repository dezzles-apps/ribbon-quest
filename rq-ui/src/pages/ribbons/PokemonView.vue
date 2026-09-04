<script setup lang="ts">
import Ribbons from '@/components/Ribbons.vue';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { Pokemon, PokemonRibbon, Response } from '@/types/api';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import PokemonInfo from '@/components/PokemonInfo.vue';
import GameInfo from '@/components/GameInfo.vue';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const pokemon = ref<Pokemon | null>(null);
const ribbonMap = ref<Map<string, PokemonRibbon>>(new Map<string, PokemonRibbon>())
async function fetchPokemon() {
  try {
    const response = await apiFetch(API.Ribbons.GetPokemon(route.params.pokemon as string));
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

onMounted(fetchPokemon);
</script>



<template>
  <div v-if="loading">
    <progress class="progress is-large is-info" max="100">60%</progress>
  </div>
  <div v-else-if="pokemon">
    <section>
      <PokemonInfo
        :details="pokemon"
        :includeLink="false"
      />
    </section>
    <Ribbons
      title="All Ribbons"
      :ribbons="pokemon.ribbons"
      :pokemon="pokemon.pokemon"
    />
    <GameInfo
      v-for="game in pokemon.games"
      :gameKey="game.gameKey"
      :name="game.name"
      :includeLink="true"
    >
      <template v-slot:content>
        <Ribbons
          :key="game.gameKey"
          :ribbons="game.ribbons.map(ribbonKey => ribbonMap.get(ribbonKey)!).filter(ribbon => ribbon !== undefined)"
          :pokemon="pokemon.pokemon"
        />
      </template>
    </GameInfo>
  </div>
  <div v-else>
    <p>Pokemon not found.</p>
  </div>
</template>

