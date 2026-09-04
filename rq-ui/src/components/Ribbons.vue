<script setup lang="ts">
import type { PokemonRibbon } from '@/types/api';
import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';
import { useApi } from '@/composables/useApi';
import API from '@/composables/endpoints';

const authStore = useAuthStore();
const api = useApi();

const props = defineProps<{
  title?: string;
  pokemon: string;
  ribbons: PokemonRibbon[]
}>()

const loadingRibbons = ref(new Map<string, boolean>());

function getRibbonClass(ribbon: PokemonRibbon): string[] {
  const classes: string[] = [];
  if (ribbon.achieved) {
    classes.push('ribbon-achieved');
  } else {
    classes.push('ribbon-not-achieved');
  }
  if (authStore.isAuthenticated) {
    if (loadingRibbons.value.get(ribbon.ribbonKey)) {
      classes.push('ribbon-loading');
    } else {
      classes.push('ribbon-clickable');
    }
  }
  classes.push(`ribbon-${ribbon.category.toLowerCase()}`);

  return classes;
}

function toggleRibbon(ribbon: PokemonRibbon) {
  if (!authStore.isAuthenticated) {
    return;
  }

  const isLoading = loadingRibbons.value.get(ribbon.ribbonKey) || false;
  if (isLoading) {
    return;
  }

  loadingRibbons.value.set(ribbon.ribbonKey, true);
  let method = ribbon.achieved ? 'DELETE' : 'POST';
  api.apiFetch(API.Ribbons.UpdateRibbon(props.pokemon, ribbon.ribbonKey), {
    method: method,
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(async response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      let newRibbon = await response.json();
      ribbon.achieved = newRibbon.data.achieved;
      ribbon.achievedAt = newRibbon.data.achievedAt;
    })
    .catch(error => {
      console.error('Error toggling ribbon:', error);
    })
    .finally(() => {
      loadingRibbons.value.set(ribbon.ribbonKey, false);
    });
}


</script>

<template>
  <div class="box mb-4">
    <slot name="title">
      <h2 class="title is-4">{{ props.title }}</h2>
    </slot>
    <div style="display: inline-flex; flex-wrap: wrap;">
      <div
        class="ribbon"
        :class="getRibbonClass(ribbon)"
        v-for="ribbon in props.ribbons"
        :key="ribbon.ribbonKey"
        @click="toggleRibbon(ribbon)"
      >
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
  align-items: center;
  text-align: center;
  width: 110px;
  height: 100px;
  outline-width: 5px;
  outline-style: solid;
  border-radius: 10px;
}

.ribbon-loading {
  cursor: wait;
}

.ribbon-clickable {
  cursor: pointer;
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