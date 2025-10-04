import { createMemoryHistory, createRouter } from 'vue-router'

import HomeView from './views/HomeView.vue'

const routes = [
  { path: '/', component: HomeView },
]

export default createRouter({
  history: createMemoryHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }

    const options = { top: 0 };

    if (to.hash !== '') {
      options.top = 100;          // apply offset so it's not covered by the nav
      options.el = to.hash;
    }

    return options;
  }
})
