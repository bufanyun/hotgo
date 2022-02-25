import request from '@/utils/request'

// 查询操作日志列表
export function list (query) {
  return request({
    url: '/log/list',
    method: 'get',
    params: query
  })
}

// 删除操作日志
export function delOperlog (operId) {
  return request({
    url: '/monitor/operlog/' + operId,
    method: 'delete'
  })
}

// 清空操作日志
export function cleanOperlog () {
  return request({
    url: '/log/clear',
    method: 'post'
  })
}

// 导出操作日志
export function exportOperlog (query) {
  return request({
    url: '/log/export',
    method: 'get',
    params: query
  })
}
