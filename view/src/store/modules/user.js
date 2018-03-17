import {
  USER_ACTION_CHECK,
  USER_ACTION_LOGIN,
  USER_ACTION_REGISTER,
  USER_MUTATION_SET_ID,
  USER_MUTATION_SET_TOKEN,
  USER_MUTATION_SET_USERNAME
} from '../types';

import { login, register, check } from '@/network/api/user';

const store = {
  state: {
    check: false,
    id: null,
    username: null,
    token: null
  },
  getters: {
    check: state => state.check,
    username: state => state.username
  },
  mutations: {
    [USER_MUTATION_SET_ID]: (state, payload) => {
      state.id = payload;
    },
    [USER_MUTATION_SET_TOKEN]: (state, payload) => {
      state.token = payload;
    },
    [USER_MUTATION_SET_USERNAME]: (state, payload) => {
      state.username = payload;
    }
  },
  actions: {
    [USER_ACTION_LOGIN]: ({ dispatch, commit, state }, payload) => {
      return new Promise((resolve, reject) => {
        const { username, password } = payload;

        login(username, password)
          .then(res => {
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    },
    [USER_ACTION_REGISTER]: ({ dispatch, commit, state }, payload) => {
      return new Promise((resolve, reject) => {
        const { username, password } = payload;

        register(username, password)
          .then(res => {
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    },
    [USER_ACTION_CHECK]: ({ dispatch, commit, state }, payload) => {
      return new Promise((resolve, reject) => {
        check()
          .then(res => {
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    }
  }
};

export default store;
