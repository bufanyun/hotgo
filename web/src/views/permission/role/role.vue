<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="角色管理"> 在这里可以管理你权限下的角色权限</n-card>
    </div>
    <n-card :bordered="false" class="mt-4 proCard">
      <BasicTable
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        @update:checked-row-keys="onCheckedRow"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加角色
          </n-button>
        </template>

        <template #action>
          <TableAction />
        </template>
      </BasicTable>
    </n-card>

    <n-modal v-model:show="showModal" :show-icon="false" preset="dialog" :title="editRoleTitle">
      <div class="py-3 menu-list">
        <n-tree
          block-line
          cascade
          checkable
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

    <n-modal v-model:show="showModal2" :show-icon="false" preset="dialog" title="添加角色">
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-form-item label="角色名称" path="name">
          <n-input placeholder="请输入名称" v-model:value="formParams.name" />
        </n-form-item>
        <n-form-item label="权限编码" path="key">
          <n-input placeholder="请输入" v-model:value="formParams.key" />
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
  </div>
</template>

<script lang="ts" setup>
  import { h, onMounted, reactive, ref, unref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { Delete, Edit, GetPermissions, getRoleList, UpdatePermissions } from '@/api/system/role';
  import { getMenuList } from '@/api/system/menu';
  import { columns } from './columns';
  import { PlusOutlined } from '@vicons/antd';
  import { getTreeAll } from '@/utils';
  import { useRouter } from 'vue-router';
  import { statusOptions } from '@/enums/optionsiEnum';
  import { copyObj } from '@/utils/array';

  const router = useRouter();
  const formRef: any = ref(null);
  const message = useMessage();
  const actionRef = ref();
  const dialog = useDialog();
  const showModal2 = ref(false);
  const showModal = ref(false);
  const formBtnLoading = ref(false);
  const formBtnLoading2 = ref(false);
  const checkedAll = ref(false);
  const editRoleTitle = ref('');
  const treeData = ref([]);
  const expandedKeys = ref([]);
  const checkedKeys = ref([]);

  const updatePermissionsParams = ref({});

  const rules = {
    name: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入名称',
    },
    address: {
      required: true,
      trigger: ['blur', 'input'],
      message: '请输入地址',
    },
    date: {
      type: 'number',
      required: true,
      trigger: ['blur', 'change'],
      message: '请选择日期',
    },
  };
  let formParams = reactive({
    id: 0,
    name: '',
    key: '',
    remark: null,
    status: 1,
    sort: 0,
    dataScope: 0,
    deptCheckStrictly: 0,
    menuCheckStrictly: 0,
  });

  const params = reactive({
    pageSize: 5,
    name: 'xiaoMa',
  });

  const actionColumn = reactive({
    width: 250,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction, {
        style: 'button',
        actions: [
          {
            label: '菜单权限',
            onClick: handleMenuAuth.bind(null, record),
            // 根据业务控制是否显示 isShow 和 auth 是并且关系
            ifShow: () => {
              // console.log('ifShow record:'+JSON.stringify(record))
              return record.key !== 'super';
            },
            // 根据权限控制是否显示: 有权限，会显示，支持多个
            // auth: ['basic_list'],
          },
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
            ifShow: () => {
              return record.key !== 'super';
            },
            // auth: ['basic_list'],
          },
          {
            label: '删除',
            // icon: 'ic:outline-delete-outline',
            onClick: handleDelete.bind(null, record),
            // 根据业务控制是否显示 isShow 和 auth 是并且关系
            ifShow: () => {
              return record.key !== 'super';
            },
            // 根据权限控制是否显示: 有权限，会显示，支持多个
            // auth: ['basic_list'],
          },
        ],
      });
    },
  });

  const loadDataTable = async (res: any) => {
    let _params = {
      ...unref(params),
      ...res,
    };
    return await getRoleList(_params);
  };

  function onCheckedRow(rowKeys: any[]) {
    console.log(rowKeys);
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function confirmForm(e: any) {
    console.log('checkedKeys.value:' + JSON.stringify(checkedKeys.value));
    console.log('updatePermissionsParams.value:' + JSON.stringify(updatePermissionsParams.value));
    e.preventDefault();
    formBtnLoading.value = true;
    UpdatePermissions({
      ...{
        id: updatePermissionsParams.value.id,
        menuIds:
          checkedKeys.value === undefined || checkedKeys.value == null ? [] : checkedKeys.value,
      },
    })
      .then((_res) => {
        console.log('_res:' + JSON.stringify(_res));
        message.success('操作成功');
        reloadTable();
        showModal.value = false;
        formBtnLoading.value = false;
      })
      .catch((e: Error) => {
        message.error(e.message ?? '操作失败');
      });
  }

  function confirmForm2(e) {
    e.preventDefault();
    formBtnLoading2.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        console.log('formParams:' + JSON.stringify(formParams));
        Edit(formParams)
          .then((_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            setTimeout(() => {
              showModal2.value = false;
              reloadTable();
            });
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
          });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading2.value = false;
    });
  }

  function addTable() {
    showModal2.value = true;
  }

  function handleEdit(record: Recordable) {
    console.log('点击了编辑', record);
    showModal2.value = true;
    formParams = copyObj(formParams, record);
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
            message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        // message.error('不确定');
      },
    });
  }

  async function handleMenuAuth(record: Recordable) {
    editRoleTitle.value = `分配 ${record.name} 的菜单权限`;
    const data = await GetPermissions({ ...{ id: record.id } });
    console.log('data:' + JSON.stringify(data));
    checkedKeys.value = data.menuIds; //record.menu_keys;
    updatePermissionsParams.value.id = record.id;
    showModal.value = true;
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
      expandedKeys.value = treeData.value.map((item: any) => item.key) as [];
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
    const treeMenuList = await getMenuList();
    expandedKeys.value = treeMenuList.list.map((item) => item.key);
    treeData.value = treeMenuList.list;
  });
</script>

<style lang="less" scoped></style>
