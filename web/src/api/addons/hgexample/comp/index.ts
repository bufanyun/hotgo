import { http } from '@/utils/http/axios';
import type { UploadFileParams } from '@/utils/http/axios/types';

// 导入Excel
export function ImportExcel(params: UploadFileParams) {
  return http.uploadFile(
    {
      url: '/hgexample/comp/importExcel',
      method: 'post',
    },
    params
  );
}
