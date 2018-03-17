import Vue from 'vue';
import Vuex from 'vuex';

import sampleModule from './modules/sample';

Vue.use(Vuex);

const store = new Vuex.Store({

    modules: {

        sample: sampleModule

    }

});

export default store;