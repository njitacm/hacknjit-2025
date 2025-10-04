
import { createApp } from 'vue'
import router from './router.js'
import App from './App.vue'
import PrimeVue from 'primevue/config';

// Vueform
import Vueform from '@vueform/vueform/plugin'
import vueformTheme from '@vueform/vueform/themes/vueform'
import en from '@vueform/vueform/locales/en'

// Vueform CSS
import '@vueform/vueform/dist/vueform.css'

const app = createApp(App);

app.use(PrimeVue);
app.use(router);
app.use(Vueform, {
  theme: vueformTheme,
  locales: { en },
  locale: 'en'
});

app.mount('#app');