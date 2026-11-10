import type { Meta, StoryObj } from '@storybook/vue3-vite'
import { expect, within } from 'storybook/test'
import Ribbon from './Ribbon.vue'
import { RibbonTheme, RibbonState } from '@/types/enums.ts'

const ribbons = [
  { ribbonKey: 'a', name: 'Alpha Ribbon', achieved: true, achievedAt: '2024-01-01', category: 'Contest' },
  { ribbonKey: 'b', name: 'Beta Ribbon', achieved: false, achievedAt: '', category: 'Battle' },
]

const meta = {
  title: 'Components/Ribbon',
  component: Ribbon,
  args: { },
} satisfies Meta<typeof Ribbon>
export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    bottomText: 'Effort',
    theme: RibbonTheme.Stats,
    achieved: true,
    icon: 'mdi-seal',
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(canvas.getByText('Effort')).toBeVisible()
  },
}

export const TopText: Story = {
  args: {
    topText: 'Effort',
    theme: RibbonTheme.Stats,
    achieved: true,
    icon: 'mdi-seal',
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(canvas.getByText('Effort')).toBeVisible()
  },
}

export const BothText: Story = {
  args: {
    topText: 'Ribbon',
    bottomText: 'Effort',
    theme: RibbonTheme.Stats,
    achieved: true,
    icon: 'mdi-seal',
  },
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(canvas.getByText('Ribbon')).toBeVisible()
    await expect(canvas.getByText('Effort')).toBeVisible()
  },
}

export const LongText: Story = {
  args: {
    bottomText: 'Toughness Master',
    theme: RibbonTheme.Contest,
    icon: 'mdi-seal',
  },
}

export const NotAchieved: Story = {
  args: {
    bottomText: 'Jeff',
    theme: RibbonTheme.Champion,
    icon: 'mdi-seal',
    achieved: false,
  },
}

export const Loading: Story = {
  args: {
    bottomText: 'Jeff',
    theme: RibbonTheme.Champion,
    icon: 'mdi-seal',
    state: RibbonState.Loading,
  },
}

export const Clickable: Story = {
  args: {
    bottomText: 'Jeff',
    theme: RibbonTheme.Champion,
    icon: 'mdi-seal',
    state: RibbonState.Clickable,
  },
}

export const Highlighted: Story = {
  args: {
    bottomText: 'Jeff',
    theme: RibbonTheme.Champion,
    highlighted: true,
    icon: 'mdi-seal',
  }
}

export const PokedexEntry: Story = {
  args: {
    achieved: true,
    topText: 'NAT-0001',
    bottomText: 'Bulbasaur',
    theme: RibbonTheme.Champion,
    image: 'Dezzles.png',
  }
}