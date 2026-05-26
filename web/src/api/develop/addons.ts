import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/addons/list',
    method: 'get',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/addons/view',
    method: 'GET',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/addons/status',
    method: 'POST',
    params,
  });
}

export function Build(params) {
  return http.request({
    url: '/addons/build',
    method: 'post',
    params,
  });
}

export function Install(params) {
  return http.request({
    url: '/addons/install',
    method: 'post',
    params,
  });
}

export function Upgrade(params) {
  return http.request({
    url: '/addons/upgrade',
    method: 'post',
    params,
  });
}

export function UnInstall(params) {
  return http.request({
    url: '/addons/uninstall',
    method: 'post',
    params,
  });
}
