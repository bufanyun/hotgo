import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/cron/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/cron/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/cron/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/cron/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/cron/view',
    method: 'GET',
    params,
  });
}

export function GroupList(params) {
  return http.request({
    url: '/cronGroup/list',
    method: 'get',
    params,
  });
}

export function GroupDelete(params) {
  return http.request({
    url: '/cronGroup/delete',
    method: 'POST',
    params,
  });
}

export function GroupEdit(params) {
  return http.request({
    url: '/cronGroup/edit',
    method: 'POST',
    params,
  });
}

export function GroupStatus(params) {
  return http.request({
    url: '/cronGroup/status',
    method: 'POST',
    params,
  });
}

export function GroupView(params) {
  return http.request({
    url: '/cronGroup/view',
    method: 'GET',
    params,
  });
}

export function getSelect(params) {
  return http.request({
    url: '/cronGroup/select',
    method: 'GET',
    params,
  });
}

export function OnlineExec(params) {
  return http.request({
    url: '/cron/onlineExec',
    method: 'POST',
    params,
  });
}
