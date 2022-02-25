import request from '@/utils/request'

// 查询部门列表
export function listDept (query, deptId, expandLevel) {
  if (deptId == null || deptId === '') {
    deptId = '0'
  }
  if (expandLevel == null || expandLevel === '') {
    expandLevel = '1'
  }
  //  + expandLevel + '/' + deptId
  return request({
    url: '/dept/list/',
    method: 'get',
    params: query
  })
}

// 查询部门树列表（排除节点）
export function listDeptExcludeChild (deptId) {
  // return request({
  //   url: '/system/dept/listTreeExcludeChild/10/0/' + deptId,
  //   method: 'get'
  // })
  return request({
    url: '/dept/list/',
    method: 'get',
    params: { id: deptId }
  })
}

// 查询部门树列表（排除当前节点及子节点）
export function listDeptTree (deptId, expandLevel) {
	// if (deptId == null || deptId === '') {
	// 	deptId = '0'
	// }
  // if (expandLevel == null || expandLevel === '') {
  //   expandLevel = '1'
  // }
  // return request({
  //   url: '/system/dept/listTree/' + expandLevel + '/' + deptId,
  //   method: 'get'
  // })

  return request({
    url: '/dept/list_tree/',
    method: 'get',
    params: { id: deptId }
  })
}

// 查询部门详细
export function getDept (deptId) {
  return request({
    url: '/dept/view',
    method: 'get',
    params: { id: deptId }
  })
}

// 根据角色ID查询部门树结构
export function roleDeptTreeselect (roleId) {
  return request({
    url: '/system/dept/roleDeptTreeselect/' + roleId,
    method: 'get'
  })
}

// 新增部门
export function addDept (data) {
  return request({
    url: '/system/dept',
    method: 'post',
    data: data
  })
}

// 修改部门
export function updateDept (data) {
  return request({
    url: '/system/dept',
    method: 'put',
    data: data
  })
}

// 删除部门
export function delDept (deptId) {
  return request({
    url: '/system/dept/' + deptId,
    method: 'delete'
  })
}

// 查询部门下拉树结构
export function findMaxSort(parentId) {
  return request({
    url: '/dept/max_sort',
    method: 'get',
    params: { id: parentId }
  })
}

// 校验部门名称是否存在
export function validateDeptNameUnique (deptName, parentId, id) {
  if (id === undefined) {
    id = ''
  }
  return request({
    url: '/system/dept/validateDeptNameUnique/' + deptName + '/' + parentId + '/' + id,
    method: 'get'
  })
}
// 部门树检索
export function searchDept (searchInfo) {
  return request({
    url: '/system/dept/search',
    method: 'get',
    params: searchInfo
  })
}
// 部门树检索
export function searchDeptList (searchInfo) {
  // return request({
  //   url: '/system/dept/searchDeptList',
  //   method: 'get',
  //   params: searchInfo
  // })
  return request({
    url: '/dept/list/',
    method: 'get',
    params: searchInfo
  })
}
// 按部门分组人员树
export function userSelectTree (deptId, expandLevel) {
	if (deptId == null || deptId === '') {
		deptId = '0'
	}
  if (expandLevel == null || expandLevel === '') {
    expandLevel = '1'
  }
  // return request({
  //   url: '/dept/userSelectList/' + expandLevel + '/' + deptId,
  //   method: 'get'
  // })
  return request({
    url: '/dept/list/',
    method: 'get'
  })
}
// 按部门树检索用户
export function searchDeptUserList (searchInfo) {
  return request({
    url: '/system/dept/searchDeptUserList',
    method: 'get',
    params: searchInfo
  })
}

// 查询部门详细
export function getDeptInfoByIds (userIds) {
  // return request({
  //   url: '/system/dept/getDeptInfoByIds',
  //   method: 'post',
  //   data: userIds
  // })
  return request({
    url: '/dept/list/',
    method: 'get',
    params: userIds
  })
}
