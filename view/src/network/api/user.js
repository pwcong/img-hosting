import request from '../request';

import { BASE_API } from '@/const/config';

function login(username, password) {
  return request(
    `${BASE_API}/user/login`,
    'POST',
    {
      username,
      password
    },
    {}
  );
}

function register(username, password) {
  return request(
    `${BASE_API}/user/register`,
    'POST',
    {
      username,
      password
    },
    {}
  );
}

function check() {
  return request(`${BASE_API}/user/check`, 'POST', {}, {});
}
