<template>
  <n-scrollbar style="max-height: 360px">
    <n-list>
      <n-list-item v-for="(item, index) in list" :key="item.id" @click="handleRead(index)">
        <n-thing class="px-15px" :class="{ 'opacity-30': item.isRead }">
          <template #avatar>
            <n-avatar round v-if="item.senderAvatar" :size="28" :src="item.senderAvatar" />
            <n-icon-wrapper v-else :size="28" :border-radius="10">
              <n-icon :size="20" :component="getIcon(item)" />
            </n-icon-wrapper>
          </template>
          <template #header>
            <n-ellipsis :line-clamp="1">
              {{ item.title }}
              <template #tooltip>
                {{ item.title }}
              </template>
            </n-ellipsis>
          </template>
          <template v-if="item.tagTitle" #header-extra>
            <n-tag v-bind="item.tagProps" size="small">{{ item.tagTitle }}</n-tag>
          </template>
          <template #description>
            <div v-if="item.content" class="description-box">
              <span v-html="item.content" class="description-html"> </span>
            </div>

            <p>{{ item.createdAt }}</p>
          </template>
        </n-thing>
      </n-list-item>
    </n-list>
  </n-scrollbar>
</template>
<script lang="ts" setup>
  import { MessageRow, getIcon } from '@/enums/systemMessageEnum';
  interface Props {
    list?: MessageRow[];
  }

  withDefaults(defineProps<Props>(), {
    list: () => [],
  });

  interface Emits {
    (e: 'read', val: number): void;
  }

  const emit = defineEmits<Emits>();

  function handleRead(index: number) {
    emit('read', index);
  }
</script>

<style lang="less" scoped>
  :deep(.description-box) {
    height: 100%;
    display: flex;
    align-items: center;
    margin-right: 10px;
  }
  :deep(.description-html) {
    height: 100%;
  }
  :deep(.px-15px) {
    padding-left: 15px;
    padding-right: 15px;
  }

  :deep(.text-34px) {
    font-size: 34px;
  }
</style>
