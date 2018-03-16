import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

import Home from '../pages/Home';
import About from '../pages/About';
import Contact from '../pages/Contact';
import API from '../pages/API';
import Own from '../pages/Own';

const routes = [
  {
    path: '/',
    component: Home
  },
  {
    path: '/about',
    component: About
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
    path: '/own',
    component: Own
  }
];

const router = new VueRouter({
  routes
});

export default router;
