const isProd = process.env.NODE_ENV == 'production';
const urlPrefix = isProd ? 'http://' + window.location.host : 'http://localhost';

export default {

    URL_API_IMG_UPLOAD: urlPrefix + '/img/upload',
    URL_API_USER_REGISTER: urlPrefix + "/user/register",
	URL_API_USER_LOGIN: urlPrefix + "/user/login",
	URL_API_USER_UPDATE: urlPrefix +  "/user/update"
}