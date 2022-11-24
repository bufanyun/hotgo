import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/provinces/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/provinces/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/provinces/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/provinces/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/provinces/view',
    method: 'GET',
    params,
  });
}
