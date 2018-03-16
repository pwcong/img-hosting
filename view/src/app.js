import Vue from 'vue';

import App from './pages/App';

import router from './routes';
import store from './store';

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
});