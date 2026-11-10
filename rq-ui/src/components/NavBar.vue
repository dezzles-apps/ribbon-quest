<script setup lang="ts">

import { computed, ref } from 'vue';
import { useTheme } from 'vuetify';
import { useAuthStore } from '@/stores/authStore';
import Crumbs from '@/components/Crumbs.vue'

const drawer = ref(false);

const authStore = useAuthStore();
const theme = useTheme();
const isDark = computed(() => theme.global.current.value.dark);

function toggleTheme() {
  const nextTheme = isDark.value ? 'light' : 'dark';
  theme.global.name.value = nextTheme;
  localStorage.setItem('theme', nextTheme);
}

function logout() {
  authStore.clearToken();
  drawer.value = false;
}

</script>

<template>
  <v-app-bar>
    <template v-slot:prepend>
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    </template>

    <v-app-bar-title>Dezzles' Pokemon Challenges</v-app-bar-title>

    <v-btn
      icon
      :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
      :title="isDark ? 'Switch to light mode' : 'Switch to dark mode'"
      @click="toggleTheme"
    >
      <v-icon>{{ isDark ? 'mdi-weather-sunny' : 'mdi-weather-night' }}</v-icon>
    </v-btn>
  </v-app-bar>
  <v-navigation-drawer
    v-model="drawer"
    :location="$vuetify.display.mobile ? 'bottom' : undefined"
    temporary
    width="300"
  >
    <v-list>
      <v-list-item href="/">Home</v-list-item>
      <v-list-group value="Ribbon Quest">
        <template v-slot:activator="{ props }">
          <v-list-item
            v-bind="props"
            prepend-icon="mdi-seal"
            title="Ribbon Quest"
          ></v-list-item>
        </template>

        <v-list-item href="/ribbons">Info</v-list-item>
        <v-list-item href="/ribbons/games">By Game</v-list-item>
        <v-list-item href="/ribbons/pokemon">By Pokemon</v-list-item>
      </v-list-group>
      <v-list-group>
        <template v-slot:activator="{ props }">
          <v-list-item
            v-bind="props"
            prepend-icon="mdi-notebook"
            title="A Dex of My Own"
          ></v-list-item>
        </template>
        <v-list-item
          href="/dexofmyown"
          title="Info"
        />
        <v-list-item
          href="/dexofmyown/pokedexes"
          title="Pokedexes"
        />

      </v-list-group>
      <v-divider />
      <v-list-group>
        <template v-slot:activator="{ props }">
          <v-list-item
            v-bind="props"
            prepend-icon="mdi-shimmer"
            title="Shinies"
          ></v-list-item>
        </template>
        <v-list-item
          href="/shinies"
          title="Stats"
        />
        <v-list-item
          href="/shinies/catches"
          title="Shinies Caught"
        />

      </v-list-group>
      <v-divider />
      <v-list-item
        href="/login"
        v-if="!authStore.isAuthenticated"
      >
        Login
      </v-list-item>
      <v-list-group
        value="Admin"
        v-if="authStore.isAuthenticated"
      >
        <template v-slot:activator="{ props }">
          <v-list-item
            v-bind="props"
            prepend-icon="mdi-shield-crown"
            title="Admin"
          ></v-list-item>
        </template>

        <v-list-item href="/admin/ribbons/pokemon">Ribbon Pokemon</v-list-item>
        <v-list-item href="/admin/shinies/catch">Catch Shiny Pokemon</v-list-item>
        <v-list-item link @click="logout">Logout</v-list-item>
      </v-list-group>
    </v-list>
  </v-navigation-drawer>
  <Crumbs />
</template>