import type { Meta, StoryObj } from '@storybook/vue3-vite'
import { expect, within } from 'storybook/test'
import Loading from './Loading.vue'

const meta = { title: 'Components/Loading', component: Loading } satisfies Meta<typeof Loading>
export default meta

export const Default: StoryObj<typeof meta> = {
  play: async ({ canvasElement }) => {
    const canvas = within(canvasElement)
    await expect(canvas.getByRole('progressbar')).toBeInTheDocument()
    await expect(canvas.getByText('Loading')).toBeVisible()
  },
}
