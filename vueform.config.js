import { defineConfig } from '@vueform/vueform'
import vueformTheme from '@vueform/vueform/themes/vueform'
import en from '@vueform/vueform/locales/en'
import '@vueform/vueform/dist/vueform.css'

export default defineConfig({
    theme: vueformTheme,
    locales: { en },
    locale: 'en',
    showRequired: ['label'],
});