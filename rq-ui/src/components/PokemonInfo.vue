<script setup lang="ts">

import type { PropType } from 'vue';
import type { PokemonDetails } from '@/types/api.ts';
import { useDates } from '@/composables/useDates';

const dates = useDates();

const props = defineProps({
  details: {
    type: Object as PropType<PokemonDetails>,
    required: true
  },
  current: Number,
  total: Number,
  includeLink:{
    type: Boolean,
    default: false
  },
  includeDescription:{
    type: Boolean,
    default: true
  }
});

function getCounterClass(): string[] {
  if (props.current === undefined || props.total === undefined) {
    return [];
  }
  const classes: string[] = [];
  if (props.current >= props.total) {
    classes.push('is-success');
  } else {
    classes.push('is-danger');
  }

  return classes;
}

function getPokemonImage(): string {
  return `/sprites/${props.details.pokemon.toLowerCase()}${props.details.shiny ? '-s' : ''}.png`;
}
</script>


<template>
  <v-card
    :title="props.details.nickname"
    :subtitle="props.details.pokemon"
  >
    <v-card-text>
      <table>
        <tbody>
          <tr>
            <td colspan="2">
              {{ props.details.nickname }} is representing the {{ props.details.region }} region.
            </td>
          </tr>
          <tr v-if="props.details.notes">
            <td>{{ props.details.notes }}</td>
          </tr>
          <tr v-if="props.details.nature">
            <td>{{props.details.nickname }} is {{ props.details.nature ?? 'Unknown' }}</td>
          </tr>
          <tr v-if="props.details.characteristic">
            <td colspan="2">{{ props.details.characteristic }}</td>
          </tr>
          <tr v-if="props.details.caughtAt">
            <td>Caught at:</td>
            <td>{{ dates.toLocal(props.details.caughtAt) }}</td>
          </tr>
        </tbody>
      </table>
      <slot name="content" />
    </v-card-text>
    <template v-slot:prepend>
      <img
        :src="getPokemonImage()"
        :alt="props.details.pokemon"
        class="pokemon-image"
        :class="!props.details.caughtAt ? 'pokemon-image-not-caught': ''"
      />
    </template>
    <template v-slot:append>
      <v-chip
        v-if="props.current !== undefined && props.total !== undefined"
        variant="flat"
        :color="props.current == props.total ? 'primary' : 'secondary'"
      >
        {{ props.current }} / {{ props.total }}
      </v-chip>
      <v-btn
        v-if="props.includeLink"
        variant="tonal"
        :to="{ name: 'ribbons-pokemon', params: { pokemon: props.details.pokemon }}"
      >
        More Details
      </v-btn>
    </template>
  </v-card>



</template>

<style lang="css" scoped>
.pokemon-image {
  max-height: 75px;
  display: block;
  margin: auto;
}

.pokemon-image-not-caught {
  opacity: 0.3;
}
</style>