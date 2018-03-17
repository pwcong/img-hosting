import Vue from 'vue';

import 'normalize.css';
import './assets/css/github-markdown.css';

import store from './store';
import router from './routes';

import App from './pages/App';

new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App)
});
