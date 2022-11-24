import { http } from '@/utils/http/axios';

export function getDeptList(params?) {
  return http.request({
    url: '/dept/list',
    method: 'GET',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/dept/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/dept/status',
    method: 'POST',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/dept/delete',
    method: 'POST',
    params,
  });
}
