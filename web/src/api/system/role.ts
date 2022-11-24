import { http } from '@/utils/http/axios';

/**
 * @description: 角色列表
 */
export function getRoleList() {
  return http.request({
    url: '/role/list',
    method: 'GET',
  });
}

export function Edit(params) {
  return http.request({
    url: '/role/edit',
    method: 'POST',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/role/delete',
    method: 'POST',
    params,
  });
}

export function UpdatePermissions(params) {
  return http.request({
    url: '/role/updatePermissions',
    method: 'POST',
    params,
  });
}

export function GetPermissions(params) {
  return http.request({
    url: '/role/getPermissions',
    method: 'GET',
    params,
  });
}
