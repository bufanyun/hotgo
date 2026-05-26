<template>
  <n-drawer
    v-model:show="isDrawer"
    :width="dialogWidth"
    :label-placement="settingStore.isMobile ? 'top' : 'left'"
  >
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
              v-for="status in dict.getOptionUnRef('sys_normal_disable')"
              :key="status.value"
              :value="status.value"
              :label="status.label"
            />
          </n-radio-group>
        </n-form-item>
      </n-form>

      <template #footer>
        <n-space>
          <n-button type="primary" :loading="btnLoading" @click="formSubmit">提交</n-button>
          <n-button @click="closeDrawer">取消</n-button>
        </n-space>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { useMessage } from 'naive-ui';
  import { EditDict } from '@/api/dict/dict';
  import { useDictStore } from '@/store/modules/dict';
  import type { FormRules } from 'naive-ui/es/form/src/interface';
  import { State } from '@/views/system/dict/model';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';

  const rules: FormRules = {
    name: {
      required: true,
      message: '请输入类型名称',
      trigger: 'blur',
    },
    type: {
      required: true,
      message: '请输入类型编码',
      trigger: 'blur',
    },
  };

  interface Props {
    title: '添加顶级菜单';
    optionTreeData: [];
  }

  const emit = defineEmits(['loadData']);
  const props = withDefaults(defineProps<Props>(), {});
  const dict = useDictStore();
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const formRef: any = ref(null);
  const formParams = ref<State>(new State());
  const btnLoading = ref(false);
  const isDrawer = ref(false);

  const dialogWidth = computed(() => {
    return adaModalWidth(500);
  });

  function openDrawer(state: State) {
    isDrawer.value = true;
    formParams.value = state;
  }

  function closeDrawer() {
    isDrawer.value = false;
  }

  function formSubmit() {
    formRef.value.validate((errors) => {
      if (!errors) {
        btnLoading.value = true;
        EditDict(formParams.value)
          .then(async (_res) => {
            message.success('操作成功');
            emit('loadData');
            closeDrawer();
          })
          .finally(() => {
            btnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 处理选项更新
  function handleUpdateValue(value: number) {
    formParams.value.pid = value;
  }

  defineExpose({
    openDrawer,
  });
</script>
