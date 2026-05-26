import { http, jumpExport } from '@/utils/http/axios';

// 获取CURD列表列表
export function List(params) {
  return http.request({
    url: '/curdDemo/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除CURD列表
export function Delete(params) {
  return http.request({
    url: '/curdDemo/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑CURD列表
export function Edit(params) {
  return http.request({
    url: '/curdDemo/edit',
    method: 'POST',
    params,
  });
}

// 修改CURD列表状态
export function Status(params) {
  return http.request({
    url: '/curdDemo/status',
    method: 'POST',
    params,
  });
}

// 操作CURD列表开关
export function Switch(params) {
  return http.request({
    url: '/curdDemo/switch',
    method: 'POST',
    params,
  });
}

// 获取CURD列表指定详情
export function View(params) {
  return http.request({
    url: '/curdDemo/view',
    method: 'GET',
    params,
  });
}

// 获取CURD列表最大排序
export function MaxSort() {
  return http.request({
    url: '/curdDemo/maxSort',
    method: 'GET',
  });
}

// 导出CURD列表
export function Export(params) {
  jumpExport('/curdDemo/export', params);
}