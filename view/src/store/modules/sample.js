import types from '../types';

const store = {

    state: {
        counts: 0
    },
    getters: {
        doubleCounts: state => 2 * state.counts
    },
    mutations: {
        [types.SAMPLE_MUTATIONS_PLUS]: (state, payload) => {
            state.counts++;
        }
    },
    actions: {

        [types.SAMPLE_ACTIONS_PLUS]: ({
            commit,
            state
        }) => {

            commit(types.SAMPLE_MUTATIONS_PLUS);

        },
        [types.SAMPLE_ACTIONS_PLUS_ASYNC]: ({
            dispatch,
            commit,
            state
        }, payload) => {

            return new Promise((resolve, reject) => {

                setTimeout(() => {

                    dispatch(types.SAMPLE_ACTIONS_PLUS);

                }, payload.time)

            });

        }

    }


}

export default store;