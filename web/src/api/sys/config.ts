import { http } from '@/utils/http/axios';

export function getConfig(params) {
  return http.request({
    url: '/config/get',
    method: 'get',
    params,
  });
}
export function updateConfig(params) {
  return http.request({
    url: '/config/update',
    method: 'post',
    params,
  });
}

export function sendTestEmail(params) {
  return http.request({
    url: '/ems/sendTest',
    method: 'post',
    params,
  });
}
