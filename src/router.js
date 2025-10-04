import { createMemoryHistory, createRouter } from 'vue-router'

import HomeView from './views/HomeView.vue'
// import RouterView from './views/RouterView.vue'

const routes = [
  { path: '/', component: HomeView },
  // { path: '/registration', component: RouterView },
]

export default createRouter({
  history: createMemoryHistory(),
  routes,
})
