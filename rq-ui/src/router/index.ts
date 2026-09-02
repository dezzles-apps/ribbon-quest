/**
 * router/index.ts
 *
 * Manual routes for ./src/pages/*.vue
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/pages/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Index,
    },
    {
      path: '/ribbons/pokemon',
      name: 'all-pokemon',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('@/pages/ribbon-quest/AllPokemon.vue'),
    },
    {
      path: '/ribbons',
      name: 'ribbons-home',
      component: () => import('@/pages/ribbon-quest/HomeView.vue'),
    },
    {
      path: '/ribbons/pokemon/:pokemon',
      name: 'pokemon',
      component: () => import('@/pages/ribbon-quest/PokemonView.vue'),
    },
    {
      path: '/ribbons/games',
      name: 'all-games',
      component: () => import('@/pages/ribbon-quest/AllGamesView.vue'),
    },
    {
      path: '/ribbons/games/:game',
      name: 'game',
      component: () => import('@/pages/ribbon-quest/GameView.vue'),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/pages/LoginView.vue'),
    },
    {
      path: '/admin/pokemon',
      name: 'admin-pokemon',
      component: () => import('@/pages/admin/PokemonView.vue'),
    },
  ],
})

export default router
