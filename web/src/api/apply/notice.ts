import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/notice/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/notice/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/notice/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/notice/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/notice/view',
    method: 'GET',
    params,
  });
}
