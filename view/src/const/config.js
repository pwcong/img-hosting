export const IS_PROD = process.env.NODE_ENV == 'production';

export const BASE_API = IS_PROD ? window.location.host : 'http://127.0.0.1:7001';
