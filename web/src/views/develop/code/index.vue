<template>
  <div>
    <n-card :bordered="false" class="proCard" title="代码生成">
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
            立即生成
          </n-button>
          <n-button
            type="error"
            @click="batchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
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

      <n-modal
        v-model:show="showModal"
        :mask-closable="false"
        :show-icon="false"
        preset="dialog"
        title="立即生成"
        :style="{
          width: dialogWidth,
        }"
      >
        <!--      <n-alert :show-icon="false" type="info">-->
        <!--        注意：！-->
        <!--      </n-alert>-->
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="生成类型" path="genType">
            <n-select
              placeholder="请选择"
              :options="selectList.genType"
              v-model:value="formParams.genType"
              :on-update:value="onUpdateValueGenType"
            />
          </n-form-item>

          <n-form-item
            label="数据库"
            path="dbName"
            v-if="formParams.genType >= 10 && formParams.genType < 20"
          >
            <n-select
              placeholder="请选择"
              :options="selectList.db"
              v-model:value="formParams.dbName"
              @update:value="handleDbUpdateValue"
            />
          </n-form-item>
          <n-form-item
            label="数据库表"
            path="tableName"
            v-if="formParams.genType >= 10 && formParams.genType < 20"
          >
            <n-select
              filterable
              tag
              :loading="tablesLoading"
              placeholder="请选择"
              :options="selectList.tables"
              v-model:value="formParams.tableName"
              @update:value="handleTableUpdateValue"
              :disabled="formParams.dbName === ''"
            />
          </n-form-item>

          <n-form-item label="生成模板" path="genTemplate">
            <n-select
              placeholder="请选择"
              :options="genTemplateOptions"
              v-model:value="formParams.genTemplate"
              :onFocus="onFocusGenTemplate"
              @update:value="handleGenTemplateUpdateValue"
            />
          </n-form-item>

          <n-form-item label="选择插件" path="addonName" v-show="selectAddon">
            <n-select
              placeholder="请选择"
              :options="selectList.addons"
              v-model:value="formParams.addonName"
            />
          </n-form-item>

          <n-form-item
            label="菜单名称"
            path="tableComment"
            v-show="formParams.genType >= 10 && formParams.genType < 20"
          >
            <n-input placeholder="请输入" v-model:value="formParams.tableComment" />
          </n-form-item>

          <n-form-item label="实体命名" path="varName">
            <n-input placeholder="请输入" v-model:value="formParams.varName" />
          </n-form-item>
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">生成配置</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { h, onBeforeMount, reactive, ref } from 'vue';
  import { NTag, useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { List, Delete, Edit, Selects, TableSelect } from '@/api/develop/code';
  import { useRouter } from 'vue-router';
  import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
  import { newState } from '@/views/develop/code/components/model';

  const selectList = ref({
    db: [],
    genType: [],
    status: [],
    tables: [],
    formMode: [],
    formRole: [],
    dictMode: [],
    whereMode: [],
    addons: [],
  });
  const columns = [
    {
      title: 'ID',
      key: 'id',
      width: 60,
    },
    {
      title: '模板',
      key: 'genTemplate',
      render(row) {
        return h('p', { id: 'app' }, [
          h('div', {
            innerHTML: '<div><p>' + row.genTemplateGroup + '</p></div>',
          }),
          h('div', {
            innerHTML:
              '<div><p>' + getOptionLabel(selectList.value.genType, row.genType) + '</p></div>',
          }),
        ]);
      },
      width: 120,
    },
    {
      title: '数据库',
      key: 'dbName',
      width: 280,
      render(row) {
        return h('p', { id: 'app' }, [
          h('div', {
            innerHTML: '<div><p>' + row.dbName + '</p></div>',
          }),
          h('div', {
            innerHTML: '<div><p>' + row.tableName + '</p></div>',
          }),
        ]);
      },
    },
    {
      title: '实体命名',
      key: 'varName',
      render(row) {
        return row.varName;
      },
      width: 140,
    },
    {
      title: '生成名称',
      key: 'tableComment',
      width: 140,
    },
    {
      title: '生成状态',
      key: 'status',
      render(row) {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: row.status == 1 ? 'success' : 'warning',
            bordered: false,
          },
          {
            default: () => getOptionLabel(selectList.value.status, row.status),
          }
        );
      },
      width: 150,
    },
    {
      title: '创建时间',
      key: 'createdAt',
      width: 180,
    },
  ];

  const dialog = useDialog();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const searchFormRef = ref<any>();

  const schemas = ref<FormSchema[]>([
    {
      field: 'genType',
      component: 'NSelect',
      label: '生成类型',
      componentProps: {
        placeholder: '请选择生成类型',
        options: [],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'varName',
      component: 'NInput',
      label: '实体命名',
      componentProps: {
        placeholder: '请输入实体命名',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ trigger: ['blur'] }],
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '生成状态',
      componentProps: {
        placeholder: '请选择状态码',
        options: [],
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ]);

  const router = useRouter();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const actionRef = ref();
  const formParams = ref<any>();

  const rules = {
    varName: {
      required: true,
      trigger: ['blur', 'input'],
      message: '实体命名不能为空,首字母大写',
    },
  };

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
            label: '生成配置',
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

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
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

  const loadDataTable = async (res) => {
    mapWidth();
    return await List({ ...res, ...searchFormRef.value?.formModel });
  };

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleEdit(record: Recordable) {
    router.push({ name: 'develop_code_deploy', params: { id: record.id } });
  }

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formParams.value).then((res) => {
          message.success('生成成功，正在前往配置');
          setTimeout(() => {
            showModal.value = false;
            router.push({ name: 'develop_code_deploy', params: { id: res.id } });
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  const dialogWidth = ref('50%');
  function mapWidth() {
    let val = document.body.clientWidth;
    const def = 840; // 默认宽度
    if (val < def) {
      dialogWidth.value = '100%';
    } else {
      dialogWidth.value = def + 'px';
    }
    return dialogWidth.value;
  }

  function getOptionLabel(options, value) {
    if (!options || options?.length === 0) {
      return `unknown`;
    }
    for (const item of options) {
      if (item.value == value) {
        return item.label;
      }
    }
    return `unknown`;
  }

  const loadSelect = async () => {
    selectList.value = await Selects({});
    for (const item of schemas.value) {
      switch (item.field) {
        case 'status':
          item.componentProps.options = selectList.value.status;
          break;
        case 'genType':
          item.componentProps.options = selectList.value.genType;
          break;
      }
    }
  };

  const tablesLoading = ref(false);
  // 处理选项更新
  async function handleDbUpdateValue(value: string | number | Array<string | number> | null) {
    tablesLoading.value = true;
    await loadTableSelect(value);
    tablesLoading.value = false;
  }

  async function loadTableSelect(value) {
    selectList.value.tables = await TableSelect({ name: value });
  }

  function handleTableUpdateValue(value, option) {
    formParams.value.varName = option?.defVarName as string;
    formParams.value.daoName = option?.daoName as string;
    formParams.value.tableComment = option?.defTableComment as string;
  }

  const genTemplateOptions = ref([]);
  function onFocusGenTemplate() {
    for (let i = 0; i < selectList.value.genType?.length; i++) {
      if (selectList.value.genType[i].value === formParams.value.genType) {
        genTemplateOptions.value = selectList.value.genType[i].templates;
        formParams.value.genTemplate = null;
      }
    }
  }

  function onUpdateValueGenType(value) {
    formParams.value.genType = value;
    onFocusGenTemplate();
  }

  const selectAddon = ref(false);
  function handleGenTemplateUpdateValue(value, option) {
    formParams.value.genTemplate = value;
    selectAddon.value = option.isAddon === true;
  }

  onBeforeMount(async () => {
    await loadSelect();
  });
</script>

<style lang="less" scoped></style>
