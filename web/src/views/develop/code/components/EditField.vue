<template>
  <basicModal
    @register="modalRegister"
    ref="modalRef"
    class="basicModal"
    @on-ok="okModal"
    @on-close="close"
    :style="{
      width: dialogWidth,
    }"
  >
    <template #default>
      <n-form
        ref="formRef"
        inline
        :label-width="90"
        :model="formValue"
        size="medium"
        label-placement="left"
      >
        <n-tabs type="line" animated>
          <n-tab-pane name="basic" tab="基本信息">
            <n-grid x-gap="24" :cols="2">
              <n-gi :span="1">
                <n-form-item label="字段描述" path="dc">
                  <n-input v-model:value="formValue.dc" placeholder="字段描述" />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="字段列名" path="name">
                  <n-input disabled v-model:value="formValue.name" placeholder="字段列名" />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="表单组件" path="formMode">
                  <n-select
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.formMode"
                    :options="getFormModeOptions(formValue.tsType)"
                    placeholder="表单组件"
                  />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="默认值" path="defaultValue">
                  <n-input
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.defaultValue"
                    placeholder="默认值"
                  />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="显示状态" path="listShow">
                  <n-select
                    v-model:value="formValue.listShow"
                    :options="props.selectList?.listShow ?? []"
                    placeholder="显示状态"
                  />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="字典类型" path="dictType">
                  <n-treeSelect
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.dictType"
                    :options="props.selectList?.dictMode ?? []"
                    placeholder="字典类型"
                    clearable
                  />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="单元格宽度" path="width">
                  <n-input v-model:value="formValue.width" placeholder="单元格宽度" />
                </n-form-item>
              </n-gi>
              <n-gi :span="1">
                <n-form-item label="表单验证" path="formRole">
                  <n-select
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.formRole"
                    :options="props.selectList?.formRole ?? []"
                    placeholder="表单验证"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
            <n-grid x-gap="12" :cols="7">
              <n-gi>
                <n-form-item label="是否编辑" path="isEdit">
                  <n-checkbox
                    :disabled="formValue.name === 'id'"
                    v-model:checked="formValue.isEdit"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="是否必填" path="required">
                  <n-checkbox
                    :disabled="formValue.name === 'id'"
                    v-model:checked="formValue.required"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="是否唯一" path="unique">
                  <n-checkbox
                    :disabled="formValue.name === 'id'"
                    v-model:checked="formValue.unique"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="列表显示" path="isList">
                  <n-checkbox v-model:checked="formValue.isList" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="字段导出" path="isExport">
                  <n-checkbox v-model:checked="formValue.isExport" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="字段查询" path="isQuery">
                  <n-checkbox v-model:checked="formValue.isQuery" />
                </n-form-item>
              </n-gi>
            </n-grid>
            <n-grid x-gap="12" :cols="1">
              <n-gi :span="6">
                <n-form-item label="查询条件" path="queryWhere">
                  <n-select
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.queryWhere"
                    :options="props.selectList?.whereMode ?? []"
                    placeholder="查询条件"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-tab-pane>
          <n-tab-pane name="expand" tab="拓展信息">
            <n-grid x-gap="24" :cols="1">
              <n-gi :span="1">
                <n-form-item label="文本说明" path="placeholder">
                  <n-input
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.placeholder"
                    placeholder="输入placeholder，占位符"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
            <n-grid x-gap="24" :cols="1">
              <n-gi :span="1">
                <n-form-item label="显示条件" path="showCondition">
                  <n-input
                    :disabled="formValue.name === 'id'"
                    v-model:value="formValue.showCondition"
                    placeholder="输入框显示条件 如 params.status == 1"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
            <n-grid x-gap="24" :cols="1">
              <n-gi :span="1">
                <n-form-item label="其他属性" path="attribute">
                  <n-checkbox-group v-model:value="formValue.attribute">
                    <n-checkbox value="sort" label="表头排序" />
                  </n-checkbox-group>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-tab-pane>
        </n-tabs>
      </n-form>
    </template>
  </basicModal>
</template>

<script lang="ts" setup>
  import { ref, computed, watch } from 'vue';
  import { basicModal, useModal } from '@/components/Modal';
  import { cloneDeep } from 'lodash-es';
  import { selectListObj } from '@/views/develop/code/components/model';

  const dialogWidth = ref('40%');
  const emit = defineEmits(['updateShowModal', 'updateFieldInfo']);

  interface Props {
    showModal: boolean;
    fieldInfo?: object;
    selectList: any;
  }

  const props = withDefaults(defineProps<Props>(), {
    showModal: false,
    fieldInfo: () => {
      return {};
    },
    selectList: selectListObj,
  });

  const isShowModal = computed({
    get: () => {
      return props.showModal;
    },
    set: (value) => {
      emit('updateShowModal', value);
    },
  });

  const formValue = ref<any>(cloneDeep(props.fieldInfo));

  const [modalRegister, { openModal, closeModal }] = useModal({
    title: '更新字段',
  });

  async function okModal() {
    emit('updateFieldInfo', formValue);
    closeModal();
  }

  function showModal() {
    openModal();
  }

  function close() {
    isShowModal.value = false;
  }

  watch(
    () => props.showModal,
    (value) => {
      if (isShowModal.value) {
        formValue.value = cloneDeep(props.fieldInfo);
        showModal();
      }
    }
  );

  function getFormModeOptions(type: string) {
    const options = cloneDeep(props.selectList.formMode ?? []);
    if (options.length === 0) {
      return [];
    }
    switch (type) {
      case 'number':
        for (let i = 0; i < options.length; i++) {
          const allows = ['InputNumber', 'Radio', 'Select', 'Switch', 'Rate'];
          if (!allows.includes(options[i].value)) {
            options[i].disabled = true;
          }
        }
        break;
      default:
    }
    return options;
  }
</script>

<style lang="less" scoped></style>
