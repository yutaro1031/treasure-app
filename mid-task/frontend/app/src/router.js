import Vue from 'vue';
import Router from 'vue-router';
import Mixins from './views/Mixins.vue';
import Tables from './views/Tables.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'tables',
      component: Tables
    },
    {
      path: '/mixtest',
      name: 'mixtest',
      component: Mixins
    },
    {
      path: '/tables',
      name: 'tables',
      component: Tables
    }
  ]
});
