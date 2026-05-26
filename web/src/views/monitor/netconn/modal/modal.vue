<template>
  <div>
    <n-modal
      v-model:show="showModal"
      style="width: 80%; height: 90%; min-height: 820px"
      :show-icon="false"
      preset="dialog"
      :title="title"
    >
      <Index />
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import Index from './index.vue';

  const showModal = ref(false);
  const title = ref('许可证列表');

  function openDrawer() {
    showModal.value = true;
  }

  defineExpose({ openDrawer });
</script>

<style scoped></style>
