<template>
  <div>
    <n-card :bordered="false" class="proCard">
      <div class="n-layout-page-header">
        <n-card :bordered="false" title="生成演示列表">
          <!--  这里有系统自动生成的CURD表格 -->
        </n-card>
      </div>
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
            v-if="hasPermission(['/demoVar/export'])"
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
    <Edit
      @reloadTable="reloadTable"
      @updateShowModal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
    />
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { Delete, List, Status, Export } from '@/api/curdDemo';
  import { State, columns, schemas, options, newState } from './model';
  import { DeleteOutlined, PlusOutlined, ExportOutlined } from '@vicons/antd';
  import { useRouter } from 'vue-router';
  import { getOptionLabel } from '@/utils/hotgo';
  import Edit from './edit.vue';
  const { hasPermission } = usePermission();
  const router = useRouter();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);
  const formParams = ref<State>();

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

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleView(record: Recordable) {
    router.push({ name: 'curdDemoView', params: { id: record.id } });
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
      onNegativeClick: () => {
        // message.error('取消');
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
          message.success('删除成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
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
