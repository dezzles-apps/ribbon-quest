<script setup lang="ts">

import type { PropType } from 'vue';
import type { PokemonDetails } from '@/types/api.ts';

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
  return `/sprites/${props.details.pokemon.toLowerCase()}.png`;
}
</script>


<template>
  <div>
    <div class="columns">
      <div class="column">
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h1 class="title is-3">{{ props.details.nickname }}</h1>
              <span :class="['icon', 'has-text-success' ]" style="margin-left: 15px" v-if="props.details.shiny">
                <i class="mdi mdi-shimmer mdi-36px"></i>
              </span>
            </div>
            <div class="level-item">
              <h2 class="subtitle is-5">{{ props.details.pokemon }}</h2>
            </div>
          </div>
        </div>
        <div>
          <table>
            <tr>
              <td colspan="2">
                {{ props.details.nickname }} is representing the {{ props.details.region }} region.
              </td>
            </tr>
            <tr v-if="props.details.nature">
              <td>{{props.details.nickname }} is {{ props.details.nature ?? 'Unknown' }}</td>
            </tr>
            <tr v-if="props.details.characteristic">
              <td colspan="2">{{ props.details.characteristic }}</td>
            </tr>
            
          </table>
        </div>
        <div class="level" v-if="props.includeLink">
          <div class="level-left">
            <div class="level-item">
              <router-link :to="`/pokemon/${props.details.pokemon}`" class="button is-primary">View</router-link>
            </div>
          </div>
        </div>
      </div>
      <div class="column is-narrow">
          <div class="level-right" v-if="props.current !== undefined && props.total !== undefined">
            <div class="level-item">
              <span class="tag" :class="getCounterClass()">{{ props.current }} / {{ props.total }}</span>
            </div>
          </div>
        <img
          :src="getPokemonImage()"
          :alt="props.details.pokemon"
          class="pokemon-image"
          :class="!props.details.caughtAt ? 'pokemon-image-not-caught': ''"
        />
      </div>
    </div>

  </div>
</template>

<style lang="css" scoped>
.pokemon-image {
  max-height: 150px;
  display: block;
  margin: auto;
}

.pokemon-image-not-caught {
  opacity: 0.3;
}
</style>