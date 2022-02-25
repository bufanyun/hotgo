import request from '@/utils/request'

// 查询数据权限方法维护列表
export function listSysDataPermissionsMethod (query) {
  return request({
    url: '/monitor/sysDataPermissionsMethod/list',
    method: 'get',
    params: query
  })
}

// 查询数据权限方法维护详细
export function getSysDataPermissionsMethod (id) {
  return request({
    url: '/monitor/sysDataPermissionsMethod/' + id,
    method: 'get'
  })
}

// 新增数据权限方法维护
export function addSysDataPermissionsMethod (data) {
  return request({
    url: '/monitor/sysDataPermissionsMethod',
    method: 'post',
    data: data
  })
}

// 修改数据权限方法维护
export function updateSysDataPermissionsMethod (data) {
  return request({
    url: '/monitor/sysDataPermissionsMethod',
    method: 'put',
    data: data
  })
}

// 删除数据权限方法维护
export function delSysDataPermissionsMethod (id) {
  return request({
    url: '/monitor/sysDataPermissionsMethod/' + id,
    method: 'delete'
  })
}

// 导出数据权限方法维护
export function exportSysDataPermissionsMethod (query) {
  return request({
    url: '/monitor/sysDataPermissionsMethod/export',
    method: 'get',
    params: query
  })
}

// 新增数据权限方法维护
export function getAllMapperData () {
  return request({
    url: '/monitor/sysDataPermissionsMethod/getAllMapperData',
    method: 'get'
  })
}

// 新增数据权限方法维护
export function getMapperName (searchInfo) {
  return request({
    url: '/monitor/sysDataPermissionsMethod/getMapperName',
    method: 'get',
    params: searchInfo
  })
}

export function getMethodHtml (searchInfo) {
  return request({
    url: '/monitor/sysDataPermissionsMethod/getMethodHtml',
    method: 'get',
    params: searchInfo
  })
}
