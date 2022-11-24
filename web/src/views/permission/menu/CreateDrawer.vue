<template>
  <n-drawer v-model:show="isDrawer" :width="width" :placement="placement">
    <n-drawer-content :title="title" closable>
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="100"
      >
        <n-divider title-placement="left">基本设置</n-divider>
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
        <n-form-item
          :label="
            formParams.type === 1 ? '上级目录' : formParams.type === 2 ? '上级菜单' : '上级按钮'
          "
          path="pid"
        >
          <n-tree-select
            :options="optionTreeData"
            default-value="0"
            @update:value="handleUpdateValue"
          />
        </n-form-item>
        <n-form-item
          :label="
            formParams.type === 1 ? '目录名称' : formParams.type === 2 ? '菜单名称' : '按钮名称'
          "
          path="title"
        >
          <n-input
            :placeholder="
              formParams.type === 1 ? '目录名称' : formParams.type === 2 ? '菜单名称' : '按钮名称'
            "
            v-model:value="formParams.title"
          />
        </n-form-item>
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
        <n-form-item label="" path="name">
          <div style="width: 120px">
            <span>
              <n-tooltip trigger="hover">
                <template #trigger>
                  <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                </template>
                对应路由配置文件中 `name` 只能是唯一性，配置 `http(s)://` 开头地址 则会新窗口打开
              </n-tooltip>
              <span>&nbsp;&nbsp;路由别名 </span>
            </span>
          </div>
          <n-input placeholder="路由别名" v-model:value="formParams.name" />
        </n-form-item>
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
        <n-form-item label="" path="redirect" v-show="formParams.type === 1">
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
        <n-divider title-placement="left">功能设置</n-divider>
        <n-form-item label="API权限" path="permissions">
          <n-input
            placeholder="请输入API权限，多个权限用,分割"
            v-model:value="formParams.permissions"
          />
        </n-form-item>
        <!--        <n-form-item label="权限名称" path="permissionName">-->
        <!--          <n-input placeholder="权限名称" v-model:value="formParams.permissionName" />-->
        <!--        </n-form-item>-->
        <n-form-item label="高亮路由" path="activeMenu">
          <n-input placeholder="高亮路由" v-model:value="formParams.activeMenu" />
        </n-form-item>
        <n-form-item label="排序" path="sort">
          <n-input-number v-model:value="formParams.sort" clearable />
        </n-form-item>

        <n-grid x-gap="24" :cols="2">
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
        </n-grid>

        <n-grid x-gap="24" :cols="2">
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

        <n-grid x-gap="24" :cols="2">
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
        </n-grid>

        <n-grid x-gap="24" :cols="2">
          <n-gi>
            <n-form-item label="外部地址" path="frameSrc" v-show="formParams.isFrame === true">
              <n-input placeholder="内联外部地址" v-model:value="formParams.frameSrc" />
            </n-form-item>
          </n-gi>
          <n-gi />
        </n-grid>
      </n-form>

      <template #footer>
        <n-space>
          <n-button type="primary" :loading="subLoading" @click="formSubmit">提交</n-button>
          <n-button @click="handleReset">重置</n-button>
        </n-space>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<script lang="ts">
  import { defineComponent, reactive, ref, toRefs } from 'vue';
  import { TreeSelectOption, useMessage } from 'naive-ui';
  import { QuestionCircleOutlined } from '@vicons/antd';
  import { EditMenu } from '@/api/system/menu';

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
  export default defineComponent({
    name: 'CreateDrawer',
    components: {},
    props: {
      title: {
        type: String,
        default: '添加顶级菜单',
      },
      optionTreeData: {
        type: Object,
        // eslint-disable-next-line vue/require-valid-default-prop
        default: [],
      },
    },
    emits: ['loadData'],
    setup(_props, context) {
      const message = useMessage();
      const formRef: any = ref(null);
      const defaultValueRef = () => ({
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

      const state = reactive({
        width: 700,
        isDrawer: false,
        subLoading: false,
        formParams: defaultValueRef(),
        placement: 'right',
        icon: '',
        alertText:
          '该功能主要实时预览各种布局效果，更多完整配置在 projectSetting.ts 中设置，建议在生产环境关闭该布局预览功能。',
      });

      function openDrawer() {
        if (document.body.clientWidth < 700) {
          state.width = document.body.clientWidth;
        }
        state.isDrawer = true;
      }

      function closeDrawer() {
        state.isDrawer = false;
      }

      function formSubmit() {
        formRef.value.validate((errors) => {
          if (!errors) {
            console.log('state.formParams:' + JSON.stringify(state.formParams));
            EditMenu({ ...state.formParams })
              .then(async (_res) => {
                console.log('_res:' + JSON.stringify(_res));
                message.success('操作成功');
                handleReset();
                await context.emit('loadData');
                closeDrawer();
              })
              .catch((e: Error) => {
                message.error(e.message ?? '操作失败');
              });
          } else {
            message.error('请填写完整信息');
          }
        });
      }

      function handleReset() {
        formRef.value.restoreValidation();
        state.formParams = Object.assign(state.formParams, defaultValueRef());
      }

      // 处理选项更新
      function handleUpdateValue(
        value: string | number | Array<string | number> | null,
        option: TreeSelectOption | null | Array<TreeSelectOption | null>
      ) {
        console.log(value, option);
        state.formParams.pid = value;
      }

      return {
        ...toRefs(state),
        formRef,
        rules,
        formSubmit,
        handleReset,
        openDrawer,
        closeDrawer,
        menuTypes,
        switchStatusMap,
        statusMap,
        hiddenMap,
        handleUpdateValue,
        QuestionCircleOutlined,
      };
    },
  });
</script>
