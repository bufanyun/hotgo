import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/cash/list',
    method: 'get',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/cash/edit',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/cash/view',
    method: 'GET',
    params,
  });
}

export function Apply(params) {
  return http.request({
    url: '/cash/apply',
    method: 'POST',
    params,
  });
}

export function Payment(params) {
  return http.request({
    url: '/cash/payment',
    method: 'POST',
    params,
  });
}
