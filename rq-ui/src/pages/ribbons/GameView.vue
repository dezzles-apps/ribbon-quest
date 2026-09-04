<script setup lang="ts">
import Ribbons from '@/components/Ribbons.vue';
import { Ribbons as apiRibbons } from '@/composables/endpoints';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { Response, Game } from '@/types/api';
import PokemonInfo from '@/components/PokemonInfo.vue';
import GameInfo from '@/components/GameInfo.vue';
import { useApi } from '@/composables/useApi';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const game = ref<Game | null>(null);
async function fetchGame() {
  try {
    const response = await apiFetch(apiRibbons.GetGame(route.params.game as string));
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
    <PokemonInfo
      v-for="pokemon in game.pokemon"
      :details="pokemon"
      :includeDescription="false"
      :includeLink="true"
    >
      <template v-slot:content>
        <Ribbons
          :key="pokemon.pokemon"
          :ribbons="pokemon.ribbons"
          :pokemon="pokemon.pokemon"
        />
      </template>
    </PokemonInfo>


  </div>
  <div v-else>
    <p>Pokemon not found.</p>
  </div>
</template>

