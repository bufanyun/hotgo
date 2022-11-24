<template>
  <n-card :bordered="false" class="proCard">
    <BasicForm @register="register" @submit="handleSubmit" @reset="handleReset">
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
        <n-button type="error" @click="batchDelete" :disabled="batchDeleteDisabled">
          <template #icon>
            <n-icon>
              <DeleteOutlined />
            </n-icon>
          </template>
          批量删除
        </n-button>
      </template>

      <template #toolbar>
        <n-button type="primary" @click="reloadTable">sms-log刷新数据</n-button>
      </template>
    </BasicTable>
  </n-card>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { getLogList, Delete } from '@/api/log/log';
  import { columns } from './columns';
  import { useRouter } from 'vue-router';
  import { DeleteOutlined } from '@vicons/antd';

  const dialog = useDialog();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

  const schemas: FormSchema[] = [
    {
      field: 'member_id',
      component: 'NInput',
      label: '操作人员',
      componentProps: {
        placeholder: '请输入操作人员ID',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'url',
      component: 'NInput',
      label: '访问路径',
      componentProps: {
        placeholder: '请输入手机访问路径',
        onInput: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'ip',
      component: 'NInput',
      label: '访问IP',
      componentProps: {
        placeholder: '请输入IP地址',
        onInput: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'method',
      component: 'NSelect',
      label: '请求方式',
      componentProps: {
        placeholder: '请选择请求方式',
        options: [
          {
            label: 'GET',
            value: 'GET',
          },
          {
            label: 'POST',
            value: 'POST',
          },
        ],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'created_at',
      component: 'NDatePicker',
      label: '访问时间',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        // defaultValue: [new Date() - 86400000 * 30, new Date()],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'take_up_time',
      component: 'NSelect',
      label: '请求耗时',
      componentProps: {
        placeholder: '请选择请求耗时',
        options: [
          {
            label: '50ms内',
            value: '50',
          },
          {
            label: '100ms内',
            value: '100',
          },
          {
            label: '200ms内',
            value: '200',
          },
          {
            label: '500ms内',
            value: '500',
          },
        ],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'error_code',
      component: 'NSelect',
      label: '状态码',
      componentProps: {
        placeholder: '请选择状态码',
        options: [
          {
            label: '0 成功',
            value: '0',
          },
          {
            label: '-1 失败',
            value: '-1',
          },
        ],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const router = useRouter();
  const formRef: any = ref(null);
  const message = useMessage();
  const actionRef = ref();
  const formParams = ref({});

  const params = ref({
    pageSize: 10,
  });

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
            label: '查看详情',
            onClick: handleEdit.bind(null, record),
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

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
    if (rowKeys.length > 0) {
      batchDeleteDisabled.value = false;
    } else {
      batchDeleteDisabled.value = true;
    }

    checkedIds.value = rowKeys;
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

  const loadDataTable = async (res) => {
    return await getLogList({ ...formParams.value, ...params.value, ...res });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    router.push({ name: 'sms_view', params: { id: record.id } });
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    formParams.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    console.log(values);
    formParams.value = {};
    reloadTable();
  }
</script>

<style lang="less" scoped></style>
