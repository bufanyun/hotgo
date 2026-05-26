import { http } from '@/utils/http/axios';

export function getLogList(params) {
  return http.request({
    url: '/smsLog/list',
    method: 'get',
    params,
  });
}
export function Delete(params) {
  return http.request({
    url: '/smsLog/delete',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/smsLog/view',
    method: 'GET',
    params,
  });
}
