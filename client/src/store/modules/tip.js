import axios from 'axios';

import * as api from '../../api.js';
import * as types from '../types.js';

const store = {

    state: {
        content: '',
        show: false,
        type: 0
    },
    getters: {
        
    },
    mutations: {

        [types.MUTATION_TIP_SHOWTIPS](state, payload){
            state.content = payload.content;
            state.type = payload.type;
            state.show = true;
            setTimeout(() => {
                state.show = false;
            }, 2000);
        },
        
    },
    actions: {
        

    }
    

}

export default store;