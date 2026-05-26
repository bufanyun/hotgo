<template>
  <div>
    <n-modal
      v-model:show="isShowModal"
      :style="{
        width: dialogWidth,
      }"
      :show-icon="false"
      preset="dialog"
      title="链接转图片"
    >
      <n-alert type="info"> 将外部图片链接下载并转存至平台的存储驱动中 </n-alert>
      <n-form
        :model="formParams"
        ref="formPacketRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-spin :show="loading" description="请稍候...">
          <n-form-item label="图片链接" path="url">
            <n-input v-model:value="formParams.url" />
          </n-form-item>

          <n-form-item label="预览图片" v-if="formParams.url != ''">
            <n-image width="150" :src="formParams.url" />
          </n-form-item>
        </n-spin>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="closeForm">关闭</n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmForm">上传 </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { adaModalWidth } from '@/utils/hotgo';
  import { useMessage } from 'naive-ui';
  import { ImageTransferStorage } from '@/api/apply/attachment';

  const emit = defineEmits(['reloadTable']);
  const loading = ref(false);
  const isShowModal = ref(false);
  const dialogWidth = ref(adaModalWidth(640));
  const formBtnLoading = ref(false);
  const formPacketRef = ref();
  const message = useMessage();
  const formParams = ref({
    url: '',
  });

  function reloadTable() {
    emit('reloadTable');
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    ImageTransferStorage(formParams.value)
      .then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          reloadTable();
          closeForm();
        });
      })
      .finally(() => {
        formBtnLoading.value = false;
      });
  }

  function closeForm() {
    isShowModal.value = false;
  }

  function showModal() {
    isShowModal.value = true;
    formParams.value.url = '';
  }

  defineExpose({ showModal });
</script>

<style lang="less"></style>
