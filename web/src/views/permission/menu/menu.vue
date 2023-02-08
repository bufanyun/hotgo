<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="菜单管理">
        在这里可以管理编辑系统下的所有菜单导航和分配相应的菜单权限
      </n-card>
    </div>
    <n-grid class="mt-4" cols="1 s:1 m:1 l:3 xl:3 2xl:3" responsive="screen" :x-gap="12">
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
                添加菜单
              </n-button>
              <n-button
                type="info"
                icon-placement="left"
                @click="openChildCreateDrawer"
                :disabled="!isEditMenu"
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
            <n-input type="input" v-model:value="pattern" placeholder="输入菜单名称搜索">
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
        <n-card :segmented="{ content: true }" :bordered="false" size="small">
          <template #header>
            <n-space>
              <n-icon size="18">
                <FormOutlined />
              </n-icon>
              <span>编辑菜单{{ treeItemTitle ? `：${treeItemTitle}` : '' }}</span>
              <span style="font-size: 14px">{{ treeItemTitle }}</span>
            </n-space>
          </template>

          <n-result
            v-show="!isEditMenu"
            status="info"
            title="提示"
            description="从菜单列表中选择一项进行编辑"
          />
          <n-form
            :model="formParams"
            :rules="rules"
            ref="formRef"
            label-placement="left"
            :label-width="100"
            v-if="isEditMenu"
            class="py-4"
          >
            <n-divider title-placement="left">基本设置</n-divider>

            <n-grid cols="2 300:1 600:2">
              <n-gi>
                <n-form-item label="类型" path="type">
                  <n-radio-group v-model:value="formParams.type" name="type">
                    <n-radio-button
                      v-for="menuType in menuTypes"
                      :key="menuType.value"
                      :value="menuType.value"
                      :label="menuType.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="上级目录" path="pid">
                  <n-tree-select
                    :options="optionTreeData"
                    :value="formParams.pid"
                    @update:value="handleUpdateValue"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid cols="2 300:1 600:2">
              <n-gi>
                <n-form-item
                  :label="
                    formParams.type === 1
                      ? '目录名称'
                      : formParams.type === 2
                      ? '菜单名称'
                      : '按钮名称'
                  "
                  path="title"
                >
                  <n-input placeholder="请输入" v-model:value="formParams.title" />
                </n-form-item>
              </n-gi>
              <n-gi v-if="formParams.type !== 3">
                <n-form-item path="icon">
                  <IconSelector style="width: 100%" v-model:value="formParams.icon" option="antd" />
                  <template #label>
                    <n-tooltip trigger="hover">
                      <template #trigger>
                        <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                      </template>
                      请填写图标编码，可以参考图标库，也可以不填使用默认图标
                    </n-tooltip>
                    菜单图标
                  </template>
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid cols="2 300:1 600:2">
              <n-gi v-if="formParams.type !== 3">
                <n-form-item path="path">
                  <n-input placeholder="路由地址" v-model:value="formParams.path" />
                  <template #label>
                    <n-tooltip trigger="hover">
                      <template #trigger>
                        <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                      </template>
                      请路由地址，如：user
                    </n-tooltip>
                    路由地址
                  </template>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item path="name">
                  <n-input placeholder="路由别名" v-model:value="formParams.name" />
                  <template #label>
                    <n-tooltip trigger="hover">
                      <template #trigger>
                        <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                      </template>
                      对应路由配置文件中 `name` 只能是唯一性，配置 `http(s)://` 开头地址
                      则会新窗口打开
                    </n-tooltip>
                    路由别名
                  </template>
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid cols="2 300:1 600:2" v-if="formParams.type !== 3">
              <n-gi>
                <n-form-item label="组件路径" path="component">
                  <n-input placeholder="组件路径" v-model:value="formParams.component" />
                  <template #feedback>
                    主目录填 `LAYOUT`;多级父目录填
                    `ParentLayout`;页面填具体的组件路径，如：`/system/menu/menu`
                  </template>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formParams.type === 1">
                <n-form-item label="默认跳转" path="redirect">
                  <n-input placeholder="默认路由跳转地址" v-model:value="formParams.redirect" />
                  <template #feedback
                    >默认跳转路由地址，如：`/system/menu/menu` 多级路由情况下适用
                  </template>
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-divider title-placement="left">功能设置</n-divider>

            <n-grid cols="1 ">
              <n-gi>
                <n-form-item label="分配权限" path="permissions">
                  <n-input
                    placeholder="请输入分配权限，多个权限用,分割"
                    v-model:value="formParams.permissions"
                  />
                  <template #label>
                    <n-tooltip trigger="hover">
                      <template #trigger>
                        <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                      </template>
                      请填写API路由地址，可同时作用于服务端和web端。多个权限用,分割
                    </n-tooltip>
                    分配权限
                  </template>
                </n-form-item>
              </n-gi>
              <!--              <n-gi>-->
              <!--                <n-form-item label="权限名称" path="permissionName">-->
              <!--                  <n-input placeholder="权限名称" v-model:value="formParams.permissionName" />-->
              <!--                  <template #feedback>分配权限存在多个时，权限名称只绑定到第一个权限</template>-->
              <!--                </n-form-item>-->
              <!--              </n-gi>-->
            </n-grid>

            <n-grid cols="2 300:1 600:2">
              <n-gi v-if="formParams.type !== 3">
                <n-form-item label="高亮路由" path="activeMenu">
                  <n-input placeholder="高亮路由" v-model:value="formParams.activeMenu" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="菜单排序" path="sort">
                  <n-input-number style="width: 100%" v-model:value="formParams.sort" clearable />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid cols="4 300:1 400:2 600:3 800:4" v-if="formParams.type !== 3">
              <n-gi>
                <n-form-item label="根路由" path="isRoot">
                  <n-radio-group v-model:value="formParams.isRoot" name="isRoot">
                    <n-radio-button
                      v-for="switchStatus in switchStatusMap"
                      :key="switchStatus.value"
                      :value="switchStatus.value"
                      :label="switchStatus.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="简化路由" path="alwaysShow">
                  <n-radio-group v-model:value="formParams.alwaysShow" name="alwaysShow">
                    <n-radio-button
                      v-for="switchStatus in switchStatusMap"
                      :key="switchStatus.value"
                      :value="switchStatus.value"
                      :label="switchStatus.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="缓存路由" path="keepAlive">
                  <n-radio-group v-model:value="formParams.keepAlive" name="keepAlive">
                    <n-radio-button
                      v-for="switchStatus in switchStatusMap"
                      :key="switchStatus.value"
                      :value="switchStatus.value"
                      :label="switchStatus.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="是否隐藏" path="hidden">
                  <n-radio-group v-model:value="formParams.hidden" name="hidden">
                    <n-radio-button
                      v-for="hidden in hiddenMap"
                      :key="hidden.value"
                      :value="hidden.value"
                      :label="hidden.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid cols="4 300:1 400:2 600:3 800:4">
              <n-gi>
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formParams.status" name="status">
                    <n-radio-button
                      v-for="status in statusMap"
                      :key="status.value"
                      :value="status.value"
                      :label="status.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formParams.type !== 3">
                <n-form-item label="是否外链" path="isFrame">
                  <n-radio-group v-model:value="formParams.isFrame" name="isFrame">
                    <n-radio-button
                      v-for="switchStatus in switchStatusMap"
                      :key="switchStatus.value"
                      :value="switchStatus.value"
                      :label="switchStatus.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi v-if="formParams.type !== 3">
                <n-form-item label="外部地址" path="frameSrc" v-show="formParams.isFrame === 1">
                  <n-input placeholder="内联外部地址" v-model:value="formParams.frameSrc" />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-form-item path="auth" style="margin-left: 100px">
              <n-space>
                <n-button type="primary" :loading="subLoading" @click="formSubmit"
                  >保存修改
                </n-button>
                <n-button @click="handleReset">重置</n-button>
                <n-button @click="handleDel">删除</n-button>
              </n-space>
            </n-form-item>
          </n-form>
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
  import { FormItemRule, TreeSelectOption, useDialog, useMessage } from 'naive-ui';
  import {
    AlignLeftOutlined,
    FormOutlined,
    PlusOutlined,
    QuestionCircleOutlined,
    SearchOutlined,
  } from '@vicons/antd';
  import { DeleteMenu, EditMenu, getMenuList } from '@/api/system/menu';
  import { getTreeItem } from '@/utils';
  import CreateDrawer from './CreateDrawer.vue';
  import IconSelector from '@/components/IconSelector/index.vue';
  import { newState, State } from '@/views/permission/menu/model';

  const menuTypes = [
    {
      value: 1,
      label: '目录',
    },
    {
      value: 2,
      label: '菜单',
    },
    {
      value: 3,
      label: '按钮',
    },
  ].map((s) => {
    return s;
  });

  const switchStatusMap = [
    {
      value: 2,
      label: '关闭',
    },
    {
      value: 1,
      label: '开启',
    },
  ].map((s) => {
    return s;
  });

  const statusMap = [
    {
      value: 0,
      label: '禁用',
    },
    {
      value: 1,
      label: '启用',
    },
  ].map((s) => {
    return s;
  });

  const hiddenMap = [
    {
      value: 0,
      label: '否',
    },
    {
      value: 1,
      label: '是',
    },
  ].map((s) => {
    return s;
  });

  const rules = {
    title: {
      required: true,
      message: '请输入名称',
      trigger: 'blur',
    },
    label: {
      required: true,
      message: '请输入标题',
      trigger: 'blur',
    },
    path: {
      required: false,
      message: '请输入路由地址',
      trigger: 'blur',
      validator: function (_rule: FormItemRule, value: any, callback: Function) {
        if (formParams.type != 3 && !value) {
          callback(new Error('请输入路由地址'));
        }
      },
    },
  };

  const formRef: any = ref(null);
  const createDrawerRef = ref();
  const message = useMessage();
  const dialog = useDialog();
  let treeItemKey = ref([]);
  let expandedKeys = ref([]);
  const treeData = ref([]);
  const loading = ref(true);
  const subLoading = ref(false);
  const isEditMenu = ref(false);
  const treeItemTitle = ref('');
  const pattern = ref('');
  const drawerTitle = ref('');
  const optionTreeData = ref<any>([]);

  const formParams = reactive<State>(newState(null));

  function openCreateDrawer() {
    drawerTitle.value = '添加菜单';
    const { openDrawer } = createDrawerRef.value;
    openDrawer(0);
  }

  function openChildCreateDrawer() {
    drawerTitle.value = '添加菜单';
    const { openDrawer } = createDrawerRef.value;
    openDrawer(formParams.id);
  }

  function selectedTree(keys) {
    if (keys.length) {
      const treeItem = getTreeItem(unref(treeData), keys[0]);
      treeItemKey.value = keys;
      treeItemTitle.value = treeItem.label + ' #' + treeItem.id;
      Object.assign(formParams, treeItem);
      isEditMenu.value = true;
    } else {
      isEditMenu.value = false;
      treeItemKey.value = [];
      treeItemTitle.value = '';
    }
  }

  function handleDel() {
    dialog.warning({
      title: '提示',
      content: `您确定要删除此菜单吗?`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteMenu({ ...formParams }).then(async (_res) => {
          message.success('操作成功');
          await loadData();
        });
      },
      onNegativeClick: () => {
        message.error('已取消');
      },
    });
  }

  function handleReset() {
    const treeItem = getTreeItem(unref(treeData), treeItemKey.value[0]);
    Object.assign(formParams, treeItem);
  }

  function formSubmit() {
    formRef.value.validate((errors: boolean) => {
      if (!errors) {
        subLoading.value = true;
        EditMenu({ ...formParams })
          .then(async (_res) => {
            subLoading.value = false;
            message.success('操作成功');
            await loadData();
          })
          .catch((_e: Error) => {
            subLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  function packHandle() {
    if (expandedKeys.value.length) {
      expandedKeys.value = [];
    } else {
      expandedKeys.value = unref(treeData).map((item: any) => item.key as string) as [];
    }
  }

  // 处理选项更新
  function handleUpdateValue(
    value: string | number | Array<string | number> | null,
    _option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    formParams.pid = value as number;
  }

  onMounted(async () => {
    await loadData();
  });

  async function loadData() {
    const treeMenuList = await getMenuList();
    const keys = treeMenuList.list.map((item) => item.key);
    Object.assign(formParams, keys);
    treeData.value = [];
    treeData.value = treeMenuList.list;
    optionTreeData.value = [
      {
        id: 0,
        key: 0,
        label: '根目录',
        pid: 0,
        title: '根目录',
        type: 1,
      },
    ];
    optionTreeData.value = optionTreeData.value.concat(treeMenuList.list);
    loading.value = false;
  }

  function onExpandedKeys(keys) {
    expandedKeys.value = keys;
  }
</script>
