<template>
  <div>
    <n-form
      :model="formParams"
      :rules="rules"
      ref="formRef"
      label-placement="left"
      :label-width="100"
      class="py-4"
    >
      <n-grid cols="2 300:1 600:2">
        <n-gi>
          <n-form-item label="菜单类型" path="type">
            <n-radio-group
              v-model:value="formParams.type"
              name="type"
              :on-update:value="handleUpdateType"
            >
              <n-radio-button
                v-for="menuType in dict.getOptionUnRef('sys_menu_types')"
                :key="menuType.value"
                :value="menuType.value"
                :label="menuType.label"
              />
            </n-radio-group>
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item label="上级目录" path="pid">
            <n-cascader
              clearable
              filterable
              :options="filterTreeOption"
              v-model:value="formParams.pid"
              value-field="key"
              label-field="label"
            />
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item :label="titleLabel" path="title">
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
                对应路由配置文件中 `name` 只能是唯一性，配置 `http(s)://` 开头地址 则会新窗口打开
              </n-tooltip>
              路由别名
            </template>
          </n-form-item>
        </n-gi>
      </n-grid>

      <n-grid cols="2 300:1 600:2" v-if="formParams.type !== 3">
        <n-gi v-if="formParams.type === 1">
          <n-form-item label="目录组件" path="component">
            <n-select
              v-if="formParams.type === 1"
              :options="dict.getOptionUnRef('sys_menu_component')"
              v-model:value="formParams.component"
              placeholder="请选择目录组件"
              clearable
              tag
            />
          </n-form-item>
        </n-gi>

        <n-gi v-if="formParams.type === 2">
          <n-form-item label="组件路径" path="component">
            <n-input placeholder="组件路径" v-model:value="formParams.component" />
            <template #feedback> 填 Vue 组件路径，如：`/system/menu/menu` </template>
          </n-form-item>
        </n-gi>
        <n-gi v-if="formParams.type === 1">
          <n-form-item label="默认跳转" path="redirect">
            <n-input placeholder="默认路由跳转地址" v-model:value="formParams.redirect" />
            <template #feedback>当目录下存在多个同级菜单时适用，如：`/system/menu/menu`</template>
          </n-form-item>
        </n-gi>
      </n-grid>

      <n-divider title-placement="left">功能设置</n-divider>

      <n-grid cols="1">
        <n-gi>
          <n-form-item label="分配权限" path="permissions">
            <n-dynamic-tags
              v-model:value="permissions"
              :on-update:value="handleUpdatePermissions"
              type="success"
            />
            <template #label>
              <n-tooltip trigger="hover">
                <template #trigger>
                  <n-icon :component="QuestionCircleOutlined" :size="18" :depth="3" />
                </template>
                请填写API路径地址，可同时作用于server端接口鉴权和web端细粒度权限
              </n-tooltip>
              分配权限
            </template>
          </n-form-item>
        </n-gi>
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
        <n-gi v-if="formParams.type !== 3">
          <n-form-item label="内嵌链接" path="frameSrc">
            <n-input-group>
              <n-select
                :style="{ width: '33%', minWidth: '80px' }"
                :options="dict.getOptionUnRef('sys_switch')"
                v-model:value="formParams.isFrame"
              />
              <n-input
                placeholder="格式：http://www.xxx.cn 或 https://www.xxx.cn"
                v-model:value="formParams.frameSrc"
              />
            </n-input-group>
          </n-form-item>
        </n-gi>
        <n-gi>
          <n-form-item label="菜单状态" path="status">
            <n-radio-group v-model:value="formParams.status" name="status">
              <n-radio-button
                v-for="status in dict.getOptionUnRef('sys_normal_disable')"
                :key="status.value"
                :value="status.value"
                :label="status.label"
              />
            </n-radio-group>
          </n-form-item>
        </n-gi>
      </n-grid>

      <n-form-item label="高级选项" path="senior" v-if="formParams.type !== 3">
        <n-checkbox-group v-model:value="senior" :on-update:value="handleSeniorChecked">
          <n-space item-style="display: flex;">
            <n-checkbox value="keepAlive" label="缓存路由" />
            <n-checkbox value="hidden" label="隐藏菜单" />
            <n-checkbox value="alwaysShow" label="简化路由" />
            <n-checkbox value="affix" label="页签固定" />
            <n-checkbox value="isRoot" label="根路由" />
          </n-space>
        </n-checkbox-group>
      </n-form-item>

      <n-form-item path="auth" style="margin-left: 100px" v-if="formParams.id > 0">
        <n-space>
          <n-button type="primary" :loading="loading" @click="formSubmit">保存修改 </n-button>
          <n-button @click="handleReset">重置</n-button>
          <n-button @click="handleDel">删除</n-button>
        </n-space>
      </n-form-item>
    </n-form>
  </div>
</template>

<script setup lang="ts">
  import { QuestionCircleOutlined } from '@vicons/antd';
  import IconSelector from '@/components/IconSelector/index.vue';
  import { computed, ref } from 'vue';
  import { newState, State } from '@/views/permission/menu/model';
  import { FormItemRule, useDialog, useMessage } from 'naive-ui';
  import { cloneDeep } from 'lodash-es';
  import { DeleteMenu, EditMenu } from '@/api/system/menu';
  import { findTreeNode } from '@/utils';
  import { useDictStore } from '@/store/modules/dict';
  import type { FormRules } from 'naive-ui/es/form/src/interface';

  const rules: FormRules = {
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
      validator: function (_rule: FormItemRule, value: any, callback: Function): boolean | Error {
        if (formParams.value.type != 3 && !value) {
          callback(new Error('请输入路由地址'));
        }
        return true;
      },
    },
  };

  const emit = defineEmits(['reloadTable', 'closeForm']);
  const dict = useDictStore();
  const message = useMessage();
  const dialog = useDialog();
  const loading = ref(false);
  const formRef = ref();
  const formParams = defineModel<State>('formParams', { required: true });
  const treeOption = defineModel<any[]>('treeOption');
  const titleLabel = computed(() => {
    const type = formParams.value.type as number;
    if (type == 1) {
      return '目录名称';
    }
    if (type == 2) {
      return '菜单名称';
    }
    return '按钮名称';
  });

  const filterTreeOption = computed(() => {
    const list = cloneDeep(treeOption.value) as any[];
    for (let i = 0; i < list.length; i++) {
      const item = list[i];
      if (item.id === formParams.value.id) {
        item.disabled = true;
        if (item.children) {
          setChildrenDisabled(item.children);
        }
      } else {
        if (item.children) {
          const foundChild = findItemById(item.children, formParams.value.id);
          if (foundChild) {
            foundChild.disabled = true;
            setChildrenDisabled(foundChild.children);
          }
        }
      }
    }
    return list;
  });

  const senior = computed(() => {
    let opts: string[] = [];
    if (formParams.value.keepAlive == 1) {
      opts.push('keepAlive');
    }
    if (formParams.value.hidden == 1) {
      opts.push('hidden');
    }
    if (formParams.value.alwaysShow == 1) {
      opts.push('alwaysShow');
    }
    if (formParams.value.affix == 1) {
      opts.push('affix');
    }
    if (formParams.value.isRoot == 1) {
      opts.push('isRoot');
    }
    return opts;
  });

  function handleSeniorChecked(_: string, opt: { actionType: 'check' | 'uncheck'; value: string }) {
    if (opt.actionType == 'check') {
      formParams.value[opt.value] = 1;
    } else {
      formParams.value[opt.value] = 0;
    }
  }

  function setChildrenDisabled(children: State[]) {
    if (!children) return;
    for (let i = 0; i < children.length; i++) {
      const child = children[i];
      child.disabled = true;
      if (child.children) {
        setChildrenDisabled(child.children);
      }
    }
  }

  function findItemById(children: State[], id: number) {
    if (!children) {
      return null;
    }
    for (let i = 0; i < children.length; i++) {
      const item = children[i];
      if (item.id === id) {
        return item;
      }

      if (item.children) {
        const child = findItemById(item.children, id);
        if (child) {
          return child;
        }
      }
    }
    return null;
  }

  function getFormLoading() {
    return loading.value;
  }

  function handleDel() {
    dialog.warning({
      title: '提示',
      content: `您确定要删除 ` + formParams.value.title + ` 菜单吗?`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        DeleteMenu(formParams.value).then((_res) => {
          message.success('操作成功');
          formParams.value = newState(null);
          emit('reloadTable');
        });
      },
    });
  }

  function handleReset() {
    const item = findTreeNode(treeOption.value, formParams.value.id);
    formParams.value = newState(item);
  }

  function formSubmit() {
    formRef.value.validate((errors: boolean) => {
      if (!errors) {
        loading.value = true;
        EditMenu(formParams.value)
          .then((_res) => {
            message.success('操作成功');
            emit('reloadTable');
            emit('closeForm');
          })
          .finally(() => {
            loading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  const permissions = computed(() => {
    if (formParams.value.permissions.length == 0) {
      return [];
    }
    return formParams.value.permissions.split(',');
  });

  function handleUpdatePermissions(values: string[]) {
    formParams.value.permissions = Array.from(new Set(values)).join(',');
  }

  function handleUpdateType(value: number) {
    formParams.value.type = value;
    if (value == 1) {
      formParams.value.component = null;
    } else {
      formParams.value.component = '';
    }

    const item = findTreeNode(treeOption.value, formParams.value.id);
    if (item && item.type == value) {
      formParams.value.component = item.component;
    }
  }

  defineExpose({
    handleReset,
    formSubmit,
    getFormLoading,
  });
</script>

<style scoped lang="less"></style>
