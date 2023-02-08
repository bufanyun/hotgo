import { http } from '@/utils/http/axios';

// 获取验证码
export function GetCaptcha() {
  return http.request({
    url: '/site/captcha',
    method: 'get',
  });
}

// 上传图片
export function UploadImage(params) {
  return http.request({
    url: '/upload/image',
    method: 'post',
    params,
  });
}
