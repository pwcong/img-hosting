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
            //     progress: 0
            // },
        ]
    },
    getters: {

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
                progress: 0
            });
        },
        [types.MUTATION_IMG_REMOVEIMG](state, payload){
            imgList.forEach((img, index) => {
                if(img.key == payload.key){
                    imgList.splice(index, 1);
                    return;
                }
            });
        }
        
    },
    actions: {

        [types.ACTION_IMG_TOUPLOADIMG]({commit, state}, key){
            

        }


    }
    

}

export default store;