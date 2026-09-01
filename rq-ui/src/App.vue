<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'
import { ref } from 'vue'
const isActive = ref(false);
const authStore = useAuthStore();
const router = useRouter()
function toggleNavbar() {
  isActive.value = !isActive.value;
}

function logout() {
  authStore.clearToken();
  router.push('/');
  toggleNavbar();
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
        <a class="navbar-item" href="/games" @click="toggleNavbar">
          Games
        </a>
        <a class="navbar-item" href="/pokemon" @click="toggleNavbar">
          Pokemon
        </a>
        <a class="navbar-item" href="/admin/pokemon" v-if="authStore.isAuthenticated" @click="toggleNavbar">
          Admin
        </a>
        <a class="navbar-item" href="/login" v-if="!authStore.isAuthenticated" @click="toggleNavbar">
          Login
        </a>
        <a class="navbar-item" @click="logout" v-if="authStore.isAuthenticated">
          Logout
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
