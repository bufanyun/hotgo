import { http } from '@/utils/http/axios';

export function OnlineList(params) {
  return http.request({
    url: '/monitor/userOnlineList',
    method: 'get',
    params,
  });
}

export function Offline(params) {
  return http.request({
    url: '/monitor/userOffline',
    method: 'POST',
    params,
  });
}

export function NetOnlineList(params) {
  return http.request({
    url: '/monitor/netOnlineList',
    method: 'get',
    params,
  });
}

export function NetOffline(params) {
  return http.request({
    url: '/monitor/netOffline',
    method: 'POST',
    params,
  });
}

export function NetOption() {
  return http.request({
    url: '/monitor/netOption',
    method: 'get',
  });
}
