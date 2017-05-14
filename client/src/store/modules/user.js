import * as types from '../types.js';

const store = {

    state: {
        uid: '',
        token: '',
    },
    getters: {
        
    },
    mutations: {
        [types.MUTATION_USER_SETUID](state, payload){
            state.uid = payload.uid;
        },
        [types.MUTATION_USER_SETTOKEN](state, payload){
            state.token = payload.token;
        }
    },
    actions: {

        [types.ACTION_USER_TOLOGIN]({ commit }, user){
            


        },
        [types.ACTION_USER_TOREGISTER]({ commit }, user){


        }


    }
    

}

export default store;