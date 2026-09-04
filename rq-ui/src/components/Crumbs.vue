<script setup lang="ts">

import { useCrumbs } from '@/composables/useCrumbs';

const crumbs = useCrumbs();

</script>

<template>
  <v-breadcrumbs :items="crumbs.getCrumbs()"/>
</template>