<script setup lang="ts">

const props = defineProps<{
  pokemon: string,
  nickname: string,
  region: string,
  current?: number,
  target?: number,
  includeLink: boolean
}>()

function getCounterClass(): string[] {
  if (props.current === undefined || props.target === undefined) {
    return [];
  }
  const classes: string[] = [];
  if (props.current >= props.target) {
    classes.push('is-success');
  } else {
    classes.push('is-danger');
  }

  return classes;
}

function getPokemonImage(): string {
  return `/sprites/${props.pokemon.toLowerCase()}.png`;
}
</script>


<template>
  <div>
    <div class="columns">
      <div class="column">
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h1 class="title is-3">{{ props.nickname }}</h1>
            </div>
          </div>
        </div>
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h2 class="subtitle is-5">{{ props.pokemon }}</h2>
            </div>
          </div>
        </div>
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              This {{ props.pokemon }} is originally from the {{ props.region }} region.
            </div>
          </div>
        </div>
        <div class="level" v-if="props.includeLink">
          <div class="level-left">
            <div class="level-item">
              <router-link :to="`/pokemon/${props.pokemon}`" class="button is-primary">View</router-link>
            </div>
          </div>
        </div>
      </div>
      <div class="column is-narrow">
          <div class="level-right" v-if="props.current !== undefined && props.target !== undefined">
            <div class="level-item">
              <span class="tag" :class="getCounterClass()">{{ props.current }} / {{ props.target }}</span>
            </div>
          </div>
        <img :src="getPokemonImage()" :alt="props.pokemon" class="pokemon-image" />
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
</style>