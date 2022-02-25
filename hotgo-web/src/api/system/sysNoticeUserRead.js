import request from '@/utils/request'

// 查询通知公告用户阅读列表
export function listSysNoticeUserRead (query) {
  return request({
    url: '/system/sysNoticeUserRead/list',
    method: 'get',
    params: query
  })
}

// 查询通知公告用户阅读详细
export function getSysNoticeUserRead (id) {
  return request({
    url: '/system/sysNoticeUserRead/' + id,
    method: 'get'
  })
}

// 新增通知公告用户阅读
export function addSysNoticeUserRead (data) {
  return request({
    url: '/system/sysNoticeUserRead',
    method: 'post',
    data: data
  })
}

// 修改通知公告用户阅读
export function updateSysNoticeUserRead (data) {
  return request({
    url: '/system/sysNoticeUserRead',
    method: 'put',
    data: data
  })
}

// 删除通知公告用户阅读
export function delSysNoticeUserRead (id) {
  return request({
    url: '/system/sysNoticeUserRead/' + id,
    method: 'delete'
  })
}

// 导出通知公告用户阅读
export function exportSysNoticeUserRead (query) {
  return request({
    url: '/system/sysNoticeUserRead/export',
    method: 'get',
    params: query
  })
}
