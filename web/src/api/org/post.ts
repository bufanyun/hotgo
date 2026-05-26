import { http } from '@/utils/http/axios';

export function getPostList(params?) {
  return http.request({
    url: '/post/list',
    method: 'GET',
    params,
  });
}

export function getPostOption(params?) {
  return getPostList(params);
}

export function Edit(params) {
  return http.request({
    url: '/post/edit',
    method: 'POST',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/post/delete',
    method: 'POST',
    params,
  });
}
