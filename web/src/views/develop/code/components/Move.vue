<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="移动字段"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-card
          :bordered="false"
          :content-style="{ padding: '0px' }"
          :header-style="{ padding: 'px' }"
          :segmented="true"
        >
          请通过拖拽来移动字段的位置。
          <div class="mt-8"></div>
          <Draggable
            class="draggable-ul"
            animation="300"
            :list="columns"
            group="people"
            itemKey="name"
          >
            <template #item="{ element }">
              <div class="cursor-move draggable-li">
                <n-tag type="default" size="small" style="font-weight: 800">{{
                  element.name
                }}</n-tag
                ><span class="ml-2">{{ element.dc }}</span>
              </div>
            </template>
          </Draggable>
        </n-card>
      </n-scrollbar>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import Draggable from 'vuedraggable';
  import { adaModalWidth } from '@/utils/hotgo';

  const showModal = ref(false);
  const columns = defineModel<[]>('columns');
  const dialogWidth = computed(() => {
    return adaModalWidth(360);
  });

  function openModal() {
    showModal.value = true;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped>
  .draggable-ul {
    width: 100%;
    overflow: hidden;
    margin-top: -8px;

    .draggable-li {
      width: 100%;
      padding: 8px 4px;
      color: #333;
      border-bottom: 1px solid #efeff5;
    }

    .draggable-li:hover {
      background-color: rgba(229, 231, 235, var(--tw-border-opacity));
    }
  }
</style>
