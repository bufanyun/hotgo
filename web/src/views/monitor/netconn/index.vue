<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="在线服务">
        在这里，您可以方便地查看、维护和管理系统中的在线服务和授权许可证
      </n-card>
    </div>

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
        :scroll-x="1280"
      >
        <template #tableTitle>
          <n-button type="info" @click="openGroupModal">
            <template #icon>
              <n-icon>
                <TrademarkOutlined />
              </n-icon>
            </template>
            许可证列表
          </n-button>
        </template>
      </BasicTable>
    </n-card>

    <GroupModal ref="GroupModalRef" />

    <Edit
      @reloadTable="reloadTable"
      @updateShowModal="updateShowModal"
      :showModal="showModal"
      :formParams="formEditParams"
    />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { NetOnlineList, NetOffline } from '@/api/monitor/monitor';
  import { columns } from './columns';
  import { TrademarkOutlined } from '@vicons/antd';
  import GroupModal from './modal/modal.vue';
  import Edit from '@/views/monitor/netconn/modal/edit.vue';
  import { newState, options, State } from '@/views/monitor/netconn/modal/model';
  import { defRangeShortcuts } from '@/utils/dateUtil';

  const message = useMessage();
  const dialog = useDialog();
  const showModal = ref(false);
  const formEditParams = ref<State>(newState(null));
  const actionRef = ref();
  const formParams = ref({});

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
            label: '编辑许可证',
            onClick: handleEditLicense.bind(null, record),
            type: 'primary',
          },
          {
            label: '下线',
            onClick: handleDelete.bind(null, record),
            type: 'error',
          },
        ],
      });
    },
  });

  const schemas: FormSchema[] = [
    {
      field: 'name',
      component: 'NInput',
      label: '应用名称',
      componentProps: {
        placeholder: '请输入应用名称',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'group',
      component: 'NSelect',
      label: '授权分组',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择授权分组',
        options: options.value.group,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'appId',
      component: 'NInput',
      label: 'APPID',
      componentProps: {
        placeholder: '请输入许可证APPID',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'addr',
      component: 'NInput',
      label: '登录地址',
      componentProps: {
        placeholder: '请输入登录地址',
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

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '确认是否要下线该服务？下线操作将中断当前连接，但服务在下线后仍可重新连接',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        NetOffline(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  const loadDataTable = async (res) => {
    return await NetOnlineList({ ...formParams.value, ...res });
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

  const GroupModalRef = ref();
  function openGroupModal() {
    const { openDrawer } = GroupModalRef.value;
    openDrawer();
  }

  function handleEditLicense(record: Recordable) {
    formEditParams.value = newState({ id: record.licenseId } as State);
    updateShowModal(true);
  }

  function updateShowModal(value) {
    showModal.value = value;
  }
</script>

<style lang="less" scoped></style>
