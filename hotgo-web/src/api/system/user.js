import request from '@/utils/request'
import { praseStrEmpty } from '@/utils/aidex'

// 查询用户列表
export function listUser(query) {
  return request({
    url: '/member/list',
    method: 'get',
    params: query
  })
}

// 查询用户详细
export function getUser(userId) {
  return request({
    url: '/member/view',
    method: 'get',
    params: { id: praseStrEmpty(userId) }
  })
}

// 新增用户
export function addUser(data) {
  return request({
    url: '/member/edit',
    method: 'post',
    data: data
  })
}

// 修改用户
export function updateUser(data) {
  return request({
    url: '/member/edit',
    method: 'post',
    data: data
  })
}

// 删除用户
export function delUser(userId) {
  return request({
    url: '/member/delete',
    method: 'post',
    params: { id: praseStrEmpty(userId) }
  })
}

// 导出用户
export function exportUser(query) {
  return request({
    url: '/member/export',
    method: 'get',
    params: query
  })
}

// 用户密码重置
export function resetUserPwd(id, password) {
  const data = {
    id,
    password
  }
  return request({
    url: '/member/reset_pwd',
    method: 'post',
    data: data
  })
}

// 用户状态修改
export function changeUserStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/member/changeStatus',
    method: 'put',
    data: data
  })
}

// 查询用户个人信息
export function getUserProfile() {
  return request({
    url: '/member/profile',
    method: 'get'
  })
}

// 修改用户个人信息
export function updateUserProfile(data) {
  return request({
    url: '/member/update_profile',
    method: 'post',
    data: data
  })
}

// 用户密码重置
export function updateUserPwd(oldPassword, newPassword) {
  const data = {
    oldPassword,
    newPassword
  }
  return request({
    url: '/member/update_pwd',
    method: 'post',
    params: data
  })
}

// 用户头像上传
export function uploadAvatar(data) {
  return request({
    url: '/member/profile/avatar',
    method: 'post',
    data: data
  })
}

// 下载用户导入模板
export function importTemplate() {
  return request({
    url: '/system/user/importTemplate',
    method: 'get'
  })
}

// 校验用户名称唯一性
export function checkUserNameUnique(data) {
  return request({
    url: '/member/name_unique',
    method: 'get',
    params: data
  })
}

export function checkEmailUnique(data) {
  return request({
    url: '/member/email_unique',
    method: 'get',
    params: data
  })
}

export function checkPhoneUnique(data) {
  return request({
    url: '/member/mobile_unique',
    method: 'get',
    params: data
  })
}

// 查询用户详细
export function getUserInfoByIds(userIds) {
  return request({
    url: '/member/getUserInfoByIds',
    method: 'post',
    data: userIds
  })
}

// 查询角色下用户列表
export function getRoleUserList(query) {
  return request({
    url: '/role/member_list',
    method: 'get',
    params: query
  })
}

// 插入角色用户
export function saveRoleUser(data) {
  return request({
    url: '/member/addRoleUser',
    method: 'post',
    data: data
  })
}
