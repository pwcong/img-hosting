import Vue from 'vue';
import Vuex from 'vuex';

import img from './modules/img';
import user from './modules/user';
import tip from './modules/tip';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    img,
    user,
    tip
  }
});

export default store;
