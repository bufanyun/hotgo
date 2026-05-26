<template>
  <div>
    <n-modal
      title="设置选项字段"
      v-model:show="showModal"
      :block-scroll="false"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
    >
      <n-form
        :model="formValue"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="100"
        class="py-4"
      >
        <n-form-item label="选项值" path="valueColumn">
          <n-select
            filterable
            tag
            placeholder="请选择"
            :options="columnsOption"
            v-model:value="formValue.valueColumn"
          />
        </n-form-item>

        <n-form-item label="选项名称" path="labelColumn">
          <n-select
            filterable
            tag
            placeholder="请选择"
            :options="columnsOption"
            v-model:value="formValue.labelColumn"
          />
        </n-form-item>
      </n-form>
      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="info" @click="confirmForm">保存</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { rules } from '@/views/addons/stock/itemBrand/model';
  import { cloneDeep } from 'lodash-es';
  import { useMessage } from 'naive-ui';
  import { Edit } from '@/api/addons/stock/itemClass';

  interface Props {
    columnsOption: any;
  }

  const props = withDefaults(defineProps<Props>(), {
    columnsOption: [],
  });

  const rules = {
    valueColumn: {
      required: true,
      trigger: ['blur', 'input'],
      type: 'string',
      message: '选项值不能为空',
    },
    labelColumn: {
      required: true,
      trigger: ['blur', 'input'],
      type: 'string',
      message: '选项名称不能为空',
    },
  };

  const message = useMessage();
  const emit = defineEmits(['update']);
  const showModal = ref(false);
  const formRef = ref();
  const formValue = ref({
    valueColumn: null,
    labelColumn: null,
  });

  function openModal(state) {
    showModal.value = true;
    if (!state) {
      state = {
        valueColumn: null,
        labelColumn: null,
      };
    }

    if (!state.valueColumn) {
      const item = props.columnsOption.find((item) => item.value === 'id');
      if (item) {
        state.valueColumn = item.value;
      }
    }

    if (!state.labelColumn) {
      const item = props.columnsOption.find((item) => item.value === 'title' || item.value === 'name');
      if (item) {
        state.labelColumn = item.value;
      }
    }
    formValue.value = cloneDeep(state);
    emit('update', formValue.value);
  }

  function closeForm() {
    showModal.value = false;
  }

  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        emit('update', formValue.value);
        closeForm();
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  defineExpose({
    openModal,
  });
</script>
