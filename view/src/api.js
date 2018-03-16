const isProd = process.env.NODE_ENV == 'production';
const urlPrefix = isProd ? 'http://127.0.0.1:7001' : '';

export const URL_API_IMG_UPLOAD = urlPrefix + '/img/upload';
export const URL_API_IMG_UPLOAD_PRIVATE = urlPrefix + '/img/upload/private';
export const URL_API_IMG_DELETE = urlPrefix + '/img/delete';
export const URL_API_IMG_LIST = urlPrefix + '/imgs';

export const URL_API_USER_REGISTER = urlPrefix + '/user/register';
export const URL_API_USER_LOGIN = urlPrefix + '/user/login';
export const URL_API_USER_CHECK = urlPrefix + '/user/check';
