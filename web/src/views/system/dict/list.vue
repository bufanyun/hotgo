<template>
  <n-card :bordered="false" class="proCard">
    <BasicForm @register="register" @submit="handleSubmit" @reset="handleReset">
      <template #statusSlot="{ model, field }">
        <n-input v-model:value="model[field]" />
      </template>
    </BasicForm>

    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      @update:checked-row-keys="onCheckedRow"
      :scroll-x="1090"
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

    <n-modal
      v-model:show="showModal"
      :show-icon="false"
      preset="dialog"
      :title="formParams?.id > 0 ? '编辑数据' : '添加数据'"
    >
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-form-item label="字典类型" path="typeId">
          <n-tree-select
            :options="typeList"
            :default-value="formParams.typeId"
            :default-expand-all="true"
            @update:value="handleUpdateTypeIdValue"
          />
        </n-form-item>
        <n-form-item label="标签" path="label">
          <n-input placeholder="请输入标签名称" v-model:value="formParams.label" />
        </n-form-item>
        <n-form-item label="标签样式" path="listClass">
          <n-select
            :render-tag="renderTag"
            v-model:value="formParams.listClass"
            :options="labelOptions"
          />
        </n-form-item>
        <n-form-item label="字典键值" path="value">
          <n-input placeholder="请输入键值" v-model:value="formParams.value" />
        </n-form-item>
        <n-form-item label="键值类型" path="valueType">
          <n-select v-model:value="formParams.valueType" :options="options" />
        </n-form-item>
        <n-form-item label="排序" path="sort">
          <n-input-number placeholder="请输入" v-model:value="formParams.sort" />
        </n-form-item>
        <n-form-item label="状态" path="status">
          <n-radio-group v-model:value="formParams.status" name="status">
            <n-radio-button
              v-for="status in statusOptions"
              :key="status.value"
              :value="status.value"
              :label="status.label"
            />
          </n-radio-group>
        </n-form-item>

        <n-form-item label="备注" path="remark">
          <n-input type="textarea" placeholder="请输入备注" v-model:value="formParams.remark" />
        </n-form-item>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showModal = false)">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </n-card>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, watch, onMounted } from 'vue';
  import { TreeSelectOption, useMessage, useDialog, NTag, SelectRenderTag } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { getDataList, getDictSelect, EditData, DeleteData } from '@/api/dict/dict';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { statusOptions } from '@/enums/optionsiEnum';
  import { TypeSelect } from '@/api/sys/config';
  import { Option } from '@/utils/hotgo';
  const options = ref<Option>();
  interface Props {
    checkedId?: number;
  }

  const props = withDefaults(defineProps<Props>(), { checkedId: 0 });
  const typeList = ref([]);
  const rules = {
    label: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入标签名称',
    },
    value: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入键值',
    },
  };

  const schemas: FormSchema[] = [
    {
      field: 'label',
      component: 'NInput',
      label: '标签',
      componentProps: {
        placeholder: '请输入标签名称',
        onInput: (e: any) => {
          console.log(e);
          params.value.label = e;
        },
      },
      rules: [{ message: '请输入字典标签名称', trigger: ['blur'] }],
    },
  ];

  const renderTag: SelectRenderTag = ({ option }) => {
    return h(
      NTag,
      {
        type: option.type as 'success' | 'warning' | 'error' | 'info' | 'primary' | 'default',
      },
      { default: () => option.label }
    );
  };

  const labelOptions = ref([
    {
      label: '绿色',
      value: 'success',
      type: 'success',
    },
    {
      label: '橙色',
      value: 'warning',
      type: 'warning',
    },
    {
      label: '红色',
      value: 'error',
      type: 'error',
    },
    {
      label: '蓝色',
      value: 'info',
      type: 'info',
    },
    {
      label: '灰色',
      value: 'default',
      type: 'default',
    },
    {
      label: '主题色',
      value: 'primary',
      type: 'primary',
    },
  ]);

  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const actionRef = ref();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formParams = ref<any>({ typeId: 0 });
  const params = ref({
    pageSize: 10,
    typeId: props.checkedId,
    label: '',
  });

  const actionColumn = reactive({
    width: 220,
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
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
      });
    },
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:1 l:2 xl:2 2xl:2' },
    labelWidth: 80,
    schemas,
  });

  function addTable() {
    showModal.value = true;
    formParams.value = {
      typeId: props.checkedId,
      label: '',
      value: '',
      listClass: 'default',
      valueType: 'string',
      sort: 0,
      status: 1,
      remark: '',
    };
  }

  const loadDataTable = async (res) => {
    return await getDataList({ ...params.value, ...res });
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        EditData(formParams.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showModal.value = false;
            reloadTable();
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleDelete(record: Recordable) {
    console.log('点击了删除', record);
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteData(record).then((_res) => {
          message.success('操作成功');
          reloadTable();
        });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = record;
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    params.value.label = '';
    reloadTable();
  }

  watch(props, (_newVal, _oldVal) => {
    params.value.typeId = _newVal.checkedId;
    formParams.value.typeId = _newVal.checkedId;
    actionRef.value.reload();
    setDictSelect();
  });

  async function setDictSelect() {
    typeList.value = await getDictSelect({});
    if (typeList.value === undefined || typeList.value === null) {
      typeList.value = [];
    }
  }

  function handleUpdateTypeIdValue(
    value: string | number | Array<string | number> | null,
    _option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    formParams.value.typeId = value;
  }

  async function loadOptions() {
    options.value = await TypeSelect();
  }

  onMounted(async () => {
    await setDictSelect();
    await loadOptions();
  });
</script>

<style lang="less" scoped></style>
