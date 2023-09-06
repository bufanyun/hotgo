<template>
  <n-space vertical>
    <n-button @click="showModal = true" size="small">
      <template #icon
        ><n-icon><PlusOutlined /></n-icon> </template
      >选择{{ buttonText }}
    </n-button>
    <div class="w-full">
      <div class="upload">
        <div class="upload-card" v-if="fileList.length > 0">
          <div
            class="upload-card-item"
            :style="getCSSProperties"
            v-for="(item, index) in fileList"
            :key="`img_${index}`"
          >
            <div class="upload-card-item-info">
              <div class="img-box">
                <template v-if="fileType === 'image'">
                  <img :src="item" @error="errorImg($event)" alt="" />
                </template>
                <template v-else>
                  <n-avatar :style="fileAvatarCSS">{{ getFileExt(item) }}</n-avatar>
                </template>
              </div>
              <div class="img-box-actions">
                <template v-if="fileType === 'image'">
                  <n-icon size="18" class="mx-2 action-icon" @click="handlePreview(item)">
                    <EyeOutlined />
                  </n-icon>
                </template>
                <template v-else>
                  <n-icon size="18" class="mx-2 action-icon" @click="handleDownload(item)">
                    <CloudDownloadOutlined />
                  </n-icon>
                </template>
                <n-icon size="18" class="mx-2 action-icon" @click="handleRemove(index)">
                  <DeleteOutlined />
                </n-icon>
              </div>
            </div>
          </div> </div></div
    ></div>
  </n-space>

  <n-modal
    v-model:show="showModal"
    :on-after-leave="handleCancel"
    :style="{
      width: dialogWidth,
    }"
  >
    <n-card title="文件选择">
      <template #header-extra>
        <n-space>
          <n-button @click="handleUpload" ghost>
            <template #icon>
              <n-icon :component="UploadOutlined" />
            </template>
            上传文件
          </n-button>
        </n-space>
      </template>
      <n-card style="overflow: auto" content-style="padding: 0;">
        <Chooser
          ref="chooserRef"
          :file-type="fileType"
          :maxNumber="maxNumber"
          :fileList="fileList"
          @saveChange="saveChange"
        />
      </n-card>
      <template #footer>
        <n-space justify="end">
          <n-button @click="handleCancel"> 取消 </n-button>
          <n-button type="primary" @click="handleSelectFile"> 确定 </n-button>
        </n-space>
      </template>
    </n-card>
  </n-modal>

  <FileUpload
    ref="fileUploadRef"
    :width="dialogWidth"
    :finish-call="handleFinishCall"
    max-upload="20"
  />

  <Preview ref="previewRef" />
</template>

<script lang="ts" setup>
  import { NButton, NSpace, NCard, NModal, NIcon, useDialog } from 'naive-ui';
  import { cloneDeep } from 'lodash-es';
  import FileUpload from '@/components/FileChooser/src/Upload.vue';
  import Chooser from '@/components/FileChooser/src/Chooser.vue';
  import Preview from '@/components/FileChooser/src/Preview.vue';
  import { computed, onMounted, ref, watch } from 'vue';
  import { adaModalWidth, errorImg } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';
  import {
    UploadOutlined,
    PlusOutlined,
    CloudDownloadOutlined,
    DeleteOutlined,
    EyeOutlined,
  } from '@vicons/antd';
  import { Attachment, FileType, getFileType } from '@/components/FileChooser/src/model';
  import { isArrayString, isString } from '@/utils/is';

  export interface Props {
    value: string | string[] | null;
    maxNumber?: number;
    fileType?: FileType;
    width?: number;
    height?: number;
  }

  const props = withDefaults(defineProps<Props>(), {
    value: '',
    maxNumber: 1,
    fileType: 'default',
    width: 100,
    height: 100,
  });

  const emit = defineEmits(['update:value']);
  const fileUploadRef = ref();
  const dialogWidth = ref('85%');
  const dialog = useDialog();
  const showModal = ref(false);
  const chooserRef = ref();
  const previewRef = ref();
  const fileList = ref<string[]>([]);

  const getCSSProperties = computed(() => {
    return {
      width: `${props.width}px`,
      height: `${props.height}px`,
    };
  });

  const fileAvatarCSS = computed(() => {
    return {
      '--n-merged-size': `var(--n-avatar-size-override, ${props.width * 0.8}px)`,
      '--n-font-size': `18px`,
    };
  });

  const buttonText = computed(() => {
    return getFileType(props.fileType);
  });

  // 预览
  function handlePreview(url: string) {
    previewRef.value.openPreview(url);
  }

  // 下载
  function handleDownload(url: string) {
    window.open(url);
  }

  // 删除
  function handleRemove(index: number) {
    dialog.info({
      title: '提示',
      content: '你确定要删除吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        fileList.value.splice(index, 1);
        if (props.maxNumber === 1) {
          emit('update:value', '');
        } else {
          emit('update:value', fileList.value);
        }
      },
      onNegativeClick: () => {},
    });
  }

  function handleSelectFile() {
    showModal.value = false;
    if (props.maxNumber === 1) {
      emit('update:value', fileList.value.length > 0 ? fileList.value[0] : '');
    } else {
      emit('update:value', fileList.value);
    }
  }

  function handleCancel() {
    showModal.value = false;
    loadImage();
  }

  function handleUpload() {
    fileUploadRef.value.openModal();
  }

  function handleFinishCall(result: Attachment, success: boolean) {
    if (success) {
      chooserRef.value.reloadTable();
    }
  }

  function saveChange(list: string[]) {
    fileList.value = list;
  }

  function loadImage() {
    const value = cloneDeep(props.value);
    if (props.maxNumber === 1) {
      fileList.value = [];
      if (value !== '') {
        if (!isString(value)) {
          console.warn(
            'When the file picker is currently in single-file mode, but the passed value is not of type string, there may be potential issues.'
          );
        }
        fileList.value.push(value as string);
      }
    } else {
      if (!isArrayString(value)) {
        console.warn(
          'When the file picker is currently in multiple-file mode, but the passed value is not of type string array, there may be potential issues.'
        );
      }
      if (!value) {
        fileList.value = [];
      } else {
        fileList.value = value as string[];
      }
    }
  }

  watch(
    () => props.value,
    () => {
      loadImage();
    },
    {
      immediate: true,
      deep: true,
    }
  );

  onMounted(async () => {
    adaModalWidth(dialogWidth, 1080);
    loadImage();
  });
</script>

<style lang="less">
  .n-upload {
    width: auto; /**  居中 */
  }

  .upload {
    width: 100%;
    overflow: hidden;

    &-card {
      width: auto;
      height: auto;
      display: flex;
      flex-wrap: wrap;
      align-items: center;

      &-item {
        margin: 0 8px 8px 0;
        position: relative;
        padding: 8px;
        border: 1px solid #d9d9d9;
        border-radius: 2px;
        display: flex;
        justify-content: center;
        flex-direction: column;
        align-items: center;

        &:hover {
          background: 0 0;

          .upload-card-item-info::before {
            opacity: 1;
          }

          &-info::before {
            opacity: 1;
          }
        }

        &-info {
          position: relative;
          height: 100%;
          padding: 0;
          overflow: hidden;

          &:hover {
            .img-box-actions {
              opacity: 1;
            }
          }

          &::before {
            position: absolute;
            z-index: 1;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 0, 0, 0.5);
            opacity: 0;
            transition: all 0.3s;
            content: ' ';
          }

          .img-box {
            position: relative;
            //padding: 8px;
            //border: 1px solid #d9d9d9;
            border-radius: 2px;
          }

          .img-box-actions {
            position: absolute;
            top: 50%;
            left: 50%;
            z-index: 10;
            white-space: nowrap;
            transform: translate(-50%, -50%);
            opacity: 0;
            transition: all 0.3s;
            display: flex;
            align-items: center;
            justify-content: space-between;

            &:hover {
              background: 0 0;
            }

            .action-icon {
              color: rgba(255, 255, 255, 0.85);

              &:hover {
                cursor: pointer;
                color: #fff;
              }
            }
          }
        }
      }

      &-item-select-picture {
        border: 1px dashed #d9d9d9;
        border-radius: 2px;
        cursor: pointer;
        background: #fafafa;
        color: #666;

        .upload-title {
          color: #666;
        }
      }
    }
  }
</style>
