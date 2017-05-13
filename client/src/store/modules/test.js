import * as types from '../types.js';

const store = {

    state: {
        name: 'Test'
    },
    getters: {
        name: state => state.name
    },
    mutations: {
        [types.TEST_SETNAME](state, payload){
            state.name = payload.name;
        }
    },
    actions: {

        asyncSetName({commit, state}, name){

            return new Promise((resolve, reject) => {

                setTimeout(() => {

                    commit(types.TEST_SETNAME, {
                        name
                    });

                    resolve();

                }, 2000);

            });

        }


    }
    

}

export default store;