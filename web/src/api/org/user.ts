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
    url: '/member/resetPwd',
    method: 'POST',
    params,
  });
}

// 获取可选的后台用户选项
export function GetMemberOption() {
  return http.request({
    url: '/member/option',
    method: 'GET',
  });
}
