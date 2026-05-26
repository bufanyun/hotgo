import { http, jumpExport } from '@/utils/http/axios';

// 获取服务许可证列表
export function List(params) {
  return http.request({
    url: '/serveLicense/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除服务许可证
export function Delete(params) {
  return http.request({
    url: '/serveLicense/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑服务许可证
export function Edit(params) {
  return http.request({
    url: '/serveLicense/edit',
    method: 'POST',
    params,
  });
}

// 修改服务许可证状态
export function Status(params) {
  return http.request({
    url: '/serveLicense/status',
    method: 'POST',
    params,
  });
}

// 获取服务许可证指定详情
export function View(params) {
  return http.request({
    url: '/serveLicense/view',
    method: 'GET',
    params,
  });
}

// 分配服务许可证路由
export function AssignRouter(params) {
  return http.request({
    url: '/serveLicense/assignRouter',
    method: 'POST',
    params,
  });
}

// 导出服务许可证
export function Export(params) {
  jumpExport('/serveLicense/export', params);
}
