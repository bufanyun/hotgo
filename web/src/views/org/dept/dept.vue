<template>
  <div>
    <n-card :bordered="false" class="proCard" title="部门管理">
      <BasicForm
        ref="searchFormRef"
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <BasicTable
        ref="actionRef"
        openChecked
        :columns="columns"
        :pagination="false"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        :actionColumn="actionColumn"
        :scroll-x="1280"
        :resizeHeightOffset="-10000"
        :cascade="false"
        :expanded-row-keys="expandedKeys"
        @update:expanded-row-keys="updateExpandedKeys"
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="handleOnCheckedRow"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/dept/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加部门
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            class="min-left-space"
            v-if="hasPermission(['/dept/delete'])"
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
            icon-placement="left"
            @click="handleAllExpanded"
            class="min-left-space"
          >
            全部{{ expandedKeys.length ? '收起' : '展开' }}
            <template #icon>
              <div class="flex items-center">
                <n-icon size="14">
                  <AlignLeftOutlined />
                </n-icon>
              </div>
            </template>
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onMounted } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { getDeptList, Delete } from '@/api/org/dept';
  import { PlusOutlined, DeleteOutlined, AlignLeftOutlined } from '@vicons/antd';
  import { columns, schemas, loadOptions, newState, filterIds } from './model';
  import { convertListToTree } from '@/utils/hotgo';
  import Edit from './edit.vue';

  const dialog = useDialog();
  const message = useMessage();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const checkedIds = ref([]);
  const expandedKeys = ref([]);
  const allTreeKeys = ref([]);

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
            auth: ['/dept/edit'],
          },
          {
            label: '添加子部门',
            onClick: handleAdd.bind(null, record),
            auth: ['/dept/edit'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/dept/delete'],
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

  // 加载普通数表数据
  const loadDataTable = async (res = {}) => {
    filterIds.value = [];
    const params = { ...(searchFormRef.value?.formModel ?? {}), ...res, pagination: false };
    const dataSource = await getDeptList(params);
    allTreeKeys.value = expandedKeys.value = dataSource.list.map((item) => item.id);
    dataSource.list = convertListToTree(dataSource.list, 'id');
    filterIds.value = dataSource.ids;
    return dataSource;
  };

  // 更新选中的行
  function handleOnCheckedRow(rowKeys) {
    checkedIds.value = rowKeys;
  }

  // 重新加载表格数据
  function reloadTable() {
    actionRef.value?.reload();
  }

  // 添加数据
  function addTable() {
    editRef.value.openModal(null);
  }

  // 添加树节点下级数据
  function handleAdd(record: Recordable) {
    const state = newState(null);
    state.pid = record.id;
    editRef.value.openModal(state);
  }

  // 编辑数据
  function handleEdit(record: Recordable) {
    editRef.value.openModal(record);
  }

  // 单个删除
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

  // 批量删除
  function handleBatchDelete() {
    if (checkedIds.value.length < 1) {
      message.error('请至少选择一项要删除的数据');
      return;
    }

    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  // 收起/展开全部树节点
  function handleAllExpanded() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = allTreeKeys.value;
    }
  }

  // 更新展开的树节点
  function updateExpandedKeys(openKeys: never[]) {
    expandedKeys.value = openKeys;
  }

  onMounted(() => {
    loadOptions();
  });
</script>

<style lang="less" scoped></style>
