<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useApi } from '@/composables/useApi';
import Loading from '@/components/Loading.vue';
import type { Event } from '@/types/events';
import type { Page, Response } from '@/types/responses';
import RibbonColours from '@/composables/ribbons';
import { useDates } from '@/composables/useDates';
import API from '@/composables/endpoints';
const loading = ref(true);
const events = ref([] as Event[] );
const next = ref('' as string | null);
const { apiFetch } = useApi();
const allLoaded = ref(false)
const dates = useDates();

function loadEvents() {
  let url = API.Events.GetAllEvents
  if (next.value) {
    url = url + '?' + new URLSearchParams({
      key: next.value
    })
  }
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
  <h1 class="text-center">Latest Updates</h1>
  <Loading v-if="loading" />
  <div v-else>
    <v-timeline
      side="end"
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
                Caught a{{ ['A','E', 'I', 'O', 'U'].indexOf(item.metadata.pokemon[0]) == 0 ? 'n' : '' }}
                <RouterLink :to="{ name: 'ribbons-pokemon', params: { pokemon: item.metadata.pokemon }}">{{  item.metadata.pokemon }}</RouterLink>
                <div v-if="item.metadata.nickname != item.metadata.pokemon" style="display: inline-block">
                  &nbsp;named {{ item.metadata.nickname }}
                </div>
                for the <RouterLink :to="{ name: 'ribbons' }">Ribbon Quest</RouterLink>
              </div>
              <div v-else>
                Unknown event type: {{  item.eventType }}
              </div>
            </div>
          </div>
        </v-alert>
      </v-timeline-item>
    </v-timeline>
    <div class="ma-auto" style="text-align: center;">
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
  </div>
</template>