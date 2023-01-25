<template>
  <div class="w-full">
    <div class="upload">
      <div class="upload-card">
        <!--图片列表-->
        <div
          class="upload-card-item"
          :style="getCSSProperties"
          v-for="(item, index) in imgList"
          :key="`img_${index}`"
        >
          <div class="upload-card-item-info">
            <div class="img-box">
              <template v-if="fileType === 'image'">
                <img :src="item" @error="errorImg($event)" />
              </template>
              <template v-else>
                <n-avatar :style="fileAvatarCSS">{{ getFileExt(item) }}</n-avatar>
              </template>
            </div>
            <div class="img-box-actions">
              <template v-if="fileType === 'image'">
                <n-icon size="18" class="mx-2 action-icon" @click="preview(item)">
                  <EyeOutlined />
                </n-icon>
              </template>
              <template v-else>
                <n-icon size="18" class="mx-2 action-icon" @click="download(item)">
                  <CloudDownloadOutlined />
                </n-icon>
              </template>
              <n-icon size="18" class="mx-2 action-icon" @click="remove(index)">
                <DeleteOutlined />
              </n-icon>
            </div>
          </div>
        </div>

        <!--上传图片-->
        <div
          class="upload-card-item upload-card-item-select-picture"
          :style="getCSSProperties"
          v-if="imgList.length < maxNumber"
        >
          <n-upload
            v-bind="$props"
            :file-list-style="{ display: 'none' }"
            @before-upload="beforeUpload"
            @finish="finish"
          >
            <div class="flex flex-col justify-center">
              <n-icon size="18" class="m-auto">
                <PlusOutlined />
              </n-icon>
              <span class="upload-title">{{ uploadTitle }}</span>
            </div>
          </n-upload>
        </div>
      </div>
    </div>

    <!--上传图片-->
    <n-space>
      <n-alert title="提示" type="info" v-if="helpText" class="flex w-full">
        {{ helpText }}
      </n-alert>
    </n-space>
  </div>

  <!--预览图片-->
  <n-modal
    v-model:show="showModal"
    preset="card"
    title="预览"
    :bordered="false"
    :style="{ width: '520px' }"
  >
    <img :src="previewUrl" />
  </n-modal>
</template>

<script lang="ts">
  import { defineComponent, toRefs, reactive, computed, watch, onMounted, ref } from 'vue';
  import { EyeOutlined, DeleteOutlined, PlusOutlined, CloudDownloadOutlined } from '@vicons/antd';
  import { basicProps } from './props';
  import { useMessage, useDialog } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import componentSetting from '@/settings/componentSetting';
  import { useGlobSetting } from '@/hooks/setting';
  import { isJsonString, isNullOrUnDef } from '@/utils/is';
  import { getFileExt } from '@/utils/urlUtils';
  import { errorImg } from '@/utils/hotgo';
  const globSetting = useGlobSetting();

  export default defineComponent({
    name: 'BasicUpload',

    components: { EyeOutlined, DeleteOutlined, PlusOutlined, CloudDownloadOutlined },
    props: {
      ...basicProps,
    },
    emits: ['uploadChange', 'delete'],
    setup(props, { emit }) {
      const getCSSProperties = computed(() => {
        return {
          width: `${props.width}px`,
          height: `${props.height}px`,
        };
      });

      const message = useMessage();
      const dialog = useDialog();
      const uploadTitle = ref(props.fileType === 'image' ? '上传图片' : '上传附件');
      const fileAvatarCSS = computed(() => {
        return {
          '--n-merged-size': `var(--n-avatar-size-override, ${props.width * 0.8}px)`,
          '--n-font-size': `18px`,
        };
      });

      const state = reactive({
        showModal: false,
        previewUrl: '',
        originalImgList: [] as string[],
        imgList: [] as string[],
      });

      //赋值默认图片显示
      watch(
        () => props.value,
        () => {
          loadValue(props.value);
          return;
        }
      );

      watch(
        () => props.values,
        () => {
          loadValue(props.values);
          return;
        }
      );

      // 加载默认
      function loadValue(value: any) {
        if (value === null) {
          return;
        }

        let data: string[] = [];
        if (isJsonString(value)) {
          value = JSON.parse(value);
        }

        // 单图模式
        if (typeof value === 'string') {
          if (value !== '') {
            data.push(value);
          }
        } else {
          // 多图模式
          data = value;
        }

        state.imgList = data.map((item) => {
          return getImgUrl(item);
        });
        state.originalImgList = state.imgList;
      }

      //预览
      function preview(url: string) {
        state.showModal = true;
        state.previewUrl = url;
      }
      //下载
      function download(url: string) {
        window.open(url);
      }

      //删除
      function remove(index: number) {
        dialog.info({
          title: '提示',
          content: '你确定要删除吗？',
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: () => {
            state.imgList.splice(index, 1);
            state.originalImgList.splice(index, 1);
            if (props.maxNumber === 1) {
              emit('uploadChange', '');
            } else {
              emit('uploadChange', state.originalImgList);
            }
            emit('delete', state.originalImgList);
          },
          onNegativeClick: () => {},
        });
      }

      //组装完整图片地址
      function getImgUrl(url: string): string {
        const { imgUrl } = globSetting;
        return /(^http|https:\/\/)/g.test(url) ? url : `${imgUrl}${url}`;
      }

      function checkFileType(map: string[], fileType: string) {
        if (isNullOrUnDef(map)) {
          return true;
        }
        return map.includes(fileType);
      }

      //上传之前
      function beforeUpload({ file }) {
        const fileInfo = file.file;
        // 设置最大值，则判断
        if (props.maxSize && fileInfo.size / 1024 / 1024 >= props.maxSize) {
          message.error(`上传文件最大值不能超过${props.maxSize}M`);
          return false;
        }

        // 设置类型,则判断
        const fileType =
          props.fileType === 'image'
            ? componentSetting.upload.imageType
            : componentSetting.upload.fileType;
        if (!checkFileType(fileType, fileInfo.type)) {
          console.log('checkFileType fileInfo.type:' + fileInfo.type);
          message.error(`只能上传文件类型为${fileType.join(',')}`);
          return false;
        }

        return true;
      }

      //上传结束
      function finish({ event: Event }) {
        const res = eval('(' + Event.target.response + ')');
        const infoField = componentSetting.upload.apiSetting.infoField;
        const imgField = componentSetting.upload.apiSetting.imgField;
        const { code } = res;
        const msg = res.msg || res.message || '上传失败';
        const result = res[infoField];
        //成功
        if (code === ResultEnum.SUCCESS) {
          let imgUrl: string = getImgUrl(result[imgField]);
          state.imgList.push(imgUrl);
          state.originalImgList = state.imgList;
          if (props.maxNumber === 1) {
            emit('uploadChange', imgUrl);
          } else {
            emit('uploadChange', state.originalImgList);
          }
        } else {
          message.error(msg);
        }
      }

      onMounted(async () => {
        setTimeout(function () {
          if (props.maxNumber === 1) {
            loadValue(props.value);
          } else {
            loadValue(props.values);
          }
        }, 50);
      });
      return {
        errorImg,
        ...toRefs(state),
        finish,
        preview,
        download,
        remove,
        beforeUpload,
        getCSSProperties,
        uploadTitle,
        fileAvatarCSS,
        getFileExt,
      };
    },
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
