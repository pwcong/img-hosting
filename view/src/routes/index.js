import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from '../pages/Home';
import Contact from '../pages/Contact';
import API from '../pages/API';
import About from '../pages/About';
import List from '../pages/List';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/contact',
    component: Contact
  },
  {
    path: '/api',
    component: API
  },
  {
    path: '/list',
    component: List
  },
  {
    path: '/about',
    component: About
  }
];

const router = new VueRouter({
  routes
});

export default router;
