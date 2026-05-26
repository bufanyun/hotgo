import { http, jumpExport } from '@/utils/http/axios';

// 获取测试分类列表
export function List(params) {
  return http.request({
    url: '/testCategory/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除测试分类
export function Delete(params) {
  return http.request({
    url: '/testCategory/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑测试分类
export function Edit(params) {
  return http.request({
    url: '/testCategory/edit',
    method: 'POST',
    params,
  });
}

// 修改测试分类状态
export function Status(params) {
  return http.request({
    url: '/testCategory/status',
    method: 'POST',
    params,
  });
}

// 获取测试分类指定详情
export function View(params) {
  return http.request({
    url: '/testCategory/view',
    method: 'GET',
    params,
  });
}

// 获取测试分类最大排序
export function MaxSort() {
  return http.request({
    url: '/testCategory/maxSort',
    method: 'GET',
  });
}