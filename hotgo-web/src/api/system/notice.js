import request from '@/utils/request'

// 查询公告列表
export function listNotice(query) {
  return request({
    url: '/notice/list',
    method: 'get',
    params: query
  })
}

// 查询公告详细
export function getNotice(noticeId) {
  return request({
    url: '/notice/view',
    method: 'get',
    params: { id: noticeId }
  })
}

// 新增公告
export function saveNotice(data) {
  return request({
    url: '/notice/edit',
    method: 'post',
    data: data
  })
}

// 删除公告
export function delNotice(noticeId) {
  return request({
    url: '/notice/delete',
    method: 'post',
    params: { id: noticeId }
  })
}

// 查看页面查询公告详细
export function getNoticeView(noticeId) {
  return request({
    url: '/system/notice/getNoticeView/' + noticeId,
    method: 'get'
  })
}

// 查询个人公告阅读列表
export function listNoticeByUser(query) {
  return request({
    url: '/system/notice/listNoticeByUser/page',
    method: 'get',
    params: query
  })
}

// 新增公告
export function updateNoticeToRead(noticeIds) {
  return request({
    url: '/system/notice/updateNoticeToRead/' + noticeIds,
    method: 'put'
  })
}
