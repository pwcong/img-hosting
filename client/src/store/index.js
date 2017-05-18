import Vue from 'vue';
import Vuex from 'vuex';

import img from './modules/img.js';
import user from './modules/user.js';
import tip from './modules/tip.js';

Vue.use(Vuex);

const store = new Vuex.Store({
    modules: {
        img,
        user,
        tip
    }
});


export default store;