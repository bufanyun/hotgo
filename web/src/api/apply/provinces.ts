import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/provinces/list',
    method: 'get',
    params,
  });
}

export function Delete(params) {
  return http.request({
    url: '/provinces/delete',
    method: 'POST',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/provinces/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/provinces/status',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/provinces/view',
    method: 'GET',
    params,
  });
}

// 获取最大排序
export function MaxSort() {
  return http.request({
    url: '/provinces/maxSort',
    method: 'GET',
  });
}

/**
 * 获取省市区关系树选项列表
 */
export function getProvincesTree() {
  return http.request({
    url: '/provinces/tree',
    method: 'GET',
  });
}

/**
 * 获取省市区下级列表
 */
export function getProvincesChildrenList(params) {
  return http.request({
    url: '/provinces/childrenList',
    method: 'GET',
    params,
  });
}

/**
 * 唯一省市区ID
 */
export function CheckProvincesUniqueId(params) {
  return http.request({
    url: '/provinces/uniqueId',
    method: 'GET',
    params,
  });
}
