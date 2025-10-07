import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router.js'
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';

// Vueform
import Vueform from '@vueform/vueform/plugin'
import vueformConfig from './../vueform.config.js';

const pinia = createPinia();
const app = createApp(App);
app.config.globalProperties.window = window;

app.use(router);
app.use(pinia);
app.use(PrimeVue, {
  theme: {
    preset: Aura
  }
});
app.use(Vueform, vueformConfig);

app.mount('#app');
