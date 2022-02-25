import request from '@/utils/request'

// 查询数据权限规则维护子表列表
export function listSysDataPermissionsRule (query) {
  return request({
    url: '/monitor/sysDataPermissionsRule/list',
    method: 'get',
    params: query
  })
}

// 查询数据权限规则维护子表详细
export function getSysDataPermissionsRule (id) {
  return request({
    url: '/monitor/sysDataPermissionsRule/' + id,
    method: 'get'
  })
}

// 新增数据权限规则维护子表
export function addSysDataPermissionsRule (data) {
  return request({
    url: '/monitor/sysDataPermissionsRule',
    method: 'post',
    data: data
  })
}

// 修改数据权限规则维护子表
export function updateSysDataPermissionsRule (data) {
  return request({
    url: '/monitor/sysDataPermissionsRule',
    method: 'put',
    data: data
  })
}

// 删除数据权限规则维护子表
export function delSysDataPermissionsRule (id) {
  return request({
    url: '/monitor/sysDataPermissionsRule/' + id,
    method: 'delete'
  })
}

// 导出数据权限规则维护子表
export function exportSysDataPermissionsRule (query) {
  return request({
    url: '/monitor/sysDataPermissionsRule/export',
    method: 'get',
    params: query
  })
}

// 导出数据权限规则维护子表
export function getDataPermissionsMethodInfo (methodId) {
  return request({
    url: '/monitor/sysDataPermissionsRule/getDataPermissionsMethodInfo/' + methodId,
    method: 'get'
  })
}
