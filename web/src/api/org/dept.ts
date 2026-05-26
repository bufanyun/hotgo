import { http } from '@/utils/http/axios';

export function getDeptList(params?) {
  return http.request({
    url: '/dept/list',
    method: 'GET',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/dept/edit',
    method: 'POST',
    params,
  });
}

// 部门状态
export function Status(params) {
  return http.request({
    url: '/dept/status',
    method: 'POST',
    params,
  });
}

// 获取管理员_部门指定详情
export function View(params) {
  return http.request({
    url: '/dept/view',
    method: 'GET',
    params,
  });
}

// 获取管理员_部门最大排序
export function MaxSort() {
  return http.request({
    url: '/dept/maxSort',
    method: 'GET',
  });
}

export function Delete(params) {
  return http.request({
    url: '/dept/delete',
    method: 'POST',
    params,
  });
}

export function getDeptOption() {
  const params = { pageSize: 100 };
  return http.request({
    url: '/dept/option',
    method: 'GET',
    params,
  });
}

// 部门关系树选项
export function TreeOption() {
  return http.request({
    url: '/dept/treeOption',
    method: 'GET',
  });
}
