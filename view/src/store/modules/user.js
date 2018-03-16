import axios from 'axios';

import * as api from '../../api';
import * as types from '../types';

const store = {
  state: {
    uid: '',
    token: '',
    logined: false,
    logining: false,
    registering: false
  },
  getters: {},
  mutations: {
    [types.MUTATION_USER_LOGIN_START](state) {
      state.logined = false;
      state.logining = true;
      state.registering = false;
      state.uid = '';
      state.token = '';

      localStorage.token = '';
    },
    [types.MUTATION_USER_LOGIN_SUCCESS](state, payload) {
      state.logined = true;
      state.logining = false;
      state.registering = false;
      state.uid = payload.uid;
      state.token = payload.token;

      localStorage.token = payload.token;
    },
    [types.MUTATION_USER_LOGIN_FAILED](state) {
      state.logined = false;
      state.logining = false;
      state.registering = false;
      state.uid = '';
      state.token = '';

      localStorage.token = '';
    },
    [types.MUTATION_USER_REGISTER_START](state, payload) {
      state.logined = false;
      state.logining = false;
      state.registering = true;
      state.uid = '';
      state.token = '';

      localStorage.token = '';
    },
    [types.MUTATION_USER_REGISTER_FAILED](state) {
      state.logined = false;
      state.logining = false;
      state.registering = false;
      state.uid = '';
      state.token = '';

      localStorage.token = '';
    },
    [types.MUTATION_USER_LOGOUT](state) {
      state.logined = false;
      state.logining = false;
      state.registering = false;
      state.uid = '';
      state.token = '';

      localStorage.token = '';
    }
  },
  actions: {
    [types.ACTION_USER_TOLOGIN]({ commit }, payload) {
      return new Promise((resolve, reject) => {
        let formData = new FormData();
        formData.append('uid', payload.uid);
        formData.append('pwd', payload.pwd);

        commit(types.MUTATION_USER_LOGIN_START);

        axios
          .request({
            url: api.URL_API_USER_LOGIN,
            method: 'post',
            data: formData
          })
          .then(res => {
            let token = res.data.token;

            commit(types.MUTATION_USER_LOGIN_SUCCESS, {
              uid: payload.uid,
              token: token
            });

            resolve();
          })
          .catch(err => {
            commit(types.MUTATION_USER_LOGIN_FAILED);

            reject();
          });
      });
    },
    [types.ACTION_USER_CHECK]({ commit }) {
      let token = localStorage.token;

      if (token && token != '') {
        let formData = new FormData();
        formData.append('token', token);

        axios
          .request({
            url: api.URL_API_USER_CHECK,
            method: 'post',
            data: formData
          })
          .then(res => {
            let uid = res.data.uid;

            commit(types.MUTATION_USER_LOGIN_SUCCESS, {
              uid: uid,
              token: token
            });
          })
          .catch(err => {
            commit(types.MUTATION_USER_LOGIN_FAILED);
          });
      } else {
        commit(types.MUTATION_USER_LOGIN_FAILED);
      }
    },
    [types.ACTION_USER_TOLOGOUT]({ commit }) {
      commit(types.MUTATION_USER_LOGOUT);
    },
    [types.ACTION_USER_TOREGISTER]({ commit }, payload) {
      return new Promise((resolve, reject) => {
        let formData = new FormData();
        formData.append('uid', payload.uid);
        formData.append('pwd', payload.pwd);

        commit(types.MUTATION_USER_REGISTER_START);

        axios
          .request({
            url: api.URL_API_USER_REGISTER,
            method: 'post',
            data: formData
          })
          .then(res => {
            let token = res.data.token;

            commit(types.MUTATION_USER_LOGIN_SUCCESS, {
              uid: payload.uid,
              token: token
            });

            resolve();
          })
          .catch(err => {
            commit(types.MUTATION_USER_REGISTER_FAILED);

            reject();
          });
      });
    }
  }
};

export default store;
