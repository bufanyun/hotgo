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
        :openChecked="true"
        :columns="columns"
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
            添加任务
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
        :title="formParams?.id > 0 ? '编辑任务 #' + formParams.id : '添加任务'"
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

          <n-form-item label="任务标题" path="title">
            <n-input placeholder="请输入任务标题" v-model:value="formParams.title" />
          </n-form-item>

          <n-form-item label="执行方法" path="name">
            <n-input placeholder="请输入执行方法" v-model:value="formParams.name" />
            <template #feedback>go方法名称</template>
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

          <n-form-item label="执行次数" path="count" v-if="formParams.policy === 4">
            <n-input placeholder="请输入执行次数" v-model:value="formParams.count" />
          </n-form-item>

          <n-form-item label="表达式" path="pattern">
            <n-input placeholder="请输入定时表达式" v-model:value="formParams.pattern" />
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

    <GroupModal ref="GroupModalRef" @reloadGroupOption="reloadGroupOption" />

    <n-modal
      v-model:show="showStdout"
      style="width: 72%"
      :show-icon="false"
      preset="dialog"
      title="调度日志"
    >
      <n-form ref="formRef" label-placement="left" :label-width="80" class="py-4">
        <n-form-item label="日志文件" path="fileName">
          <n-input placeholder="" v-model:value="log.fileName" readonly />
        </n-form-item>

        <n-form-item label="文件大小" path="sizeFormat">
          <n-input placeholder="" v-model:value="log.sizeFormat" readonly />
        </n-form-item>

        <n-form-item label="日志内容" path="log" class="n-code-container">
          <n-code
            class="n-code-contents"
            id="n-code"
            :word-wrap="true"
            :code="log?.contents ?? ''"
            language="html"
          />
        </n-form-item>
      </n-form>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref, onBeforeMount, computed } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Delete, Edit, getSelect, List, Status, OnlineExec, DispatchLog } from '@/api/sys/cron';
  import { columns } from './columns';
  import { DeleteOutlined, GroupOutlined, PlusOutlined } from '@vicons/antd';
  import GroupModal from './modal/modal.vue';
  import { adaTableScrollX } from '@/utils/hotgo';

  const optionTreeData = ref<any>([]);
  const defaultValueRef = () => ({
    id: 0,
    groupId: 0,
    title: '',
    name: '',
    params: '',
    pattern: '',
    policy: 1,
    count: 1,
    sort: 0,
    remark: '',
    status: 1,
  });
  const params = ref<any>({
    pageSize: 10,
    title: '',
    content: '',
    status: null,
  });

  const rules = {};

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

  const statusActions = [
    {
      label: '设为运行中',
      key: 1,
    },
    {
      label: '设为已结束',
      key: 2,
    },
    {
      label: '删除任务',
      key: 3,
    },
  ];

  const groupOptions = ref<any>([]);
  const showStdout = ref(false);
  const log = ref();
  const schemas: FormSchema[] = [
    {
      field: 'groupId',
      component: 'NTreeSelect',
      label: '任务分组',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择分组',
        options: groupOptions,
        clearable: true,
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
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  let formParams = ref<any>(defaultValueRef());

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
            type: 'success',
          },
          {
            label: '调度日志',
            onClick: handleDispatchLog.bind(null, record),
            type: 'default',
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
          },
        ],
        dropDownActions: statusActions,
        select: (key) => {
          if (key < 3) {
            updateStatus(record.id, key);
            return;
          }
          if (key == 3) {
            handleDelete(record);
            return;
          }
        },
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

  function addTable() {
    showModal.value = true;
    formParams.value = defaultValueRef();
  }

  const loadDataTable = async (res) => {
    return await List({ ...params.value, ...res, ...searchFormRef.value?.formModel });
  };

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
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
        Edit(formParams.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showModal.value = false;
            reloadTable();
            formParams.value = ref(defaultValueRef());
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = record;
  }

  function handleDispatchLog(record: Recordable) {
    DispatchLog(record).then((res) => {
      showStdout.value = true;
      log.value = res;

      setTimeout(() => {
        const container = document.getElementById('n-code');
        container.scrollTop = container.scrollHeight;
      }, 100);
    });
  }

  function handleExecute(record: Recordable) {
    dialog.info({
      title: '提示',
      content: '提交成功后将立即执行一次，你确定要执行吗？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        OnlineExec(record).then((_res) => {
          message.success('提交成功，执行结果请登录控制台查看日志！');
          reloadTable();
        });
      },
    });
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

  function handleSubmit(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function handleReset(values: Recordable) {
    params.value = values;
    reloadTable();
  }

  function updateStatus(id, status) {
    Status({ id: id, status: status }).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  const GroupModalRef = ref();
  function openGroupModal() {
    const { openDrawer } = GroupModalRef.value;
    openDrawer();
  }

  async function reloadGroupOption() {
    const tmp = await getSelect({});
    optionTreeData.value = tmp.list;
    if (optionTreeData.value === undefined || optionTreeData.value === null) {
      optionTreeData.value = [];
    }
    groupOptions.value = transformGroupData(optionTreeData.value);
  }

  function transformGroupData(data: any[]) {
    return data.map((item) => {
      return {
        key: item.key,
        value: item.key,
        label: item.label,
        children: item.children !== null ? transformGroupData(item.children) : null,
      };
    });
  }

  onBeforeMount(async () => {
    await reloadGroupOption();
  });

  // 处理选项更新
  function handleUpdateValue(value: string | number) {
    formParams.value.groupId = value;
  }
</script>

<style lang="less" scoped>
  .n-code-container {
    height: 450px;
    width: 100%;
  }

  .n-code-contents {
    height: 450px;
    width: 100%;
    overflow: auto;
    background: #2f3129;
    color: wheat;
    padding: 10px;
  }
</style>
