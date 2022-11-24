import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/blacklist/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/blacklist/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/blacklist/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/blacklist/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/blacklist/view',
    method: 'GET',
    params,
  });
}
