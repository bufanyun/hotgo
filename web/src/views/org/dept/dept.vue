<template>
  <div>
    <n-card :bordered="false" title="部门管理">
      <BasicForm
        @register="register"
        @submit="handleSubmit"
        @reset="handleReset"
        @keyup.enter="handleSubmit"
        ref="formRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>
      <n-space vertical :size="12">
        <n-space>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加部门
          </n-button>
        </n-space>

        <n-data-table
          :columns="columns"
          :data="data"
          :row-key="rowKey"
          :loading="loading"
          :resizeHeightOffset="-20000"
          default-expand-all
        />
      </n-space>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        preset="dialog"
        :title="formParams?.id > 0 ? '编辑部门 #' + formParams?.id : '添加部门'"
      >
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-form-item label="上级部门" path="pid">
            <n-tree-select
              key-field="id"
              :options="options"
              :default-value="optionsDefaultValue"
              :default-expand-all="true"
              @update:value="handleUpdateValue"
            />
          </n-form-item>

          <n-form-item label="部门名称" path="name">
            <n-input placeholder="请输入名称" v-model:value="formParams.name" />
          </n-form-item>
          <n-form-item label="部门编码" path="code">
            <n-input placeholder="请输入部门编码" v-model:value="formParams.code" />
          </n-form-item>

          <n-form-item label="负责人" path="leader">
            <n-input placeholder="请输入负责人" v-model:value="formParams.leader" />
          </n-form-item>
          <n-form-item label="联系电话" path="phone">
            <n-input placeholder="请输入联系电话" v-model:value="formParams.phone" />
          </n-form-item>
          <n-form-item label="邮箱" path="email">
            <n-input placeholder="请输入邮箱" v-model:value="formParams.email" />
          </n-form-item>

          <!--          <n-form-item label="排序" path="sort">-->
          <!--            <n-input-number v-model:value="formParams.sort" clearable />-->
          <!--          </n-form-item>-->

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
        </n-form>

        <template #action>
          <n-space>
            <n-button @click="() => (showModal = false)">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-card>
  </div>
</template>

<script lang="ts" setup name="org_dept">
  import { h, onMounted, ref } from 'vue';
  import { DataTableColumns, NButton, NTag, useDialog, useMessage } from 'naive-ui';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { PlusOutlined } from '@vicons/antd';
  import { TableAction } from '@/components/Table';
  import { statusActions, statusOptions } from '@/enums/optionsiEnum';
  import { Delete, Edit, getDeptList, Status } from '@/api/org/dept';
  import { cloneDeep } from 'lodash-es';
  import { renderIcon, renderTooltip } from '@/utils';
  import { HelpCircleOutline } from '@vicons/ionicons5';

  const rules = {
    name: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入名称',
    },
    code: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入编码',
    },
  };

  const schemas: FormSchema[] = [
    {
      field: 'name',
      component: 'NInput',
      label: '部门名称',
      componentProps: {
        placeholder: '请输入部门名称',
        onInput: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入部门名称', trigger: ['blur'] }],
    },
    {
      field: 'code',
      component: 'NInput',
      label: '部门编码',
      componentProps: {
        placeholder: '请输入部门编码',
        showButton: false,
        onInput: (e: any) => {
          console.log(e);
        },
      },
    },
  ];

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  const options = ref<any>([]);
  const optionsDefaultValue = ref<any>(null);
  const loading = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);

  const defaultState = {
    id: 0,
    pid: 0,
    name: '',
    code: '',
    type: '',
    leader: '',
    phone: '',
    email: '',
    sort: 0,
    status: 1,
    createdAt: '',
    updatedAt: '',
  };
  let formParams = ref<any>();

  type RowData = {
    createdAt: string;
    status: number;
    name: string;
    id: number;
    children?: RowData[];
  };
  const data = ref<any>([]);
  const columns: DataTableColumns<RowData> = [
    {
      title(_column) {
        return renderTooltip(
          h(
            NButton,
            {
              strong: true,
              size: 'small',
              text: true,
              iconPlacement: 'right',
            },
            { default: () => '部门', icon: renderIcon(HelpCircleOutline) }
          ),
          '支持上下级部门，点击列表中左侧 > 按钮可展开下级部门列表'
        );
      },
      key: 'name',
      render(row) {
        return h(
          NTag,
          {
            type: 'info',
          },
          {
            default: () => row.name,
          }
        );
      },
      width: 200,
    },
    // {
    //   title: '部门ID',
    //   key: 'index',
    //   width: 100,
    // },
    {
      title: '部门编码',
      key: 'code',
      width: 100,
    },
    {
      title: '负责人',
      key: 'leader',
      width: 100,
    },
    {
      title: '联系电话',
      key: 'phone',
      width: 150,
    },
    {
      title: '邮箱',
      key: 'email',
      width: 150,
    },
    {
      title: '状态',
      key: 'status',
      width: 80,
      render(row) {
        return h(
          NTag,
          {
            style: {
              marginRight: '6px',
            },
            type: row.status == 1 ? 'info' : 'error',
            bordered: false,
          },
          {
            default: () => (row.status == 1 ? '正常' : '已禁用'),
          }
        );
      },
    },
    {
      title: '创建时间',
      key: 'createdAt',
      width: 150,
      render: (rows, _) => {
        return rows.createdAt;
      },
    },
    {
      title: '操作',
      key: 'actions',
      width: 220,
      fixed: 'right',
      render(record: any) {
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
          dropDownActions: statusActions,
          select: (key) => {
            updateStatus(record.id, key);
          },
        });
      },
    },
  ];

  const rowKey = (row: RowData) => row.id;

  function addTable() {
    showModal.value = true;
    formParams.value = cloneDeep(defaultState);
    optionsDefaultValue.value = null;
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = cloneDeep(record);
    formParams.value.children = null;
    optionsDefaultValue.value = formParams.value.pid;
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record)
          .then((_res) => {
            message.success('操作成功');
            loadDataTable({});
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('取消');
      },
    });
  }

  function updateStatus(id, status) {
    Status({ id: id, status: status })
      .then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          loadDataTable({});
        });
      })
      .catch((e: Error) => {
        message.error(e.message ?? '操作失败');
      });
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
            loadDataTable({});
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  async function handleSubmit(values: Recordable) {
    await loadDataTable(values);
  }

  function handleReset(_values: Recordable) {}

  const loadDataTable = async (res) => {
    loading.value = true;
    const tmp = await getDeptList({ ...res, ...formRef.value?.formModel });
    data.value = tmp?.list;
    if (data.value === undefined || data.value === null) {
      data.value = [];
    }
    options.value = [
      {
        index: 0,
        key: 0,
        label: '顶级部门',
        children: data.value,
      },
    ];

    loading.value = false;
  };

  onMounted(async () => {
    await loadDataTable({});
  });

  function handleUpdateValue(value) {
    formParams.value.pid = value;
  }
</script>
