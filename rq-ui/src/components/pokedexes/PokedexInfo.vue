<script setup lang="ts">

import type { PropType } from 'vue';
import type { Pokedex } from '@/types/pokedex';
import { useDates } from '@/composables/useDates';

const dates = useDates();

const props = defineProps({
  details: {
    type: Object as PropType<Pokedex>,
    required: true
  },
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
  if (props.details.stats.caught === undefined || props.details.stats.total === undefined) {
    return [];
  }
  const classes: string[] = [];
  if (props.details.stats.caught >= props.details.stats.caught) {
    classes.push('is-success');
  } else {
    classes.push('is-danger');
  }

  return classes;
}

function getImage(): string {
  return props.details.image
}
</script>


<template>
  <v-card
    :title="`${props.details.name} Pokedex`"
  >
    <v-card-text>
      <table>
        <tbody>
          <tr v-if="props.details.notes">
            <td colspan="2">{{ props.details.notes }}</td>
          </tr>
        </tbody>
      </table>
      <slot name="content" />
    </v-card-text>
    <template v-slot:prepend>
      <img
        :src="getImage()"
        :alt="props.details.name"
        class="pokemon-image"
      />
    </template>
    <template v-slot:append>
      <v-chip
        v-if="props.details.stats.caught !== undefined && props.details.stats.total !== undefined"
        variant="flat"
        :color="props.details.stats.caught == props.details.stats.total ? 'primary' : 'secondary'"
      >
        {{ props.details.stats.caught }} / {{ props.details.stats.total }}
      </v-chip>
      <v-btn
        v-if="props.includeLink"
        variant="tonal"
        :to="{ name: 'dexofmyown-pokedex', params: { pokedex: props.details.name }}"
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