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
            :options="dict.getOptionUnRef('tagTypeOptions')"
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
              v-for="status in dict.getOptionUnRef('sys_normal_disable')"
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
  import { useMessage, useDialog, NTag, SelectRenderTag } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { getDataList, getDictSelect, EditData, DeleteData } from '@/api/dict/dict';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { TypeSelect } from '@/api/sys/config';
  import { findTreeNode } from '@/utils';
  import { cloneDeep } from 'lodash-es';
  import { Option, useDictStore } from '@/store/modules/dict';
  import { State } from '@/views/system/dict/model';

  interface Props {
    checkedId?: number;
  }

  const dict = useDictStore();
  const props = withDefaults(defineProps<Props>(), { checkedId: 0 });
  const typeList = ref<any>([]);
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
      label: '字典标签',
      componentProps: {
        placeholder: '请输入字典标签',
        onInput: (e: any) => {
          console.log(e);
          params.value.label = e;
        },
      },
      rules: [{ message: '请输入字典标签', trigger: ['blur'] }],
    },
    {
      field: 'value',
      component: 'NInput',
      label: '字典键值',
      componentProps: {
        placeholder: '请输入字典键值',
        onInput: (e: any) => {
          console.log(e);
          params.value.value = e;
        },
      },
      rules: [{ message: '请输入字典键值', trigger: ['blur'] }],
    },
  ];

  const renderTag: SelectRenderTag = ({ option }) => {
    return h(
      NTag,
      {
        type: option.listClass as 'success' | 'warning' | 'error' | 'info' | 'primary' | 'default',
      },
      { default: () => option.label }
    );
  };

  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const actionRef = ref();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formParams = ref<State>(new State());
  const options = ref<Option>();
  const typeId = ref(0);

  const params = ref({
    pageSize: 10,
    typeId: props.checkedId,
    label: '',
    value: '',
  });

  const actionColumn = reactive({
    width: 140,
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

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  function addTable() {
    showModal.value = true;
    formParams.value = new State();
    if (typeId.value > 0) {
      formParams.value.typeId = typeId.value;
    }
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

    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        EditData(formParams.value)
          .then((_res) => {
            message.success('操作成功');
            showModal.value = false;
            reloadTable();
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  function handleDelete(record: Recordable) {
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
    });
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = cloneDeep(record) as unknown as State;
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    params.value.label = '';
    params.value.value = '';
    reloadTable();
  }

  async function setDictSelect() {
    const tmp = await getDictSelect({});
    typeList.value = tmp.list;
    if (typeList.value === undefined || typeList.value === null) {
      typeList.value = [];
    }
  }

  function handleUpdateTypeIdValue(value) {
    const row = findTreeNode(typeList.value, value, 'id');
    if (!row) {
      message.error('未找到该节点数据');
      return;
    }
    formParams.value.typeId = value;
    formParams.value.type = row.type;
  }

  function loadOptions() {
    dict.loadOptions(['sys_normal_disable']);
    TypeSelect().then((res) => {
      options.value = res;
    });
  }

  watch(props, (newVal, _oldVal) => {
    params.value.typeId = newVal.checkedId;
    typeId.value = newVal.checkedId;
    actionRef.value.reload();
    setDictSelect();
  });

  onMounted(async () => {
    loadOptions();
    await setDictSelect();
  });
</script>

<style lang="less" scoped></style>
