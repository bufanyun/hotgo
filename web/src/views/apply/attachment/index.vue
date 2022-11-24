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
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <UploadOutlined />
              </n-icon>
            </template>
            上传附件
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
        title="上传附件"
      >
        <n-upload
          multiple
          directory-dnd
          :action="`${uploadUrl}/admin/upload/image`"
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
                <archive-icon />
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
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Delete, Edit, List, Status } from '@/api/apply/attachment';
  import { columns } from './columns';
  import { DeleteOutlined, UploadOutlined } from '@vicons/antd';
  import { statusActions, statusOptions } from '@/enums/optionsiEnum';
  import { useGlobSetting } from '@/hooks/setting';
  import { useUserStoreWidthOut } from '@/store/modules/user';
  import componentSetting from '@/settings/componentSetting';
  import { ResultEnum } from '@/enums/httpEnum';

  const useUserStore = useUserStoreWidthOut();

  const globSetting = useGlobSetting();

  const { uploadUrl } = globSetting;

  const uploadHeaders = reactive({
    Authorization: useUserStore.token,
  });

  const fileList = ref<UploadFileInfo[]>([
    // {
    //   id: 'c',
    //   name: '图片.png',
    //   status: 'finished',
    //   url: 'https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg',
    // },
  ]);

  const driveOptions = [
    {
      value: 'local',
      label: '本地',
    },
  ].map((s) => {
    return s;
  });

  const params = ref({
    pageSize: 10,
    title: '',
    content: '',
    status: null,
  });

  const rules = {
    title: {
      // required: true,
      trigger: ['blur', 'input'],
      message: '请输入标题',
    },
  };

  const schemas: FormSchema[] = [
    {
      field: 'member_id',
      component: 'NInput',
      label: '用户ID',
      componentProps: {
        placeholder: '请输入用户ID',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入用户ID', trigger: ['blur'] }],
    },
    {
      field: 'drive',
      component: 'NSelect',
      label: '驱动',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择驱动',
        options: driveOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '状态',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择类型',
        options: statusOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref({});
  const formRef = ref({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

  const resetFormParams = {
    basicLogo: '',
    id: 0,
    title: '',
    name: '',
    type: 1,
    receiver: '',
    remark: '',
    sort: 0,
    status: 1,
    created_at: '',
    updated_at: '',
  };
  let formParams = ref(resetFormParams);

  const actionColumn = reactive({
    width: 220,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
        dropDownActions: statusActions,
        select: (key) => {
          updateStatus(record.id, key);
        },
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function addTable() {
    showModal.value = true;
    fileList.value = [];
  }

  const loadDataTable = async (res) => {
    return await List({ ...params.value, ...res, ...searchFormRef.value.formModel });
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
    if (rowKeys.length > 0) {
      batchDeleteDisabled.value = false;
    } else {
      batchDeleteDisabled.value = true;
    }

    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        console.log('formParams:' + JSON.stringify(formParams.value));
        Edit(formParams.value)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            setTimeout(() => {
              showModal.value = false;
              reloadTable();
              formParams.value = ref(resetFormParams);
            });
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    showModal.value = true;
    formParams.value = record;
  }

  function handleDelete(record: Recordable) {
    console.log('点击了删除', record);
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete(record)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value })
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    params.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function updateStatus(id, status) {
    Status({ id: id, status: status })
      .then((_res) => {
        console.log('_res:' + JSON.stringify(_res));
        message.success('操作成功');
        setTimeout(() => {
          reloadTable({});
        });
      })
      .catch((e: Error) => {
        message.error(e.message ?? '操作失败');
      });
  }

  //上传之前
  function beforeUpload({ file }) {
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
