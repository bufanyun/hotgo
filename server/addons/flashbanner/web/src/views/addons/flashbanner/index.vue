<template>
  <div>
    <n-card :bordered="false" class="proCard" title="轮播图管理">
      <BasicForm
        @register="register"
        @submit="handleQuery"
        @reset="resetForm"
        @keyup.enter="handleQuery"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable
        :openChecked="true"
        :columns="columns"
        :actionColumn="actionColumn"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="tableRef"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="handleAdd"
            class="min-left-space"
            v-if="hasPermission(['/member/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
        </template>
      </BasicTable>
      <BasicEdit ref="editRef" @on-refresh="onRefresh" />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { h, onMounted, ref, computed, watch, reactive } from 'vue';
import { register, defaultColumns } from './components/model';
import BasicEdit from './components/Edit.vue';
import { BasicForm } from '@/components/Form/index';
import { BasicTable, TableAction } from '@/components/Table';
import { usePermission } from '@/hooks/web/usePermission';
import { List, Delete } from '@/api/addons/flashbanner/index';
import { PlusOutlined } from '@vicons/antd';
import { useDialog, useMessage } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { adaTableScrollX } from '@/utils/hotgo';

interface TableActionState {
  reload: () => void;
}

const { hasPermission } = usePermission();
const dialog = useDialog();
const message = useMessage();
const formParams = ref({});
const tableRef = ref<TableActionState>();
const editRef = ref();
const columns = ref(defaultColumns);
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
          auth: ['/member/delete'],
        },
        {
          label: '删除',
          onClick: handleDelete.bind(null, record),
          auth: ['/member/delete'],
        },
      ],
    });
  },
});

const scrollX = computed(() => {
  return adaTableScrollX(defaultColumns, actionColumn.width);
});

const loadDataTable = async (res) => {
  return await List({ ...formParams.value, ...res });
};

// 刷新table
const onRefresh = () => {
  tableRef.value?.reload();
};

// 重置查询框
const resetForm = () => {
  formParams.value = {};
  onRefresh();
};

// 查询
const handleQuery = (e: any) => {
  formParams.value = { ...e };
  onRefresh();
};

// 添加
const handleAdd = () => {
  if (editRef.value) {
    editRef.value.showModal = true;
  }
};

// 编辑
const handleEdit = (record: Recordable) => {
  if (editRef.value) {
    editRef.value.showModal = true;
    editRef.value.formParams = cloneDeep(record);
  }
}

// 删除
const handleDelete = (record: Recordable) => {
  dialog.warning({
    title: '警告',
    content: '你确定要删除？',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      Delete({ id: record.id }).then((_res) => {
        message.success('操作成功');
        onRefresh();
      });
    },
  });
}
</script>

<style lang="less"></style>
