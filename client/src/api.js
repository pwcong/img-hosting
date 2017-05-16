const isProd = process.env.NODE_ENV == 'production';
const urlPrefix = isProd ? 'http://' + window.location.host : 'http://localhost';

export const URL_API_IMG_UPLOAD = urlPrefix + '/img/upload';

export const URL_API_USER_REGISTER =  urlPrefix + "/user/register";
export const URL_API_USER_LOGIN = urlPrefix + "/user/login";
