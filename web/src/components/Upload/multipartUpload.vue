<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      :on-after-leave="handleRemove"
      :style="{
        width: width,
      }"
      title="上传大文件"
    >
      <n-upload
        directory-dnd
        :custom-request="handleUpload"
        :on-remove="handleRemove"
        name="file"
        :disabled="uploadStatus != 0 && uploadStatus != 3"
      >
        <n-upload-dragger>
          <div style="margin-bottom: 12px">
            <n-icon size="48" :depth="3">
              <FileAddOutlined />
            </n-icon>
          </div>
          <template v-if="uploadStatus == 0 || uploadStatus == 3">
            <n-text style="font-size: 16px">点击或者拖动{{ typeTag }}到该区域来上传</n-text>
            <n-p depth="3" style="margin: 8px 0 0 0">支持大文件分片上传，支持断点续传</n-p>
          </template>
          <template v-else-if="uploadStatus == 1">
            <span style="font-weight: 600">解析中，请稍候...</span>
          </template>
          <template v-else-if="uploadStatus == 2">
            <span style="font-weight: 600">正在上传({{ progress }}%)...</span>
            <n-p depth="3" style="margin: 8px 0 0 0">文件大小：{{ sizeFormat }}</n-p>
          </template>
        </n-upload-dragger>
      </n-upload>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { NModal, UploadCustomRequestOptions, useMessage, useDialog } from 'naive-ui';
  import { FileAddOutlined } from '@vicons/antd';
  import SparkMD5 from 'spark-md5';
  import { Attachment, FileType, getFileType } from '@/components/FileChooser/src/model';
  import { CheckMultipart, UploadPart } from '@/api/base';
  import type { UploadFileParams } from '@/utils/http/axios/types';

  export interface Props {
    width?: string;
    uploadType?: FileType;
  }

  const props = withDefaults(defineProps<Props>(), {
    width: '60%',
    uploadType: 'default',
  });

  const emit = defineEmits(['onFinish']);
  const message = useMessage();
  const dialog = useDialog();
  const showModal = ref(false);
  const chunkSize = 2 * 1024 * 1024; // 每个分片大小限制，默认2M
  const uploadStatus = ref(0); // 上传状态 0等待上传 1解析中 2上传中 3已取消
  const progress = ref(0);
  const sizeFormat = ref('0B');
  const typeTag = computed(() => {
    return getFileType(props.uploadType);
  });

  // 取消上传
  function handleRemove() {
    if (uploadStatus.value == 1 || uploadStatus.value == 2) {
      uploadStatus.value = 3;
      dialog.info({
        title: '提示',
        content: '已取消大文件上传，已上传的文件不会自动删除，重新操作可进行断点续传',
        positiveText: '确定',
      });
    }
  }

  // 开始上传
  function handleUpload(options: UploadCustomRequestOptions) {
    uploadStatus.value = 1;

    // 初始化上传进度
    updateProgress(options, 0);

    const file = options.file.file as File;
    const fileReader = new FileReader();
    fileReader.readAsArrayBuffer(file);
    fileReader.onload = async (e) => {
      const spark = new SparkMD5.ArrayBuffer();
      spark.append(e.target.result);
      let md5 = spark.end();
      let start = 0;
      let end = 0;
      let index = 0;
      let shards: any[] = [];
      while (end < file.size) {
        start = index * chunkSize;
        end = (index + 1) * chunkSize;

        const params: UploadFileParams = {
          uploadType: props.uploadType,
          md5: md5,
          index: index + 1,
          fileName: file.name,
          file: file.slice(start, end),
        };

        const shard = { index: index + 1, params: params };
        shards.push(shard);
        index++;
      }

      uploadStatus.value = 2;

      const params = {
        uploadType: props.uploadType,
        fileName: file.name,
        size: file.size,
        md5: md5,
        shardCount: shards.length,
      };

      CheckMultipart(params)
        .then(async (res) => {
          // 已存在
          if (!res.waitUploadIndex || res.waitUploadIndex.length == 0) {
            onFinish(options, res.attachment);
            return;
          }

          // 断点续传，过滤掉已上传成功的分片文件
          shards = shards.filter((shard) => res.waitUploadIndex.includes(shard.index));
          if (shards.length == 0) {
            onFinish(options, res.attachment);
            return;
          }

          // 导入断点续传进度
          updateProgress(options, res.progress);
          sizeFormat.value = res.sizeFormat;

          for (const item of shards) {
            if (uploadStatus.value == 3) {
              break;
            }
            item.params.uploadId = res.uploadId;
            await handleUploadPart(options, item);
          }
        })
        .catch(() => {
          uploadStatus.value = 0;
          options.onError();
        });
    };
  }

  // 上传分片文件
  async function handleUploadPart(options: UploadCustomRequestOptions, item) {
    const res = await UploadPart(item.params);
    updateProgress(options, res.progress);
    if (res.finish) {
      onFinish(options, res.attachment);
    }
  }

  // 更新上传进度
  function updateProgress(options: UploadCustomRequestOptions, value: number) {
    options.onProgress({ percent: value });
    progress.value = value;
  }

  // 上传成功后的回调
  function onFinish(options: UploadCustomRequestOptions, result: Attachment) {
    options.onFinish();
    message.success('上传成功');
    uploadStatus.value = 0;
    emit('onFinish', result, true);
  }

  function openModal() {
    showModal.value = true;
    uploadStatus.value = 0;
  }

  defineExpose({
    openModal,
  });
</script>
