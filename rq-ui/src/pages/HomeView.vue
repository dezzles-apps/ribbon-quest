<script setup lang="ts">
import Events from '@/components/Events.vue'
</script>

<template>
      <v-sheet>
        <div class="box ma-4">
          <h1 class="title is-4 text-center">Dezzles' Pokemon Challenges</h1>
          <div class="w-50 ma-auto">
            2025 was all about catching shiny Pokemon and if you want to see what I caught, check out the 
            <a href="http://shiny-hunt.dezzles-apps.com" target="_blank">shiny hunt mini-site</a>!
            <br/>
            I've decided that from now one I'm going to start centralising all of my insane Pokemon
            challenges into a single website. At the moment this site is running the Ribbon Quest I'm
            currently on, but will expand as I come up with more silly ideas.
          </div>
          <Events />
        </div>

      </v-sheet>
</template>
