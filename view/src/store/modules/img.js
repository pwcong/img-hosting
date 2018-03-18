import Cookies from 'js-cookie';

import {
  IMG_ACTION_UPLOAD,
  IMG_ACTION_REMOVE,
  IMG_ACTION_GETLIST,
  IMG_MUTATION_SET_IMGLIST
} from '../types';

import { getList } from '@/network/api/img';

const store = {
  state: {
    imgList: []
  },
  getters: {
    imgList: state => state.imgList
  },
  mutations: {
    [IMG_MUTATION_SET_IMGLIST]: (state, payload) => {
      state.imgList = payload;
    }
  },
  actions: {
    [IMG_ACTION_GETLIST]: ({ dispatch, commit, state }, payload) => {
      return new Promise((resolve, reject) => {
        const { pageSize, pageNo } = payload;

        getList(pageSize, pageNo)
          .then(res => {
            commit(IMG_MUTATION_SET_IMGLIST, res.payload.data || []);
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
