<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :on-after-leave="cancelForm"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      :title="formParams?.id > 0 ? '编辑 #' + formParams?.id : '添加'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="100"
        class="py-8"
      >
        <n-grid x-gap="24" :cols="1">
          <n-gi>
            <n-form-item label="名称" path="name">
              <n-input placeholder="请输入名称" v-model:value="formParams.name" />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="上传图片" path="cover">
              <UploadImage v-model:value="formParams.cover" />
            </n-form-item>
          </n-gi>
          <n-gi>
            <n-form-item label="链接地址" path="link">
              <n-input placeholder="请输入链接地址" v-model:value="formParams.link" />
            </n-form-item>
          </n-gi>
        </n-grid>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="cancelForm">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { h, reactive, ref, computed, defineExpose, defineEmits } from 'vue';
import { useDialog, useMessage } from 'naive-ui';
import UploadImage from '@/components/Upload/uploadImage.vue';
import { rules } from './model';
import { Edit, Add } from '@/api/addons/flashbanner/index';
import { adaModalWidth } from '@/utils/hotgo';
import { cloneDeep } from 'lodash-es';

const emit = defineEmits(['on-refresh']);

const defaultState = {
  name: '',
  cover: '',
  link: '',
};

const message = useMessage();
const showModal = ref(false);
const formBtnLoading = ref(false);
const formRef = ref<any>();
const formParams = ref<any>(cloneDeep(defaultState));

const dialogWidth = computed(() => {
  return adaModalWidth();
});

// 关闭表单
const cancelForm = () => {
  showModal.value = false;
  formParams.value = cloneDeep(defaultState);
};

// 新增或编辑
function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  const Request = formParams.value.id > 0 ? Edit : Add;
  formRef.value.validate((errors) => {
    if (!errors) {
      Request(formParams.value).then((_res) => {
        message.success('操作成功');
        cancelForm();
        emit('on-refresh');
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}

defineExpose({ showModal, formParams });
</script>

<style lang="less" scoped></style>
