import request from '../request';

import { BASE_API } from '@/const/config';

export function getList(pageSize, pageNo) {
  return request(`${BASE_API}/imgs?page_size=${pageSize}&page_no=${pageNo}`, 'GET', {}, {});
}

export function removeImage(id) {
  return request(
    `${BASE_API}/img/remove`,
    'POST',
    {
      ids: [id]
    },
    {}
  );
}
