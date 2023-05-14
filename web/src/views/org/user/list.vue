<template>
  <div>
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
      :scroll-x="1800"
    >
      <template #tableTitle>
        <n-button
          type="primary"
          @click="addTable"
          class="min-left-space"
          v-if="hasPermission(['/member/edit'])"
        >
          <template #icon>
            <n-icon>
              <PlusOutlined />
            </n-icon>
          </template>
          添加用户
        </n-button>
        <n-button
          type="error"
          @click="batchDelete"
          :disabled="batchDeleteDisabled"
          class="min-left-space"
          v-if="hasPermission(['/member/delete'])"
        >
          <template #icon>
            <n-icon>
              <DeleteOutlined />
            </n-icon>
          </template>
          批量删除
        </n-button>

        <n-button
          type="success"
          @click="handleInviteQR(userStore.info?.inviteCode)"
          class="min-left-space"
          v-if="userStore.loginConfig?.loginRegisterSwitch === 1"
        >
          <template #icon>
            <n-icon>
              <QrCodeOutline />
            </n-icon>
          </template>
          邀请注册
        </n-button>
      </template>
    </BasicTable>

    <n-modal
      v-model:show="showModal"
      :show-icon="false"
      preset="dialog"
      :title="formParams?.id > 0 ? '编辑用户 #' + formParams?.id : '添加用户'"
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
            <n-form-item label="姓名" path="realName">
              <n-input placeholder="请输入姓名" v-model:value="formParams.realName" />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="用户名" path="username">
              <n-input placeholder="请输入登录用户名" v-model:value="formParams.username" />
            </n-form-item>
          </n-gi>
        </n-grid>

        <n-grid x-gap="24" :cols="2">
          <n-gi>
            <n-form-item label="绑定角色" path="roleId">
              <n-tree-select
                key-field="id"
                :options="options.role"
                :default-value="formParams.roleId"
                :default-expand-all="true"
                @update:value="handleUpdateRoleValue"
              />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="所属部门" path="deptId">
              <n-tree-select
                key-field="id"
                :options="options.dept"
                :default-value="formParams.deptId"
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
                :options="options.post"
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

    <AddBalance
      @reloadTable="reloadTable"
      @updateShowModal="updateBalanceShowModal"
      :showModal="showBalanceModal"
      :formParams="formParams"
    />

    <AddIntegral
      @reloadTable="reloadTable"
      @updateShowModal="updateIntegralShowModal"
      :showModal="showIntegralModal"
      :formParams="formParams"
    />

    <n-modal v-model:show="showQrModal" :show-icon="false" preset="dialog" title="邀请注册二维码">
      <n-form class="py-4">
        <div class="text-center">
          <qrcode-vue :value="qrParams.qrUrl" :size="220" class="canvas" style="margin: 0 auto" />
        </div>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showQrModal = false)">关闭</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { h, reactive, ref } from 'vue';
  import { SelectOption, TreeSelectOption, useDialog, useMessage } from 'naive-ui';
  import { ActionItem, BasicTable, TableAction } from '@/components/Table';
  import { BasicForm } from '@/components/Form/index';
  import { Delete, Edit, List, Status, ResetPwd } from '@/api/org/user';
  import { columns } from './columns';
  import { PlusOutlined, DeleteOutlined } from '@vicons/antd';
  import { QrCodeOutline } from '@vicons/ionicons5';
  import { sexOptions, statusOptions } from '@/enums/optionsiEnum';
  import { adaModalWidth } from '@/utils/hotgo';
  import { getRandomString } from '@/utils/charset';
  import { cloneDeep } from 'lodash-es';
  import QrcodeVue from 'qrcode.vue';
  import AddBalance from './addBalance.vue';
  import AddIntegral from './addIntegral.vue';
  import { addNewState, addState, options, register, defaultState } from './model';
  import { usePermission } from '@/hooks/web/usePermission';
  import { useRouter } from 'vue-router';
  import { useUserStore } from '@/store/modules/user';
  import { LoginRoute } from '@/router';

  interface Props {
    type?: string;
  }

  const props = withDefaults(defineProps<Props>(), {
    type: '-1',
  });

  const rules = {
    username: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入用户名',
    },
  };

  const { hasPermission } = usePermission();
  const router = useRouter();
  const userStore = useUserStore();
  const showIntegralModal = ref(false);
  const showBalanceModal = ref(false);
  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const searchFormRef = ref<any>({});
  const formRef = ref<any>({});
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const dialogWidth = ref('50%');
  const formParams = ref<any>();
  const showQrModal = ref(false);
  const qrParams = ref({
    name: '',
    qrUrl: '',
  });

  const actionColumn = reactive({
    width: 220,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      const downActions = getDropDownActions(record);
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '已启用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1 && record.id !== 1;
            },
            auth: ['/member/status'],
          },
          {
            label: '已禁用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2 && record.id !== 1;
            },
            auth: ['/member/status'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            auth: ['/member/edit'],
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            auth: ['/member/delete'],
          },
        ],
        dropDownActions: downActions,
        select: (key) => {
          if (key === 0) {
            return handleResetPwd(record);
          }
          if (key === 100) {
            return handleAddBalance(record);
          }
          if (key === 101) {
            return handleAddIntegral(record);
          }
          if (key === 102) {
            if (userStore.loginConfig?.loginRegisterSwitch !== 1) {
              message.error('管理员暂未开启此功能');
              return;
            }
            return handleInviteQR(record.inviteCode);
          }
        },
      });
    },
  });

  function getDropDownActions(record: Recordable): ActionItem[] {
    if (record.id === 1) {
      return [];
    }

    let list = [
      {
        label: '重置密码',
        key: 0,
      },
      {
        label: '变更余额',
        key: 100,
      },
      {
        label: '变更积分',
        key: 101,
      },
    ];

    if (userStore.loginConfig?.loginRegisterSwitch === 1) {
      list.push({
        label: 'TA的邀请码',
        key: 102,
      });
    }

    return list;
  }

  function addTable() {
    showModal.value = true;
    formParams.value = cloneDeep(defaultState);
  }

  const loadDataTable = async (res) => {
    adaModalWidth(dialogWidth);
    return await List({ ...res, ...searchFormRef.value?.formModel, ...{ roleId: props.type } });
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
    formParams.value = cloneDeep(record);
  }

  function handleResetPwd(record: Recordable) {
    record.password = getRandomString(12);
    dialog.warning({
      title: '警告',
      content: '你确定要重置密码？\r\n重置成功后密码为：' + record.password + '\r\n 请先保存',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        ResetPwd(record).then((_res) => {
          message.success('操作成功');
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

  function handleSubmit(_values: Recordable) {
    reloadTable();
  }

  function handleReset(_values: Recordable) {
    reloadTable();
  }

  function handleStatus(record: Recordable, status) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  function handleUpdateDeptValue(
    value: string | number | Array<string | number> | null,
    _option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    formParams.value.deptId = Number(value);
  }

  function handleUpdateRoleValue(
    value: string | number | Array<string | number> | null,
    _option: SelectOption | null | Array<SelectOption | null>
  ) {
    formParams.value.roleId = Number(value);
  }

  function handleUpdatePostValue(
    value: string | number | Array<string | number> | null,
    _option: SelectOption | null | Array<SelectOption | null>
  ) {
    formParams.value.postIds = value;
  }

  function updateBalanceShowModal(value) {
    showBalanceModal.value = value;
  }

  function handleAddBalance(record: Recordable) {
    showBalanceModal.value = true;
    formParams.value = addNewState(record as addState);
  }

  function updateIntegralShowModal(value) {
    showIntegralModal.value = value;
  }

  function handleAddIntegral(record: Recordable) {
    showIntegralModal.value = true;
    formParams.value = addNewState(record as addState);
  }

  function handleInviteQR(code: string) {
    const w = window.location;
    const domain = w.protocol + '//' + w.host + w.pathname + '#';
    qrParams.value.qrUrl = domain + LoginRoute.path + '?scope=register&inviteCode=' + code;
    showQrModal.value = true;
  }
</script>

<style lang="less" scoped></style>
