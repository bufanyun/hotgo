import request from '@/utils/request'

// 查询用户主题信息记录列表
export function listSysThemeConfig (query) {
  return request({
    url: '/sysThemeConfig/sysThemeConfig/list',
    method: 'get',
    params: query
  })
}

// 查询用户主题信息记录详细
export function getSysThemeConfig (id) {
  return request({
    url: '/sysThemeConfig/sysThemeConfig/' + id,
    method: 'get'
  })
}

// 新增用户主题信息记录
export function addSysThemeConfig (data) {
  return request({
    url: '/sysThemeConfig/sysThemeConfig',
    method: 'post',
    data: data
  })
}

// 修改用户主题信息记录
export function updateSysThemeConfig (data) {
  return request({
    url: '/sysThemeConfig/sysThemeConfig',
    method: 'put',
    data: data
  })
}

// 删除用户主题信息记录
export function delSysThemeConfig (id) {
  return request({
    url: '/sysThemeConfig/sysThemeConfig/' + id,
    method: 'delete'
  })
}

// 查询最大编号
export function findMaxSort () {
  return request({
    url: '/sysThemeConfig/sysThemeConfig/findMaxSort',
    method: 'get'
  })
}
// 导出用户主题信息记录
export function exportSysThemeConfig (query) {
  return request({
    url: '/sysThemeConfig/sysThemeConfig/export',
    method: 'get',
    params: query
  })
}
