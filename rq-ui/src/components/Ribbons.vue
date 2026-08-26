<script setup lang="ts">
import type { PokemonRibbon } from '@/types/api';

const props = defineProps<{
  title?: string;
  ribbons: PokemonRibbon[]
}>()

function getRibbonClass(ribbon: PokemonRibbon): string[] {
  const classes: string[] = [];
  if (ribbon.achieved) {
    classes.push('ribbon-achieved');
  } else {
    classes.push('ribbon-not-achieved');
  }
  classes.push(`ribbon-${ribbon.category.toLowerCase()}`);

  return classes;
}
</script>

<template>
  <div class="box mb-4">
    <slot name="title">
      <h2 class="title is-4">{{ props.title }}</h2>
    </slot>
    <div style="display: inline-flex; flex-wrap: wrap;">
      <div class="ribbon" :class="getRibbonClass(ribbon)" v-for="ribbon in props.ribbons" :key="ribbon.ribbonKey">
        <span :class="['icon', ribbon.achieved ? 'has-text-success' : 'has-text-grey']">
          <i class="mdi mdi-seal mdi-36px"></i>
        </span>
        <div class="ribbon-name">
          {{ ribbon.name }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ribbon {
  margin: 10px;
  
  padding-top: 15px;
  align-items: center;
  text-align: center;
  width: 115px;
  height: 100px;
  outline-width: 5px;
  outline-style: solid;
  border-radius: 10px;
}
.ribbon-name {
  font-weight: bold;
  margin-right: 10px;
  width: 100%;
  color: darkslategray;
}

.ribbon-not-achieved {
  opacity: 0.5;
}

.ribbon-julie {
  background-color: #DCB0F2;
  outline-color: #a35ec9;
}
.ribbon-champion {
  background-color: #66C5CC;
  outline-color: #358d95;
}
.ribbon-battle {
  background-color: #F89C74;
  outline-color: #754834;
}
.ribbon-contest {
  background-color: #FE88B1;
  outline-color: #8e3d5a;
}
.ribbon-stats {
  background-color: #9EB9F3;
  outline-color: #62749b;
}
.ribbon-shopping {
  background-color: #F6CF71;
  outline-color: #d5a32c;
}
</style>