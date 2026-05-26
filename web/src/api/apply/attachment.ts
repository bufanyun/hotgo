import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/attachment/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/attachment/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/attachment/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/attachment/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/attachment/view',
    method: 'GET',
    params,
  });
}

export function ClearKind(params) {
  return http.request({
    url: '/attachment/clearKind',
    method: 'POST',
    params,
  });
}

// 链接图片转存
export function ImageTransferStorage(params) {
  return http.request({
    url: '/upload/imageTransferStorage',
    method: 'POST',
    params,
  });
}
