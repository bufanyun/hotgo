import { http, jumpExport } from '@/utils/http/axios';

// 列表
export function List(params) {
  return http.request({
    url: '/hgexample/treeTable/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除
export function Delete(params) {
  return http.request({
    url: '/hgexample/treeTable/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑
export function Edit(params) {
  return http.request({
    url: '/hgexample/treeTable/edit',
    method: 'POST',
    params,
  });
}

// 修改状态
export function Status(params) {
  return http.request({
    url: '/hgexample/treeTable/status',
    method: 'POST',
    params,
  });
}

// 操作开关
export function Switch(params) {
  return http.request({
    url: '/hgexample/treeTable/switch',
    method: 'POST',
    params,
  });
}

// 详情
export function View(params) {
  return http.request({
    url: '/hgexample/treeTable/view',
    method: 'GET',
    params,
  });
}

// 获取最大排序
export function MaxSort() {
  return http.request({
    url: '/hgexample/treeTable/maxSort',
    method: 'GET',
  });
}

// 树形选项
export function Select() {
  return http.request({
    url: '/hgexample/treeTable/select',
    method: 'get',
  });
}

// 导出
export function Export(params) {
  jumpExport('/hgexample/treeTable/export', params);
}
