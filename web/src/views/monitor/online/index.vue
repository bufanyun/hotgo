<template>
  <div>
    <n-card :bordered="false" class="proCard" title="在线用户">
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
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { OnlineList, Offline } from '@/api/monitor/monitor';
  import { columns } from './columns';
  import { defRangeShortcuts } from '@/utils/dateUtil';
  import { adaTableScrollX } from '@/utils/hotgo';

  const schemas: FormSchema[] = [
    {
      field: 'userId',
      component: 'NInput',
      label: '用户ID',
      componentProps: {
        placeholder: '请输入用户ID',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'username',
      component: 'NInput',
      label: '用户名',
      componentProps: {
        placeholder: '请输入用户名',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'ip',
      component: 'NInput',
      label: '登录IP',
      componentProps: {
        placeholder: '请输入登录IP',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'firstTime',
      component: 'NDatePicker',
      label: '登录时间',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        shortcuts: defRangeShortcuts(),
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const dialog = useDialog();
  const message = useMessage();
  const actionRef = ref();
  const formParams = ref({});

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
            label: '强制退出',
            onClick: handleDelete.bind(null, record),
            type: 'error',
          },
        ],
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要强制退出该用户？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Offline(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  const loadDataTable = async (res) => {
    return await OnlineList({ ...formParams.value, ...res });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleSubmit(values: Recordable) {
    formParams.value = values;
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    formParams.value = {};
    reloadTable();
  }
</script>

<style lang="less" scoped></style>
