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

const route = useRoute();
const loading = ref(true);
const { apiFetch } = useApi();

const pokedex = ref<Pokedex | null>(null);
const entries = ref<PokedexEntry[][]>([])


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
    <Box
      v-for="box in entries"
      :entries="box"
    />
    <!--v-expansion-panels
      gap="16"
    >
      <v-expansion-panel
        title="Latest Events"
      >
        <v-expansion-panel-text>
          <Events category="RIBBONS" :pokemon="route.params.pokemon as string"/>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>

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
    </GameInfo-->
  </div>
  <div v-else>
    <p>Pokedex not found.</p>
  </div>
</template>

