<template>
  <n-card :bordered="false" class="proCard">
    <n-result
      v-show="checkedId <= 0"
      status="info"
      title="提示"
      description="请选择一个想要编辑的省市区"
    />

    <div v-show="checkedId > 0">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>

      <BasicTable
        :columns="listColumns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加数据
          </n-button>
        </template>
      </BasicTable>
    </div>

    <Edit
      @reloadTable="reloadTable"
      @updateShowModal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
      :optionTreeData="optionTreeData"
      :isUpdate="isUpdate"
    />
  </n-card>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, watch, computed } from 'vue';
  import { useMessage, useDialog } from 'naive-ui';
  import { BasicColumn, BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { listColumns, newState, State } from './model';
  import { PlusOutlined } from '@vicons/antd';
  import { getProvincesChildrenList, Delete } from '@/api/apply/provinces';
  import Edit from './edit.vue';
  import { adaTableScrollX } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);

  interface Props {
    checkedId?: number;
    optionTreeData: any;
  }

  const props = withDefaults(defineProps<Props>(), { checkedId: 0, optionTreeData: [] });
  const searchFormRef = ref<any>({});
  const message = useMessage();
  const dialog = useDialog();
  const actionRef = ref();
  const isUpdate = ref(false);
  const showModal = ref(false);
  const formParams = ref<State>(newState(null));
  const params = ref({
    pageSize: 10,
    pid: props.checkedId,
    label: '',
  });

  const schemas: FormSchema[] = [
    {
      field: 'id',
      component: 'NInput',
      label: '地区ID',
      componentProps: {
        placeholder: '请输入地区ID',
        onInput: (e: any) => {
          console.log(e);
          params.value.label = e;
        },
      },
    },
    {
      field: 'title',
      component: 'NInput',
      label: '地区名称',
      componentProps: {
        placeholder: '请输入地区名称',
        onInput: (e: any) => {
          console.log(e);
          params.value.label = e;
        },
      },
    },
  ];

  const actionColumn = reactive<BasicColumn>({
    width: 150,
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
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(listColumns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:1 l:2 xl:2 2xl:2' },
    labelWidth: 80,
    schemas,
  });

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
    formParams.value.pid = props.checkedId;
    isUpdate.value = false;
  }

  const loadDataTable = async (res) => {
    if (props.checkedId <= 0) {
      return [];
    }
    return await getProvincesChildrenList({
      ...{ pid: props.checkedId },
      ...searchFormRef.value?.formModel,
      ...res,
    });
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
  }

  function reloadTable() {
    actionRef.value.reload();
    emit('reloadTable');
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '您确定想删除吗？',
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

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = newState(record as State);
    isUpdate.value = true;
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    params.value.label = '';
    reloadTable();
  }

  watch(props, (_newVal, _oldVal) => {
    if (params.value.pid === _newVal.checkedId) {
      return;
    }
    params.value.pid = _newVal.checkedId;
    formParams.value.pid = Number(_newVal.checkedId);
    if (_newVal.checkedId > 0) {
      reloadTable();
    }
  });

  function updateShowModal(value) {
    showModal.value = value;
  }
</script>

<style lang="less" scoped></style>
