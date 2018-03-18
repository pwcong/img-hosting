import Cookies from 'js-cookie';

import request from '@/network/request';
import { BASE_API } from '@/const/config';

import {
  IMG_ACTION_UPLOAD,
  IMG_ACTION_REMOVE,
  IMG_ACTION_GETLIST,
  IMG_MUTATION_SET_IMGLIST,
  IMG_MUTATION_PUSH_IMG,
  IMG_MUTATION_POP_IMG,
  IMG_MUTATION_UPLOAD_START,
  IMG_MUTATION_UPLOAD_SUCCESS,
  IMG_MUTATION_UPLOAD_FAILED,
  IMG_MUTATION_UPLOAD_DOING
} from '../types';

import { getList } from '@/network/api/img';

const store = {
  state: {
    imgList: [],
    uploadList: [
      // {
      //   symbol: '',
      //   img: null,
      //   uploading: false,
      //   uploaded: false,
      //   success: true,
      //   url: ''
      // }
    ]
  },
  getters: {
    imgList: state => state.imgList,
    uploadList: state => state.uploadList
  },
  mutations: {
    [IMG_MUTATION_SET_IMGLIST]: (state, payload) => {
      state.imgList = payload;
    },
    [IMG_MUTATION_PUSH_IMG]: (state, payload) => {
      state.uploadList.push(payload);
    },
    [IMG_MUTATION_POP_IMG]: (state, payload) => {
      state.uploadList = state.uploadList.filter(img => img.symbol !== payload.symbol);
    },
    [IMG_MUTATION_UPLOAD_START]: (state, payload) => {
      state.uploadList.forEach(t => {
        if (t.symbol === payload.symbol) {
          t.uploading = true;
          t.uploaded = false;
          t.success = true;
          t.url = '';
          t.progress = 0;
        }
      });
    },
    [IMG_MUTATION_UPLOAD_SUCCESS]: (state, payload) => {
      state.uploadList.forEach(t => {
        if (t.symbol === payload.symbol) {
          t.uploading = false;
          t.uploaded = true;
          t.success = true;
          t.url = payload.url;
          t.progress = 100;
        }
      });
    },
    [IMG_MUTATION_UPLOAD_FAILED]: (state, payload) => {
      state.uploadList.forEach(t => {
        if (t.symbol === payload.symbol) {
          t.uploading = false;
          t.uploaded = false;
          t.success = false;
          t.url = '';
          t.progress = 0;
        }
      });
    },
    [IMG_MUTATION_UPLOAD_DOING]: (state, payload) => {
      state.uploadList.forEach(t => {
        if (t.symbol === payload.symbol) {
          t.progress = payload.progress;
        }
      });
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
    },
    [IMG_ACTION_UPLOAD]: ({ dispatch, commit, state }, payload) => {
      return new Promise((resolve, reject) => {
        var list = [];

        if (payload) {
          list.push(payload);
        } else {
          list = state.uploadList;
        }

        list.filter(t => !t.uploaded && !t.uploading).forEach(img => {
          const formData = new FormData();
          formData.append('img', img.file);

          commit(IMG_MUTATION_UPLOAD_START, img);

          request(
            `${BASE_API}/img/upload`,
            'POST',
            formData,
            {},
            {
              onUploadProgress: function(progressEvent) {
                const loaded = progressEvent.loaded;
                const total = progressEvent.total;

                commit(IMG_MUTATION_UPLOAD_DOING, {
                  ...img,
                  progress: 100 * loaded / total
                });
              }
            }
          )
            .then(res => {
              commit(IMG_MUTATION_UPLOAD_SUCCESS, {
                ...img,
                url: BASE_API + res.payload.url
              });
            })
            .catch(err => {
              commit(IMG_MUTATION_UPLOAD_FAILED, img);
            });
        });

        resolve();
      });
    }
  }
};

export default store;
