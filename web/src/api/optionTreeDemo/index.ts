import { http, jumpExport } from '@/utils/http/axios';

// 获取选项树表列表
export function List(params) {
  return http.request({
    url: '/optionTreeDemo/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除选项树表
export function Delete(params) {
  return http.request({
    url: '/optionTreeDemo/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑选项树表
export function Edit(params) {
  return http.request({
    url: '/optionTreeDemo/edit',
    method: 'POST',
    params,
  });
}

// 获取选项树表指定详情
export function View(params) {
  return http.request({
    url: '/optionTreeDemo/view',
    method: 'GET',
    params,
  });
}

// 获取选项树表最大排序
export function MaxSort() {
  return http.request({
    url: '/optionTreeDemo/maxSort',
    method: 'GET',
  });
}

// 获取选项树表关系树选项
export function TreeOption() {
  return http.request({
    url: '/optionTreeDemo/treeOption',
    method: 'GET',
  });
}