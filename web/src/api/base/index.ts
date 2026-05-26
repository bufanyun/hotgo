import { http } from '@/utils/http/axios';
import { useUserStoreWidthOut } from '@/store/modules/user';
import type { UploadFileParams } from '@/utils/http/axios/types';

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
    'Content-Type': 'multipart/form-data',
  };
  return http.request({
    url: '/upload/file',
    method: 'post',
    params,
    headers,
  });
}

// 检查文件分片
export function CheckMultipart(params) {
  return http.request({
    url: '/upload/checkMultipart',
    method: 'post',
    params,
  });
}

// 分片上传
export function UploadPart(params: UploadFileParams) {
  return http.uploadFile(
    {
      url: '/upload/uploadPart',
      method: 'post',
    },
    params
  );
}
