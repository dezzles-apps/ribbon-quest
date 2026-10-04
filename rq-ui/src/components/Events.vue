<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useApi } from '@/composables/useApi';
import Loading from '@/components/Loading.vue';
import type { Event } from '@/types/events';
import type { Page, Response } from '@/types/responses';
import RibbonColours from '@/composables/ribbons';
import { useDates } from '@/composables/useDates';
import API from '@/composables/endpoints';

const props = defineProps({
  pokemon: {
    type: String,
    required: false
  },
  category: {
    type: String,
    required: false
  }
})
const loading = ref(true);
const events = ref([] as Event[] );
const next = ref('' as string | null);
const { apiFetch } = useApi();
const allLoaded = ref(false)
const dates = useDates();

function loadEvents() {
  let url = API.Events.GetAllEvents
  let params = {} as any
  if (next.value) {
    params.key = next.value
  }
  if (props.category) {
    params.category = props.category
  }
  if (props.pokemon) {
    params.pokemon = props.pokemon
  }
  url = url + '?' + new URLSearchParams(params)
  return apiFetch(url)
    .then(async resp => {
      if (!resp.ok) {
        throw new Error("Failed to load")
      }
      const data = await resp.json() as Response<Page<Event>>
      data.data.items.forEach(item => {
        events.value.push(item)
      })
      next.value = data.data.next
      if (!next.value) {
        allLoaded.value = true
      }
      loading.value = false
    })
}

function getDotColour(event : Event) {
  if (event.eventType === 'RIBBON') {
    return RibbonColours[event.metadata.ribbonCategory.toLowerCase()].background
  }
  return 'grey'
}


function getIcon(event : Event) {
  switch (event.eventType) {
    case 'RIBBON':
      return 'mdi-seal'
    case 'RIBBON_CATCH':
      return 'mdi-pokeball'
    default:
      return 'mdi-help-circle-outline'
  }
}

onMounted(loadEvents)

</script>


<template>
  <Loading v-if="loading" />
  <div v-else style="width: 100%;">
    <v-timeline
      side="end"
      class="mx-auto"
    >
      <v-timeline-item
        v-for="(item, idx) in events"
        :key="idx"
        :dot-color="getDotColour(item)"
        size="small"
      >
        <v-alert
          :icon="getIcon(item)"
          :value="true"
        >
          <div class="d-flex">
            <strong class="me-4">{{ dates.toLocal(item.eventTime) }}</strong>
            <div class="text-body-large me-4">
              <div v-if="item.eventType == 'RIBBON'">
                <RouterLink :to="{ name: 'ribbons-pokemon', params: { pokemon: item.metadata.pokemon }}">{{  item.metadata.nickname }}</RouterLink>
                earned the {{ item.metadata.ribbonName }} {{ item.metadata.ribbonType.toLowerCase() }} 
              </div>
              <div v-else-if="item.eventType == 'RIBBON_CATCH'" style="display: inline">
                Caught a{{ ['A','E', 'I', 'O', 'U'].indexOf(item.metadata.pokemon[0]) != -1 ? 'n' : '' }}
                <RouterLink :to="{ name: 'ribbons-pokemon', params: { pokemon: item.metadata.pokemon }}">{{  item.metadata.pokemon }}</RouterLink>
                <div v-if="item.metadata.nickname != item.metadata.pokemon" style="display: inline-block">
                  &nbsp;named {{ item.metadata.nickname }}
                </div>
                for the <RouterLink :to="{ name: 'ribbons' }">Ribbon Quest</RouterLink>
              </div>
              <div v-else-if="item.eventType == 'POKEDEX_CATCH'" style="display: inline;">
                 Caught a{{ ['A','E', 'I', 'O', 'U'].indexOf(item.metadata.pokemon[0]) != -1 ? 'n' : '' }} {{  item.metadata.pokemon }} for the 
                <RouterLink :to="{ name: 'dexofmyown-pokedex', params: { pokedex: item.metadata.pokedex }}">{{  item.metadata.pokedex }} Pokedex</RouterLink>
              </div>
              <div v-else>
                Unknown event type: {{  item.eventType }}
              </div>
            </div>
          </div>
        </v-alert>
      </v-timeline-item>
    </v-timeline>
  </div>
    <div class="ma-auto pb-3" style="text-align: center;">
      <v-btn
        v-if="!allLoaded"
        color="green"
        type="flat"
        :loading="loading"
        @click="loadEvents()"
      >
        Load More...
      </v-btn>
      <div v-if="allLoaded">No more events...</div>
    </div>
</template>