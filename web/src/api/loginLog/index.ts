import { http, jumpExport } from '@/utils/http/axios';

// 获取登录日志列表
export function List(params) {
  return http.request({
    url: '/loginLog/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除登录日志
export function Delete(params) {
  return http.request({
    url: '/loginLog/delete',
    method: 'POST',
    params,
  });
}

// 导出登录日志
export function Export(params) {
  jumpExport('/loginLog/export', params);
}
