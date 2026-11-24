import { createMemoryHistory, createRouter, createWebHistory } from 'vue-router'

import HomeView from './views/HomeView.vue'
import RegisterView from './views/RegisterView.vue'
import PageNotFoundView from './views/PageNotFoundView.vue'

const routes = [
  {
    path: '/',
    component: HomeView
  },
  {
    path: '/registration',
    component: RegisterView
  },
  {
    path: '/:pathMatch(.*)*',
    component: PageNotFoundView,
  }
]

export default createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }

    const options = { top: 0 };

    if (to.hash !== '') {
      options.top = 80;          // apply offset so it's not covered by the nav
      options.el = to.hash;
    }

    return options;
  }
});