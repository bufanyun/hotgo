<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="服务日志">
        在这里开发者可以快速定位服务端在运行时产生的重要日志，方便排查系统异常和日常运维工作
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
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1800"
        :resizeHeightOffset="-10000"
        size="small"
      >
        <template #tableTitle>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/serveLog/delete'])"
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
            v-if="hasPermission(['/serveLog/delete'])"
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

      <n-modal v-model:show="showModal" :show-icon="false" preset="dialog" style="width: 920px">
        <n-card
          :bordered="false"
          title="日志内容"
          class="proCard mt-4"
          size="small"
          :segmented="{ content: true }"
        >
          <n-alert type="error" :show-icon="false">
            {{ preview?.content }}
          </n-alert>
        </n-card>

        <n-card
          :bordered="false"
          class="proCard mt-4"
          size="small"
          :segmented="{ content: true }"
          title="堆栈打印"
        >
          <JsonViewer
            :value="JSON.parse(preview?.stack)"
            :expand-depth="10"
            sort
            style="width: 100%; min-width: 3.125rem"
          />
        </n-card>
        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">关闭</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { List, Export, Delete } from '@/api/serveLog';
  import { columns, schemas } from './model';
  import { ExportOutlined, DeleteOutlined } from '@vicons/antd';
  import { useRouter } from 'vue-router';
  import { JsonViewer } from 'vue3-json-viewer';
  import 'vue3-json-viewer/dist/index.css';

  const { hasPermission } = usePermission();
  const router = useRouter();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);

  const actionColumn = reactive({
    width: 180,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '访问日志',
            onClick: handleView.bind(null, record),
            ifShow: record.sysLogId > 0,
            type: 'default',
          },
          {
            label: '堆栈',
            onClick: handleStack.bind(null, record),
            type: 'primary',
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/serveLog/delete'],
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

  const preview = ref<Recordable>();
  function handleStack(record: Recordable) {
    showModal.value = true;
    preview.value = record;
  }

  function handleView(record: Recordable) {
    router.push({ name: 'log_view', params: { id: record.sysLogId } });
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
</script>

<style lang="less" scoped></style>
