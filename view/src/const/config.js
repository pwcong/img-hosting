const IS_PROD = process.env.NODE_ENV == 'production';

const BASE_API = IS_PROD ? '' : 'http://127.0.0.1:7001';

export default {
  IS_PROD,
  BASE_API
};
