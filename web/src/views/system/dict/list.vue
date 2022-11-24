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
          新建
        </n-button>
      </template>
    </BasicTable>

    <n-modal v-model:show="showModal" :show-icon="false" preset="dialog" title="新建">
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
        <n-form-item label="字典键值" path="value">
          <n-input placeholder="请输入键值" v-model:value="formParams.value" />
        </n-form-item>
        <n-form-item label="表格回显" path="listClass">
          <n-input placeholder="请输入表格回显样式" v-model:value="formParams.listClass" />
        </n-form-item>
        <n-form-item label="排序" path="sort">
          <n-input placeholder="请输入" v-model:value="formParams.sort" />
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
  import { TreeSelectOption, useMessage, useDialog } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { getDataList, getDictSelect, EditData, DeleteData } from '@/api/dict/dict';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { statusOptions } from '@/enums/optionsiEnum';

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
      // labelMessage: '请输入字典标签名称',
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

  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const actionRef = ref();
  const showModal = ref(false);
  const formBtnLoading = ref(false);

  const resetFormParams = {
    typeId: props.checkedId,
    label: '',
    value: '',
    listClass: '',
    sort: 0,
    status: 1,
    remark: '',
  };
  const formParams = ref(resetFormParams);

  const params = ref({
    pageSize: 10,
    typeId: props.checkedId,
    label: '',
  });

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
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
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
    formParams.value = resetFormParams;
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
        console.log('formParams:' + JSON.stringify(formParams.value));
        EditData(formParams.value)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            setTimeout(() => {
              showModal.value = false;
              reloadTable();
              formParams.value = ref(resetFormParams);
            });
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
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
      negativeText: '不确定',
      onPositiveClick: () => {
        DeleteData(record)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    showModal.value = true;
    formParams.value = record;
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    reloadTable();
  }

  function handleReset(values: Recordable) {
    console.log(values);
    params.value.label = '';
    reloadTable();
  }

  watch(props, (_newVal, _oldVal) => {
    console.log('_newVal:' + JSON.stringify(_newVal));
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
    option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    console.log(value, option);

    formParams.value.typeId = value;
  }

  onMounted(() => {
    setDictSelect();
  });
</script>

<style lang="less" scoped></style>
