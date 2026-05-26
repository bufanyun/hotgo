<template>
  <div>
    <n-card :bordered="false" class="proCard" title="角色管理">
      <n-space vertical :size="12">
        <n-space>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加角色
          </n-button>
        </n-space>

        <n-data-table
          v-if="dataSource.length > 0 || !loading"
          :columns="columns.concat(actionColumn)"
          :data="dataSource"
          :row-key="(row) => row.id"
          :loading="loading"
          :resizeHeightOffset="-20000"
          default-expand-all
        />
      </n-space>
    </n-card>

    <EditRole ref="editRoleRef" @reloadTable="reloadTable" />
    <EditMenuAuth ref="editMenuAuthRef" @reloadTable="reloadTable" />
    <EditDataAuth ref="editDataAuthRef" @reloadTable="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref } from 'vue';
  import { NButton, useDialog, useMessage } from 'naive-ui';
  import { BasicColumn, TableAction } from '@/components/Table';
  import { Delete, getRoleList } from '@/api/system/role';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import EditRole from './editRole.vue';
  import EditMenuAuth from './editMenuAuth.vue';
  import EditDataAuth from './editDataAuth.vue';
  import { newState } from '@/views/permission/role/model';

  const message = useMessage();
  const dialog = useDialog();
  const loading = ref(false);
  const dataSource = ref<any>([]);
  const editRoleRef = ref();
  const editMenuAuthRef = ref();
  const editDataAuthRef = ref();

  const actionColumn = reactive<BasicColumn>({
    width: 200,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction, {
        style: 'primary',
        actions: [
          {
            label: '菜单权限',
            onClick: handleMenuAuth.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            type: 'default',
          },
          {
            label: '数据权限',
            onClick: handleDataAuth.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            type: 'default',
          },
          {
            label: '添加子角色',
            onClick: handleAddSub.bind(null, record),
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
          },
        ],
      });
    },
  });

  function loadDataTable() {
    loading.value = true;
    getRoleList({ pageSize: 100, page: 1 }).then((res) => {
      dataSource.value = res.list ?? [];
      loading.value = false;
    });
  }

  function reloadTable() {
    loadDataTable();
  }

  function addTable() {
    editRoleRef.value.openModal(null);
  }

  function handleEdit(record: Recordable) {
    editRoleRef.value.openModal(record);
  }

  function handleAddSub(record: Recordable) {
    let state = newState(null);
    state.pid = record.id;
    editRoleRef.value.openModal(state);
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
    });
  }

  async function handleMenuAuth(record: Recordable) {
    editMenuAuthRef.value.openModal(record);
  }

  function handleDataAuth(record: Recordable) {
    editDataAuthRef.value.openModal(record);
  }

  onMounted(() => {
    loadDataTable();
  });
</script>

<style lang="less" scoped></style>
