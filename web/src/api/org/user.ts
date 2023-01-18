import { http } from '@/utils/http/axios';

export function List(params?) {
  return http.request({
    url: '/member/list',
    method: 'GET',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/member/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/member/status',
    method: 'POST',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/member/delete',
    method: 'POST',
    params,
  });
}

export function ResetPwd(params) {
  return http.request({
    url: '/member/reset_pwd',
    method: 'POST',
    params,
  });
}
