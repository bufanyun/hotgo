<template>
  <n-modal
    v-model:show="showModal"
    preset="card"
    title="预览"
    :bordered="false"
    :style="{ width: '520px' }"
  >
    <n-image preview-disabled :src="previewUrl" :on-error="errorImg" />
  </n-modal>
</template>

<script setup lang="ts">
  import { errorImg } from '@/utils/hotgo';
  import { ref } from 'vue';
  const showModal = ref(false);
  const previewUrl = ref('');

  function openPreview(url) {
    showModal.value = true;
    previewUrl.value = url;
  }

  defineExpose({
    openPreview,
  });
</script>

<style scoped lang="less"></style>
