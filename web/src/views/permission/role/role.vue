<template>
  <div>
    <n-card :bordered="false" title="角色管理">
      <n-space vertical :size="12">
        <n-space>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加角色
          </n-button>
        </n-space>

        <n-data-table
          v-if="data.length > 0 || !loading"
          :columns="columns.concat(actionColumn)"
          :data="data"
          :row-key="(row) => row.id"
          :loading="loading"
          :resizeHeightOffset="-20000"
          default-expand-all
        />
      </n-space>
    </n-card>

    <n-modal
      v-model:show="showModal"
      :show-icon="false"
      :mask-closable="false"
      preset="dialog"
      :title="editRoleTitle"
    >
      <div class="py-3 menu-list" :style="{ maxHeight: '90vh', height: '70vh' }">
        <n-tree
          block-line
          checkable
          :check-on-click="true"
          :default-expand-all="true"
          :virtual-scroll="true"
          :data="treeData"
          :expandedKeys="expandedKeys"
          :checked-keys="checkedKeys"
          style="max-height: 950px; overflow: hidden"
          @update:checked-keys="checkedTree"
          @update:expanded-keys="onExpandedKeys"
        />
      </div>
      <template #action>
        <n-space>
          <n-button type="info" ghost icon-placement="left" @click="packHandle">
            全部{{ expandedKeys.length ? '收起' : '展开' }}
          </n-button>

          <n-button type="info" ghost icon-placement="left" @click="checkedAllHandle">
            全部{{ checkedAll ? '取消' : '选择' }}
          </n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm">提交</n-button>
        </n-space>
      </template>
    </n-modal>

    <n-modal
      v-model:show="showModal2"
      :show-icon="false"
      preset="dialog"
      :title="formParams.id > 0 ? '编辑角色 #' + formParams.id : '添加角色'"
    >
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-form-item label="上级角色" path="pid">
          <n-tree-select
            :options="optionTreeData"
            :default-value="formParams.pid"
            key-field="id"
            label-field="name"
            :on-update:value="onUpdateValuePid"
          />
        </n-form-item>
        <n-form-item label="角色名称" path="name">
          <n-input placeholder="请输入名称" v-model:value="formParams.name" />
        </n-form-item>
        <n-form-item label="权限编码" path="key">
          <n-input placeholder="请输入" v-model:value="formParams.key" />
        </n-form-item>
        <!--        <n-form-item label="排序" path="sort">-->
        <!--          <n-input-number v-model:value="formParams.sort" clearable />-->
        <!--        </n-form-item>-->

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
          <n-input type="textarea" placeholder="请输入" v-model:value="formParams.remark" />
        </n-form-item>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showModal2 = false)">取消</n-button>
          <n-button type="info" :loading="formBtnLoading2" @click="confirmForm2">确定</n-button>
        </n-space>
      </template>
    </n-modal>

    <n-modal
      v-model:show="showDataModal"
      :show-icon="false"
      preset="dialog"
      :title="'修改 ' + dataForm?.name + ' 的数据权限'"
    >
      <n-form
        :model="dataForm"
        ref="dataFormRef"
        label-placement="left"
        :label-width="120"
        class="py-4"
      >
        <n-form-item label="数据范围" path="dataScope">
          <n-select v-model:value="dataForm.dataScope" :options="dataScopeOption" />
        </n-form-item>
        <n-form-item label="自定义权限" path="customDept" v-if="dataForm.dataScope === 4">
          <n-tree-select
            multiple
            key-field="id"
            :options="deptList"
            v-model:value="dataForm.customDept"
            :default-expand-all="true"
            @update:value="handleUpdateDeptValue"
          />
        </n-form-item>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="() => (showDataModal = false)">取消</n-button>
          <n-button type="info" :loading="dataFormBtnLoading" @click="confirmDataForm"
            >确定</n-button
          >
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref } from 'vue';
  import { NButton, useDialog, useMessage } from 'naive-ui';
  import { BasicColumn, TableAction } from '@/components/Table';
  import {
    Delete,
    Edit,
    GetPermissions,
    getRoleList,
    UpdatePermissions,
    DataScopeSelect,
    DataScopeEdit,
  } from '@/api/system/role';
  import { EditMenu, getMenuList } from '@/api/system/menu';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { getAllExpandKeys, getTreeAll } from '@/utils';
  import { statusOptions } from '@/enums/optionsiEnum';
  import { cloneDeep } from 'lodash-es';
  import { getDeptList } from '@/api/org/dept';

  const formRef: any = ref(null);
  const message = useMessage();
  const dialog = useDialog();
  const showModal2 = ref(false);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formBtnLoading2 = ref(false);
  const checkedAll = ref<any>(false);
  const editRoleTitle = ref('');
  const treeData = ref([]);
  const expandedKeys = ref([]);
  const checkedKeys = ref<any>([]);
  const updatePermissionsParams = ref<any>({});
  const optionTreeData = ref<any>([]);
  const dataScopeOption = ref<any>();
  const deptList = ref<any>([]);
  const dataFormRef = ref<any>();
  const dataFormBtnLoading = ref(false);
  const showDataModal = ref(false);
  const dataForm = ref<any>();
  const loading = ref(false);
  const data = ref<any>([]);

  const rules = {
    name: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入名称',
    },
    key: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入角色编码',
    },
  };

  const defaultState = {
    id: 0,
    pid: 0,
    level: 1,
    tree: '',
    name: '',
    key: '',
    remark: null,
    status: 1,
    sort: 0,
    dataScope: 1,
    customDept: [],
  };

  const formParams = ref<any>(cloneDeep(defaultState));

  const actionColumn = reactive<BasicColumn>({
    width: 300,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction, {
        style: 'primary',
        actions: [
          {
            label: '菜单权限',
            onClick: handleMenuAuth.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            type: 'default',
          },
          {
            label: '数据权限',
            onClick: handleDataAuth.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
            type: 'default',
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
            ifShow: () => {
              return record.id !== 1;
            },
          },
        ],
      });
    },
  });

  const loadDataTable = async (res: any) => {
    loading.value = true;
    const tmp = await getRoleList({ ...res, ...{ pageSize: 100, page: 1 } });
    data.value = tmp.list ?? [];
    loading.value = false;
  };

  function reloadTable() {
    loadDataList();
    loadDataTable({});
  }

  function confirmForm(e: any) {
    e.preventDefault();
    formBtnLoading.value = true;
    UpdatePermissions({
      ...{
        id: updatePermissionsParams.value.id,
        menuIds:
          checkedKeys.value === undefined || checkedKeys.value == null ? [] : checkedKeys.value,
      },
    }).then((_res) => {
      message.success('操作成功');
      reloadTable();
      showModal.value = false;
      formBtnLoading.value = false;
    });
  }

  function confirmForm2(e) {
    e.preventDefault();
    formBtnLoading2.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formParams.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showModal2.value = false;
            reloadTable();
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading2.value = false;
    });
  }

  function addTable() {
    showModal2.value = true;
    formParams.value = cloneDeep(defaultState);
  }

  function handleEdit(record: Recordable) {
    showModal2.value = true;
    formParams.value = cloneDeep(record);
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

  async function handleMenuAuth(record: Recordable) {
    editRoleTitle.value = `分配 ${record.name} 的菜单权限`;
    checkedKeys.value = [];
    checkedAll.value = false;
    const data = await GetPermissions({ ...{ id: record.id } });
    checkedKeys.value = data.menuIds;
    updatePermissionsParams.value.id = record.id;
    showModal.value = true;
  }

  function handleDataAuth(record: Recordable) {
    dataForm.value = cloneDeep(record);
    showDataModal.value = true;
  }

  function handleUpdateDeptValue(value: string | number | Array<string | number> | null) {
    dataForm.value.customDept = value;
  }

  function confirmDataForm(e) {
    e.preventDefault();
    dataFormBtnLoading.value = true;
    dataFormRef.value.validate((errors) => {
      if (!errors) {
        DataScopeEdit(dataForm.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            showDataModal.value = false;
            reloadTable();
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      dataFormBtnLoading.value = false;
    });
  }

  function checkedTree(keys) {
    checkedKeys.value = keys;
  }

  function onExpandedKeys(keys) {
    expandedKeys.value = keys;
  }

  function packHandle() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = getAllExpandKeys(treeData) as [];
    }
  }

  function checkedAllHandle() {
    if (!checkedAll.value) {
      checkedKeys.value = getTreeAll(treeData.value);
      checkedAll.value = true;
    } else {
      checkedKeys.value = [];
      checkedAll.value = false;
    }
  }

  onMounted(async () => {
    loadDataList();
    loadMenuList();
    loadDeptList();
    loadDataScopeSelect();
    await loadDataTable({});
  });

  function loadDataList() {
    getRoleList({ pageSize: 100, page: 1 }).then((res) => {
      optionTreeData.value = [
        {
          id: 0,
          key: 0,
          label: '顶级角色',
          pid: 0,
          name: '顶级角色',
        },
      ];
      optionTreeData.value = optionTreeData.value.concat(res.list);
    });
  }

  function loadMenuList() {
    getMenuList().then((res) => {
      expandedKeys.value = getAllExpandKeys(res.list) as [];
      treeData.value = res.list;
    });
  }

  function loadDeptList() {
    getDeptList({}).then((res) => {
      if (res.list) {
        deptList.value = res.list;
      }
    });
  }

  function loadDataScopeSelect() {
    DataScopeSelect().then((res) => {
      if (res.list) {
        dataScopeOption.value = res.list;
      }
    });
  }

  function onUpdateValuePid(value: string | number) {
    formParams.value.pid = value;
  }
</script>

<style lang="less" scoped></style>
