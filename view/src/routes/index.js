import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from '../pages/Home.vue';
import Counter from '../pages/Counter.vue';

Vue.use(VueRouter);

const routes = [

    {
        path: '/',
        component: Home
    },
    {
        path: '/counter',
        component: Counter
    }

]

const router = new VueRouter({
    routes
});

export default router;