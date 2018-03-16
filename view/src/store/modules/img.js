import axios from 'axios';

import * as api from '../../api';
import * as types from '../types';

const store = {
  state: {
    imgCount: 0,
    imgList: [
      // {
      //     flag: 'img-1',
      //     file: '',
      //     img: '/imgs/test.jpg',
      //     uploading: false,
      //     uploaded: false,
      //     progress: 0,
      //     url: ''
      // },
    ],
    ownList: [
      // {
      //     "filename": "59432931_p0.jpg",
      //     "year": "2017",
      //     "month": "05",
      //     "day": "16",
      //     "storename": "b45b2db6425f.jpg",
      //     "path": "/2017/05/16/b45b2db6425f.jpg",
      //     "url": "http://localhost/public/2017/05/16/b45b2db6425f.jpg"
      // }
    ]
  },
  getters: {
    imgList: state => state.imgList,
    imgListLength: state => state.imgList.length,

    ownList: state => state.ownList,
    ownListLength: state => state.ownList.length,

    canUploadAll: state => {
      for (let i = 0; i < state.imgList.length; i++) {
        if (!state.imgList[i].uploading && !state.imgList[i].uploaded) return true;
      }

      return false;
    }
  },
  mutations: {
    [types.MUTATION_IMG_PUSHIMG](state, payload) {
      state.imgList.push({
        flag: 'img-' + state.imgCount,
        file: payload.file,
        img: window.URL.createObjectURL(payload.file),
        uploading: false,
        uploaded: false,
        progress: 0,
        url: ''
      });

      state.imgCount++;
    },
    [types.MUTATION_IMG_REMOVEIMG](state, payload) {
      state.imgList.forEach((img, index) => {
        if (img.flag == payload.flag) {
          state.imgList.splice(index, 1);
          return;
        }
      });
    },
    [types.MUTATION_IMG_REMOVEOWNIMG](state, payload) {
      state.ownList.forEach((img, index) => {
        if (
          img.year == payload.year &&
          img.month == payload.month &&
          img.day == payload.day &&
          img.storename == payload.storename
        ) {
          state.ownList.splice(index, 1);
          return;
        }
      });
    },
    [types.MUTATION_IMG_SETOWNIMGS](state, payload) {
      if (payload.imgs) state.ownList = payload.imgs;
      else state.ownList = [];
    }
  },
  actions: {
    [types.ACTION_IMG_TOUPLOADIMG]({ commit, state }, payload) {
      state.imgList.forEach(img => {
        if (img.flag == payload.flag) {
          let formData = new FormData();
          formData.append('img', img.file);
          formData.append('uid', payload.uid);

          img.uploading = true;

          axios
            .request({
              url: api.URL_API_IMG_UPLOAD,
              method: 'post',
              data: formData,
              onUploadProgress: function(progressEvent) {
                let loaded = progressEvent.loaded;
                let total = progressEvent.total;

                img.progress = 100 * loaded / total;
              }
            })
            .then(res => {
              img.url = res.data.url;
              img.uploaded = true;
              img.uploading = false;
            })
            .catch(err => {
              img.uploading = false;
            });
        }
      });
    },
    [types.ACTION_IMG_TOUPLOADALLIMG]({ dispatch, commit, state }, payload) {
      state.imgList.forEach(img => {
        if (!img.uploaded && !img.uploading) {
          dispatch(types.ACTION_IMG_TOUPLOADIMG, {
            flag: img.flag,
            uid: payload.uid
          });
        }
      });
    },
    [types.ACTION_IMG_GETLIST]({ commit, state }, payload) {
      let formData = new FormData();
      formData.append('token', payload.token);

      axios
        .request({
          url: api.URL_API_IMG_LIST,
          method: 'post',
          data: formData
        })
        .then(res => {
          let imgs = res.data.imgs;

          commit(types.MUTATION_IMG_SETOWNIMGS, {
            imgs
          });
        })
        .catch(err => {});
    },
    [types.ACTION_IMG_REMOVEOWN]({ commit, state }, payload) {
      let formData = new FormData();
      formData.append('token', payload.token);
      formData.append('year', payload.year);
      formData.append('month', payload.month);
      formData.append('day', payload.day);
      formData.append('storename', payload.storename);

      axios
        .request({
          url: api.URL_API_IMG_DELETE,
          method: 'post',
          data: formData
        })
        .then(res => {
          commit(types.MUTATION_IMG_REMOVEOWNIMG, {
            year: payload.year,
            month: payload.month,
            day: payload.day,
            storename: payload.storename
          });
        })
        .catch(err => {});
    }
  }
};

export default store;
