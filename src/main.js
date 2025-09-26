import { createApp } from 'vue'
import App from './App.vue'
<<<<<<< HEAD

var app = createApp(App);
app.config.globalProperties.window = window
=======
import PrimeVue from 'primevue/config';
// Vueform
import Vueform from '@vueform/vueform/plugin'
import vueformTheme from '@vueform/vueform/themes/vueform'
import en from '@vueform/vueform/locales/en'

// Vueform CSS
import '@vueform/vueform/dist/vueform.css'

const app = createApp(App)

app.use(PrimeVue)
app.use(Vueform, {
  theme: vueformTheme,
  locales: { en },
  locale: 'en'
})
>>>>>>> 202e8af (Added primevue)

app.mount('#app')
