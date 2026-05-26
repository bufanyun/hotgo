<template>
  <div>
    <n-modal
      title="菜单管理"
      v-model:show="showModal"
      :block-scroll="false"
      :mask-closable="false"
      :show-icon="false"
      :on-after-leave="handleCancel"
      preset="card"
      :style="{
        width: dialogWidth,
      }"
    >
      <Menu v-model:isModal="isModal" />
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import Menu from './menu.vue';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const showModal = ref(false);
  const isModal = ref(true);
  const dialogWidth = computed(() => {
    return adaModalWidth(1280);
  });

  function openModal() {
    showModal.value = true;
  }

  function handleCancel() {
    emit('reloadTable');
  }

  defineExpose({
    openModal,
  });
</script>
