<script setup lang="ts">

const props = defineProps({
  gameKey: String,
  name: String,
  current: Number,
  total: Number,
  includeLink:{
    type: Boolean,
    default: false
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


</script>


<template>
  <v-card
    :title="props.name"
  >
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
        :to="{ name: 'ribbons-game', params: { game: props.gameKey }}"
      >
        More Details
      </v-btn>
    </template>
    <v-card-text>
      <slot name="content" />
    </v-card-text>
  </v-card>
</template>