<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="导入excel">
        将excel表格数据导入，可解析出完整的表格内容，包括所有的sheet和行列数据
      </n-card>
    </div>
    <n-card :bordered="false" class="mt-4 proCard">
      <n-upload
        directory-dnd
        :custom-request="handleUpload"
        name="file"
        type="file"
        accept=".xlsx, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
      >
        <n-upload-dragger>
          <div style="margin-bottom: 12px">
            <n-icon size="48" :depth="3">
              <DownloadOutlined />
            </n-icon>
          </div>
          <n-text style="font-size: 16px"> 点击或者拖动.xlsx文件到该区域来上传</n-text>
          <n-p depth="3" style="margin: 8px 0 0 0"> 单次上传数据最大不建议超过5000条</n-p>
        </n-upload-dragger>
      </n-upload>
    </n-card>

    <n-card
      :bordered="false"
      class="proCard mt-4"
      size="small"
      :segmented="{ content: true }"
      title="表格数据"
    >
      <n-scrollbar style="max-height: 520px">
        <JsonViewer :value="response" :expand-depth="5" copyable boxed sort class="json-width" />
      </n-scrollbar>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { UploadCustomRequestOptions, useMessage } from 'naive-ui';
  import { DownloadOutlined } from '@vicons/antd';
  import { ImportExcel } from '@/api/addons/hgexample/comp';
  import type { UploadFileParams } from '@/utils/http/axios/types';
  import { JsonViewer } from 'vue3-json-viewer';
  import 'vue3-json-viewer/dist/vue3-json-viewer.css';

  const message = useMessage();
  const response = ref<any>({});

  function handleUpload(options: UploadCustomRequestOptions) {
    message.loading('正在导入，请稍候...', { duration: 1200 });
    const params: UploadFileParams = {
      file: options.file.file as File,
    };

    ImportExcel(params).then((res) => {
      response.value = res;
      message.destroyAll();
      message.success('解析成功');
    });
  }
</script>

<style lang="less"></style>
