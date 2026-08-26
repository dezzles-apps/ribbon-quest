<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'
const isActive = ref(false)

function toggleNavbar() {
  isActive.value = !isActive.value
}
</script>

<template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <a class="navbar-item" href="/">
        <img src="/favicon.svg" />
        Dezzles' Ribbon Quest
      </a>

      <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="rq-navbar" @click="toggleNavbar">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

    <div id="rq-navbar" class="navbar-menu" :class="isActive ? 'is-active' : ''">
      <div class="navbar-start">
        <a class="navbar-item" href="/">
          Home
        </a>

        <a class="navbar-item" href="/pokemon">
          Pokemon
        </a>

      </div>

    </div>
  </nav>
  <section class="section">
    <div class="container">
      <RouterView />
    </div>
  </section>
  <footer class="footer has-background-dark has-text-light">
    <div class="content has-text-centered">
      <p>
        <strong>Dezzles' Ribbon Quest</strong> © 2026
      </p>
    </div>
  </footer>

</template>

<style scoped>

</style>
