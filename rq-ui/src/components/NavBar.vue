<script setup>
  import { ref, watch } from 'vue'

  const drawer = ref(false)
  const group = ref(null)

  watch(group, () => {
    drawer.value = false
  })
</script>

<template>
  <v-app-bar
    color="primary"
    density="compact"
  >
    <template v-slot:prepend>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    </template>

    <v-app-bar-title>Dezzles Pokemon Stuff</v-app-bar-title>

    <template v-slot:append>
      <v-btn icon="mdi-dots-vertical"></v-btn>
    </template>
  </v-app-bar>
  <v-navigation-drawer
    v-model="drawer"
    :location="$vuetify.display.mobile ? 'bottom' : undefined"
    temporary
  >
    <v-list>
      <v-list-item href="/" title="Home" />
      <v-list-item href="/ribbons" title="Ribbon Quest" />
  </v-list>
  </v-navigation-drawer>
</template>