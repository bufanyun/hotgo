import request from '@/utils/request'

// 查询个性化配置详细
export function getSysTableConfig (tableKey) {
  return request({
    url: '/system/sysTableConfig/getInfoByTableKey/' + tableKey,
    method: 'get'
  })
}

// 新增个性化配置
export function addSysTableConfig (data) {
  return request({
    url: '/system/sysTableConfig',
    method: 'post',
    data: data
  })
}

// 修改个性化配置
export function updateSysTableConfig (data) {
  return request({
    url: '/system/sysTableConfig',
    method: 'put',
    data: data
  })
}
