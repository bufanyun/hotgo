import request from '@/utils/request'

// 查询工作台小页管理列表
export function listSysPortlet (query) {
  return request({
    url: '/system/sysPortlet/list',
    method: 'get',
    params: query
  })
}

// 查询工作台小页管理详细
export function getSysPortlet (id) {
  return request({
    url: '/system/sysPortlet/' + id,
    method: 'get'
  })
}

// 新增工作台小页管理
export function addSysPortlet (data) {
  return request({
    url: '/system/sysPortlet',
    method: 'post',
    data: data
  })
}

// 修改工作台小页管理
export function updateSysPortlet (data) {
  return request({
    url: '/system/sysPortlet',
    method: 'put',
    data: data
  })
}

// 删除工作台小页管理
export function delSysPortlet (id) {
  return request({
    url: '/system/sysPortlet/' + id,
    method: 'delete'
  })
}

// 查询最大编号
export function findMaxSort () {
  return request({
    url: '/system/sysPortlet/findMaxSort',
    method: 'get'
  })
}
// 校验小页编码是否存在
export function checkCodeUnique (id, code) {
  if (id === undefined) {
    id = ''
  }
  return request({
    url: '/system/sysPortlet/checkCodeUnique/' + code + '/' + id,
    method: 'get'
  })
}

// 导出工作台小页管理
export function exportSysPortlet (query) {
  return request({
    url: '/system/sysPortlet/export',
    method: 'get',
    params: query
  })
}
// 查询工作台小页管理列表
export function listSysPortletByRoleId (query) {
  return request({
    url: '/system/sysPortlet/getSysPortletByRoleId',
    method: 'get',
    params: query
  })
}
