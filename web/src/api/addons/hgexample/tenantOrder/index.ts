import { http, jumpExport } from '@/utils/http/axios';

// 获取多租户功能演示列表
export function List(params) {
  return http.request({
    url: '/hgexample/tenantOrder/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除多租户功能演示
export function Delete(params) {
  return http.request({
    url: '/hgexample/tenantOrder/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑多租户功能演示
export function Edit(params) {
  return http.request({
    url: '/hgexample/tenantOrder/edit',
    method: 'POST',
    params,
  });
}

// 获取多租户功能演示指定详情
export function View(params) {
  return http.request({
    url: '/hgexample/tenantOrder/view',
    method: 'GET',
    params,
  });
}

// 导出多租户功能演示
export function Export(params) {
  jumpExport('/hgexample/tenantOrder/export', params);
}