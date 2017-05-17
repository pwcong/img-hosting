import axios from 'axios';

import * as api from '../../api.js';
import * as types from '../types.js';

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

        ownList: state => state.ownList,
        ownListLenght: state => state.ownList.length,

        imgList: state => state.imgList,
        imgListLength: state => state.imgList.length,

        canUploadAll: state => {
            for(let i = 0; i < state.imgList.length; i++){
                if(!state.imgList[i].uploading && !state.imgList[i].uploaded)
                    return true;
            }

            return false;
        }

    },
    mutations: {

        [types.MUTATION_IMG_PUSHIMG](state, payload){
            state.imgList.push({
                flag: "img-" + state.imgCount,
                file: payload.file,
                img: window.URL.createObjectURL(payload.file),
                uploading: false,
                uploaded: false,
                progress: 0,
                url: ''
            });

            state.imgCount++;
        },
        [types.MUTATION_IMG_REMOVEIMG](state, payload){
            state.imgList.forEach((img, index) => {
                if(img.flag == payload.flag){
                    state.imgList.splice(index, 1);
                    return;
                }
            });
        },
        [types.MUTATION_IMG_SETOWNIMGS](state, payload){
            state.ownList = payload.imgs;
        },
        
    },
    actions: {

        [types.ACTION_IMG_TOUPLOADIMG]({commit, state}, payload){

            state.imgList.forEach( img => {

                if(img.flag == payload.flag){

                    let formData = new FormData();
                    formData.append("img", img.file);
                    formData.append("uid", payload.uid);

                    img.uploading = true;
            

                    axios.request({
                        url: api.URL_API_IMG_UPLOAD,
                        method: 'post',
                        data: formData,
                        onUploadProgress: function(progressEvent){
                            let loaded = progressEvent.loaded;
                            let total = progressEvent.total;

                            img.progress = 100 * loaded / total;

                        }

                    }).then( res => {
                        img.url = res.data.url;
                        img.uploaded = true;
                        img.uploading = false;
                    }).catch( err => {
                        img.uploading = false;
                    });

                }   

            });
            

        },
        [types.ACTION_IMG_TOUPLOADALLIMG]({dispatch, commit, state}, payload){
            
            state.imgList.forEach( img => {

                if(!img.uploaded && !img.uploading){
                    dispatch(types.ACTION_IMG_TOUPLOADIMG, {
                        flag: img.flag,
                        uid: payload.uid
                    });
                }

            });

        },
        [types.ACTION_IMG_GETLIST]({commit, state}, payload){

            let formData = new FormData();
            formData.append('token', payload.token);

            axios.request({
                url: api.URL_API_IMG_LIST,
                method: 'post',
                data: formData

            }).then( res => {
                let imgs = res.data.imgs;

                commit(types.MUTATION_IMG_SETOWNIMGS, {
                    imgs
                });

            }).catch( err => {
                
            });



        }


    }
    

}

export default store;