import { http } from '@/utils/http/axios';

// 获取验证码
export function GetCaptcha() {
  return http.request({
    url: '/site/captcha',
    method: 'get',
  });
}
