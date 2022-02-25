import request from '@/utils/request'

// 查询模板配置列表
export function listTemplate (query) {
  return request({
    url: '/system/template/list',
    method: 'get',
    params: query
  })
}

// 查询模板配置详细
export function getTemplate (id) {
  return request({
    url: '/system/template/' + id,
    method: 'get'
  })
}

// 新增模板配置
export function addTemplate (data) {
  return request({
    url: '/system/template',
    method: 'post',
    data: data
  })
}

// 修改模板配置
export function updateTemplate (data) {
  return request({
    url: '/system/template',
    method: 'put',
    data: data
  })
}

// 删除模板配置
export function delTemplate (id) {
  return request({
    url: '/system/template/' + id,
    method: 'delete'
  })
}

// 导出模板配置
export function exportTemplate (query) {
  return request({
    url: '/system/template/export',
    method: 'get',
    params: query
  })
}

// 状态修改
export function changeStatus (id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/system/template/changeStatus',
    method: 'put',
    data: data
  })
}

// 状态修改
export function changeTemplateDefault (id, templateDefault) {
  const data = {
    id,
    templateDefault
  }
  return request({
    url: '/system/template/changeTemplateDefault',
    method: 'put',
    data: data
  })
}

// 查询最大编号
export function findMaxSort () {
  return request({
    url: '/system/template/findMaxSort',
    method: 'get'
  })
}

// 校验模板名称是否存在
export function checkTemplateUnique (id, templateName) {
  if (id === undefined) {
    id = ''
  }
  return request({
    url: '/system/template/checkTemplateNameUnique/' + templateName + '/' + id,
    method: 'get'
  })
}
