import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Vueform from '@vueform/vueform/vite'

export default defineConfig({
  floatPlaceholders: false,
  plugins: [
    vue(),
    Vueform(),   // only the vite plugin
  ],
})

