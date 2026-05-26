<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="通知公告">
        在这里，您可以向平台中的用户发送通知、公告和私信消息
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        @keyup.enter="handleSubmit"
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
        :resizeHeightOffset="-20000"
      >
        <template #tableTitle>
          <n-button
            type="warning"
            @click="addTable(1)"
            class="min-left-space"
            v-if="hasPermission(['/notice/editNotify'])"
          >
            <template #icon>
              <n-icon>
                <NotificationOutlined />
              </n-icon>
            </template>
            发通知
          </n-button>

          <n-button
            type="error"
            @click="addTable(2)"
            class="min-left-space"
            v-if="hasPermission(['/notice/editNotice'])"
          >
            <template #icon>
              <n-icon>
                <BellOutlined />
              </n-icon>
            </template>
            发公告
          </n-button>

          <n-button
            type="info"
            @click="addTable(3)"
            class="min-left-space"
            v-if="hasPermission(['/notice/editLetter'])"
          >
            <template #icon>
              <n-icon>
                <SendOutlined />
              </n-icon>
            </template>
            发私信
          </n-button>

          <n-button
            type="error"
            @click="batchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/notice/delete'])"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit ref="editRef" @reload-table="reloadTable" />
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, onMounted, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { Delete, List, Status } from '@/api/apply/notice';
  import { BellOutlined, DeleteOutlined, NotificationOutlined, SendOutlined } from '@vicons/antd';
  import { adaTableScrollX } from '@/utils/hotgo';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useDictStore } from '@/store/modules/dict';
  import { columns } from './columns';
  import { schemas, loadOptions } from './model';
  import Edit from './edit.vue';

  const { hasPermission } = usePermission();
  const message = useMessage();
  const dict = useDictStore();
  const actionRef = ref();
  const dialog = useDialog();
  const searchFormRef = ref<any>({});
  const editRef = ref();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);

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
            label: '已启用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/notice/status'],
          },
          {
            label: '已禁用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/notice/status'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/notice/edit'],
            type: 'primary',
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/notice/delete'],
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

  function addTable(type) {
    editRef.value.openModal(null, type);
  }

  const loadDataTable = async (res) => {
    return await List({ ...res, ...searchFormRef.value?.formModel });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleEdit(record: Recordable) {
    editRef.value.openModal(record, record.type);
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

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
    });
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  function handleStatus(record: Recordable, status: number) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  onMounted(() => {
    loadOptions();
  });
</script>

<style lang="less" scoped></style>
