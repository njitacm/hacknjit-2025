import { createApp } from 'vue'
import App from './App.vue'
import router from './router.js'

// Vueform
import Vueform from '@vueform/vueform/plugin'
import vueformTheme from '@vueform/vueform/themes/vueform'
import en from '@vueform/vueform/locales/en'

// Vueform CSS
import '@vueform/vueform/dist/vueform.css'

const app = createApp(App)


app.config.globalProperties.window = window

app.use(Vueform, {
  theme: vueformTheme,
  locales: { en },
  locale: 'en'
})

app.use(router)
app.mount('#app')

