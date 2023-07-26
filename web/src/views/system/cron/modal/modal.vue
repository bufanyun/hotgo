<template>
  <div>
    <n-modal
      v-model:show="showModal"
      style="width: 80%; height: 700px"
      :show-icon="false"
      preset="dialog"
      :title="title"
    >
      <Index @reloadGroupOption="reloadGroupOption" />
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import Index from './index.vue';

  const emit = defineEmits(['reloadGroupOption']);
  const showModal = ref(false);
  const title = ref('管理分组');

  function openDrawer() {
    showModal.value = true;
  }

  function reloadGroupOption() {
    emit('reloadGroupOption');
  }

  defineExpose({ openDrawer });
</script>

<style scoped></style>
