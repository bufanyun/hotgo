import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/notice/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/notice/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/notice/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/notice/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/notice/view',
    method: 'GET',
    params,
  });
}

// 获取最大排序
export function MaxSort() {
  return http.request({
    url: '/notice/maxSort',
    method: 'GET',
  });
}

export function EditNotify(params) {
  return http.request({
    url: '/notice/editNotify',
    method: 'POST',
    params,
  });
}

export function EditNotice(params) {
  return http.request({
    url: '/notice/editNotice',
    method: 'POST',
    params,
  });
}

export function EditLetter(params) {
  return http.request({
    url: '/notice/editLetter',
    method: 'POST',
    params,
  });
}

export function ReadAll(params) {
  return http.request({
    url: '/notice/readAll',
    method: 'POST',
    params,
  });
}

export function PullMessages() {
  return http.request({
    url: '/notice/pullMessages',
    method: 'get',
  });
}

export function UpRead(params) {
  return http.request({
    url: '/notice/upRead',
    method: 'POST',
    params,
  });
}

export function MessageList(params) {
  return http.request({
    url: '/notice/messageList',
    method: 'get',
    params,
  });
}
