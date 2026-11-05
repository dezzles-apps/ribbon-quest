import type { Meta, StoryObj } from '@storybook/vue3-vite'
import { expect, within } from 'storybook/test'
import { reactive } from 'vue'
import { createMemoryHistory, createRouter, routeLocationKey, routerKey } from 'vue-router'
import Crumbs from './Crumbs.vue'

const Empty = { render: () => null }

const meta: Meta<typeof Crumbs> = {
  title: 'Components/Crumbs',
  component: Crumbs,
  decorators: [
    story => {
      const router = createRouter({
        history: createMemoryHistory(),
        routes: [
          { path: '/', name: 'home', component: Empty, meta: { crumbs: { parent: null, title: 'Home', href: '/' } } },
          { path: '/ribbons', name: 'ribbons', component: Empty, meta: { crumbs: { parent: 'home', title: 'Ribbon Quest', href: '/ribbons' } } },
        ],
      })
      const route = router.resolve('/ribbons')
      return {
        components: { story },
        provide: { [routerKey as symbol]: router, [routeLocationKey as symbol]: reactive(route) },
        template: '<story />',
      }
    },
  ],
}
export default meta

export const Nested: StoryObj<typeof Crumbs> = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(await canvas.findByText('Home')).toBeVisible()
    await expect(await canvas.findByText('Ribbon Quest')).toBeVisible()
  },
}
