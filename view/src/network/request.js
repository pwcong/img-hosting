import axios from 'axios';

import { Message } from 'element-ui';

import Cookies from 'js-cookie';

const HEADERS = {};

export default function(url, method, data, headers, options) {
  return new Promise((resolve, reject) => {
    axios(
      Object.assign(
        {},
        {
          url,
          method,
          data,
          headers: Object.assign(
            {
              Token: Cookies.get('token')
            },
            HEADERS,
            headers
          )
        },
        options
      )
    )
      .then(res => {
        if (res.status === 200 && res.data && res.data.code === 20000) {
          resolve(res.data);
        } else {
          Message({
            message: res.data ? res.data.message : '请求失败',
            duration: 1500,
            type: 'error'
          });
          reject(res);
        }
      })
      .catch(err => {
        reject(err);
      });
  });
}
