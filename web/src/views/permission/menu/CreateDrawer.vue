<template>
  <n-drawer v-model:show="isDrawer" :width="width" :placement="placement" :mask-closable="false">
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
        <n-form-item label="上级目录" path="pid">
          <n-tree-select
            filterable
            :options="optionTreeData"
            :default-value="formParams.pid"
            @update:value="handleUpdateValue"
          />
        </n-form-item>
        <n-form-item
          :label="
            formParams.type === 1 ? '目录名称' : formParams.type === 2 ? '菜单名称' : '按钮名称'
          "
          path="title"
        >
          <n-input placeholder="请输入" v-model:value="formParams.title" />
        </n-form-item>

        <n-form-item path="icon" v-if="formParams.type !== 3">
          <IconSelector style="width: 100%" v-model:value="formParams.icon" option="antd" />
          <template #label>
            <n-tooltip trigger="hover">
              <template #trigger>
                <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
              </template>
              请填写图标编码，可以参考图标库，也可以不填使用默认图标
            </n-tooltip>
            菜单图标</template
          >
        </n-form-item>
        <n-form-item path="path" v-if="formParams.type !== 3">
          <n-input placeholder="路由地址" v-model:value="formParams.path" />
          <template #label>
            <n-tooltip trigger="hover">
              <template #trigger>
                <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
              </template>
              请路由地址，如：user
            </n-tooltip>
            路由地址</template
          >
        </n-form-item>
        <n-form-item path="name">
          <n-input placeholder="路由别名" v-model:value="formParams.name" />
          <template #label>
            <n-tooltip trigger="hover">
              <template #trigger>
                <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
              </template>
              对应路由配置文件中 `name` 只能是唯一性，配置 `http(s)://` 开头地址 则会新窗口打开
            </n-tooltip>
            路由别名</template
          >
        </n-form-item>
        <n-form-item label="组件路径" path="component" v-if="formParams.type !== 3">
          <n-input placeholder="组件路径" v-model:value="formParams.component" />
          <template #feedback>
            主目录填 `LAYOUT`;多级父目录填
            `ParentLayout`;页面填具体的组件路径，如：`/system/menu/menu`</template
          >
        </n-form-item>
        <n-form-item label="默认跳转" path="redirect" v-if="formParams.type === 1">
          <n-input placeholder="默认路由跳转地址" v-model:value="formParams.redirect" />
          <template #feedback
            >默认跳转路由地址，如：`/system/menu/menu` 多级路由情况下适用</template
          >
        </n-form-item>
        <n-divider title-placement="left">功能设置</n-divider>
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
            分配权限</template
          >
        </n-form-item>
        <!--                <n-form-item label="权限名称" path="permissionName">-->
        <!--                  <n-input placeholder="权限名称" v-model:value="formParams.permissionName" />-->
        <!--                </n-form-item>-->
        <n-form-item label="高亮路由" path="activeMenu" v-if="formParams.type !== 3">
          <n-input placeholder="高亮路由" v-model:value="formParams.activeMenu" />
        </n-form-item>
        <n-form-item label="菜单排序" path="sort">
          <n-input-number style="width: 100%" v-model:value="formParams.sort" clearable />
        </n-form-item>

        <n-grid x-gap="24" :cols="2" v-if="formParams.type !== 3">
          <n-gi>
            <n-form-item label="根路由" path="isRoot">
              <n-radio-group v-model:value="formParams.isRoot" name="isRoot">
                <n-radio-button
                  v-for="switchStatus in statusMap"
                  :key="switchStatus.value"
                  :value="switchStatus.value"
                  :label="switchStatus.label"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="页签固定" path="affix">
              <n-radio-group v-model:value="formParams.affix" name="affix">
                <n-radio-button
                  v-for="switchStatus in statusMap"
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
                  v-for="switchStatus in statusMap"
                  :key="switchStatus.value"
                  :value="switchStatus.value"
                  :label="switchStatus.label"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
        </n-grid>

        <n-grid x-gap="24" :cols="2" v-if="formParams.type !== 3">
          <n-gi>
            <n-form-item label="缓存路由" path="keepAlive">
              <n-radio-group v-model:value="formParams.keepAlive" name="keepAlive">
                <n-radio-button
                  v-for="switchStatus in statusMap"
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

        <n-grid x-gap="24" :cols="2" v-if="formParams.type !== 3">
          <n-gi>
            <n-form-item label="是否外链" path="isFrame">
              <n-radio-group v-model:value="formParams.isFrame" name="isFrame">
                <n-radio-button
                  v-for="switchStatus in statusMap"
                  :key="switchStatus.value"
                  :value="switchStatus.value"
                  :label="switchStatus.label"
                />
              </n-radio-group>
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="外部地址" path="frameSrc" v-show="formParams.isFrame === 1">
              <n-input placeholder="内联外部地址" v-model:value="formParams.frameSrc" />
            </n-form-item>
          </n-gi>
        </n-grid>

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
      </n-form>

      <template #footer>
        <n-space>
          <n-button type="primary" :loading="subLoading" @click="formSubmit">确认添加</n-button>
          <n-button @click="handleReset">重置</n-button>
          <n-button @click="closeDrawer">取消</n-button>
        </n-space>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<script lang="ts">
  import { defineComponent, reactive, ref, toRefs } from 'vue';
  import { FormItemRule, TreeSelectOption, useMessage } from 'naive-ui';
  import { QuestionCircleOutlined } from '@vicons/antd';
  import { EditMenu } from '@/api/system/menu';
  import { newState } from '@/views/permission/menu/model';

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

  export default defineComponent({
    name: 'CreateDrawer',
    components: {},
    props: {
      title: {
        type: String,
        default: '添加顶级菜单',
      },
      optionTreeData: {
        type: Object || Array,
        default: [],
      },
    },
    emits: ['loadData'],
    setup(_props, context) {
      const message = useMessage();
      const formRef: any = ref(null);
      const state = reactive<any>({
        width: 700,
        isDrawer: false,
        subLoading: false,
        formParams: newState(null),
        placement: 'right',
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
            if (state.formParams.type != 3 && !value) {
              callback(new Error('请输入路由地址'));
            }
          },
        },
      };

      function openDrawer(pid: number) {
        if (document.body.clientWidth < 700) {
          state.width = document.body.clientWidth;
        }
        state.isDrawer = true;
        state.formParams = newState(null);
        state.formParams.pid = pid;
        if (pid > 0) {
          state.formParams.type = 2;
        }
      }

      function closeDrawer() {
        state.isDrawer = false;
      }

      function formSubmit() {
        formRef.value.validate((errors) => {
          if (!errors) {
            state.subLoading = true;
            EditMenu({ ...state.formParams })
              .then(async (_res) => {
                state.subLoading = false;
                message.success('操作成功');
                handleReset();
                await context.emit('loadData');
                closeDrawer();
              })
              .catch((_e: Error) => {
                state.subLoading = false;
              });
          } else {
            message.error('请填写完整信息');
          }
        });
      }

      function handleReset() {
        formRef.value.restoreValidation();
        state.formParams = newState(null);
      }

      // 处理选项更新
      function handleUpdateValue(
        value: string | number | Array<string | number> | null,
        _option: TreeSelectOption | null | Array<TreeSelectOption | null>
      ) {
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
