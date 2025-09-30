import { createMemoryHistory, createRouter } from 'vue-router'

import HomeView from './views/HomeView.vue'
// import RouterView from './views/RouterView.vue'

const routes = [
  {
    path: '/',
    component: HomeView
  },
  {
    path: '/registration',
    component: RegisterView
  },
]

export default createRouter({
  history: createMemoryHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return {
        top: 0,
      }
    }
  }
})
