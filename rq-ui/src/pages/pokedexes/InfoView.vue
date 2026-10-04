<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { Pokedex } from '@/types/pokedex';
import type { Response } from '@/types/responses';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';
import Events from '@/components/Events.vue'
import Loading from '@/components/Loading.vue';
const loading = ref(true);
const caught = ref(0);
const total = ref(1);
const { apiFetch } = useApi();

function loadStats() {
  return apiFetch(API.Pokedexes.GetAllPokedexes).then(async resp => {
    if (!resp.ok) {
      throw new Error("An error occurred");
    }
    loading.value = false;
    const responseBody = await resp.json() as Response<Pokedex[]>;
    caught.value = responseBody.data.map(v => v.stats.caught).reduce((p, n) => p + n)
    total.value = responseBody.data.map(v => v.stats.total).reduce((p, n) => p + n)
  })
}

onMounted(loadStats);

</script>

<template>
  <v-sheet
    class="d-flex flex-wrap mx-auto px-10"
    elevation="2"
    rounded
  >
    <div class="box mb-4">
      <h1 class="title is-4 text-center">A Dex Of My Own</h1>
      <div class="d-flex align-center ma-auto justify-space-between" style="max-width: 300px;">
        <div class="mt-n2">
          <v-card-title>Progress</v-card-title>
        </div>
        <Loading v-if="loading" />
        <div
          class="text-center"
          v-else
        >
          <v-progress-circular
            :model-value="100 * caught / total"
            :size="100"
            :width="12"
            bg-color="surface-light"
            class="ma-3"
            color="orange-accent-2"
            reveal
            rounded
          >
            <v-avatar color="surface-light" size="70">{{caught}} / {{total}}</v-avatar>
          </v-progress-circular>
        </div>
      </div>
      <div>
        Every time a new Pokemon game comes out, I speed through completing my Living Pokedex for the game
        as well as updating (where necessary) my National Pokedex. This little area is for tracking
        that progress.
      </div>
      <br>
      <div>
        Any game not currently listed here from Generation 8 onwards, you can pretty safely assume that the
        Pokedex for that game already exists. However, while I <b>do</b> have a full Living National Dex at
        moment, it has one thing that upsets me: A lot of the Pokemon came from Wonder Trades in generations
        six and seven.
      </div>
      <br>
      <div>
        So here we go... Going through and updating my National Living Dex to have all of my Original Trainer
        ID and being named correctly. If an entry on this page has a date earlier than October 2026, all it
        means is that the Pokemon was caught and catalogued prior to me starting this project.
      </div>
      <h1 class="text-center">Latest Updates</h1>
    </div>
    <Events category="POKEDEX" />
  </v-sheet>
</template>
