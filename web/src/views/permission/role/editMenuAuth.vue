<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :show-icon="false" preset="dialog">
      <n-drawer-content closable :title="`分配 ${formValue.name} 的菜单权限`">
        <n-spin :show="loading" description="请稍候...">
          <div :style="{ maxHeight: '78vh', height: '78vh' }">
            <n-input v-model:value="pattern" placeholder="输入菜单名称或权限路径搜索" class="mb-2">
              <template #suffix>
                <n-icon size="18" class="cursor-pointer">
                  <SearchOutlined />
                </n-icon>
              </template>
            </n-input>
            <n-tree
              block-line
              checkable
              check-on-click
              default-expand-all
              virtual-scroll
              :data="treeData"
              :pattern="pattern"
              :filter="filterTreeNode"
              :expandedKeys="expandedKeys"
              :checked-keys="checkedKeys"
              style="max-height: 950px; overflow: hidden"
              @update:checked-keys="checkedTree"
              @update:expanded-keys="onExpandedKeys"
            />
          </div>
        </n-spin>
        <template #footer>
          <n-space v-if="showImportSelect">
            <n-input-group>
              <n-tree-select
                placeholder="请选择一个要导入的角色"
                :consistent-menu-width="false"
                clearable
                filterable
                default-expand-all
                :options="editRoleOption"
                key-field="id"
                label-field="name"
                :on-update:value="handleImportSelect"
                style="width: 300px"
              />
              <div class="mr-2"></div>
              <n-button ghost @click="showImportSelect = false"> 取消 </n-button>
            </n-input-group>
          </n-space>

          <n-space v-if="!showImportSelect">
            <n-button ghost @click="showImportSelect = true"> 导入权限 </n-button>
            <n-button type="info" ghost icon-placement="left" @click="packHandle">
              全部{{ expandedKeys.length ? '收起' : '展开' }}
            </n-button>
            <n-button type="info" ghost icon-placement="left" @click="checkedAllHandle">
              全部{{ checkedAll ? '取消' : '选择' }}
            </n-button>

            <n-popconfirm @positive-click="confirmForm">
              <template #trigger>
                <n-button type="primary" :loading="formBtnLoading">提交</n-button>
              </template>
              你正在修改 {{ formValue.name }} 的菜单权限，确定要提交吗？
            </n-popconfirm>
          </n-space>
        </template>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { GetPermissions, getRoleList, UpdatePermissions } from '@/api/system/role';
  import { NButton, useMessage } from 'naive-ui';
  import { adaModalWidth, getTreeKeys } from '@/utils/hotgo';
  import { findTreeNode, getAllExpandKeys } from '@/utils';
  import { getMenuList } from '@/api/system/menu';
  import { SearchOutlined } from '@vicons/antd';
  import { State, newState } from '@/views/permission/role/model';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formBtnLoading = ref(false);
  const rawRoleOption = ref<State[]>([]);
  const checkedAll = ref(false);
  const treeData = ref([]);
  const expandedKeys = ref<any[]>([]);
  const checkedKeys = ref<any[]>([]);
  const pattern = ref('');
  const showImportSelect = ref(false);

  const editRoleOption = computed<State[]>(() => {
    if (!rawRoleOption.value) {
      return [];
    }
    const role = findTreeNode(rawRoleOption.value, formValue.value.id, 'id');
    if (role) {
      role.disabled = true;
    }
    return rawRoleOption.value;
  });

  const dialogWidth = computed(() => {
    return adaModalWidth(630);
  });

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    const params = {
      id: formValue.value.id,
      menuIds: checkedKeys.value ?? [],
    };
    UpdatePermissions(params)
      .then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          showModal.value = false;
          emit('reloadTable');
        });
      })
      .finally(() => {
        formBtnLoading.value = false;
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
      checkedKeys.value = getTreeKeys(treeData.value);
      checkedAll.value = true;
    } else {
      checkedKeys.value = [];
      checkedAll.value = false;
    }
  }

  // 按名称和权限搜索
  function filterTreeNode(pattern: string, node: any) {
    if (!pattern) return true;
    const searchText = pattern.toLowerCase();

    const label = (node.label || node.title || '').toLowerCase();
    if (label.includes(searchText)) {
      return true;
    }

    const permissions = node.permissions || '';
    if (permissions) {
      const permissionsLower = permissions.toLowerCase();
      if (permissionsLower.includes(searchText)) {
        return true;
      }
    }
    return false;
  }

  function handleImportSelect(key: number) {
    showImportSelect.value = false;
    showModal.value = true;
    getPermissions(key);

    // 默认全部展开
    expandedKeys.value = getAllExpandKeys(treeData);
    message.success('导入成功，提交后生效');
  }

  async function loadMenuList() {
    const res = await getMenuList();
    expandedKeys.value = getAllExpandKeys(res.list) as [];
    treeData.value = res.list;
  }

  async function getPermissions(id: number) {
    checkedKeys.value = [];
    checkedAll.value = false;
    const res = await GetPermissions({ id: id });
    checkedKeys.value = res.menuIds;
  }

  async function loadDataList() {
    const res = await getRoleList({ pageSize: 100, page: 1 });
    rawRoleOption.value = res.list;
  }

  async function openModal(record: Recordable) {
    loading.value = true;
    formValue.value = newState(record);
    showModal.value = true;
    await loadMenuList();
    await getPermissions(record.id);
    await loadDataList();
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>
