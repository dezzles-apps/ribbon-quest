<script setup lang="ts">

import { useRoute, useRouter } from 'vue-router';
import { useCrumbs } from '@/composables/useCrumbs';

const crumbs = useCrumbs();

</script>

<template>
  <v-breadcrumbs :items="crumbs.getCrumbs()"/>
</template>