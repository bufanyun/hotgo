<template>
  <div>
    <n-card :bordered="false" class="proCard" title="附件管理">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        @keyup.enter="handleSubmit"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>

      <BasicTable
        :openChecked="true"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1800"
        :resizeHeightOffset="-20000"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <UploadOutlined />
              </n-icon>
            </template>
            上传图片
          </n-button>
          &nbsp;
          <n-button type="primary" @click="addFileTable">
            <template #icon>
              <n-icon>
                <UploadOutlined />
              </n-icon>
            </template>
            上传文件
          </n-button>
          &nbsp;
          <n-button type="error" @click="batchDelete" :disabled="batchDeleteDisabled">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        preset="dialog"
        style="width: 60%"
        title="上传图片"
      >
        <n-upload
          multiple
          directory-dnd
          :action="`${uploadUrl}${urlPrefix}/upload/image`"
          :headers="uploadHeaders"
          :data="{ type: 0 }"
          @before-upload="beforeUpload"
          @finish="finish"
          name="file"
          :max="20"
          :default-file-list="fileList"
          list-type="image"
        >
          <n-upload-dragger>
            <div style="margin-bottom: 12px">
              <n-icon size="48" :depth="3">
                <CloudUploadOutlined />
              </n-icon>
            </div>
            <n-text style="font-size: 16px"> 点击或者拖动图片到该区域来上传</n-text>
            <n-p depth="3" style="margin: 8px 0 0 0"> 单次最多允许20个图片</n-p>
          </n-upload-dragger>
        </n-upload>
      </n-modal>

      <n-modal
        v-model:show="showFileModal"
        :show-icon="false"
        preset="dialog"
        style="width: 60%"
        title="上传文件"
      >
        <n-upload
          multiple
          directory-dnd
          :action="`${uploadUrl}${urlPrefix}/upload/file`"
          :headers="uploadHeaders"
          :data="{ type: 0 }"
          @before-upload="beforeUpload"
          @finish="finish"
          name="file"
          :max="20"
          :default-file-list="fileList"
          list-type="image"
        >
          <n-upload-dragger>
            <div style="margin-bottom: 12px">
              <n-icon size="48" :depth="3">
                <FileAddOutlined />
              </n-icon>
            </div>
            <n-text style="font-size: 16px"> 点击或者拖动文件到该区域来上传</n-text>
            <n-p depth="3" style="margin: 8px 0 0 0"> 单次最多允许20个文件</n-p>
          </n-upload-dragger>
        </n-upload>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { UploadFileInfo, useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { Delete, List } from '@/api/apply/attachment';
  import { columns, schemas } from './columns';
  import {
    DeleteOutlined,
    UploadOutlined,
    FileAddOutlined,
    CloudUploadOutlined,
  } from '@vicons/antd';
  import { useGlobSetting } from '@/hooks/setting';
  import { useUserStoreWidthOut } from '@/store/modules/user';
  import componentSetting from '@/settings/componentSetting';
  import { ResultEnum } from '@/enums/httpEnum';

  const useUserStore = useUserStoreWidthOut();
  const globSetting = useGlobSetting();
  const { uploadUrl } = globSetting;
  const urlPrefix = globSetting.urlPrefix || '';
  const uploadHeaders = reactive({
    Authorization: useUserStore.token,
  });

  const fileList = ref<UploadFileInfo[]>([]);
  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showFileModal = ref(false);
  const showModal = ref(false);
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

  const actionColumn = reactive({
    width: 150,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '下载',
            onClick: handleDown.bind(null, record),
            type: 'default',
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function addTable() {
    showFileModal.value = false;
    showModal.value = true;
    fileList.value = [];
  }

  function addFileTable() {
    showModal.value = false;
    showFileModal.value = true;
    fileList.value = [];
  }

  const loadDataTable = async (res) => {
    return await List({ ...res, ...searchFormRef.value?.formModel });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleDown(record: Recordable) {
    window.open(record.fileUrl);
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          console.log('_res:' + JSON.stringify(_res));
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  //上传之前
  function beforeUpload({ _file }) {
    return true;
  }

  //上传结束
  function finish({ event: Event }) {
    const res = eval('(' + Event.target.response + ')');
    const infoField = componentSetting.upload.apiSetting.infoField;
    const { code } = res;
    const msg = res.msg || res.message || '上传失败';
    const result = res[infoField];

    //成功
    if (code === ResultEnum.SUCCESS) {
      fileList.value.push({
        id: result.id,
        name: result.name,
        status: 'finished',
        type: result.naiveType,
      });

      message.success('上传' + result.name + '成功');
      reloadTable();
    } else {
      message.error(msg);
    }
  }
</script>

<style lang="less" scoped></style>
