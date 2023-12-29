<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="生成演示">
        <!--  这是由系统生成的CURD表格，你可以将此行注释改为表格的描述 -->
      </n-card>
    </div>
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
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
        :resizeHeightOffset="-10000"
        size="small"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/curdDemo/edit'])"
          >
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
            v-if="hasPermission(['/curdDemo/delete'])"
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
            v-if="hasPermission(['/curdDemo/export'])"
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
    <Edit @reloadTable="reloadTable" ref="editRef" />
    <View ref="viewRef" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export, Delete, Status } from '@/api/curdDemo';
  import { columns, schemas, options } from './model';
  import { PlusOutlined, ExportOutlined, DeleteOutlined } from '@vicons/antd';
  import { getOptionLabel } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import View from './view.vue';

  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const viewRef = ref();
  const editRef = ref();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

  const actionColumn = reactive({
    width: 300,
    title: '操作',
    key: 'action',
    // fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/curdDemo/edit'],
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/curdDemo/status'],
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/curdDemo/status'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/curdDemo/delete'],
          },
        ],
        dropDownActions: [
          {
            label: '查看详情',
            key: 'view',
            auth: ['/curdDemo/view'],
          },
        ],
        select: (key) => {
          if (key === 'view') {
            return handleView(record);
          }
        },
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  const loadDataTable = async (res) => {
    return await List({ ...searchFormRef.value?.formModel, ...res });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function addTable() {
    editRef.value.openModal(null);
  }

  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  function handleView(record: Recordable) {
    viewRef.value.openModal(record);
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
      message.success('设为' + getOptionLabel(options.value.sys_normal_disable, status) + '成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }
</script>

<style lang="less" scoped></style>