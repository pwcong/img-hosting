import Vue from 'vue';
import Vuex from 'vuex';

import img from './modules/img.js';
import user from './modules/user.js';

Vue.use(Vuex);

const store = new Vuex.Store({
    modules: {
        img,
        user
    }
});


export default store;