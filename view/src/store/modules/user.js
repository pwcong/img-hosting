import Cookies from 'js-cookie';

import {
  USER_ACTION_CHECK,
  USER_ACTION_LOGIN,
  USER_ACTION_REGISTER,
  USER_ACTION_LOGOUT,
  USER_MUTATION_SET_ID,
  USER_MUTATION_SET_TOKEN,
  USER_MUTATION_SET_USERNAME,
  USER_MUTATION_SET_CHECK
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
    [USER_MUTATION_SET_CHECK]: (state, payload) => {
      state.check = payload;
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
            const { id, token, username } = res.payload;

            commit(USER_MUTATION_SET_CHECK, true);
            commit(USER_MUTATION_SET_ID, id);
            commit(USER_MUTATION_SET_TOKEN, token);
            commit(USER_MUTATION_SET_USERNAME, username);

            Cookies.set('token', token);

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
            const { id, token, username } = res.payload;
            commit(USER_MUTATION_SET_CHECK, true);
            commit(USER_MUTATION_SET_ID, id);
            commit(USER_MUTATION_SET_TOKEN, token);
            commit(USER_MUTATION_SET_USERNAME, username);

            Cookies.set('token', token);
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
            const { id, token, username } = res.payload;
            commit(USER_MUTATION_SET_CHECK, true);
            commit(USER_MUTATION_SET_ID, id);
            commit(USER_MUTATION_SET_TOKEN, token);
            commit(USER_MUTATION_SET_USERNAME, username);

            Cookies.set('token', token);
            resolve(res);
          })
          .catch(err => {
            reject(err);
          });
      });
    },
    [USER_ACTION_LOGOUT]: ({ dispatch, commit, state }) => {
      return new Promise((resolve, reject) => {
        commit(USER_MUTATION_SET_CHECK, false);
        commit(USER_MUTATION_SET_ID, null);
        commit(USER_MUTATION_SET_TOKEN, null);
        commit(USER_MUTATION_SET_USERNAME, null);

        Cookies.remove('token');

        resolve();
      });
    }
  }
};

export default store;
