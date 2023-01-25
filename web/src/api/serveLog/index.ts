import { http, jumpExport } from '@/utils/http/axios';

// 获取服务日志列表
export function List(params) {
  return http.request({
    url: '/serveLog/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除服务日志
export function Delete(params) {
  return http.request({
    url: '/serveLog/delete',
    method: 'POST',
    params,
  });
}

// 获取服务日志指定详情
export function View(params) {
  return http.request({
    url: '/serveLog/view',
    method: 'GET',
    params,
  });
}

// 导出服务日志
export function Export(params) {
  jumpExport('/serveLog/export', params);
}
