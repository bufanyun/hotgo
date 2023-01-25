<template>
  <div>
    <n-card :bordered="false" class="proCard" title="访问黑名单">
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
        :scroll-x="1090"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加策略
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
        :title="formParams?.id > 0 ? '编辑策略 #' + formParams.id : '添加策略'"
        style="width: 720px"
      >
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="IP策略" path="ip">
            <n-input type="textarea" placeholder="请输入IP策略" v-model:value="formParams.ip" />
            <template #feedback>
              <p>支持添加IP：如果添加多个IP请用","隔开</p>
              <p>支持添加IP段,如：192.168.0.0/24</p>
              <p>支持添加IP范围,格式如：192.168.1.xx-192.168.1.xx</p>
              <br />
            </template>
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-radio-group v-model:value="formParams.status" name="status">
              <n-radio-button
                v-for="status in blacklistOptions"
                :key="status.value"
                :value="status.value"
                :label="status.label"
              />
            </n-radio-group>
          </n-form-item>

          <n-form-item label="备注" path="remark">
            <n-input type="textarea" placeholder="请输入备注" v-model:value="formParams.remark" />
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref } from 'vue';
  import { NTag, useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Delete, Edit, List, Status } from '@/api/sys/blacklist';
  import { DeleteOutlined, PlusOutlined } from '@vicons/antd';
  import { statusActions, statusOptions } from '@/enums/optionsiEnum';
  import { Dict } from '@/api/dict/dict';
  import { getOptionLabel, getOptionTag } from '@/utils/hotgo';

  const blacklistOptions = [
    {
      value: 1,
      label: '封禁中',
      listClass: 'warning',
    },
    {
      value: 2,
      label: '已解封',
      listClass: 'success',
    },
  ].map((s) => {
    return s;
  });

  const options = ref({
    status: blacklistOptions,
  });

  const columns = [
    {
      title: 'ID',
      key: 'id',
    },
    {
      title: 'IP地址',
      key: 'ip',
    },
    {
      title: '备注',
      key: 'remark',
    },
    {
      title: '状态',
      key: 'status',
      render(row) {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: getOptionTag(options.value.status, row.status),
            bordered: false,
          },
          {
            default: () => getOptionLabel(options.value.status, row.status),
          }
        );
      },
    },

    {
      title: '创建时间',
      key: 'createdAt',
    },
  ];

  const params = ref<any>({
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

  const schemas = ref<FormSchema[]>([
    {
      field: 'ip',
      component: 'NInput',
      label: 'IP地址',
      componentProps: {
        placeholder: '请输入IP地址',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入IP地址', trigger: ['blur'] }],
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '状态',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择类型',
        options: blacklistOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ]);

  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

  const resetFormParams = {
    id: 0,
    ip: '',
    remark: '',
    sort: 0,
    status: 1,
  };
  let formParams = ref<any>(resetFormParams);

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
    formParams.value = resetFormParams;
  }

  const loadDataTable = async (res) => {
    return await List({ ...params.value, ...res, ...searchFormRef.value?.formModel });
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
    batchDeleteDisabled.value = rowKeys.length <= 0;
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
        Edit(formParams.value)
          .then((_res) => {
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
    showModal.value = true;
    formParams.value = record;
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
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
        Delete({ id: checkedIds.value })
          .then((_res) => {
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleSubmit(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function updateStatus(id, status) {
    Status({ id: id, status: status }).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }
</script>

<style lang="less" scoped></style>
