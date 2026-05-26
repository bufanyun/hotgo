<template>
  <div>
    <div class="n-layout-page-header" v-if="!isModal">
      <n-card :bordered="false" title="菜单管理">
        在这里可以管理编辑系统下的所有菜单导航和分配相应的菜单权限
      </n-card>
    </div>
    <n-grid
      cols="1 s:1 m:1 l:3 xl:3 2xl:3"
      responsive="screen"
      :x-gap="12"
    >
      <n-gi span="1">
        <n-card :segmented="{ content: true }" :bordered="false" size="small" class="proCard">
          <template #header>
            <n-space>
              <n-button type="info" icon-placement="left" @click="openCreateDrawer">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <PlusOutlined />
                    </n-icon>
                  </div>
                </template>
                添加菜单
              </n-button>
              <n-button
                type="info"
                icon-placement="left"
                @click="openChildCreateDrawer"
                :disabled="formParams.id == 0"
              >
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <PlusOutlined />
                    </n-icon>
                  </div>
                </template>
                添加子菜单
              </n-button>
              <n-button type="primary" icon-placement="left" @click="packHandle">
                全部{{ expandedKeys.length ? '收起' : '展开' }}
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <AlignLeftOutlined />
                    </n-icon>
                  </div>
                </template>
              </n-button>
            </n-space>
          </template>
          <div class="w-full menu">
            <n-input type="text" v-model:value="pattern" placeholder="输入菜单名称或权限路径搜索">
              <template #suffix>
                <n-icon size="18" class="cursor-pointer">
                  <SearchOutlined />
                </n-icon>
              </template>
            </n-input>
            <div class="py-3 menu-list">
              <template v-if="loading">
                <div class="flex items-center justify-center py-4">
                  <n-spin size="medium" />
                </div>
              </template>
              <template v-else>
                <n-tree
                  block-line
                  cascade
                  checkable
                  :virtual-scroll="true"
                  :pattern="pattern"
                  :filter="filterTreeNode"
                  :data="treeOption"
                  :expandedKeys="expandedKeys"
                  style="max-height: 650px; overflow: hidden"
                  @update:selected-keys="selectedTree"
                  @update:expanded-keys="onExpandedKeys"
                />
              </template>
            </div>
          </div>
        </n-card>
      </n-gi>
      <n-gi span="2">
        <n-card :segmented="{ content: true }" :bordered="false" size="small" class="proCard">
          <template #header>
            <n-space>
              <n-icon size="18">
                <FormOutlined />
              </n-icon>
              <span>编辑菜单{{ treeItemTitle ? `：${treeItemTitle}` : '' }}</span>
            </n-space>
          </template>

          <n-result
            v-show="formParams.id == 0"
            status="info"
            title="提示"
            description="从菜单列表中选择一项进行编辑"
          />
          <EditForm
            v-if="formParams.id > 0"
            v-model:formParams="formParams"
            v-model:treeOption="treeOption"
            @reloadTable="loadTreeOption"
          />
        </n-card>
      </n-gi>
    </n-grid>
    <AddModal ref="addModalRef" v-model:treeOption="treeOption" @reloadTable="loadTreeOption" />
  </div>
</template>
<script lang="ts" setup>
  import { computed, onMounted, ref, unref } from 'vue';
  import { AlignLeftOutlined, FormOutlined, PlusOutlined, SearchOutlined } from '@vicons/antd';
  import { getMenuList } from '@/api/system/menu';
  import { newState, State, loadOptions } from '@/views/permission/menu/model';
  import EditForm from '@/views/permission/menu/editForm.vue';
  import AddModal from '@/views/permission/menu/addModal.vue';

  const addModalRef = ref();
  const loading = ref(false);
  const treeOption = ref([]);
  const pattern = ref('');
  const expandedKeys = ref([]);
  const formParams = ref<State>(newState(null));
  const treeItemTitle = computed(() => {
    if (formParams.value.id > 0) {
      return formParams.value.label + ' #' + formParams.value.id;
    }
    return '';
  });
  const isModal = defineModel<boolean>('isModal', { default: false });

  function openCreateDrawer() {
    addModalRef.value.openModal(null);
  }

  function openChildCreateDrawer() {
    const state = newState(null);
    state.pid = formParams.value.id;
    state.type = formParams.value.type;
    addModalRef.value.openModal(state);
  }

  function selectedTree(keys: number[], option: any[]) {
    let item = null;
    if (keys.length) {
      item = option[0];
    }
    formParams.value = newState(item);
  }

  function packHandle() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = unref(treeOption).map((item: any) => item.key as string) as [];
    }
  }

  function onExpandedKeys(keys) {
    expandedKeys.value = keys;
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

  // 加载菜单选项树
  function loadTreeOption() {
    const needLoading = treeOption.value.length == 0;
    if (needLoading) {
      loading.value = true;
    }

    getMenuList()
      .then((res) => {
        if (res.list && res.list.length > 0) {
          treeOption.value = res.list;
        } else {
          treeOption.value = [];
        }
      })
      .finally(() => {
        if (needLoading) {
          loading.value = false;
        }
      });
  }

  onMounted(() => {
    loadTreeOption();
    loadOptions();
  });
</script>
