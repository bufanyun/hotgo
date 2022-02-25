import request from '@/utils/request'

// 查询菜单列表
export function listMenu(query, menuId, expandLevel) {
  // console.log('query:' + JSON.stringify(query))
  if (menuId == null || menuId === '') {
    menuId = 0
  }
  if (expandLevel == null || expandLevel === '') {
    expandLevel = '2'
  }
  if (query === undefined) {
    query = {}
  }
  query.pid = menuId
  return request({
    url: '/menu/list',
    method: 'get',
    params: query
  })
}

// 查询菜单详细
export function getMenu(menuId) {
  return request({
    url: '/menu/view',
    method: 'get',
    params: { id: menuId }
  })
}

// 查询菜单下拉树结构
export function treeselect(menuId, expandLevel) {
  if (menuId == null || menuId === '') {
    menuId = '0'
  }
  if (expandLevel == null || expandLevel === '') {
    expandLevel = '2'
  }
  return request({
    url: '/system/menu/treeselect/' + expandLevel + '/' + menuId,
    method: 'get'
  })
}

// 查询菜单下拉树结构
export function menuTreeExcludeButton(menuId, expandLevel) {
  if (menuId == null || menuId === '') {
    menuId = '0'
  }
  if (expandLevel == null || expandLevel === '') {
    expandLevel = '2'
  }
  return request({
    url: '/system/menu/menuTreeExcludeButton/' + expandLevel + '/' + menuId,
    method: 'get'
  })
}

// 根据角色ID查询菜单下拉树结构
export function roleMenuTreeselect(roleId) {
  return request({
    url: '/menu/role_list',
    method: 'get',
    params: { id: roleId }
  })
}

// 新增菜单
export function saveMenu(data) {
  return request({
    url: '/menu/edit',
    method: 'post',
    data: data
  })
}

// 删除菜单
export function delMenu(menuId) {
  return request({
    url: '/menu/delete',
    method: 'post',
    params: { id: menuId }
  })
}

// 菜单树检索
export function searchMenuList(searchInfo) {
  return request({
    url: '/menu/search_list',
    method: 'get',
    params: searchInfo
  })
}

// 菜单名称唯一校验
export function checkMenuNameUnique(data) {
  return request({
    url: '/menu/name_unique',
    method: 'get',
    params: data
  })
}

// 路由地址唯一校验
export function checkMenuCodeUnique(data) {
  return request({
    url: '/menu/code_unique',
    method: 'get',
    params: data
  })
}

// 查询菜单同层最大排序
export function findMaxSort(parentId) {
  return request({
    url: '/menu/max_sort',
    method: 'get',
    params: { id: parentId }
  })
}
