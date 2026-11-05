import type { Meta, StoryObj } from '@storybook/vue3-vite'
import { expect, within } from 'storybook/test'
import Ribbons from './Ribbons.vue'

const ribbons = [
  { ribbonKey: 'a', name: 'Alpha Ribbon', achieved: true, achievedAt: '2024-01-01', category: 'Contest' },
  { ribbonKey: 'b', name: 'Beta Ribbon', achieved: false, achievedAt: '', category: 'Battle' },
]

const meta = {
  title: 'Components/Ribbons',
  component: Ribbons,
  args: { title: 'Test Ribbons', pokemon: 'pikachu', ribbons },
} satisfies Meta<typeof Ribbons>
export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(canvas.getByText('Test Ribbons')).toBeVisible()
    await expect(canvas.getByText('Alpha Ribbon')).toBeVisible()
    await expect(canvas.getByText('Beta Ribbon')).toBeVisible()
  },
}

export const HideCompleted: Story = {
  args: { filter: { hideCompletePokemon: false, hideCompleteRibbons: true } },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(canvas.queryByText('Alpha Ribbon')).toBeNull()
    await expect(canvas.getByText('Beta Ribbon')).toBeVisible()
  },
}
