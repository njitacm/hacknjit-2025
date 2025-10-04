
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router.js'
import PrimeVue from 'primevue/config';

// Vueform
import Vueform from '@vueform/vueform/plugin'
import vueformTheme from '@vueform/vueform/themes/vueform'
import en from '@vueform/vueform/locales/en'

// Vueform CSS
import '@vueform/vueform/dist/vueform.css'

const pinia = createPinia();
const app = createApp(App);
app.config.globalProperties.window = window;

app.use(router);
app.use(pinia);
app.use(PrimeVue);
app.use(Vueform, {
  theme: vueformTheme,
  locales: { en },
  locale: 'en'
});

app.mount('#app');