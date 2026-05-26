import { http, jumpExport } from '@/utils/http/axios';

// 获取资产变动列表
export function List(params) {
  return http.request({
    url: '/creditsLog/list',
    method: 'get',
    params,
  });
}

// 导出资产变动
export function Export(params) {
  jumpExport('/creditsLog/export', params);
}

// 获取变动状态选项
export function Option() {
  return http.request({
    url: '/creditsLog/option',
    method: 'GET',
  });
}
