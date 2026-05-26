<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <n-result
        v-show="checkedId <= 0"
        status="info"
        title="提示"
        description="请选择一个想要编辑的树"
      />

      <div v-show="checkedId > 0">
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
          :checked-row-keys="checkedIds"
          @update:checked-row-keys="onCheckedRow"
          :scroll-x="scrollX"
          :resizeHeightOffset="-10000"
          size="small"
          @update:sorter="handleUpdateSorter"
        >
          <template #tableTitle>
            <n-button type="primary" @click="addTable" class="min-left-space">
              <template #icon>
                <n-icon>
                  <PlusOutlined />
                </n-icon>
              </template>
              添加
            </n-button>
            <n-button
              type="error"
              @click="handleBatchDelete"
              :disabled="batchDeleteDisabled"
              class="min-left-space"
            >
              <template #icon>
                <n-icon>
                  <DeleteOutlined />
                </n-icon>
              </template>
              批量删除
            </n-button>
            <n-button type="primary" @click="handleExport" class="min-left-space">
              <template #icon>
                <n-icon>
                  <ExportOutlined />
                </n-icon>
              </template>
              导出
            </n-button>
          </template>
        </BasicTable>
      </div>
    </n-card>

    <Edit
      @reload-table="reloadTable"
      @update-show-modal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
      :optionTreeData="optionTreeData"
    />
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, reactive, ref, watch } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { useSorter } from '@/hooks/common';
  import { Delete, List, Status, Export } from '@/api/addons/hgexample/treeTable';
  import { State, columns, schemas, newState } from './model';
  import { DeleteOutlined, PlusOutlined, ExportOutlined } from '@vicons/antd';
  import { adaTableScrollX } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import { useDictStore } from '@/store/modules/dict';

  interface Props {
    checkedId?: number;
    optionTreeData: any;
  }

  const props = withDefaults(defineProps<Props>(), { checkedId: 0, optionTreeData: [] });
  const emit = defineEmits(['reloadTable']);
  const dict = useDictStore();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);
  const formParams = ref<State>();
  const actionRef = ref();
  const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);
  const pid = ref(0);

  const actionColumn = reactive({
    width: 200,
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
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
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
    if (pid.value <= 0) {
      return [];
    }
    return await List({
      ...searchFormRef.value?.formModel,
      ...{ sorters: sortStatesRef.value },
      ...{ pid: pid.value },
      ...res,
    });
  };

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
    formParams.value.pid = pid.value;
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    emit('reloadTable');
    actionRef.value.reload();
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = newState(record as State);
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
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          batchDeleteDisabled.value = true;
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }

  function handleStatus(record: Recordable, status: number) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('设为' + dict.getLabel('sys_normal_disable', status) + '成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  watch(props, (_newVal, _oldVal) => {
    if (pid.value === _newVal.checkedId) {
      return;
    }

    pid.value = Number(_newVal.checkedId);
    if (_newVal.checkedId > 0) {
      reloadTable();
    }
  });
</script>

<style lang="less" scoped></style>
