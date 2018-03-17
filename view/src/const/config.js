export const IS_PROD = process.env.NODE_ENV == 'production';

export const BASE_API = IS_PROD ? '' : 'http://127.0.0.1:7001';
