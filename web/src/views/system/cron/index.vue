<template>
  <div>
    <n-card :bordered="false" class="proCard" title="定时任务">
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
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="1090"
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
          &nbsp;
          <n-button type="error" @click="batchDelete" :disabled="batchDeleteDisabled">
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
          &nbsp;
          <n-button type="info" @click="openGroupModal">
            <template #icon>
              <n-icon>
                <GroupOutlined />
              </n-icon>
            </template>
            任务分组
          </n-button>
        </template>
      </BasicTable>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        preset="dialog"
        title="新建"
        style="width: 720px"
      >
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="任务分组" path="groupId">
            <n-tree-select
              :options="optionTreeData"
              :default-value="formParams.groupId"
              @update:value="handleUpdateValue"
            />
          </n-form-item>

          <n-form-item label="任务名称" path="name">
            <n-input placeholder="请输入公告标题" v-model:value="formParams.name" />
            <template #feedback> go函数名称</template>
          </n-form-item>

          <n-form-item label="执行参数" path="params">
            <n-input
              type="textarea"
              placeholder="请输入执行参数，如果函数需要多个参数请用,隔开"
              v-model:value="formParams.params"
            />
          </n-form-item>

          <n-form-item label="执行策略" path="policy">
            <n-radio-group v-model:value="formParams.policy" name="policy">
              <n-radio-button
                v-for="type in policyOptions"
                :key="type.value"
                :value="type.value"
                :label="type.label"
              />
            </n-radio-group>
          </n-form-item>

          <n-form-item label="执行次数" path="count">
            <n-input placeholder="请输入执行次数" v-model:value="formParams.count" />
            <template #feedback> 仅在单次、多次策略时生效</template>
          </n-form-item>

          <n-form-item label="定时表达式" path="pattern">
            <n-input placeholder="请输入表达式" v-model:value="formParams.pattern" />
            <template #feedback>
              表达式语法参考：<a
                target="_blank"
                href="https://goframe.org/pages/viewpage.action?pageId=30736411"
                >https://goframe.org/pages/viewpage.action?pageId=30736411</a
              >
            </template>
          </n-form-item>

          <n-form-item label="排序" path="sort">
            <n-input-number v-model:value="formParams.sort" clearable />
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

    <GroupModal ref="GroupModalRef" />
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref, onBeforeMount } from 'vue';
  import { TreeSelectOption, useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Delete, Edit, getSelect, List, Status } from '@/api/sys/cron';
  import { columns } from './columns';
  import { DeleteOutlined, GroupOutlined, PlusOutlined } from '@vicons/antd';
  import { statusActions } from '@/enums/optionsiEnum';
  import GroupModal from './modal/modal.vue';

  const optionTreeData = ref([]);
  const defaultValueRef = () => ({
    id: 0,
    groupId: 0,
    name: '',
    params: '',
    pattern: '',
    policy: 1,
    count: 1,
    sort: 0,
    remark: '',
    status: 1,
  });
  const params = ref({
    pageSize: 10,
    title: '',
    content: '',
    status: null,
  });

  const rules = {
    name: {
      // required: true,
      trigger: ['blur', 'input'],
      message: '请输入任务名称',
    },
  };

  const policyOptions = [
    {
      value: 1,
      label: '并行策略',
    },
    {
      value: 2,
      label: '单例策略',
    },
    {
      value: 3,
      label: '单次策略',
    },
    {
      value: 4,
      label: '多次策略',
    },
  ].map((s) => {
    return s;
  });

  const statusOptions = [
    {
      value: 1,
      label: '运行中',
    },
    {
      value: 2,
      label: '已结束',
    },
  ].map((s) => {
    return s;
  });

  const groupOptions = ref([]);

  const schemas: FormSchema[] = [
    {
      field: 'groupId',
      component: 'NSelect',
      label: '任务分组',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择分组',
        options: groupOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'name',
      component: 'NInput',
      label: '任务名称',
      componentProps: {
        placeholder: '请输入任务名称',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入任务名称', trigger: ['blur'] }],
    },
    {
      field: 'policy',
      component: 'NSelect',
      label: '执行策略',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择策略',
        options: policyOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '运行状态',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择类型',
        options: statusOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref({});
  const formRef = ref({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  let formParams = ref(defaultValueRef());

  const actionColumn = reactive({
    width: 320,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '在线执行',
            onClick: handleExecute.bind(null, record),
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
        dropDownActions: statusActions,
        select: (key) => {
          updateStatus(record.id, key);
        },
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
    formParams.value = defaultValueRef();
  }

  const loadDataTable = async (res) => {
    return await List({ ...params.value, ...res, ...searchFormRef.value.formModel });
  };

  function onCheckedRow(rowKeys) {
    console.log(rowKeys);
    if (rowKeys.length > 0) {
      batchDeleteDisabled.value = false;
    } else {
      batchDeleteDisabled.value = true;
    }

    checkedIds.value = rowKeys;
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
        Edit(formParams.value)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            setTimeout(() => {
              showModal.value = false;
              reloadTable();
              formParams.value = ref(defaultValueRef());
            });
          })
          .catch((_e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    showModal.value = true;
    formParams.value = record;
  }

  function handleExecute(record: Recordable) {
    console.log('点击了handleExecute', record);
  }

  function handleDelete(record: Recordable) {
    console.log('点击了删除', record);
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete(record)
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

  function batchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '不确定',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value })
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            reloadTable();
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  function handleSubmit(values: Recordable) {
    console.log(values);
    params.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function updateStatus(id, status) {
    Status({ id: id, status: status })
      .then((_res) => {
        console.log('_res:' + JSON.stringify(_res));
        message.success('操作成功');
        setTimeout(() => {
          reloadTable({});
        });
      })
      .catch((e: Error) => {
        message.error(e.message ?? '操作失败');
      });
  }

  const GroupModalRef = ref();

  function openGroupModal() {
    const { openDrawer } = GroupModalRef.value;
    openDrawer();
  }

  async function setDictSelect() {
    optionTreeData.value = await getSelect({});
    if (optionTreeData.value === undefined || optionTreeData.value === null) {
      optionTreeData.value = [];
    }

    groupOptions.value = [];
    for (let i = 0; i < optionTreeData.value?.length; i++) {
      groupOptions.value.push({
        value: optionTreeData.value[i].key,
        label: optionTreeData.value[i].label,
      });
    }
  }

  onBeforeMount(async () => {
    await setDictSelect();
  });

  // 处理选项更新
  function handleUpdateValue(
    value: string | number | Array<string | number> | null,
    option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    console.log(value, option);
    formParams.value.groupId = value;
  }
</script>

<style lang="less" scoped></style>
