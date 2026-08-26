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
  <div>
    <div class="columns">
      <div class="column">
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <h1 class="title is-3">{{ props.name }}</h1>
            </div>
          </div>
        </div>
        <div class="level" v-if="props.includeLink">
          <div class="level-left">
            <div class="level-item">
              <router-link :to="`/games/${props.gameKey}`" class="button is-primary">View</router-link>
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