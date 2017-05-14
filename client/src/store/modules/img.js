import * as types from '../types.js';

const store = {

    state: {
        imgList: [
            {
                flag: 'img-1',
                file: '/imgs/test.jpg',
                uploading: false,
                uploaded: false,
                progress: 0
            },
            {
                flag: 'img-1',
                file: '/imgs/test.jpg',
                uploading: false,
                uploaded: false,
                progress: 0
            }
        ]
    },
    getters: {

        imgList: state => state.imgList,
        imgListLength: state => state.imgList.length,

        canUploadAll: state => {
            for(let i = 0; i < imgList.length; i++){
                if(!imgList[i].uploading && !imgList[i].uploaded)
                    return true;
            }

            return false;
        }

    },
    mutations: {

        [types.MUTATION_IMG_PUSHIMG](state, payload){
            imgList.push(payload.img);
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