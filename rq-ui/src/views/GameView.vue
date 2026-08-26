<script setup lang="ts">
import Ribbons from '@/components/Ribbons.vue';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { Response, Game, GamePokemon, PokemonRibbon } from '@/types/api';
import PokemonInfo from '@/components/PokemonInfo.vue';
import GameInfo from '@/components/GameInfo.vue';
import { useApi } from '@/composables/useApi';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const game = ref<Game | null>(null);
async function fetchGame() {
  try {
    const response = await apiFetch(`/api/games/v1/${route.params.game}`);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json() as Response<Game>;
    game.value = data.data;
  } catch (error) {
    console.error('Error fetching Game:', error);
  } finally {
    loading.value = false;
  }
}

onMounted(fetchGame);
</script>

<template>
  <div v-if="loading">
    <progress class="progress is-large is-info" max="100">60%</progress>
  </div>
  <div v-else-if="game">
    <section>
      <GameInfo
        :gameKey="game.gameKey"
        :name="game.name"
        :includeLink="false"
      />
    </section>
    <Ribbons
      v-for="pokemon in game.pokemon"
      :key="pokemon.pokemon"
      :ribbons="pokemon.ribbons"
    >
      <template v-slot:title>
        <PokemonInfo
          :pokemon="pokemon.pokemon"
          :nickname="pokemon.nickname"
          :region="pokemon.region"
          :includeDescription="false"
          :includeLink="true"
        />
      </template>
    </Ribbons>
  </div>
  <div v-else>
    <p>Pokemon not found.</p>
  </div>
</template>

