import { http } from '@/utils/http/axios';

/**
 * 获取字典类型tree
 * @param params
 */
export function getDictTree(params?) {
  return http.request({
    url: '/dict_type/tree',
    method: 'GET',
    params,
  });
}

/**
 * 编辑字典类型
 * @param params
 * @constructor
 */
export function EditDict(params?) {
  return http.request({
    url: '/dict_type/edit',
    method: 'POST',
    params,
  });
}

/**
 * 删除字典类型
 * @param params
 * @constructor
 */
export function DeleteDict(params?) {
  return http.request({
    url: '/dict_type/delete',
    method: 'POST',
    params,
  });
}

/**
 * 获取字典类型选项
 * @param params
 */
export function getDictSelect(params?) {
  return http.request({
    url: '/dict_type/select',
    method: 'GET',
    params,
  });
}

/**
 * 编辑字典数据
 * @param params
 * @constructor
 */
export function EditData(params?) {
  return http.request({
    url: '/dict_data/edit',
    method: 'POST',
    params,
  });
}

/**
 * 删除字典数据
 * @param params
 * @constructor
 */
export function DeleteData(params?) {
  return http.request({
    url: '/dict_data/delete',
    method: 'POST',
    params,
  });
}

/**
 * 获取字典数据列表
 * @param params
 */
export function getDataList(params?) {
  return http.request({
    url: '/dict_data/list',
    method: 'GET',
    params,
  });
}
