import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router.js'

const pinia = createPinia()
const app = createApp(App)
app.use(pinia)

app.config.globalProperties.window = window

app.use(router)
app.mount('#app')
