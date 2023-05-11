import { http, jumpExport } from '@/utils/http/axios';

// 获取交易退款列表
export function List(params) {
  return http.request({
    url: '/payRefund/list',
    method: 'get',
    params,
  });
}

// 导出交易退款
export function Export(params) {
  jumpExport('/payRefund/export', params);
}
