<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="菜单管理"> 在这里可以管理编辑系统下的所有菜单导航</n-card>
    </div>
    <n-grid class="mt-4" cols="1 s:1 m:1 l:3 xl:3 2xl:3" responsive="screen" :x-gap="12">
      <n-gi span="1">
        <n-card :segmented="{ content: true }" :bordered="false" size="small">
          <template #header>
            <n-space>
              <!--              <n-dropdown trigger="hover" @select="selectAddMenu" :options="addMenuOptions">-->
              <!--                <n-button type="info" ghost icon-placement="right">-->
              <!--                  添加菜单-->
              <!--                  <template #icon>-->
              <!--                    <div class="flex items-center">-->
              <!--                      <n-icon size="14">-->
              <!--                        <DownOutlined />-->
              <!--                      </n-icon>-->
              <!--                    </div>-->
              <!--                  </template>-->
              <!--                </n-button>-->
              <!--              </n-dropdown>-->

              <n-button type="info" ghost icon-placement="left" @click="openCreateDrawer">
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <PlusOutlined />
                    </n-icon>
                  </div>
                </template>
                添加菜单
              </n-button>
              <n-button type="info" ghost icon-placement="left" @click="packHandle">
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
              <span style="font-size: 14px">{{
                treeItemTitle ? '' : '从菜单列表选择一项后，进行编辑'
              }}</span>
            </n-space>
          </template>
          <!--          <n-alert type="info" closable> 从菜单列表选择一项后，进行编辑</n-alert>-->
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

            <n-grid x-gap="24" :cols="2">
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
                <n-form-item
                  :label="
                    formParams.type === 1
                      ? '上级目录'
                      : formParams.type === 2
                      ? '上级菜单'
                      : '上级按钮'
                  "
                  path="pid"
                >
                  <n-tree-select
                    :options="optionTreeData"
                    :value="formParams.pid"
                    @update:value="handleUpdateValue"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid x-gap="24" :cols="2">
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
                  <n-input
                    :placeholder="
                      formParams.type === 1
                        ? '目录名称'
                        : formParams.type === 2
                        ? '菜单名称'
                        : '按钮名称'
                    "
                    v-model:value="formParams.title"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="" path="icon">
                  <div style="width: 120px">
                    <span>
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                        </template>
                        请填写图标编码，可以参考图标库，也可以不填使用默认图标
                      </n-tooltip>
                      <span>&nbsp;&nbsp;图标 </span>
                    </span>
                  </div>
                  <n-input placeholder="图标映射路径" v-model:value="formParams.icon" />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid x-gap="24" :cols="2">
              <n-gi>
                <n-form-item label="" path="path">
                  <div style="width: 120px">
                    <span>
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                        </template>
                        路由地址，如：user
                      </n-tooltip>
                      <span>&nbsp;&nbsp;路由地址 </span>
                    </span>
                  </div>
                  <n-input placeholder="路由地址" v-model:value="formParams.path" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="" path="name">
                  <div style="width: 120px">
                    <span>
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                        </template>
                        对应路由配置文件中 `name` 只能是唯一性，配置 `http(s)://` 开头地址
                        则会新窗口打开
                      </n-tooltip>
                      <span>&nbsp;&nbsp;路由别名 </span>
                    </span>
                  </div>
                  <n-input placeholder="路由别名" v-model:value="formParams.name" />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid x-gap="24" :cols="2">
              <n-gi>
                <n-form-item label="" path="component">
                  <div style="width: 120px">
                    <span>
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                        </template>
                        访问的组件路径，如：`/system/menu/menu`，默认在`views`目录下，默认 `LAYOUT`
                        如果是多级菜单 `ParentLayout`
                      </n-tooltip>
                      <span>&nbsp;&nbsp;组件路径 </span>
                    </span>
                  </div>
                  <n-input placeholder="组件路径" v-model:value="formParams.component" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="" path="redirect">
                  <div style="width: 120px">
                    <span>
                      <n-tooltip trigger="hover">
                        <template #trigger>
                          <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                        </template>
                        默认跳转路由地址，如：`/system/menu/menu` 多级路由情况下适用
                      </n-tooltip>
                      <span>&nbsp;&nbsp;默认跳转 </span>
                    </span>
                  </div>
                  <n-input placeholder="默认路由跳转地址" v-model:value="formParams.redirect" />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-divider title-placement="left">功能设置</n-divider>
            <n-form-item label="排序" path="sort">
              <n-input-number v-model:value="formParams.sort" clearable />
            </n-form-item>
            <n-grid x-gap="24" :cols="2">
              <n-gi>
                <n-form-item label="API权限" path="permissions">
                  <n-input
                    placeholder="请输入API权限，多个权限用,分割"
                    v-model:value="formParams.permissions"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <!--                <n-form-item label="权限名称" path="permissionName">-->
                <!--                  <n-input placeholder="权限名称" v-model:value="formParams.permissionName" />-->
                <!--                </n-form-item>-->
                <n-form-item label="高亮路由" path="activeMenu">
                  <n-input placeholder="高亮路由" v-model:value="formParams.activeMenu" />
                </n-form-item>
              </n-gi>
            </n-grid>

            <n-grid x-gap="24" :cols="4">
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

            <n-grid x-gap="24" :cols="4">
              <n-gi>
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
              <n-gi>
                <n-form-item label="外部地址" path="frameSrc" v-show="formParams.isFrame === true">
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
  import { TreeSelectOption, useDialog, useMessage } from 'naive-ui';
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
      value: 0,
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
    label: {
      required: true,
      message: '请输入标题',
      trigger: 'blur',
    },
    path: {
      required: true,
      message: '请输入路径',
      trigger: 'blur',
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

  const formParams = reactive({
    id: 0,
    pid: 0,
    title: '',
    name: '',
    path: '',
    label: '',
    icon: '',
    type: 1,
    redirect: '',
    permissions: '',
    permissionName: '',
    component: '',
    alwaysShow: 1,
    activeMenu: '',
    isRoot: 0,
    isFrame: 0,
    frameSrc: '',
    keepAlive: 0,
    hidden: 0,
    affix: 0,
    status: 1,
    sort: 10,
  });

  function openCreateDrawer() {
    drawerTitle.value = '添加菜单';
    const { openDrawer } = createDrawerRef.value;
    openDrawer();
  }

  function selectedTree(keys) {
    if (keys.length) {
      const treeItem = getTreeItem(unref(treeData), keys[0]);
      treeItemKey.value = keys;
      treeItemTitle.value = treeItem.label;
      Object.assign(formParams, treeItem);
      isEditMenu.value = true;
    } else {
      isEditMenu.value = false;
      treeItemKey.value = [];
      treeItemTitle.value = '';
    }
  }

  function handleDel() {
    dialog.info({
      title: '提示',
      content: `您确定想删除此权限吗?`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        console.log('DeleteMenu formParams:' + JSON.stringify(formParams));
        DeleteMenu({ ...formParams })
          .then(async (_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            // handleReset();
            await loadData();
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
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
        console.log('formParams:' + JSON.stringify(formParams));
        // message.error('抱歉，您没有该权限');
        EditMenu({ ...formParams })
          .then(async (_res) => {
            console.log('_res:' + JSON.stringify(_res));
            message.success('操作成功');
            // handleReset();
            await loadData();
          })
          .catch((e: Error) => {
            message.error(e.message ?? '操作失败');
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
    option: TreeSelectOption | null | Array<TreeSelectOption | null>
  ) {
    formParams.pid = value;
    console.log(value, option);
  }

  onMounted(async () => {
    await loadData();
  });

  async function loadData() {
    const treeMenuList = await getMenuList();
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

  const editConfirm = (val) => {
    console.log(val);
  };
</script>
