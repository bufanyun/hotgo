<template>
  <n-modal
    v-model:show="showFileModal"
    :show-icon="false"
    preset="dialog"
    :style="{
      width: width,
    }"
    :title="'上传' + typeTag"
  >
    <n-upload
      multiple
      directory-dnd
      :action="`${uploadUrl}${urlPrefix}/upload/file`"
      :headers="uploadHeaders"
      :data="{ type: 0 }"
      @finish="finish"
      name="file"
      :max="maxUpload"
      :default-file-list="fileList"
      list-type="image"
    >
      <n-upload-dragger>
        <div style="margin-bottom: 12px">
          <n-icon size="48" :depth="3">
            <FileAddOutlined />
          </n-icon>
        </div>
        <n-text style="font-size: 16px"> 点击或者拖动{{ typeTag }}到该区域来上传</n-text>
        <n-p depth="3" style="margin: 8px 0 0 0"> 单次最多允许{{ maxUpload }}个{{ typeTag }}</n-p>
      </n-upload-dragger>
    </n-upload>
  </n-modal>
</template>

<script lang="ts" setup>
  import { computed, reactive, ref } from 'vue';
  import { FileAddOutlined } from '@vicons/antd';
  import { useUserStoreWidthOut } from '@/store/modules/user';
  import { useGlobSetting } from '@/hooks/setting';
  import { NModal, UploadFileInfo, useMessage } from 'naive-ui';
  import componentSetting from '@/settings/componentSetting';
  import { ResultEnum } from '@/enums/httpEnum';
  import { Attachment, FileType, getFileType, UploadTag } from '@/components/FileChooser/src/model';

  export interface Props {
    width?: string;
    maxUpload?: number;
    finishCall?: Function | null;
    uploadType?: FileType;
  }

  const props = withDefaults(defineProps<Props>(), {
    width: '60%',
    maxUpload: 20,
    finishCall: null,
    uploadType: 'default',
  });

  const fileList = ref<UploadFileInfo[]>([]);
  const showFileModal = ref(false);
  const message = useMessage();
  const useUserStore = useUserStoreWidthOut();
  const globSetting = useGlobSetting();
  const { uploadUrl } = globSetting;
  const urlPrefix = globSetting.urlPrefix || '';
  const uploadHeaders = reactive({
    Authorization: useUserStore.token,
    uploadType: props.uploadType,
  });

  const typeTag = computed(() => {
    return getFileType(props.uploadType);
  });

  //上传结束
  function finish({ event: Event }) {
    const res = eval('(' + Event.target.response + ')');
    const infoField = componentSetting.upload.apiSetting.infoField;
    const { code } = res;
    const msg = res.msg || res.message || '上传失败';
    const result = res[infoField] as Attachment;

    //成功
    if (code === ResultEnum.SUCCESS) {
      fileList.value.push({
        id: result.md5,
        name: result.name,
        status: 'finished',
        type: result.naiveType,
      });

      message.success('上传' + result.name + '成功');
      if (props.finishCall !== null) {
        props.finishCall(result, true);
      }
    } else {
      message.error(msg);
      if (props.finishCall !== null) {
        props.finishCall(result, false);
      }
    }
  }

  function openModal() {
    showFileModal.value = true;
    fileList.value = [];
  }

  defineExpose({
    openModal,
  });
</script>
