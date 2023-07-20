import { http, jumpExport } from '@/utils/http/axios';

// 创建充值订单
export function Create(params) {
  return http.request({
    url: '/order/create',
    method: 'post',
    params,
  });
}

// 获取充值订单列表
export function List(params) {
  return http.request({
    url: '/order/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除充值订单
export function Delete(params) {
  return http.request({
    url: '/order/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑充值订单
export function Edit(params) {
  return http.request({
    url: '/order/edit',
    method: 'POST',
    params,
  });
}

// 修改充值订单状态
export function Status(params) {
  return http.request({
    url: '/order/status',
    method: 'POST',
    params,
  });
}

// 操作充值订单开关
export function Switch(params) {
  return http.request({
    url: '/order/switch',
    method: 'POST',
    params,
  });
}

// 获取充值订单指定详情
export function View(params) {
  return http.request({
    url: '/order/view',
    method: 'GET',
    params,
  });
}

// 获取充值订单最大排序
export function MaxSort() {
  return http.request({
    url: '/order/maxSort',
    method: 'GET',
  });
}

// 获取订单状态选项
export function Option() {
  return http.request({
    url: '/order/option',
    method: 'GET',
  });
}

// 申请订单退款
export function ApplyRefund(params) {
  return http.request({
    url: '/order/applyRefund',
    method: 'post',
    params,
  });
}

// 受理订单退款
export function AcceptRefund(params) {
  return http.request({
    url: '/order/acceptRefund',
    method: 'post',
    params,
  });
}

// 导出充值订单
export function Export(params) {
  jumpExport('/order/export', params);
}
