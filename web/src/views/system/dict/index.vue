<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="字典管理">
        可代替后台管理系统，设置的大量枚举值和配置，统一标准化管理，随时修改或增加
      </n-card>
    </div>
    <n-grid class="mt-6" cols="1 s:1 m:1 l:4 xl:4 2xl:4" responsive="screen" :x-gap="12">
      <n-gi span="1">
        <n-card :segmented="{ content: true }" :bordered="false" size="small">
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
                添加
              </n-button>
              <n-button
                type="info"
                icon-placement="left"
                @click="openEditDrawer"
                :disabled="formParams.id === 0"
              >
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <EditOutlined />
                    </n-icon>
                  </div>
                </template>
                编辑
              </n-button>
              <n-button type="error" icon-placement="left" @click="handleDel">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <DeleteOutlined />
                    </n-icon>
                  </div>
                </template>
                删除
              </n-button>
              <n-button type="info" icon-placement="left" @click="packHandle">
                {{ expandedKeys.length ? '收起' : '展开' }}
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
            <n-input type="input" v-model:value="pattern" placeholder="输入字典名称搜索">
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
                  :data="treeData"
                  :expandedKeys="expandedKeys"
                  style="max-height: 95%; overflow: hidden"
                  @update:selected-keys="selectedTree"
                  @update:expanded-keys="onExpandedKeys"
                />
              </template>
            </div>
          </div>
        </n-card>
      </n-gi>
      <n-gi span="3">
        <n-card :segmented="{ content: true }" :bordered="false" size="small">
          <template #header>
            <n-space>
              <n-icon size="18">
                <FormOutlined />
              </n-icon>
              <span>编辑字典{{ treeItemTitle ? `：${treeItemTitle}` : '' }}</span>
              <span style="font-size: 14px">{{
                treeItemTitle ? '' : '从列表选择一项后，进行编辑'
              }}</span>
            </n-space>
          </template>

          <List :checkedId="checkedId" />
        </n-card>
      </n-gi>
    </n-grid>
    <CreateDrawer
      ref="createDrawerRef"
      :title="drawerTitle"
      :optionTreeData="optionTreeData"
      @loadData="loadData"
    />
  </div>
</template>
<script lang="ts" setup>
  import { onMounted, reactive, ref, unref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import {
    AlignLeftOutlined,
    FormOutlined,
    PlusOutlined,
    EditOutlined,
    SearchOutlined,
    DeleteOutlined,
  } from '@vicons/antd';
  import { getTreeItem } from '@/utils';
  import CreateDrawer from './CreateDrawer.vue';
  import List from './list.vue';
  import { DeleteDict, getDictTree } from '@/api/dict/dict';

  const createDrawerRef = ref();
  const message = useMessage();
  const dialog = useDialog();
  let treeItemKey = ref([]);
  let expandedKeys = ref([]);
  const treeData = ref([]);
  const loading = ref(true);
  const isEditMenu = ref(false);
  const treeItemTitle = ref('');
  const checkedId = ref(0);
  const pattern = ref('');
  const drawerTitle = ref('');
  const optionTreeData = ref([
    {
      id: 0,
      key: 0,
      label: '根目录',
      pid: 0,
      title: '根目录',
      type: 1,
    },
  ]);
  const defaultValueRef = () => ({
    id: 0,
    pid: 0,
    type: '',
    name: '',
    remark: '',
    status: 1,
    sort: 10,
  });

  const formParams = reactive(defaultValueRef());

  function openCreateDrawer() {
    drawerTitle.value = '添加字典类型';
    const { openDrawer } = createDrawerRef.value;
    openDrawer(defaultValueRef());
  }

  function openEditDrawer() {
    drawerTitle.value = '编辑字典类型';
    const { openDrawer } = createDrawerRef.value;
    openDrawer(formParams);
  }

  function selectedTree(keys) {
    if (keys.length) {
      const treeItem = getTreeItem(unref(treeData), keys[0]);
      // console.log('选择treeItem:' + JSON.stringify(treeItem));
      treeItemKey.value = keys;
      treeItemTitle.value = treeItem.label;
      Object.assign(formParams, treeItem);
      isEditMenu.value = true;
      checkedId.value = treeItem.id;
    } else {
      isEditMenu.value = false;
      treeItemKey.value = [];
      treeItemTitle.value = '';
    }
  }

  function handleDel() {
    dialog.info({
      title: '提示',
      content: `您确定想删除此类型吗?`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        console.log('DeleteMenu formParams:' + JSON.stringify(formParams));
        DeleteDict({ ...formParams })
          .then(async (_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            // handleReset();
            await loadData();
          })
          .catch((_e: Error) => {
            // message.error(e.message ?? '操作失败');
          });
      },
      onNegativeClick: () => {
        message.error('已取消');
      },
    });
  }

  function packHandle() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = unref(treeData).map((item: any) => item.key as string) as [];
    }
  }

  onMounted(async () => {
    await loadData();
  });

  async function loadData() {
    const treeMenuList = await getDictTree();
    const keys = treeMenuList.list.map((item) => item.key);
    Object.assign(formParams, keys);
    treeData.value = [];
    optionTreeData.value = [];
    treeData.value = treeMenuList.list;
    optionTreeData.value = optionTreeData.value.concat(treeMenuList.list);
    loading.value = false;
  }

  function onExpandedKeys(keys) {
    expandedKeys.value = keys;
  }
</script>
