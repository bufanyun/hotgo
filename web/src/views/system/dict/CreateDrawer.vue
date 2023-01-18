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
        <n-form-item label="上级字典" path="pid">
          <n-tree-select
            :options="optionTreeData"
            :default-value="formParams.pid"
            @update:value="handleUpdateValue"
          />
        </n-form-item>
        <n-form-item label="类型名称" path="name">
          <n-input placeholder="请输入类型名称" v-model:value="formParams.name" />
        </n-form-item>
        <n-form-item label="类型编码" path="type">
          <n-input placeholder="请输入类型编码" v-model:value="formParams.type" />
        </n-form-item>
        <n-form-item label="排序" path="sort">
          <n-input-number v-model:value="formParams.sort" clearable />
        </n-form-item>
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
  import { EditDict } from '@/api/dict/dict';

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
        type: Object || Array,
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
        type: '',
        name: '',
        remark: '',
        status: 1,
        sort: 10,
      });

      const state = reactive<any>({
        width: 500,
        isDrawer: false,
        subLoading: false,
        formParams: defaultValueRef(),
        placement: 'right',
        icon: '',
        alertText:
          '该功能主要实时预览各种布局效果，更多完整配置在 projectSetting.ts 中设置，建议在生产环境关闭该布局预览功能。',
      });

      function openDrawer(form) {
        if (document.body.clientWidth < 500) {
          state.width = document.body.clientWidth;
        }
        state.isDrawer = true;
        state.formParams = Object.assign(state.formParams, form);
      }

      function closeDrawer() {
        state.isDrawer = false;
      }

      function formSubmit() {
        formRef.value.validate((errors) => {
          if (!errors) {
            EditDict({ ...state.formParams }).then(async (_res) => {
              message.success('操作成功');
              handleReset();
              await context.emit('loadData');
              closeDrawer();
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
        statusMap,
        handleUpdateValue,
        QuestionCircleOutlined,
      };
    },
  });
</script>
