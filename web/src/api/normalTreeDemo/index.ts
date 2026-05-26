import { http, jumpExport } from '@/utils/http/axios';

// 获取普通树表列表
export function List(params) {
  return http.request({
    url: '/normalTreeDemo/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除普通树表
export function Delete(params) {
  return http.request({
    url: '/normalTreeDemo/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑普通树表
export function Edit(params) {
  return http.request({
    url: '/normalTreeDemo/edit',
    method: 'POST',
    params,
  });
}

// 获取普通树表指定详情
export function View(params) {
  return http.request({
    url: '/normalTreeDemo/view',
    method: 'GET',
    params,
  });
}

// 获取普通树表最大排序
export function MaxSort() {
  return http.request({
    url: '/normalTreeDemo/maxSort',
    method: 'GET',
  });
}

// 获取普通树表关系树选项
export function TreeOption() {
  return http.request({
    url: '/normalTreeDemo/treeOption',
    method: 'GET',
  });
}