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
      path: '/admin/ribbons/pokemon',
      name: 'admin-pokemon',
      component: () => import('@/pages/admin/PokemonView.vue'),
      meta: {
        crumbs: {
          parent: 'home',
          title: 'Ribbon Pokemon',
          href: '/admin/ribbons/pokemon'
        }
      }
    },
    {
      path: '/admin/ribbons/pokemon/new',
      name: 'admin-add-ribbon-pokemon',
      component: () => import('@/pages/admin/AddRibbonPokemonView.vue'),
      meta: {
        crumbs: {
          parent: 'admin-pokemon',
          title: 'New Ribbon Pokemon',
          href: '/admin/ribbons/pokemon/new'
        }
      }
    },
    {
      path: '/admin/ribbons/pokemon/:pokemonId',
      name: 'admin-edit-ribbon-pokemon',
      component: () => import('@/pages/admin/EditRibbonPokemonView.vue'),
      meta: {
        crumbs: {
          parent: 'admin-pokemon',
          title: (i: any) => (`Edit ${i.pokemonId}`),
          href: (i: any) => (`/admin/ribbons/pokemon/${i.pokemonId}`)
        }
      }
    },
    {
      path: '/dexofmyown',
      name: 'dexofmyown-info',
      component: () => import('@/pages/pokedexes/InfoView.vue'),
      meta: {
        crumbs: {
          parent: 'home',
          title: 'Dex of My Own',
          href: '/dexofmyown'
        }
      }
    },
    {
      path: '/dexofmyown/pokedexes',
      name: 'dexofmyown-pokedexes',
      component: () => import('@/pages/pokedexes/PokedexesView.vue'),
      meta: {
        crumbs: {
          parent: 'dexofmyown-info',
          title: 'Pokedexes',
          href: '/dexofmyown/pokedexes'
        }
      }
    },
    {
      path: '/dexofmyown/pokedexes/:pokedex',
      name: 'dexofmyown-pokedex',
      component: () => import('@/pages/pokedexes/PokedexView.vue'),
      meta: {
        crumbs: {
          parent: 'dexofmyown-pokedexes',
          title: (i: any) => (`${i.pokedex}`),
          href: (i: any) => (`/dexofmyown/pokedexes/${i.pokedex}`)
        }
      }
    }
  ],
})

export default router
