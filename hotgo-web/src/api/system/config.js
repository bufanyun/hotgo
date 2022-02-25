import request from '@/utils/request'

// 查询参数列表
export function listConfig (query) {
  return request({
    url: '/system/config/page',
    method: 'get',
    params: query
  })
}

// 查询参数详细
export function getConfig (configId) {
  return request({
    url: '/system/config/' + configId,
    method: 'get'
  })
}

// 根据参数键名查询参数值
export function getConfigKey (configKey) {
  return request({
    url: '/config/get_value/',
    method: 'get',
    params: { key: configKey }
  })
}

// 新增参数配置
export function saveConfig (data) {
  return request({
    url: '/system/config',
    method: 'post',
    data: data
  })
}

// 删除参数配置
export function delConfig (configId) {
  return request({
    url: '/system/config/' + configId,
    method: 'delete'
  })
}

// 清理参数缓存
export function refreshCache () {
  return request({
    url: '/system/config/refreshCache',
    method: 'delete'
  })
}

// 导出参数
export function exportConfig (query) {
  return request({
    url: '/system/config/export',
    method: 'get',
    params: query
  })
}

// 参数列表唯一校验
export function checkConfigKeyUnique (data) {
  return request({
    url: 'system/config/checkConfigKeyUnique',
    method: 'get',
    params: data
  })
}
