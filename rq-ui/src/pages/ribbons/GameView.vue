<script setup lang="ts">
import Ribbons from '@/components/Ribbons.vue';
import { Ribbons as apiRibbons } from '@/composables/endpoints';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import type { Response } from '@/types/responses';
import type { Game, GamePokemon } from '@/types/ribbons';
import type { RibbonFilter } from '@/types/filters';
import Loading from '@/components/Loading.vue';
import PokemonInfo from '@/components/PokemonInfo.vue';
import GameInfo from '@/components/GameInfo.vue';
import { useApi } from '@/composables/useApi';
const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const game = ref<Game | null>(null);
const filter = ref({} as RibbonFilter)
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

function showPokemon(pokemon : GamePokemon) {
  if (!filter.value.hideCompletePokemon)
    return true;
  return pokemon.ribbons.filter(v => !v.achieved).length > 0;
}
onMounted(fetchGame);
</script>

<template>
  <Loading v-if="loading"/>
  <div v-else-if="game">
    <section>
      <GameInfo
        :gameKey="game.gameKey"
        :name="game.name"
        :includeLink="false"
      />
      <v-container>
        <v-row gap="0">
          <v-col>
            <v-checkbox
              class="pa-2"
              v-model="filter.hideCompletePokemon"
              label="Hide Complete Pokemon"
            ></v-checkbox>
          </v-col>
          <v-col>
            <v-checkbox class="pa-2"
              v-model="filter.hideCompleteRibbons"
              label="Hide Complete Ribbons"
            ></v-checkbox>
          </v-col>

          <v-responsive width="100%"></v-responsive>
        </v-row>
      </v-container>
    </section>
    <PokemonInfo
      v-for="pokemon in game.pokemon.filter(showPokemon)"
      :details="pokemon"
      :includeDescription="false"
      :includeLink="true"
    >
      <template v-slot:content>
        <Ribbons
          :key="pokemon.pokemon"
          :ribbons="pokemon.ribbons"
          :pokemon="pokemon.pokemon"
          :filter="filter"
        />
      </template>
    </PokemonInfo>


  </div>
  <div v-else>
    <p>Pokemon not found.</p>
  </div>
</template>

