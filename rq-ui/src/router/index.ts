/**
 * router/index.ts
 *
 * Manual routes for ./src/pages/*.vue
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/pages/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Index,
      meta: {
        crumbs: {
          parent: null,
          title: 'Home',
          href: '/'
        }
      }
    },
    {
      path: '/ribbons',
      name: 'ribbons',
      component: () => import('@/pages/ribbons/RibbonsView.vue'),
      meta: {
        crumbs: {
          parent: 'home',
          title: 'Ribbon Quest',
          href: '/ribbons'
        }
      }
    },
    {
      path: '/ribbons/pokemon',
      name: 'ribbons-all-pokemon',
      component: () => import('@/pages/ribbons/AllPokemon.vue'),
      meta: {
        crumbs: {
          parent: 'ribbons',
          title: 'Pokemon',
          href: '/ribbons/pokemon'
        }
      }
    },
    {
      path: '/ribbons/pokemon/:pokemon',
      name: 'ribbons-pokemon',
      component: () => import('@/pages/ribbons/PokemonView.vue'),
      meta: {
        crumbs: {
          parent: 'ribbons-all-pokemon',
          title: i => i.pokemon,
          href: i => (`/ribbons/pokemon/${i.pokemon}`)
        }
      }
    },
    {
      path: '/ribbons/games',
      name: 'ribbons-all-games',
      component: () => import('@/pages/ribbons/AllGamesView.vue'),
    },
    {
      path: '/ribbons/games/:game',
      name: 'ribbons-game',
      component: () => import('@/pages/ribbons/GameView.vue'),
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
    }
  ],
})

export default router
