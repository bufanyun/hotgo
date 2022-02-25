import request from '@/utils/request'

// 查询字典数据列表
export function listData(query) {
  return request({
    url: '/dict_data/list',
    method: 'get',
    params: query
  })
}

// 查询字典数据详细
export function getData(dictCode) {
  return request({
    url: '/dict_data/view',
    method: 'get',
    params: { id: dictCode }
  })
}

// 根据字典类型查询字典数据信息
export function getDicts(dictType) {
  const params = { type: dictType }
  return request({
    url: '/dict/attribute',
    method: 'get',
    params: params
  })
}

// 根据字典类型查询字典数据信息
export function getAllDicts(dictType) {
  return request({
    url: '/system/dict/data/all/type/' + dictType,
    method: 'get'
  })
}

// 新增字典数据
export function saveData(data) {
  return request({
    url: '/dict_data/edit',
    method: 'post',
    data: data
  })
}

// 删除字典数据
export function delData(dictCode) {
  return request({
    url: '/dict_data/delete',
    method: 'post',
    params: { id: dictCode }
  })
}

// 导出字典数据
export function exportData(query) {
  return request({
    url: '/system/dict/data/export',
    method: 'get',
    params: query
  })
}

// 查询字典类型列表
export function checkDictDataValueUnique(data) {
  console.log('data:' + JSON.stringify(data))
  return request({
    url: '/dict_data/unique',
    method: 'get',
    params: data
  })
}

// 查询最大排序
export function findMaxSort(dictType) {
  return request({
    url: '/dict_data/max_sort',
    method: 'get',
    params: { type: dictType }
  })
}
