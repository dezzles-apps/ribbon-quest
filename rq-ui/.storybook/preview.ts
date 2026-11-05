import type { Preview } from '@storybook/vue3-vite'
import { setup } from '@storybook/vue3-vite'
import { createPinia } from 'pinia'
import vuetify from '../src/plugins/vuetify'

setup(app => {
  app.use(vuetify)
  app.use(createPinia())
})

const preview: Preview = {
  parameters: {
    a11y: { test: 'todo' },
  },
}

export default preview
