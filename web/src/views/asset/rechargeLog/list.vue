<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <BasicForm
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
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
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      >
        <template #tableTitle>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/order/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>

          <n-button
            type="primary"
            @click="handleExport"
            class="min-left-space"
            v-if="hasPermission(['/order/export'])"
          >
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>

    <ApplyRefund
      @reload-table="reloadTable"
      @update-show-modal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
    />

    <AcceptRefund
      @reload-table="reloadTable"
      @update-show-modal="updateAcceptShowModal"
      :showModal="showAcceptModal"
      :formParams="formParams"
    />
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export, Delete } from '@/api/order';
  import { State, columns, schemas, newState } from './model';
  import { ExportOutlined, DeleteOutlined } from '@vicons/antd';
  import ApplyRefund from './applyRefund.vue';
  import AcceptRefund from './acceptRefund.vue';
  import { adaTableScrollX } from '@/utils/hotgo';

  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '-1',
  });
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);
  const showAcceptModal = ref(false);
  const formParams = ref<State>();

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
            type: 'warning',
            label: '受理退款',
            onClick: handleAcceptRefund.bind(null, record),
            auth: ['/order/acceptRefund'],
            ifShow: () => {
              return record.status == 6;
            },
          },
          {
            type: 'default',
            label: '申请退款',
            onClick: handleApplyRefund.bind(null, record),
            auth: ['/order/applyRefund'],
            ifShow: () => {
              return record.status == 4;
            },
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/order/delete'],
            ifShow: () => {
              return record.status == 5;
            },
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

  const loadDataTable = async (res) => {
    return await List({ ...searchFormRef.value?.formModel, ...res, ...{ status: props.type } });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  function handleBatchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？只有已关闭的订单才能被删除',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  function handleApplyRefund(record: Recordable) {
    showModal.value = true;
    formParams.value = newState(record as State);
  }

  function updateAcceptShowModal(value) {
    showAcceptModal.value = value;
  }

  function handleAcceptRefund(record: Recordable) {
    showAcceptModal.value = true;
    formParams.value = newState(record as State);
  }

  defineExpose({
    reloadTable,
  });
</script>

<style lang="less" scoped></style>
