import Vue from 'vue';
import Router from 'vue-router';
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
  ]
});
