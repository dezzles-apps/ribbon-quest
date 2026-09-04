/**
 * router/index.ts
 *
 * Manual routes for ./src/pages/*.vue
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/pages/HomeView.vue'
import { Games } from '@/data/games';

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
          title: (i: any) => i.pokemon,
          href: (i: any) => (`/ribbons/pokemon/${i.pokemon}`)
        }
      }
    },
    {
      path: '/ribbons/games',
      name: 'ribbons-all-games',
      component: () => import('@/pages/ribbons/AllGamesView.vue'),
      meta: {
        crumbs: {
          parent: 'ribbons',
          title: 'Games',
          href: '/ribbons/games'
        }
      }
    },
    {
      path: '/ribbons/games/:game',
      name: 'ribbons-game',
      component: () => import('@/pages/ribbons/GameView.vue'),
      meta: {
        crumbs: {
          parent: 'ribbons-all-games',
          title: (i: any) => Games[i.game] || i.game,
          href: (i: any) => (`/ribbons/games/${i.game}`)
        }
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/pages/LoginView.vue'),
      meta: {
        crumbs: {
          parent: 'home',
          title: 'Login',
          href: '/login'
        }
      }
    },
    {
      path: '/admin/pokemon',
      name: 'admin-pokemon',
      component: () => import('@/pages/admin/PokemonView.vue'),
      meta: {
        crumbs: {
          parent: 'home',
          title: 'Ribbon Pokemon',
          href: '/admin/pokemon'
        }
      }
    }
  ],
})

export default router
