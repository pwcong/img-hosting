import axios from 'axios';

const HEADERS = {

}

export default function (url, method, data, headers) {

    return new Promise((resolve, reject) => {

        axios({
            url,
            method,
            data,
            headers: Object.assign({}, HEADERS, headers)
        }).then(res => {

            if(res.status === 200){
                resolve(res);
            }else{
                reject(res);
            }

        }).catch(err => {
            reject(err);
        });


    });


}