import { http } from '@/utils/http/axios';

export function OnlineList(params) {
  return http.request({
    url: '/monitor/onlineList',
    method: 'get',
    params,
  });
}

export function Offline(params) {
  return http.request({
    url: '/monitor/offline',
    method: 'POST',
    params,
  });
}
