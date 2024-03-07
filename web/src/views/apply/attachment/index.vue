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
        :scroll-x="1280"
        :resizeHeightOffset="-20000"
      >
        <template #tableTitle>
          <n-button type="primary" @click="handleUpload" class="ml-2">
            <template #icon>
              <n-icon>
                <UploadOutlined />
              </n-icon>
            </template>
            上传文件
          </n-button>
          <n-button type="success" @click="handleMultipartUpload" class="ml-2">
            <template #icon>
              <n-icon>
                <FileAddOutlined />
              </n-icon>
            </template>
            上传大文件
          </n-button>
          <n-button type="primary" @click="handleUploadImage" class="ml-2">
            <template #icon>
              <n-icon>
                <FileImageOutlined />
              </n-icon>
            </template>
            上传图片
          </n-button>
          <n-button type="primary" @click="handleUploadDoc" class="ml-2">
            <template #icon>
              <n-icon>
                <FileWordOutlined />
              </n-icon>
            </template>
            上传文档
          </n-button>
          <n-button type="error" @click="batchDelete" :disabled="batchDeleteDisabled" class="ml-2">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>
    </n-card>

    <FileUpload ref="fileUploadRef" :finish-call="handleFinishCall" />
    <FileUpload ref="imageUploadRef" :finish-call="handleFinishCall" upload-type="image" />
    <FileUpload ref="docUploadRef" :finish-call="handleFinishCall" upload-type="doc" />
    <MultipartUpload ref="multipartUploadRef" @onFinish="handleFinishCall" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { Delete, List } from '@/api/apply/attachment';
  import { columns, schemas } from './columns';
  import {
    DeleteOutlined,
    UploadOutlined,
    FileWordOutlined,
    FileImageOutlined,
    FileAddOutlined,
  } from '@vicons/antd';
  import FileUpload from '@/components/FileChooser/src/Upload.vue';
  import MultipartUpload from '@/components/Upload/multipartUpload.vue';
  import { Attachment } from '@/components/FileChooser/src/model';

  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const fileUploadRef = ref();
  const imageUploadRef = ref();
  const docUploadRef = ref();
  const multipartUploadRef = ref();

  const actionColumn = reactive({
    width: 120,
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
    labelWidth: 100,
    schemas,
  });

  function handleUpload() {
    fileUploadRef.value.openModal();
  }

  function handleUploadImage() {
    imageUploadRef.value.openModal();
  }

  function handleUploadDoc() {
    docUploadRef.value.openModal();
  }

  function handleMultipartUpload() {
    multipartUploadRef.value.openModal();
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

  function handleFinishCall(result: Attachment, success: boolean) {
    if (success) {
      reloadTable();
    }
  }
</script>

<style lang="less" scoped></style>
