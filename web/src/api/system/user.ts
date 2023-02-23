import { http } from '@/utils/http/axios';
import { ApiEnum } from '@/enums/apiEnum';

export interface BasicResponseModel<T = any> {
  code: number;
  message: string;
  data: T;
}

export interface BasicPageParams {
  pageNumber: number;
  pageSize: number;
  total: number;
}

export function getConfig() {
  return http.request({
    url: ApiEnum.SiteConfig,
    method: 'get',
    headers: { hostname: location.hostname },
  });
}

/**
 * @description: 获取用户信息
 */
export function getUserInfo() {
  return http.request({
    url: ApiEnum.MemberInfo,
    method: 'get',
  });
}

export function updateMemberProfile(params) {
  return http.request({
    url: '/member/updateProfile',
    method: 'post',
    params,
  });
}

export function updateMemberPwd(params) {
  return http.request({
    url: '/member/updatePwd',
    method: 'post',
    params,
  });
}

export function updateMemberMobile(params) {
  return http.request({
    url: '/member/updateMobile',
    method: 'post',
    params,
  });
}

export function updateMemberEmail(params) {
  return http.request({
    url: '/member/updateEmail',
    method: 'post',
    params,
  });
}

export function SendBindEmail() {
  return http.request({
    url: '/ems/sendBind',
    method: 'post',
  });
}

export function SendBindSms() {
  return http.request({
    url: '/sms/sendBind',
    method: 'post',
  });
}

export function updateMemberCash(params) {
  return http.request({
    url: '/member/updateCash',
    method: 'post',
    params,
  });
}

/**
 * @description: 用户登录
 */
export function login(params) {
  return http.request<BasicResponseModel>(
    {
      url: ApiEnum.SiteLogin,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户修改密码
 */
export function changePassword(params, uid) {
  return http.request(
    {
      url: `/user/u${uid}/changepw`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户登出
 */
export function logout(params) {
  return http.request({
    url: '/login/logout',
    method: 'POST',
    params,
  });
}
