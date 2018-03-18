import Vue from 'vue';
import Vuex from 'vuex';

import user from './modules/user';
import img from './modules/img';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    user,
    img
  }
});

export default store;
