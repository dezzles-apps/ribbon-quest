<script setup lang="ts">

import { ref } from 'vue';
import { useAuthStore } from '@/stores/authStore';

const drawer = ref(false);
const open = ref(['Ribbon Quest'])

const authStore = useAuthStore();

import Crumbs from '@/components/Crumbs.vue'

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
  </v-app-bar>
  <v-navigation-drawer
    v-model="drawer"
    :location="$vuetify.display.mobile ? 'bottom' : undefined"
    temporary
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

        <v-list-item href="/admin/pokemon">Ribbon Pokemon</v-list-item>
        <v-list-item link @click="logout">Logout</v-list-item>
      </v-list-group>
    </v-list>
  </v-navigation-drawer>
  <Crumbs />
</template>