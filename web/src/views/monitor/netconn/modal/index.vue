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
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      >
        <template #tableTitle>
          <n-button
            type="primary"
            @click="addTable"
            class="min-left-space"
            v-if="hasPermission(['/serveLicense/edit'])"
          >
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加许可证
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
            v-if="hasPermission(['/serveLicense/delete'])"
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
            v-if="hasPermission(['/serveLicense/delete'])"
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
      @reload-table="reloadTable"
      @update-show-modal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
    />

    <n-modal
      v-model:show="showRoutesModal"
      :show-icon="false"
      preset="dialog"
      :title="'分配路由 #' + formParams?.id"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-alert :show-icon="false" type="info">
        如果许可证未分配任何路由，则客户端可以访问所有服务路由接口
      </n-alert>
      <n-form
        :model="formParams"
        ref="formRef"
        label-placement="left"
        :label-width="100"
        class="py-4"
      >
        <n-transfer
          ref="transfer"
          v-model:value="formParams.routes"
          virtual-scroll
          :options="routesOptions"
          source-filterable
          :render-source-label="renderLabel"
          :render-target-label="renderLabel"
        />
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, onMounted, reactive, ref } from 'vue';
  import { useDialog, useMessage, NTag } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { usePermission } from '@/hooks/web/usePermission';
  import { Delete, Export, List, Status, AssignRouter } from '@/api/serveLicense';
  import { columns, newState, loadOptions, schemas, State } from './model';
  import { DeleteOutlined, ExportOutlined, PlusOutlined } from '@vicons/antd';
  import { adaModalWidth, adaTableScrollX } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import { NetOption } from '@/api/monitor/monitor';
  import { useDictStore } from '@/store/modules/dict';

  const dict = useDictStore();
  const { hasPermission } = usePermission();
  const actionRef = ref();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);
  const formParams = ref<State>(newState(null));
  const showRoutesModal = ref(false);
  const formBtnLoading = ref(false);
  const formRef = ref<any>({});
  const routesOptions = ref([]);
  const dialogWidth = computed(() => {
    return adaModalWidth();
  });

  const actionColumn = reactive({
    width: 300,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '分配路由',
            onClick: handleAssignRouter.bind(null, record),
            auth: ['/serveLicense/assignRouter'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            auth: ['/serveLicense/edit'],
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
            auth: ['/serveLicense/status'],
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
            auth: ['/serveLicense/status'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            auth: ['/serveLicense/delete'],
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

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        AssignRouter(formParams.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showRoutesModal.value = false;
            reloadTable();
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function closeForm() {
    showRoutesModal.value = false;
  }

  function handleAssignRouter(record: Recordable) {
    showRoutesModal.value = true;
    formParams.value = newState(record as State);
  }

  function renderLabel({ option }) {
    return h(
      'div',
      {
        style: {
          display: 'flex',
          margin: '6px 0',
        },
      },
      {
        default: () => [
          h(
            NTag,
            {
              style: {
                marginRight: '6px',
              },
              type: option.isRPC ? 'success' : 'info',
              bordered: false,
            },
            {
              default: () => (option.isRPC ? 'RPC' : 'TCP'),
            }
          ),
          h(
            'div',
            {
              style: {
                display: 'flex',
                marginLeft: '6px',
                alignSelf: 'center',
              },
            },
            { default: () => option.label }
          ),
        ],
      }
    );
  }

  function loadRoutesOptions() {
    NetOption().then((res) => {
      routesOptions.value = res.routes;
    });
  }

  onMounted(async () => {
    loadOptions();
    loadRoutesOptions();
  });
</script>

<style lang="less" scoped></style>
