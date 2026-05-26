<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="添加菜单"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <EditForm
            ref="editFormRef"
            v-model:formParams="formParams"
            v-model:treeOption="treeOption"
            @reloadTable="reloadTable"
            @closeForm="closeForm"
          />
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm"> 关闭 </n-button>
          <n-button @click="handleReset"> 重置 </n-button>
          <n-button type="primary" :loading="formLoading" @click="formSubmit"> 确定添加 </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';
  import EditForm from '@/views/permission/menu/editForm.vue';

  const emit = defineEmits(['reloadTable']);
  const formParams = ref<State>(newState(null));
  const treeOption = defineModel<any[]>('treeOption');
  const loading = ref(false);
  const showModal = ref(false);
  const editFormRef = ref();
  const formLoading = computed(() => {
    return editFormRef.value?.getFormLoading();
  });

  const dialogWidth = computed(() => {
    return adaModalWidth(960);
  });

  function reloadTable() {
    emit('reloadTable');
  }

  function openModal(state: State) {
    showModal.value = true;
    formParams.value = newState(state);
  }

  function formSubmit(e) {
    e.preventDefault();
    editFormRef.value.formSubmit();
  }

  function handleReset() {
    editFormRef.value.handleReset();
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
