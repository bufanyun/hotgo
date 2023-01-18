import { http, jumpExport } from '@/utils/http/axios';

// 获取生成演示列表
export function List(params) {
  return http.request({
    url: '/curdDemo/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除生成演示
export function Delete(params) {
  return http.request({
    url: '/curdDemo/delete',
    method: 'POST',
    params,
  });
}


// 新建/编辑生成演示
export function Edit(params) {
  return http.request({
    url: '/curdDemo/edit',
    method: 'POST',
    params,
  });
}


// 修改生成演示状态
export function Status(params) {
  return http.request({
    url: '/curdDemo/status',
    method: 'POST',
    params,
  });
}


// 操作生成演示开关
export function Switch(params) {
  return http.request({
    url: '/curdDemo/switch',
    method: 'POST',
    params,
  });
}


// 获取生成演示指定详情
export function View(params) {
  return http.request({
    url: '/curdDemo/view',
    method: 'GET',
    params,
  });
}


// 获取生成演示最大排序
export function MaxSort() {
  return http.request({
    url: '/curdDemo/maxSort',
    method: 'GET',
  });
}


// 导出生成演示
export function Export(params) {
  jumpExport('/curdDemo/export', params);
}