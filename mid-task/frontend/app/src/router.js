import Vue from 'vue';
import Router from 'vue-router';
import Books from './views/Books.vue';
import Search from './views/Search.vue';
import Tags from './views/Tags.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'books',
      component: Books
    },
    {
      path: '/search',
      name: 'search',
      component: Search
    },
    {
      path: '/tags',
      name: 'tags',
      component: Tags
    }
  ]
});
