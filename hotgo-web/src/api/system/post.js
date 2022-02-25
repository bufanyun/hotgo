import request from '@/utils/request'

// 查询岗位列表
export function listPost (query) {
  return request({
    url: '/post/list',
    method: 'get',
    params: query
  })
}

// 查询岗位详细
export function getPost (id) {
  return request({
    url: '/post/view',
    method: 'get',
    params: { id: id }
  })
}

// 新增岗位
export function savePost (data) {
  return request({
    url: '/post/edit',
    method: 'post',
    data: data
  })
}

// 删除岗位
export function delPost (id) {
  return request({
    url: '/post/delete',
    method: 'post',
    data: { id: id }
  })
}

// 导出岗位
export function exportPost (query) {
  return request({
    url: '/system/post/export',
    method: 'get',
    params: query
  })
}

// 岗位编码唯一校验
export function checkPostCodeUnique (data) {
  return request({
    url: '/post/code_unique',
    method: 'get',
    params: data
  })
}

// 岗位名称唯一校验
export function checkPostNameUnique (data) {
  return request({
    url: '/post/name_unique',
    method: 'get',
    params: data
  })
}

// 查询岗位最大排序
export function findMaxSort () {
  return request({
    url: '/post/max_sort',
    method: 'get'
  })
}
