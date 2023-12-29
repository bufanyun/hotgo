import { http } from '@/utils/http/axios';
import { useUserStoreWidthOut } from '@/store/modules/user';

// 获取验证码
export function GetCaptcha() {
  return http.request({
    url: '/site/captcha',
    method: 'get',
  });
}

// 上传图片
export function UploadImage(params) {
  const useUserStore = useUserStoreWidthOut();
  const headers = {
    Authorization: useUserStore.token,
    uploadType: 'default',
  };
  return http.request({
    url: '/upload/file',
    method: 'post',
    params,
    headers,
  });
}
