import { http } from '@/utils/http/axios';

export function getLogList(params) {
  return http.request({
    url: '/log/list',
    method: 'get',
    params,
  });
}
export function Delete(params) {
  return http.request({
    url: '/log/delete',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/log/view',
    method: 'GET',
    params,
  });
}
