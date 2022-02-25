import request from '@/utils/request'

// 查询多栏目门户配置列表
export function listSysPortalConfig (query) {
  return request({
    url: '/system/sysPortalConfig/list',
    method: 'get',
    params: query
  })
}

// 查询多栏目门户配置详细
export function getSysPortalConfig (id) {
  return request({
    url: '/system/sysPortalConfig/' + id,
    method: 'get'
  })
}

// 新增多栏目门户配置
export function addSysPortalConfig (data) {
  return request({
    url: '/system/sysPortalConfig',
    method: 'post',
    data: data
  })
}

// 修改多栏目门户配置
export function updateSysPortalConfig (data) {
  return request({
    url: '/system/sysPortalConfig',
    method: 'put',
    data: data
  })
}

// 删除多栏目门户配置
export function delSysPortalConfig (id) {
  return request({
    url: '/system/sysPortalConfig/' + id,
    method: 'delete'
  })
}

// 查询最大编号
export function findMaxSort () {
  return request({
    url: '/system/sysPortalConfig/findMaxSort',
    method: 'get'
  })
}
// 校验小页编码是否存在
export function checkCodeUnique (id, code) {
  if (id === undefined) {
    id = ''
  }
  return request({
    url: '/system/sysPortalConfig/checkCodeUnique/' + code + '/' + id,
    method: 'get'
  })
}

// 导出多栏目门户配置
export function exportSysPortalConfig (query) {
  return request({
    url: '/system/sysPortalConfig/export',
    method: 'get',
    params: query
  })
}
// 查询多栏目门户配置详细
export function getConfigAndPortalList (id) {
  return request({
    url: '/system/sysPortalConfig/getConfigAndPortalList/' + id,
    method: 'get'
  })
}
// 修改多栏目门户配置
export function updateDefaultPortalConfig (data) {
  return request({
    url: '/system/sysPortalConfig/updateDefaultPortalConfig',
    method: 'put',
    data: data
  })
}
// 查询模板列表
export function getPortalTemplateList () {
  return request({
    url: '/system/sysPortalConfig/getPortalTemplateList',
    method: 'get'
  })
}
