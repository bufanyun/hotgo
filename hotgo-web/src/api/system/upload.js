import request from '@/utils/request'
// 上传附件
export function baseUpload (data) {
  return request({
    url: '/common/attach/baseupload/v1',
    method: 'post',
    data: data
  })
}
// 仅上传附件到磁盘不保存数据库
export function uploadDisk (data) {
  return request({
    url: '/common/attach/uploadDisk/v1',
    method: 'post',
    data: data
  })
}
