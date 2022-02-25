import request from '@/utils/request'

// 查询默认规则维护列表
export function listSysDataPermissionsDefRule (query) {
  return request({
    url: '/monitor/sysDataPermissionsDefRule/list',
    method: 'get',
    params: query
  })
}

// 查询默认规则维护详细
export function getSysDataPermissionsDefRule (id) {
  return request({
    url: '/monitor/sysDataPermissionsDefRule/' + id,
    method: 'get'
  })
}

// 新增默认规则维护
export function addSysDataPermissionsDefRule (data) {
  return request({
    url: '/monitor/sysDataPermissionsDefRule',
    method: 'post',
    data: data
  })
}

// 修改默认规则维护
export function updateSysDataPermissionsDefRule (data) {
  return request({
    url: '/monitor/sysDataPermissionsDefRule',
    method: 'put',
    data: data
  })
}

// 删除默认规则维护
export function delSysDataPermissionsDefRule (id) {
  return request({
    url: '/monitor/sysDataPermissionsDefRule/' + id,
    method: 'delete'
  })
}

// 导出默认规则维护
export function exportSysDataPermissionsDefRule (query) {
  return request({
    url: '/monitor/sysDataPermissionsDefRule/export',
    method: 'get',
    params: query
  })
}

// 查询菜单同层最大排序
export function findMaxSort () {
  return request({
    url: '/monitor/sysDataPermissionsDefRule/findMaxSort',
    method: 'get'
  })
}

// 获得全部的且已启用的默认规则
export function getAllDefaultRule () {
  return request({
    url: '/monitor/sysDataPermissionsDefRule/getAllDefaultRule',
    method: 'get'
  })
}
