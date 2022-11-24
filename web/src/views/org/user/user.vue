<template>
  <div>
    <n-card :bordered="false" class="proCard" title="后台用户">
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
        </template>
      </BasicTable>

      <n-modal
        v-model:show="showModal"
        :show-icon="false"
        preset="dialog"
        title="新建"
        :style="{
          width: dialogWidth,
        }"
      >
        <n-form
          :model="formParams"
          :rules="rules"
          ref="formRef"
          label-placement="left"
          :label-width="80"
          class="py-4"
        >
          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="姓名" path="realname">
                <n-input placeholder="请输入姓名" v-model:value="formParams.realname" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="登录用户名" path="username">
                <n-input placeholder="请输入登录用户名" v-model:value="formParams.username" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="绑定角色" path="role">
                <n-select
                  :default-value="formParams.role"
                  :options="roleList"
                  @update:value="handleUpdateRoleValue"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="所属部门" path="dept_id">
                <n-tree-select
                  :options="deptList"
                  :default-value="formParams.dept_id"
                  :default-expand-all="true"
                  @update:value="handleUpdateDeptValue"
                />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="绑定岗位" path="postIds">
                <n-select
                  :default-value="formParams.postIds"
                  multiple
                  :options="postList"
                  @update:value="handleUpdatePostValue"
                />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="密码" path="password">
                <n-input
                  type="password"
                  :placeholder="formParams.id === 0 ? '请输入' : '不填则不修改'"
                  v-model:value="formParams.password"
                />
              </n-form-item>
            </n-gi>
          </n-grid>
          <n-divider title-placement="left">填写更多信息（可选)</n-divider>
          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="手机号" path="mobile">
                <n-input placeholder="请输入" v-model:value="formParams.mobile" />
              </n-form-item>
            </n-gi>
            <n-gi>
              <n-form-item label="邮箱" path="email">
                <n-input placeholder="请输入" v-model:value="formParams.email" />
              </n-form-item>
            </n-gi>
          </n-grid>

          <n-grid x-gap="24" :cols="2">
            <n-gi>
              <n-form-item label="性别" path="sex">
                <n-radio-group v-model:value="formParams.sex" name="sex">
                  <n-radio-button
                    v-for="status in sexOptions"
                    :key="status.value"
                    :value="status.value"
                    :label="status.label"
                  />
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi>
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
            </n-gi>
          </n-grid>

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
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { SelectOption, TreeSelectOption, useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, FormSchema, useForm } from '@/components/Form/index';
  import { Delete, Edit, List, Status } from '@/api/org/user';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { sexOptions, statusActions, statusOptions } from '@/enums/optionsiEnum';
  import { getDeptList } from '@/api/org/dept';
  import { getRoleList } from '@/api/system/role';
  import { getPostList } from '@/api/org/post';

  const params = ref({
    pageSize: 10,
    name: '',
    code: '',
    status: null,
  });

  const rules = {
    name: {
      // required: true,
      trigger: ['blur', 'input'],
      message: '请输入名称',
    },
  };

  const schemas: FormSchema[] = [
    {
      field: 'username',
      component: 'NInput',
      label: '用户名',
      componentProps: {
        placeholder: '请输入用户名',
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
      rules: [{ message: '请输入用户名', trigger: ['blur'] }],
    },
    {
      field: 'realname',
      component: 'NInput',
      label: '姓名',
      componentProps: {
        placeholder: '请输入姓名',
        showButton: false,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'mobile',
      component: 'NInputNumber',
      label: '手机',
      componentProps: {
        placeholder: '请输入手机号码',
        showButton: false,
        onInput: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'status',
      component: 'NSelect',
      label: '状态',
      defaultValue: null,
      componentProps: {
        placeholder: '请选择类型',
        options: statusOptions,
        onUpdateValue: (e: any) => {
          console.log(e);
        },
      },
    },
    {
      field: 'created_at',
      component: 'NDatePicker',
      label: '创建时间',
      componentProps: {
        type: 'datetimerange',
        clearable: true,
        // defaultValue: [new Date() - 86400000 * 30, new Date()],
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
  const deptList = ref([]);
  const roleList = ref([]);
  const postList = ref([]);

  const resetFormParams = {
    id: 0,
    role: null,
    realname: '',
    username: '',
    password: '',
    dept_id: null,
    postIds: null,
    mobile: '',
    email: '',
    sex: 1,
    leader: '',
    phone: '',
    sort: 0,
    status: 1,
    created_at: '',
    updated_at: '',
  };
  let formParams = ref(resetFormParams);

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
    formParams.value = resetFormParams;
  }

  const loadDataTable = async (res) => {
    mapWidth();
    deptList.value = await getDeptList({});
    if (deptList.value === undefined || deptList.value === null) {
      deptList.value = [];
    }

    roleList.value = [];
    let roleLists = await getRoleList();
    if (roleLists.list === undefined || roleLists.list === null) {
      roleLists = [];
    } else {
      roleLists = roleLists.list;
    }
    if (roleLists.length > 0) {
      for (let i = 0; i < roleLists.length; i++) {
        roleList.value[i] = {};
        roleList.value[i].label = roleLists[i].name;
        roleList.value[i].value = roleLists[i].id;
      }
    }

    postList.value = [];
    let postLists = await getPostList();
    if (postLists.list === undefined || postLists.list === null) {
      postLists = [];
    } else {
      postLists = postLists.list;
    }
    if (postLists.length > 0) {
      for (let i = 0; i < postLists.length; i++) {
        postList.value[i] = {};
        postList.value[i].label = postLists[i].name;
        postList.value[i].value = postLists[i].id;
      }
    }
    console.log('post.value:' + JSON.stringify(postList.value));

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
              formParams.value = ref(resetFormParams);
            });
          })
          .catch((e: Error) => {
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
          .catch((_e: Error) => {
            // message.error(_e.message ?? '操作失败');
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
          reloadTable();
        });
      })
      .catch((e: Error) => {
        message.error(e.message ?? '操作失败');
      });
  }

  const dialogWidth = ref('50%');

  function mapWidth() {
    let val = document.body.clientWidth;
    const def = 720; // 默认宽度
    if (val < def) {
      dialogWidth.value = '100%';
    } else {
      dialogWidth.value = def + 'px';
    }

    return dialogWidth.value;
  }

  function handleUpdateDeptValue(
    value: string | number | Array<string | number> | null,
    option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    console.log(value, option);

    formParams.value.dept_id = value;
  }

  function handleUpdateRoleValue(
    value: string | number | Array<string | number> | null,
    option: SelectOption | null | Array<SelectOption | null>
  ) {
    console.log(value, option);

    formParams.value.role = value;
  }

  function handleUpdatePostValue(
    value: string | number | Array<string | number> | null,
    option: SelectOption | null | Array<SelectOption | null>
  ) {
    console.log(value, option);

    formParams.value.postIds = value;
  }
</script>

<style lang="less" scoped></style>
