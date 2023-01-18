import { http } from '@/utils/http/axios';

export function List(params) {
  return http.request({
    url: '/genCodes/list',
    method: 'get',
    params,
  });
}
export function Delete(params) {
  return http.request({
    url: '/genCodes/delete',
    method: 'POST',
    params,
  });
}

export function View(params) {
  return http.request({
    url: '/genCodes/view',
    method: 'GET',
    params,
  });
}

export function Edit(params) {
  return http.request({
    url: '/genCodes/edit',
    method: 'POST',
    params,
  });
}

export function Status(params) {
  return http.request({
    url: '/genCodes/status',
    method: 'POST',
    params,
  });
}
export function Selects(params) {
  return http.request({
    url: '/genCodes/selects',
    method: 'get',
    params,
  });
}

export function TableSelect(params) {
  return http.request({
    url: '/genCodes/tableSelect',
    method: 'get',
    params,
  });
}

export function ColumnSelect(params) {
  return http.request({
    url: '/genCodes/columnSelect',
    method: 'get',
    params,
  });
}

export function ColumnList(params) {
  return http.request({
    url: '/genCodes/columnList',
    method: 'get',
    params,
  });
}

export function Preview(params) {
  return http.request({
    url: '/genCodes/preview',
    method: 'post',
    params,
  });
}

export function Build(params) {
  return http.request({
    url: '/genCodes/build',
    method: 'post',
    params,
  });
}
