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
    />
  </n-card>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { OnlineList, Offline } from '@/api/monitor/monitor';
  import { columns } from './columns';
  import { useRouter } from 'vue-router';

  const dialog = useDialog();
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
          // {
          //   label: '查看详情',
          //   onClick: handleEdit.bind(null, record),
          // },
          {
            label: '强制退出',
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

  function handleDelete(record: Recordable) {
    console.log('点击了删除', record);
    dialog.warning({
      title: '警告',
      content: '你确定要强制退出该用户？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Offline(record)
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

  const loadDataTable = async (res) => {
    return await OnlineList({ ...formParams.value, ...params.value, ...res });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    router.push({ name: 'serve_log_view', params: { id: record.id } });
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
