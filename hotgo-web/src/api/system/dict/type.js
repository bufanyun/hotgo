import request from '@/utils/request'

// 查询字典类型列表
export function listType(query) {
  return request({
    url: '/dict_type/list',
    method: 'get',
    params: query
  })
}

// 查询字典类型详细
export function getType(dictId) {
  return request({
    url: '/dict_type/view',
    method: 'get',
    params: { id: dictId }
  })
}

// 新增字典类型
export function saveType(data) {
  return request({
    url: '/dict_type/edit',
    method: 'post',
    data: data
  })
}

// 删除字典类型
export function delType(dictId) {
  return request({
    url: '/dict_type/delete',
    method: 'post',
    params: { id: dictId }
  })
}

// 刷新字典缓存
export function refreshCache() {
  return request({
    url: '/dict_type/refresh_cache',
    method: 'get'
  })
}

// 导出字典类型
export function exportType(query) {
  return request({
    url: '/dict_type/export',
    method: 'get',
    params: query
  })
}

// 获取字典选择框列表
export function optionselect() {
  return request({
    url: '/system/dict/type/optionselect',
    method: 'get'
  })
}

// 查询字典类型列表
export function checkDictTypeUnique(data) {
  return request({
    url: '/dict_type/unique',
    method: 'get',
    params: data
  })
}
