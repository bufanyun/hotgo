<template>
  <n-tabs v-model:value="currentTab" type="line" justify-content="space-evenly">
    <n-tab-pane
      v-for="(item, index) in notificationStore.getMessages"
      :key="item.key"
      :name="index"
    >
      <template #tab>
        <div>
          <span>{{ item.name }}</span>
          <n-badge
            v-bind="item.badgeProps"
            :value="item.list.filter((message) => !message.isRead).length"
            :max="99"
            show-zero
          />
        </div>
      </template>
      <n-spin :show="loading">
        <n-empty v-show="item.list.length === 0" description="无数据" :show-icon="false">
          <template #extra>
            <n-button size="small" @click="handleLoadMore"> 查看更多</n-button>
          </template>
        </n-empty>

        <message-list :list="item.list" @read="handleRead" />
      </n-spin>
    </n-tab-pane>
  </n-tabs>
  <n-space v-if="showAction" justify="center" size="large" class="flex border-t">
    <n-button class="act-btn" size="small" @click="handleClear">清空</n-button>
    <n-button class="act-btn" size="small" @click="handleAllRead">全部已读</n-button>
    <n-button class="act-btn" size="small" @click="handleLoadMore">查看更多</n-button>
  </n-space>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import MessageList from './MessageList.vue';
  import { notificationStoreWidthOut } from '@/store/modules/notification';
  import { ReadAll, UpRead } from '@/api/apply/notice';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const notificationStore = notificationStoreWidthOut();
  const loading = ref(false);
  const currentTab = ref(0);
  const showAction = computed(
    () => notificationStore.getMessages[currentTab.value].list.length > 0
  );

  function handleRead(index: number) {
    loading.value = true;
    const message = notificationStore.getMessages[currentTab.value].list[index];
    UpRead({ id: message.id })
      .then(() => {
        message.isRead = true;
        if (!message.isRead) {
          switch (message.type) {
            case 1:
              notificationStore.notifyUnread--;
              break;
            case 2:
              notificationStore.noticeUnread--;
              break;
            case 3:
              notificationStore.letterUnread--;
              break;
          }
        }
      })
      .finally(() => {
        loading.value = false;
      });
  }

  function handleAllRead() {
    loading.value = true;
    ReadAll({ type: notificationStore.getMessages[currentTab.value].key })
      .then(() => {
        notificationStore.getMessages[currentTab.value].list.forEach((item) =>
          Object.assign(item, { isRead: true })
        );
        switch (notificationStore.getMessages[currentTab.value].key) {
          case 1:
            notificationStore.notifyUnread = 0;
            break;
          case 2:
            notificationStore.noticeUnread = 0;
            break;
          case 3:
            notificationStore.letterUnread = 0;
            break;
        }
      })
      .finally(() => {
        loading.value = false;
      });
  }

  function handleClear() {
    notificationStore.getMessages[currentTab.value].list = [];
    switch (notificationStore.getMessages[currentTab.value].key) {
      case 1:
        notificationStore.notifyUnread = 0;
        break;
      case 2:
        notificationStore.noticeUnread = 0;
        break;
      case 3:
        notificationStore.letterUnread = 0;
        break;
    }
  }

  function handleLoadMore() {
    router.push({
      name: 'home_message',
      query: {
        type: notificationStore.getMessages[currentTab.value].key,
      },
    });
  }
</script>
<style scoped>
  .act-btn {
    margin-top: 8px;
  }
</style>
